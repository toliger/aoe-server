package server

import (
	"flag"
	"log"
	"net"
	"os"
	"testing"

	pb "git.unistra.fr/AOEINT/server/grpc"
	"google.golang.org/grpc"
)

var client *grpc.Server

func clientInit() {
	lis, err := net.Listen("tcp", ":50011")
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

func TestServerInit(t *testing.T) {
	InitListenerServer(nil)
}

func TestMain(m *testing.M) {
	clientInit()
	flag.Parse()
	exitCode := m.Run()

	StopListenerServer()
	client.GracefulStop()
	os.Exit(exitCode)
}
