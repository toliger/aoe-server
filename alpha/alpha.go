// à remplir et vérifier la correcte implémentation des fonctions pour l'alpha.
package alpha

import "fmt"
//import npc "server/npc"
import ressource "server/ressource"

func AlphaTest(){
    res := ressource.Create("water",1,1)
    fmt.Println("Une ressource de type :", ressource.GetType(res),"est crée sur la case x:", ressource.GetX(res), "y:",ressource.GetY(res))
}
