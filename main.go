package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
	"github.com/gcleroux/IFT630-SCAM/pkg/people"
	"github.com/gcleroux/IFT630-SCAM/pkg/utils"
)

func main() {
	// Load global config
	conf, err := utils.LoadConfig("./conf/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Init des personnages
	people.MayorInit(conf.Budget)
	people.OuvrierInit(conf.NbOuvrier, conf.TravailOuvrier)
	people.CitoyenInit(conf.NbCitoyen)

	//Start le timer
	start := time.Now()

	//Set up the Randomiser
	rand.Seed(time.Now().UnixNano())

	// Création des variables
	// qui déterminent le nombre de Goroutines
	var nbOuvriers int = conf.NbOuvrier
	var nbCitoyens int = conf.NbCitoyen
	// de statistiques
	var nbCitoyensPerdus int = 0
	// de ressources secondaires
	var tauxJoie int = 50
	var tauxSante int = 50

	for jour := 1; jour <= conf.NbJour; jour++ {
		// Le WaitGroup sert a synchroniser toutes les goroutine pour termine proprement une journee
		var wg sync.WaitGroup

		// Affichage de la journee
		fmt.Printf("\nJour #%d\n=========\n", jour)

		// Un obtient le channel qui sera ferme a la fin d'une journee
		// Les composantes qui sont dependants de la longueur d'une journee doivent
		// recevoir le channel en parametre
		done := utils.DayTime(time.Duration(conf.DayTime) * time.Second)

		//NOTE: Une facon alternative de gérer la sync serait d'avoir des tick comme sur un moteur de jeu.
		//      On pourrait avoir 5 ticks par jours, donc présumément 5 steps par jours et c'est ça qui assurerait
		//      la synchronisation des goroutine. J'ai mis du temps pour que ce soit plus facile a faire pour l'instant
		//      mais faire des ticks serait pas trop compliqué selon moi

		// Pour ajouter des types de personnes dans la simulation, on a juste a lancer une goroutine
		// avec une methode Step() pour qu'il soit intégré au pipeline

		wg.Add(1)
		go batiment.RegistreStep(&wg, done)

		wg.Add(1)
		go people.MayorStep(&wg, done)

		wg.Add(nbOuvriers)
		for i := 0; i < nbOuvriers; i++ {
			go people.OuvrierStep(&wg, i)
		}

		wg.Add(nbCitoyens)
		for i := 0; i < nbCitoyens; i++ {
			go people.CitoyenStep(&wg, i)
		}

		// De nouveaux ouvriers et citoyens sont potentiellement ajoutés
		if rand.Intn(100) < conf.TauxRecrutementOuvrier { // 0 <= f < 100, donc 10% d'ajouter un ouvrier
			if nbCitoyens > 0 {
				nbOuvriers += 1
				nbCitoyens -= 1
				fmt.Println("Un citoyen devient un ouvrier.")
			}
		}
		if rand.Intn(100) < conf.TauxNaissance { // 0 <= f < 100, donc 50% d'ajouter un citoyen
			nbCitoyens += 1
			fmt.Println("Un citoyen est né dans la métropole.")
		}

		// On attend que tout le monde dans la ville termine sa journee
		wg.Wait()

		//Si une Ressource Secondaire est à 0 à la fin d’une journée, quelques citoyens sont perdus.
		if tauxJoie < 10 {
			perte := rand.Intn(5) + 1 //1 à 5
			if nbCitoyens < perte {
				perte = nbCitoyens
			}
			nbCitoyens -= perte
			nbCitoyensPerdus += perte
			if perte == 0 {
				fmt.Println("Le taux de Joie dans la ville est à ", tauxJoie, "%. La ville n'a aucun citoyen a perdre.")
			} else if perte == 1 {
				fmt.Println("Le taux de Joie dans la ville est à ", tauxJoie, "%. ", perte, " citoyen est perdu.")
			} else {
				fmt.Println("Le taux de Joie dans la ville est à ", tauxJoie, "%. ", perte, " citoyens sont perdus.")
			}
		}
		if tauxSante < 10 {
			perte := rand.Intn(5) + 1 //2 à 5
			if nbCitoyens < perte {
				perte = nbCitoyens
			}
			nbCitoyens -= perte
			nbCitoyensPerdus += perte
			if perte == 0 {
				fmt.Println("Le taux de Sante dans la ville est à ", tauxSante, "%. La ville n'a aucun citoyen a perdre.")
			} else if perte == 1 {
				fmt.Println("Le taux de Sante dans la ville est à ", tauxSante, "%. ", perte, " citoyen est perdu.")
			} else {
				fmt.Println("Le taux de Sante dans la ville est à ", tauxSante, "%. ", perte, " citoyens sont perdus.")
			}
		}
	}

	// Faire un cleanup des channels avec les fonctions MayorEnd, CitoyenEnd, etc.
	people.MayorEnd()
	batiment.RegistreEnd()

	// Calcul du score
	Score := 0
	Score += nbCitoyens * 5
	Score -= nbCitoyensPerdus * 10
	Score += len(batiment.GetBatiments()) * 20
	//Score += moyenneJoie //TODO: Implémenter le calcul de la moyenne de la ressource Joie
	//Score += moyenneSante //TODO: Implémenter le calcul de la moyenne de la ressource Joie
	Score += people.GetBudgetVille() / 100

	fmt.Println()
	fmt.Println("=== Fin de la simulation ===")
	fmt.Println("Nombre de jours simulés: ", conf.NbJour)
	fmt.Println("Nombre final d'ouvriers: ", nbOuvriers)
	fmt.Println("Nombre final de citoyens: ", nbCitoyens)
	fmt.Println("Nombre de citoyens perdus: ", nbCitoyensPerdus)
	fmt.Println("Budget restant: ", people.GetBudgetVille())
	fmt.Println("Nombre de bâtiments construits: ", len(batiment.GetBatiments()))
	fmt.Println("Liste des batiments dans la ville:")
	for _, b := range batiment.GetBatiments() {
		fmt.Println(b)
	}
	fmt.Println("Score: ", Score) //TODO: Implémenter système de score

	fmt.Println("\nTemps total d'exécution du programme:", time.Since(start))
}
