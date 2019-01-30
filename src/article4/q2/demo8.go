package main

import (
	"flag"
	"fmt"
)

func main()  {
	var name  = getTheFlag() //类型推断利于代码重构
	flag.Parse()
	fmt.Printf("Hello %v!\n",*	name)
}



func getTheFlag() *string {

	return flag.String("name","everyone","the greeting object")
}
