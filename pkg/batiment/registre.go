package batiment

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sync"
)

// Tous les batiments disponibles dans la simulation
var TypesBatiments []Batiment = loadBatimentsInfos("./conf/batiments/")

// Batiments contenus dans la ville
var batimentsVille BatimentVille = BatimentVille{BatimentsVille: []Batiment{}}

// Tous les batiments activement en construction
var idProjet int = 0

var projets ProjetVille = ProjetVille{ProjetsVille: []Projet{}}

// On keep track de l'assignation des ouvriers
var jobBoardVille JobBoard

// Channels
var EnConstruction = make(chan Batiment)
var JourneeTravail = make(chan Travail)

// Le registre reste ouvert toute la journee
func RegistreStep(wg *sync.WaitGroup, done <-chan interface{}) {
	defer wg.Done()

	for {
		select {
		case b := <-EnConstruction:
			// On ajoute la demande du maire au projet en cours
			projets.Append(Projet{idProjet, b, 0})
			idProjet++
		case t := <-JourneeTravail:
			CheckWorkDone(t)
		case <-done:
			// La journee est terminee
			return
		}
	}
}

// Fermer les channels
func RegistreEnd() {
	close(EnConstruction)
	close(JourneeTravail)
}

// Trouve le prix le moins cher des batiments de la ville
func TrouveBatimentMoinsCher() int {
	min := math.MaxInt
	for _, b := range TypesBatiments {
		if b.Price < min {
			min = b.Price
		}
	}
	return min
}

// Retourne une liste des batiments qu'on peut se permettre selon un budget
func GetBatimentsAbordables(budget int) []Batiment {
	res := []Batiment{}

	for _, b := range TypesBatiments {
		if b.Price <= budget {
			res = append(res, b)
		}
	}
	return res
}

// Assigne un projet à un ouvrier, s'il n'est pas déjà sur un autre projet
func DemandeTravail(id int) (Projet, error) {
	projetsLength := projets.Length()
	if projetsLength == 0 {
		return Projet{}, errors.New("Pas de projet en cours")
	}

	// On regarde si l'ouvrier est deja associe a un projet
	if proj, ok := jobBoardVille.Get(id); ok {
		return proj, nil
	}

	// On assigne un nouveau projet a l'employe
	var newProj = projets.Get(rand.Intn(projetsLength))
	jobBoardVille.Set(id, newProj)

	//TODO: Il serait bien d'utilser capacite dans le batiment pour limiter le nombre d'ouvrier sur un projet
	return newProj, nil
}

func CheckWorkDone(t Travail) {
	//TODO: Maintenant qu'on a un jobBoard, on devrait plutot acceder au projet de cette facon
	for idx, p := range projets.GetAll() {
		if p.Id == t.Id {
			// On ajoute le travail de l'ouvrier au projet
			p.Travail += t.Effort
			projets.Set(idx, p)

			// Le batiment est complete, on l'enleve des projets pour le mettre dans les complets
			if p.Travail >= p.Batiment.Work {
				jobBoardVille.Delete(p.Id)

				fmt.Println("[REGISTRE]: La construction de", p.Batiment.Name, " (", p.Id, ") est terminée!")
				projets.Delete(idx)
				batimentsVille.Append(p.Batiment)
			}

		}
	}
}

func VisiteBatiment() (Batiment, error) {
	batimentsLength := batimentsVille.Length()
	if batimentsLength == 0 {
		return Batiment{}, errors.New("Pas de batiment dans la ville")
	}
	//TODO: Prendre en compte la capacite des batiments
	batiment := batimentsVille.Get(rand.Intn(batimentsLength))

	// On retourne un batiment a visiter au hasard
	return batiment, nil
}

// Retourne la liste des batiments de la ville
func GetBatiments() []Batiment {
	return batimentsVille.GetAll()
}
