package npc

import (
	"fmt"
	"git.unistra.fr/AOEINT/server/carte"
	"git.unistra.fr/AOEINT/server/ressource"
	"git.unistra.fr/AOEINT/server/constants"
	"time"
	"sync"
	"git.unistra.fr/AOEINT/server/data"
	"strconv"
)

type Npc struct {
    x int
    y int
    pv int
    vitesse int
    vue int
    portee int
    offensive bool//true=soldier else harvester
    size int
    damage int
    selectable bool //false=villager
	typ int // 0:villager, 1:harvester, 2:soldier
	TeamFlag bool
	ressourceChannel chan []int
	hasOrder bool //Si un déplacement a dejà été demandé par le joueur (disable auto movement)
	PlayerUUID string
}
//Crée un nouveau Npc avec les paramètres fourni
func New(x int,y int,pv int, vitesse int, vue int, portee int, offensive bool,size int, damage int,selectable bool, typ int,flag bool, channel chan []int) Npc{
	pnj:=Npc{x,y,pv,vitesse,vue,portee,offensive,size,damage,selectable,typ,flag,channel,false,""}
	return pnj
}

//Crée un Npc du type fourni
func Create(class string,x int,y int, flag bool,channel chan []int) (Npc,string){
	var pnj Npc
	switch class{
		case "soldier":
			pnj=New(x,y,constants.SoldierPv,constants.SoldierVitesse,constants.SoldierVue,
				constants.SoldierPortee, true,constants.SoldierSize,constants.SoldierDamage,true,2,flag,channel)
		case "harvester":
			pnj=New(x,y,constants.HarvesterPv,constants.HarvesterVitesse,constants.HarvesterVue,
				constants.HarvesterVillPortee,false,constants.HarvesterSize,constants.HarvesterDamage,true,1,flag,channel)
		default:
			pnj=New(x,y,constants.VillagerPv,constants.VillagerVitesse,constants.VillagerVue,
				constants.HarvesterVillPortee,false,constants.VillagerSize,constants.VillagerDamage,false,0,flag,channel)
	}
	id:=(&data.IdMap).AddObject(&pnj)
	pnj.Transmit(id)
    return pnj,id
}

func (pnj Npc)stringify() map[string]string{
	res:=make(map[string]string)
	res["pv"]=strconv.Itoa(pnj.pv)
	res["x"]=strconv.Itoa(pnj.x)
	res["y"]=strconv.Itoa(pnj.y)
	res["vitesse"]=strconv.Itoa(pnj.vitesse)
	res["type"]=strconv.Itoa(pnj.typ)
	res["damage"]=strconv.Itoa(pnj.damage)
	res["offensive"]=strconv.FormatBool(pnj.offensive)
	res["vue"]=strconv.Itoa(pnj.vue)
	res["portee"]=strconv.Itoa(pnj.portee)
	res["TeamFlag"]=strconv.FormatBool(pnj.TeamFlag)
	res["PlayerUUID"]=pnj.PlayerUUID
	return res
}

//Ajoute le npc au buffer de communication
func (pnj Npc) Transmit(id string){
	arr:=pnj.stringify()
	for k,e := range arr{
		data.AddNewAction(constants.ACTION_NEWNPC,id,k,e)
	}
}

//Npc
func (pnj Npc) GetX() int{
	return pnj.x
}

func (pnj Npc) GetY() int{
	return pnj.y
}

func (pnj Npc) GetVue() int{
	return pnj.vue
}

func (pnj Npc) GetType() int{
	return pnj.typ
}

func (pnj Npc) GetPv() int{
	return pnj.pv
}

func (pnj *Npc)deplacement(path []carte.Case, wg *sync.WaitGroup) {
	if(path!=nil) {
		ndep:=len(path)-1
		vdep:=(1000000000/pnj.vitesse)
		for i:=0;i<=ndep;i++{
			time.Sleep(time.Duration(vdep))
			pnj.x=path[i].GetPathX()
			pnj.y=path[i].GetPathY()
		}
		if (wg != nil) {
			wg.Done()
		}
	}
}
func (pnj *Npc) MoveTo(c carte.Carte, destx int, desty int, wg *sync.WaitGroup) []carte.Case{
	path:= c.GetPathFromTo(pnj.x,pnj.y,destx,desty)
	go pnj.deplacement(path, wg)
	return path
}

func Abs(x int) int {
	if (x < 0) {
		return -x
	}
	return x
}
func (pnj Npc)GetSpeed() int{
	return pnj.vitesse
}

// Renvoie vrai si le villageois peut accéder à une case pour recolter la ressource en x,y
func RecoltePossible(c carte.Carte, x int, y int) bool{
	for i := x-1; i <= x+1; i++{
		for j := y-1; j <= y+1; j++{
			if (c.IsEmpty(i, j)){
				return true
			}
		}
	}
	return false
}



