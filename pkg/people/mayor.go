package people

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gcleroux/IFT630-SCAM/pkg/registre"
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
func MayorStep(wg *sync.WaitGroup, done <-chan interface{}, tauxSante float64, tauxJoie float64, nbCitoyenMatin int) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	// On retrouve la liste des batiments abordables
	abordables := registre.GetBatimentsAbordables(budgetVille)

	tauxEmploiment := float64(nbCitoyenMatin) / float64(registre.GetCapacitéEmploieVille())

	if len(abordables) != 0 {
		if tauxJoie < 60 || tauxSante < 60 {
			if tauxJoie < tauxSante && !registre.ProjetsGenereJoie() {
				// On retrouve la liste des batiments qui augmente la joie
				batimentJoie := registre.GetBatimentJoyeux(abordables)
				if len(batimentJoie) != 0 {
					choix := batimentJoie[rand.Intn(len(batimentJoie))]
					nbProjets++
					budgetVille -= choix.Price
					fmt.Println("[MAYOR]: Le maire demande la construction d'un", choix.Name)
					registre.EnConstruction <- choix
				}
			} else if !registre.ProjetsGenereSante() {
				// On retrouve la liste des batiments qui augmente la sante
				batimentSante := registre.GetBatimentSante(abordables)
				if len(batimentSante) != 0 {
					// fmt.Println("[DEBUG]: len(batimentSante)", len(batimentSante))
					// fmt.Println("[DEBUG]: batimentSante", batimentSante)
					choix := batimentSante[rand.Intn(len(batimentSante))]
					nbProjets++
					budgetVille -= choix.Price
					fmt.Println("[MAYOR]: Le maire demande la construction d'un", choix.Name)
					registre.EnConstruction <- choix
				}
			}
		} else if tauxEmploiment > 0.90 && len(registre.GetProjets()) == 0 {
			// On choisit un batiment qui produit de l'argent dans la liste des batiments abordables
			batimentBudget := registre.GetBatimentBudget(abordables)
			if len(batimentBudget) != 0 {
				choix := batimentBudget[rand.Intn(len(batimentBudget))]
				nbProjets++
				budgetVille -= choix.Price
				fmt.Println("[MAYOR]: Le maire demande la construction d'un", choix.Name)
				registre.EnConstruction <- choix
			}
		} else {
			//On ne bati rien pour l'instant
			fmt.Println("[MAYOR]: Le maire choisi de ne rien commander. Le budget, la joie et la sante sont assez haut.")
		}
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
