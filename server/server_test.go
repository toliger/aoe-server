package server

import (
	
	"fmt"
	"log"
	"time"
	"context"
	//"flag"
	
	"net"
	//"os"
	"testing"

	pb "git.unistra.fr/AOEINT/server/grpc"
	"google.golang.org/grpc"
)

var client *grpc.Server

const (
	addrClientCo = "localhost:50010"
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJncm91cCI6InBsYXllciIsIm5hbWUiOiJQaWVycmUgQyIsInV1aWQiOiJiMzNkOTU0Zi1jNjNlLTRiNDgtODhlYi04YjVlODZkOTQyNDYiLCJpYXQiOjE1MTYyMzkwMjJ9.0btft-GVpqZSFvO_8o9qy5Nl9rNFgePXBfwz6bfR-P8"
)

func clientHello() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addrClientCo, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Salut"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("Reception d'un SayHello")

	log.Println(reply)
}

func clientRightClick() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addrClientCo, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewInteractionsClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := pb.AskUpdateRequest{Token: token}

	reply, err := c.AskUpdate(ctx, &request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("Reception d'un AskUpdateReply")

	log.Println(reply)
}

func clientInit() {
	lis, err := net.Listen("tcp", ":50020")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Print("Server listen !")
	}

	// Initialization of gRPC server
	arg := Arguments{}
	client = grpc.NewServer()

	// Registration of services Hello and Interactions
	pb.RegisterHelloServer(client, &arg)
	pb.RegisterInteractionsServer(client, &arg)

	// Make listen the server
	if err := client.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}

func TestMain(m *testing.M) {

	s := Arguments{}
	ctx := context.WithTimeout(context.Background(), time.Second)

	fmt.Println("Test de SayHello")
	s.SayHello(ctx, )

	fmt.Println("Test de RightClick")

	fmt.Println("Test de AskUpdate")

	/*go clientInit()
	var g game.Game
	go InitListenerServer(&g)
	flag.Parse()
	exitCode := m.Run()

	StopListenerServer()
	client.GracefulStop()
	os.Exit(exitCode)*/
}
