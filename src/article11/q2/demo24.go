package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main()  {
	for  i := 0; i < 10; i++ {
		example()

	}
	example2()
}

func example(){
	intChans := [3]chan int {
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	index := rand.Intn(3)
	intChans[index] <- index

	select {
	case <-intChans[0]:
		fmt.Println("The first candidate case is selected. ")
	case <-intChans[1]:
		fmt.Println("The second candidate case is selected. ")
	case elem := <-intChans[2]:
		fmt.Printf("The last candidate case is selected. the element is %v\n",elem)
	default:
		fmt.Println("No candidate case is selected. ")
	}
}

func example2()  {
	intChan := make(chan int, 1)
	//一秒后关闭通道
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})
	select {
	case _, ok := <- intChan:
		if !ok {
			fmt.Println("The channel close.")
			break
		}
		fmt.Println("The channel selected.")

	}
}
