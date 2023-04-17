package batiment

import (
	"errors"
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

// Retourne le projet à l'index du tableau
func (projets *ProjetVille) Get(index int) Projet {
	projets.projetsVilleMutex.RLock()
	defer projets.projetsVilleMutex.RUnlock()
	return projets.projetsVille[index]
}

// Change la valeur d'un projet à l'index indiqué
func (projets *ProjetVille) Set(index int, val Projet) {
	projets.projetsVilleMutex.Lock()
	defer projets.projetsVilleMutex.Unlock()
	projets.projetsVille[index] = val
}

// Retourne toute la liste des projets
func (projets *ProjetVille) GetAll() []Projet {
	projets.projetsVilleMutex.RLock()
	defer projets.projetsVilleMutex.RUnlock()
	return projets.projetsVille
}

// Trouve et ajoute un travail au jobBoard pour un ouvrier, s'il n'y a pas de travail retourne un projet vide et une erreur
func (projets *ProjetVille) FindWork(idOuvrier int, jobBoard JobBoard) (Projet, error) {
	for _, proj := range projets.projetsVille {
		if proj.Capacity < proj.Batiment.WorkerCapacity {
			proj.Capacity++
			jobBoard.Set(idOuvrier, proj)
			return proj, nil
		}
	}
	return Projet{}, errors.New("Pas de projet disponible")
}
