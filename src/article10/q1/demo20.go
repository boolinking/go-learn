package main

import "fmt"

func main()  {
	ch1 := make(chan int , 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	//ch1 <- 4
	elem1 := <- ch1
	fmt.Printf("The first element from received from channel ch1: %v\n",elem1)

}
