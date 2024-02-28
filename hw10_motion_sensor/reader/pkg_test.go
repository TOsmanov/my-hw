package reader

import (
	"sync"
	"testing"

	"github.com/TOsmanov/my-hw/hw10_motion_sensor/sensor"
)

func TestReader(t *testing.T) {
	sens := make(chan int)
	data := make(chan int)
	max := 100
	count := 10000
	wg := sync.WaitGroup{}
	go sensor.Sensor(int64(max), sens)
	go Reader(count, sens, data)
	output := <-data

	if output > max {
		t.Error(output)
	}

	wg.Wait()
}
