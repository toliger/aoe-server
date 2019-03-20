// Package server :
// All for the clients interactions
package server

import (
	"git.unistra.fr/AOEINT/server/npc"
	"git.unistra.fr/AOEINT/server/data"
	"git.unistra.fr/AOEINT/server/game"

	/*"git.unistra.fr/AOEINT/server/ressource"
	"git.unistra.fr/AOEINT/server/joueur"
	"git.unistra.fr/AOEINT/server/batiment"*/
	"fmt"
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "git.unistra.fr/AOEINT/server/grpc"
)

///////////////////////////////////////////////////////////////////////////////
// General
///////////////////////////////////////////////////////////////////////////////

var server *grpc.Server

// Arguments :
// Data structure used in the gRPC method's
type Arguments struct {
	g *game.Game
	UpdateBuffer []pb.UpdateAsked
}

// InitListenerServer :
//	Function starting gRPC interactions
//	Blocking function
func InitListenerServer(g *game.Game) {

	// Initialization of the socket
	// TODO Use of a variable of environment to set de port
	lis, err := net.Listen("tcp", ":50010")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initialization of gRPC server
	arg := Arguments{g: g}
	server = grpc.NewServer()

	// Registration of services Hello and Interactions
	pb.RegisterHelloServer(server, &arg)
	pb.RegisterInteractionsServer(server, &arg)

	// Make listen the server
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}

// StopListenerServer :
// Function stopping the gRPC interactions (clean stop)
func StopListenerServer() {
	server.GracefulStop()
}

// KillListenerServer :
// Function stopping the gRPC interactions (dirty stop)
func KillListenerServer() {
	server.Stop()
}

///////////////////////////////////////////////////////////////////////////////
// Client -> Server
///////////////////////////////////////////////////////////////////////////////

// SayHello :
// Function of the service Hello: SayHello
func (s *Arguments) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Reception d'un HelloRequest et envoie d'un HelloReply")
	return &pb.HelloReply{}, nil
}

// RightClick :
// Function of the service Interactions: RightClick
func (s *Arguments) RightClick(ctx context.Context, in *pb.RightClickRequest) (*pb.RightClickReply, error) {

	// Loop on each entity
	var tmpCoord []*pb.Coordinates
	sendPath := make(map[string]*pb.RPCoordinates, len(in.EntitySelectionUUID))
	for i:=0 ; i<len(in.EntitySelectionUUID) ; i++ {

		// Get the entity
		entity := data.IDMap.GetObjectFromID(in.EntitySelectionUUID[i]).(*npc.Npc)

		// Get the path of the entity
		path := entity.MoveTo(s.g.Carte, int(in.Point.X), int(in.Point.Y), nil)

		// Creating the array for the message
		if len(path) != 0 {
			tmpCoord = make([]*pb.Coordinates, len(path))
			tmp := make([]pb.Coordinates, len(path))
			for j:=0 ; j<len(path) ; j++ {
				tmp[j].X = int32(path[j].GetPathX())
				tmp[j].Y = int32(path[j].GetPathY())
				tmpCoord[j] = &tmp[j]
			}
		} else {
			tmpCoord = make([]*pb.Coordinates, 1)
			tmpCoord[0] = &pb.Coordinates{X: int32(-1), Y: int32(-1)}
		}

		// Linking the array to the message
		sendPath[in.EntitySelectionUUID[i]] = &pb.RPCoordinates{Coord: tmpCoord}
	}

	// Put data in UpdateBuffer
	for uuid, rpcoord := range sendPath {
		lenght := len(rpcoord.Coord)
		if lenght > 1 {
			tmp := pb.UpdateAsked{Type: 3, EntityUUID: uuid}
			tmp.Arg = make([]*pb.Param, 0)
			tmp.Arg = append(tmp.Arg, &pb.Param{Key: "x",Value: string(rpcoord.Coord[0].X)})
			tmp.Arg = append(tmp.Arg, &pb.Param{Key: "y",Value: string(rpcoord.Coord[0].Y)})
			tmp.Arg = append(tmp.Arg, &pb.Param{Key: "xDest",Value: string(rpcoord.Coord[lenght-1].X)})
			tmp.Arg = append(tmp.Arg, &pb.Param{Key: "yDest",Value: string(rpcoord.Coord[lenght-1].Y)})

			s.UpdateBuffer = append(s.UpdateBuffer, tmp)
		}
	}

	return &pb.RightClickReply{Path: sendPath}, nil
}

// AskUpdate :
// Function of the service Interactions: AskUpdate
func (s *Arguments) AskUpdate(ctx context.Context, in *pb.AskUpdateRequest) (*pb.AskUpdateReply, error) {

	toSend := make([]*pb.UpdateAsked, 0)

	if s.UpdateBuffer != nil {
		for i:=0 ; i<len(s.UpdateBuffer) ; i++ {
			toSend = append(toSend, &s.UpdateBuffer[i])
		}
	}

	return &pb.AskUpdateReply{Array: toSend}, nil
}
