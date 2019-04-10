package server

import (
	
	"fmt"
	"log"
	"time"
	"context"
	"testing"

	pb "git.unistra.fr/AOEINT/server/grpc"
	"google.golang.org/grpc"
)

var client *grpc.Server

const (
	addrClientCo = "localhost:50010"
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJncm91cCI6InBsYXllciIsIm5hbWUiOiJQaWVycmUgQyIsInV1aWQiOiJiMzNkOTU0Zi1jNjNlLTRiNDgtODhlYi04YjVlODZkOTQyNDYiLCJpYXQiOjE1MTYyMzkwMjJ9.0btft-GVpqZSFvO_8o9qy5Nl9rNFgePXBfwz6bfR-P8"
)

func TestSayHello(t *testing.T) {
	
	s := Arguments{}
	ctx ,cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("Lancement des tests de SayHello")

	reply, err := s.SayHello(ctx, &pb.HelloRequest{});
	if err != nil {
		log.Println(err)
		t.Errorf("")
	}
	fmt.Println(reply)
}

func TestRightClick(t *testing.T) {
	
	s := Arguments{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("Lancement des tests de RightClick")

	reply, err := s.RightClick(ctx, &pb.RightClickRequest{});
	if err != nil {
		log.Println(err)
		t.Errorf("")
	}
	fmt.Println(reply)
}

func TestAskUpdate(t *testing.T) {
	
	s := Arguments{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("Lancement des tests de AskUpdate")

	reply, err := s.AskUpdate(ctx, &pb.AskUpdateRequest{});
	if err != nil {
		log.Println(err)
		t.Errorf("")
	}
	fmt.Println(reply)
}

func TestMain(m *testing.M) {

	TestSayHello(&testing.T{})

	TestRightClick(&testing.T{})

	TestAskUpdate(&testing.T{})
}
