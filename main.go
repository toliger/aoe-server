package main

import (
  "fmt"
  "git.unistra.fr/AOEINT/server/affichage"
  "git.unistra.fr/AOEINT/server/falseclient"
  "git.unistra.fr/AOEINT/server/game"
  "git.unistra.fr/AOEINT/server/data"
  "git.unistra.fr/AOEINT/server/client"
)

func main() {
	var g game.Game
	data.IdMap=data.NewObjectID()
	g.GameRunning=true
	(&g).GetPlayerData()
	data:=game.ExtractData()
	(&g).GenerateMap(data)
	fmt.Println("Data struct extracted from json:",data)

	// On lance le faux client pour tester les fonctions de liaison
	go falseclient.StartClient(&(g.GameRunning))
	affichage.ImprimerCarte(g.Carte)
	(&g).GameLoop()

  // Listen
  client.InitListenerServer()
}
