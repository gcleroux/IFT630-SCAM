package batiment

import "math"

type BatimentInfo struct {
	IdBatiment     int
	NomBatiment    string
	PrixBatiment   int
	EffortBatiment int
}

var Projets = make(chan BatimentInfo, 100)
var Complets = make(chan string, 100)
var VilleContenu = []string{}

var ChoixBatiments = []BatimentInfo{
	{IdBatiment: 1, NomBatiment: "Parc", PrixBatiment: 250, EffortBatiment: 20},
	{IdBatiment: 2, NomBatiment: "Hopital", PrixBatiment: 500, EffortBatiment: 50},
	{IdBatiment: 3, NomBatiment: "Hotel", PrixBatiment: 400, EffortBatiment: 40}}

func TrouveBatimentMoinsCher() int {
	plusPetitPrix := math.MaxInt
	for i := 0; i < len(ChoixBatiments); i++ {
		if ChoixBatiments[i].PrixBatiment < plusPetitPrix {
			plusPetitPrix = ChoixBatiments[i].PrixBatiment
		}
	}
	return plusPetitPrix
}
