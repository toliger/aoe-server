package main

import (
  "fmt"
  "git.unistra.fr/AOEINT/server/affichage"
  simulateClient "git.unistra.fr/AOEINT/server/falseclient"
  "git.unistra.fr/AOEINT/server/game"
  "git.unistra.fr/AOEINT/server/data"

  "context"
  "log"
  "net"
  "google.golang.org/grpc"
  pb "git.unistra.fr/AOEINT/server/serveur"
)

const (
	port = ":50067"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received !")
	return &pb.HelloReply{}, nil
}

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

  // Listen
  lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
    log.Print("Server listen on ", port)
  }
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
  }
}
