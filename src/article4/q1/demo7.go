package q1

import (
	"flag"
	"fmt"
)
func main()  {
	//var name string
	//flag.StringVar(&name,"name","everyone","the greeting object")
	//var name = *flag.String("name","everyone","the greeting object") //类型推断
	name := *flag.String("name","everyone","the greeting object") //短变量声明

	flag.Parse()
	fmt.Printf("Hello %v!\n",name)

}
