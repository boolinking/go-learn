package main

import (
	"io"
	"os"
	"fmt"
)

func main()  {
	var err error
	n ,err := io.WriteString(os.Stdout,"Hello everyone!\n	") //这里对err进行了重声明
	if err != nil {
		fmt.Printf("Error, %v\n",err)
	}
	fmt.Printf("%d byte(s) were written .", n)

}
