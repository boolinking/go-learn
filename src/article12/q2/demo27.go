package main

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

func calculate(x , y int , op operate) (int , error)    {
	if op == nil {
		return 0, errors.New("Invaild operation")

	}
	return op(x,y) , nil
}

type calculateFunc func(x , y int ) (int , error)

func genCalculator(op operate) calculateFunc  {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("Invaild operation")
		}
		return op(x,y) ,nil
	}
}

func main()  {
	x , y := 12 , 13
	op := func(x , y int) int {
		return x + y
	}

	result , err := calculate(x , y , op)
	fmt.Printf("The result: %d (error: %v)\n",result,err)
	result , err = calculate(x , y , nil)
	fmt.Printf("The result: %d (error: %v)\n",result,err)


	x , y = 12 , 16
	calc := genCalculator(op)
	result , err = calc(x , y)
	fmt.Printf("The result: %d (error: %v)\n",result,err)

}
