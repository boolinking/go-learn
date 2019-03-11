package main

import "fmt"

func main(){
	array1 := [3]string{"a","b","c"}
	fmt.Printf("The array: %v\n",array1)
	//数组传给函数会被复制  数组是值类型
	array2 := modifyArray(array1)
	fmt.Printf("The array: %v\n",array1)
	fmt.Printf("The array: %v\n",array2)

	//引用类型（切片，字典，通道）传给函数发生浅复制。
	silce := []string{"1","2","3"}
	fmt.Printf("The silce: %v\n",silce)
	silce2 := modifySilce(silce)
	fmt.Printf("The silce: %v\n",silce)
	fmt.Printf("The silce: %v\n",silce2)

	complexArray1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}

	fmt.Printf("The complex array: %v\n", complexArray1)
	complexArray2 := modifyComplexArray(complexArray1)
	fmt.Printf("The modified complex array: %v\n", complexArray2)
	fmt.Printf("The original complex array: %v\n", complexArray1)
}


func modifyArray(a [3]string)  [3]string {
	a[1] = "a"
	return a
}

func modifySilce( s []string ) []string {
	s[1] = "1"
	return s
}

// 示例3。
func modifyComplexArray(a [3][]string) [3][]string {
	a[1][1] = "s"
	a[2] = []string{"o", "p", "q"}
	return a
}