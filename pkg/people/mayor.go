package people

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
)

func MayorHello() string {
	return "Hello, Mayor!"
}

var villeContenu = []string{}

func MayorStart(budget int, nbOuvrier int) int {
	plusPetitPrix := batiment.TrouveBatimentMoinsCher()

	for budget > plusPetitPrix {
		batimentChoisi := rand.Intn(len(batiment.ChoixBatiments))
		if budget > batiment.ChoixBatiments[batimentChoisi].PrixBatiment {
			commande := batiment.ChoixBatiments[batimentChoisi]
			budget = budget - commande.PrixBatiment
			OuvrierConstruit(commande)
			villeContenu = append(villeContenu, commande.NomBatiment)
		}
	}
	return budget
}

func MayorEnd(budget int) {
	fmt.Println("Le Mayor prend sa retraite avec un budget restant de " + strconv.Itoa(budget) + "$")
	fmt.Println("La ville contient les bâtiments suivants:")
	for i := 0; i < len(villeContenu); i++ {
		fmt.Println(villeContenu[i])
	}
}
