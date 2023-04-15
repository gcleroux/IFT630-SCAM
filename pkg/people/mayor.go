package people

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
)

func MayorHello() string {
	return "Hello, Mayor!"
}

// Attributs propres au maire de la ville
var nbProjets int
var budgetVille int

func MayorInit(budget int) {
	nbProjets = 0
	budgetVille = budget
}

func MayorStep(wg *sync.WaitGroup) {
	defer wg.Done()
	// On retrouve la liste des batiments abordables
	abordables := batiment.GetBatimentsAbordables(budgetVille)

	if len(abordables) == 0 {
		// On ne peut rien faire ajd
		return
	}

	// On choisit un batiment au hasard dans la liste des batiments abordables
	choix := abordables[rand.Intn(len(abordables))]
	nbProjets++
	budgetVille -= choix.Price
	fmt.Println("[MAYOR]: Le maire a demande la construction d'un", choix.Name)
	batiment.EnConstruction <- choix
}

// 	for i := 0; i < nbOuvrier; i++ {
// 		go Ouvrier(batiment.Projets, batiment.Complets)
// 	}
//
// 	for i := 0; i < nbCitoyen; i++ {
// 		go Population(i)
// 	}
//
// 	for budgetVille >= plusPetitPrix {
// 		batimentChoisi := rand.Intn(len(batiment.ChoixBatiments))
// 		if budgetVille > batiment.ChoixBatiments[batimentChoisi].PrixBatiment {
// 			commande := batiment.ChoixBatiments[batimentChoisi]
// 			budgetVille -= commande.PrixBatiment
//
// 			nbProjets++
// 			batiment.Projets <- commande
// 		}
// 	}
// 	close(batiment.Projets)
//
// 	for i := 0; i < nbProjets; i++ {
// 		msg := <-batiment.Complets
// 		fmt.Println(msg)
// 	}
// 	close(batiment.Complets)
// }

// func MayorEnd() {
// 	fmt.Println("Le Mayor prend sa retraite avec un budget restant de " + strconv.Itoa(budgetVille) + "$")
// 	fmt.Println("La ville contient les bâtiments suivants:")
// 	for i := 0; i < len(batiment.VilleContenu); i++ {
// 		fmt.Println(batiment.VilleContenu[i])
// 	}
// 	nbVisites := NbVisites()
// 	fmt.Println("La population à utiliser les services offert par la ville " + strconv.Itoa(nbVisites) + " fois")
// }
