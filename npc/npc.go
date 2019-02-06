package npc

import (
	"fmt"
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

func Test() {
	entity:=create("soldier",1,1)
	fmt.Println("Hello, playground")
	fmt.Println(entity.pv)
}