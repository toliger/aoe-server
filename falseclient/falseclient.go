package falseclient
import tl "github.com/JoelOtter/termloop"

func StartClient(gameLoop *bool){
	game:=tl.NewGame()
	game.SetDebugOn(true)
	game.SetEndKey(tl.KeyArrowUp)
	screen:=game.Screen()
	//canvas:=tl.NewCanvas(64,64)
	screen.SetFps(30)
	x,y:=screen.Size()
	game.Log("debug:",game.DebugOn)
	game.Log("screen:",x,",",y)
	//game.Start()
	(*gameLoop)=false
}

//Fonctions a exporter dans client.go plus tard

