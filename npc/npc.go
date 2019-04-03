package npc


import (
	"time"
	"sync"
	"strconv"
	"log"
	"git.unistra.fr/AOEINT/server/utils"
	"git.unistra.fr/AOEINT/server/carte"
	"git.unistra.fr/AOEINT/server/ressource"
	"git.unistra.fr/AOEINT/server/constants"
	"git.unistra.fr/AOEINT/server/data"
)


//Npc :
type Npc struct {
  x *safeNumberInt
  y *safeNumberInt
  pv *safeNumberInt
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
	// isMoving *safeNumber
	PlayerUUID string
	moveAction map[int](chan bool)
}


type safeNumber struct {
	val bool
	m   sync.Mutex
}


type safeNumberInt struct {
	val int
	m   sync.Mutex
}


//New : new NPC
func New(x *safeNumberInt,y *safeNumberInt,pv *safeNumberInt,vitesse int, vue int, portee int, offensive bool,size int, damage int,selectable bool, typ int,flag bool, channel *chan []int) Npc{
	pnj:=Npc{x,y,pv,vitesse,vue,portee,offensive,size,damage,selectable,typ,flag,*channel,false,"", make(map[int](chan bool))}
	return pnj
}


//Create : generate a new NPC
func Create(class string,x int,y int, flag bool,channel *chan []int) (Npc,string){
	var pnj Npc
	sfPv := &safeNumberInt{}
	sfX := &safeNumberInt{}
	sfX.val = x
	sfY := &safeNumberInt{}
	sfY.val = y
	switch class{
		case "soldier":
			sfPv.val = constants.SoldierPv
			pnj=New(sfX,sfY,sfPv,constants.SoldierVitesse,constants.SoldierVue,
				constants.SoldierPortee, true,constants.SoldierSize,constants.SoldierDamage,true,2,flag,channel)
		case "harvester":
			sfPv.val = constants.HarvesterPv
			pnj=New(sfX,sfY,sfPv,constants.HarvesterVitesse,constants.HarvesterVue,
				constants.HarvesterVillPortee,false,constants.HarvesterSize,constants.HarvesterDamage,true,1,flag,channel)
		default:
			sfPv.val = constants.VillagerPv
			pnj=New(sfX,sfY,sfPv,constants.VillagerVitesse,constants.VillagerVue,
				constants.HarvesterVillPortee,false,constants.VillagerSize,constants.VillagerDamage,false,0,flag,channel)
	}
	id:=(&data.IDMap).AddObject(&pnj)
	//pnj.Transmit(id)
    return pnj,id
}


