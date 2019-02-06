package carte

import tuile "server/carte/tuile"
import "fmt"

type carte struct{
	size int
	matrice[][] tuile.Tuile
}
//Création de la carte
func New(size int) carte{
	var mat carte
	mat.size=size
	mat.matrice=make([][]tuile.Tuile,size)
	for i := 0; i < size; i++ {
		mat.matrice[i]= make([]tuile.Tuile,size)
		mat.matrice[i][i%size]=tuile.New()
	}
	return mat
}
//Affichage sur terminal
func Debug(mat carte){
	for i:=0;i<mat.size;i++{
		for j:=0;j<mat.size;j++{
			fmt.Print(tuile.GetType(mat.matrice[i][j]))
			fmt.Print(" ");
		}
		fmt.Println("")
	}
}