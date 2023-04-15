package batiment

import (
	"fmt"
	"math"
	"sync"
)

// Tous les batiments disponibles dans la simulation
var TypesBatiments []Batiment = loadBatimentsInfos("./conf/batiments/")

// Batiments contenus dans la ville
var BatimentsVille []Batiment = []Batiment{}

// Channels
var EnConstruction = make(chan Batiment, 100)

var Projects = make(chan Batiment, 100)
var Complets = make(chan Batiment, 100)
var Registre = make(chan Batiment, 100)

// Le registre reste ouvert toute la journee
func RegistreStep(wg *sync.WaitGroup, done <-chan interface{}) {
	defer wg.Done()

	for {
		select {
		case b := <-EnConstruction:
			fmt.Println("Le maire a demande un batiment", b.Name)
		case <-done:
			// La journee est terminee
			return
		}
	}
}

// Trouve le prix le moins cher des batiments de la ville
func TrouveBatimentMoinsCher() int {
	min := math.MaxInt
	for _, b := range TypesBatiments {
		if b.Price < min {
			min = b.Price
		}
	}
	return min
}

// Retourne une liste des batiments qu'on peut se permettre selon un budget
func GetBatimentsAbordables(budget int) []Batiment {
	res := []Batiment{}

	for _, b := range TypesBatiments {
		if b.Price <= budget {
			res = append(res, b)
		}
	}
	return res
}

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
