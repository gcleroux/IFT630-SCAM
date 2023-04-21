package batiment

import (
	"sync"
)

// JobBoard contient une map de tous les ouvrier et à quel projet ils sont assignés
// Utilise un sync.Map qui gère la concurrence entre les threads
type JobBoard struct {
	projetBoard sync.Map
}

// Retourne le projet associé à un ouvrier
func (board *JobBoard) GetProjet(idOuvrier int) (Projet, bool) {
	if value, ok := board.projetBoard.Load(idOuvrier); ok {
		proj := value.(Projet)
		return proj, true
	}
	return Projet{}, false
}

// Changer le projet sur lequel un ouvrier travail
func (board *JobBoard) Set(idOuvrier int, proj Projet) {
	board.projetBoard.Store(idOuvrier, proj)
}

// Supprimer toutes les tâches d'ouvrier associé à un projet
func (board *JobBoard) DeleteProject(idProjet int) {
	board.projetBoard.Range(func(key, value interface{}) bool {
		proj, ok := value.(Projet)
		if ok && proj.Id == idProjet {
			board.projetBoard.Delete(key)
		}
		return true
	})
}

// Supprimer une tâche associée à un ouvrier
func (board *JobBoard) DeleteOuvrier(idOuvrier int) {
	// Si l'ouvrier est associé à un projet
	if _, ok := board.projetBoard.Load(idOuvrier); ok {
		// Supprimer la tâche de l'ouvrier
		board.projetBoard.Delete(idOuvrier)
	}
}
