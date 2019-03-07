package tuile

import batiment "git.unistra.fr/AOEINT/server/batiment"
import ressource "git.unistra.fr/AOEINT/server/ressource"
type Tuile struct{//Batiment, ressource ou vide
	typ int //0 vide 1 batiment 2 ressource
	bat *batiment.Batiment
	res *ressource.Ressource
}

func New() Tuile{
	return (Tuile{0,nil,nil})
}
func (t Tuile) GetType() int{
	return t.typ
}

func (t Tuile)ExtractData(typ *int,bat *batiment.Batiment,res *ressource.Ressource){
	*typ=t.typ
	bat=t.bat
	res=t.res
}
func (t *Tuile)AddBuilding(bat *batiment.Batiment){
	(*t).typ=1
	(*t).bat=bat
	(*t).res=nil
}

func (t *Tuile)AddRessource(res *ressource.Ressource){
	(*t).typ=2
	(*t).res=res
	(*t).bat=nil
}
