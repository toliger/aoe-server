package main

import "fmt"
//import npc "server/npc"
import carte "server/carte"
import tests "server/test"
import "encoding/json"
import "io/ioutil"
import "server/joueur"

type Data struct{
	size int
	joueurs []joueur.Joueur
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

func ExtractData Game(){

}