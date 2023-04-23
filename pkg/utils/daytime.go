package utils

import "time"

// Cette fonction est responsable de la duree d'une journee
// Les composantes dependantes seront alors notifiées par le close
// et quitteront leur step
func DayTime(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}
