package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("开始goroutine")
	runtime.GOMAXPROCS(1) //分配一个逻辑处理器给调度器使用
	//创建两个等待WaitGroup
	var wg sync.WaitGroup
	wg.Add(2)

	//创建一个协程打印字母表三次
	go func() {
		//执行完后通知WaitGroup这个协程执行完了
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	//创建一个协程打印小写字母表三次
	go func() {
		//执行完后通知WaitGroup这个协程执行完了
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("等待协程执行完")
	wg.Wait()
	fmt.Println("执行完")
}
