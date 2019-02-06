package joueur

import npc "server/npc"
import batiment "server/batiment"
import constants "server/constants"

type Joueur struct{
	faction bool //true: faction 1, false: faction 2
	nom string
	nbats int
	batiments[] batiment.Batiment
	nelems int
	entities[] npc.Npc
	id byte
	stone int
	wood int
	food int
}
var model byte =1//Permet d'obtenir des id uniques lors d'une partie

//Crée un joueur
func Create(faction bool,nom string) Joueur{
	return Joueur{faction,nom,0,make([]batiment.Batiment,constants.MaxBuildings),0,make([]npc.Npc,constants.MaxEntities),model++,StartingStone,StartingWood,StartingFood}
}

func GetFaction(j Joueur){
	return j.faction
}

func add