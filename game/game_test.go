package game

import (
	d "git.unistra.fr/AOEINT/server/data"
    "testing"
    "time"
    "log"
	"sync"
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
	log.Println("Player 1")
    for _,pnj := range player1.GetEntities() {
        if (pnj == nil){
            continue
        }
        log.Printf("type %v  a : %v pv et est à la position (%v, %v) ",
        pnj.GetType(),  pnj.GetPv(), pnj.GetX(), pnj.GetY())
	}
	player1.EntityListMutex.RUnlock()

	player2.EntityListMutex.RLock()
	log.Println("Player 2")
    for _,pnj := range player2.GetEntities() {
        if (pnj == nil){
            continue
        }
        log.Printf("type %v  a : %v pv et est à la position (%v, %v) ",
        pnj.GetType(),  pnj.GetPv(), pnj.GetX(), pnj.GetY())
	}
	player2.EntityListMutex.RUnlock()
	var wg sync.WaitGroup
	wg.Add(1)
	player1.GetPointerNpc(0).MoveTo(g.Carte, 11, 7, &wg)
	// Wait for moveTo to finish
	wg.Wait()
	// Wait for the fight to finish
    time.Sleep(time.Duration(time.Millisecond * 5550))
	//(&g).DeleteNpc(player1.GetPointerNpc(2))
	player1.EntityListMutex.RLock()
	player2.EntityListMutex.RLock()
	if (player2.GetPointerNpc(5) != nil && player1.GetPointerNpc(0) != nil){
		t.Error("au moins un des deux npc devrait etre mort")
	}
	player2.EntityListMutex.RUnlock()
	player1.EntityListMutex.RUnlock()
	log.Println("-------------------After fight-------------------")
    //error := 0.
	log.Println("Player 1")
	player1.EntityListMutex.RLock()
    for i,pnj := range (*player1).GetEntities() {
        if (pnj == nil && i != 0 && i <= 5){
			t.Error("le npc numero", i,"ne devrait pas etre supprimé")
        }
		if (pnj != nil){
			log.Printf("type %v  a : %v pv et est à la position (%v, %v) ",
	        pnj.GetType(),  pnj.GetPv(), pnj.GetX(), pnj.GetY())
		}
	}
	player1.EntityListMutex.RUnlock()
	t.Log("Player 2")
	player2.EntityListMutex.RLock()
    for i,pnj := range (player2).GetEntities() {
		log.Print(i)
		if (pnj != nil){
			log.Printf("type %v  a : %v pv et est à la position (%v, %v) ",
	        pnj.GetType(),  pnj.GetPv(), pnj.GetX(), pnj.GetY())
			continue
		}
        if (pnj == nil && i < 4){
			t.Error("le npc numero", i,"ne devrait pas etre supprimé")
			continue
        }
	}
	player2.EntityListMutex.RUnlock()
}

func TestMain(m *testing.M) {
	cst.Testing=true
	var g Game
	(&g).GetPlayerData()
	d.InitiateActionBuffer()
	TestDestruction(&testing.T{})
	TestAutoFight(&testing.T{})
}
