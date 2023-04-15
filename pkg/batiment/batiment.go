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
	Id       int    `yaml:"id"`
	Name     string `yaml:"name"`
	Price    int    `yaml:price`
	Work     int    `yaml:work`
	Capacity int    `yaml:capacity`
	Income   int    `yaml:income`
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
