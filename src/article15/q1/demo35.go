package main

type Named interface {
	Name() string
}

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string)  {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main()  {

	//不可寻址的规律 不可变，临时结果，不安全
	const num = 222
	//_ = &num  //常量不可寻址
	//_ = &(123) //基本类型字面量不可寻址

	var str = "abc"
	_ = str
	//_ = &(str[0]) //字符串变量的索引结果不可寻址
	//_ = &(str[0:2]) //字符串变量的切片结果不可寻址
	str2 := str[0]
	_ = &str2

	//_ = &(234+123) //算术操作的结果值不可寻址。
	num2 := 234
	_ = num2
	//_ = &(num2 + num) //算术操作的结果值不可寻址。
	//_ = &([3]int{1,2,3}[0]) //对数组字面量的索引结果值不可寻址。
	//_ = &([3]int{1, 2, 3}[0:1]) //对数组字面量的切片结果不可寻址
	_ = &([]int{1,2,3}[0]) //切片字面量索引结果值可寻址
	//_ = &([]int{1,2,3}[0:1]) //不可寻址
	//_ = &(map[int]string{1: "a"}[0])

	var map1 = map[int]string{1:"a",2:"b",3:"c"}
	_ = map1
	//_ = &(map1[0])
	//_ = &(func(x, y int) int {
	//	return x + y
	//}) // 字面量代表的函数不可寻址。
	//_ = &(fmt.Sprintf) // 标识符代表的函数不可寻址。
	//_ = &(fmt.Sprintln("abc")) // 对函数的调用结果值不可寻址。

	dog := Dog{"little pig"}
	_ = dog
	//_ = &(dog.Name) // 标识符代表的函数不可寻址。
	//_ = &(dog.Name()) // 对方法的调用结果值不可寻址。

	//_ = &(Dog{"little pig"}.name) // 结构体字面量的字段不可寻址。

	//_ = &(interface{}(dog)) // 类型转换表达式的结果值不可寻址。
	dogI := interface{}(dog)
	_ = dogI
	//_ = &(dogI.(Named)) // 类型断言表达式的结果值不可寻址。
	named := dogI.(Named)
	_ = named
	//_ = &(named.(Dog)) // 类型断言表达式的结果值不可寻址。

	var chan1 = make(chan int, 1)
	chan1 <- 1
	//_ = &(<-chan1) // 接收表达式的结果值不可寻址。

}
