package people

import (
	"fmt"

	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
)

var dateHier int = -1

func Ouvrier(projets <-chan batiment.BatimentInfo, complets chan<- string, calendrier <-chan int /*, buildingBoard <-chan []batiment.BatimentInfo*/, numOuvrier int) {

	contrat := -1 //IdUniqueBatiment du batiment sur lequel l'ouvrier travaille.
	dateAuj = <-calendrier

	for {
		//dateAuj = GetDateAuj()
		if dateHier < dateAuj { //Une fois par jour
			if len(batiment.GetBuildingBoard()) == 0 {
				dateHier++
			} else if contrat != -1 { //Si l'Ouvrier a déjà un contrat
				for _, commande := range batiment.GetBuildingBoard() { //Il essaie de retrouver son contrat
					if contrat == commande.IdUniqueBatiment { //Il retrouve son contrat
						if commande.ConstructionBatiment < commande.EffortBatiment { //Si le batiment n'est pas terminé d'être construit
							codeRetour := batiment.Ajoute50Construction(contrat, numOuvrier) //L'ouvrier met 50 points de construction dans ce bâtiment.
							if codeRetour == 2 {                                             //Si l'ouvrier fini la construction du batiment
								batiment.VilleContenu = append(batiment.VilleContenu, commande.NomBatiment) //Le bâtiment est ajouté à la liste des bâtiment fini de la ville.
								batiment.RemoveFromBuildingBoard(commande.IdUniqueBatiment)                 //Remove from buildingBoard.
								contrat = -1                                                                //Il reset son contrat
							} else if codeRetour == 1 {

							} else {
								fmt.Println("Erreur dans la logique d'un ouvrier")
							}
						}
						break //Il arrête de chercher son contrat
						//} else {
						//	i++ //Il continue de chercher son contrat
					}
				}
				contrat = -1
			} else { //Si l'Ouvrier n'a pas déjà un contrat
				//i := 0
				for _, commande := range batiment.GetBuildingBoard() { //Il passe à travers les contrats possibles
					if commande.ConstructionBatiment < commande.EffortBatiment { //Il trouve le premier non-complété
						codeRetour := batiment.Ajoute50Construction(commande.IdUniqueBatiment, numOuvrier) //L'ouvrier met 50 points de construction dans ce bâtiment.
						if codeRetour == 1 {                                                               //S'il ne termine pas la construction du Batiment
							contrat = commande.IdUniqueBatiment
						} else if codeRetour == 2 { //S'il termine la construction du Batiment
							batiment.VilleContenu = append(batiment.VilleContenu, commande.NomBatiment) //Le bâtiment est ajouté à la liste des bâtiment fini de la ville.
							batiment.RemoveFromBuildingBoard(commande.IdUniqueBatiment)                 //Remove from buildingBoard.
							contrat = -1                                                                //Il reset son contrat
							//complet <- msg
						}
						break
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
