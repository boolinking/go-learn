package main

import (
	"fmt"
	"reflect"
)

type Pet interface {
	Name() string
	Category() string

}

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string)  {
	dog.name = name
}

func (dog Dog) Name() string  {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main()  {
	var dog *Dog
	fmt.Println("The first is nil")
	dog2 := dog
	fmt.Println("The second is nil")
	var pet Pet = dog2 //接口数据结构包含类型和数值
	if pet == nil{
		fmt.Println("The pet is nil")
	} else {
		fmt.Println("The pet is not nil")
	}
	fmt.Printf("The type of pet is %T.\n",pet)
	fmt.Printf("The type of pet is %s.\n",reflect.TypeOf(pet).String())
	fmt.Printf("The type of second dog is %T.\n",dog2)

	wrap := func(dog *Dog) Pet {
		if dog == nil {
			return nil
		}
		return dog
	}

	//dog2 == nil
	pet = wrap(dog2)
	if pet == nil{
		fmt.Println("The pet is nil")
	} else {
		fmt.Println("The pet is not nil")
	}

}




