package main

import "fmt"

var block = "package"

func main()  {
	block := "function"

	{
		block := "inner" //不同的代码块可以对重名变量进行声明，Java不可
		fmt.Printf("the block is %s .\n",block)
	}

	fmt.Printf("the block is %s .\n",block )
}
