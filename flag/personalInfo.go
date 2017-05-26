package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("name", "", "输入姓名")
	age := flag.Int("age", 18, "输入年龄")
	var married bool
	flag.BoolVar(&married, "married", false, "是否结婚")
	flag.Parse()
	fmt.Printf("姓名：%s, 年龄：%d, 是否结婚：%t", *username, *age, married)
}
