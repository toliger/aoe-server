package main

import "fmt"
import npc "server/npc"
import carte "server/carte"
import testsAlpha "server/alpha"
func main() {
	loopBoolean:=false;
	mat:=carte.New(10)
	//carte.Debug(mat)
	fmt.Println("test")
	npc.Test()
	fmt.Println("La case 0 0 est elle libre ?")
	if mat.IsEmpty(0,0) {
		fmt.Println("oui")
	} else {
		fmt.Println("non")
	}

	testsAlpha.AlphaTest(mat)

	gameLoop(mat,&loopBoolean)
}

func gameLoop(Terrain carte.Carte, gameRunning *bool){
	for *gameRunning{

	}

}
