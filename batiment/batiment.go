package batiment

import (
  cst "git.unistra.fr/AOEINT/server/constants"
  "git.unistra.fr/AOEINT/server/data"
  "git.unistra.fr/AOEINT/server/utils"
  "strconv"
  "log"
)


//Batiment : Structure contenant tous les éléments nécessaires pour la gestion d'un batiment
type Batiment struct{
	X int
	Y int
	Pv int
	Typ int //auberge: 0, caserne:1, établi:2 ...
	Longueur int
	Largeur int
	PlayerUID string
    batimentChannel *chan int
}


//New : Constructeur de l'objet Batiment
func New(x int,y int, typ int, long int, larg int, pv int) Batiment{
    buffer:=make(chan int,cst.BatimentBufferSize)
	return (Batiment{x,y,pv,typ,long,larg,"",&buffer })
}


//Create : Crée une Instance de batiment
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
    go (&bat).batimentUpdate()
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


//Transmit : Adds the corresponding action to ActionBuffer
func (bat Batiment) Transmit(id string){
	arr:=bat.stringify(id)
	for k,e := range arr{
		data.AddToAllAction(cst.ActionNewBuilding,id,k,e)
	}
}


/*batimentUpdate : Met automatiquement a jour les pv du batiment à partir du channel du batiment
*/
func (bat *Batiment) batimentUpdate(){
	utils.Debug("batiment:channel actif")
    var res int
	for{
        res= <- *(bat.batimentChannel)
        //log.Print(res)
        bat.SubPv(res)
        log.Print((*bat).GetPv())
	}
	utils.Debug("batiment:channel inactif")
}


//DestroyBuilding : "Detruit" l'objet batiment si il n'y a plus de pv
func (bat *Batiment)DestroyBuilding(){
	bat = nil //nil permet assigner la valeur nul à un pointeur
}


//GetChannel retourne le channel de ressource du joueur
func(bat *Batiment)GetChannel() *(chan int){
	return bat.batimentChannel
}


//GetPv : Retourne les pv d'un bâtiment
func (bat Batiment)GetPv() int{
	return bat.Pv
}

//SetPv : change des pv d'un bâtiment
func (bat *Batiment)SetPv(val int){
    bat.Pv = val
}

//SubPv : decrement les pv d'un bâtiment de val
func (bat *Batiment)SubPv(val int){
    bat.Pv -= val
}


//GetLongueur : Retourne la longueur d'un batiment
func (bat Batiment)GetLongueur() int{
	return bat.Longueur
}


//GetLargeur : Retourne la largeur d'un batiment
func (bat Batiment)GetLargeur() int{
	return bat.Largeur
}


//GetX : Retourne un coordonnée x de Batiment
func (bat Batiment)GetX() int{
	return bat.X
}


//GetY : Retourne un coordonnée y de Batiment
func (bat Batiment)GetY() int{
	return bat.Y
}
