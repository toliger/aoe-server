package joueur

import "fmt"
import "sync"
//import wait "k8s.io/apimachinery/pkg/util/wait"
import time "time"
import npc "git.unistra.fr/AOEINT/server/npc"
import batiment "git.unistra.fr/AOEINT/server/batiment"
import constants "git.unistra.fr/AOEINT/server/constants"
import carte "git.unistra.fr/AOEINT/server/carte"
import ressource "git.unistra.fr/AOEINT/server/ressource"

type Joueur struct{
	faction bool //true: faction 1, false: faction 2
	nom string
	Uid string
	nbats int
	batiments[] *batiment.Batiment
	nelems int
	entities[] *npc.Npc
	id byte
	stone int
	wood int
	food int
}

type position struct{
	x int
	y int
}

var model byte =0//Permet d'obtenir des id uniques lors d'une partie

//Crée un joueur
func Create(faction bool,nom string,uid string) Joueur{
	res :=Joueur{faction,nom,uid,0,make([](*batiment.Batiment),constants.MaxBuildings),0,make([](*npc.Npc),constants.MaxEntities),model,constants.StartingStone,constants.StartingWood,constants.StartingFood}
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

func Abs(x int) int {
	if (x < 0) {
		return -x
	}
	return x
}

// Renvoie vrai si le villageois peut accéder à une case pour recolter la ressource en x,y
func (joueur Joueur) RecoltePossible(c carte.Carte, x int, y int) bool{
	for i := x-1; i <= x+1; i++{
		for j := y-1; j <= y+1; j++{
			if (c.IsEmpty(i, j)){
				return true
			}
		}
	}
	return false
}



//Recolte de ressources (se deplace vers la ressource la plus proche dans la vue du villageois)
func (joueur Joueur) DeplacementRecolte(vill npc.Npc, c carte.Carte){
	var i, j int
	var ress *ressource.Ressource
	distance := 2000
	if (vill.GetType() == 2){
		fmt.Println("Un soldat ne peut pas recolter de ressources")
		return
	}
	for i = vill.GetX() - vill.GetVue(); i <= vill.GetX() + vill.GetVue() || i > c.GetSize(); i++{
		if (i < 0){
			i = 0
		}
		for j = vill.GetY() - vill.GetVue(); j <= vill.GetY() + vill.GetVue() || j > c.GetSize(); j++{
			if (j < 0){
				j = 0
			}
			if (c.GetTile(i, j).GetType() == 2){
				if (c.GetTile(i, j).GetRess().GetType() == 2 && vill.GetType() != 0){
					fmt.Println("Seul un harvester peut recolter de la pierre")
					continue;
				}
				if ((Abs(i - vill.GetX()) + Abs(j - vill.GetY())) < distance &&
					joueur.RecoltePossible(c, i, j)){
					distance = Abs(i - vill.GetX()) + Abs(j - vill.GetY())
					ress = c.GetTile(i, j).GetRess()
				}
			}
		}
	}
	//fmt.Println(ress == nil)

	// pas de ressources dans la vue du villageois
	if (distance == 2000){
		return
	}

	var posRecolteVillX, posRecolteVillY int
	distance = 2000

	for i = ress.GetX() - 1; i <= ress.GetX() + 1; i++{
		for j = ress.GetY() - 1; j <= ress.GetY() + 1; j++{
			if ( (Abs(i - vill.GetX()) + Abs(j - vill.GetY()) ) < distance &&
				c.IsEmpty(i, j)){
				distance = Abs(i - vill.GetX()) + Abs(j - vill.GetY())
				posRecolteVillX = i
				posRecolteVillY = j
			}
		}
	}
	// pas d'accès possible pour recolter la ressource
	if (distance == 2000){
		return
	}
	// on attends que le villageois est finit son déplacement
	var wg sync.WaitGroup
	wg.Add(1)
    go (&vill).MoveTo(c, posRecolteVillX, posRecolteVillY, &wg)
	wg.Wait()

    // fmt.Printf("posRecolteVillX : %d, posRecolteVillY : %d\n", posRecolteVillX, posRecolteVillY)
	// fmt.Printf("villX : %d, villY: %d\n", vill.GetX(), vill.GetY())

	// Le villageois se trouve bien à l'emplacement de la recolte?
	if (vill.GetX() == (posRecolteVillX-1) && vill.GetY() == posRecolteVillY-1){
		 go joueur.Recolte(vill, c, ress, posRecolteVillX, posRecolteVillY)
	}
}

// Effectue la recolte de la ressource (x par seconde)
func (joueur Joueur) Recolte(vill npc.Npc, c carte.Carte, ress *ressource.Ressource,
	posRecolteVillX int, posRecolteVillY int){
	uptimeTicker := time.NewTicker(1 * time.Second)
	tps_ecoule := 0
	for {
		// La ressource est épuisée ou le villageois est mort
		if (tps_ecoule == ress.GetPv() || vill.GetPv() == 0){
			break
		}
		// Le villageois ne se trouve plus à l'emplacement de la ressource
		if (vill.GetX() != (posRecolteVillX-1) || vill.GetY() != posRecolteVillY-1){
			break;
		}
		select {
		case <-uptimeTicker.C:
			tps_ecoule++
			switch ress.GetType(){
			case 1:
				joueur.AddWood(constants.TauxRecolte)
			case 2:
				joueur.AddStone(constants.TauxRecolte)
			case 3:
				joueur.AddFood(constants.TauxRecolte)
			default:
				joueur.AddFood(constants.TauxRecolte)
			}
		}
	}
}
