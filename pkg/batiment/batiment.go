package batiment

import (
	"fmt"
	"math"
)

type BatimentInfo struct {
	IdBatiment           int
	IdUniqueBatiment     int
	NomBatiment          string
	PrixBatiment         int
	EffortBatiment       int
	ConstructionBatiment int
	TauxProfitBatiment   float32
}

var buildingBoard []BatimentInfo
var Projets = make(chan BatimentInfo, 100)
var Complets = make(chan string, 100)
var Calendrier = make(chan int, 100)
var VilleContenu = []string{}

var ChoixBatiments = []BatimentInfo{
	{IdBatiment: 1, NomBatiment: "Parc", PrixBatiment: 250, EffortBatiment: 200, TauxProfitBatiment: 0},
	{IdBatiment: 2, NomBatiment: "Hopital", PrixBatiment: 500, EffortBatiment: 500, TauxProfitBatiment: 0.15},
	{IdBatiment: 3, NomBatiment: "Hotel", PrixBatiment: 400, EffortBatiment: 400, TauxProfitBatiment: 0.25},
	{IdBatiment: 4, NomBatiment: "Bar", PrixBatiment: 300, EffortBatiment: 250, TauxProfitBatiment: 0.10},
}

func TrouveBatimentMoinsCher() int {
	plusPetitPrix := math.MaxInt
	for i := 0; i < len(ChoixBatiments); i++ {
		if ChoixBatiments[i].PrixBatiment < plusPetitPrix {
			plusPetitPrix = ChoixBatiments[i].PrixBatiment
		}
	}
	return plusPetitPrix
}

func GetBuildingBoard() []BatimentInfo {
	return buildingBoard
}

func SetBuildingBoard(bBrecu []BatimentInfo) {
	buildingBoard = bBrecu
}

// Trouvez le batiment avec cet idUniqueRecu et lui ajouter 50 points de construction.
func Ajoute50Construction(idUniqueRecu int, numOuvrier int) int {
	for i, commande := range buildingBoard { //Il essaie de retrouver son contrat
		if idUniqueRecu == commande.IdUniqueBatiment { //Il retrouve son contrat
			buildingBoard[i].ConstructionBatiment += 50
			if commande.ConstructionBatiment == 0 {
				fmt.Println("L'ouvrier ", numOuvrier, " commence la construction du Batiment #", i, ":", commande.NomBatiment, ". Le bâtiment est maintenant à ", buildingBoard[i].ConstructionBatiment, "/", commande.EffortBatiment, "de construit.")
				return 1
			}
			if buildingBoard[i].ConstructionBatiment >= commande.EffortBatiment {
				fmt.Println("L'ouvrier ", numOuvrier, " a terminer la construction du Batiment #", i, ":", commande.NomBatiment)
				return 2
			} else {
				fmt.Println("L'ouvrier ", numOuvrier, " travaille sur la construction du Batiment #", i, ":", commande.NomBatiment, ". Le bâtiment est maintenant à ", buildingBoard[i].ConstructionBatiment, "/", commande.EffortBatiment, "de construit.")
				return 1
			}
		}
	}
	return -1
}

func RemoveFromBuildingBoard(idUniqueRecu int) {
	buildingBoard[idUniqueRecu] = buildingBoard[len(buildingBoard)-1]
	buildingBoard = buildingBoard[:len(buildingBoard)-1]
}
