package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	//只收不发通道
	var uselessChan = make(<-chan int, 1)
	//只发不收通道
	var anotherChan = make(chan<- int, 1)
	fmt.Printf("The useless channels: %v, %v\n",uselessChan,anotherChan)

	intChan := make(chan int , 1)
	SendInt(intChan)

	intChan2 := getIntChan()
	for elem := range intChan2 {
		fmt.Printf("the element in inChan2: %v\n",elem)
	}

	_ = GetIntChan(getIntChan)
	

}

func SendInt(ch chan<- int )  {
	ch <- rand.Intn(1000)
}

type Notifier interface {
	SendInt(ch chan<- int)
}

func getIntChan() <-chan int {
	num := 5
	ch := make(chan int , num)
	for i := 0; i < num; i++  {
		ch <- i
	}
	close(ch)
	return ch
}

type GetIntChan func() <-chan int