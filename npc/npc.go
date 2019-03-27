package npc

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"git.unistra.fr/AOEINT/server/carte"
	"git.unistra.fr/AOEINT/server/constants"
	"git.unistra.fr/AOEINT/server/data"
	"git.unistra.fr/AOEINT/server/ressource"
	"git.unistra.fr/AOEINT/server/utils"
)

//Npc :
type Npc struct {
	x                int
	y                int
	px               float64
	py               float64
	pv               int
	vitesse          int
	vue              int
	portee           int
	offensive        bool //true=soldier else harvester
	size             int
	damage           int
	selectable       bool //false=villager
	typ              int  // 0:villager, 1:harvester, 2:soldier
	TeamFlag         int
	ressourceChannel chan []int
	hasOrder         bool //Si un déplacement a dejà été demandé par le joueur (disable auto movement)
	// isMoving *safeNumber
	PlayerUUID string
}

type safeNumber struct {
	val bool
	m   sync.Mutex
}

//New : new NPC
func New(x int, y int, pv int, vitesse int, vue int, portee int, offensive bool, size int, damage int, selectable bool, typ int, flag int, channel *chan []int) Npc {
	pnj := Npc{x, y, 0, 0, pv, vitesse, vue, portee, offensive, size, damage, selectable, typ, flag, *channel, false, ""}
	return pnj
}

//Create : generate a new NPC
func Create(class string, x int, y int, flag int, channel *chan []int) (Npc, string) {
	var pnj Npc
	switch class {
	case "soldier":
		pnj = New(x, y, constants.SoldierPv, constants.SoldierVitesse, constants.SoldierVue,
			constants.SoldierPortee, true, constants.SoldierSize, constants.SoldierDamage, true, 2, flag, channel)
	case "harvester":
		pnj = New(x, y, constants.HarvesterPv, constants.HarvesterVitesse, constants.HarvesterVue,
			constants.HarvesterVillPortee, false, constants.HarvesterSize, constants.HarvesterDamage, true, 1, flag, channel)
	default:
		pnj = New(x, y, constants.VillagerPv, constants.VillagerVitesse, constants.VillagerVue,
			constants.HarvesterVillPortee, false, constants.VillagerSize, constants.VillagerDamage, false, 0, flag, channel)
	}
	id := (&data.IDMap).AddObject(&pnj)
	//pnj.Transmit(id)
	return pnj, id
}

//Stringify : create a map[string]string of the main arguments of a NPC
func (pnj Npc) Stringify() map[string]string {
	res := make(map[string]string)
	res["pv"] = strconv.Itoa(pnj.pv)
	res["x"] = strconv.Itoa(pnj.x)
	res["y"] = strconv.Itoa(pnj.y)
	res["vitesse"] = strconv.Itoa(pnj.vitesse)
	res["type"] = strconv.Itoa(pnj.typ)
	res["damage"] = strconv.Itoa(pnj.damage)
	res["offensive"] = strconv.FormatBool(pnj.offensive)
	res["vue"] = strconv.Itoa(pnj.vue)
	res["portee"] = strconv.Itoa(pnj.portee)
	res["TeamFlag"] = strconv.Itoa(pnj.TeamFlag)
	res["PlayerUUID"] = pnj.PlayerUUID
	res["px"] = fmt.Sprintf("%f", pnj.px)
	res["py"] = fmt.Sprintf("%f", pnj.py)
	return res
}

//Transmit : add the npc to the communcation's buffer
func (pnj Npc) Transmit(id string) {
	arr := pnj.Stringify()
	for k, e := range arr {
		data.AddToAllAction(constants.ActionNewNpc, id, k, e)
	}
}

//GetX : return the position X
func (pnj Npc) GetX() int {
	return pnj.x
}

//GetY : return the position Y
func (pnj Npc) GetY() int {
	return pnj.y
}

//GetVue : return villager's vision
func (pnj Npc) GetVue() int {
	return pnj.vue
}

