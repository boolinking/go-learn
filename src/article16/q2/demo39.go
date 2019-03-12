package main

import (
	"fmt"
	"time"
)

func main()  {
	num := 10
	sign := make(chan struct{},num)
	for i := 0; i < num ; i++ {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}
	for i := 0 ; i < num ; i++ {
		<- sign
	}
	//time.Sleep(time.Millisecond * 500)
}
