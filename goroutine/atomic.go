package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)

	go intCounter(1)
	go intCounter(2)

	wg.Wait()
	fmt.Println("counter:", counter)
}

func intCounter(id int) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}
