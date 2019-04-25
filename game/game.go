package game

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
	"sync"

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
	GameRunning bool
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
		datafileName = "../data/SmallTestMap.json"
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
	//On retire le pnj de la liste des pnj du jeu
	data.IDMap.DeleteObjectFromID(id)
	//On retire le pnj de la liste des pnj du joueur
	if !g.GetPlayerFromUID(pnj.PlayerUUID).DeleteNpcFromList(pnj.Get64X(), pnj.Get64Y(), pnj.GetType(), pnj.GetPv()) {
		return false
	}
	data.AddToAllAction(constants.ActionDelNpc, id, "useless", "useless")
	return true
}

//LaunchAutomaticFight : launch the AutomaticFight for all inactive npc
func (g *Game) LaunchAutomaticFight() {
	var wgObjectAcess sync.WaitGroup
	uptimeTicker := time.NewTicker(time.Duration(100 * time.Millisecond))
	for {
		select {
		case <-uptimeTicker.C:
			for _,player := range g.Joueurs{
				if ((*player).GetEntities() == nil){
					break
				}
				for i := 0; i < len((*player).GetEntities()); i++{
					// delete the dead npcs
					for _,playerTemp := range g.Joueurs{
						for _,pnjTemp := range (*playerTemp).GetEntities(){
							if (pnjTemp == nil){
								break
							}
							if (pnjTemp.GetPv() == 0){
								wgObjectAcess.Wait()
								g.DeleteNpc(pnjTemp)
							}
						}
					}
					if ((*player).GetPointerNpc(i) == nil){
						break
					}
					if ((*player).GetNpc(i).IsActive() == false){
						for _,p := range g.Joueurs{
							// log.Print("player :", (*player).GetUID())
							// log.Print("other player :", p.GetUID())
							//Seacch for ennemies npc
							if (player.GetUID() != p.GetUID()){
								pnjToFight := p.IsThereNpcInRange(player.GetPointerNpc(i))
								//log.Print("staticFight npc du type :", pnjToFight.GetType())
								if(pnjToFight != nil){
									go (*player).GetPointerNpc(i).StaticFightNpc(pnjToFight, &wgObjectAcess)
								}
							}
						}
					}
				}
			}
		}
	}
}


//EndOfGame : Interromps la boucle principale du jeu
func (g *Game) EndOfGame() {
	(*g).GameRunning = false
}

//GameLoop : fonction contenant la boucle principale du jeu
func (g *Game) GameLoop() {
	for (*g).GameRunning {
		time.Sleep(time.Duration(1000000000))
	}
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
			if i < 2 {
				(*g).Joueurs[i].AddAndCreateNpc("villager", i, j)
			} else {
				(*g).Joueurs[i].AddAndCreateNpc("villager", g.Carte.GetSize()-1, g.Carte.GetSize()-1)
			}
		}
	}
	//ajout de npc pour tester le combat
	 (*g).Joueurs[0].AddAndCreateNpc("villager", 0, 0)
	// (*g).Joueurs[0].AddAndCreateNpc("villager", 0, 0)
	//
	// (*g).Joueurs[1].AddAndCreateNpc("villager", 0, 1)
	(*g).Joueurs[0].AddAndCreateNpc("villager", 15, 15)
	(*g).Joueurs[1].AddAndCreateNpc("villager", 15, 16)
	(*g).Joueurs[0].AddAndCreateNpc("villager", 30, 31)
	(*g).Joueurs[1].AddAndCreateNpc("villager", 31, 30)
	(*g).Joueurs[0].AddAndCreateNpc("soldier", 20, 20)
	(*g).Joueurs[1].AddAndCreateNpc("soldier", 20, 21)

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
	j0 := joueur.Create(0, "Bob", "b33d954f-c63e-4b48-88eb-8b5e86d94246")
	j1 := joueur.Create(1, "Alice", "1982N19N2")
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
