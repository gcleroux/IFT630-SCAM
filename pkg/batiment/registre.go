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
var BatimentsVille []Batiment = []Batiment{}

// Tous les batiments activement en construction
var idProjet int = 0
var Projets []Projet = []Projet{}

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
			Projets = append(Projets, Projet{idProjet, b, 0})
			idProjet++
		case t := <-JourneeTravail:
			CheckWorkDone(t)
		case <-done:
			// La journee est terminee
			return
		}
	}
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

func DemandeTravail() (Projet, error) {
	if len(Projets) == 0 {
		return Projet{}, errors.New("Pas de projet en cours")
	}
	//TODO: Il serait bien d'utilser capacite dans le batiment pour limiter le nombre d'ouvrier sur un projet
	// On retourne un projet au hasard
	return Projets[rand.Intn(len(Projets))], nil
}

func CheckWorkDone(t Travail) {
	for idx, p := range Projets {
		if p.Id == t.Id {
			// On ajoute le travail de l'ouvrier au projet
			p.Travail += t.Effort
			Projets[idx] = p

			// Le batiment est complete, on l'enleve des projets pour le mettre dans les complets
			if p.Travail >= p.Batiment.Work {
				fmt.Println("[REGISTRE]: La construction de", p.Batiment.Name, "est terminée!")
				Projets = append(Projets[:idx], Projets[idx+1:]...)
				BatimentsVille = append(BatimentsVille, p.Batiment)
			}
		}
	}
}

func VisiteBatiment() (Batiment, error) {
	if len(BatimentsVille) == 0 {
		return Batiment{}, errors.New("Pas de batiment dans la ville")
	}
	//TODO: Prendre en compte la capacite des batiments

	// On retourne un batiment a visiter au hasard
	return BatimentsVille[rand.Intn(len(BatimentsVille))], nil
}
