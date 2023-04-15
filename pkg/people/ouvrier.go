package people

//
// import (
// 	"fmt"
// 	"time"
//
// 	"github.com/gcleroux/IFT630-SCAM/pkg/batiment"
// )
//
// func Ouvrier(projets <-chan batiment.BatimentInfo, complets chan<- string) {
//
// 	for commande := range projets {
//
// 		var effortTotal = commande.EffortBatiment
// 		fmt.Println("Un ouvrier commence la construction de : ", commande.NomBatiment)
// 		for i := 0; i < effortTotal; i++ {
// 			time.Sleep(time.Millisecond * 50)
// 		}
// 		batiment.VilleContenu = append(batiment.VilleContenu, commande.NomBatiment)
//
// 		msg := fmt.Sprint("Un ouvrier a terminer la construction de : ", commande.NomBatiment)
// 		complets <- msg
// 	}
// }
