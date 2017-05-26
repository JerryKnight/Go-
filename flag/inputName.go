package main

import (
	"flag"
	"fmt"
)

func main() {
	//第一个参数：参数名称，第二个参数：默认值，第三个参数：参数说明
	username := flag.String("name", "", "请输入你的名字")
	flag.Parse()
	fmt.Println("你的名字：", *username)
}
