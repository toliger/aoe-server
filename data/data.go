package data

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"strconv"
	"strings"
	"git.unistra.fr/AOEINT/server/constants"
	"git.unistra.fr/AOEINT/server/utils"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/machinebox/graphql"
	"context"
)

//Action classe detaillant une action de ActionBuffer
type Action struct {
	Description map[string]map[string]string
}

//ActionBuffer variable détaillant les actions à envoyer au client
//	Exemple: [type:int].Description["UUID"]["Key"]="value"
// Modification: ["PlayerUID"]->[type:int].Description["UUID"]["Key"]="value"
var ActionBuffer map[string]([]Action)
//InitiateActionBuffer Initialisation du buffer d'actions
func InitiateActionBuffer() {
	ActionBuffer = make(map[string]([]Action), 4)
	ActionBuffer[constants.PlayerUID1] = make([]Action, constants.MaxActions)
	ActionBuffer[constants.PlayerUID2] = make([]Action, constants.MaxActions)
	if constants.PlayerUID3 != "DEFAULT" {
		ActionBuffer[constants.PlayerUID3] = make([]Action, constants.MaxActions)
	}
	if constants.PlayerUID4 != "DEFAULT" {
		ActionBuffer[constants.PlayerUID4] = make([]Action, constants.MaxActions)
	}
	if(actionChannel ==nil){
		actionChannel= make(chan request,constants.ActionChannelSize)
		go bufferLoop()
	}
}

type request struct{
	typ int
	uuid string
	key string
	description string
}

var actionChannel chan(request)

//AjoutConcurrent Permet d'effectuer un AddToAllAction de manière
func AjoutConcurrent(typ int, uuid string, key string, description string){
	req := request{typ,uuid,key,description}
	actionChannel <- req
}

func bufferLoop(){
	for {
		req:= <- actionChannel
		if(req.typ == -1){
			break
		}
		AddToAllAction(req.typ,req.uuid,req.key,req.description)
	}
}

//AddNewAction Ajoute une Action(type int, clee string, description string) au buffer
func AddNewAction(PlayerUID string, typ int, uuid string, key string, description string) {
	elem, ok := ActionBuffer[PlayerUID][typ].Description[uuid]
	if !ok {
		elem = make(map[string]string)
		if ActionBuffer[PlayerUID][typ].Description == nil {
			ActionBuffer[PlayerUID][typ].Description = make(map[string]map[string]string)
		}
		ActionBuffer[PlayerUID][typ].Description[uuid] = elem
	}
	ActionBuffer[PlayerUID][typ].Description[uuid][key] = description
}

//AddToAllAction Ajoute une action pour tous les joueurs
func AddToAllAction(typ int, uuid string, key string, description string) {
	for k := range ActionBuffer {
		AddNewAction(k, typ, uuid, key, description)
	}
}

//CleanActionBuffer vide le buffer ActionBuffer
func CleanActionBuffer() {
	ActionBuffer = nil //throw to garbage collector
	InitiateActionBuffer()
}

//CleanPlayerActionBuffer vide le buffer du joueur correspondant
func CleanPlayerActionBuffer(uuid string) {
	ActionBuffer[uuid] = nil //throw to garbage collector
	ActionBuffer[uuid] = make([]Action, constants.MaxActions)
}

//ObjectID Structure générique associant chaque batiment/ressource/pnj à son id
type ObjectID struct {
	IDOffset int
	IDArray  map[string]interface{}
	m *sync.RWMutex
}

//IDIsType renvoie un booleen indiquant si un id correspond a un objet du type fourni
func IDIsType(id string, T interface{}) bool {
	return IDMap.GetObjectFromID(id) == T
}

//NewObjectID Cree une instance ObjectId
func NewObjectID() ObjectID {
	var m sync.RWMutex
	res := (ObjectID{0, nil,&m})
	res.IDArray = make(map[string]interface{}, constants.MAXOBJECTS)
	return res
}

//IDMap buffer associant id et objet
var IDMap ObjectID

//AddObject Fonction  permettant d'ajouter un objet générique à ObjectId. Retourne l'id de l'objet
func (o *ObjectID) AddObject(obj interface{}) string {
	o.m.Lock()
	key := strconv.Itoa((*o).IDOffset)
	(*o).IDArray[key] = obj
	(*o).IDOffset++
	o.m.Unlock()
	return key
}

