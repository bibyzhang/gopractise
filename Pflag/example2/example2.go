package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

var (
	ip = pflag.IntP("flagname", "f", 1234, "help message")
)

func main() {
	pflag.Lookup("flagname").NoOptDefVal = "4321"

	// 把用户传递的命令行参数解析为对应变量的值
	pflag.Parse()

	fmt.Println("flagname=", *ip)
}
