package ressource

import "git.unistra.fr/AOEINT/server/constants"
import "git.unistra.fr/AOEINT/server/data"
import "strconv"
import "sync"

//Ressource :
type Ressource struct{
    X int
    Y int
    Pv *safeNumber
    Typ int // 0:water, 1:tree, 2:rock, 3 food ...
}

type safeNumber struct {
	val int
	m   sync.Mutex
}

//InitiatePV remplis la structure pv d'une ressource
func (ress *Ressource)InitiatePV(){
	i := &safeNumber{}
	i.set(100)
	ress.Pv=i
}

func (i *safeNumber) get() int {
	i.m.Lock()
	defer i.m.Unlock()
	return i.val
}

func (i *safeNumber) set(val int) {
	i.m.Lock()
	defer i.m.Unlock()
	i.val = val
}

func (i *safeNumber) sub(val int) {
	i.m.Lock()
	defer i.m.Unlock()
	i.val -= val
}

func new(x int, y int, pv *safeNumber, typ int) Ressource {
    return (Ressource{x,y,pv,typ,})
}

//Create : generate a new npc
func Create(class string, x int, y int) Ressource {
    var res Ressource
	i := &safeNumber{}
    switch class{
        case "water":
			i.set(100)
            res=new(x, y, i, 0)
        case "tree":
			i.set(100)
            res=new(x, y, i,  1)
        case "rock":
			i.set(100)
            res=new(x, y, i,  2)
        case "food":
			i.set(100)
            res=new(x, y, i, 3)
        default:
            res=new(x, y, i,  0) //water
    }
    return res
}

func (ress Ressource)stringify(id string)map[string]string{
	result:=make(map[string]string)
	result["x"]=strconv.Itoa(ress.X)
	result["y"]=strconv.Itoa(ress.Y)
	result["pv"]=strconv.Itoa(ress.GetPv())
	result["type"]=strconv.Itoa(ress.Typ)
	result["id"]=id
	return result
}

//Transmit :
func (ress Ressource) Transmit(id string){
	arr:=ress.stringify(id)
	for k,e := range arr{
		data.AddNewAction(constants.ActionNewRessource,id,k,e)
	}
}

//GetType : return the ress type
func (ress Ressource)GetType() int{
	return ress.Typ
}

//Damage inflige x degats a la ressource
func (ress *Ressource)Damage(x int){
	ress.Pv.sub(x)
}

//GetX : return position X
func (ress Ressource)GetX() int{
	return ress.X
}

//GetY : return position Y
func (ress Ressource)GetY() int{
	return ress.Y
}

//GetPv : return PV
func (ress Ressource)GetPv() int{
	return ress.Pv.get()
}

//IsHarvestable : is the ress harvestable?
func (ress Ressource)IsHarvestable() bool{
	return ress.Typ!=0
}
