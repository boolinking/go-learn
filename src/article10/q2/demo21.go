package main

func main()  {
	ch1 := make(chan int , 	1)
	ch1 <- 1
	//ch1 <- 2  //通道阻塞

	ch2 := make(chan int , 1)
	//elem ,ok := <- ch2

	ch2 <- 1

	var ch3 chan int
	//ch3 <- 1 //通道的值为nil，因此这里会造成永久的阻塞！
	<- ch3 //这里会造成永久的阻塞


}
