package npc

import (
	"testing"
	"time"

	"git.unistra.fr/AOEINT/server/carte"
	d "git.unistra.fr/AOEINT/server/data"
	"git.unistra.fr/AOEINT/server/ressource"
)

/*
func TestDeplacement(t *testing.T){
	d.IDMap=d.NewObjectID()
	d.InitiateActionBuffer()
	bip := make(chan[]int,10)
	pnj,_:=Create("harvester",0,0,false,&bip)
	c:=carte.New(50)
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	x:=r1.Intn(50)
	y:=r1.Intn(50)
	path:=pnj.MoveTo(c,x,y,nil)
	if(path==nil){
		t.Error("Erreur: pas de chemin de (0,0) vers (",x,",",y,")")
	}
	//Petit délai de 2*10-6 seconde pour marge d'erreur
	dep:=(1000000000/pnj.vitesse)*(len(path)+1)
	time.Sleep(time.Duration(dep))
	if(pnj.x!=x || pnj.y!=y){
		t.Log(path)
		t.Error("L'entité ne s'est pas déplacée à temps: ",pnj.x,":",pnj.y,"!=",x,":",y)
	}
}
*/

func TestRecolteContraintes(t *testing.T) {
	d.IDMap = d.NewObjectID()
	d.InitiateActionBuffer()
	bip := make(chan []int, 100)
	pnj, _ := Create("harvester", 0, 0, 0, &bip)
	c := carte.New(50)
	ress := ressource.Create("tree", 2, 2)
	obstacle1 := ressource.Create("tree", 1, 1)
	obstacle2 := ressource.Create("tree", 1, 2)
	obstacle3 := ressource.Create("tree", 2, 1)
	obstacle4 := ressource.Create("tree", 2, 3)
	obstacle5 := ressource.Create("tree", 3, 2)
	obstacle6 := ressource.Create("tree", 3, 3)
	obstacle7 := ressource.Create("tree", 3, 1)
	obstacle8 := ressource.Create("tree", 1, 3)
	c.AddNewRessource(&obstacle1)
	c.AddNewRessource(&obstacle2)
	c.AddNewRessource(&obstacle3)
	c.AddNewRessource(&obstacle4)
	c.AddNewRessource(&obstacle5)
	c.AddNewRessource(&obstacle6)
	c.AddNewRessource(&obstacle7)
	c.AddNewRessource(&obstacle8)
	c.AddNewRessource(&ress)
	go (pnj).MoveHarvestTarget(c, &ress)
	time.Sleep(time.Duration(5 * time.Second))
	if pnj.GetX() != 0 || pnj.GetY() != 0 {
		t.Error("le pnj n'est pas sencé s'etre deplacé car bloqué par des obstacles")
	}
}

func TestRecolte(t *testing.T) {
	d.IDMap = d.NewObjectID()
	d.InitiateActionBuffer()
	bip := make(chan []int, 100)
	pnj, _ := Create("harvester", 0, 0, 0, &bip)
	c := carte.New(50)
	ress := ressource.Create("tree", 2, 2)
	ress2 := ressource.Create("tree", 32, 18)
	c.AddNewRessource(&ress)
	c.AddNewRessource(&ress2)
	go (pnj).MoveHarvestTarget(c, &ress2)
	time.Sleep(time.Duration(5 * time.Second))
	if ress2.GetPv() != 100 {
		t.Error("la ressource n'est pas sencé avoir perdu de pv car hors de vue")
	}
	go (pnj).MoveHarvestTarget(c, &ress)
	time.Sleep(time.Duration(2 * time.Second))
	go (pnj).MoveTo(c, 10, 10, nil)
	time.Sleep(time.Duration(1 * time.Second))
	go (pnj).MoveTo(c, 16, 18, nil)
	time.Sleep(time.Duration(5 * time.Second))
	if pnj.GetX() != 16 || pnj.GetY() != 18 {
		t.Error("le pnj n'est pas au bon endroit")
		t.Log("pnjX :", pnj.GetX(), " pnjY :", pnj.GetY())
	}
	if ress.GetPv() == 100 {
		t.Error("la ressource n'a pas perdu de Pv")
		t.Log("pv: ", ress.GetPv())
	}
}

/*
func TestFightBuilding(t *testing.T){
	d.IDMap=d.NewObjectID()
	d.InitiateActionBuffer()
	bip1 := make(chan[]int,100)
	pnj1,_:= Create("soldier",10,10,false,&bip1)
	b := batiment.Create("auberge", 14,12)
	c := carte.New(50)
	if(c.AddNewBuilding(&b)!=true){
		t.Error("cannot add building")
	}
	ch := make(chan bool, 2)
	go (&pnj1).MoveFightBuilding(c, &b, &ch)
	time.Sleep(time.Duration(6*time.Second))

	if((&b).Pv==constants.PVAuberge){
		t.Error("la cible n'a pas perdu de Pv blbl")
		t.Log("pv de la cible: ", b.GetPv())
	}
}
*/

/*
func TestFightNpc(t *testing.T) {
	d.IDMap = d.NewObjectID()
	d.InitiateActionBuffer()
	bip1 := make(chan []int, 100)
	bip2 := make(chan []int, 100)
	pnj1, _ := Create("soldier", 10.0, 10.0, 0, &bip1)
	pnj2, _ := Create("soldier", 14.0, 13.0, 1, &bip2)
	c := carte.New(50)
	ress := ressource.Create("tree", 16, 18)
	c.AddNewRessource(&ress)
	//ch4 := make(chan bool, 2)
	go (pnj1).MoveFight(c, pnj2)
	time.Sleep(time.Duration(4 * time.Second))
	// go (&pnj2).MoveTo(c, 17,15, nil)
	// time.Sleep(time.Duration(4 * time.Second))

	// if pnj1.GetX() != 17 || pnj1.GetY() != 15{
	// 	t.Error("mauvais deplacement apres un fight")
	// 	t.Log("pnjX :", pnj1.GetX(), " pnjY :", pnj1.GetY())
	// }
	if pnj2.GetPv() == constants.SoldierPv {
		t.Error("la cible n'a pas perdu de Pv")
		t.Log("pv de la cible: ", pnj2.GetPv())
	}

	go (pnj1).MoveHarvestTarget(c, &ress)
	time.Sleep(time.Duration(6 * time.Second))
	go (pnj1).MoveFight(c, pnj2)
	time.Sleep(time.Duration(4 * time.Second))
	if pnj2.GetPv() == constants.SoldierPv {
		t.Error("la cible n'a pas perdu de Pv")
		t.Log("pv de la cible: ", pnj2.GetPv())
	}
	// if(pnj1.GetPv()==constants.SoldierPv){
	// 	t.Error("l'agresseur n'a pas perdu de Pv")
	// 	t.Log("pv de l'agresseur': ", pnj1.GetPv())
	// }
}
*/
