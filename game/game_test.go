package game

import (
	d "git.unistra.fr/AOEINT/server/data"
    "testing"
    "time"
	"log"
    cst "git.unistra.fr/AOEINT/server/constants"
)

func TestDestruction(t *testing.T) {
	log.Println("running TestDestruction")
	var g Game
	d.IDMap = d.NewObjectID()
	cExit:=make(chan(bool))
	g.GameRunning = cExit
	(&g).GetPlayerData()
	g.Joueurs[0].AddAndCreateNpc("villager",0,0)
	g.Joueurs[0].AddAndCreateNpc("villager",0,0)
	npc:=g.Joueurs[0].GetPointerNpc(0)
	npc2:=g.Joueurs[0].GetPointerNpc(1)
	id := d.IDMap.GetIDFromObject(npc)
	id2 := d.IDMap.GetIDFromObject(npc2)
	if id =="-1" {
		t.Error("wrong id for npc1")
	}
	if id2 == "-1"{
		t.Error("wrong id for npc2")
	}
	if !g.DeleteNpc(npc) {
		t.Error("Echec de la suppression")
	}
	if g.Joueurs[0].GetPointerNpc(0) != nil || d.IDMap.GetObjectFromID(id) != nil {
		t.Error("Npc toujours existant")
	}
	if g.Joueurs[0].GetPointerNpc(1) == nil || d.IDMap.GetObjectFromID(id2) == nil {
		t.Error("npc2 supprimé")
	}
	//Maintenant avec AddAndCreateNpc
	(&g).Joueurs[1].AddAndCreateNpc("villager", 0, 0)
	npc = g.Joueurs[1].GetPointerNpc(0)
	id = d.IDMap.GetIDFromObject(npc)
	if id == "-1" {
		t.Error("wrong id")
	}
	g.DeleteNpc(npc)
	if g.Joueurs[1].GetPointerNpc(0) != nil || d.IDMap.GetObjectFromID(id) != nil {
		t.Error("Npc toujours existant")
	}
}


func TestAutoFight(t *testing.T) {
	log.Println("running TestAutoFight")
	var g Game
	d.IDMap = d.NewObjectID()
	cExit:=make(chan(bool))
	g.GameRunning = cExit
	(&g).GetPlayerData()
	data := ExtractData()
	(&g).GenerateMap(data)
	go (&g).LaunchAutomaticFight()
    player1 := g.GetPlayerFromUID(d.ExtractFromToken(cst.Player1JWT).UID)
    player2 := g.GetPlayerFromUID(d.ExtractFromToken(cst.Player2JWT).UID)
	player1.EntityListMutex.RLock()
    for _,pnj := range player1.GetEntities() {
        if (pnj == nil){
            break
        }
        t.Log("Player 1")
        t.Logf("type %v  a : %v pv et est à la position (%v, %v) ",
        pnj.GetType(),  pnj.GetPv(), pnj.GetX(), pnj.GetY())
	}
	player1.EntityListMutex.RUnlock()

    time.Sleep(time.Duration(time.Millisecond * 4550))
	t.Log("After fight")
    error := 0.
	t.Log("Player 1")
	player1.EntityListMutex.RLock()
    for _,pnj := range (*player1).GetEntities() {
        if (pnj == nil){
            break
        }
		if(pnj.GetType() == 0){
	        if(pnj.GetPv() > 0){
	            t.Logf("type %v  a : %v pv et est à la position (%v, %v) ",
	            pnj.GetType(),  pnj.GetPv(), pnj.GetX(), pnj.GetY() )
	            error++
	        }
		}
		if (pnj.GetType() == 2){
			if(pnj.GetPv() > 0){
	            t.Logf("type %v  a : %v pv et est à la position (%v, %v) ",
	            pnj.GetType(),  pnj.GetPv(), pnj.GetX(), pnj.GetY() )
	            error++
	        }
		}
	}
	player1.EntityListMutex.RUnlock()
	t.Log("Player 2")
	player2.EntityListMutex.RLock()
    for _,pnj := range (*player2).GetEntities() {
        if (pnj == nil){
            break
        }
		if(pnj.GetType() == 0){
	        if(pnj.GetPv() > 0){
	            t.Logf("type %v  a : %v pv et est à la position (%v, %v) ",
	            pnj.GetType(),  pnj.GetPv(), pnj.GetX(), pnj.GetY() )
	            error += 0.5
	        }
		}
		if (pnj.GetType() == 2){
			if(pnj.GetPv() > 0){
				t.Logf("type %v  a : %v pv et est à la position (%v, %v) ",
	            pnj.GetType(),  pnj.GetPv(), pnj.GetX(), pnj.GetY() )
	             error += 0.5
	        }
		}
	}
	player2.EntityListMutex.RUnlock()
    if error >= 1{
        t.Error("les npc n'ont pas perdu de pv")
	}
	// On lance le faux client pour tester les fonctions de liaison
	//go (&g).GameLoop()
}

func TestMain(m *testing.M) {
	cst.Testing=true
	var g Game
	(&g).GetPlayerData()
	d.InitiateActionBuffer()
	TestDestruction(&testing.T{})
	TestAutoFight(&testing.T{})
}
