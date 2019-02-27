package main

import "fmt"
//import npc "server/npc"
import carte "git.unistra.fr/AOEINT/server/carte"
//import tests "server/test"
import "encoding/json"
import "io/ioutil"
import "git.unistra.fr/AOEINT/server/batiment"
import "os"
import "git.unistra.fr/AOEINT/server/ressource"
import "git.unistra.fr/AOEINT/server/joueur"
import "git.unistra.fr/AOEINT/server/affichage"
import "git.unistra.fr/AOEINT/server/constants"

type Data struct{
	Size int
	Buildings []batiment.Batiment
	Ressources []ressource.Ressource
}

type Game struct{
	joueurs []joueur.Joueur
	carte carte.Carte
	GameRunning bool
}

func (g Game)GetPlayerFromUID(uid string) *joueur.Joueur{
	for i:=0;i<len(g.joueurs);i++{
		if(g.joueurs[i].Uid==uid){
			return &(g.joueurs[i])
		}
	}
	return nil
}

func main() {
	var game Game
	game.GameRunning=true
	(&game).GetPlayerData()
	data:=ExtractData()
	(&game).generateMap(data)
	fmt.Println("Data struct extracted from json:",data)
	affichage.ImprimerCarte(game.carte)
}

func (g *Game)EndOfGame(){
	(*g).GameRunning=false
}

func (g *Game)gameLoop(){
	for (*g).GameRunning{

	}

}

func (g *Game)generateMap(data Data){
	(*g).carte =carte.New(data.Size)
	//On attribue les auberges
	if(len((*g).joueurs)==2){//Si Seulement 2 joueurs fournis, fait en sorte de leur donner des bases adverses
		(*g).joueurs[0].AddBuilding(data.Buildings[0])
		if((*g).carte.AddNewBuilding(&(data.Buildings[0]))==false){
			fmt.Println("Erreur lors du placement d'une auberge")
			os.Exit(1)
		}
		(*g).joueurs[1].AddBuilding(data.Buildings[2])
		if((*g).carte.AddNewBuilding(&(data.Buildings[2]))==false){
			fmt.Println("Erreur lors du placement d'une auberge")
			os.Exit(1)
		}
	}else{//sinon 4 joueurs classiques dans l'ordre des bases fournies (blue blue red red)
		for i:=0;i<4;i++{
			(*g).joueurs[i].AddBuilding(data.Buildings[i])
			if((*g).carte.AddNewBuilding(&(data.Buildings[i]))==false){
				fmt.Println("Erreur lors du placement d'une auberge")
				os.Exit(1)
			}
		}
	}
	for i:=0;i<len(data.Ressources);i++{
		if((*g).carte.AddNewRessource(&(data.Ressources[i]))==false){
			fmt.Println("Erreur lors du placement d'une ressource")
			os.Exit(1)
		}
	}
}

func (g *Game)GetPlayerData(){
	args:=os.Args[1:] //On récupère les paramètres du programme
	fmt.Println(args)
	if(len(args)!=4 && len(args)!=8){
		fmt.Println("Nombre d'arguments incorrect",len(args))
		fmt.Println("usage: ./server PlayerName UidPlayer ...")
		os.Exit(1)
	}
	(*g).joueurs=make([]joueur.Joueur,len(args)/2)
	fmt.Println("nb joueurs=",len(args)/2)
	nbP:=0
	faction:=false
	for i:=0;i<len(args);i+=2{
		if(i>len(args)/2){
			faction=!faction
		}
		(*g).joueurs[nbP]=joueur.Create(faction,args[i],args[i+1])
	}
}

func ExtractData() Data{
	datafileName:="data/GameData.json"
	if(constants.UseSmallMap){
		datafileName="data/SmallTestMap.json"
	}
	jsonFile, err:= os.Open(datafileName)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()
	byteValue,_ := ioutil.ReadAll(jsonFile)
	var newGame Data
	json.Unmarshal(byteValue, &newGame)
	return newGame
}
