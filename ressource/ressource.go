package ressource

import "fmt"

type Ressource struct{

    x int
    y int
    pv int
    type int // 0:water, 1:tree, 2:rock, 3 food,

}
//Crée la Ressource
func New(x int, y int, pv int, type int) Ressource {
    res:=Ressource{x,y,pv,tp}
    return res
}

func Create(class string, x int, y int) Ressource {
    var res Ressource
    switch class{
        case "water":
            res=New(x, y, 100,  0)
        case "tree":
            res=New(x, y, 100,  1)
        case "rock":
            res=New(x, y, 100,  2)
        case "food":
            res=New(x, y, 100, 3)
        default:
            return -1, errors.New("Class of Ressource not specified correctly")
    }
    return res
}
