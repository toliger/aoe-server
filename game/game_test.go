package game

import (
	"log"
	"sync"
	"testing"
	"time"

	"git.unistra.fr/AOEINT/server/constants"
	cst "git.unistra.fr/AOEINT/server/constants"
	d "git.unistra.fr/AOEINT/server/data"
)

func TestDestruction(t *testing.T) {
	log.Println("running TestDestruction")
	var g Game
	d.IDMap = d.NewObjectID()
	cExit := make(chan (bool))
	g.GameRunning = cExit
	(&g).GetPlayerData()
	g.Joueurs[0].AddAndCreateNpc("villager", 0, 0)
	g.Joueurs[0].AddAndCreateNpc("villager", 0, 0)
	npc := g.Joueurs[0].GetPointerNpc(0)
	npc2 := g.Joueurs[0].GetPointerNpc(1)
	id := d.IDMap.GetIDFromObject(npc)
	id2 := d.IDMap.GetIDFromObject(npc2)
	if id == "-1" {
		t.Error("wrong id for npc1")
	}
	if id2 == "-1" {
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
	cExit := make(chan (bool))
	g.GameRunning = cExit
	(&g).GetPlayerData()
	data := ExtractData()
	(&g).GenerateMap(data)
	player1 := g.GetPlayerFromUID(cst.PlayerUID1)
	player2 := g.GetPlayerFromUID(cst.PlayerUID2)
	pnj1 := player1.GetPointerNpc(0)
	pnj2 := player2.GetPointerNpc(5)
	log.Println("-------------------Before fight-------------------")
	player1.EntityListMutex.RLock()
	log.Println("Player 1")
	for _, pnj := range player1.GetEntities() {
		if pnj == nil {
			continue
		}
		log.Printf("type %v  a : %v pv et est à la position (%v, %v) ",
			pnj.GetType(), pnj.GetPv(), pnj.GetX(), pnj.GetY())
	}
	player1.EntityListMutex.RUnlock()

	player2.EntityListMutex.RLock()
	log.Println("Player 2")
	for _, pnj := range player2.GetEntities() {
		if pnj == nil {
			continue
		}
		log.Printf("type %v  a : %v pv et est à la position (%v, %v) ",
			pnj.GetType(), pnj.GetPv(), pnj.GetX(), pnj.GetY())
	}
	player2.EntityListMutex.RUnlock()
	log.Println("-------------------Start movement-------------------")
	var wg sync.WaitGroup
	wg.Add(1)
	pnj1.MoveTo(g.Carte, 11, 7, &wg)
	// Wait for moveTo to finish
	wg.Wait()
	log.Println("-------------------During fight-------------------")
	// Wait for the fight to finish
	time.Sleep(time.Duration(time.Millisecond * 5550))
	//(&g).DeleteNpc(player1.GetPointerNpc(2))
	if pnj2 != nil && pnj1 != nil {
		t.Error("au moins un des deux npc devrait etre mort")
	}
	log.Println("-------------------After fight-------------------")
	//error := 0.
	log.Println("Player 1")
	player1.EntityListMutex.RLock()
	for i, pnj := range (*player1).GetEntities() {
		if pnj == nil && i != 0 && i <= 5 {
			t.Error("le npc numero", i, "ne devrait pas etre supprimé")
		}
		if pnj != nil {
			log.Printf("type %v  a : %v pv et est à la position (%v, %v) ",
				pnj.GetType(), pnj.GetPv(), pnj.GetX(), pnj.GetY())
		}
	}
	player1.EntityListMutex.RUnlock()
	t.Log("Player 2")
	player2.EntityListMutex.RLock()
	for i, pnj := range (player2).GetEntities() {
		if pnj != nil {
			log.Printf("type %v  a : %v pv et est à la position (%v, %v) ",
				pnj.GetType(), pnj.GetPv(), pnj.GetX(), pnj.GetY())
			continue
		}
		if pnj == nil && i < 4 {
			t.Error("le npc numero", i, "ne devrait pas etre supprimé")
			continue
		}
	}
	player2.EntityListMutex.RUnlock()
}

func TestMoveTargetNpc(t *testing.T) {
	log.Println("running TestMoveTargetNpc")
	var g Game
	d.IDMap = d.NewObjectID()
	cExit := make(chan (bool))
	g.GameRunning = cExit
	(&g).GetPlayerData()
	data := ExtractData()
	(&g).GenerateMap(data)
	player1 := g.GetPlayerFromUID(cst.PlayerUID1)
	player2 := g.GetPlayerFromUID(cst.PlayerUID2)
	pnj1 := player1.GetPointerNpc(0)
	pnj2 := player2.GetPointerNpc(5)
	log.Println("-------------------Before fight-------------------")
	player1.EntityListMutex.RLock()
	log.Println("Player 1")
	for _, pnj := range player1.GetEntities() {
		if pnj == nil {
			continue
		}
		log.Printf("type %v  a : %v pv et est à la position (%v, %v) ",
			pnj.GetType(), pnj.GetPv(), pnj.GetX(), pnj.GetY())
	}
	player1.EntityListMutex.RUnlock()

	player2.EntityListMutex.RLock()
	log.Println("Player 2")
	for _, pnj := range player2.GetEntities() {
		if pnj == nil {
			continue
		}
		log.Printf("type %v  a : %v pv et est à la position (%v, %v) ",
			pnj.GetType(), pnj.GetPv(), pnj.GetX(), pnj.GetY())
	}
	log.Printf("pnj 5 du player 2 (%v, %v)", pnj2.Get32X(), pnj2.Get32Y())
	player2.EntityListMutex.RUnlock()

	log.Println("-------------------Start movement-------------------")
	var wg sync.WaitGroup
	wg.Add(1)
	go pnj1.MoveTargetNpc(g.Carte, pnj2, &wg)
	// Wait for moveTarget to finish
	wg.Wait()
	if !pnj1.VerifyPos(pnj2.Get32X()-1, pnj2.Get32Y()-1) {
		t.Error("mauvais emplacement du pnj1")
		log.Printf("le pnj1 est à la pos (%v, %v)", pnj1.GetX(), pnj1.GetY())
	}
	log.Println("-------------------During fight-------------------")
	// Start fight
	time.Sleep(time.Duration(time.Millisecond * 1550))
	if pnj2.GetPv() == constants.VillagerPv || pnj1.GetPv() == constants.VillagerPv {
		t.Error("les deux pnj devrait avoir perdu des pv")
	}
	if pnj1 == nil || pnj2 == nil {
		t.Error("les deux pnj ne devraient pas etre morts")
	}
	log.Printf("pnj 5 du player 2 (%v, %v)", pnj2.Get32X(), pnj2.Get32Y())
	// MoveTo During fight
	log.Println("-------------------Pnj1 movement-------------------")
	wg.Add(1)
	pnj1.MoveTo(g.Carte, pnj2.Get32X()-3, pnj2.Get32Y()-3., &wg)
	wg.Wait()
	//Verify pnj1 position
	if !pnj1.VerifyPos(pnj2.Get32X()-3, pnj2.Get32Y()-3) {
		t.Error("mauvaise position du pnj1")
		log.Printf("le pnj1 est à la pos (%v, %v)", pnj1.GetX(), pnj1.GetY())
	}
	// Wait for the fight to finish
	time.Sleep(time.Duration(time.Millisecond * 4550))
	//(&g).DeleteNpc(player1.GetPointerNpc(2))
	log.Println("-------------------After fight-------------------")
	if pnj2 != nil && pnj1 != nil {
		t.Error("au moins un des deux npc devrait etre mort")
	}
	log.Println("Player 1")
	player1.EntityListMutex.RLock()
	for i, pnj := range (*player1).GetEntities() {
		if pnj == nil && i != 0 && i <= 5 {
			t.Error("le npc numero", i, "ne devrait pas etre supprimé")
		}
		if pnj != nil {
			log.Printf("type %v  a : %v pv et est à la position (%v, %v) ",
				pnj.GetType(), pnj.GetPv(), pnj.GetX(), pnj.GetY())
		}
	}
	player1.EntityListMutex.RUnlock()
	t.Log("Player 2")
	player2.EntityListMutex.RLock()
	for i, pnj := range (player2).GetEntities() {
		if pnj != nil {
			log.Printf("type %v  a : %v pv et est à la position (%v, %v) ",
				pnj.GetType(), pnj.GetPv(), pnj.GetX(), pnj.GetY())
			continue
		}
		if pnj == nil && i < 4 {
			t.Error("le npc numero", i, "ne devrait pas etre supprimé")
			continue
		}
	}
	player2.EntityListMutex.RUnlock()
}

func TestMoveTargetBuilding(t *testing.T) {
	log.Println("running TestMoveTargetBuilding")
	var g Game
	d.IDMap = d.NewObjectID()
	cExit := make(chan (bool))
	g.GameRunning = cExit
	(&g).GetPlayerData()
	data := ExtractData()
	(&g).GenerateMap(data)
	player1 := g.GetPlayerFromUID(cst.PlayerUID1)
	player2 := g.GetPlayerFromUID(cst.PlayerUID2)
	pnj1 := player1.GetPointerNpc(5)
	building1 := player2.GetPointerBuilding(0)
	log.Println("-------------------Before fight-------------------")
	for i, building := range player2.GetBuildings() {
		if building == nil {
			continue
		}
		log.Printf("Le batiment numero %v a la position (%v, %v)", i, building.GetX(), building.GetY())
	}
	log.Println("-------------------Start movement-------------------")
	var wg sync.WaitGroup
	wg.Add(1)
	if pnj1 == nil || building1 == nil {
		t.Error("le batiment ou le npc n'existe pas")
	}
	go pnj1.MoveTargetBuilding(g.Carte, player2.GetPointerBuilding(20), &wg)
	// Wait for moveTarget to finish
	wg.Wait()
	log.Println("-------------------During fight-------------------")
	// Wait for the fight to begin
	time.Sleep(time.Duration(time.Millisecond * 5150))
	log.Println("-------------------After fight-------------------")
	if player2.GetPointerBuilding(20) == nil {
		t.Error("le batiment ne devrait pas etre detruit")
	}
	if player2.GetPointerBuilding(20).GetPv() != cst.PVAuberge-5 {
		t.Error("le batiment devrait avoir perdu 5")
		t.Log("Il reste ", player2.GetPointerBuilding(20).GetPv(), " au batiment")
	}
	log.Println("Player 1")
	if player1.GetPointerNpc(20) == nil {
		t.Error("le npc ne devrait pas etre mort")
	}
}

func TestMain(m *testing.M) {
	cst.Testing = true
	var g Game
	(&g).GetPlayerData()
	d.InitiateActionBuffer()
	TestDestruction(&testing.T{})
	go (&g).LaunchAutomaticFight()
	TestAutoFight(&testing.T{})
	TestMoveTargetNpc(&testing.T{})
	TestMoveTargetBuilding(&testing.T{})
}
