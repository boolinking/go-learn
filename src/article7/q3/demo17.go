package main

import "fmt"

func main(){
	a1 := [7]int{1 , 2, 3, 4, 5, 6, 7}
	fmt.Printf("a1: %v (len: %d, cap: %d)\n",a1,len(a1),cap(a1))

	s9 := a1[1:4]
	//s9[0] = 1
	fmt.Printf("s9: %v (len: %d, cap: %d)\n",
		s9, len(s9), cap(s9))

	for i := 1; i <= 5; i++ {
		s9 = append(s9, i)		//append 追加函数时底层数组在需要扩容时不会改变，只会生成新数组
								//只要长度不超过源容量，append就不会引起扩容。原底层数组的元素会被替换掉
		fmt.Printf("s9(%d): %v (len: %d, cap: %d)\n",
			i, s9, len(s9), cap(s9))
	}
	fmt.Printf("a1: %v (len: %d, cap: %d)\n",
		a1, len(a1), cap(a1))
	fmt.Println()
}
