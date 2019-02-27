//Contient toutes les fonctions pour les echanges de donnees client/serveur
package client

import (
	"git.unistra.fr/AOEINT/server/npc"
	"git.unistra.fr/AOEINT/server/carte"
	"git.unistra.fr/AOEINT/server/ressource"
	"git.unistra.fr/AOEINT/server/joueur"
	"git.unistra.fr/AOEINT/server/batiment"
)

//serveur vers client

//Envoie toutes les donnees necessaires à la mise en place de la partie en debut de jeu
//A envoyer: donnees des joueurs, structure data(map), entites de depart..
func initGameState(){

}

//Maj les ressources du joueur à partir de l'uid correspondant
func updatePlayerRessources(playerUID string,stone int,wood int,food int){}

//Maj: Indique la destruction d'un Batiment au client pour qu'il soit retire
func BuildingDestroyed(playerUID string,x int, y int){

}

//Permet de Maj la liste des npcs visibles en indiquant leur mort au client
func PlayerNpcsDied(playerUID string,npc []npc.Npc){

}

//client vers serveur

//demande la creation d'un batiment à partir de l'uid du joueur, une position et un type de batiment
//class: "auberge","caserne","etabli"
func TryToBuild(playerUID string, x int, y int, class string) bool{
	return false
}

//Demande le deplacement des npc selectionnes
func MoveSelectedNpc(playerUID string, liste []npc.Npc, x int, y int){

}

//Demande la suppression par le joueur de l'un de ses batiments
func EraseBuilding(playerUID string, x int, y int){

}

//Averti le serveur de la creation d'une entite: verification des ressources necessaires
func AddNewNpc(playerUID string, x int, y int, typ int) bool{
	return false
}
//Enleve des Pv a un batiment
func  DamageBuilding(playerUID string, x int, y int, attack int){

}
