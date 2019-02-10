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
func (t Tuile)AddBuilding(bat *batiment.Batiment){
	t.typ=1
	t.bat=bat
	t.res=nil
}

func (t Tuile)AddRessource(res *ressource.Ressource){
	t.typ=2
	t.res=res
	t.bat=nil
}