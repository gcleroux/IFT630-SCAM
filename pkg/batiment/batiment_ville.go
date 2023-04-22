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
func (batiments *BatimentVille) Visite() (Batiment, error) {
	batiments.batimentsVilleMutex.Lock()
	defer batiments.batimentsVilleMutex.Unlock()
	// Trouver un batiment qui peut accueilir un citoyen et l'ajouter
	for index, batiment := range batiments.batimentsVille {
		if batiment.Visitors < batiment.Capacity {
			batiments.batimentsVille[index].Visitors += 1
			return batiment, nil
		}
	}
	return Batiment{}, errors.New("Pas de batiment disponible")
}

func (batiments *BatimentVille) ResetVisites() {
	batiments.batimentsVilleMutex.Lock()
	defer batiments.batimentsVilleMutex.Unlock()
	// fmt.Println("Nombre de visiteur dans les batiments")
	for index := range batiments.batimentsVille {
		batiments.batimentsVille[index].Visitors = 0
	}
}