//GetType : return the villager's type
func (pnj Npc) GetType() int {
	return pnj.typ
}

//GetPv : return Pv
func (pnj Npc) GetPv() int {
	return pnj.pv
}

func (pnj *Npc) deplacement(path []carte.Case, wg *sync.WaitGroup) {
	if path != nil {
		ndep := len(path) - 1
		vdep := (1000000000 / pnj.vitesse)
		for i := 0; i <= ndep; i++ {
			time.Sleep(time.Duration(vdep))
			pnj.x = path[i].GetPathX()
			pnj.y = path[i].GetPathY()
		}
		if wg != nil {
			wg.Done()
		}
	}
}

//MoveTo : move a npc from his position(x,y) to another position(x,y)
func (pnj *Npc) MoveTo(c carte.Carte, destx int, desty int, wg *sync.WaitGroup) []carte.Case {
	var path []carte.Case
	path = nil
	if c.GetTile(destx, desty).GetType() == 0 {
		path = c.GetPathFromTo(pnj.x, pnj.y, destx, desty)
		go pnj.deplacement(path, wg)
	}
	return path
}

//Abs : utility function
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//GetSpeed : return the npc's speed
func (pnj Npc) GetSpeed() int {
	return pnj.vitesse
}

//RecoltePossible : return true if te villager can acces to a tile to harvest the resource in x, y
func RecoltePossible(c carte.Carte, x int, y int) bool {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if c.IsEmpty(i, j) {
				return true
			}
		}
	}
	return false
}

//MoveHarvest : (move to the nreast ressource in the villagers's vision)
func (pnj *Npc) MoveHarvest(c carte.Carte) {
	var i, j int
	var ress *ressource.Ressource
	distance := 2000
	if pnj.GetType() == 2 {
		utils.Debug("Un soldat ne peut pas recolter de ressources")
		return
	}
	for i = pnj.GetX() - pnj.GetVue(); i <= pnj.GetX()+pnj.GetVue() || i > c.GetSize(); i++ {
		if i < 0 {
			i = 0
		}
		for j = pnj.GetY() - pnj.GetVue(); j <= pnj.GetY()+pnj.GetVue() || j > c.GetSize(); j++ {
			if j < 0 {
				j = 0
			}
			if c.GetTile(i, j).GetType() == 2 {
				if c.GetTile(i, j).GetRess().GetType() == 2 && pnj.GetType() != 0 {
					utils.Debug("Seul un harvester peut recolter de la pierre")
					continue
				}
				if (Abs(i-pnj.GetX())+Abs(j-pnj.GetY())) < distance &&
					RecoltePossible(c, i, j) {
					distance = Abs(i-pnj.GetX()) + Abs(j-pnj.GetY())
					ress = c.GetTile(i, j).GetRess()
				}
			}
		}
	}

	// pas de ressources dans la vue du villageois
	if distance == 2000 {
		return
	}

	var posRecolteVillX, posRecolteVillY int
	distance = 2000

	for i = ress.GetX() - pnj.portee; i <= ress.GetX()+pnj.portee; i++ {
		for j = ress.GetY() - pnj.portee; j <= ress.GetY()+pnj.portee; j++ {
			if (Abs(i-pnj.GetX())+Abs(j-pnj.GetY())) < distance &&
				c.IsEmpty(i, j) {
				distance = Abs(i-pnj.GetX()) + Abs(j-pnj.GetY())
				posRecolteVillX = i
				posRecolteVillY = j
			}
		}
	}
	// pas d'accès possible pour recolter la ressource
	if distance == 2000 {
		return
	}
	// on attends que le villageois ait finit son déplacement
	var wg sync.WaitGroup
	wg.Add(1)
	go pnj.MoveTo(c, posRecolteVillX, posRecolteVillY, &wg)
	wg.Wait()

	// Le villageois se trouve bien à l'emplacement de la recolte?
	if pnj.GetX() == (posRecolteVillX) && pnj.GetY() == posRecolteVillY {
		go (pnj).Harvest(c, ress, posRecolteVillX, posRecolteVillY)
	}
}

