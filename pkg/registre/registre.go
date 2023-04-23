package registre

import (
	"errors"
	"fmt"
	"math"
	"sync"

	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
)

// Tous les batiments disponibles dans la simulation
var TypesBatiments []batiment.Batiment = batiment.LoadBatimentsInfos("./conf/batiments/")

// Batiments contenus dans la ville
var batimentsVille batiment.BatimentVille

// Tous les batiments activement en construction
var idProjet int = 0

// Les projets contenus dans la ville
var projets batiment.ProjetVille

// On keep track de l'assignation des ouvriers
var jobBoardVille batiment.JobBoard = *batiment.NewJobBoard()

// Le travail accompli par un travailleur dans une journée
var workUnitPerDay int

// Channels
var EnConstruction = make(chan batiment.Batiment)
var JourneeTravail = make(chan batiment.Travail)

func RegisterInit(workPerDay int) {
	workUnitPerDay = workPerDay
}

// Le registre reste ouvert toute la journee
func RegistreStep(wg *sync.WaitGroup, done <-chan interface{}) {
	defer wg.Done()

	for {
		select {
		case b := <-EnConstruction:
			// On ajoute la demande du maire au projet en cours
			projets.Append(batiment.Projet{Id: idProjet, Batiment: b, Travail: 0, Capacity: 0})
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
func GetBatimentsAbordables(budget int) []batiment.Batiment {
	res := []batiment.Batiment{}

	for _, b := range TypesBatiments {
		if b.Price <= budget {
			res = append(res, b)
		}
	}
	return res
}

// Retoune une liste des batiments qui génèrent de la joie
func GetBatimentJoyeux(abordables []batiment.Batiment) []batiment.Batiment {
	res := []batiment.Batiment{}

	for _, b := range abordables {
		if b.GenerationJoie > 0 {
			res = append(res, b)
		}
	}
	return res
}

// Retoune une liste des batiments qui génèrent de la joie
func GetBatimentSante(abordables []batiment.Batiment) []batiment.Batiment {
	res := []batiment.Batiment{}

	for _, b := range abordables {
		if b.GenerationSante > 0 {
			res = append(res, b)
		}
	}
	return res
}

// Retoune une liste des batiments qui génèrent de l'argent
func GetBatimentBudget(abordables []batiment.Batiment) []batiment.Batiment {
	res := []batiment.Batiment{}

	for _, b := range abordables {
		if b.Income > 0 {
			res = append(res, b)
		}
	}
	return res
}

// Assigne un projet à un ouvrier, s'il n'est pas déjà sur un autre projet
func DemandeTravail(idOuvrier int) (batiment.Projet, error) {
	if projets.Length() == 0 {
		return batiment.Projet{}, errors.New("Pas de projet en cours")
	}

	// On regarde si l'ouvrier est deja associe a un projet
	if proj, ok := jobBoardVille.Get(idOuvrier); ok {
		travailFait, err := projets.GetWorkDoneProjet(proj.Id)
		if err != nil {
			return batiment.Projet{}, err
		}

		travailAFaire, err := projets.GetWorkProjet(proj.Id)
		if err != nil {
			return batiment.Projet{}, err
		}

		// Vérifier si le projet est déjà terminé. Si non, l'ouvrier continue le projet
		if travailFait < travailAFaire {
			return proj, nil
		}
	}

	// On assigne un nouveau projet a l'ouvrier, s'il a un emploi de disponible
	newProj, err := projets.FindWork(idOuvrier, workUnitPerDay)

	if err != nil {
		jobBoardVille.Delete(idOuvrier) // Retirer le projet associé à l'ouvrier s'il existe
		return batiment.Projet{}, err
	}

	jobBoardVille.Set(idOuvrier, newProj)

	return newProj, nil
}

// Boucle sur tous les projets en cours et met à jout les projets en cours
//   - Ajoute le travail accompli durant la journée par les travailleurs
//   - Si un projet est terminé, il est retirer des projets en cours et les projet associés au travailleur dans
//     le jobBoard sont supprimés et ajoute dans les batiments de la ville.
func CheckWorkDone(t batiment.Travail) {
	//TODO: Maintenant qu'on a un jobBoard, on devrait plutot acceder au projet de cette facon
	for idx, p := range projets.GetAll() {
		if p.Id == t.Id {
			// On ajoute le travail de l'ouvrier au projet
			p.Travail += t.Effort
			projets.Set(idx, p)

			// Le batiment est complete, on l'enleve des projets pour le mettre dans les complets
			if p.Travail >= p.Batiment.Work {
				jobBoardVille.DeleteProjet(p.Id)

				fmt.Println("[REGISTRE]: La construction de", p.Batiment.Name, p.Id, "est terminée!")
				projets.Delete(idx)
				batimentsVille.Append(p.Batiment)
			}

		}
	}
}

func VisiteBatiment() (batiment.Batiment, error) {
	if batimentsVille.Length() == 0 {
		return batiment.Batiment{}, errors.New("Pas de batiment dans la ville")
	}

	// Le citoyen visite un batiment
	b, err := batimentsVille.Visite()

	if err != nil {
		return batiment.Batiment{}, err
	}

	// On retourne un batiment a visiter au hasard
	return b, nil
}

// Met à 0 le compte des visiteurs de chaque batiment
func ResetVisites() {
	batimentsVille.ResetVisites()
}

// Retourne la liste des batiments de la ville
func GetBatiments() []batiment.Batiment {
	return batimentsVille.GetAll()
}

func GetBatimentsList() map[string]int {
	return batimentsVille.GetBatimentsList()
}

func GetProjetsList() map[string]int {
	return projets.GetProjetsList()
}

func GetListeChantiers() (map[string]int, error) {
	listeChantiers, err := projets.GetListeChantiers()
	if err != nil {
		return listeChantiers, err
	}
	return listeChantiers, nil
}

func GetListeVisites() (map[string]int, error) {
	mapVisite := batimentsVille.GetListeVisites()
	empty := true
	for _, count := range mapVisite {
		if count > 0 {
			empty = false
		}
	}
	if empty {
		return map[string]int{}, errors.New("aucune visite durant la jounée")
	}
	return mapVisite, nil
}

func GetProjets() []batiment.Projet {
	return projets.GetAll()
}

func ProjetsGenereJoie() bool {
	return projets.GenereJoie()
}

func ProjetsGenereSante() bool {
	return projets.GenereSante()
}

func GetCapacitéEmploieVille() int {
	total := batimentsVille.CalculCapacitéEmploieVille()
	if total > 0 {
		return total
	}
	return 1
}
