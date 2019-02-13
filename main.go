package main

import(
	"fmt"
	"server/carte"
	testsA "server/alpha"
	"encoding/json"
	"server/joueur"

)

type Game struct{
	players []joueur.Joueur
	carte care.Carte
	GameRunning bool
}

func main() {
	loopBoolean:=true;
	mat:=carte.New(10)
	//carte.Debug(mat)
	//npc.Test(mat)

	tests.Test(mat)

	gameLoop(mat,&loopBoolean)
}

func (g *Game)EndOfGame(){
	(*g).GameRunning=false
}

func (g *Game)gameLoop(){
	for (*g).GameRunning{

	}

}

func (g *Game)ExtractGameData(){
	jsonFile, err := os.Open("GameData.json")
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

}
