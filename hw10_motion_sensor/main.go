package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func Sensor(max int64, delay time.Duration) chan int {
	start := time.Now()
	stop := start.Add(delay)
	c := make(chan int)
	go func() {
		for {
			r, err := rand.Int(rand.Reader, big.NewInt(max))
			if err != nil {
				break
			}
			c <- int(r.Int64())
			if time.Now().After(stop) {
				break
			}
		}
		close(c)
	}()
	return c
}

func Reader(l int, inputCh chan int) chan int {
	c := make(chan int)
	container := make([]int, 0, l)
	go func() {
		for {
			a, ok := <-inputCh
			if !ok {
				break
			}
			container = append(container, a)
			if len(container) == l {
				var mid float64
				for i := 0; i < l; i++ {
					mid += float64(container[i])
				}
				mid = mid/float64(l) + 0.5
				c <- int(mid)
				container = make([]int, 0, l)
			}
		}
		close(c)
	}()
	return c
}

func main() {
	emulationSensor := Sensor(1000, time.Minute)
	readData := Reader(10, emulationSensor)

wait:
	for {
		select {
		case _, ok := <-emulationSensor:
			if !ok {
				fmt.Println("The processing of the motion sensor data is completed.")
				break wait
			}
		case output, ok := <-readData:
			if !ok {
				break wait
			}
			fmt.Println(output)
		}
	}
}
