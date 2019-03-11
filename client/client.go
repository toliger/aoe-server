//Contient toutes les fonctions pour les echanges de donnees client/serveur
package client

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
	pb "git.unistra.fr/AOEINT/server/serveur"
)

///////////////////////////////////////////////////////////////////////////////
// Général
///////////////////////////////////////////////////////////////////////////////

var server *grpc.Server

type ServerArguments struct {
	g *game.Game 
}

// Fonction demarrant la gestion des intéractions gRPC
// Fonction bloquante, à lancer en concurrence
func InitListenerServer(g *game.Game) {

	// Initialisation du socket d'écoute réseau
	// TODO Utiliser une variable d'environement pour pouvoir redéfinir le port
	lis, err := net.Listen("tcp", ":50010")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initialisation du serveur gRPC
	arg := ServerArguments{g: g}
	server = grpc.NewServer()

	// Enregistement des services Hello, Map et Interactions sur le serveur
	pb.RegisterHelloServer(server, &arg)
	pb.RegisterInteractionsServer(server, &arg)

	// Mise en écoute du serveur
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}

// Fonction arrêtant la gestion des intéractions gRPC (arrêt propre)
func StopListenerServer() {
	server.GracefulStop()
}

// Fonction arrêtant la gestion des intéractions gRPC (arrêt brutal)
func KillListenerServer() {
	server.Stop()
}


///////////////////////////////////////////////////////////////////////////////
// Serveur -> Client
///////////////////////////////////////////////////////////////////////////////

//Envoie toutes les donnees necessaires à la mise en place de la partie en debut de jeu
//A envoyer: donnees des joueurs, structure data(map), entites de depart..
func InitGameState() {}

//Maj les ressources du joueur à partir de l'uid correspondant
func updatePlayerRessources(playerUID string,stone int,wood int,food int){}

//Maj: Indique la destruction d'un Batiment au client pour qu'il soit retire
func BuildingDestroyed(playerUID string,x int, y int){

}

//Permet de Maj la liste des npcs visibles en indiquant leur mort au client
func PlayerNpcsDied(playerUID string,npc []npc.Npc){

}

///////////////////////////////////////////////////////////////////////////////
// Client -> Serveur
///////////////////////////////////////////////////////////////////////////////

// Fonction du service Hello: SayHello
func (s *ServerArguments) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Reception d'un HelloRequest et envoie d'un HelloReply")
	return &pb.HelloReply{}, nil
}

// Fonction du service Interactions: RightClick
func (s *ServerArguments) RightClick(ctx context.Context, in *pb.RightClickRequest) (*pb.RightClickReply, error) {

	// Boucle pour chaque entité
	sendPath := make(map[string]*pb.RPCoordinates, len(in.EntitySelectionUUID))
	for i:=0 ; i<len(in.EntitySelectionUUID) ; i++ {

		// Obtention de l'entité
		entity := data.IdMap.GetObjectFromId(in.EntitySelectionUUID[i]).(*npc.Npc)

		// Obtention du path pour l'entité
		path := entity.MoveTo(s.g.Carte, int(in.Point.X), int(in.Point.Y), nil)

		// Affectation au tableau
		tmpCoord := make([]*pb.Coordinates, len(path))
		tmp := make([]pb.Coordinates, len(path))
		for j:=0 ; j<len(path) ; j++ {
			tmp[j].X = int32(path[j].GetPathX())
			tmp[j].Y = int32(path[j].GetPathY())
			tmpCoord[j] = &tmp[j]
		}

		// Lien au reply
		sendPath[in.EntitySelectionUUID[i]] = &pb.RPCoordinates{Coord: tmpCoord}
	}

	return &pb.RightClickReply{Path: sendPath}, nil
}

// Fonction du service Interactions: AskUpdate
func (s *ServerArguments) AskUpdate(ctx context.Context, in *pb.AskUpdateRequest) (*pb.AskUpdateReply, error) {
	fmt.Println("Reception d'un AskUpdateRequest et envoie d'un AskUpdateReply")
	return &pb.AskUpdateReply{}, nil
}

//demande la creation d'un batiment à partir de l'uid du joueur, une position et un type de batiment
//class: "auberge","caserne","etabli"
func TryToBuild(playerUID string, x int, y int, class string) bool{
	return false
}

//Demande le deplacement des npc selectionnes
func MoveSelectedNpc(playerUID string, liste []npc.Npc, x int, y int){

}

//Demande la suppression par le joueur de l'un de ses batiments
func EraseBuilding(playerUID string, x int, y int){

}

//Averti le serveur de la creation d'une entite: verification des ressources necessaires
func AddNewNpc(playerUID string, x int, y int, typ int) bool{
	return false
}
//Enleve des Pv a un batiment
func  DamageBuilding(playerUID string, x int, y int, attack int){

}
