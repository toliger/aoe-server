// à remplir et vérifier la correcte implémentation des fonctions pour l'alpha.
package alpha

import "fmt"
//import npc "server/npc"
import ressource "server/ressource"
import carte "server/carte"

func AlphaTest(c carte.Carte){
    res := ressource.Create("tree",1,1)
    fmt.Println("Une ressource de type :", ressource.GetType(res),"est crée sur la case x:", ressource.GetX(res), "y:",ressource.GetY(res))
	var path []carte.Case
	path= c.GetPathFromTo(0,0,3,1)
	fmt.Printf("Chemin de (0,0) à (3,1) l=%d\n",len(path))
	for i:= range path{
		fmt.Printf("%d:(%d,%d)\n",i,path[i].GetPathX(),path[i].GetPathY())
	}
}
