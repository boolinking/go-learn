package main

import (
	"fmt"
	"flag"
	"os"
)

var name string

func init()  {
	//自定义参数说明
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name,"name","everyone","The greet object")
}

func main()  {
	flag.Parse()
	fmt.Printf("Hello %s !",name)
}
