// Package server :
// All for the clients interactions
package server

import (
	"context"
	"fmt"
	"log"
	"net"


	"git.unistra.fr/AOEINT/server/utils"
	"git.unistra.fr/AOEINT/server/constants"
	"git.unistra.fr/AOEINT/server/data"
	"git.unistra.fr/AOEINT/server/game"

	"git.unistra.fr/AOEINT/server/npc"
	"git.unistra.fr/AOEINT/server/batiment"

	pb "git.unistra.fr/AOEINT/server/grpc"
	"google.golang.org/grpc"
)

///////////////////////////////////////////////////////////////////////////////
// General
///////////////////////////////////////////////////////////////////////////////

var server *grpc.Server

// Arguments :
// Data structure used in the gRPC method's
type Arguments struct {
	g *game.Game
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
	// For Debug Mode
	utils.Debug("Reception d'un HelloRequest et envoie d'un HelloReply")
	return &pb.HelloReply{}, nil
}

// RightClick :
// Function of the service Interactions: RightClick
func (s *Arguments) RightClick(ctx context.Context, in *pb.RightClickRequest) (*pb.RightClickReply, error) {

	// For Debug Mode
	utils.Debug("Reception d'un RightClickRequest et envoie d'un RightClickReply")

	// Loop on each entity
	for i := 0; i < len(in.EntitySelectionUUID); i++ {

		// Get the entity
		entity := data.IDMap.GetObjectFromID(in.EntitySelectionUUID[i]).(*npc.Npc)

		// Get the path of the entity
		path := entity.MoveTo(s.g.Carte, int(in.Point.X), int(in.Point.Y), nil, nil)

		// Filling ActionBuffer with the right data
		entityData := entity.Stringify()
		data.AddToAllAction(constants.ActionAlterationNpc, in.EntitySelectionUUID[i], "pv", entityData["pv"])
		data.AddToAllAction(constants.ActionAlterationNpc, in.EntitySelectionUUID[i], "x", entityData["x"])
		data.AddToAllAction(constants.ActionAlterationNpc, in.EntitySelectionUUID[i], "y", entityData["y"])
		data.AddToAllAction(constants.ActionAlterationNpc, in.EntitySelectionUUID[i], "vitesse", entityData["vitesse"])
		data.AddToAllAction(constants.ActionAlterationNpc, in.EntitySelectionUUID[i], "damage", entityData["damage"])
		data.AddToAllAction(constants.ActionAlterationNpc, in.EntitySelectionUUID[i], "vue", entityData["vue"])
		data.AddToAllAction(constants.ActionAlterationNpc, in.EntitySelectionUUID[i], "portee", entityData["portee"])
		data.AddToAllAction(constants.ActionAlterationNpc, in.EntitySelectionUUID[i], "pv", entityData["pv"])
		if len(path) != 0 {
			data.AddToAllAction(constants.ActionAlterationNpc, in.EntitySelectionUUID[i], "destX", fmt.Sprintf("%f", in.Point.X))
			data.AddToAllAction(constants.ActionAlterationNpc, in.EntitySelectionUUID[i], "destY", fmt.Sprintf("%f", in.Point.Y))
		} else {
			data.AddToAllAction(constants.ActionAlterationNpc, in.EntitySelectionUUID[i], "destX", "-1")
			data.AddToAllAction(constants.ActionAlterationNpc, in.EntitySelectionUUID[i], "destY", "-1")
		}
	}

	return &pb.RightClickReply{}, nil
}

// AskUpdate :
// Function of the service Interactions: AskUpdate
func (s *Arguments) AskUpdate(ctx context.Context, in *pb.AskUpdateRequest) (*pb.AskUpdateReply, error) {

	// For Debug Mode
	utils.Debug("Reception d'un AskUpdateRequest et envoie d'un AskUpdateReply")

	// Extract data from token to get player's UUID
	playerUUID := data.ExtractFromToken(in.Token)
	if playerUUID == nil {
		log.Print("Token invalide dans AskUpdate")
		return &pb.AskUpdateReply{Array: nil}, nil
	}

	toSend := make([]*pb.UpdateAsked, 0)

	// Verify if the playerUUID exist in the map
	if _, isFilled := data.ActionBuffer[playerUUID.UUID]; isFilled {
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

// AskCreation :
// Function of creation of a building or NPC
func (s *Arguments) AskCreation(ctx context.Context, in *pb.AskCreationRequest) (*pb.AskCreationReply, error) {
	
	// Extract data from token to get player's UUID
	playerUUID := data.ExtractFromToken(in.Token)
	if playerUUID == nil {
		log.Print("Token invalide dans AskCreation")
		return &pb.AskCreationReply{Validation: false}, nil
	}

	actionType := in.Type
	switch actionType {
	case constants.ActionNewNpc:
		
		// Define class asked
		var class string
		if (in.TypeUnit == 0) {
			class = "villager"
		} else if (in.TypeUnit == 1) {
			class = "harvester"
		} else if (in.TypeUnit == 2) {
			class = "soldier"
		} else {
			log.Print("TypeUnit invalide dans AskCreation")
			return &pb.AskCreationReply{Validation: false}, nil
		}

		// Create NPC into the right player and update ActionBuffer
		player := s.g.GetPlayerFromUID(playerUUID.UUID)
		player.AddAndCreateNpc(class, int(in.Case.X), int(in.Case.Y))
		fmt.Println(int(in.Case.X), int(in.Case.Y))

	case constants.ActionNewBuilding:

		// Define class asked and create it to the right player
		var class string
		if (in.TypeUnit == 0) {
			class = "auberge"
		} else if (in.TypeUnit == 1) {
			class = "caserne"
		} else if (in.TypeUnit == 2) {
			class = "etabli"
		} else {
			log.Print("TypeUnit invalide dans AskCreation")
			return &pb.AskCreationReply{Validation: false}, nil
		}

		// Create NPC into the right player and update ActionBuffer
		player := s.g.GetPlayerFromUID(playerUUID.UUID)
		b := batiment.Create(class, int(in.Case.X), int(in.Case.Y))
		if s.g.Carte.AddNewBuilding(&b) != true {
			log.Print("Erreur, peut pas créer un batiment dans AskCreation")
			return &pb.AskCreationReply{Validation: false}, nil
		}
		player.AddBuilding(&b)


	default:
		log.Print("Format de requête invalide dans AskCreation")
		log.Print("Voici in : ", in)
		return &pb.AskCreationReply{Validation: false}, nil
	}

	return &pb.AskCreationReply{Validation: true}, nil
}