package sensor

import (
	"crypto/rand"
	"math/big"
)

func Sensor(max int64, ch chan int) error {
	// r, err := rand.Int(rand.Reader, big.NewInt(max))
	// if err != nil {
	// 	return err
	// }
	// r := rand.New(rand.NewSource(seed))
	for {
		r, err := rand.Int(rand.Reader, big.NewInt(max))
		if err != nil {
			return err
		}
		s := r.Int64()
		ch <- int(s)
	}
}
