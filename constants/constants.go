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
	DEBUG      = false
	Player1JWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI5MDdmZjMwNS00OGRhLTRiMWEtYjI2Mi1hZWQxYzEwMzYzZjkiLCJpYXQiOjE1NTY3MjY3NjUsImV4cCI6MTU1NjczMzk2NX0.tcnpzN-ZDUOzvh10ovs1jCUAEW39j-nNfEFLQ5gSvhk"
	Player2JWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI2N2EyOGVlNi0xYTVjLTQ3NWQtYjY0Zi1hNGRjOGYwNDBkYzEiLCJpYXQiOjE1NTY3MjY4MTQsImV4cCI6MTU1NjczNDAxNH0.0Me2dBnn28ON6BvnO28sd2xeU4ub7hX_lCc99Dqs8BE"
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

	//ActionStartOfGame
	ActionStartOfGame = 11
)

//MaxActions constante maxActions
const MaxActions = 12

//ActionChannelSize Nombre d'actions simultanées stockées dans le channel
const ActionChannelSize = 150

//Epsilon : nombre pr comparaison float
const Epsilon = 0.0000000001

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

//TimeBeforeExit Time before closing the server after the end of a game
const TimeBeforeExit = 15

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
const PrixStoneCaserne = 50

//PrixWoodCaserne cout caserne en bois
const PrixWoodCaserne = 50

//PrixStoneEtabli cout etabli en stone
const PrixStoneEtabli = 0

//PrixWoodEtabli cout etabli en bois
const PrixWoodEtabli = 50

//DataRecup consts
var (
	//GameUUID id de partie
	GameUUID string = utils.Getenv("GAMEUUID", "2b6a0353-6a88-4060-93bc-f1208c623e80")

	//APIHost API
	APIHost string = utils.Getenv("APIHOST", "DEFAULT")

	//Token Token
	Token string = utils.Getenv("TOKEN", "DEFAULT")

	//TokenSecret Token secret
	TokenSecret string = utils.Getenv("TOKEN_SECRET", "DEFAULT")
)

const (
	//ExpiringTime time(in seconds) before the game closes if the correct amount of players is not obtained
	ExpiringTime = 60

	//MaxGameTime maximum duration of a game in seconds
	MaxGameTime = 600

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
const SoldierVitesse = 1

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
const HarvesterVitesse = 2

//HarvesterVue vue d'un harvester
const HarvesterVue = 10

//HarvesterSize taille d'un harvester
const HarvesterSize = 1

//VillagerPv pv d'un villager
const VillagerPv = 4

//VillagerVitesse vitesse d'un villager
const VillagerVitesse = 2

//VillagerVue vue d'un villager
const VillagerVue = 10

//VillagerSize taille d'un villager
const VillagerSize = 1

//VillagerDamage degats d'un villager
const VillagerDamage = 5

//VillagerFoodCost : cout en nourriture d'un villageois
const VillagerFoodCost = 25

//HarvesterFoodCost : cout en nourriture d'un harvester
const HarvesterFoodCost = 30

//HarvesterWoodCost : cout en bois d'un harvester
const HarvesterWoodCost = 20

//SoldierFoodCost : cout en nourriture d'un soldat
const SoldierFoodCost = 50

//SoldierWoodCost : cout en bois d'un soldat
const SoldierWoodCost = 10

//===== Default API values ======

//APIHOST default API host
var APIHOST = utils.Getenv("APIHOST", "https://game.api.archisme.com")

//APIPORT default API port
var APIPORT = utils.Getenv("APIPORT", "443")

//Testing used to enable functionalities for testing
var Testing = false
