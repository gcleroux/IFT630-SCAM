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
	people.MayorInit(conf.Budget)
	// people.OuvrierInit(conf.NbOuvrier)
	// people.CitoyenInit(conf.NbCitoyen)

	start := time.Now()

	for jour := 1; jour <= conf.NbJour; jour++ {
		// Le WaitGroup sert a synchroniser toutes les goroutine pour termine proprement une journee
		var wg sync.WaitGroup

		// Affichage de la journee
		fmt.Printf("\nJour #%d\n=========\n", jour)

		// Un channel qui sera ferme a la fin d'une journee
		done := utils.DayTime(time.Duration(conf.DayTime) * time.Second)

		wg.Add(1)
		go batiment.RegistreStep(&wg, done)

		wg.Add(1)
		go people.MayorStep(&wg)

		// On attend que tout le monde dans la ville termine sa journee
		wg.Wait()

	}

	// people.MayorEnd()

	elapsed := time.Since(start)
	fmt.Print("Temps total d'exécution du programme:")
	fmt.Println(elapsed)
}
