package people

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
)

type CitoyenLock struct {
	mutex        sync.Mutex
	nombreVisite int
}

func (c *CitoyenLock) Visite() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.nombreVisite++
	//c.mutex.Unlock()
}
func (c *CitoyenLock) NombreVisites() int {
	//nbVisites := 0
	c.mutex.Lock()
	// nbVisites = c.nombreVisite
	// c.mutex.Unlock()
	defer c.mutex.Unlock()
	return c.nombreVisite
	//return nbVisites
}

type Citoyen struct {
	IdCitoyen int
}

var c CitoyenLock
var nombreVisite = 0

func Visite(idCitoyen int) {

	for {
		installations := batiment.VilleContenu

		if len(installations) < 2 { // Pour l'instant si on met < 1, les citoyens vont tous dans le même batîment (trop long avant d'en avoir un 2e)
			time.Sleep(time.Millisecond * 500)
		} else {
			choix := installations[rand.Intn(len(installations))]

			var index int
			for i := range batiment.ChoixBatiments {
				if batiment.ChoixBatiments[i].NomBatiment == choix {
					index = i
				}
			}

			installation := batiment.ChoixBatiments[index]
			dureeVisite := installation.EffortBatiment
			fmt.Println("Le citoyen", idCitoyen, "utilise les services offert par ", installation.NomBatiment)
			time.Sleep(time.Millisecond * time.Duration(dureeVisite*50))
			c.Visite()
		}
	}
}

func NbVisites() int {
	return c.NombreVisites()
}
