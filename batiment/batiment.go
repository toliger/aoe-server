package batiment

import cst "git.unistra.fr/AOEINT/server/constants"
import "git.unistra.fr/AOEINT/server/data"
import "strconv"

type Batiment struct{
	X int
	Y int
	Pv int
	Typ int //auberge: 0, caserne:1, établi:2 ...
	Longueur int
	Largeur int
	PlayerUID string
}
//Constructeur de l'objet Batiment
func New(x int,y int, typ int, long int, larg int, pv int) Batiment{
	return (Batiment{x,y,pv,typ,long,larg,""})
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

func (bat Batiment)stringify(id string)map[string]string{
	res:=make(map[string]string)
	res["x"]=strconv.Itoa(bat.X)
	res["y"]=strconv.Itoa(bat.Y)
	res["pv"]=strconv.Itoa(bat.Pv)
	res["type"]=strconv.Itoa(bat.Typ)
	res["PlayerUUID"]=bat.PlayerUID
	res["id"]=id
	return res
}

func (bat Batiment) Transmit(id string){
	arr:=bat.stringify(id)
	for k,e := range arr{
		data.AddNewAction(cst.ACTION_NEWBUILDING,id,k,e)
	}
}

//"Detruit" l'objet batiment si il n'y a plus de pv
func (bat *Batiment)DestroyBuilding(){
	bat = nil //nil permet assigner la valeur nul à un pointeur
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
