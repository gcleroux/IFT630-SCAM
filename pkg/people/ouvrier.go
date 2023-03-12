package people

import (
	"fmt"

	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
)

var dateHier int = -1
var dateAuj int = 0

func Ouvrier(projets <-chan batiment.BatimentInfo, complets chan<- string, calendrier <-chan int /*, buildingBoard <-chan []batiment.BatimentInfo*/) {

	contrat := -1
	dateAuj = <-calendrier

	for {
		if dateHier < dateAuj { //Une fois par jour
			if contrat != -1 { //Si l'Ouvrier a déjà un contrat
				i := 0
				for commande := range projets { //Il essaie de retrouver son contrat
					if i == contrat { //Il retrouve son contrat
						if commande.ConstructionBatiment < commande.EffortBatiment { //Si le batiment n'est pas terminé d'être construit
							commande.ConstructionBatiment += 50
							fmt.Println("Un ouvrier travaille sur la construction du Batiment #", i, ":", commande.NomBatiment, ". Le bâtiment est maintenant à ", commande.ConstructionBatiment, "/", commande.EffortBatiment, "de construit.")
							if commande.ConstructionBatiment >= commande.EffortBatiment { //Si l'ouvrier fini la construction du batiment
								batiment.VilleContenu = append(batiment.VilleContenu, commande.NomBatiment)
								msg := fmt.Sprint("Un ouvrier a terminer la construction du Batiment #", i, ":", commande.NomBatiment)
								complets <- msg
								contrat = -1 //Il reset son contrat
							}
						}
						break //Il arrête de chercher son contrat
					} else {
						i++ //Il continue de chercher son contrat
					}
				}

			} else { //Si l'Ouvrier n'a pas déjà un contrat
				i := 0
				for commande := range projets { //Il passe à travers les contrats possible
					if commande.ConstructionBatiment < commande.EffortBatiment { //Il trouve le premier non-complété
						if commande.ConstructionBatiment == 0 { //S'il n'a pas encore été touché du tout
							commande.ConstructionBatiment += 50
							fmt.Println("Un ouvrier commence la construction du Batiment #", i, ":", commande.NomBatiment, ". Le bâtiment est maintenant à ", commande.ConstructionBatiment, "/", commande.EffortBatiment, "de construit.")
							contrat = i
							break
						} else { //Si le contrat a été commencé mais pas complété
							commande.ConstructionBatiment += 50
							if commande.ConstructionBatiment >= commande.EffortBatiment { //S'il termine la construction du Batiment
								batiment.VilleContenu = append(batiment.VilleContenu, commande.NomBatiment)
								msg := fmt.Sprint("Un ouvrier a terminer la construction du Batiment #", i, ":", commande.NomBatiment)
								complets <- msg
								contrat = -1 //Il reset son contrat
							} else { //S'il ne termine pas la construction du Batiment
								fmt.Println("Un ouvrier travaille sur la construction du Batiment #", i, ":", commande.NomBatiment, ". Le bâtiment est maintenant à ", commande.ConstructionBatiment, "/", commande.EffortBatiment, "de construit.")
								contrat = 1 //Il mémorise de quel batiment il s'agit pour y travailler le lendemain
							}
							break
						}
					} else {
						i++ // Il regarde le prochain contrat
					}
				}
			}
			dateHier++
		} else {
			//fmt.Println("Un ouvrier n'est pas en train de travailler")
		}
		//fmt.Println("Un ouvrier termine sa journée de travail")
	}

	// for commande := range projets {

	// 	var effortTotal = commande.EffortBatiment
	// 	fmt.Println("Un ouvrier commence la construction de : ", commande.NomBatiment)
	// 	for i := 0; i < effortTotal; i++ {
	// 		time.Sleep(time.Millisecond * 50)
	// 	}
	// 	batiment.VilleContenu = append(batiment.VilleContenu, commande.NomBatiment)

	// 	msg := fmt.Sprint("Un ouvrier a terminer la construction de : ", commande.NomBatiment)
	// 	complets <- msg
	// }
}
