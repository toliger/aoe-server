package tuile

import batiment "git.unistra.fr/AOEINT/server/batiment"
import ressource "git.unistra.fr/AOEINT/server/ressource"

//Tuile : Structure d'une tuile(case de la carte)
type Tuile struct{//Batiment, ressource ou vide
	typ int //0 vide 1 batiment 2 ressource
	bat *batiment.Batiment
	res *ressource.Ressource
	//entities []*npc.Npc
}

//New : création d'une tuile
func New() Tuile{
	return (Tuile{0,nil,nil})
}

//GetType : renvoie le type d'une tuile
func (t Tuile) GetType() int{
	return t.typ
}

//GetRess : renvoie la ressource contenue par une tuile(nil si aucune ressource) 
func (t Tuile) GetRess() *ressource.Ressource{
	return t.res
}

//ExtractData : Permet de récupérer le type, la ressource et le batiment d'une tuile par pointeurs
func (t Tuile)ExtractData(typ *int,bat *batiment.Batiment,res *ressource.Ressource){
	*typ=t.typ
	bat=t.bat
	res=t.res
}

//AddBuilding : Ajoute un pointeur sur batiment à une tuile
func (t *Tuile)AddBuilding(bat *batiment.Batiment){
	(*t).typ=1
	(*t).bat=bat
	(*t).res=nil
}

//AddRessource : Ajoute un pointeur sur ressource à une tuile
func (t *Tuile)AddRessource(res *ressource.Ressource){
	(*t).typ=2
	(*t).res=res
	(*t).bat=nil
}

//Empty : Vide une tuile de son contenu
func (t *Tuile)Empty(){
	(*t).typ=0
	(*t).bat=nil
	(*t).res=nil
}