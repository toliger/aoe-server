package joueur

import (
	"log"
	"strconv"

	"sync"

	"git.unistra.fr/AOEINT/server/batiment"
	"git.unistra.fr/AOEINT/server/carte"
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
	EntityListMutex  *sync.RWMutex
}

//GetChannel retourne le channel de ressource du joueur
func (j *Joueur) GetChannel() *(chan []int) {
	return &j.ressourceChannel
}

//Create : generate a player
func Create(faction int, nom string, uid string) Joueur {
	buffer := make(chan []int, constants.RessourceBufferSize)
	var m sync.RWMutex
	res := Joueur{faction, nom, uid, 0, make([](*batiment.Batiment), constants.MaxBuildings), 0, make([](*npc.Npc), constants.MaxEntities), constants.StartingStone, constants.StartingWood, constants.StartingFood, buffer, &m}
	go (&res).ressourceUpdate()
	return res
}

//GetUID :
func (j Joueur) GetUID() string {
	return j.UID
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
		data.AjoutJoueurConcurrent(j.UID, constants.ActionPlayerRessource, j.UID, k, e)
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
		if res[0] != -1 {
			j.AddWood(res[0])
			j.AddStone(res[1])
			j.AddFood(res[2])
			data.AjoutJoueurConcurrent(j.UID, constants.ActionPlayerRessource, j.UID, "wood", strconv.Itoa(j.GetWood()))
			data.AjoutJoueurConcurrent(j.UID, constants.ActionPlayerRessource, j.UID, "food", strconv.Itoa(j.GetFood()))
			data.AjoutJoueurConcurrent(j.UID, constants.ActionPlayerRessource, j.UID, "stone", strconv.Itoa(j.GetStone()))
			log.Print("ressourceUpdate")
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

//GetEntities :
func (j Joueur) GetEntities() []*npc.Npc {
	return (j.entities)
}

//GetBuildings renvoie la liste des batiments du jouer
func (j Joueur) GetBuildings() []*batiment.Batiment {
	return j.batiments
}

//GetPointerBuilding : renvoie un pointeur sur un batiment
func (j Joueur) GetPointerBuilding(i int) *batiment.Batiment {
	return j.batiments[i]
}

//GetNpc :
func (j Joueur) GetNpc(i int) npc.Npc {
	return *(j.entities[i])
}

//GetPointerNpc :
func (j Joueur) GetPointerNpc(i int) *npc.Npc {
	return j.entities[i]
}

//GetPointerNpcByPos :
func (j Joueur) GetPointerNpcByPos(x int, y int) *npc.Npc {
	for _, pnj := range j.entities {
		if pnj.GetX() == x && pnj.GetY() == y {
			return pnj
		}
	}
	return nil
}

//DeleteNpcFromList retire un pnj de la liste du joueur
func (j *Joueur) DeleteNpcFromList(x float32, y float32, typ int, pv int, id string) bool {
	j.EntityListMutex.RLock()
	index := -1
	for i := range j.entities {
		if j.entities[i] == nil {
			continue
		}
		if j.entities[i].Get32X() == x && j.entities[i].Get32Y() == y && j.entities[i].GetType() == typ && j.entities[i].GetPv() == pv {
			if data.IDMap.GetIDFromObject(j.entities[i]) == id {
				index = i
				break
			}
		}
	}
	j.EntityListMutex.RUnlock()
	if index == -1 {
		return false
	}
	j.EntityListMutex.Lock()
	j.entities[index].SetPv(0)
	j.entities[index] = nil
	j.EntityListMutex.Unlock()
	return true
}

//GetBatiment :
func (j Joueur) GetBatiment(i int) batiment.Batiment {
	return *(j.batiments[i])
}

//DeleteBatimentFromList retire un batiment de la liste du joueur
func (j *Joueur) DeleteBatimentFromList(x int, y int, typ int) bool {
	if j == nil {
		return false
	}
	for i := range j.batiments {
		if j.batiments[i] != nil {
			if j.batiments[i].X == x && j.batiments[i].Y == y && j.batiments[i].Typ == typ {
				j.batiments[i] = nil
				return true
			}
		}
	}
	return false
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
	log.Print("addFood")
}

//AddBuilding : add a new building to the player
func (j *Joueur) AddBuilding(b *batiment.Batiment) {
	b.SetPlayerUID(j.UID)
	(*j).batiments = append(j.batiments, b)
	b.PlayerUID = j.GetUID()
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

/*AddAndCreateNpcVerification : create and add a new NPC to the player
* Also verify that the npc is created by a building
 */
func (j *Joueur) AddAndCreateNpcVerification(class string, x int, y int) {
	for i := range j.batiments {
		if j.GetPointerBuilding(i) != nil {
			if j.GetPointerBuilding(i).GetX()-1 == x && j.GetPointerBuilding(i).GetY() == y {
				break
			} else if j.GetPointerBuilding(i).GetX()+1 == x && j.GetPointerBuilding(i).GetY() == y {
				break
			} else if j.GetPointerBuilding(i).GetX() == x && j.GetPointerBuilding(i).GetY()-1 == y {
				break
			} else if j.GetPointerBuilding(i).GetX() == x && j.GetPointerBuilding(i).GetY()+1 == y {
				break
			} else {
				return
			}
		}
	}
	entity, id := npc.Create(class, float32(x), float32(y), j.faction, &j.ressourceChannel)
	j.AddNpc(entity)
	entity.Transmit(id, constants.ActionNewNpc)
}

//AddAndCreateNpc : create and add a new NPC to the player
func (j *Joueur) AddAndCreateNpc(class string, x int, y int) {
	entity, id := npc.Create(class, float32(x), float32(y), j.faction, &j.ressourceChannel)
	j.AddNpc(entity)
	entity.Transmit(id, constants.ActionNewNpc)
}

/*AddAndCreateNpcByBuilding : create and add a new NPC to the player by a building
* The NPC is created preferably at left or at right of the building and cost ressources
 */
func (j *Joueur) AddAndCreateNpcByBuilding(c carte.Carte, posX int, posY int, class int) {
	var b *batiment.Batiment
	b = nil
	for i := 0; i < len((*j).entities); i++ {
		if j.batiments[i] == nil {
			continue
		}
		if j.batiments[i].GetX() == posX && j.batiments[i].GetY() == posY {
			b = (*j).batiments[i]
			break
		}
	}
	if b == nil {
		return
	}
	if b.GetPlayerUID() != j.UID {
		utils.Debug("batiment choisi n'appartenant pas au joueur")
		return
	}
	y := b.GetY()
	x := b.GetX()
	//left->right->before->behind
	if b.GetX() > c.GetSize()/2 {
		if c.IsEmpty(b.GetX()-1, b.GetY()) {
			x = b.GetX() - 1
		} else if c.IsEmpty(b.GetX()+1, b.GetY()) {
			x = b.GetX() + 1
		} else if c.IsEmpty(b.GetX(), b.GetY()+1) {
			y = b.GetY() + 1
		} else if c.IsEmpty(b.GetX(), b.GetY()-1) {
			y = b.GetY() - 1
		} else {
			utils.Debug("no empty tiles near the building to create NPC")
			return
		}
		//right->left->before->behind
	} else {
		if c.IsEmpty(b.GetX()+1, b.GetY()) {
			x = b.GetX() + 1
		} else if c.IsEmpty(b.GetX()-1, b.GetY()) {
			x = b.GetX() - 1
		} else if c.IsEmpty(b.GetX(), b.GetY()+1) {
			x = b.GetY() + 1
		} else if c.IsEmpty(b.GetX(), b.GetY()-1) {
			y = b.GetY() - 1
		} else {
			utils.Debug("no empty tiles near the building to create NPC")
			return
		}
	}
	switch b.GetType() {
	case 0:
		if class != 0 {
			utils.Debug("Erreur: Mauvais type npc création")
			return
		}
		entity, id := npc.Create("villager", float32(x), float32(y), j.faction, &j.ressourceChannel)
		j.AddNpc(entity)
		entity.Transmit(id, constants.ActionNewNpc)
		tabRessources := make([]int, 3) //0 bois 1 pierre 2 nourriture
		tabRessources[2] = -(constants.VillagerFoodCost)
		j.ressourceChannel <- tabRessources
	case 1:
		if class != 2 {
			utils.Debug("Erreur: Mauvais type npc création")
			return
		}
		entity, id := npc.Create("soldier", float32(x), float32(y), j.faction, &j.ressourceChannel)
		j.AddNpc(entity)
		entity.Transmit(id, constants.ActionNewNpc)
		tabRessources := make([]int, 3) //0 bois 1 pierre 2 nourriture
		tabRessources[2] = -(constants.SoldierFoodCost)
		tabRessources[0] = -(constants.SoldierWoodCost)
		j.ressourceChannel <- tabRessources
	case 2:
		if class != 1 {
			utils.Debug("Erreur: Mauvais type npc création")
			return
		}
		entity, id := npc.Create("harvester", float32(x), float32(y), j.faction, &j.ressourceChannel)
		j.AddNpc(entity)
		entity.Transmit(id, constants.ActionNewNpc)
		tabRessources := make([]int, 3) //0 bois 1 pierre 2 nourriture
		tabRessources[2] = -(constants.HarvesterFoodCost)
		tabRessources[0] = -(constants.HarvesterWoodCost)
		j.ressourceChannel <- tabRessources
	default:
		utils.Debug("Type du batiment non reconnu")
	}
}

//IsThereNpcInRange : returns the first npc of the player in range of the given npc if there is one else nil
func (j *Joueur) IsThereNpcInRange(pnj *npc.Npc) *npc.Npc {
	if (*j).entities == nil {
		return nil
	}
	for i := 0; i < len((*j).entities); i++ {
		if (*j).entities[i] == nil {
			continue
		}
		for x := pnj.GetX() - pnj.GetPortee(); x <= pnj.GetX()+pnj.GetPortee(); x++ {
			for y := pnj.GetY() - pnj.GetPortee(); y <= pnj.GetY()+pnj.GetPortee(); y++ {
				if ((*j).entities[i].GetX() == x) && ((*j).entities[i].GetY() == y) {
					return (*j).entities[i]
				}
			}
		}

	}
	return nil
}

//IsThereBuildingInRange : returns the first building of the player in range of the given npc if there is one else nil
func (j *Joueur) IsThereBuildingInRange(pnj *npc.Npc) *batiment.Batiment {
	if (*j).batiments == nil {
		return nil
	}
	for i := 0; i < len((*j).batiments); i++ {
		if (*j).batiments[i] == nil {
			continue
		}
		for x := pnj.GetX() - pnj.GetPortee(); x <= pnj.GetX()+pnj.GetPortee(); x++ {
			for y := pnj.GetY() - pnj.GetPortee(); y <= pnj.GetY()+pnj.GetPortee(); y++ {
				if (*j).batiments[i] != nil {
					if (*j).batiments[i].GetX() == x && j.batiments[i].GetY() == y {
						return (*j).batiments[i]
					}
				}
			}
		}
	}
	return nil
}
