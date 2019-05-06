package main

import (
	d "git.unistra.fr/AOEINT/server/data"
	"git.unistra.fr/AOEINT/server/game"
	"git.unistra.fr/AOEINT/server/server"
)

func main() {
	d.ExtractFromToken("aubvfauipva.eyJncm91cCI6InBsYXllciIsIm5hbWUiOiJQaWVycmUgQyIsInV1aWQiOiJiMzNkOTU0Zi1jNjNlLTRiNDgtODhlYi04YjVlODZkOTQyNDYiLCJpYXQiOjE1MTYyMzkwMjJ9.oaougf")
	var g game.Game
	d.IDMap = d.NewObjectID()
	cExit := make(chan(bool))
	g.GameRunning = cExit

	// Listen
	server.InitListenerServer(&g)
	go initialize(&g)
}

func initialize(g *game.Game){
	g.GetPlayerData()
	d.InitiateActionBuffer()
	data := game.ExtractData()
	g.GenerateMap(data)
	go g.LaunchAutomaticFight()
	go g.BrokenBuildingsCollector()
	// On lance le faux client pour tester les fonctions de liaison
	go g.GameLoop()
}
