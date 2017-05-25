//这个例子演示如何在程序中造成竞争状态
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	runtime.GOMAXPROCS(1) //分配一个逻辑器给调度器使用
	wg.Add(2)

	go intCounter(1)
	go intCounter(2)

	wg.Wait()
	fmt.Println("counter:", counter)
}

func intCounter(id int) {
	defer wg.Done() //函数执行完后通知WaitGroup
	for i := 0; i < 2; i++ {
		value := counter
		runtime.Gosched() //当前goroutine从线程退出，并放回队列
		value++
		counter = value
	}
}
