package tuile

import(
    "testing"
    d "git.unistra.fr/AOEINT/server/data"
    b "git.unistra.fr/AOEINT/server/batiment"
    "git.unistra.fr/AOEINT/server/ressource"
)

func TestCreation(t *testing.T){
    d.IDMap=d.NewObjectID()
	d.InitiateActionBuffer()
    tuile := New()
    rock := ressource.Create("rock",23,25)
    (&tuile).AddRessource(&rock)
    auberge1 := b.Create("auberge",5,5)
    (&tuile).AddBuilding(&auberge1)
    (&tuile).Empty()
    // rien à réellement tester pour l'instant
}
