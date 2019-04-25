package ressource

import (
    "testing"
    d "git.unistra.fr/AOEINT/server/data"
    "git.unistra.fr/AOEINT/server/constants"
)

func TestCreation(t *testing.T){
    d.IDMap=d.NewObjectID()
	d.InitiateActionBuffer()
    tree := Create("tree",10,10)
    d.IDMap.AddObject(&tree)
    (&tree).InitiatePV();
    (&tree).Damage(5);
    if ((&tree).GetPv() != (constants.RessourcePv - 5) ){
        t.Error("the ressource didn't lose the good amount of HP")
    }
}
