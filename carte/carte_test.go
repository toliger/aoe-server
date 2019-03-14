package carte

import "testing"
import "time"
import batiment "git.unistra.fr/AOEINT/server/batiment"
import ressource "git.unistra.fr/AOEINT/server/ressource"
import d "git.unistra.fr/AOEINT/server/data"
import "math/rand"

//Verifie si la carte creee est de la bonne taille et vide
func TestCreation(t *testing.T){
	d.IDMap=d.NewObjectID()
	d.InitiateActionBuffer()
	c :=New(10)
	var typ int
	var bat *batiment.Batiment
	var res *ressource.Ressource
	if(c.GetSize()!=10){
		t.Error("Expected size 10, got ",c.GetSize())
	}
	for _,arr := range c.matrice{
		for _,tile :=range arr{
			tile.ExtractData(&typ,bat,res)
			if(typ!=0 || bat!=nil || res!=nil){
				t.Error("Expected an empty tile, got ",tile)
			}
		}
	}
}

//Verifie le placement de batiments et de ressources
func TestPlacement(t *testing.T){
	d.IDMap=d.NewObjectID()
	d.InitiateActionBuffer()
	res:=ressource.Create("tree",1,1)
	bat:=batiment.Create("auberge",2,2)
	c:=New(10)
	if(c.AddNewRessource(&res)!=true){
		t.Error("cannot add ressource")
	}
	if(c.GetTile(1,1).GetType()!=2){
		t.Error("ressource wasn't added")
	}
	if(c.AddNewBuilding(&bat)!=true){
		t.Error("cannot add building")
	}
	if(c.GetTile(2,2).GetType()!=1 || c.GetTile(3,2).GetType()!=1 || c.GetTile(3,3).GetType()!=1 || c.GetTile(2,3).GetType()!=1){
		t.Error("building wasn't added")
	}
}

func TestChemin(t *testing.T){
	d.IDMap=d.NewObjectID()
	d.InitiateActionBuffer()
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	c :=New(50)
	for k:=0;k<10;k++{
		px:=0
		py:=0
		x:=r1.Intn(50)
		y:=r1.Intn(50)
		path:=c.GetPathFromTo(0,0,x,y)
		if(path==nil){
			t.Error("Pas de chemin pour (0,0)->(",x,",",y,")")
		}
		for i:=0;i<len(path);i++{
			if path[i].x>px+1 || path[i].x <px-1 || path[i].y >py+1 || path[i].y<py-1 {
				t.Error("Chemin discontinu:",k," ",i," (",px,",",py,") -> (",path[i].x,",",path[i].y,") L=",len(path))
			}
			px=path[i].x
			py=path[i].y
		}
	}
}