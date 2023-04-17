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

// Channels
var Revenus = make(chan int)

func MayorInit(budget int) {
	nbProjets = 0
	budgetVille = budget
}

// Le maire fait des demandes de projets si le budget le permet
func MayorStep(wg *sync.WaitGroup, done <-chan interface{}) {
	defer wg.Done()
	// On retrouve la liste des batiments abordables
	abordables := batiment.GetBatimentsAbordables(budgetVille)

	if len(abordables) != 0 {
		// On choisit un batiment au hasard dans la liste des batiments abordables
		choix := abordables[rand.Intn(len(abordables))]
		nbProjets++
		budgetVille -= choix.Price
		fmt.Println("[MAYOR]: Le maire a demande la construction d'un", choix.Name)
		batiment.EnConstruction <- choix
	}

	for {
		select {
		case r := <-Revenus:
			budgetVille += r
		case <-done:
			// La journee est terminee
			return
		}
	}
}

// Fermer les channels
func MayorEnd() {
	close(Revenus)
}
