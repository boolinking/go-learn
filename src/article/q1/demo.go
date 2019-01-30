package main

import (
	"fmt"
	"flag"
)

//实现命令源码文件接受参数
var name string

func init()  {
	flag.StringVar(&name,"name","everyone","The greet object")

}
func main()  {
	var ip = flag.Int("flagname", 1234, "help message for flagname")
	flag.Parse()
	fmt.Printf("Hello %s !",name)
	fmt.Println("ip has value ", *ip)
}
