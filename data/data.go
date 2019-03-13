package data

import "strconv"
import "git.unistra.fr/AOEINT/server/constants"

type Action struct{
	Description map[string]map[string]string
}
//variable détaillant les actions à envoyer au client 
//	Exemple: [type:int].Description["UUID"]["Key"]="value"  
var ActionBuffer []Action

//Initialisation du buffer d'actions
func InitiateActionBuffer(){
	ActionBuffer=make([]Action,constants.MAXACTIONS)
}
//Ajoute une Action(type int, clee string, description string) au buffer
func AddNewAction(typ int,uuid string, key string, description string){
	
	elem, ok := ActionBuffer[typ].Description[uuid]
    if !ok {
       elem = make(map[string]string)
	   if(ActionBuffer[typ].Description == nil){
			ActionBuffer[typ].Description=make(map[string]map[string]string)
	   }
       ActionBuffer[typ].Description[uuid] = elem
    }
	ActionBuffer[typ].Description[uuid][key]=description
}

func CleanActionBuffer(){
	ActionBuffer=nil //throw to garbage collector
	InitiateActionBuffer()
}

//Structure générique associant chaque batiment/ressource/pnj à son id
type ObjectId struct{
	IdOffset int
	IdArray map[string]interface{}
}
//Crée une instance ObjectId
func NewObjectID() ObjectId{
	res:=(ObjectId{0,nil})
	res.IdArray=make(map[string]interface{},constants.MAXOBJECTS)
	return res
}

var IdMap ObjectId

//Fonction  permettant d'ajouter un objet générique à ObjectId. Retourne l'id de l'objet
func (o *ObjectId)AddObject(obj interface{}) string{
	key:=strconv.Itoa((*o).IdOffset)
	(*o).IdArray[key]=obj
	(*o).IdOffset++
	return key
}

//Fonction permettant de retirer un objet à partir de son id
func (o *ObjectId) DeleteObjectFromId(id string){
	delete((*o).IdArray,id)
}

//Retire un objet de la liste à partir de son propre pointeur
func (o *ObjectId) DeleteObject(obj interface{}) bool{
	for i,e := range (*o).IdArray{
		if(e==obj){
			delete((*o).IdArray,i)
			return true
		}
	}
	return false
}

//Renvoie un pointeur sur l'obj correspondant à l'id fourni
func (o *ObjectId) GetObjectFromId(id string) interface{}{
	return (*o).IdArray[id]
}

//Renvoie l'id d'un objet à partir de son pointeur
func (o *ObjectId) GetIdFromObject(obj interface{}) string{
	for i,e:=range (*o).IdArray{
		if(e==obj){
			return i
		}
	}
	return "-1"
}