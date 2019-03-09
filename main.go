package main

import "fmt"
//import npc "server/npc"
//import tests "server/test"
import "git.unistra.fr/AOEINT/server/affichage"
import simulateClient "git.unistra.fr/AOEINT/server/falseclient"
import "git.unistra.fr/AOEINT/server/game"
//import tests "server/test"


func main() {
	var g game.Game
	g.GameRunning=true
	(&g).GetPlayerData()
	data:=game.ExtractData()
	(&g).GenerateMap(data)
	fmt.Println("Data struct extracted from json:",data)

	//On lance le faux client pour tester les fonctions de liaison
	go simulateClient.StartClient(&(g.GameRunning))
	affichage.ImprimerCarte(g.Carte)
	(&g).GameLoop()
}
