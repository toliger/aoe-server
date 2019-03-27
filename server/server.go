// Package server :
// All for the clients interactions
package server


import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	"git.unistra.fr/AOEINT/server/utils"
	"git.unistra.fr/AOEINT/server/npc"
	"git.unistra.fr/AOEINT/server/data"
	"git.unistra.fr/AOEINT/server/game"
	"git.unistra.fr/AOEINT/server/constants"
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
	} else {
    log.Print("Server listen !")
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
	utils.Debug("Reception d'un HelloRequest et envoie d'un HelloReply")
	return &pb.HelloReply{}, nil
}


// RightClick :
// Function of the service Interactions: RightClick
func (s *Arguments) RightClick(ctx context.Context, in *pb.RightClickRequest) (*pb.RightClickReply, error) {

	// Loop on each entity
	for i:=0 ; i<len(in.EntitySelectionUUID) ; i++ {

		// Get the entity
		entity := data.IDMap.GetObjectFromID(in.EntitySelectionUUID[i]).(*npc.Npc)

		// Get the path of the entity
		path := entity.MoveTo(s.g.Carte, int(in.Point.X), int(in.Point.Y), nil)

		// Filling ActionBuffer with the right data
		entityData := entity.Stringify()
		data.AddToAllAction(constants.ActionHarmNpc, in.EntitySelectionUUID[i], "pv", entityData["pv"])
		data.AddToAllAction(constants.ActionHarmNpc, in.EntitySelectionUUID[i], "x", entityData["x"])
		data.AddToAllAction(constants.ActionHarmNpc, in.EntitySelectionUUID[i], "y", entityData["y"])
		data.AddToAllAction(constants.ActionHarmNpc, in.EntitySelectionUUID[i], "vitesse", entityData["vitesse"])
		data.AddToAllAction(constants.ActionHarmNpc, in.EntitySelectionUUID[i], "damage", entityData["damage"])
		data.AddToAllAction(constants.ActionHarmNpc, in.EntitySelectionUUID[i], "vue", entityData["vue"])
		data.AddToAllAction(constants.ActionHarmNpc, in.EntitySelectionUUID[i], "portee", entityData["portee"])
		data.AddToAllAction(constants.ActionHarmNpc, in.EntitySelectionUUID[i], "pv", entityData["pv"])
		if len(path) != 0 {
			data.AddToAllAction(constants.ActionHarmNpc, in.EntitySelectionUUID[i], "destX", string(path[len(path)-1].GetPathX()))
			data.AddToAllAction(constants.ActionHarmNpc, in.EntitySelectionUUID[i], "destY",  string(path[len(path)-1].GetPathY()))
		} else {
			data.AddToAllAction(constants.ActionHarmNpc, in.EntitySelectionUUID[i], "destX", "-1")
			data.AddToAllAction(constants.ActionHarmNpc, in.EntitySelectionUUID[i], "destY",  "-1")
		}
	}

	return &pb.RightClickReply{}, nil
}


// AskUpdate :
// Function of the service Interactions: AskUpdate
func (s *Arguments) AskUpdate(ctx context.Context, in *pb.AskUpdateRequest) (*pb.AskUpdateReply, error) {

	playerUUID := data.ExtractFromToken(in.Token)
	if playerUUID == nil {
		log.Print("Token invalide dans AskUpdate")
		return &pb.AskUpdateReply{Array: nil}, nil
	}

	toSend := make([]*pb.UpdateAsked, 0)

	// Verify if the playerUUID exist in the map
	if _, isFilled := data.ActionBuffer[playerUUID.UUID] ; isFilled {
		for actionType := range data.ActionBuffer[playerUUID.UUID] {
			for entityUUID := range data.ActionBuffer[playerUUID.UUID][actionType].Description {
				upAsk := pb.UpdateAsked{Type: int32(actionType), EntityUUID: entityUUID}

				for key, value := range data.ActionBuffer[playerUUID.UUID][actionType].Description[entityUUID] {
					upAsk.Arg = append(upAsk.Arg, &pb.Param{Key: key, Value: value})
				}

				toSend = append(toSend, &upAsk)
			}
		}
	} else {
		log.Print("PlayerUUID invalide dans AskUpdate")
		return &pb.AskUpdateReply{Array: nil}, nil
	}

	// Deletin the historic of updates waiting for the client
	data.CleanPlayerActionBuffer(playerUUID.UUID)

	return &pb.AskUpdateReply{Array: toSend}, nil
}
