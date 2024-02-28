package sensor

import (
	"sync"
	"testing"
)

func TestSensor(t *testing.T) {
	wg := sync.WaitGroup{}
	sens := make(chan int)
	max := 100
	go Sensor(int64(max), sens)
	output := <-sens
	for i := 0; i < 10; i++ {
		if output > max {
			t.Error(i)
		}
	}
	wg.Wait()
}
