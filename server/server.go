// Package server :
// All for the clients interactions
package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"reflect"

	"git.unistra.fr/AOEINT/server/constants"
	"git.unistra.fr/AOEINT/server/data"
	"git.unistra.fr/AOEINT/server/game"
	"git.unistra.fr/AOEINT/server/utils"

	"git.unistra.fr/AOEINT/server/batiment"
	"git.unistra.fr/AOEINT/server/npc"
	"git.unistra.fr/AOEINT/server/ressource"

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

	// For Testing Mode
	if len(in.EntitySelectionUUID) == 0 {
		msg := "Erreur, aucune entité envoié à RightClick"
		log.Println(msg)
		return &pb.RightClickReply{}, errors.New(msg)
	}

	// Extract data from token to get player's UUID
	playerUUID := data.ExtractFromToken(in.Token)
	if playerUUID == nil {
		msg := "Token invalide dans AskUpdate"
		log.Print(msg)
		return &pb.RightClickReply{}, errors.New(msg)
	}
	
	if in.Target == "" { // MoveTo request
		// Loop on each entity
		for i := 0; i < len(in.EntitySelectionUUID); i++ {

			// Get the entity
			entity := data.IDMap.GetObjectFromID(in.EntitySelectionUUID[i])
			if entity == nil {
				msg := "Erreur, une entity n'est pas trouvé dans RightClick"
				log.Println(msg)
				return &pb.RightClickReply{}, errors.New(msg)
			}

			// Verify if the asker can move the NPC
			if entity.(*npc.Npc).PlayerUUID != playerUUID.UID {
				msg := "Erreur, une entity n'est pas au joueur"
				log.Println(msg)
				continue
			}

			go entity.(*npc.Npc).MoveTo(s.g.Carte, int(in.Point.X), int(in.Point.Y), nil)
			
		}

	} else { // Attack request
		// Loop on each entity
		for i := 0; i < len(in.EntitySelectionUUID); i++ {

			// Get the entities
			entity := data.IDMap.GetObjectFromID(in.EntitySelectionUUID[i])
			if entity == nil {
				msg := "Erreur, une entity n'est pas trouvé dans RightClick"
				log.Println(msg)
				return &pb.RightClickReply{}, errors.New(msg)
			}

			// Verify if the asker can move the NPC
			if entity.(*npc.Npc).PlayerUUID != playerUUID.UID {
				msg := "Erreur, une entity n'est pas au joueur"
				log.Println(msg)
				continue
			}

			target := data.IDMap.GetObjectFromID(in.Target)
			if target == nil {
				msg := "Erreur, target n'est pas trouvé dans RightClick"
				log.Println(msg)
				return &pb.RightClickReply{}, errors.New(msg)
			}

			switch reflect.TypeOf(target) {
			case reflect.TypeOf(&npc.Npc{}):
				go entity.(*npc.Npc).MoveTargetNpc(s.g.Carte, target.(*npc.Npc), nil)

			case reflect.TypeOf(&batiment.Batiment{}):
				go entity.(*npc.Npc).MoveTargetBuilding(s.g.Carte, target.(*batiment.Batiment), nil)

			case reflect.TypeOf(&ressource.Ressource{}):
				go entity.(*npc.Npc).MoveHarvestTarget(s.g.Carte, target.(*ressource.Ressource))

			default:
				msg := "Erreur, target est invalide dans RightClick"
				log.Println(msg)
				return &pb.RightClickReply{}, errors.New(msg)
			}
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
		msg := "Token invalide dans AskUpdate: vide"
		log.Print(msg)
		return &pb.AskUpdateReply{Array: nil}, errors.New(msg)
	}

	toSend := make([]*pb.UpdateAsked, 0)

	// Verify if the playerUUID exist in the map
	if _, isFilled := data.ActionBuffer[playerUUID.UID]; isFilled {
		for actionType := range data.ActionBuffer[playerUUID.UID] {
			for entityUUID := range data.ActionBuffer[playerUUID.UID][actionType].Description {
				upAsk := pb.UpdateAsked{Type: int32(actionType), EntityUUID: entityUUID}

				for key, value := range data.ActionBuffer[playerUUID.UID][actionType].Description[entityUUID] {
					upAsk.Arg = append(upAsk.Arg, &pb.Param{Key: key, Value: value})
				}
				//log.Print("envoyé: joueur: ",playerUUID.UID," type: ",actionType)
				toSend = append(toSend, &upAsk)
			}
		}
	} else {
		log.Print("PlayerUUID invalide dans AskUpdate: ",playerUUID.UID)
		log.Print("wanted:")
		for key:=range data.ActionBuffer{
			log.Print(key)
		}
		return &pb.AskUpdateReply{Array: nil}, nil
	}

	// Deletin the historic of updates waiting for the client
	data.CleanPlayerActionBuffer(playerUUID.UID)

	return &pb.AskUpdateReply{Array: toSend}, nil
}

// AskCreation :
// Function of creation of a building or NPC
func (s *Arguments) AskCreation(ctx context.Context, in *pb.AskCreationRequest) (*pb.AskCreationReply, error) {

	// For Debug Mode
	utils.Debug("Reception d'un AskCreationRequest et envoie d'un AskCreationReply")

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
		if in.TypeUnit == 0 {
			class = "villager"
		} else if in.TypeUnit == 1 {
			class = "harvester"
		} else if in.TypeUnit == 2 {
			class = "soldier"
		} else {
			log.Print("TypeUnit invalide dans AskCreation")
			return &pb.AskCreationReply{Validation: false}, nil
		}

		// Create NPC into the right player and update ActionBuffer
		player := s.g.GetPlayerFromUID(playerUUID.UID)
		player.AddAndCreateNpc(class, int(in.Case.X), int(in.Case.Y))
		fmt.Println(int(in.Case.X), int(in.Case.Y))

	case constants.ActionNewBuilding:

		// Define class asked and create it to the right player
		var class string
		if in.TypeUnit == 0 {
			class = "auberge"
		} else if in.TypeUnit == 1 {
			class = "caserne"
		} else if in.TypeUnit == 2 {
			class = "etabli"
		} else {
			log.Print("TypeUnit invalide dans AskCreation")
			return &pb.AskCreationReply{Validation: false}, nil
		}

		// Create NPC into the right player and update ActionBuffer
		player := s.g.GetPlayerFromUID(playerUUID.UID)
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