//Stringify : create a map[string]string of the main arguments of a NPC
func (pnj Npc)Stringify() map[string]string{
	res:=make(map[string]string)
	res["pv"]=strconv.Itoa(pnj.GetPv())
	res["x"]=strconv.Itoa(pnj.GetX())
	res["y"]=strconv.Itoa(pnj.GetY())
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


//Transmit : add the npc to the communcation's buffer
func (pnj Npc) Transmit(id string){
	arr:=pnj.Stringify()
	for k,e := range arr{
		data.AddToAllAction(constants.ActionNewNpc,id,k,e)
	}
}


func (i *safeNumberInt) get() int {
	i.m.Lock()
	defer i.m.Unlock()
	return i.val
}


func (i *safeNumberInt) set(val int) {
	i.m.Lock()
	defer i.m.Unlock()
	i.val = val
}


func (i *safeNumberInt) sub(val int) {
	i.m.Lock()
	defer i.m.Unlock()
	i.val -= val
}


//GetPv : return the HP
func (pnj Npc) GetPv() int {
	return pnj.pv.get()
}


//SetPv : set the npc's HP value
func (pnj *Npc) SetPv(val int) {
	pnj.pv.set(val)
}


//SubPv : decrement the npc's HP value
func (pnj *Npc) SubPv(val int) {
	pnj.pv.sub(val)
}


//GetX : return the position X
func (pnj Npc) GetX() int{
	return pnj.x.get()
}


//SetX : set the npc's X value
func (pnj *Npc) SetX(val int) {
	pnj.x.set(val)
}


//GetY : return the position Y
func (pnj Npc) GetY() int{
	return pnj.y.get()
}


//SetY : set the npc's Y value
func (pnj *Npc) SetY(val int) {
	pnj.y.set(val)
}


//GetVue : return villager's vision
func (pnj Npc) GetVue() int{
	return pnj.vue
}


//GetType : return the villager's type
func (pnj Npc) GetType() int{
	return pnj.typ
}


//GetSpeed : return the npc's speed
func (pnj Npc)GetSpeed() int{
	return pnj.vitesse
}



func (pnj *Npc)deplacement(path []carte.Case, wg *sync.WaitGroup, moveA *chan bool) {
	if(path!=nil) {
		// Cancel the old movement
		index:=len(pnj.moveAction)-1
		pnj.moveAction[index] = make(chan bool, 2)
		pnj.moveAction[index] <- true
		pnj.moveAction[index] <- true
		pnj.moveAction[index] = *moveA
		delete(pnj.moveAction,index)
		//var initialposX, initialposY int
		ndep:=len(path)-1
		vdep:=(1000000000/pnj.vitesse)
		for i:=0;i<=ndep;i++{
			select {
			case <- *moveA:
				wg.Done()
				return
			default:
				//initialposX, initialposY = pnj.GetX(), pnj.GetY()
				time.Sleep(time.Duration(vdep))
				/*
				if(pnj.GetX() != initialposX || pnj.GetY() != initialposY){
					wg.Done()
					return
				}
				*/
				pnj.SetX(path[i].GetPathX())
				pnj.SetY(path[i].GetPathY())
			}
		}
		if (wg != nil) {
			wg.Done()
		}
	}
}


//MoveTo : move a npc from his position(x,y) to another position(x,y)
func (pnj *Npc) MoveTo(c carte.Carte, destx int, desty int, wg *sync.WaitGroup, moveA *chan bool) []carte.Case{
	var path []carte.Case
	path=nil
	if(c.GetTile(destx,desty).GetType()==0){
		path= c.GetPathFromTo(pnj.GetX(),pnj.GetY(),destx,desty)
		go pnj.deplacement(path, wg, moveA)
	}
	return path
}


//Abs : utility function
func Abs(x int) int {
	if (x < 0) {
		return -x
	}
	return x
}


//RecoltePossible : return true if te villager can acces to a tile to harvest the resource in x, y
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


/*MoveFight : attack a given npc and also trigger the fight-back
* Both the aggressor and the target while fight and chase unless the player orders
another action or loses vision of the other NPC
*/
func (pnj *Npc)MoveFight(c carte.Carte, target *Npc, moveA *chan bool){

	if (pnj.GetVue() < (Abs(target.GetX() - pnj.GetX()) + Abs(target.GetY() - pnj.GetY())) ){
		log.Print("Le npc ciblé n'est pas dans la vue du npc")
		return
	}
	initialPosTargetX, initialPosTargetY  := target.GetX(), target.GetY()
	var posFightPnjX, posFightPnjY int

	var i, j int

	distance := 2000

	for i = target.GetX() - pnj.portee; i <= target.GetX() + pnj.portee; i++{
		for j = target.GetY() - pnj.portee; j <= target.GetY() + pnj.portee; j++{
			if ( (Abs(i - pnj.GetX()) + Abs(j - pnj.GetY()) ) < distance &&
				c.IsEmpty(i, j)){
				distance = Abs(i - pnj.GetX()) + Abs(j - pnj.GetY())
				posFightPnjX = i
				posFightPnjY = j
			}
		}
	}


	// No space available to attack the ennemy
	if (distance == 2000){
		return
	}
	// Wait that the npc is in the range to attack
	go pnj.MoveTo(c, posFightPnjX, posFightPnjY, nil, moveA)

	/* Verify each x ms that the target didn't move from his initial position
	*  if he did move, do MoveTo to the new position, if not fight him when the
	*  movement is finished
	*/
	uptimeTicker := time.NewTicker(time.Duration(10 * time.Millisecond))
	for{
		select {
		case <-uptimeTicker.C:
			// if the target is not in the aggressor's vision anymore, he stops chasing him
			if (pnj.GetVue() < (Abs(target.GetX() - pnj.GetX()) + Abs(target.GetY() - pnj.GetY())) ){
				log.Print("Le npc ciblé n'est pas dans la vue du npc")
				return
			}

			ch := make(chan bool, 2)

			if (initialPosTargetX != target.GetX() || initialPosTargetY != target.GetY()){
				distance = 2000

				for i = target.GetX() - pnj.portee; i <= target.GetX() + pnj.portee; i++{
					for j = target.GetY() - pnj.portee; j <= target.GetY() + pnj.portee; j++{
						if ( (Abs(i - pnj.GetX()) + Abs(j - pnj.GetY()) ) < distance &&
							c.IsEmpty(i, j)){
							distance = Abs(i - pnj.GetX()) + Abs(j - pnj.GetY())
							posFightPnjX = i
							posFightPnjY = j
						}
					}
				}

				// No space available to attack the ennemy
				if (distance == 2000){
					return
				}
				// Wait that the npc is in the range to attack
				go pnj.MoveTo(c, posFightPnjX, posFightPnjY,nil, &ch)
			}
			// The aggressor finished his movement and so can start fighting him
			if (pnj.GetX() == (posFightPnjX) && pnj.GetY() == posFightPnjY){
				ch := make(chan bool, 2)
				chTarget := make(chan bool, 2)
				//Fight
				go pnj.Fight(c, target, posFightPnjX, posFightPnjY, &ch)
				//The target fights back after a short delay
				time.Sleep(time.Duration((1 * time.Second)/4))
				target.Fight(c, pnj, target.GetX(), target.GetY(), &chTarget)
				break;
			}
		}
	}
}

//Fight : attack a npc
func (pnj *Npc)Fight(c carte.Carte, target *Npc, posFightPnjX int,
	 posFightPnjY int, moveA *chan bool){
	uptimeTicker := time.NewTicker(time.Duration(1 * time.Second))
	initialPosTargetX, initialPosTargetY  := target.GetX(), target.GetY()
	for {
		// if the target is not in the aggressor's vision anymore, he stops chasing him
		if (pnj.GetVue() < (Abs(target.GetX() - pnj.GetX()) + Abs(target.GetY() - pnj.GetY())) ){
			log.Print("Le npc ciblé n'est pas dans la vue du npc")
			return
		}
		// if the target moved start chasing him again
		if (initialPosTargetX != target.GetX() || initialPosTargetY != target.GetY()){
			go pnj.MoveFight(c, target, moveA)
			return
		}
		//The target or the attacker is dead
		if (target.GetPv() == 0 || pnj.GetPv() == 0){
			break
		}
		//The attacker moved
		if (pnj.GetX() != (posFightPnjX) || pnj.GetY() != posFightPnjY){
			break
		}
		//The target moved
		if (target.GetX() != initialPosTargetX || target.GetY() != initialPosTargetY){
			break
		}

		select {
		case <- *moveA:
			return
		case <-uptimeTicker.C:
			target.SubPv(pnj.damage)
		}
	}
}

/*
//MoveHarvest : (move to the neareast ressource in the villagers's vision)
func (pnj *Npc)MoveHarvest(c carte.Carte){
	var i, j int
	var ress *ressource.Ressource
	distance := 2000
	if (pnj.GetType() == 2){
		utils.Debug("Un soldat ne peut pas recolter de ressources")
		return
	}
	for i = pnj.GetX() - pnj.GetVue(); i <= pnj.GetX() + pnj.GetVue() || i > c.GetSize(); i++{
		if (i < 0){
			i = 0
		}
		for j = pnj.GetY() - pnj.GetVue(); j <= pnj.GetY() + pnj.GetVue() || j > c.GetSize(); j++{
			if (j < 0){
				j = 0
			}
			if (c.GetTile(i, j).GetType() == 2){
				if (c.GetTile(i, j).GetRess().GetType() == 2 && pnj.GetType() != 0){
					utils.Debug("Seul un harvester peut recolter de la pierre")
					continue;
				}
				if ((Abs(i - pnj.GetX()) + Abs(j - pnj.GetY())) < distance &&
					RecoltePossible(c, i, j)){
					distance = Abs(i - pnj.GetX()) + Abs(j - pnj.GetY())
					ress = c.GetTile(i, j).GetRess()
				}
			}
		}
	}

	// pas de ressources dans la vue du villageois
	if (distance == 2000){
		return
	}

	var posRecolteVillX, posRecolteVillY int
	distance = 2000

	for i = ress.GetX() - pnj.portee; i <= ress.GetX() + pnj.portee; i++{
		for j = ress.GetY() - pnj.portee; j <= ress.GetY() + pnj.portee; j++{
			if ( (Abs(i - pnj.GetX()) + Abs(j - pnj.GetY()) ) < distance &&
				c.IsEmpty(i, j)){
				distance = Abs(i - pnj.GetX()) + Abs(j - pnj.GetY())
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
    go pnj.MoveTo(c, posRecolteVillX, posRecolteVillY, &wg)
	wg.Wait()


	// Le villageois se trouve bien à l'emplacement de la recolte?
	if (pnj.GetX() == (posRecolteVillX) && pnj.GetY() == posRecolteVillY){
		 go (pnj).Harvest(c, ress, posRecolteVillX, posRecolteVillY)
	}
}
*/

/*MoveHarvestTarget : (move to the nreast ressource in the villagers's vision).
Triggered when a villager is inactive, cancelled when the player moves the npc
*/
func (pnj *Npc)MoveHarvestTarget(c carte.Carte, ress *ressource.Ressource, moveA *chan bool){
	var i, j int
	//Verify the parameters
	if (pnj.GetType() == 2){
		log.Print("Un soldat ne peut pas recolter de ressources")
		return
	}
	if (ress.GetType() == 2 && pnj.GetType() != 0){
		log.Print("Seul un harvester peut recolter de la pierre")
		return
	}
	if (pnj.GetVue() < (Abs(ress.GetX() - pnj.GetX()) + Abs(ress.GetY() - pnj.GetY())) ){
		log.Print("La ressource n'est pas dans la vue du npc")
		return
	}

	var posRecolteVillX, posRecolteVillY int
	distance := 2000

	for i = ress.GetX() - pnj.portee; i <= ress.GetX() + pnj.portee; i++{
		for j = ress.GetY() - pnj.portee; j <= ress.GetY() + pnj.portee; j++{
			if ( (Abs(i - pnj.GetX()) + Abs(j - pnj.GetY()) ) < distance &&
				c.IsEmpty(i, j)){
				distance = Abs(i - pnj.GetX()) + Abs(j - pnj.GetY())
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
    go pnj.MoveTo(c, posRecolteVillX, posRecolteVillY, &wg, moveA)
	wg.Wait()

	// Le villageois se trouve bien à l'emplacement de la recolte?
	if (pnj.GetX() == (posRecolteVillX) && pnj.GetY() == posRecolteVillY){
		 go (pnj).Harvest(c, ress, posRecolteVillX, posRecolteVillY, moveA)
	}
}


//Harvest : Harvesting of the ressource
func (pnj *Npc)Harvest(c carte.Carte, ress *ressource.Ressource, posRecolteVillX int,
	 posRecolteVillY int, moveA *chan bool){
	uptimeTicker := time.NewTicker(time.Duration(1 * time.Second))
	tpsEcoule := 0
	for {
		// La ressource est épuisée ou le villageois est mort
		if (tpsEcoule == ress.GetPv() || pnj.GetPv() == 0){
			break
		}
		// Le villageois ne se trouve plus à l'emplacement de la ressource
		if (pnj.GetX() != (posRecolteVillX) || pnj.GetY() != posRecolteVillY){
			break
		}

		select {
		case <- *moveA:
			return
		case <-uptimeTicker.C:
			tpsEcoule++
			switch ress.GetType(){
			case 1:
				tabRessources:=make([]int,3) //0 bois 1 pierre 2 nourriture
				if((*ress).GetPv()<=0){
					c.GetTile(ress.X,ress.Y).Empty()
				}else{
					ress.Damage(pnj.damage)
					tabRessources[0]=pnj.damage
					pnj.ressourceChannel<-tabRessources
				}
			case 2:
				tabRessources:=make([]int,3) //0 bois 1 pierre 2 nourriture
				if((*ress).GetPv()<=0){
					c.GetTile(ress.X,ress.Y).Empty()
				}else{
					ress.Damage(pnj.damage)
					tabRessources[1]=pnj.damage
					pnj.ressourceChannel<-tabRessources
				}
			case 3:
				tabRessources:=make([]int,3) //0 bois 1 pierre 2 nourriture
				if((*ress).GetPv()<=0){
					c.GetTile(ress.X,ress.Y).Empty()
				}else{
					tabRessources[2]=pnj.damage
					ress.Damage(pnj.damage)
					pnj.ressourceChannel<-tabRessources
				}
			default:
				utils.Debug("recolte:ressource inconnue")
			}
		}
	}
}
