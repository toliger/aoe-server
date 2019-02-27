package joueur

import npc "git.unistra.fr/AOEINT/server/npc"
import batiment "git.unistra.fr/AOEINT/server/batiment"
import constants "git.unistra.fr/AOEINT/server/constants"

type Joueur struct{
	faction bool //true: faction 1, false: faction 2
	nom string
	Uid string
	nbats int
	batiments[] batiment.Batiment
	nelems int
	entities[] npc.Npc
	id byte
	stone int
	wood int
	food int
}
var model byte =0//Permet d'obtenir des id uniques lors d'une partie

//Crée un joueur
func Create(faction bool,nom string,uid string) Joueur{
	res :=Joueur{faction,nom,uid,0,make([]batiment.Batiment,constants.MaxBuildings),0,make([]npc.Npc,constants.MaxEntities),model,constants.StartingStone,constants.StartingWood,constants.StartingFood}
	model++
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
//Retourne l'id jouer
func (j Joueur) GetId() byte{
	return j.id
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
	(*j).batiments=append(j.batiments,b)
}

func Abs(x int) int {
	if (x < 0) {
		return -x
	}
	return x
}

//Recolte de ressources (se deplace vers la ressource la plus proche dans la vue du villageois)
func (joueur Joueur) recolte(vill npc.Npc, c carte.Carte){
	var i, j int
	var ress *ressource.Ressource
	distance := 2000
	if (vill.GetType() == 2){
		fmt.Println("Un soldat ne peut pas recolter de ressources")
		return
	}
	for i = vill.GetX() - vill.GetVue(); i <= vill.GetX() + vill.GetVue(); i++{
		for j = vill.GetY() - vill.GetVue(); j <= vill.GetY() + vill.GetVue(); j++{
			if (c.GetTile(i, j).GetType() == 2){
				if (c.GetTile(i, j).GetRess().GetType() == 2 && vill.GetType() != 0){
					fmt.Println("Seul un harvester peut recolter de la pierre")
					continue;
				}
				if (Abs(i - vill.GetX()) + Abs(j - vill.GetY()) < distance){
					distance = Abs(i - vill.GetX()) + Abs(j - vill.GetY())
					ress = c.GetTile(i, j).GetRess()
				}
			}
		}
	}
	// pas de ressources dans la vue du villageois
	if (distance == 2000){
		return
	}
	go vill.MoveTo(c, ress.GetX(), ress.GetY());
	var elapsed time.Duration
	start := time.Now()
	for{
		elapsed = time.Since(start)
		if (elapsed % 1 == 0){
			// Le villageois ne se trouve pas (ou plus) à l'emplacement de la ressource
			if (vill.GetX() != ress.GetX() || vill.GetY() != ress.GetY()){
				break
			}
			switch ress.GetType(){
			case 1:
				joueur.AddWood(cst.TauxRecolte)
			case 2:
				joueur.AddStone(cst.TauxRecolte)
			case 3:
				joueur.AddFood(cst.TauxRecolte)
			default:
				joueur.AddFood(cst.TauxRecolte)
			}
		}
	}
}
