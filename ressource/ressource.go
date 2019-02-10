package ressource

//import "fmt"

type Ressource struct{

    x int
    y int
    pv int
    typ int // 0:water, 1:tree, 2:rock, 3 food ...

}
//Crée la Ressource
func new(x int, y int, pv int, typ int) Ressource {
    return (Ressource{x,y,pv,typ})
}

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

func (res Ressource)GetType() int{
	return res.typ
}

func (res Ressource)GetX() int{
	return res.x
}

func (res Ressource)GetY() int{
	return res.y
}

func (res Ressource)IsHarvestable() bool{
	return res.typ!=0
}
