package main

import (
  "git.unistra.fr/AOEINT/server/game"
  d "git.unistra.fr/AOEINT/server/data"
  "git.unistra.fr/AOEINT/server/server"
)

func main() {
	d.ExtractFromToken("aubvfauipva.eyJncm91cCI6InBsYXllciIsIm5hbWUiOiJQaWVycmUgQyIsInV1aWQiOiJiMzNkOTU0Zi1jNjNlLTRiNDgtODhlYi04YjVlODZkOTQyNDYiLCJpYXQiOjE1MTYyMzkwMjJ9.oaougf")
	var g game.Game
	d.IDMap=d.NewObjectID()
	d.InitiateActionBuffer()
	g.GameRunning=true
	(&g).GetPlayerData()
	d.InitiateActionBuffer()
	data:=game.ExtractData()
	(&g).GenerateMap(data)

	// On lance le faux client pour tester les fonctions de liaison
	go (&g).GameLoop()

	// Listen
	server.InitListenerServer(&g)
}
