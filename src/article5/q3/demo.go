package main

import . "fmt" // fmt包中的公开的程序实体 视为当前包的程序实体
import . "article5/q3/lib"

// var Name  int 编译错误

func main()  {
	var Name = ""
	Printf("name is %s .\n",Name)
}

func hello()  {
	Printf("name is %s .\n",Name)
}