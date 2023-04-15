package batiment

import "fmt"

// Tous les batiments disponibles dans la simulation
var TypesBatiments []Batiment = loadBatimentsInfos("./conf/batiments/")

// Public functions
func RegistreBatiment() {
	// if len(typesBatiments) == 0 {
	// 	typesBatiments = loadBatimentsInfos("./conf/batiments/")
	// }

	fmt.Println(TypesBatiments)
}

// Channels
var Projects = make(chan Batiment, 100)
var Complets = make(chan Batiment, 100)
var Registre = make(chan Batiment, 100)

// func TrouveBatimentMoinsCher() int {
// 	plusPetitPrix := math.MaxInt
// 	for i := 0; i < len(ChoixBatiments); i++ {
// 		if ChoixBatiments[i].PrixBatiment < plusPetitPrix {
// 			plusPetitPrix = ChoixBatiments[i].PrixBatiment
// 		}
// 	}
// 	return plusPetitPrix
// }
//
// func GetRandomBatiment() Batiment {
// 	idx := rand.Intn(len(ChoixBatiments))
// 	return ChoixBatiments[idx]
// }
//
// func GetVilleContenu() []string {
// 	return v.villeContenu()
// }
//
// type safeVilleContenu struct {
// 	mutex        sync.Mutex
// 	VilleContenu []string
// }
//
// // Internal code
// var v safeVilleContenu
//
// var ChoixBatiments = []Batiment{
// 	{IdBatiment: 1, NomBatiment: "Parc", PrixBatiment: 250, EffortBatiment: 20, TauxProfitBatiment: 0},
// 	{IdBatiment: 2, NomBatiment: "Hopital", PrixBatiment: 500, EffortBatiment: 50, TauxProfitBatiment: 0.15},
// 	{IdBatiment: 3, NomBatiment: "Hotel", PrixBatiment: 400, EffortBatiment: 40, TauxProfitBatiment: 0.25},
// }
//
// func (v *safeVilleContenu) appendToVille(nom string) {
// 	v.mutex.Lock()
// 	v.VilleContenu = append(v.VilleContenu, nom)
// 	v.mutex.Unlock()
// }
//
// func (v *safeVilleContenu) villeContenu() []string {
// 	v.mutex.Lock()
// 	ret := v.VilleContenu
// 	v.mutex.Unlock()
// 	return ret
// }
