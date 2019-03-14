package constants

//Listes des differentes constantes de jeu
//constantes de debug
const (
	//Utilisation de la map 50*50 de test
	UseSmallMap = true
	//Mode verbose eventuel
	DEBUG = false
)
//constantes liees aux actions
const (
//NPC
//ActionNewNpc action creation npc
	ActionNewNpc = 0
//ActionDelNpc action destruction npc
	ActionDelNpc = 1
	// ActionHarmNpc degats infliges a un npc
	ActionHarmNpc = 2
	
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

//caracteristiques d'une partie

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
//Batiments

//LongueurAuberge longueur auberge
const LongueurAuberge = 2
//LargeurAuberge largeur auberge
const LargeurAuberge = 2
//LongueurCaserne longueur caserne
const LongueurCaserne = 2
//LargeurCaserne largeur caserne
const LargeurCaserne = 2
//LongueurEtabli longueur etabli
const LongueurEtabli = 2
//LargeurEtabli largeur etabli
const LargeurEtabli = 2
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
	GameUUID string
//APIHost API
	APIHost string
//Token Token
	Token string
//TokenSecret Token secret
	TokenSecret string
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
	RessourceBufferSize=50
)


//Default PNJ values

//HarvesterVillPortee  portee d'un harvester
const HarvesterVillPortee=1
//SoldierPv pv d'un  soldat
const SoldierPv=8
//SoldierVitesse vitesse d'un soldat
const SoldierVitesse=3
//SoldierVue vue d'un soldat
const SoldierVue=10
//SoldierPortee portee d'un soldat
const SoldierPortee=1
//SoldierSize taille d'un soldat
const SoldierSize=1
//SoldierDamage degats d'un soldat
const SoldierDamage=2
//HarvesterPv pv d'un harvester
const HarvesterPv=6
//HarvesterVitesse vitesse d'un harvester
const HarvesterVitesse=4
//HarvesterVue vue d'un harvester
const HarvesterVue=10
//HarvesterSize taille d'un harvester
const HarvesterSize=1
//HarvesterDamage degats d'un harvester
const HarvesterDamage=7
//VillagerPv pv d'un villager
const VillagerPv=4
//VillagerVitesse vitesse d'un villager
const VillagerVitesse=4
//VillagerVue vue d'un villager
const VillagerVue=10
//VillagerSize taille d'un villager
const VillagerSize=1
//VillagerDamage degats d'un villager
const VillagerDamage=5
