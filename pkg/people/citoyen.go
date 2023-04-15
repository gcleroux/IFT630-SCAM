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

func CitoyenStep(wg *sync.WaitGroup) {
	defer wg.Done()

	// On demande au registre quel chantier rejoindre pour la journee
	batiment, err := batiment.VisiteBatiment()

	if err != nil {
		// On a pas de batiment à visiter dans la journee
		return
	}

	fmt.Println("Un citoyen visite le batiment,", batiment.Name)

	// Envoi des revenus au maire
	Revenus <- batiment.Income
}

//
//
// type SafeCounter struct {
// 	mutex        sync.Mutex
// 	nombreVisite int
// }
//
// func (c *SafeCounter) Visite() {
// 	c.mutex.Lock()
// 	c.nombreVisite++
// 	c.mutex.Unlock()
// }
// func (c *SafeCounter) NombreVisites() int {
// 	nbVisites := 0
// 	c.mutex.Lock()
// 	nbVisites = c.nombreVisite
// 	c.mutex.Unlock()
// 	return nbVisites
// }
//
// type Citoyen struct {
// 	IdCitoyen int
// }
//
// var c SafeCounter
// var nombreVisite = 0
//
// func Population(idCitoyen int) {
// 	visiteCompletee := false
//
// 	for visiteCompletee == false {
// 		installations := batiment.VilleContenu
//
// 		if len(installations) < 2 {
// 			time.Sleep(time.Millisecond * 500)
// 		} else {
// 			choix := rand.Intn(len(installations))
// 			installation := batiment.ChoixBatiments[choix]
// 			dureeVisite := installation.EffortBatiment
// 			fmt.Println("Le citoyen", idCitoyen, "utilise les services offert par ", installation.NomBatiment)
// 			time.Sleep(time.Millisecond * time.Duration(dureeVisite))
// 			c.Visite()
// 			visiteCompletee = true
// 		}
// 	}
// }
//
// func NbVisites() int {
// 	return c.NombreVisites()
// }
