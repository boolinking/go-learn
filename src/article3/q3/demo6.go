package main

import (
	"flag"
	"article3/q3/lib"
	// in "article3/q3/lib/internal" 无法通过编译
	//"os"
)

var name string

func init()  {
	flag.StringVar(&name , "name" ,"everyone","the greet object")
}

func main()  {
	flag.Parse()
	lib.Hello(name)
	//in.Hello(os.Stdout ,name)
}
