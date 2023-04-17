package batiment

import (
	"sync"
)

// JobBoard contient une map de tous les ouvrier et à quel projet ils sont assignés
// Utilise un sync.Map qui gère la concurrence entre les threads
type JobBoard struct {
	projetBoard sync.Map
}

// Retourne le projet associé à l'index
func (board *JobBoard) Get(index int) (Projet, bool) {
	if value, ok := board.projetBoard.Load(index); ok {
		proj := value.(Projet)
		return proj, true
	}
	return Projet{}, false
}

// Change la valeur de l'élément index au projet en paramètre
func (board *JobBoard) Set(index int, proj Projet) {
	board.projetBoard.Store(index, proj)
}

// Supprimer les tâches d'ouvrier associé à un projet
func (board *JobBoard) Delete(index int) {
	board.projetBoard.Range(func(key, value interface{}) bool {
		proj, ok := value.(Projet)
		if ok && proj.Id == index {
			board.projetBoard.Delete(key)
		}
		return true
	})
}
