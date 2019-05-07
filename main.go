package main

import (
	d "git.unistra.fr/AOEINT/server/data"
	"git.unistra.fr/AOEINT/server/game"
	"git.unistra.fr/AOEINT/server/server"
	"log"
)

func main() {
	d.ExtractFromToken("aubvfauipva.eyJncm91cCI6InBsYXllciIsIm5hbWUiOiJQaWVycmUgQyIsInV1aWQiOiJiMzNkOTU0Zi1jNjNlLTRiNDgtODhlYi04YjVlODZkOTQyNDYiLCJpYXQiOjE1MTYyMzkwMjJ9.oaougf")
	var g game.Game
	d.IDMap = d.NewObjectID()
	cExit := make(chan(bool))
	g.GameRunning = cExit
	g.GetPlayerData()
	d.InitiateActionBuffer()
	data := game.ExtractData()
	g.GenerateMap(data)
	go g.LaunchAutomaticFight()
	go g.BrokenBuildingsCollector()
	// Listen
	go g.GameLoop()
	log.Println(g.Joueurs[0].UID)
	log.Println(g.Joueurs[1].UID)
	server.InitListenerServer(&g)
	//go initialize(&g)
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
