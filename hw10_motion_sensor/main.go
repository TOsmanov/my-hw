package main

import (
	"sync"
	"time"

	"github.com/TOsmanov/my-hw/hw10_motion_sensor/printer"
	"github.com/TOsmanov/my-hw/hw10_motion_sensor/reader"
	"github.com/TOsmanov/my-hw/hw10_motion_sensor/sensor"
)

func main() {
	wg := sync.WaitGroup{}
	sens := make(chan int)
	data := make(chan int)

	go sensor.Sensor(1000, sens)
	go reader.Reader(10, sens, data)
	go printer.Printer(data)

	time.Sleep(60 * time.Second)
	wg.Wait()
}
