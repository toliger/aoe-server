package batiment

import(
    "testing"
    d "git.unistra.fr/AOEINT/server/data"
)

func TestCreation(t *testing.T){
    d.IDMap=d.NewObjectID()
	d.InitiateActionBuffer()
    caserne := Create("caserne", 15,15)
    caserne.GetPv()
    (&caserne).DestroyBuilding()
    // rien à réellement tester pour l'instant
}
