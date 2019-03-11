package main

import "fmt"

type Printer func(contens string) (n int , err error )

func printToStd(contents string) (bytesNum int , err error)  {
	return fmt.Println(contents)
}

func main(){
	var p  Printer
	p = printToStd
	p("something")
}
