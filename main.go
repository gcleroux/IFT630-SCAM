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
			go people.CitoyenStep(&wg)
		}

		// On attend que tout le monde dans la ville termine sa journee
		wg.Wait()

	}

	//TODO: Faire un cleanup des channels avec les fcts MayorEnd, CitoyenEnd, etc.

	fmt.Println("Liste des batiments dans la ville")
	fmt.Println("=================================")
	for _, b := range batiment.BatimentsVille {
		fmt.Println(b)
	}

	fmt.Println("\nTemps total d'exécution du programme:", time.Since(start))
}
