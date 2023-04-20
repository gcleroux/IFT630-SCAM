package people

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
)

func MayorHello() string {
	return "Hello, Mayor!"
}

// Attributs propres au maire de la ville
var nbProjets int
var budgetVille int
var joieJournaliere float64
var santeJournaliere float64

// Channels
var Revenus = make(chan int)
var Joie = make(chan float64)
var Sante = make(chan float64)

func MayorInit(budget int) {
	nbProjets = 0
	budgetVille = budget
}

// Le maire fait des demandes de projets si le budget le permet
func MayorStep(wg *sync.WaitGroup, done <-chan interface{}) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	// On retrouve la liste des batiments abordables
	abordables := batiment.GetBatimentsAbordables(budgetVille)

	if len(abordables) != 0 {
		// On choisit un batiment au hasard dans la liste des batiments abordables
		choix := abordables[rand.Intn(len(abordables))]
		nbProjets++
		budgetVille -= choix.Price
		fmt.Println("[MAYOR]: Le maire demande la construction d'un", choix.Name)
		batiment.EnConstruction <- choix
	}

	// Reset journalière à 0 des ressources secondaires
	joieJournaliere = 0.0
	santeJournaliere = 0.0

	for {
		select {
		case r := <-Revenus:
			// Revenu peut être négatif
			budgetVille += r
		case j := <-Joie:
			joieJournaliere += j
		case s := <-Sante:
			santeJournaliere += s
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

// Get le budgetVille
func GetBudgetVille() int {
	return budgetVille
}

// Get la joieJournaliere
func GetJoieJournaliere() float64 {
	return joieJournaliere
}

// Get la santeJournaliere
func GetSanteJournaliere() float64 {
	return santeJournaliere
}
