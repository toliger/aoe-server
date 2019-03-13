package main

import (
  "fmt"
  "git.unistra.fr/AOEINT/server/falseclient"
  "git.unistra.fr/AOEINT/server/game"
  d "git.unistra.fr/AOEINT/server/data"
  "git.unistra.fr/AOEINT/server/client"
)

func main() {
	var g game.Game
	d.IdMap=d.NewObjectID()
	d.InitiateActionBuffer()
	g.GameRunning=true
	(&g).GetPlayerData()
	data:=game.ExtractData()
	(&g).GenerateMap(data)
	fmt.Println("Data struct extracted from json:",data)
	fmt.Println("buffer",d.ActionBuffer)
	// On lance le faux client pour tester les fonctions de liaison
	go falseclient.StartClient(&(g.GameRunning))
	(&g).GameLoop()

	// Listen
	client.InitListenerServer(&g)
}
