package main

import "fmt"

func main(){
	ch1 := make(chan int , 2)

	go func(){
		for i := 0; i < 10; i++ {
			fmt.Printf("Sender sending element %v...\n",i)
			ch1 <- i
		}
		fmt.Println("Sender: close the channel. ")
		close(ch1)
	}()

	for  {
		elem ,ok := <- ch1
		if !ok {
			fmt.Println("Receiver: close channel. ")
			break
		}
		fmt.Printf("Receiver: receive a element %v\n",elem)
	}
}
