package main

import (
	"fmt"
	"time"

	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
	"github.com/gcleroux/IFT630-SCAM/pkg/people"
)

func main() {
	start := time.Now()

	var budget int = 2000
	var nbOuvrier int = 3

	fmt.Println("Choix des batiments: ", batiment.ChoixBatiments)

	budget = people.MayorStart(budget, nbOuvrier)

	people.MayorEnd(budget)

	elapsed := time.Since(start)
	fmt.Print("Temps total d'exécution du programme:")
	fmt.Println(elapsed)
}
