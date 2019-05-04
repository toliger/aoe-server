package game

import (
	d "git.unistra.fr/AOEINT/server/data"
    "testing"
    "time"
    "log"
	"sync"
    cst "git.unistra.fr/AOEINT/server/constants"
)
/*
func TestDestruction(t *testing.T) {
	var g Game
	d.IDMap = d.NewObjectID()
	d.InitiateActionBuffer()
	(&g).Joueurs = make([]*joueur.Joueur, 2)
	j0 := joueur.Create(0, "Bob", "b33d954f-c63e-4b48-88eb-8b5e86d94246")
	j1 := joueur.Create(1, "Alice", "1982N19N2")
	(&g).Joueurs[0] = &j0
	(&g).Joueurs[1] = &j1
	npc, id := npc.Create("villager", 0, 0, j0.GetFaction(), j0.GetChannel())
	(&g).Joueurs[0].AddNpc(npc)
	test := d.IDMap.GetIDFromObject(npc)
	if test != id {
		t.Error("id:", id, " got:", test)
	}
	if !g.DeleteNpc(npc) {
		t.Error("Echec de la suppression")
	}
	if g.Joueurs[0].GetPointerNpc(0) != nil || d.IDMap.GetObjectFromID("0") != nil {
		t.Error("Npc toujours existant")
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
*/

func TestAutoFight(t *testing.T) {
	cst.Testing=true
	d.ExtractFromToken("aubvfauipva.eyJncm91cCI6InBsYXllciIsIm5hbWUiOiJQaWVycmUgQyIsInV1aWQiOiJiMzNkOTU0Zi1jNjNlLTRiNDgtODhlYi04YjVlODZkOTQyNDYiLCJpYXQiOjE1MTYyMzkwMjJ9.oaougf")
	var g Game
	d.IDMap = d.NewObjectID()
	d.InitiateActionBuffer()
	cExit:=make(chan(bool))
	g.GameRunning = cExit
	(&g).GetPlayerData()
	d.InitiateActionBuffer()
	data := ExtractData()
	(&g).GenerateMap(data)
	go (&g).LaunchAutomaticFight()
    player1 := g.GetPlayerFromUID("b33d954f-c63e-4b48-88eb-8b5e86d94246")
    player2 := g.GetPlayerFromUID("1982N19N2")
	player1.EntityListMutex.RLock()
	log.Println("Player 1")
    for _,pnj := range player1.GetEntities() {
        if (pnj == nil){
            break
        }
        log.Printf("type %v  a : %v pv et est à la position (%v, %v) ",
        pnj.GetType(),  pnj.GetPv(), pnj.GetX(), pnj.GetY())
	}
	player1.EntityListMutex.RUnlock()

	player2.EntityListMutex.RLock()
	log.Println("Player 2")
    for _,pnj := range player2.GetEntities() {
        if (pnj == nil){
            break
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
	log.Println("Player 2")
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
