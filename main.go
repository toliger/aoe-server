package main

import (
	d "git.unistra.fr/AOEINT/server/data"
	"git.unistra.fr/AOEINT/server/game"
	"git.unistra.fr/AOEINT/server/server"
	"log"
  "net/http"
        "time"

        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promauto"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
        go func() {
                for {
                        opsProcessed.Inc()
                        time.Sleep(2 * time.Second)
                }
        }()
}

var (
        opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
                Name: "players_count",
                Help: "The total number of players",
        })
)

func main() {

  recordMetrics()

  http.Handle("/metrics", promhttp.Handler())
  go http.ListenAndServe(":2112", nil)
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
	//for _,g.Buildings
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
