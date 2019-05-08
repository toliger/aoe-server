package main

import (
	"log"
	d "git.unistra.fr/AOEINT/server/data"
	"git.unistra.fr/AOEINT/server/game"
	"git.unistra.fr/AOEINT/server/server"
	"git.unistra.fr/AOEINT/server/constants"
)

func main() {
	var g game.Game
	g.GameInitialisationTime=constants.ExpiringTime
	g.GameTimeLeft=constants.MaxGameTime
	g.BeginGame=make(chan(bool))
	g.BeginTimer=make(chan(bool))
	d.IDMap = d.NewObjectID()
	cExit := make(chan(bool))
	g.GameRunning = cExit
	/*g.GetPlayerData()
	d.InitiateActionBuffer()
	data := game.ExtractData()
	g.GenerateMap(data)
	go g.LaunchAutomaticFight()
	go g.BrokenBuildingsCollector()*/
	// Listen
	//go g.GameLoop()
	go server.InitListenerServer(&g)
	log.Println("En attente des joueurs")
	<-g.BeginGame //remplacer par g.GameInitialisationTime=-1 pour les tests solo en local
	log.Println("Démarrage de la partie")
	go startGame(&g)
}

func startGame(g *game.Game){
	d.InitiateActionBuffer()
	data := game.ExtractData()
	g.GenerateMap(data)
	go g.LaunchAutomaticFight()
	go g.BrokenBuildingsCollector()
	go g.GameLoop()
}