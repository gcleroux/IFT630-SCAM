// package test

// import (
// 	"testing"

// 	"github.com/gcleroux/IFT630-SCAM/pkg/people"
// 	"github.com/gcleroux/IFT630-SCAM/pkg/utils"
// )

// func TestCitoyen(t *testing.T) {
// 	config := utils.LoadConfig("./files/test_conf.yml")

// 	for i := 0; i < conf.NbCitoyen; i++ {
// 		go people.Population(i)
// 	}

// 	for {
// 		if people.NbVisites() == got.NbCitoyen {
// 			break;
// 		}
// 	}

// 	if got != expected {
// 		t.Errorf("Population() = %s; want %s", got, expected)
// 	}
// }