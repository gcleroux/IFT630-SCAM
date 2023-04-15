package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gcleroux/IFT630-SCAM/pkg/people"
	"github.com/gcleroux/IFT630-SCAM/pkg/utils"
)

func main() {
	// Load global config
	conf, err := utils.LoadConfig("./conf/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	people.MayorStart(conf.Budget, conf.NbOuvrier, conf.NbCitoyen)
	people.MayorEnd()

	elapsed := time.Since(start)
	fmt.Print("Temps total d'exécution du programme:")
	fmt.Println(elapsed)
}
