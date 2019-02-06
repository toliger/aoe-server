package ressource

//import "fmt"

type Ressource struct{

    x int
    y int
    pv int
    typ int // 0:water, 1:tree, 2:rock, 3 food,

}
//Crée la Ressource
func new(x int, y int, pv int, typ int) Ressource {
    return (Ressource{x,y,pv,typ})
}

func Create(class string, x int, y int) Ressource {
    var res Ressource
    switch class{
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

func GetType(res Ressource) int{
	return res.typ
}

func GetX(res Ressource) int{
	return res.x
}

func GetY(res Ressource) int{
	return res.y
}

func IsHarvestable(res Ressource) bool{
	return res.typ!=0
}
