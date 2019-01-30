package main

import (
	"fmt"
	"flag"
	"os"
)

var name string

func init()  {
	flag.StringVar(&name,"name","everyone","The greet object")
}

func main()  {
	//自定义参数说明
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,"Usage of %s:\n","question")
		flag.PrintDefaults()
	}

	flag.Parse()
	fmt.Printf("Hello %s !",name)
}

