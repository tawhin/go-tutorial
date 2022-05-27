package demo

import "fmt"

func RunLoopsDemo() {
	loopToCondition()
	loopForever()
	loopCollections()
}

func loopToCondition() {
	fmt.Println("\nShow loop to condition")

	// Variable 'i' is declared at the for loop level scope, it cannot be
	// referenced outside the loop.
	for i := 0; i < 5; i++ {
		fmt.Println(i)

		if i == 2 {
			continue
		}
		fmt.Println("Continuing")
		if i == 4 {
			break
		}
	}
}

func loopForever() {
	fmt.Println("\nShow infinity loop")
	var i int

	for {

		if i == 5 {
			break
		}

		fmt.Println(i)
		i++
	}
}

func loopCollections() {
	fmt.Println("\nShow loop collection")

	slice := []int{1, 2, 3}
	fmt.Println("Slice:", slice)
	for i, v := range slice {
		fmt.Println(i, v)
	}

	myMap := map[string]int{"http": 80, "https": 443}
	fmt.Println("Map:", myMap)
	for k, v := range myMap {
		fmt.Println(k, v)
	}

	// Go compiler has special consideration for looping of collections, so you can
	// just listen to the first returned param of the range function

	for k := range myMap {
		fmt.Println("Key:", k)
	}

	// If you just want values, then you have to express a write only indicator for the keys,
	// you cannot declare a explicit variable for the key and not use it (go flags a var 'declared but not used' compilation error)
	for _, v := range myMap {
		fmt.Println("Value:", v)
	}
}
