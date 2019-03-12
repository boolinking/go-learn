package main

import (
	"unsafe"
	"fmt"
)

type Dog struct {
	name string
}
func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main()  {
	dog := Dog{"little dog"}
	p := &dog
	dogPtr := uintptr(unsafe.Pointer(p))
	namePtr := dogPtr + unsafe.Offsetof(p.name)
	nameP := (*string)(unsafe.Pointer(namePtr))
	fmt.Printf("nameP == &(p.name)? %v\n",
		nameP == &(p.name))
	fmt.Printf("The name of dog is %q.\n", *nameP)

	*nameP = "dog"
	fmt.Printf("The name of dog is %q.\n", p.name)

	numP := (*int)(unsafe.Pointer(namePtr))
	num := *numP
	fmt.Printf("This is an unexpected number: %d\n", num)

}
