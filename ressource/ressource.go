package ressource

import "fmt"

type Ressource struct{

    x int
    y int
    pv int
    typ int // 0:water, 1:tree, 2:rock, 3 food,

}
//Crée la Ressource
func new(x int, y int, pv int, type int) Ressource {
    return (Ressource{x,y,pv,tp})
}

func Create(class string, x int, y int) Ressource {
    var res Ressource
    switch class{
        case "water":
            res=new(x, y, 100,  0)
        case "tree":
            res=new(x, y, 100,  1)
        case "rock":
            res=new(x, y, 100,  2)
        case "food":
            res=new(x, y, 100, 3)
        default:
            return -1, errors.New("Class of Ressource not specified correctly")
    }
    return res
}

func IsHarvestable(Ressource res) bool{
	return res.typ!="water"
}
