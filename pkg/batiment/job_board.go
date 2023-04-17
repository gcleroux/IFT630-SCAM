package batiment

import (
	"sync"
)

// JobBoard contient un sync.Map qui gère la concurrence entre les threads
type JobBoard struct {
	data sync.Map
}

// Retourne le projet associé à l'index
func (board *JobBoard) Get(index int) (Projet, bool) {
	if value, ok := board.data.Load(index); ok {
		proj := value.(Projet)
		return proj, true
	}
	return Projet{}, false
}

// Change la valeur de l'élément index au projet en paramètre
func (board *JobBoard) Set(index int, proj Projet) {
	board.data.Store(index, proj)
}

// Supprimer un projet du board selon l'index indiqué
func (board *JobBoard) Delete(index int) {
	board.data.Range(func(key, value interface{}) bool {
		proj, ok := value.(Projet)
		if ok && proj.Id == index {
			board.data.Delete(key)
			return false
		}
		return true
	})
}
