package batiment

import (
	"errors"
	"math/rand"
	"sync"
)

// Structure d'exclusion mutuelle read/write pour gérer les proejts de la ville qui sont partagés par différents threads.
// sync.RWMutex permet la lecture simultané de plusieurs processus ou un seul processus en écriture.
type ProjetVille struct {
	projetsVille      []Projet
	projetsVilleMutex sync.RWMutex
}

// Ajouter un nouveau projet
func (projets *ProjetVille) Append(nouveauProjet Projet) {
	projets.projetsVilleMutex.Lock()
	defer projets.projetsVilleMutex.Unlock()
	projets.projetsVille = append(projets.projetsVille, nouveauProjet)
}

// Supprimer un projet à l'index indiqué
func (projets *ProjetVille) Delete(index int) {
	projets.projetsVilleMutex.Lock()
	defer projets.projetsVilleMutex.Unlock()
	projets.projetsVille = append(projets.projetsVille[:index], projets.projetsVille[index+1:]...)
}

// Retourne le nombre de projets dans le tableau
func (projets *ProjetVille) Length() int {
	projets.projetsVilleMutex.RLock()
	defer projets.projetsVilleMutex.RUnlock()
	return len(projets.projetsVille)
}

// Retourne le travail accomplit sur un projet
func (projets *ProjetVille) GetWorkProjet(idProjet int) (int, error) {
	projets.projetsVilleMutex.RLock()
	defer projets.projetsVilleMutex.RUnlock()
	for _, proj := range projets.projetsVille {
		if proj.Id == idProjet {
			return proj.Batiment.Work, nil
		}
	}
	return 0, errors.New("Aucun projet avec cet identifiant")
}

// Retourne le travail accomplit sur un projet
func (projets *ProjetVille) GetWorkDoneProjet(idProjet int) (int, error) {
	projets.projetsVilleMutex.RLock()
	defer projets.projetsVilleMutex.RUnlock()
	for _, proj := range projets.projetsVille {
		if proj.Id == idProjet {
			return proj.Travail, nil
		}
	}
	return 0, errors.New("Aucun projet avec cet identifiant")
}

// Change la valeur d'un projet à l'index indiqué
func (projets *ProjetVille) Set(index int, proj Projet) {
	projets.projetsVilleMutex.Lock()
	defer projets.projetsVilleMutex.Unlock()
	projets.projetsVille[index] = proj
}

// Retourne toute la liste des projets
func (projets *ProjetVille) GetAll() []Projet {
	projets.projetsVilleMutex.RLock()
	defer projets.projetsVilleMutex.RUnlock()
	return projets.projetsVille
}

// Trouve et ajoute un travail au jobBoard pour un ouvrier, s'il n'y a pas de travail retourne un projet vide et une erreur
func (projets *ProjetVille) FindWork(idOuvrier int, jobBoard JobBoard) (Projet, error) {
	projets.projetsVilleMutex.Lock()
	defer projets.projetsVilleMutex.Unlock()

	if len(projets.projetsVille) > 0 {
		// La moitié du temps, on tente d'assigner le travail de façon aléatoire
		if rand.Float32() < 0.5 {
			pIndex := rand.Intn(len(projets.projetsVille))
			proj := projets.projetsVille[pIndex]

			if proj.Capacity < proj.Batiment.WorkerCapacity {
				dayWork := proj.Capacity * workUnitPerDay
				if proj.Travail+dayWork < proj.Batiment.Work {
					proj.Capacity++
					jobBoard.Set(idOuvrier, proj)
					return proj, nil
				}
			}
		}
		for _, proj := range projets.projetsVille {
			if proj.Capacity < proj.Batiment.WorkerCapacity {
				dayWork := proj.Capacity * workUnitPerDay
				if proj.Travail+dayWork < proj.Batiment.Work {
					proj.Capacity++
					jobBoard.Set(idOuvrier, proj)
					return proj, nil
				}
			}
		}
	}
	return Projet{}, errors.New("Pas de projet de disponible")
}
