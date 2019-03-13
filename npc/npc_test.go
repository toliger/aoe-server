package npc

import "testing"
import "git.unistra.fr/AOEINT/server/carte"
import d "git.unistra.fr/AOEINT/server/data"
import "git.unistra.fr/AOEINT/server/ressource"
import "time"
//import "math/rand"
/*
func TestDeplacement(t *testing.T){
	d.IdMap=d.NewObjectID()
	d.InitiateActionBuffer()
	bip := make(chan[]int,10)
	pnj,_:=Create("harvester",0,0,false,bip)
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
	d.IdMap=d.NewObjectID()
	d.InitiateActionBuffer()
	bip := make(chan[]int,100)
	pnj,_:=Create("harvester",0,0,false,bip)
	c:=carte.New(50)
	ress:=ressource.Create("tree",2,2)
	c.AddNewRessource(&ress)
	(&pnj).MoveHarvest(c)
	time.Sleep(time.Duration(4000000000))
	if(ress.GetPv()!=100){
		t.Error("la ressource n'a pas perdu de Pv")
	}
	t.Log("pv: ",ress.GetPv())
}
