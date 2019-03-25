package main

import (
  "git.unistra.fr/AOEINT/server/game"
  d "git.unistra.fr/AOEINT/server/data"
  "git.unistra.fr/AOEINT/server/server"
)

func main() {
	var g game.Game
	d.IDMap=d.NewObjectID()
	d.InitiateActionBuffer()
	g.GameRunning=true
	(&g).GetPlayerData()
	data:=game.ExtractData()
	(&g).GenerateMap(data)

	// On lance le faux client pour tester les fonctions de liaison
	go (&g).GameLoop()

	// Listen
	server.InitListenerServer(&g)
}
