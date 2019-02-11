package npc

import (
	"fmt"
	"server/carte"
	"time"
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
}
//Crée un nouveau Npc avec les paramètres fourni
func New(x int,y int,pv int, vitesse int, vue int, portee int, offensive bool,size int, damage int,selectable bool) Npc{
	pnj:=Npc{x,y,pv,vitesse,vue,portee,offensive,size,damage,selectable}
	return pnj
}
//Crée un Npc du type fourni
func create(class string,x int,y int) Npc{
    var pnj Npc
	switch class{
		case "soldier":
			pnj=New(x,y,8,3,10,2,true,1,2,true)
		case "harvester":
			pnj=New(x,y,4,4,10,2,false,1,10,true)
		default:
			pnj=New(x,y,4,4,10,2,false,1,0,false)
    }
    return pnj
}

func (pnj Npc)deplacement(path []carte.Case){
	if(path!=nil){
		ndep:=len(path)-1
		vdep:=(1000000000/pnj.vitesse)
		for i:=0;i<ndep;i++{
			time.Sleep(time.Duration(vdep))
			pnj.x=path[i].GetPathX()
			pnj.y=path[i].GetPathY()
			fmt.Println("déplacement")
		}
	}
}
func (pnj Npc) MoveTo(c carte.Carte, destx int, desty int) []carte.Case{
	path:= c.GetPathFromTo(pnj.x,pnj.y,destx,desty)
	go pnj.deplacement(path)
	return path
}

func Test(c carte.Carte) {
	entity:=create("soldier",1,1)
	fmt.Println("Hello, playground")
	fmt.Println(entity.pv)
	path:=entity.MoveTo(c,5,8)
	fmt.Printf("Path len=%d\n",len(path))
}