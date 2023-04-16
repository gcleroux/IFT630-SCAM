package batiment

import (
	"sync"
)

type ProjetVille struct {
	ProjetsVille      []Projet
	ProjetsVilleMutex sync.RWMutex
}

func (projets *ProjetVille) Append(nouveauProjet Projet) {
	projets.ProjetsVilleMutex.Lock()
	defer projets.ProjetsVilleMutex.Unlock()
	projets.ProjetsVille = append(projets.ProjetsVille, nouveauProjet)
}

func (projets *ProjetVille) Delete(index int) {
	projets.ProjetsVilleMutex.Lock()
	defer projets.ProjetsVilleMutex.Unlock()
	projets.ProjetsVille = append(projets.ProjetsVille[:index], projets.ProjetsVille[index+1:]...)
}

func (projets *ProjetVille) Length() int {
	projets.ProjetsVilleMutex.RLock()
	defer projets.ProjetsVilleMutex.RUnlock()
	return len(projets.ProjetsVille)
}

func (projets *ProjetVille) Get(index int) Projet {
	projets.ProjetsVilleMutex.RLock()
	defer projets.ProjetsVilleMutex.RUnlock()
	return projets.ProjetsVille[index]
}

func (projets *ProjetVille) Set(index int, val Projet) {
	projets.ProjetsVilleMutex.Lock()
	defer projets.ProjetsVilleMutex.Unlock()
	projets.ProjetsVille[index] = val
}

func (projets *ProjetVille) GetAll() []Projet {
	projets.ProjetsVilleMutex.RLock()
	defer projets.ProjetsVilleMutex.RUnlock()
	return projets.ProjetsVille
}
