package reader

func Reader(l int, inputCh chan int, outputCh chan int) {
	container := make([]int, 0, l)
	for {
		output := <-inputCh
		container = append(container, output)
		if len(container) == l {
			var mid float64
			for i := 0; i < l; i++ {
				mid += float64(container[i])
			}
			mid = mid/float64(l) + 0.5
			outputCh <- int(mid)
			container = make([]int, 0, l)
		}
	}
}
