package batiment

import cst "server/constants"

type Batiment struct{
	x int
	y int
	pv int
	typ int //auberge: 0, caserne:1, établi:2 ...
	longueur int
	largeur int
	
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
	return bat.longueur
}

func (bat Batiment)GetLargeur() int{
	return bat.largeur
}

func (bat Batiment)GetX() int{
	return bat.x
}

func (bat Batiment)GetY() int{
	return bat.y
}