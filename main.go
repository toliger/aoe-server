package main

import (
  "fmt"
  "git.unistra.fr/AOEINT/server/game"
  d "git.unistra.fr/AOEINT/server/data"
  client "git.unistra.fr/AOEINT/server/server"
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
	(&g).GameLoop()

	// Listen
	client.InitListenerServer(&g)
}
