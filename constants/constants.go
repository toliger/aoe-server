package constants
//Listes des differentes constantes de jeu

//constantes de debug

const (
	UseSmallMap = true
)

//caracteristiques d'une partie

const MaxEntities=100
const MaxBuildings=20
const StartingSoldier=6
const StartingHarvester=0
const StartingVillager=4
const StartingWood=50
const StartingStone=0
const StartingFood=50

//Batiments

const LongueurAuberge=2
const LargeurAuberge=2
const LongueurCaserne=2
const LargeurCaserne=2
const LongueurEtabli=2
const LargeurEtabli=2
const PVAuberge=100
const PVCaserne=100
const PVEtabli=100

//Couts de construction d'un Batiment par Ressource (à déterminer)

const PrixStoneAuberge=5
const PrixWoodAuberge=5
const PrixStoneCaserne=5
const PrixWoodCaserne=5
const PrixStoneEtabli=5
const PrixWoodEtabli=5

//DataRecup consts

const DATAFILENAME="server/data/Gamedata.json"
//Default PNJ values
const SoldierRayon=5
const VillagerRayon=10
const HarvesterRayon=8
const SoldierPortee=1
const HarvesterPortee=1
const VillagerPortee=0
const TauxRecolte=2
