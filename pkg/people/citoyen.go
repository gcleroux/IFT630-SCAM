package people

import (
	"fmt"
	"sync"

	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
)

var nbCitoyens int

func CitoyenInit(nb int) {
	nbCitoyens = nb
}

func CitoyenStep(wg *sync.WaitGroup, id int, f func()) {
	defer wg.Done()

	// On demande au registre quel chantier rejoindre pour la journee
	batiment, err := batiment.VisiteBatiment()

	if err != nil {
		// On a pas de batiment à visiter dans la journee
		return
	}

	fmt.Println("Le citoyen", id, "visite le batiment,", batiment.Name)

	// Envoi des revenus au maire
	Revenus <- batiment.Income
	f()
}
