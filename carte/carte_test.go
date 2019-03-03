package carte

import "testing"
import batiment "git.unistra.fr/AOEINT/server/batiment"
import ressource "git.unistra.fr/AOEINT/server/ressource"

//Verifie si la carte creee est de la bonne taille et vide
func TestCreation(t *testing.T){
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