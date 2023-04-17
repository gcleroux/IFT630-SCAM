package batiment

import (
	"errors"
	"sync"
)

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

// Retourne le batiment à l'index du tableau
func (batiments *BatimentVille) Get(index int) Batiment {
	batiments.batimentsVilleMutex.RLock()
	defer batiments.batimentsVilleMutex.RUnlock()
	return batiments.batimentsVille[index]
}

// Retourne toute la liste des batiments
func (batiments *BatimentVille) GetAll() []Batiment {
	batimentsVille.batimentsVilleMutex.RLock()
	defer batimentsVille.batimentsVilleMutex.RUnlock()
	return batimentsVille.batimentsVille
}

// Trouve un emploi à un citoyen dans un batiment de la ville
func (batiments *BatimentVille) Visite() (Batiment, error) {
	for _, batiment := range batiments.batimentsVille {
		if batiment.Visitors < batiment.Capacity {
			// batiment.Visitors++	// TODO: Est-ce que les citoyens visite le batiment ou il se trouve un emploi permanent ?
			return batiment, nil
		}
	}
	return Batiment{}, errors.New("Pas de batiment disponible")
}
