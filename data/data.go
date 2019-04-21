package data

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"strconv"
	"strings"

	"git.unistra.fr/AOEINT/server/constants"
	"git.unistra.fr/AOEINT/server/utils"
	jwt "github.com/dgrijalva/jwt-go"
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
}

//IDIsType renvoie un booleen indiquant si un id correspond a un objet du type fourni
func IDIsType(id string, T interface{}) bool {
	return IDMap.GetObjectFromID(id) == T
}

//NewObjectID Cree une instance ObjectId
func NewObjectID() ObjectID {
	res := (ObjectID{0, nil})
	res.IDArray = make(map[string]interface{}, constants.MAXOBJECTS)
	return res
}

//IDMap buffer associant id et objet
var IDMap ObjectID

//AddObject Fonction  permettant d'ajouter un objet générique à ObjectId. Retourne l'id de l'objet
func (o *ObjectID) AddObject(obj interface{}) string {
	key := strconv.Itoa((*o).IDOffset)
	(*o).IDArray[key] = obj
	(*o).IDOffset++
	return key
}

//DeleteObjectFromID Fonction permettant de retirer un objet à partir de son id
func (o *ObjectID) DeleteObjectFromID(id string) {
	delete((*o).IDArray, id)
}

//DeleteObject Retire un objet de la liste à partir de son propre pointeur
func (o *ObjectID) DeleteObject(obj interface{}) bool {
	for i, e := range (*o).IDArray {
		if e == obj {
			delete((*o).IDArray, i)
			return true
		}
	}
	return false
}

//GetObjectFromID Renvoie un pointeur sur l'obj correspondant à l'id fourni
func (o *ObjectID) GetObjectFromID(id string) interface{} {
	obj, test := (*o).IDArray[id]
	if !test {
		return nil
	} else {
		return obj
	}
}

//ConvertToInter renvoie un interface de l'objet
func ConvertToInter(obj interface{}) interface{} {
	return obj
}

//GetIDFromObject Renvoie l'id d'un objet à partir de son pointeur
func (o *ObjectID) GetIDFromObject(obj interface{}) string {
	for i, e := range (*o).IDArray {
		if e == obj {
			return i
		}
	}
	return "-1"
}

//TokenValue structure contenant les parametres recuperes dans le segment data du token
type TokenValue struct {
	Group string
	Name  string
	UUID  string
	Iat   int
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

//Curl method POST/GET, query body ex: mutation{login(email: "gege@hotmail.fr",  password: "un")  }
func Curl(queryBody string) (string, error) {
	body := strings.NewReader(`{ "query": "` + queryBody + `" }`)
	req, err := http.NewRequest("POST", constants.APIHOST+":"+constants.APIPORT, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	var errClose error
	defer func() {
		errClose = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		return "", err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyString := string(bodyBytes)
	return bodyString, errClose
}
