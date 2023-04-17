package batiment

import (
	"sync"
)

// Structure d'exclusion mutuelle read/write pour gérer les batiments de la ville qui sont une partagés par différents threads.
// sync.RWMutex permet la lecture simultané de plusieurs processus ou un seul processus en écriture.
type BatimentVille struct {
	BatimentsVille      []Batiment
	BatimentsVilleMutex sync.RWMutex
}

// Ajouter un nouveau batiment
func (batiments *BatimentVille) Append(nouveauBatiment Batiment) {
	batiments.BatimentsVilleMutex.Lock()
	defer batiments.BatimentsVilleMutex.Unlock()
	batiments.BatimentsVille = append(batiments.BatimentsVille, nouveauBatiment)
}

// Retourne le nombre de batiments dans le tableau
func (batiments *BatimentVille) Length() int {
	batiments.BatimentsVilleMutex.RLock()
	defer batiments.BatimentsVilleMutex.RUnlock()
	return len(batiments.BatimentsVille)
}

// Retourne le batiment à l'index du tableau
func (batiments *BatimentVille) Get(index int) Batiment {
	batiments.BatimentsVilleMutex.RLock()
	defer batiments.BatimentsVilleMutex.RUnlock()
	return batiments.BatimentsVille[index]
}

// Retourne toute la liste des batiments
func (batiments *BatimentVille) GetAll() []Batiment {
	batimentsVille.BatimentsVilleMutex.RLock()
	defer batimentsVille.BatimentsVilleMutex.RUnlock()
	return batimentsVille.BatimentsVille
}
