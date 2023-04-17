package batiment

import (
	"sync"
)

// Structure d'exclusion mutuelle read/write pour gérer les proejts de la ville qui sont partagés par différents threads.
// sync.RWMutex permet la lecture simultané de plusieurs processus ou un seul processus en écriture.
type ProjetVille struct {
	ProjetsVille      []Projet
	ProjetsVilleMutex sync.RWMutex
}

// Ajouter un nouveau projet
func (projets *ProjetVille) Append(nouveauProjet Projet) {
	projets.ProjetsVilleMutex.Lock()
	defer projets.ProjetsVilleMutex.Unlock()
	projets.ProjetsVille = append(projets.ProjetsVille, nouveauProjet)
}

// Supprimer un projet à l'index indiqué
func (projets *ProjetVille) Delete(index int) {
	projets.ProjetsVilleMutex.Lock()
	defer projets.ProjetsVilleMutex.Unlock()
	projets.ProjetsVille = append(projets.ProjetsVille[:index], projets.ProjetsVille[index+1:]...)
}

// Retourne le nombre de projets dans le tableau
func (projets *ProjetVille) Length() int {
	projets.ProjetsVilleMutex.RLock()
	defer projets.ProjetsVilleMutex.RUnlock()
	return len(projets.ProjetsVille)
}

// Retourne le projet à l'index du tableau
func (projets *ProjetVille) Get(index int) Projet {
	projets.ProjetsVilleMutex.RLock()
	defer projets.ProjetsVilleMutex.RUnlock()
	return projets.ProjetsVille[index]
}

// Change la valeur d'un projet à l'index indiqué
func (projets *ProjetVille) Set(index int, val Projet) {
	projets.ProjetsVilleMutex.Lock()
	defer projets.ProjetsVilleMutex.Unlock()
	projets.ProjetsVille[index] = val
}

// Retourne toute la liste des projets
func (projets *ProjetVille) GetAll() []Projet {
	projets.ProjetsVilleMutex.RLock()
	defer projets.ProjetsVilleMutex.RUnlock()
	return projets.ProjetsVille
}
