package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//等待组，用于等待程序结束
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//创建一个无缓冲的通道
	table := make(chan int)

	wg.Add(2)

	go playPingpong("Jerry", table)
	go playPingpong("Tom", table)

	//开球
	table <- 1

	wg.Wait()
}

func playPingpong(user string, table chan int) {
	defer wg.Done()

	for {

		//等待球被击打过来
		ball, ok := <-table

		if !ok {
			//通道关闭了表示对方没有击中
			fmt.Printf("球员[%s]赢了。。。。\n", user)
			break
		}

		//获得一个随机数，判断是否击中了球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("球员[%s]失误了。。。\n", user)
			//没有击中球，关闭通道
			close(table)
			break
		}

		time.Sleep(10 * time.Millisecond)

		//击中了球并把次数加1
		fmt.Printf("球员[%s]击中了球，第%d拍\n", user, ball)
		ball++
		table <- ball
	}

}