//Recolte de ressources (se deplace vers la ressource la plus proche dans la vue du villageois)
func (vill *Npc)DeplacementRecolte(c carte.Carte){
	var i, j int
	var ress *ressource.Ressource
	distance := 2000
	if (vill.GetType() == 2){
		fmt.Println("Un soldat ne peut pas recolter de ressources")
		return
	}
	for i = vill.GetX() - vill.GetVue(); i <= vill.GetX() + vill.GetVue() || i > c.GetSize(); i++{
		if (i < 0){
			i = 0
		}
		for j = vill.GetY() - vill.GetVue(); j <= vill.GetY() + vill.GetVue() || j > c.GetSize(); j++{
			if (j < 0){
				j = 0
			}
			if (c.GetTile(i, j).GetType() == 2){
				if (c.GetTile(i, j).GetRess().GetType() == 2 && vill.GetType() != 0){
					fmt.Println("Seul un harvester peut recolter de la pierre")
					continue;
				}
				if ((Abs(i - vill.GetX()) + Abs(j - vill.GetY())) < distance &&
					RecoltePossible(c, i, j)){
					distance = Abs(i - vill.GetX()) + Abs(j - vill.GetY())
					ress = c.GetTile(i, j).GetRess()
				}
			}
		}
	}
	//fmt.Println("ressource?",ress == nil)

	// pas de ressources dans la vue du villageois
	if (distance == 2000){
		return
	}

	var posRecolteVillX, posRecolteVillY int
	distance = 2000

	for i = ress.GetX() - constants.HarvesterVillPortee; i <= ress.GetX() + constants.HarvesterVillPortee; i++{
		for j = ress.GetY() - constants.HarvesterVillPortee; j <= ress.GetY() + constants.HarvesterVillPortee; j++{
			if ( (Abs(i - vill.GetX()) + Abs(j - vill.GetY()) ) < distance &&
				c.IsEmpty(i, j)){
				distance = Abs(i - vill.GetX()) + Abs(j - vill.GetY())
				posRecolteVillX = i
				posRecolteVillY = j
			}
		}
	}
	// pas d'accès possible pour recolter la ressource
	if (distance == 2000){
		return
	}
	// on attends que le villageois ait finit son déplacement
	var wg sync.WaitGroup
	wg.Add(1)
    go vill.MoveTo(c, posRecolteVillX, posRecolteVillY, &wg)
	wg.Wait()

    // fmt.Printf("posRecolteVillX : %d, posRecolteVillY : %d\n", posRecolteVillX, posRecolteVillY)
	// fmt.Printf("villX : %d, villY: %d\n", vill.GetX(), vill.GetY())

	// Le villageois se trouve bien à l'emplacement de la recolte?
	if (vill.GetX() == (posRecolteVillX) && vill.GetY() == posRecolteVillY){
		 go Recolte(vill, c, ress, posRecolteVillX, posRecolteVillY)
	}
}

// Effectue la recolte de la ressource (x par seconde)
func Recolte(vill *Npc, c carte.Carte, ress *ressource.Ressource,
	posRecolteVillX int, posRecolteVillY int){
	uptimeTicker := time.NewTicker(time.Duration(1 * time.Second))
	tps_ecoule := 0
	for {
		// La ressource est épuisée ou le villageois est mort
		if (tps_ecoule == ress.GetPv() || vill.GetPv() == 0){
			break
		}
		// Le villageois ne se trouve plus à l'emplacement de la ressource
		if (vill.GetX() != (posRecolteVillX) || vill.GetY() != posRecolteVillY){
			break;
		}
		select {
		case <-uptimeTicker.C:
			tps_ecoule++
			switch ress.GetType(){
			case 1:
				tabRessources:=make([]int,3) //0 bois 1 pierre 2 nourriture
				tabRessources[0]=vill.damage
				(*ress).Pv-=vill.damage
				if((*ress).Pv<=0){
					c.GetTile(ress.X,ress.Y).Empty()
				}
				vill.ressourceChannel<-tabRessources
			case 2:
				tabRessources:=make([]int,3) //0 bois 1 pierre 2 nourriture
				tabRessources[1]=vill.damage
				(*ress).Pv-=vill.damage
				if((*ress).Pv<=0){
					c.GetTile(ress.X,ress.Y).Empty()
				}
				vill.ressourceChannel<-tabRessources
			case 3:
				tabRessources:=make([]int,3) //0 bois 1 pierre 2 nourriture
				tabRessources[2]=vill.damage
				(*ress).Pv-=vill.damage
				if((*ress).Pv<=0){
					c.GetTile(ress.X,ress.Y).Empty()
				}
				vill.ressourceChannel<-tabRessources
			default:
				fmt.Println("recolte:ressource inconnue")
			}
		}
	}
}
