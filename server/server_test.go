package server

import (

	"fmt"
	"log"
	"time"
	"context"
	"testing"
	"git.unistra.fr/AOEINT/server/data"
	pb "git.unistra.fr/AOEINT/server/grpc"
	"google.golang.org/grpc"
)

var client *grpc.Server

const (
	addrClientCo = "localhost:50010"
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJncm91cCI6InBsYXllciIsIm5hbWUiOiJQaWVycmUgQyIsInV1aWQiOiJiMzNkOTU0Zi1jNjNlLTRiNDgtODhlYi04YjVlODZkOTQyNDYiLCJpYXQiOjE1MTYyMzkwMjJ9.0btft-GVpqZSFvO_8o9qy5Nl9rNFgePXBfwz6bfR-P8"
)

func TestSayHello(t *testing.T) {
	
	var err error
	s := Arguments{}
	ctx ,cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("\nLancement des tests de SayHello")

	fmt.Println("1. Envoie et la réception")
	fmt.Println("Envoie d'un HelloRequest")
	_, err = s.SayHello(ctx, &pb.HelloRequest{});
	if err != nil {
		log.Println(err)
		t.Errorf("")
	}
	fmt.Println("Reception d'un HelloReply")

	fmt.Println("Validation des tests de SayHello")
}

func TestRightClick(t *testing.T) {
	
	var err error

	s := Arguments{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("\nLancement des tests de RightClick")

	fmt.Println("1. RightClick Vide")
	_, err = s.RightClick(ctx, &pb.RightClickRequest{});
	if err == nil { 
		t.Errorf("Erreur du test, message d'erreur vide")
	}

	fmt.Println("\n2. AskUpdate Token remplie (mal v1)")
	_, err = s.AskUpdate(ctx, &pb.AskUpdateRequest{
		Token: "eyJhbGciOiJIUzI1NiIs",
	});
	if err == nil {
		t.Errorf("Erreur du test, message d'erreur vide")
	}

	fmt.Println("\n2. AskUpdate Token remplie (mal v2)")
	_, err = s.AskUpdate(ctx, &pb.AskUpdateRequest{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM",
	});
	if err == nil {
		t.Errorf("Erreur du test, message d'erreur vide")
	}

	fmt.Println("\n2. AskUpdate Token remplie (mal v3)")
	_, err = s.AskUpdate(ctx, &pb.AskUpdateRequest{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxw",
	});
	if err == nil {
		t.Errorf("Erreur du test, message d'erreur vide")
	}

	fmt.Println("\n3. RightClick EntitySelectionUUID remplie (mais mal)")
	_, err = s.RightClick(ctx, &pb.RightClickRequest{
		Token: token,
		EntitySelectionUUID: []string{"5", "6"},
	})
	if err == nil {
		t.Errorf("Erreur du test, message d'erreur vide")
	}

	fmt.Println("\n4. RightClick EntitySelectionUUID et Target remplie")
	_, err = s.RightClick(ctx, &pb.RightClickRequest{
		Token: token,
		EntitySelectionUUID: []string{"5", "6"},
		Target: "152",
	})
	if err == nil {
		t.Errorf("Erreur du test, message d'erreur vide")
	}

	fmt.Println("\n5. RightClick tout remplie")
	_, err = s.RightClick(ctx, &pb.RightClickRequest{
		Token: token,
		Point: &pb.Coordinates{X: 5.2, Y: 7.3},
		EntitySelectionUUID: []string{"5", "6"},
		Target: "152",
	})
	if err == nil {
		t.Errorf("Erreur du test, message d'erreur vide")
	}

	fmt.Println("Validation des tests de RightClick")
}

func TestAskUpdate(t *testing.T) {
	
	var err error
	s := Arguments{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("\nLancement des tests de AskUpdate")

	fmt.Println("\n1. AskUpdate vide")
	_, err = s.AskUpdate(ctx, &pb.AskUpdateRequest{});
	if err == nil {
		t.Errorf("Erreur du test, message d'erreur vide")
	}

	fmt.Println("\n2. AskUpdate Token remplie (mal v1)")
	_, err = s.AskUpdate(ctx, &pb.AskUpdateRequest{
		Token: "eyJhbGciOiJIUzI1NiIs",
	});
	if err == nil {
		t.Errorf("Erreur du test, message d'erreur vide")
	}

	fmt.Println("\n2. AskUpdate Token remplie (mal v2)")
	_, err = s.AskUpdate(ctx, &pb.AskUpdateRequest{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM",
	});
	if err == nil {
		t.Errorf("Erreur du test, message d'erreur vide")
	}

	fmt.Println("\n2. AskUpdate Token remplie (mal v3)")
	_, err = s.AskUpdate(ctx, &pb.AskUpdateRequest{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxw",
	});
	if err == nil {
		t.Errorf("Erreur du test, message d'erreur vide")
	}
}

func TestAskCreation(t *testing.T) {
	
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
	data.InitiateActionBuffer()
	data.IDMap= data.NewObjectID()
	TestSayHello(&testing.T{})

	TestRightClick(&testing.T{})

	TestAskUpdate(&testing.T{})

	// TestAskCreation(&testing.T{})
}
