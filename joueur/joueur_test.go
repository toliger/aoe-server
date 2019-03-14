package joueur

import(
    "testing"
    d "git.unistra.fr/AOEINT/server/data"
    b "git.unistra.fr/AOEINT/server/batiment"
    "git.unistra.fr/AOEINT/server/npc"
)

func TestCreation(t *testing.T){
    d.IDMap=d.NewObjectID()
	d.InitiateActionBuffer()
    bip1 := make(chan[]int,100)
    bip2 := make(chan[]int,100)
    player1 := Create(true, "arnold", "0")
    player2 := Create(false, "elise", "1")
    vill1,_ := npc.Create("villager",10,10,player1.GetFaction(),(&bip1))
    vill2,_ := npc.Create("villager",10,10,player2.GetFaction(),(&bip2))
    (&player1).AddNpc(&vill1)
    (&player2).AddNpc(&vill2)
    auberge1 := b.Create("auberge",3,3)
    auberge2 := b.Create("auberge",47,47)
    (&player1).AddBuilding(&auberge1)
    (&player2).AddBuilding(&auberge2)
    if (player1.GetNpc(0) != vill1){
        t.Error("probleme dans l'ajout de NPC")
    }
}