/*MoveHarvestTarget : (move to the nreast ressource in the villagers's vision).
Triggered when a villager is inactive, cancelled when the player moves the npc
*/
func (pnj *Npc) MoveHarvestTarget(c carte.Carte, ress *ressource.Ressource) {
	var i, j int
	//Verify the parameters
	if pnj.GetType() == 2 {
		log.Print("Un soldat ne peut pas recolter de ressources")
		return
	}
	if ress.GetType() == 2 && pnj.GetType() != 0 {
		log.Print("Seul un harvester peut recolter de la pierre")
		return
	}
	if pnj.GetVue() < (Abs(ress.GetX()-pnj.GetX()) + Abs(ress.GetY()-pnj.GetY())) {
		log.Print("La ressource n'est pas dans la vue du npc")
		return
	}

	var posRecolteVillX, posRecolteVillY int
	distance := 2000

	for i = ress.GetX() - pnj.portee; i <= ress.GetX()+pnj.portee; i++ {
		for j = ress.GetY() - pnj.portee; j <= ress.GetY()+pnj.portee; j++ {
			if (Abs(i-pnj.GetX())+Abs(j-pnj.GetY())) < distance &&
				c.IsEmpty(i, j) {
				distance = Abs(i-pnj.GetX()) + Abs(j-pnj.GetY())
				posRecolteVillX = i
				posRecolteVillY = j
			}
		}
	}
	// pas d'accès possible pour recolter la ressource
	if distance == 2000 {
		return
	}
	// on attends que le villageois ait finit son déplacement
	var wg sync.WaitGroup
	wg.Add(1)
	go pnj.MoveTo(c, posRecolteVillX, posRecolteVillY, &wg)
	wg.Wait()

	// Le villageois se trouve bien à l'emplacement de la recolte?
	if pnj.GetX() == (posRecolteVillX) && pnj.GetY() == posRecolteVillY {
		go (pnj).Harvest(c, ress, posRecolteVillX, posRecolteVillY)
	}
}

//Harvest : Harvesting of the ressource
func (pnj *Npc) Harvest(c carte.Carte, ress *ressource.Ressource, posRecolteVillX int,
	posRecolteVillY int) {
	uptimeTicker := time.NewTicker(time.Duration(1 * time.Second))
	tpsEcoule := 0
	for {
		// La ressource est épuisée ou le villageois est mort
		if tpsEcoule == ress.GetPv() || pnj.GetPv() == 0 {
			break
		}
		// Le villageois ne se trouve plus à l'emplacement de la ressource
		if pnj.GetX() != (posRecolteVillX) || pnj.GetY() != posRecolteVillY {
			break
		}

		select {
		case <-uptimeTicker.C:
			tpsEcoule++
			switch ress.GetType() {
			case 1:
				tabRessources := make([]int, 3) //0 bois 1 pierre 2 nourriture
				if (*ress).GetPv() <= 0 {
					c.GetTile(ress.X, ress.Y).Empty()
				} else {
					ress.Damage(pnj.damage)
					tabRessources[0] = pnj.damage
					pnj.ressourceChannel <- tabRessources
				}
			case 2:
				tabRessources := make([]int, 3) //0 bois 1 pierre 2 nourriture
				if (*ress).GetPv() <= 0 {
					c.GetTile(ress.X, ress.Y).Empty()
				} else {
					ress.Damage(pnj.damage)
					tabRessources[1] = pnj.damage
					pnj.ressourceChannel <- tabRessources
				}
			case 3:
				tabRessources := make([]int, 3) //0 bois 1 pierre 2 nourriture
				if (*ress).GetPv() <= 0 {
					c.GetTile(ress.X, ress.Y).Empty()
				} else {
					tabRessources[2] = pnj.damage
					ress.Damage(pnj.damage)
					pnj.ressourceChannel <- tabRessources
				}
			default:
				utils.Debug("recolte:ressource inconnue")
			}
		}
	}
}
