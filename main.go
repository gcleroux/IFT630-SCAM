package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
	"github.com/gcleroux/IFT630-SCAM/pkg/people"
	"github.com/gcleroux/IFT630-SCAM/pkg/utils"
)

func main() {
	start := time.Now()
	conf, err := utils.LoadConfig("./conf/config.yml")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Choix des batiments: ", batiment.ChoixBatiments)

	for i := 0; i < conf.NbCitoyen; i++ {
		go people.Visite(i)
	}

	fmt.Println("Le Maire est embauché")
	people.MayorStart(conf.Budget, conf.NbOuvrier, conf.NbJours, conf.NbJoie, conf.NbSante, conf.NbCitoyen)

	people.MayorEnd(conf.NbJours, conf.NbJoie, conf.NbSante)

	elapsed := time.Since(start)
	fmt.Print("Temps total d'exécution du programme:")
	fmt.Println(elapsed)
}
