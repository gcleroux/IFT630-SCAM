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

	budget := people.MayorStart(conf.Budget, conf.NbOuvrier)

	people.MayorEnd(budget)

	elapsed := time.Since(start)
	fmt.Print("Temps total d'exécution du programme:")
	fmt.Println(elapsed)
}
