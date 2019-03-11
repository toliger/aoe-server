package joueur

//import wait "k8s.io/apimachinery/pkg/util/wait"
import npc "git.unistra.fr/AOEINT/server/npc"
import batiment "git.unistra.fr/AOEINT/server/batiment"
import constants "git.unistra.fr/AOEINT/server/constants"
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
	buffer:=make(chan []int,constants.RESSOURCE_BUFFER_SIZE)
	res :=Joueur{faction,nom,uid,0,make([](*batiment.Batiment),constants.MaxBuildings),0,make([](*npc.Npc),constants.MaxEntities),constants.StartingStone,constants.StartingWood,constants.StartingFood,buffer}
	go (&res).ressourceUpdate()
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

func (j *Joueur)AddBuilding(b batiment.Batiment){
	(*j).batiments=append(j.batiments,&b)
}
func (j *Joueur)AddNpc(entity npc.Npc){
	test:=false
	for i:=0;i<len(j.entities);i++{
		if(j.entities[i]==nil){
			j.entities[i]=&entity
			test=true
			break
		}
	}
	if(!test){
		(*j).entities=append(j.entities,&entity)
	}
}
