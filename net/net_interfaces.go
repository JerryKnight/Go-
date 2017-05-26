package main

import (
	"fmt"
	"net"
)

func main() {
	addrs, _ := net.InterfaceAddrs() //返回系统的网络接口的地址列表
	fmt.Println(addrs)

	interfaces, _ := net.Interfaces() //返回系统的网络接口
	fmt.Println(interfaces)

	hb := net.JoinHostPort("127.0.0.1", "8080") //将host和port合并为一个网络地址
	fmt.Println(hb)

	it, _ := net.LookupAddr("127.0.0.1") //映射到该地址的主机名序列
	fmt.Println(it)
}
