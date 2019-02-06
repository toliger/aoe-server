package carte

import tuile "server/carte/tuile"
import "fmt"

type Carte struct{
	size int
	matrice[][] tuile.Tuile
}
//Création de la Carte
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
//Affichage sur terminal
func Debug(mat Carte){
	for i:=0;i<mat.size;i++{
		for j:=0;j<mat.size;j++{
			fmt.Print(tuile.GetType(mat.matrice[i][j]))
			fmt.Print(" ");
		}
		fmt.Println("")
	}
}