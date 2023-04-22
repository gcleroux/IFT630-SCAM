package people

import (
	"sync"

	"github.com/gcleroux/IFT630-SCAM/pkg/registre"
)

var nbCitoyens int

func CitoyenInit(nb int) {
	nbCitoyens = nb
}

func CitoyenStep(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	// On demande au registre quel chantier rejoindre pour la journee
	batiment, err := registre.VisiteBatiment()

	if err != nil {
		// On a pas de batiment à visiter dans la journee
		return
	}

	// fmt.Println("[DEBUG] Le citoyen", id, "visite le batiment,", batiment.Name)

	// Envoi des ressources au maire
	Revenus <- batiment.Income
	Joie <- batiment.GenerationJoie
	Sante <- batiment.GenerationSante
}
