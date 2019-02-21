package batiment

import cst "server/constants"

type Batiment struct{
	X int
	Y int
	Pv int
	Typ int //auberge: 0, caserne:1, établi:2 ...
	Longueur int
	Largeur int
}
//Crée un nouveau bâtiment, pv = 100
func New(x int,y int, typ int, long int, larg int, pv int) Batiment{
	return (Batiment{x,y,pv,typ,long,larg})
}

func Create(class string, x int, y int ) Batiment{
	var bat Batiment
	switch class{
	case "auberge":
		bat=New(x, y, 0,cst.LongueurAuberge,cst.LargeurAuberge,cst.PVAuberge)
	case "caserne":
		bat=New(x, y, 1,cst.LongueurCaserne,cst.LargeurCaserne,cst.PVCaserne)
	case "etabli":
		bat=New(x, y, 2,cst.LongueurEtabli,cst.LargeurEtabli,cst.PVEtabli)
	default: //défaut=auberge
		bat=New(x, y, 0,cst.LongueurAuberge,cst.LargeurAuberge,cst.PVAuberge)
	}
	return bat
}

func (bat Batiment)GetLongueur() int{
	return bat.Longueur
}

func (bat Batiment)GetLargeur() int{
	return bat.Largeur
}

func (bat Batiment)GetX() int{
	return bat.X
}

func (bat Batiment)GetY() int{
	return bat.Y
}