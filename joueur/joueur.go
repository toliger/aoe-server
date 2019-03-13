package joueur

import "git.unistra.fr/AOEINT/server/npc"
import "git.unistra.fr/AOEINT/server/batiment"
import "git.unistra.fr/AOEINT/server/constants"
import "git.unistra.fr/AOEINT/server/data"
import "strconv"
import "fmt"

type Joueur struct{
	faction bool //true: faction 1, false: faction 2
	nom string
	Uid string
	nbats int
	batiments[] *batiment.Batiment
	nelems int
	entities[] *npc.Npc
	stone int
	wood int
	food int
	ressourceChannel chan []int
}

//Crée un joueur
func Create(faction bool,nom string,uid string) Joueur{
	buffer:=make(chan []int,constants.RessourceBufferSize)
	res :=Joueur{faction,nom,uid,0,make([](*batiment.Batiment),constants.MaxBuildings),0,make([](*npc.Npc),constants.MaxEntities),constants.StartingStone,constants.StartingWood,constants.StartingFood,buffer}
	go (&res).ressourceUpdate()
	res.Transmit()
	return res
}
//Retourne la faction
func (j Joueur) GetFaction() bool{
	return j.faction
}
//Retourne le Nom
func (j Joueur) GetNom() string{
	return j.nom
}

func (j Joueur) stringify() map[string]string{
	result:=make(map[string]string)
	result["nom"]=j.nom
	result["faction"]=strconv.FormatBool(j.faction)
	result["uid"]=j.Uid
	result["stone"]=strconv.Itoa(j.stone)
	result["wood"]=strconv.Itoa(j.wood)
	result["food"]=strconv.Itoa(j.food)
	return result
}

func (j Joueur) Transmit(){
	arr:=j.stringify()
	for k,e := range arr{
		data.AddNewAction(constants.ActionPlayerRessource,j.Uid,k,e)
	}
}

//Met automatiquement a jour les ressources du joueur a partir des int[3] envoyes au channel du joueur
//arrêt du thread dedié si la premiere valeur du tableau reçu par le channel est -1
func (j *Joueur)ressourceUpdate(){
	var res []int
	fmt.Println(j.nom,":channel actif")
	for{
		res=<-j.ressourceChannel
		if(res[0]!=1){
			j.AddWood(res[0])
			j.AddStone(res[1])
			j.AddFood(res[2])
		}else{
			break
		}
	}
	fmt.Println(j.nom,":channel inactif")
}
//Retourne l'uid du joueur
func (j Joueur) GetUid() string{
	return j.Uid
}

//Retourne la quantité de d'une ressource d'un joueur
func (j Joueur) GetStone() int{
	return j.stone
}
func (j Joueur) GetWood() int{
	return j.wood
}
func (j Joueur) GetFood() int{
	return j.food
}

func (j Joueur) GetNpc(i int) npc.Npc{
	return *(j.entities[i])
}

//ajout de ressources
func (j *Joueur) AddStone(s int){
	(*j).stone +=s
}
func (j *Joueur) AddWood(w int){
	(*j).wood +=w
}
func (j *Joueur) AddFood(f int){
	(*j).food+= f
}

func (j *Joueur)AddBuilding(b *batiment.Batiment){
	b.PlayerUID=j.Uid
	(*j).batiments=append(j.batiments,b)
}
func (j *Joueur)AddNpc(entity *npc.Npc){
	test:=false
	for i:=0;i<len(j.entities);i++{
		if(j.entities[i]==nil){
			j.entities[i]=entity
			test=true
			break
		}
	}
	if(!test){
		(*j).entities=append(j.entities,entity)
	}
	entity.PlayerUUID=(*j).Uid
}

///////////////////////////////////////////////////////
func (j *Joueur)GetChannel() chan []int {
    return j.ressourceChannel
}
