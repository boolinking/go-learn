package main

import "fmt"

var container   = []string{"zero","one","two"}

func main()  {
	container := map[int]string{0 : "zero",1 : "ONE", 2 : "two"} //不同代码块可重名变量，发生屏蔽
	fmt.Printf(" The element is %s .\n",container[1])

}
