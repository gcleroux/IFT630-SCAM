package test

import (
	"log"
	"testing"

	"github.com/gcleroux/IFT630-SCAM/pkg/utils"
)

func TestConfig(t *testing.T) {
	got, err := utils.LoadConfig("./files/test_conf.yml")

	if err != nil {
		log.Fatal(err)
	}

	expectedBudget := 420
	expectedNbOuvrier := 69

	if got.Budget != expectedBudget {
		t.Errorf("TestConfig() = %d; want %d", got.Budget, expectedBudget)
	}
	if got.NbOuvrier != expectedNbOuvrier {
		t.Errorf("TestConfig() = %d; want %d", got.NbOuvrier, expectedNbOuvrier)
	}
}
