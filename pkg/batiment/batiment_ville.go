package batiment

import (
	"errors"
	"sync"
)

// BatimentVille contient la liste des batiments complétés de la ville
// Structure d'exclusion mutuelle read/write pour gérer les batiments de la ville qui sont une partagés par différents threads.
// sync.RWMutex permet la lecture simultané de plusieurs processus ou un seul processus en écriture.
type BatimentVille struct {
	batimentsVille      []Batiment
	batimentsVilleMutex sync.RWMutex
}

// Ajouter un nouveau batiment
func (batiments *BatimentVille) Append(nouveauBatiment Batiment) {
	batiments.batimentsVilleMutex.Lock()
	defer batiments.batimentsVilleMutex.Unlock()
	batiments.batimentsVille = append(batiments.batimentsVille, nouveauBatiment)
}

// Retourne le nombre de batiments dans le tableau
func (batiments *BatimentVille) Length() int {
	batiments.batimentsVilleMutex.RLock()
	defer batiments.batimentsVilleMutex.RUnlock()
	return len(batiments.batimentsVille)
}

// Retourne toute la liste des batiments
func (batiments *BatimentVille) GetAll() []Batiment {
	batiments.batimentsVilleMutex.RLock()
	defer batiments.batimentsVilleMutex.RUnlock()
	return batiments.batimentsVille
}

// Trouve un emploi à un citoyen dans un batiment de la ville
func (batiments *BatimentVille) Visite(id int) (Batiment, error) {
	batiments.batimentsVilleMutex.RLock()
	defer batiments.batimentsVilleMutex.RUnlock()
	// Rempli les bâtiments en ordre de construction
	j := 0
	for _, b := range batiments.batimentsVille {
		for i := 0; i < b.Capacity; i++ {
			if i+j == id {
				return b, nil
			}
		}
		j += b.Capacity
	}

	return Batiment{}, errors.New("pas de batiment disponible")
}

// Met le nombre de Visitors pour chaque batiments de la ville à 0
func (batiments *BatimentVille) ResetVisites() {
	batiments.batimentsVilleMutex.Lock()
	defer batiments.batimentsVilleMutex.Unlock()
	for _, batiment := range batiments.batimentsVille {
		batiment.Visitors = 0
	}
}

// Retourne la somme des Capacités pour tous les bâtiments de la ville
func (batiments *BatimentVille) CalculCapacitéEmploieVille() int {
	batiments.batimentsVilleMutex.Lock()
	defer batiments.batimentsVilleMutex.Unlock()
	capacitéEmploieVille := 0
	for _, b := range batiments.batimentsVille {
		capacitéEmploieVille += b.Capacity
	}

	return capacitéEmploieVille
}
