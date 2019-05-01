package carte

import (
  tuile "git.unistra.fr/AOEINT/server/carte/tuile"
  "git.unistra.fr/AOEINT/server/utils"
  "git.unistra.fr/AOEINT/server/batiment"
  "git.unistra.fr/AOEINT/server/ressource"
  "git.unistra.fr/AOEINT/server/data"
  "git.unistra.fr/AOEINT/server/constants"
)


//Carte : Structure detaillant une carte de tuiles
type Carte struct{
	size int
	matrice[][] tuile.Tuile
}


//New : Creation d'une carte de taille n [int]
func New(size int) Carte{
	var mat Carte
	mat.size=size
	mat.matrice=make([][]tuile.Tuile,size)

	for i := 0; i < size; i++ {
		mat.matrice[i]= make([]tuile.Tuile,size)
		mat.matrice[i][i%size]=tuile.New()
	}
	return mat
}


//IsEmpty : Verifie si la cade est vide
func (c Carte) IsEmpty(x int, y int) bool{
		return c.matrice[x][y].GetType() ==0
}


//GetSize : Renvoie la taille d'une carte
func (c Carte)GetSize() int{
	return c.size
}


//GetTile : Renvoie la tuile pour la position demandee
func (c Carte) GetTile(x int, y int) *tuile.Tuile{
	return &(c.matrice[x][y])
}


//AddNewRessource : Ajoute une ressource(pointee) a la carte
func (c Carte)AddNewRessource(res *ressource.Ressource) bool{
	x:=(*res).GetX()
	y:=(*res).GetY()

	if(!c.IsEmpty(x,y)){
		return false
	}

	(c.GetTile(x,y)).AddRessource(res)
	id:=(&data.IDMap).AddObject(res)
	(*res).Transmit(id,constants.ActionNewRessource)

	return true
}


//AddNewBuilding : Ajoute un batiment(pointe) a la carte
func (c Carte)AddNewBuilding(bat *batiment.Batiment) bool{
	x:=(*bat).GetX()
	y:=(*bat).GetY()

	for i:=0;i<(*bat).GetLongueur();i++{
		for j:=0;j<(*bat).GetLargeur();j++{
			if(!c.IsEmpty(x+i,y+j)){
				return false
			}
		}
	}

	for i:=0;i<(*bat).GetLongueur();i++{
		for j:=0;j<(*bat).GetLargeur();j++{
			(c.GetTile(x+i,y+j)).AddBuilding(bat)
		}
	}

	id:=(&data.IDMap).AddObject(bat)
	(*bat).Transmit(id)

	return true
}


//Debug : Affichage de debuguage sur terminal
func Debug(mat Carte){
	for i:=0;i<mat.size;i++{
		for j:=0;j<mat.size;j++{
			utils.Debug(string(mat.matrice[i][j].GetType()))
		}
		utils.Debug("")
	}
}


//Case : Structure d'une case, []Case correspond a un chemin
type Case struct{
	x int
	y int
	tile tuile.Tuile
}


//GetPathX : Renvoie la valeur X d'une case
func (c Case) GetPathX() int{
	return c.x
}


//GetPathY : Renvoie la valeur Y d'une case
func (c Case) GetPathY() int{
	return c.y
}


//GetPathTile : Get the tile designed by a case
func (c Case) GetPathTile() tuile.Tuile{
	return c.tile
}


func validCoords(x int,y int, size int) bool{
	return ((x>=0 && x<size) && (y>=0 && y<size))
}


func  printMatrix(weightMatrix [][]int){
	for i:=0;i<len(weightMatrix);i++{
		for j:=0;j<len(weightMatrix);j++{
			utils.Debug(string(weightMatrix[i][j]))
			utils.Debug(" ")
		}
		utils.Debug("")
	}
}


//UNVISITED valeur de case non parcourue lors de l'attribution de poids
const UNVISITED=-1


//OBSTACLE valeur de case contenant un obstacle lors de l'attribution de poids
const OBSTACLE=-2


func attribuerPoids(weightMatrix [][]int,x int,y int, step int) [][]int{
	if(validCoords(x,y,len(weightMatrix))){
		if(weightMatrix[x][y]==UNVISITED){
			weightMatrix[x][y]=step+1
		}
	}
	return weightMatrix
}


