package main

import "fmt"
import "git.unistra.fr/AOEINT/server/affichage"
import simulateClient "git.unistra.fr/AOEINT/server/falseclient"
import "git.unistra.fr/AOEINT/server/game"
import "git.unistra.fr/AOEINT/server/data"

func main() {
	var g game.Game
	data.IdMap=data.NewObjectID()
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
