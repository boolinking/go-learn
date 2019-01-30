package main

import (
	"flag"
)

//实现命令源码文件接受参数
var name string

func init()  {
	flag.StringVar(&name,"name","everyone","The greet object")

}
func main()  {
	flag.Parse()
	hello(name)
}