func shortestPathAux(weightMatrix [][]int,c Carte, x int, y int, currX *int, currY *int, step int, path []Case, modif *bool) bool{
	if(validCoords(x,y,len(weightMatrix))){
		if(weightMatrix[x][y]==step-1){
			path[step-1]=(Case{x,y,*(c.GetTile(x,y))})
			*currX=x
			*currY=y
			*modif=true
			if(step==1){
				return true
			}
		}
	}
	return false
}


//Selectionne le chemin le plus court a partir de la matrice des poids
func shortestPath(weightMatrix [][]int,destx int, desty int,c Carte,path []Case) []Case{
	if(weightMatrix[destx][desty]==UNVISITED || weightMatrix[destx][desty]==OBSTACLE){
		return (nil)
	}
	path=make([]Case,weightMatrix[destx][desty]+1)
	path[len(path)-1]=(Case{destx,desty,*(c.GetTile(destx,desty))})
	currX:=destx
	currY:=desty
	modif:=false

	for step:=weightMatrix[destx][desty]+1;step>0;step--{
		modif=false
		if(shortestPathAux(weightMatrix,c, currX, currY+1, &currX, &currY, step,path, &modif)){
			break
		}else if(!modif && shortestPathAux(weightMatrix,c, currX, currY-1, &currX, &currY, step,path,&modif)){
			break
		}else if(!modif && shortestPathAux(weightMatrix,c, currX+1, currY, &currX, &currY, step,path,&modif)){
			break
		}else if(!modif && shortestPathAux(weightMatrix,c, currX-1, currY, &currX, &currY, step,path,&modif)){
			break
		}else if(!modif && shortestPathAux(weightMatrix,c, currX-1, currY+1, &currX, &currY, step,path,&modif)){
			break
		}else if(!modif && shortestPathAux(weightMatrix,c, currX-1, currY-1, &currX, &currY, step,path,&modif)){
			break
		}else if(!modif && shortestPathAux(weightMatrix,c, currX+1, currY+1, &currX, &currY, step,path,&modif)){
			break
		}else if(!modif && shortestPathAux(weightMatrix,c, currX+1, currY-1, &currX, &currY, step,path,&modif)){
			break
		}
	}
	path[len(path)-1]=(Case{destx,desty,*(c.GetTile(destx,desty))})
	return path
}


//GetPathFromTo : Renvoie le chemin le plus court entre deux cases ou nil si inatteignable
func (c Carte) GetPathFromTo(x int, y int, destx int, desty int) []Case{
	var path []Case
	var weightMatrix [][]int
	weightMatrix=make([][]int,c.size)

	for i := 0; i < c.size; i++ {
		weightMatrix[i]= make([]int,c.size)
		for j:=0;j<c.size;j++{
			if(c.IsEmpty(i,j)){
				weightMatrix[i][j]=UNVISITED
			}else{
				weightMatrix[i][j]=OBSTACLE
			}
		}
	}

	weightMatrix[x][y]=0 //Case source visitee
	//Fin de l'initialisation

	for step:=0;step<c.size*c.size;step++{
		for i:=0;i<c.size;i++{
			for j:=0;j<c.size;j++{
				if(weightMatrix[i][j]==step){
					weightMatrix=attribuerPoids(weightMatrix,i,j+1,step) //Haut
					weightMatrix=attribuerPoids(weightMatrix,i+1,j,step) //Droite
					weightMatrix=attribuerPoids(weightMatrix,i,j-1,step) //Bas
					weightMatrix=attribuerPoids(weightMatrix,i-1,j,step) //Gauche
					weightMatrix=attribuerPoids(weightMatrix,i-1,j+1,step)//Gauch Bas
					weightMatrix=attribuerPoids(weightMatrix,i-1,j-1,step)//Gauche Haut
					weightMatrix=attribuerPoids(weightMatrix,i+1,j+1,step)//Droite Bas
					weightMatrix=attribuerPoids(weightMatrix,i+1,j-1,step)//Droite Haut
				}
			}
		}
	}

	//Une matrice de poids est cree
	return shortestPath(weightMatrix,destx, desty,c,path)
}
