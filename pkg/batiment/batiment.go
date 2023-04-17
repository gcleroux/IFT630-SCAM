package batiment

// Ce fichier crée la struct pour encapsuler un batiment
// Si certains attributs des batiments sont modifiés, il est
// important d'ajuster ce fichier

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Batiment struct {
	Id             int    `yaml:"id"`           // Identifiant
	Name           string `yaml:"name"`         // Nom
	Price          int    `yaml:price`          // Coût de construction
	Work           int    `yaml:work`           // Travail requis pour construire
	Capacity       int    `yaml:capacity`       // Nombre de visiteur maximal
	WorkerCapacity int    `yaml:workercapacity` // Nombre d'ouvrier maximal pour construire
	Income         int    `yaml:income`         // Revenu
	Visitors       int    // Nombre de visiteur dans le batiment
}

type Projet struct {
	Id       int      // Identifiant
	Batiment Batiment // Batiment associé au projet
	Travail  int      // Travail accompli de la construction
	Capacity int      // Nombre d'ouvrier travaillant sur le chantier
}

type Travail struct {
	Id     int // Identifiant du projet de ce travail
	Effort int // Effort de travail fait par l'employer
}

// Load les infos des batiments a partir des fichiers YAML
func loadBatimentsInfos(prefix string) []Batiment {
	// Liste vides des batiments
	batiments := []Batiment{}

	// Get liste des fichiers YAML des batiments
	batimentFiles, err := ioutil.ReadDir(prefix)

	// Erreur directory non trouve
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, file := range batimentFiles {
		// Read the file's data
		data, err := os.ReadFile(prefix + file.Name())
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		var batiment Batiment
		err = yaml.Unmarshal(data, &batiment)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		batiments = append(batiments, batiment)
	}
	return batiments
}
