package constants

import (
	"git.unistra.fr/AOEINT/server/utils"
)

//Listes des differentes constantes de jeu
//constantes de debug
const (
	//Utilisation de la map 50*50 de test
	UseSmallMap = true

	//Mode verbose eventuel
	DEBUG = false
)

var (
	//PlayerUID1 Name of the first player
	PlayerUID1 = "DEFAULT"
	//PlayerUID2 Name of the first player
	PlayerUID2 = "DEFAULT"
	//PlayerUID3 Name of the first player
	PlayerUID3 = "DEFAULT"
	//PlayerUID4 Name of the first player
	PlayerUID4 = "DEFAULT"
)

//constantes liees aux actions
const (
	//NPC
	//ActionNewNpc action creation npc
	ActionNewNpc = 0

	//ActionDelNpc action destruction npc
	ActionDelNpc = 1

	// ActionAlterationNpc degats infliges a un npc
	ActionAlterationNpc = 2

	//PLAYER
	//ActionPlayerRessource maj d'un joueur
	ActionPlayerRessource = 3

	//RESSOURCE
	//ActionHarmRessource degats infliges a une ressource
	ActionHarmRessource = 4

	//ActionDelRessource destruction d'une ressource
	ActionDelRessource = 5

	//ActionNewRessource creation ressource
	ActionNewRessource = 6

	//RessourcePv
	RessourcePv = 100

	//BUILDING
	//ActionHarmBuilding degats infliges a un batiment
	ActionHarmBuilding = 7

	//ActionDestroyBuilding destruction batiment
	ActionDestroyBuilding = 8

	//ActionNewBuilding creation batiment
	ActionNewBuilding = 9

	//GAME
	//ActionEndOfGame fin de jeu
	ActionEndOfGame = 10
)

//MaxActions constante maxActions
const MaxActions = 10

//ActionChannelSize Nombre d'actions simultanées stockées dans le channel
const ActionChannelSize = 10
//===== caracteristiques d'une partie =====

//MaxEntities nb entites max par joueur
const MaxEntities = 100

//MaxBuildings nb batiments max par joueur
const MaxBuildings = 20

//StartingSoldier nb soldats en debut de partie
const StartingSoldier = 6

//StartingHarvester nb harvester en debut de partie
const StartingHarvester = 0

//StartingVillager nb villager en debut de partie
const StartingVillager = 4

//StartingWood nb ressource bois de depart
const StartingWood = 50

//StartingStone nb ressource stone de depart
const StartingStone = 0

//StartingFood nb ressource food de depart
const StartingFood = 50

//MAXOBJECTS nb objets du buffer objets
const MAXOBJECTS = 200

//===== Batiments =====

//LongueurAuberge longueur auberge
const LongueurAuberge = 1

//LargeurAuberge largeur auberge
const LargeurAuberge = 1

//LongueurCaserne longueur caserne
const LongueurCaserne = 1

//LargeurCaserne largeur caserne
const LargeurCaserne = 1

//LongueurEtabli longueur etabli
const LongueurEtabli = 1

//LargeurEtabli largeur etabli
const LargeurEtabli = 1

//PVAuberge pv auberge
const PVAuberge = 100

//PVCaserne pv caserne
const PVCaserne = 100

//PVEtabli pv etabli
const PVEtabli = 100

//Couts de construction d'un Batiment par Ressource (à déterminer)

//PrixStoneAuberge cout auberge en stone
const PrixStoneAuberge = -1

//PrixWoodAuberge cout auberge en bois
const PrixWoodAuberge = -1

//PrixStoneCaserne cout caserne en stone
const PrixStoneCaserne = 5

//PrixWoodCaserne cout caserne en bois
const PrixWoodCaserne = 5

//PrixStoneEtabli cout etabli en stone
const PrixStoneEtabli = 5

//PrixWoodEtabli cout etabli en bois
const PrixWoodEtabli = 5

//DataRecup consts
var (
	//GameUUID id de partie
	GameUUID string = utils.Getenv("GAMEUUID", "DEFAULT")

	//APIHost API
	APIHost string = utils.Getenv("APIHOST", "DEFAULT")

	//Token Token
	Token string = utils.Getenv("TOKEN", "DEFAULT")

	//TokenSecret Token secret
	TokenSecret string = utils.Getenv("TOKEN_SECRET", "DEFAULT")
)

const (
	//DataFileName : chemin vers les données de jeu
	DataFileName = "server/data/Gamedata.json"

	//Valeurs par defaut uniquement si non fournies en variables d'environnement (4*)

	//GameUUIDDef : id de partie par defaut
	GameUUIDDef = "DEFAULT"

	//APIHostDef : Valeur api par defaut
	APIHostDef = "DEFAULT"

	//TOKENDef valeur token par defaut
	TOKENDef = "DEFAULT"

	//TOKENSecretDef valeur token secret par defaut
	TOKENSecretDef = "DEFAULT"

	//RessourceBufferSize taille du buffer de ressources
	RessourceBufferSize = 50

	//BatimentBufferSize taille du buffer de batiments
	BatimentBufferSize = 50
)

//===== Default PNJ values ======

//HarvesterVillPortee  portee d'un harvester
const HarvesterVillPortee = 1

//MinimumDmg degats minimum d'un NPC
const MinimumDmg = 1

//TauxRecolteVill taux recolte villageois
const TauxRecolteVill = 5

//TauxRecolteHarvester taux recolte harvester
const TauxRecolteHarvester = 7

//SoldierPv pv d'un  soldat
const SoldierPv = 8

//SoldierVitesse vitesse d'un soldat
const SoldierVitesse = 3

//SoldierVue vue d'un soldat
const SoldierVue = 10

//SoldierPortee portee d'un soldat
const SoldierPortee = 1

//SoldierSize taille d'un soldat
const SoldierSize = 1

//SoldierDamage degats d'un soldat
const SoldierDamage = 2

//HarvesterPv pv d'un harvester
const HarvesterPv = 6

//HarvesterVitesse vitesse d'un harvester
const HarvesterVitesse = 4

//HarvesterVue vue d'un harvester
const HarvesterVue = 10

//HarvesterSize taille d'un harvester
const HarvesterSize = 1

//HarvesterDamage degats d'un harvester
const HarvesterDamage = 7

//VillagerPv pv d'un villager
const VillagerPv = 4

//VillagerVitesse vitesse d'un villager
const VillagerVitesse = 4

//VillagerVue vue d'un villager
const VillagerVue = 10

//VillagerSize taille d'un villager
const VillagerSize = 1

//VillagerDamage degats d'un villager
const VillagerDamage = 5

//===== Default API values ======

//APIHOST default API host
var APIHOST = utils.Getenv("APIHOST", "localhost")

//APIPORT default API port
var APIPORT = utils.Getenv("APIPORT", "4000")

//Testing used to enable functionalities for testing
var Testing = false