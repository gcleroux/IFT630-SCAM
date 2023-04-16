package people

import (
	"fmt"
	"sync"

	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
)

// La qte de travail qu'on ouvrier peut faire dans une journee
var nbOuvriers int
var travailOuvrier int

func OuvrierInit(nb int, travail int) {
	nbOuvriers = nb
	travailOuvrier = travail
}

func OuvrierStep(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	// On demande au registre quel chantier rejoindre pour la journee
	job, err := batiment.DemandeTravail(id)

	if err != nil {
		// On a pas de travail a faire pour la journee
		return
	}

	fmt.Println("Un ouvrier travaille sur le chantier du ", job.Batiment.Name)

	// On signale au registre qu'on a terminé pour la journee
	work := batiment.Travail{job.Id, travailOuvrier}
	batiment.JourneeTravail <- work
}
