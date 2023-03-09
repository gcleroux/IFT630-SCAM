package people

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
)

// type MayorLock struct {
// 	mutex  sync.Mutex
// 	budget int
// }

// func (c *MayorLock) Depense(montant int) bool {
// 	transaction := false
// 	c.mutex.Lock()
// 	if montant < c.budget {
// 		c.budget -= montant
// 		transaction = true
// 	}
// 	c.mutex.Unlock()
// 	return transaction
// }
// func (c *MayorLock) Profit(montant int) {
// 	c.mutex.Lock()
// 	c.budget += montant
// 	c.mutex.Unlock()
// }
// func (c *MayorLock) Budget() int {
// 	budgetCourant := 0
// 	c.mutex.Lock()
// 	budgetCourant = c.budget
// 	c.mutex.Unlock()
// 	return budgetCourant
// }

func MayorHello() string {
	return "Hello, Mayor!"
}

var budgetVille int

func MayorStart(budget int, nbOuvrier int, nbCitoyen int) {
	nbProjets := 0
	plusPetitPrix := batiment.TrouveBatimentMoinsCher()
	budgetVille = budget

	for i := 0; i < nbOuvrier; i++ {
		go Ouvrier(batiment.Projets, batiment.Complets)
	}

	for i := 0; i < nbCitoyen; i++ {
		go Population(i)
	}

	for budgetVille >= plusPetitPrix {
		batimentChoisi := rand.Intn(len(batiment.ChoixBatiments))
		if budgetVille > batiment.ChoixBatiments[batimentChoisi].PrixBatiment {
			commande := batiment.ChoixBatiments[batimentChoisi]
			budgetVille -= commande.PrixBatiment

			nbProjets++
			batiment.Projets <- commande
		}
	}
	close(batiment.Projets)

	for i := 0; i < nbProjets; i++ {
		msg := <-batiment.Complets
		fmt.Println(msg)
	}
	close(batiment.Complets)
}

func MayorEnd() {
	fmt.Println("Le Mayor prend sa retraite avec un budget restant de " + strconv.Itoa(budgetVille) + "$")
	fmt.Println("La ville contient les bâtiments suivants:")
	for i := 0; i < len(batiment.VilleContenu); i++ {
		fmt.Println(batiment.VilleContenu[i])
	}
	nbVisites := NbVisites()
	fmt.Println("La population à utiliser les services offert par la ville " + strconv.Itoa(nbVisites) + " fois")
}
