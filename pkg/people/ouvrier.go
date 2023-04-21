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

// Un ouvrier fait une demande de travail pour la journée et signal au régistre lorsqu'il a terminé.
func OuvrierStep(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	// On demande au registre quel chantier rejoindre pour la journee
	job, err := batiment.DemandeTravail(id)

	if err != nil {
		// On a pas de travail a faire pour la journee
		fmt.Println("L'ouvrier", id, "n'a pas de travail pour la journée")
		return
	}

	fmt.Println("L'ouvrier", id, "travaille sur le chantier du ", job.Batiment.Name, job.Id)

	// On signale au registre qu'on a terminé pour la journee
	work := batiment.Travail{Id: job.Id, Effort: travailOuvrier}
	batiment.JourneeTravail <- work
}
