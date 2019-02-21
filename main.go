package main

import "fmt"
//import npc "server/npc"
import carte "server/carte"
//import tests "server/test"
import "encoding/json"
import "io/ioutil"
import "server/batiment"
import "os"
import "server/ressource"
import "server/joueur"


type Data struct{
	Size int
	Camps []batiment.Batiment
	Ressources []ressource.Ressource
}

type Game struct{
	joueurs joueur.Joueur
	carte carte.Carte
	GameRunning bool
}

func main() {
	loopBoolean:=true;
	//npc.Test(mat)
	//var game Game
	//tests.Test(mat)
	data:=ExtractData()
	_ = data
	_ = loopBoolean
	fmt.Println(data)
}

func (g *Game)EndOfGame(){
	(*g).GameRunning=false
}

func (g *Game)gameLoop(){
	for (*g).GameRunning{
		
	}

}

func ExtractData() Data{
	jsonFile, err:= os.Open("data/GameData.json")
	if err!=nil{
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue,_ := ioutil.ReadAll(jsonFile)
	var newGame Data
	json.Unmarshal(byteValue, &newGame)
	return newGame
}