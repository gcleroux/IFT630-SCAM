package people

import (
	"fmt"
	"time"

	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
)

func OuvrierConstruit(commande batiment.BatimentInfo) {
	var effortTotal = commande.EffortBatiment
	fmt.Println("L'ouvrier commence la construction de : ", commande.NomBatiment)
	for i := 0; i < effortTotal; i++ {
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Println("L'ouvrier a terminer la construction de : ", commande.NomBatiment)
	villeContenu = append(villeContenu, commande.NomBatiment)
}
