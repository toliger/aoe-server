package game

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"git.unistra.fr/AOEINT/server/batiment"
	Carte "git.unistra.fr/AOEINT/server/carte"
	"git.unistra.fr/AOEINT/server/constants"
	"git.unistra.fr/AOEINT/server/data"
	"git.unistra.fr/AOEINT/server/joueur"
	"git.unistra.fr/AOEINT/server/npc"
	"git.unistra.fr/AOEINT/server/ressource"
	"git.unistra.fr/AOEINT/server/utils"
)

//Game : Structure contenant les donnees principales d'une partie
type Game struct {
	Joueurs     []*joueur.Joueur
	Carte       Carte.Carte
	GameRunning chan(bool)
}

//Data :Structure permettant de stocker les informations recuperees sur le fichier json
type Data struct {
	Size       int
	Buildings  []batiment.Batiment
	Ressources []ressource.Ressource
}

//ExtractData : extract data from a file (ressources, buildings)
func ExtractData() Data {
	datafileName := "data/GameData.json"
	if constants.UseSmallMap {
		datafileName = "data/SmallTestMap.json"
	}
	if constants.Testing {
		datafileName= "../" + datafileName
	}
	jsonFile, err := os.Open(datafileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var newGame Data
	err = json.Unmarshal(byteValue, &newGame)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err = jsonFile.Close()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return newGame
}

//GetPlayerFromUID : Permet de recuperer l'instance d'un joueur à partir de son uid
func (g *Game) GetPlayerFromUID(uid string) *joueur.Joueur {
	for i := 0; i < len(g.Joueurs); i++ {
		if g.Joueurs[i].UID == uid {
			return (g.Joueurs[i])
		}
	}
	return nil
}

//DeleteBuilding supprime un batiment, le retire de la liste du joueur et des ID, puis envoie une action
func (g *Game) DeleteBuilding(bat *batiment.Batiment) bool {
	//On recupere l'id du batiment
	id := data.IDMap.GetIDFromObject(bat)
	if id == "-1" {
		return false
	}
	//On retire le batiment de la liste des batiments du jeu
	data.IDMap.DeleteObjectFromID(id)
	//On retire le batiment de la liste du joueur
	if !g.GetPlayerFromUID(bat.PlayerUID).DeleteBatimentFromList(bat.X, bat.Y, bat.Typ) {
		return false
	}
	data.AddToAllAction(constants.ActionDestroyBuilding, id, "useless", "useless")
	bat = nil
	return true
}

//DeleteNpc Supprime un pnj, le retire de la liste du joueur et des ID, puis envoie une action DelNPC
func (g *Game) DeleteNpc(pnj *npc.Npc) bool {
	//On récupère l'id du npc
	id := data.IDMap.GetIDFromObject(pnj)
	if id == "-1" {
		return false
	}
	//On retire le pnj de la liste des pnj du joueur
	if !g.GetPlayerFromUID(pnj.PlayerUUID).DeleteNpcFromList(pnj.Get64X(), pnj.Get64Y(), pnj.GetType(), pnj.GetPv(),id) {
		return false
	}
	//On retire le pnj de la liste des pnj du jeu
	data.IDMap.DeleteObjectFromID(id)
	data.AddToAllAction(constants.ActionDelNpc, id, "useless", "useless")
	return true
}

//LaunchAutomaticFight : launch the AutomaticFight for all inactive npc
func (g *Game) LaunchAutomaticFight() {
	uptimeTicker := time.NewTicker(time.Duration(100 * time.Millisecond))
	deleteTicker := time.NewTicker(time.Duration(10 * time.Millisecond))
	for {
		select {
		case <- deleteTicker.C:
			// delete the dead npcs
			for _,player := range g.Joueurs{
				for _,pnj := range player.GetEntities(){
					if (pnj == nil){
						break
					}
					if (pnj.GetPv() <= 0){
						g.DeleteNpc(pnj)
						//log.Printf("delete pnj pos (%v, %v)", pnj.GetX(),pnj.GetY())
					}
				}
			}
		case <-uptimeTicker.C:
			for _,player := range g.Joueurs{
				if ((*player).GetEntities() == nil){
					break
				}
				for _,pnj := range player.GetEntities(){
					if pnj == nil{
						break
					}
					if pnj.IsActive() == false{
						for _,p := range g.Joueurs{
							//Search for ennemies npc
							if (player.GetUID() != p.GetUID()){
								pnjToFight := p.IsThereNpcInRange(pnj)
								if (pnjToFight != nil){
									//log.Printf("staticFight aggressor (%v, %v), target (%v, %v)", pnj.GetX(),pnj.GetY(), pnjToFight.GetX(),pnjToFight.GetY())
									go pnj.StaticFightNpc(pnjToFight)
								}else{
									buildingToFight := p.IsThereBuildingInRange(pnj)
									if (buildingToFight != nil){
										go pnj.StaticFightBuilding(buildingToFight)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

//BrokenBuildingsCollector deletes all buildings with less than 1 HP. Ends the game if a base is destroyed
func (g *Game)BrokenBuildingsCollector(){
	uptimeTicker := time.NewTicker(time.Duration(100 * time.Millisecond))
	for {
		select {
			case <-uptimeTicker.C:
				for _,player := range g.Joueurs{
					list:=player.GetBuildings()
					if(list==nil){
						break
					}
					for key,bat := range list{
						if bat.GetPv()<=0{
							g.DeleteBuilding(bat)
							if key==0{ //Auberge
								g.EndOfGame()
							}
						}
					}
				}
		}
	}
}

//EndOfGame : Interromps la boucle principale du jeu
func (g *Game) EndOfGame() {
	(*g).GameRunning <- false
	data.AddToAllAction(constants.ActionEndOfGame,"useless","useless","useless")
}

//GameLoop : fonction contenant la boucle principale du jeu
func (g *Game) GameLoop() {
	var test bool
	for{
		test =<-g.GameRunning
		if test==false{
			break
		}
	}
	time.Sleep(time.Duration(time.Second * constants.TimeBeforeExit))
}

//GenerateMap : Permet de generer la Carte a partir d'une structure data
func (g *Game) GenerateMap(data Data) {
	(*g).Carte = Carte.New(data.Size)
	//On attribue les auberges
	if len((*g).Joueurs) == 2 { //Si Seulement 2 Joueurs fournis, fait en sorte de leur donner des bases adverses

		if (*g).Carte.AddNewBuilding(&(data.Buildings[0])) == false {
			log.Fatal("Erreur lors du placement d'une auberge")
			os.Exit(1)
		}

		(*g).Joueurs[1].AddBuilding(&data.Buildings[2])
		if (*g).Carte.AddNewBuilding(&(data.Buildings[2])) == false {
			utils.Debug("Erreur lors du placement d'une auberge")
			os.Exit(1)
		}

	} else { //sinon 4 Joueurs classiques dans l'ordre des bases fournies (blue blue red red)
		for i := 0; i < 4; i++ {
			(*g).Joueurs[i].AddBuilding(&data.Buildings[i])
			if (*g).Carte.AddNewBuilding(&(data.Buildings[i])) == false {
				utils.Debug("Erreur lors du placement d'une auberge")
				os.Exit(1)
			}
		}
	}
 	//ajout des npc de base
	for i := 0; i < len((*g).Joueurs); i++ {
		for j := 0; j < 5; j++ {
			if i % 2 == 0 {
				(*g).Joueurs[i].AddAndCreateNpc("villager", i, j)
			} else {
				(*g).Joueurs[i].AddAndCreateNpc("villager", g.Carte.GetSize()/5-i, g.Carte.GetSize()/5-j)
			}
		}
	}
	//ajout de npc pour tester le combat
 	// (*g).Joueurs[0].AddAndCreateNpc("villager", 0, 0)
	// (*g).Joueurs[0].AddAndCreateNpc("villager", 0, 0)
	//
	// (*g).Joueurs[1].AddAndCreateNpc("villager", 0, 1)
	// (*g).Joueurs[0].AddAndCreateNpc("villager", 15, 15)
	// (*g).Joueurs[1].AddAndCreateNpc("villager", 15, 16)
	// (*g).Joueurs[0].AddAndCreateNpc("villager", 30, 31)
	// (*g).Joueurs[1].AddAndCreateNpc("villager", 31, 30)
	(*g).Joueurs[0].AddAndCreateNpc("soldier", 5, 5)
	(*g).Joueurs[1].AddAndCreateNpc("soldier", g.Carte.GetSize()/5 - 5, g.Carte.GetSize()/5 - 5)

	//Ajout des ressources
	for i := 0; i < len(data.Ressources); i++ {
		(&data.Ressources[i]).InitiatePV()
		if (*g).Carte.AddNewRessource(&(data.Ressources[i])) == false {
			utils.Debug("Erreur lors du placement d'une ressource")
			os.Exit(1)
		}
	}
}

/*GetPlayerData : Recupere les donnes des Joueurs entree en parametre du programme
Modification: Changement pour des valeurs statiques (temporaire)
*/
func (g *Game) GetPlayerData() {
	(*g).Joueurs = make([]*joueur.Joueur, 2)
	id1:= data.ExtractFromToken(constants.Player1JWT).UID
	id2:= data.ExtractFromToken(constants.Player2JWT).UID
	j0 := joueur.Create(0, "Bob", id1)
	j1 := joueur.Create(1, "Alice", id2)
	(*g).Joueurs[0] = &j0
	(*g).Joueurs[1] = &j1
	constants.PlayerUID1 = (*g).Joueurs[0].UID
	constants.PlayerUID2 = (*g).Joueurs[1].UID
	if len((*g).Joueurs) > 2 {
		constants.PlayerUID3 = (*g).Joueurs[2].UID
	}
	if len((*g).Joueurs) > 3 {
		constants.PlayerUID4 = (*g).Joueurs[3].UID
	}
	utils.Debug("joueurs:" + (*g).Joueurs[0].GetNom() + "" + (*g).Joueurs[1].GetNom())
}
