package main

func main()  {

	//go字典的键类型不能是函数、切片、字典
	//键类型必须支持判等操作

	//var badMap map[[]int]int //编译错误
	//_ = badMap
	//var badMap2 = map[interface{}]int{
	//	"1": 2,
	//	[]int{2}: 3,  //运行时panic
	//	3 : 4,
	//}
	//_ = badMap2
	var badMap3 map[[1][]string]int // 这里会引发编译错误。
	_ = badMap3

	//type badkey struct {
	//	silce []string
	//}

	//var badMap4  map[badkey]int //编译错误
	//_ = badMap4

}
