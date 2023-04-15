package people

import (
	"fmt"

	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
)

func Ouvrier(complets chan<- string, calendrier <-chan int, numOuvrier int) {

	contrat := -1 //IdUniqueBatiment du batiment sur lequel l'ouvrier travaille.
	dateAuj = <-calendrier
	var dateHier int = -1
	var aTravailleAujourdhui bool = false

	for {
		if dateHier < dateAuj { //Une fois par jour
			if len(batiment.GetBuildingBoard()) == 0 { //Si le buildingBoard est vide
				dateHier++
				IncNbThreadPret()
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
						aTravailleAujourdhui = true
						break //Il arrête de chercher son contrat
						//} else {
						//	i++ //Il continue de chercher son contrat
					}
				}
				if !aTravailleAujourdhui {
					contrat = -1
					fmt.Println("L'ouvrier ", numOuvrier, " supervise alors qu'un autre ouvrier fini le building sur lequel il travaillait.")
				}
				aTravailleAujourdhui = false
				dateHier++
				IncNbThreadPret()
			} else { //Si l'Ouvrier n'a pas déjà un contrat
				var commande batiment.BatimentInfo
				for i, commandePossible := range batiment.GetBuildingBoard() { //Il passe à travers les contrats possibles
					if (float64)(commandePossible.ConstructionBatiment/commandePossible.EffortBatiment) < 0.5 { //Il trouve le premier construit à moins de 50%
						commande = commandePossible
						break
					} else if i == len(batiment.GetBuildingBoard()) { //S'il n'en trouve pas, il travaille sur le premier (le plus ancien)
						commande = batiment.GetBuildingBoard()[0]
						break
					}
				}
				codeRetour := batiment.Ajoute50Construction(commande.IdUniqueBatiment, numOuvrier) //L'ouvrier met 50 points de construction dans ce bâtiment.
				if codeRetour == 1 {                                                               //S'il ne termine pas la construction du Batiment
					contrat = commande.IdUniqueBatiment
				} else if codeRetour == 2 { //S'il termine la construction du Batiment
					batiment.VilleContenu = append(batiment.VilleContenu, commande.NomBatiment) //Le bâtiment est ajouté à la liste des bâtiment fini de la ville.
					batiment.RemoveFromBuildingBoard(commande.IdUniqueBatiment)                 //Remove from buildingBoard.
					contrat = -1                                                                //Il reset son contrat
					//complet <- msg
				}
				IncNbThreadPret()
				dateHier++
			}
		} else {
			//fmt.Println("Un ouvrier n'est pas en train de travailler")
		}
		//fmt.Println("Un ouvrier termine sa journée de travail")
	}
}
