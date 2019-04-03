package joueur

import (
	"strconv"

	"git.unistra.fr/AOEINT/server/batiment"
	"git.unistra.fr/AOEINT/server/constants"
	"git.unistra.fr/AOEINT/server/data"
	"git.unistra.fr/AOEINT/server/npc"
	"git.unistra.fr/AOEINT/server/utils"
)

//Joueur :
type Joueur struct {
	faction          int //true: faction 1, false: faction 2
	nom              string
	UID              string
	nbats            int
	batiments        []*batiment.Batiment
	nelems           int
	entities         []*npc.Npc
	stone            int
	wood             int
	food             int
	ressourceChannel chan []int
}

//GetChannel retourne le channel de ressource du joueur
func (j *Joueur) GetChannel() *(chan []int) {
	return &j.ressourceChannel
}

//Create : generate a player
func Create(faction int, nom string, uid string) Joueur {
	buffer := make(chan []int, constants.RessourceBufferSize)
	res := Joueur{faction, nom, uid, 0, make([](*batiment.Batiment), constants.MaxBuildings), 0, make([](*npc.Npc), constants.MaxEntities), constants.StartingStone, constants.StartingWood, constants.StartingFood, buffer}
	go (&res).ressourceUpdate()
	return res
}

//GetFaction : return the faction
func (j Joueur) GetFaction() int {
	return j.faction
}

//GetNom : return the name
func (j Joueur) GetNom() string {
	return j.nom
}

func (j Joueur) stringify() map[string]string {
	result := make(map[string]string)
	result["nom"] = j.nom
	result["faction"] = strconv.Itoa(j.faction)
	result["uid"] = j.UID
	result["stone"] = strconv.Itoa(j.stone)
	result["wood"] = strconv.Itoa(j.wood)
	result["food"] = strconv.Itoa(j.food)
	return result
}

//Transmit ajoute le joueur au buffer d'action
func (j Joueur) Transmit() {
	arr := j.stringify()
	for k, e := range arr {
		data.AddNewAction(j.UID, constants.ActionPlayerRessource, j.UID, k, e)
	}
}

/*ressourceUpdate :Met automatiquement a jour les ressources du joueur a partir des int[3] envoyes au channel du joueur
arrêt du thread dedié si la premiere valeur du tableau reçu par le channel est -1
*/
func (j *Joueur) ressourceUpdate() {
	var res []int
	utils.Debug(j.nom + ":channel actif")
	for {
		res = <-j.ressourceChannel
		if res[0] != 1 {
			j.AddWood(res[0])
			j.AddStone(res[1])
			j.AddFood(res[2])
		} else {
			break
		}
	}
	utils.Debug(j.nom + ":channel inactif")
}

//Retourne la quantité de d'une ressource d'un joueur

//GetStone :
func (j Joueur) GetStone() int {
	return j.stone
}

//GetWood :
func (j Joueur) GetWood() int {
	return j.wood
}

//GetFood :
func (j Joueur) GetFood() int {
	return j.food
}

//GetNpc :
func (j Joueur) GetNpc(i int) npc.Npc {
	return *(j.entities[i])
}

//GetBatiment :
func (j Joueur) GetBatiment(i int) batiment.Batiment {
	return *(j.batiments[i])
}

//Adding ressources

//AddStone :
func (j *Joueur) AddStone(s int) {
	(*j).stone += s
}

//AddWood :
func (j *Joueur) AddWood(w int) {
	(*j).wood += w
}

//AddFood :
func (j *Joueur) AddFood(f int) {
	(*j).food += f
}

//AddBuilding : add a new building to the player
func (j *Joueur) AddBuilding(b *batiment.Batiment) {
	(*j).batiments = append(j.batiments, b)
}

//AddNpc : add a new NPC to the player
func (j *Joueur) AddNpc(entity *npc.Npc) {
	test := false
	for i := 0; i < len(j.entities); i++ {
		if j.entities[i] == nil {
			j.entities[i] = entity
			test = true
			break
		}
	}
	if !test {
		(*j).entities = append(j.entities, entity)
	}
	entity.PlayerUUID = j.UID
}

//AddAndCreateNpc : create and add a new NPC to the player
func (j *Joueur) AddAndCreateNpc(class string, x int, y int) {
	entity, id := npc.Create(class, float64(x), float64(y), j.faction, &j.ressourceChannel)
	test := false
	for i := 0; i < len(j.entities); i++ {
		if j.entities[i] == nil {
			j.entities[i] = &entity
			test = true
			break
		}
	}
	if !test {
		(*j).entities = append(j.entities, &entity)
	}
	entity.PlayerUUID = j.UID
	entity.Transmit(id)
}
