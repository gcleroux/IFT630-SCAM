package batiment

import (
	"sync"
)

// JobBoard contient une map de tous les ouvrier et à quel projet ils sont assignés
// Structure d'exclusion mutuelle read/write pour gérer les projets de la ville qui sont partagés par différents threads.
// sync.RWMutex permet la lecture simultané de plusieurs processus ou un seul processus en écriture.
type JobBoard struct {
	projetBoard      map[int]Projet
	projetBoardMutex sync.RWMutex
}

// Retourne le projet associé à l'index
func (board *JobBoard) Get(idOuvrier int) (Projet, bool) {
	board.projetBoardMutex.RLock()
	defer board.projetBoardMutex.RUnlock()
	if proj, ok := board.projetBoard[idOuvrier]; ok {
		return proj, true
	}
	return Projet{}, false
}

// Change la valeur de l'élément index au projet en paramètre
func (board *JobBoard) Set(idOuvrier int, proj Projet) {
	board.projetBoardMutex.Lock()
	defer board.projetBoardMutex.Unlock()
	board.projetBoard[idOuvrier] = proj
}

// Supprimer la tâche associé à un projet
func (board *JobBoard) Delete(idOuvrier int) {
	board.projetBoardMutex.Lock()
	defer board.projetBoardMutex.Unlock()
	delete(board.projetBoard, idOuvrier)
}

// Supprimer toutes les tâches associées à un projet
func (board *JobBoard) DeleteProjet(idProjet int) {
	board.projetBoardMutex.Lock()
	defer board.projetBoardMutex.Unlock()
	for key := range board.projetBoard {
		if board.projetBoard[key].Id == idProjet {
			delete(board.projetBoard, key)
		}
	}
}
