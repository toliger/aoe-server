package tuile

import batiment "server/batiment"
import ressource "server/ressource"

type Tuile struct{//Batiment, ressource ou vide
	typ int //0 vide 1 batiment 2 ressource
	bat *batiment.Batiment
	res *ressource.Ressource
}

func New() Tuile{
	return (Tuile{0,nil,nil})
}
func GetType(tuil Tuile) int{
	return tuil.typ
}

func Actualise(tuile Tuile){

}

