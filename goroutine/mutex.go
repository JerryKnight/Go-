package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int32
	wg      sync.WaitGroup
	mutex   sync.Mutex
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
		mutex.Lock() //加锁
		{
			value := counter
			runtime.Gosched() //goroutine从线程退出，并放入队列中
			value++
			counter = value
		}
		mutex.Unlock() //解锁
	}
}
