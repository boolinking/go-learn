package main

import "fmt"

type Cat struct {
	name string
	scientificName string
	category string
}

func New(name , scientificName, category string) Cat {
	return Cat{
		name:name,
		scientificName:scientificName,
		category:category,
	}
}

func (cat *Cat) SetName(name string)  {
	cat.name = name
}

func (cat Cat) SetNameCopy(name string)  {
	cat.name = name
}

func (cat Cat) ScientificName() string {
	return cat.scientificName
}

func (cat Cat) Category() string {
	return cat.category
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.category, cat.name)
}
func main()  {
	cat := New("little pig", "American Shorthair", "cat")
	cat.SetName("dog")
	fmt.Printf("The cat: %s\n",cat)

	cat.SetNameCopy("little dog") //cat会被复制
	fmt.Printf("The cat: %s\n",cat)

	type Pet interface {
		SetName(name string)
		ScientificName() string
		Category() string
	}
	_, ok := interface{}(cat).(Pet)
	fmt.Printf("Cat implements interface Pet: %v\n", ok)
	_, ok = interface{}(&cat).(Pet)
	fmt.Printf("Cat implements interface Pet: %v\n", ok) //指针变量包含所有方法的集合

}
