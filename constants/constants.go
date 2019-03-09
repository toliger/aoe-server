package constants
//Listes des differentes constantes de jeu
//constantes de debug

const (
	UseSmallMap = true
	DEBUG = false
)

//caracteristiques d'une partie

const MaxEntities = 100
const MaxBuildings = 20
const StartingSoldier = 6
const StartingHarvester = 0
const StartingVillager = 4
const StartingWood = 50
const StartingStone = 0
const StartingFood = 50
const MAXACTIONS = 20
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
)

//Default PNJ values
const SoldierRayon = 5
const VillagerRayon = 10
const HarvesterRayon = 8
const SoldierPortee = 1
const HarvesterPortee = 1
const VillagerPortee = 0
