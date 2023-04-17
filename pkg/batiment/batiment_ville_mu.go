package batiment

import (
	"sync"
)

type BatimentVille struct {
	BatimentsVille      []Batiment
	BatimentsVilleMutex sync.RWMutex
}

func (batiments *BatimentVille) Append(nouveauBatiment Batiment) {
	batiments.BatimentsVilleMutex.Lock()
	defer batiments.BatimentsVilleMutex.Unlock()
	batiments.BatimentsVille = append(batiments.BatimentsVille, nouveauBatiment)
}

func (batiments *BatimentVille) Length() int {
	batiments.BatimentsVilleMutex.RLock()
	defer batiments.BatimentsVilleMutex.RUnlock()
	return len(batiments.BatimentsVille)
}

func (batiments *BatimentVille) GetIndex(index int) Batiment {
	batiments.BatimentsVilleMutex.RLock()
	defer batiments.BatimentsVilleMutex.RUnlock()
	return batiments.BatimentsVille[index]
}

func GetBatimentAll() []Batiment {
	batimentsVille.BatimentsVilleMutex.RLock()
	defer batimentsVille.BatimentsVilleMutex.RUnlock()
	return batimentsVille.BatimentsVille
}
