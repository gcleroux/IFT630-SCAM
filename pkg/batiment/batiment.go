package batiment

import "math"

type BatimentInfo struct {
	IdBatiment           int
	NomBatiment          string
	PrixBatiment         int
	EffortBatiment       int
	ConstructionBatiment int
	TauxProfitBatiment   float32
}

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
