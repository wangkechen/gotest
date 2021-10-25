package main

import "fmt"

func divideByZero() {
	// Use this deferred function to handle errors.
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("HERE")
			fmt.Println(err)
			fmt.Println(0)
		}
	}()
	// Cause an error.
	// ... Go will run the defer func above.
	cat := 0
	dog := 10 / cat
	fmt.Println(dog)
}

/*func main() {
	// Create a divide by zero error and handle it.
	divideByZero()
	fmt.Println("reach")
}*/