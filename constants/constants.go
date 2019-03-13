package constants

//Listes des differentes constantes de jeu
//constantes de debug
const (
	UseSmallMap = true
	DEBUG = false
)
//constantes liées aux actions
const (
//NPC
	ACTION_NEWNPC = 0
	ACTION_DELNPC = 1
	ACTION_HARMNPC = 2
//PLAYER
	ACTION_PLAYERRESSOURCE = 3
//RESSOURCE
	ACTION_HARMRESSOURCE = 4
	ACTION_DELRESSOURCE = 5
	ACTION_NEWRESSOURCE = 6
//BUILDING
	ACTION_HARMBUILDING = 7
	ACTION_DESTROYBUILDING = 8
	ACTION_NEWBUILDING = 9
//GAME
	ACTION_ENDOFGAME = 10
)
const MAXACTIONS = 10

//caracteristiques d'une partie

const MaxEntities = 100
const MaxBuildings = 20
const StartingSoldier = 6
const StartingHarvester = 0
const StartingVillager = 4
const StartingWood = 50
const StartingStone = 0
const StartingFood = 50
const MAXOBJECTS = 100
//Batiments

const LongueurAuberge = 2
const LargeurAuberge = 2
const LongueurCaserne = 2
const LargeurCaserne = 2
const LongueurEtabli = 2
const LargeurEtabli = 2
const PVAuberge = 100
const PVCaserne = 100
const PVEtabli = 100

//Couts de construction d'un Batiment par Ressource (à déterminer)

const PrixStoneAuberge = 5
const PrixWoodAuberge = 5
const PrixStoneCaserne = 5
const PrixWoodCaserne = 5
const PrixStoneEtabli = 5
const PrixWoodEtabli = 5

//DataRecup consts
var (
	GAME_UUID string
	API_HOST string
	TOKEN string
	TOKEN_SECRET string
)
const (
	DATAFILENAME = "server/data/Gamedata.json"
	//Valeurs par défaut uniquement si non fournies en variables d'environnement (4*)
	GAME_UUID_def = "DEFAULT"
	API_HOST_def = "DEFAULT"
	TOKEN_def = "DEFAULT"
	TOKEN_SECRET_def = "DEFAULT" 
	RESSOURCE_BUFFER_SIZE=50
)

//Default PNJ values

const HarvesterVillPortee=1

const SoldierPv=8
const SoldierVitesse=3
const SoldierVue=10
const SoldierPortee=1
const SoldierSize=1
const SoldierDamage=2

const HarvesterPv=6
const HarvesterVitesse=4
const HarvesterVue=10
const HarvesterSize=1
const HarvesterDamage=7

const VillagerPv=4
const VillagerVitesse=4
const VillagerVue=10
const VillagerSize=1
const VillagerDamage=5
