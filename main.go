package main

import "fmt"
import npc "server/npc"
import carte "server/carte"
import testsAlpha "server/alpha"
func main() {
	loopBoolean:=false;
	mat:=carte.New(4)
	carte.Debug(mat)
	fmt.Println("test")
	npc.Test()
	testsAlpha.AlphaTest()
	
	gameLoop(mat,&loopBoolean)
}

func gameLoop(Terrain carte.Carte, gameRunning *bool){
	for *gameRunning{
		
	}
	
}