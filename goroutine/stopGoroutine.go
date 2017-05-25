package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)     //休眠1秒钟
	atomic.StoreInt64(&shutdown, 1) //将shutdown变量值修改为1
	wg.Wait()
}

func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Printf("doing %s work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("stop %s work\n", name)
			break
		}
	}
}
