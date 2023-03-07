package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type batimentInfo struct {
	idBatiment     int
	nomBatiment    string
	prixBatiment   int
	effortBatiment int
}

// --- Global Variable ---
var budget int = 2000
var nbOuvrier int = 3
var choixBatiments = []batimentInfo{
	{idBatiment: 1, nomBatiment: "Parc", prixBatiment: 250, effortBatiment: 20},
	{idBatiment: 2, nomBatiment: "Hopital", prixBatiment: 500, effortBatiment: 50},
	{idBatiment: 3, nomBatiment: "Hotel", prixBatiment: 400, effortBatiment: 40}}
var villeContenu = []string{}

// --- Main ---
func main() {
	start := time.Now()

	fmt.Println("Choix des batiments: ", choixBatiments)

	budget = MayorStart(budget, nbOuvrier, choixBatiments)

	MayorEnd(budget)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Print("Temps total d'exécution du programme:")
	fmt.Println(elapsed)
}

// --- Fonctions de la City ---

// --- Fonctions du Mayor ---
func MayorHello() string {
	return "Hello, Mayor!"
}

func MayorStart(budget int, nbOuvrier int, choixBatiments []batimentInfo) int {
	plusPetitPrix := 0
	go func() {
		plusPetitPrix = TrouveBatimentMoinsCher()
	}()

	for budget > plusPetitPrix {
		batimentChoisi := rand.Intn(len(choixBatiments))
		if budget > choixBatiments[batimentChoisi].prixBatiment {
			commande := choixBatiments[batimentChoisi]
			budget = budget - commande.prixBatiment
			OuvrierConstruit(commande)
		}
	}
	return budget
}

func MayorEnd(budget int) {
	fmt.Println("Le Mayor prend sa retraite avec un budget restant de " + strconv.Itoa(budget) + "$")
	fmt.Println("La ville contient les bâtiments suivants:")
	for i := 0; i < len(villeContenu); i++ {
		fmt.Println(villeContenu[i])
	}
}

// --- Fonctions des Ouvriers ---
func OuvrierConstruit(commande batimentInfo) {
	var effortTotal = commande.effortBatiment
	fmt.Println("L'ouvrier commence la construction de : ", commande.nomBatiment)
	for i := 0; i < effortTotal; i++ {
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Println("L'ouvrier a terminer la construction de : ", commande.nomBatiment)
	villeContenu = append(villeContenu, commande.nomBatiment)
}

//--- Fonctions des Citoyens ---

// --- Fonctions Globales ---
func TrouveBatimentMoinsCher() int {
	plusPetitPrix := math.MaxInt
	for i := 0; i < len(choixBatiments); i++ {
		if choixBatiments[i].prixBatiment < plusPetitPrix {
			plusPetitPrix = choixBatiments[i].prixBatiment
		}
	}
	return plusPetitPrix
}
