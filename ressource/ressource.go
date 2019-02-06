package ressource

import "fmt"

type ressource struct{

    x int
    y int
    pv int
    type int // 0:water, 1:tree, 2:rock, 3 food,

}
//Crée la ressource
func New(x int, y int, pv int, type int) ressource {
    res:=ressource{x,y,pv,tp}
    return res
}

func Create(class string, x int, y int) ressource {
    var res ressource
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
            return -1, errors.New("Class of ressource not specified correctly")
    }
    return res
}
