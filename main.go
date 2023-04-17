package main

import (
	"fmt"
	"log"
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

	start := time.Now()

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

		wg.Add(conf.NbOuvrier)
		for i := 0; i < conf.NbOuvrier; i++ {
			go people.OuvrierStep(&wg, i)
		}

		wg.Add(conf.NbCitoyen)
		for i := 0; i < conf.NbCitoyen; i++ {
			go people.CitoyenStep(&wg, i)
		}

		// On attend que tout le monde dans la ville termine sa journee
		wg.Wait()

	}

	// Faire un cleanup des channels avec les fcts MayorEnd, CitoyenEnd, etc.
	people.MayorEnd()
	batiment.RegistreEnd()

	fmt.Println()
	fmt.Println("=== Fin de la simulation ===")
	fmt.Println("Nombre de jours simulés: ", conf.NbJour)
	fmt.Println("Nombre final de citoyens: ", conf.NbCitoyen) //TODO: augmenter lorsque de nouveaux citoyens sont ajoutés
	fmt.Println("Nombre final d'ouvriers: ", conf.NbOuvrier)  //TODO: augmenter lorsque de nouveaux ouvriers sont embauchés
	fmt.Println("Budget restant: ", people.GetBudgetVille())
	fmt.Println("Nombre de bâtiments construits: ", len(batiment.GetBatiments()))
	fmt.Println("Liste des batiments dans la ville:")
	for _, b := range batiment.GetBatiments() {
		fmt.Println(b)
	}
	fmt.Println("Score: ", 0) //TODO: Implémenter système de score

	fmt.Println("\nTemps total d'exécution du programme:", time.Since(start))
}