//DeleteObjectFromID Fonction permettant de retirer un objet à partir de son id
func (o *ObjectID) DeleteObjectFromID(id string) {
	o.m.Lock()
	delete((*o).IDArray, id)
	o.m.Unlock()
}

//DeleteObject Retire un objet de la liste à partir de son propre pointeur
func (o *ObjectID) DeleteObject(obj interface{}) bool {
	id :="-1"
	o.m.RLock()
	for i, e := range (*o).IDArray {
		if e == obj {
			id=i
			break
		}
	}
	o.m.RUnlock()
	if(id=="-1"){
		return false
	}
	o.m.Lock()
	delete((*o).IDArray, id)
	o.m.Unlock()
	return true
}

//GetObjectFromID Renvoie un pointeur sur l'obj correspondant à l'id fourni
func (o *ObjectID) GetObjectFromID(id string) interface{} {
	o.m.RLock()
	obj, test := (*o).IDArray[id]
	o.m.RUnlock()
	if !test {
		return nil
	}
	return obj
}

//ConvertToInter renvoie un interface de l'objet
func ConvertToInter(obj interface{}) interface{} {
	return obj
}

//GetIDFromObject Renvoie l'id d'un objet à partir de son pointeur
func (o *ObjectID) GetIDFromObject(obj interface{}) string {
	o.m.RLock()
	id := "-1"
	for i, e := range (*o).IDArray {
		if e == obj {
			id=i
			break
		}
	}
	o.m.RUnlock()
	return id
}

//TokenValue structure contenant les parametres recuperes dans le segment data du token
type TokenValue struct {
	Group string
	Name  string
	UID  string
	Iat   int
	Exp	int
}

//ExtractFromToken retourne une map[string]string des donnees du segment d'un token
func ExtractFromToken(tokenString string) *TokenValue {
	tab := strings.Split(tokenString, ".")
	if len(tab) != 3 {
		utils.Debug("Erreur lors de l'extraction du segment data")
		return nil
	}
	segment := tab[1] //extrait le segment data du milieu
	buff, err := jwt.DecodeSegment(segment)
	if err != nil {
		utils.Debug("Erreur lors du decodage d'un token")
		return nil
	}
	var extract TokenValue
	err = json.Unmarshal(buff, &extract)
	if err != nil {
		utils.Debug("Erreur lors de la conversion du token")
		return nil
	}
	return &extract
}

//GetPlayers Returns the ids of the players of the current game
func GetPlayers()([]string,error){
	rep,err := Curl("game(id:\""+constants.GameUUID+"\"){players{id}}")
	if err!= nil{
		return nil,nil
	}
	t1:=rep["game"]
	t2:=t1.(map[string]interface{})["players"]
	t3:=t2.([]interface{})[0]
	t4:=t2.([]interface{})[1]
	res:=make([]string,2)
	res[0]=t3.(map[string]interface{})["id"].(string)
	res[1]=t4.(map[string]interface{})["id"].(string)
	return res,nil
}
//GetPlayersFromGID  get player ids from new api
func GetPlayersFromGID() ([]string,error){
	resp, err := http.Get(constants.APIHOST+"/v1/game/"+constants.GameUUID)
	if err != nil {
		utils.Debug(err.Error())
		return nil,err
	}
	tab:=make([]string,2)
	//defer resp.Body.Close()
	bodyBytes,err:= ioutil.ReadAll(resp.Body)
	if err != nil{
		utils.Debug(err.Error())
		return nil,err
	}
	var response map[string]interface{}
	err=json.Unmarshal(bodyBytes,&response)
	if err != nil {
		return nil,err
	}
	t1:=(response["Players"]).([]interface{})
	tab[0]=t1[0].(string)
	tab[1]=t1[1].(string)
	errClose:=resp.Body.Close()
	return tab,errClose
}

//Curl method POST/GET, query body ex: mutation{login(email: "gege@hotmail.fr",  password: "un")  }
func Curl(queryBody string) (map[string]interface{}, error) {
	client := graphql.NewClient(constants.APIHOST)
	req := graphql.NewRequest("query{"+queryBody+"}")
	req.Header.Set("Cache-Control", "no-cache")
	ctx := context.Background()
	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		return nil,err
	}

	return respData, nil
}
