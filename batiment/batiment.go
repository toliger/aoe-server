package batiment

type Batiment struct{
	x int
	y int
	pv int
	typ int //auberge: 0, caserne:1, établi:2 ...
}
//Crée un nouveau bâtiment, pv = 100
func New(x int,y int, t int) Batiment{
	return (Batiment{x,y,100,t})
}

func Create(class string, x int, y int ) Batiment{
	var bat Batiment
	switch class{
	case "auberge":
		bat=New(x, y, 0)
	case "caserne":
		bat=New(x, y, 1)
	case "etabli":
		bat=New(x, y, 2)
	default:
		bat=New(x, y, 0)
	}
	return bat
}
