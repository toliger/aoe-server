package joueur

import (
	"log"
	"testing"
	"time"

	b "git.unistra.fr/AOEINT/server/batiment"
	"git.unistra.fr/AOEINT/server/carte"
	"git.unistra.fr/AOEINT/server/constants"
	d "git.unistra.fr/AOEINT/server/data"
)

func TestCreation(t *testing.T) {
	constants.PlayerUID1 = "0"
	constants.PlayerUID2 = "1"
	d.IDMap = d.NewObjectID()
	d.InitiateActionBuffer()
	c := carte.New(50)
	player1 := Create(1, "arnold", "0")
	player2 := Create(0, "elise", "1")
	auberge1 := b.Create("auberge", 3, 3)
	auberge2 := b.Create("auberge", 47, 47)
	c.AddNewBuilding(&auberge1)
	c.AddNewBuilding(&auberge2)
	(&player1).AddBuilding(&auberge1)
	(&player2).AddBuilding(&auberge2)
	log.Printf("le joueur1 a commencé avec %v de nourriture", player1.GetFood())
	if &auberge1 == nil {
		t.Error("auberge non existante")
	}
	(&player1).AddAndCreateNpcByBuilding(&c, auberge1.GetX(), auberge1.GetY())
	time.Sleep(time.Duration(1 * time.Second))
	if player1.GetBuildings() == nil {
		t.Error("le building n'a pas été bien ajouté au joueur")
	}
	log.Printf("il devrait rester %v nourriture au joueur 1", constants.StartingFood-constants.VillagerFoodCost)
	// if player1.GetFood() != constants.StartingFood-constants.VillagerFoodCost {
	// 	t.Error("erreur dans la diminution des ressources du joueur")
	// 	log.Printf("le joueur1 a %v de nourriture", player1.GetFood())
	// }
}
