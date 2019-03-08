package npc

import (
	"fmt"
	"git.unistra.fr/AOEINT/server/carte"
	"time"
	"sync"
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
}
//Crée un nouveau Npc avec les paramètres fourni
func New(x int,y int,pv int, vitesse int, vue int, portee int, offensive bool,size int, damage int,selectable bool, typ int) Npc{
	pnj:=Npc{x,y,pv,vitesse,vue,portee,offensive,size,damage,selectable,typ}
	return pnj
}
//Crée un Npc du type fourni
func Create(class string,x int,y int) Npc{
	var pnj Npc
	switch class{
		case "soldier":
			pnj=New(x,y,8,3,10,2,true,1,2,true,2)
		case "harvester":
			pnj=New(x,y,4,4,10,2,false,1,10,true,1)
		default:
			pnj=New(x,y,4,4,10,2,false,1,0,false,0)
	}
	return pnj
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

func (pnj *Npc)deplacement(path []carte.Case, wg *sync.WaitGroup){
	if(path!=nil){
		ndep:=len(path)-1
		vdep:=(1000000000/pnj.vitesse)
		for i:=0;i<ndep;i++{
			time.Sleep(time.Duration(vdep))
			pnj.x=path[i].GetPathX()
			pnj.y=path[i].GetPathY()
			//fmt.Println("déplacement")
		}
		wg.Done()
	}
}
func (pnj *Npc) MoveTo(c carte.Carte, destx int, desty int, wg *sync.WaitGroup) []carte.Case{
	path:= c.GetPathFromTo(pnj.x,pnj.y,destx,desty)
	go pnj.deplacement(path, wg)
	return path
}

func Test(c carte.Carte) {
	var wg sync.WaitGroup
	entity:=Create("soldier",1,1)
	fmt.Println("Hello, playground")
	fmt.Println(entity.pv)
	path:=entity.MoveTo(c,5,8,&wg)
	fmt.Printf("Path len=%d\n",len(path))
}
