package batiment

import (
	"strconv"
	"sync"

	cst "git.unistra.fr/AOEINT/server/constants"
	"git.unistra.fr/AOEINT/server/data"
	"git.unistra.fr/AOEINT/server/utils"
)

//Batiment : Structure contenant tous les éléments nécessaires pour la gestion d'un batiment
type Batiment struct {
	X               int
	Y               int
	Pv              int
	Typ             int //auberge: 0, caserne:1, établi:2 ...
	Longueur        int
	Largeur         int
	PlayerUID       string
	batimentChannel *chan int
	m               *sync.Mutex
}

//New : Constructeur de l'objet Batiment
func New(x int, y int, typ int, long int, larg int, pv int) Batiment {
	var m sync.Mutex
	buffer := make(chan int, cst.BatimentBufferSize)
	return (Batiment{x, y, pv, typ, long, larg, "", &buffer, &m})
}

//Create : Crée une Instance de batiment
func Create(class string, x int, y int) Batiment {
	var bat Batiment
	switch class {
	case "auberge":
		bat = New(x, y, 0, cst.LongueurAuberge, cst.LargeurAuberge, cst.PVAuberge)
	case "caserne":
		bat = New(x, y, 1, cst.LongueurCaserne, cst.LargeurCaserne, cst.PVCaserne)
	case "etabli":
		bat = New(x, y, 2, cst.LongueurEtabli, cst.LargeurEtabli, cst.PVEtabli)
	default: //défaut=auberge
		bat = New(x, y, 0, cst.LongueurAuberge, cst.LargeurAuberge, cst.PVAuberge)
	}
	go (&bat).batimentUpdate()
	return bat
}

//InitMutex initializes the mutex used to protects the building's HPs
func (bat *Batiment) InitMutex() {
	if bat.m == nil {
		var mut sync.Mutex
		bat.m = &mut
	}
}

func (bat Batiment) stringify(typ int, id string) map[string]string {
	res := make(map[string]string)
	if typ == cst.ActionNewBuilding {
		res["x"] = strconv.Itoa(bat.X)
		res["y"] = strconv.Itoa(bat.Y)
		res["pv"] = strconv.Itoa(bat.GetPv())
		res["type"] = strconv.Itoa(bat.Typ)
		res["PlayerUUID"] = bat.PlayerUID
		res["id"] = id
		return res
	}
	if typ == cst.ActionDestroyBuilding {
		res["id"] = id
		return res
	}
	if typ == cst.ActionHarmBuilding {
		res["pv"] = strconv.Itoa(bat.GetPv())
		res["id"] = id
		return res
	}
	return res
}

//Transmit : Adds the corresponding action to ActionBuffer
func (bat Batiment) Transmit(typ int, id string) {
	arr := bat.stringify(typ, id)
	for k, e := range arr {
		data.AjoutConcurrent(typ, id, k, e)
	}
}

/*batimentUpdate : Met automatiquement a jour les pv du batiment à partir du channel du batiment
 */
func (bat *Batiment) batimentUpdate() {
	utils.Debug("batiment:channel actif")
	var res int
	for {
		res = <-*(bat.batimentChannel)
		bat.SubPv(res)
		if bat.GetPv() <= 0 {
			break
		}
	}
	utils.Debug("batiment:channel inactif")
}

//DestroyBuilding : "Detruit" l'objet batiment si il n'y a plus de pv
func (bat *Batiment) DestroyBuilding() {
	bat = nil //nil permet assigner la valeur nul à un pointeur
}

//GetChannel retourne le channel de ressource du joueur
func (bat *Batiment) GetChannel() *(chan int) {
	return bat.batimentChannel
}

//GetType : retourne le type du batiment
func (bat Batiment) GetType() int {
	return bat.Typ
}

//GetPv : Retourne les pv d'un bâtiment
func (bat Batiment) GetPv() int {
	bat.m.Lock()
	val := bat.Pv
	bat.m.Unlock()
	return val
}

//SetPv : change des pv d'un bâtiment
func (bat *Batiment) SetPv(val int) {
	bat.m.Lock()
	bat.Pv = val
	bat.m.Unlock()
}

//SubPv : decrement les pv d'un bâtiment de val
func (bat *Batiment) SubPv(val int) {
	bat.m.Lock()
	bat.Pv -= val
	if bat.Pv < 0 {
		bat.Pv = 0
	}
	bat.m.Unlock()
}

//GetLongueur : Retourne la longueur d'un batiment
func (bat Batiment) GetLongueur() int {
	return bat.Longueur
}

//GetLargeur : Retourne la largeur d'un batiment
func (bat Batiment) GetLargeur() int {
	return bat.Largeur
}

//GetPlayerUID : Retourne l'id du joueur auquel appartient ce batiment
func (bat Batiment) GetPlayerUID() string {
	return bat.PlayerUID
}

//SetPlayerUID : Change le playerUUID d'un batiment
func (bat *Batiment) SetPlayerUID(UID string) {
	bat.PlayerUID = UID
}

//GetX : Retourne un coordonnée x de Batiment
func (bat Batiment) GetX() int {
	return bat.X
}

//GetY : Retourne un coordonnée y de Batiment
func (bat Batiment) GetY() int {
	return bat.Y
}

//IsATile : Retourne vrai si les coordonnées correspondent à une case du batiment
func (bat Batiment) IsATile(x int, y int) bool {
	return ((x >= bat.X-bat.Largeur+1 && x <= bat.X+bat.Largeur-1) && (y >= bat.Y-bat.Longueur+1 && y <= bat.Y+bat.Longueur-1))
}
