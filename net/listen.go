package main

import (
	"fmt"
	"net"
)

func main() {

	//第一个参数必须是面向流的网络，tcp、tcp4、tcp6、unix、unixpacket
	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("监听失败")
		return
	}

	for {
		conn, err := ls.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn) //使用goroutine处理请求
	}
}

func handleConnection(conn Conn) {

}
