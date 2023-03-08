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

func MayorStart(budget int, nbOuvrier int) int {
	nbProjets := 0
	plusPetitPrix := batiment.TrouveBatimentMoinsCher()

	for i := 0; i < nbOuvrier; i++ {
		go Ouvrier(batiment.Projets, batiment.Complets)
	}

	for budget >= plusPetitPrix {
		batimentChoisi := rand.Intn(len(batiment.ChoixBatiments))
		if budget > batiment.ChoixBatiments[batimentChoisi].PrixBatiment {
			commande := batiment.ChoixBatiments[batimentChoisi]
			budget = budget - commande.PrixBatiment

			nbProjets++
			batiment.Projets <- commande

		}
	}
	close(batiment.Projets)

	for i := 0; i < nbProjets; i++ {
		msg := <-batiment.Complets
		fmt.Println(msg)
	}
	close(batiment.Complets)
	return budget
}

func MayorEnd(budget int) {
	fmt.Println("Le Mayor prend sa retraite avec un budget restant de " + strconv.Itoa(budget) + "$")
	fmt.Println("La ville contient les bâtiments suivants:")
	for i := 0; i < len(batiment.VilleContenu); i++ {
		fmt.Println(batiment.VilleContenu[i])
	}
}
