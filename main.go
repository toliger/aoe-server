package main

import "fmt"
//import npc "server/npc"
//import tests "server/test"
import "server/affichage"
import simulateClient "server/falseclient"
import "server/game"

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