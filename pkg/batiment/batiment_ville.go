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
	listeBatiments := batiments.batimentsVille
	return listeBatiments
}

// Retourne le nombre de batiment de chaque catégorie
func (batiments *BatimentVille) GetBatimentsList() map[string]int {
	batiments.batimentsVilleMutex.RLock()
	defer batiments.batimentsVilleMutex.RUnlock()
	// Compter le nombre de batiment de chaque type
	batimentMap := make(map[string]int)
	for _, b := range batiments.batimentsVille {
		if batimentMap[b.Name] == 0 {
			batimentMap[b.Name] = 1
		} else {
			batimentMap[b.Name]++
		}
	}
	return batimentMap
}

// Retourne un map des visites par type de batiment map(key: <NomBatiment>, value: <NombreVisiteur>)
func (batiments *BatimentVille) GetListeVisites() map[string]int {
	batiments.batimentsVilleMutex.RLock()
	defer batiments.batimentsVilleMutex.RUnlock()
	// Compter le nombre de visiteurs dans chaque batiment
	visiteMap := make(map[string]int)
	for _, b := range batiments.batimentsVille {
		// fmt.Println("[DEBUG] Ajout de visiteurs = ", batiments.batimentsVille[index].Visitors)
		visiteMap[b.Name] += b.Visitors
	}
	return visiteMap
}

// Trouve un emploi à un citoyen dans un batiment de la ville
func (batiments *BatimentVille) Visite() (Batiment, error) {
	batiments.batimentsVilleMutex.Lock()
	defer batiments.batimentsVilleMutex.Unlock()
	// Les citoyens visitent les batiments dans l'ordrede construction
	for index, batiment := range batiments.batimentsVille {
		if batiment.Visitors < batiment.Capacity {
			batiments.batimentsVille[index].Visitors += 1
			// fmt.Println("[DEBUG] ", batiment.Name, " visitors = ", batiments.batimentsVille[index].Visitors)
			return batiment, nil
		}
	}
	return Batiment{}, errors.New("pas de batiment disponible")
}

// Met le nombre de Visitors pour chaque batiments de la ville à 0
func (batiments *BatimentVille) ResetVisites() {
	batiments.batimentsVilleMutex.Lock()
	defer batiments.batimentsVilleMutex.Unlock()
	// Remettre le compte des visiteurs des batiments à 0
	for index := range batiments.batimentsVille {
		batiments.batimentsVille[index].Visitors = 0
	}
}

// Retourne la somme des Capacités pour tous les bâtiments de la ville
func (batiments *BatimentVille) CalculCapacitéEmploieVille() int {
	batiments.batimentsVilleMutex.RLock()
	defer batiments.batimentsVilleMutex.RUnlock()
	// Faire la somme de la capacité d'emploi dans la ville
	capacitéEmploieVille := 0
	for _, b := range batiments.batimentsVille {
		capacitéEmploieVille += b.Capacity
	}
	return capacitéEmploieVille
}
