package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	c  int
	m  sync.Mutex
	wg sync.WaitGroup
}

func (counter *Counter) incremetation(max int) {
	for i := 0; i < max; i++ {
		counter.wg.Add(1)
		go func() {
			defer counter.wg.Done()
			counter.m.Lock()
			counter.c++
			fmt.Println(counter.c)
			counter.m.Unlock()
		}()
	}
	counter.wg.Wait()
}

func main() {
	c1 := Counter{}
	c1.c = 0

	c1.incremetation(1000)
}
