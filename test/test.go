// Ces tests permetent de vérifier la correcte implémentation des fonctions..
package test

import "fmt"
//import npc "server/npc"
import ressource "server/ressource"
import carte "server/carte"

func Test(c carte.Carte){
    res := ressource.Create("tree",1,1)
    fmt.Println("Une ressource de type :", res.GetType(),"est crée de coordonées x:", res.GetX(), "y:",res.GetY())
    fmt.Println("Ajout de cette ressource à la tuile correspondante")
    c.AddNewRessource(&res)
    fmt.Println("La ressource est ajoutée sur la case x:", res.GetX(), "y:",res.GetY())
    carte.Debug(c)

	var path []carte.Case
	path= c.GetPathFromTo(0,0,5,3)
	fmt.Printf("Chemin de (0,0) à (5,3) l=%d\n",len(path))
	for i:= range path{
		fmt.Printf("%d:(%d,%d)\n",i,path[i].GetPathX(),path[i].GetPathY())
	}
	carte.Debug(c)
	fmt.Println("On recalcule le chemin:\n")
	path= c.GetPathFromTo(0,0,5,3)
	fmt.Printf("Chemin de (0,0) à (5,3) l=%d\n",len(path))
	for i:= range path{
		fmt.Printf("%d:(%d,%d)\n",i,path[i].GetPathX(),path[i].GetPathY())
	}
}
