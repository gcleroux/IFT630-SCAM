package people

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
)

func MayorHello() string {
	return "Hello, Mayor!"
}

var budgetVille int
var joursActuel int

func MayorStart(budget int, nbOuvrier int, nbJours int, nbJoie int, nbSante int) {
	nbProjets := 0
	plusPetitPrix := batiment.TrouveBatimentMoinsCher()
	budgetVille = budget
	joursActuel = 0

	// Le Maire embauche les Ouvriers de départ
	fmt.Println("Le Maire embauche ", nbOuvrier, " Ouvriers.")
	for i := 0; i < nbOuvrier; i++ {
		go Ouvrier(batiment.Projets, batiment.Complets, batiment.Calendrier /*, batiment.BuildingBoard*/)
	}

	for joursActuel = 0; joursActuel < nbJours; joursActuel++ {
		fmt.Println("Le jour # ", joursActuel, " commence.")
		// Le maire commande un nouveau bâtiment par jour
		if budgetVille >= plusPetitPrix {
			batimentChoisi := rand.Intn(len(batiment.ChoixBatiments))
			if budgetVille > batiment.ChoixBatiments[batimentChoisi].PrixBatiment {
				fmt.Println("Le Maire commande la construction d'un ", batiment.ChoixBatiments[batimentChoisi].NomBatiment)
				commande := batiment.ChoixBatiments[batimentChoisi]
				budgetVille -= commande.PrixBatiment

				nbProjets++
				batiment.Projets <- commande
			}
		}

		// ???
		// for i := 0; i < nbProjets; i++ {
		// 	msg := <-batiment.Complets
		// 	fmt.Println(msg)
		// }

		// Le Maire informe les Ouvriers et Citoyens que c'est une nouvelle journée
		batiment.Calendrier <- joursActuel

		//Le Maire attend que les Ouvriers et Citoyens aient fini de travailler pour la journée
		time.Sleep(3 * time.Second)

		// À la fin de la journée, le Maire vérifie si une Ressource Secondaire est à 0, dans quel cas la ville s'effronde
		if nbJoie <= 0 || nbSante <= 0 {
			time.Sleep(3 * time.Second)
			MayorEnd(nbJours, nbJoie, nbSante)
			os.Exit(3)
		}

		//Le Maire commence la prochaine journée.
	}
	close(batiment.Projets)
	close(batiment.Complets)
}

func MayorEnd(nbJours int, nbJoie int, nbSante int) {
	if joursActuel < nbJours {
		fmt.Print("La ville s'effronde au jour "+strconv.Itoa(joursActuel), ",")
		if nbJoie <= 0 {
			fmt.Println(" car les Citoyens haïsse la ville et se sont révolté.")
		} else if nbSante <= 0 {
			fmt.Println(" car une pandémie a ravagé la ville.")
		}
	} else {
		fmt.Println("La ville se développe jusqu'au jour " + strconv.Itoa(joursActuel) + "et devient auto-suffisante!")
	}

	fmt.Println("Le Mayor prend sa retraite avec un budget restant de " + strconv.Itoa(budgetVille) + "$")
	fmt.Println("La ville contient les bâtiments suivants:")
	for i := 0; i < len(batiment.VilleContenu); i++ {
		fmt.Println(batiment.VilleContenu[i])
	}
	nbVisites := NbVisites()
	fmt.Println("La population à utiliser les services offert par la ville " + strconv.Itoa(nbVisites) + " fois")
}
