// à remplir et vérifier la correcte implémentation des fonctions pour l'alpha.
package main

import "fmt"
import npc "server/npc"
import ressource "server/ressource"

func main(){
    res := ressource.Create(water,1,1)
    fmt.Println("Une ressource de type :" + res.type +
         " est crée sur la case x: " +res.x + " y: " res.y)
}
