package main

import "fmt"

type Animal interface {
	Category() string
	ScientificName() string
}

type Named interface {
	Name() string
}

type Pet interface {
	Animal
	Named
}

type PetTag struct {
	name string
	ower string
}

func (pet PetTag) Name() string  {
	return pet.name
}

func (pet PetTag) Ower() string {
	return pet.ower
}

type Dog struct {
	PetTag
	scientificName string
}

func (dog Dog) ScientificName()  string{
	return dog.scientificName
}

func (dog Dog) Category() string {
	return "dog"
}

func main()  {
	petTag := PetTag{name:"little dog"}
	_,ok := interface{}(petTag).(Named)
	fmt.Printf("PetTag implements interface Named: %v\n", ok)
	dog := Dog {
		PetTag : petTag,
		scientificName:"Labrador ",
	}
	_, ok = interface{}(dog).(Animal)
	fmt.Printf("Dog implements interface Animal: %v\n", ok)
	_, ok = interface{}(dog).(Named)
	fmt.Printf("Dog implements interface Named: %v\n", ok)
	_, ok = interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
}
