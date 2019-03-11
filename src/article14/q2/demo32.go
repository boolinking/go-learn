package main

import "fmt"

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
	dog := Dog{"little dog"}
	fmt.Printf("The dog name is %q\n",dog.Name())

	var pet Pet = dog // 复制
	dog.SetName("monster")
	fmt.Printf("The dog name is %q\n",dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
	fmt.Println()

	dog1 := Dog{"little dog"}
	fmt.Printf("The name of first dog is %q.\n",dog1.Name())
	dog2 := dog1
	fmt.Printf("The second of first dog is %q.\n",dog2.Name())

	dog1.name = "big dog"
	fmt.Printf("The name of first dog is %q.\n",dog1.Name())
	fmt.Printf("The second of first dog is %q.\n",dog2.Name())
	fmt.Println()

	dog = Dog{"little dog"}
	fmt.Printf("The dog name is %q\n",dog.Name())
	pet = &dog
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())

}



