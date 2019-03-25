package npc

import "testing"
import "git.unistra.fr/AOEINT/server/carte"
import d "git.unistra.fr/AOEINT/server/data"
import "git.unistra.fr/AOEINT/server/ressource"
import "time"
import "git.unistra.fr/AOEINT/server/constants"

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

func TestRecolte(t *testing.T){
	d.IDMap=d.NewObjectID()
	d.InitiateActionBuffer()
	bip := make(chan[]int,100)
	pnj,_:=Create("harvester",0,0,false,&bip)
	c:=carte.New(50)
	ress:=ressource.Create("tree",2,2)
	c.AddNewRessource(&ress)
	go (&pnj).MoveHarvestTarget(c, &ress)
	time.Sleep(time.Duration(4*time.Second))
	if(ress.GetPv()==100){
		t.Error("la ressource n'a pas perdu de Pv")
	}
	t.Log("pv: ",ress.GetPv())
}

func TestFight(t *testing.T){
	d.IDMap=d.NewObjectID()
	d.InitiateActionBuffer()
	bip1 := make(chan[]int,100)
	bip2 := make(chan[]int,100)
	pnj1,_:=Create("soldier",10,10,false,&bip1)
	pnj2,_:=Create("soldier",14,13,true,&bip2)
	c:=carte.New(50)
	go (&pnj1).MoveFight(c, &pnj2)
	time.Sleep(time.Duration(4*time.Second))
	if(pnj2.GetPv()==constants.SoldierPv){
		t.Error("la cible n'a pas perdu de Pv")
		t.Log("pv de la cible: ",pnj2.GetPv())
	}
	if(pnj1.GetPv()==constants.SoldierPv){
		t.Error("l'agresseur n'a pas perdu de Pv")
		t.Log("pv de l'agresseur': ",pnj1.GetPv())
	}
}
