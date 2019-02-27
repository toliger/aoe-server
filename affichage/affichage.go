package affichage

import "git.unistra.fr/AOEINT/server/carte"
import "github.com/fogleman/gg"

func ImprimerCarte(c carte.Carte){
	const TCase=10

	canvas := gg.NewContext(c.GetSize()*TCase,c.GetSize()*TCase)
	canvas.Clear()
	for i:=0; i<c.GetSize(); i++{
		for j:=0; j<c.GetSize(); j++{
			if(c.IsEmpty(i,j)){
				canvas.SetRGB(20,20,20)
			}else if(c.GetTile(i,j).GetType()==1){//Batiment=rouge
				canvas.SetRGB(255, 153, 153)
			}else{
				canvas.SetRGB(77, 148, 255) //ressource=bleu
			}
			canvas.DrawRectangle(float64(i*TCase), float64(j*TCase), TCase, TCase)
			canvas.Fill()
		}
	}

	canvas.SetRGB(0, 0, 0)
	canvas.SetLineWidth(1.0)
	for i:=0;i<c.GetSize();i++{
		canvas.DrawLine(float64(i*TCase),0,float64(i*TCase),float64(c.GetSize()*TCase))
		canvas.Stroke()
		canvas.DrawLine(0,float64(i*TCase),float64(c.GetSize()*TCase),float64(i*TCase))
		canvas.Stroke()
	}
	canvas.SavePNG("map.png")
}
