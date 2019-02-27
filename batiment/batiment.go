package batiment

import cst "git.unistra.fr/AOEINT/server/constants"
type Batiment struct{
	X int
	Y int
	Pv int
	Typ int //auberge: 0, caserne:1, établi:2 ...
	Longueur int
	Largeur int
}
//Constructeur de l'objet Batiment
func New(x int,y int, typ int, long int, larg int, pv int) Batiment{
	return (Batiment{x,y,pv,typ,long,larg})
}
//Crée une Instance de batiment
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
//Retourne les pv d'un bâtiment
func (bat Batiment)GetPv() int{
	return bat.Pv
}
//Retourne la longueur d'un batiment
func (bat Batiment)GetLongueur() int{
	return bat.Longueur
}
//Retourne la largeur d'un batiment
func (bat Batiment)GetLargeur() int{
	return bat.Largeur
}
//Retourne un coordonnée x de Batiment
func (bat Batiment)GetX() int{
	return bat.X
}
//Retourne un coordonnée y de Batiment
func (bat Batiment)GetY() int{
	return bat.Y
}
