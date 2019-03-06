//Contient toutes les fonctions pour les echanges de donnees client/serveur
package client

import (
	"git.unistra.fr/AOEINT/server/npc"
	/*"git.unistra.fr/AOEINT/server/carte"
	"git.unistra.fr/AOEINT/server/ressource"
	"git.unistra.fr/AOEINT/server/joueur"
	"git.unistra.fr/AOEINT/server/batiment"*/
	"git.unistra.fr/AOEINT/server/constants"
	"fmt"
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "git.unistra.fr/AOEINT/server/serveur"
	"strconv"
)

///////////////////////////////////////////////////////////////////////////////
// Général
///////////////////////////////////////////////////////////////////////////////

type ObjectId struct{
	IdOffset int
	IdArray map[string]*interface{}
}
func NewObjectID() ObjectId{
	return (ObjectId{0,make(map[string]*interface{},constants.MAXOBJECTS)})
}

func (o *ObjectId)AddObject(obj *interface{}){
	key:=strconv.Itoa((*o).IdOffset)
	(*o).IdArray[key]=obj
	(*o).IdOffset++
}

func (o *ObjectId) DeleteObject(id string){
	delete((*o).IdArray,id)
}

func (o *ObjectId) GetObjectFromId(id string) *interface{}{
	return (*o).IdArray[id]
}

var server *grpc.Server

type Server struct {}

// Fonction demarrant la gestion des intéractions gRPC
// Fonction bloquante, à lancer en concurrence
func InitListenerServer(adress string) {

	// Initialisation du socket d'écoute réseau
	lis, err := net.Listen("tcp", adress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initialisation du serveur gRPC
	server = grpc.NewServer()

	// Enregistement des services Hello, Map et Interactions sur le serveur
	pb.RegisterHelloServer(server, &Server{})
	pb.RegisterInteractionsServer(server, &Server{})

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
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Reception d'un HelloRequest et envoie d'un HelloReply")
	return &pb.HelloReply{}, nil
}

// Fonction du service Interactions: RightClick
func (s *Server) RightClick(ctx context.Context, in *pb.RightClickRequest) (*pb.RightClickReply, error) {
	fmt.Println("Reception d'un RightClickRequest et envoie d'un RightClickReply")
	return &pb.RightClickReply{}, nil
}

// Fonction du service Interactions: AskUpdate
func (s *Server) AskUpdate(ctx context.Context, in *pb.AskUpdateRequest) (*pb.AskUpdateReply, error) {
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
