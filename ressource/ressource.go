package ressource

import "git.unistra.fr/AOEINT/server/constants"
import "git.unistra.fr/AOEINT/server/data"
import "strconv"
import "sync"

//Ressource :
type Ressource struct{
    X int
    Y int
    Pv int
    Typ int // 0:water, 1:tree, 2:rock, 3 food ...
	mutex *sync.RWMutex
}

func new(x int, y int, pv int, typ int) Ressource {
    return (Ressource{x,y,pv,typ,&sync.RWMutex{}})
}

//Create : generate a new npc
func Create(class string, x int, y int) Ressource {
    var res Ressource
    switch class{
        case "water":
            res=new(x, y, 100, 0)
        case "tree":
            res=new(x, y, 100,  1)
        case "rock":
            res=new(x, y, 100,  2)
        case "food":
            res=new(x, y, 100, 3)
        default:
            res=new(x, y, 100,  0) //water
    }
    return res
}

func (res Ressource)stringify()map[string]string{
	result:=make(map[string]string)
	result["x"]=strconv.Itoa(res.X)
	result["y"]=strconv.Itoa(res.Y)
	result["pv"]=strconv.Itoa(res.GetPv())
	result["type"]=strconv.Itoa(res.Typ)
	return result
}

//Transmit :
func (res Ressource) Transmit(id string){
	arr:=res.stringify()
	for k,e := range arr{
		data.AddNewAction(constants.ActionNewRessource,id,k,e)
	}
}

//GetType : return the ress type
func (res Ressource)GetType() int{
	return res.Typ
}

//Damage inflige x degats a la ressource
func (res *Ressource)Damage(x int){
	res.mutex.Lock()
	defer res.mutex.Unlock()
	res.Pv-=x
}

//GetX : return position X
func (res Ressource)GetX() int{
	return res.X
}

//GetY : return position Y
func (res Ressource)GetY() int{
	return res.Y
}

//GetPv : return PV
func (res Ressource)GetPv() int{
	res.mutex.RLock()
	defer res.mutex.RUnlock()
	return res.Pv
}

//IsHarvestable : is the ress harvestable?
func (res Ressource)IsHarvestable() bool{
	return res.Typ!=0
}
