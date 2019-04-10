package joueur

import (
	"testing"

	b "git.unistra.fr/AOEINT/server/batiment"
	d "git.unistra.fr/AOEINT/server/data"
	"git.unistra.fr/AOEINT/server/npc"
)

func TestCreation(t *testing.T) {
	d.IDMap = d.NewObjectID()
	d.InitiateActionBuffer()
	bip1 := make(chan []int, 100)
	bip2 := make(chan []int, 100)
	player1 := Create(1, "arnold", "0")
	player2 := Create(0, "elise", "1")
	vill1, _ := npc.Create("villager", 10, 10, player1.GetFaction(), (&bip1))
	vill2, _ := npc.Create("villager", 10, 10, player2.GetFaction(), (&bip2))
	(&player1).AddNpc(&vill1)
	(&player2).AddNpc(&vill2)
	auberge1 := b.Create("auberge", 3, 3)
	auberge2 := b.Create("auberge", 47, 47)
	(&player1).AddBuilding(&auberge1)
	(&player2).AddBuilding(&auberge2)
	if player1.GetNpc(0) != vill1 || player2.GetNpc(0) != vill2 {
		t.Error("probleme dans l'ajout de NPC")
	}
}
