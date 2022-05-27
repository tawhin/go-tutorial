package demo

import "fmt"

// Function names starting with a Capital are exported from the package
func RunTypesDemo() {
	showVars()
	showPointers()
	showConsts()
	showIotaConsts()
	showArrays()
	showSliceExcplicitArray()
	showSliceImplicitArray()
}

func showVars() {
	fmt.Println("\nShow Vars:")
	// Uninitialised, explicitly typed declaration
	var i int
	i = 2
	fmt.Println(i)

	// Initialised explicity typed.
	var j int = 3
	fmt.Println(j)
	var f float32 = 3.14
	fmt.Println(f)

	// Implicitly typed via initialisation
	name := "trevor"
	age := 47
	bmi := 21.23
	male := true
	fmt.Println(name, age, bmi, male)
}
func showPointers() {
	fmt.Println("\nShow Pointers:")
	var strPtr *string = new(string)
	// Use de-reference symbol * for assignment.
	*strPtr = "First"
	// Print address and value
	fmt.Println(strPtr, *strPtr)
	// Cannot use pointer arythematic in go, i.e. strPtr++

	name := "Second"
	// Use address symbol to implicitly declare and assign a pointer.

	namePtr := &name
	fmt.Println(namePtr, *namePtr)
}

func showConsts() {
	fmt.Println("\nShow Consts")
	// Const need to be evaluated at compile time
	// Implicitly defined const
	const c = 4
	fmt.Println(1 << c)
	// Go we dynamicly type the const c to evaluate the expression
	fmt.Println(c + 1.2)

	// Explicitly defined const with a specified type, this can only be evaluated
	// within expressions with other int type values.
	const d int = 3
	fmt.Println(d)
	// Need explicit conversion operator
	fmt.Println(float32(d) + 1.2)
}

const (
	// iota increments it's value by one everytime it's called.
	zero = iota
	one  = iota
	// can be used in expressions
	six   = iota + 4
	eight = 1 << iota // 1 shifted 3 times
	// repeat the last const expression, this increments iota and performs the same bitshift operation
	sixteen // 1 shifted 4 times
)

const (
	// iota resets for each const block
	// iota increments it's value by one everytime it's called.
	zeroAgain = iota
	oneAgain  = iota
	// can be used in expressions
	sixAgain   = iota + 4
	eightAgain = 1 << iota // 1 shifted 3 times
	// repeat the last const expression, this increments iota and performs the same bitshift operation
	sixteenAgain // 1 shifted 4 times
)

func showIotaConsts() {
	fmt.Println("\nShow Iota and Consts")
	// Note these consts are still evaluated at compile time
	fmt.Println(zero, one, six, eight, sixteen)
	fmt.Println(zeroAgain, oneAgain, sixAgain, eightAgain, sixteenAgain)
}

func showArrays() {
	fmt.Println("\nShow Array")
	// Fixed size collection of similar types.
	// Long form, elements initialsed to zero
	var arr [3]int
	arr[0] = 1
	fmt.Println(arr)

	// Implicit initialisation format
	arr2 := [3]int{1}
	// arr is equal to arr2
	fmt.Print(arr2)
}

func showSliceExcplicitArray() {
	fmt.Println("\nShow Slice Explicit Array")
	// Built on top of array, but dynamically sixed
	arr := [3]int{1}
	// Contains all the elements of arr
	slice := arr[:]
	fmt.Print(arr, slice)

	// slice is a effectively a pointer to the arr
	arr[1] = 43
	slice[2] = 47
	// arr and slice still the same.
	fmt.Print(arr, slice)
}

func showSliceImplicitArray() {
	fmt.Println("\nShow Slice Implicit Array")
	// Compiler declares and manages underlying array
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	// Dynamic, go will perform an implicit copy of the underlying into a new resized array
	slice = append(slice, 4, 43, 5)
	fmt.Println(slice)

	// Slice of slice, starting from index 1
	s2 := slice[1:]
	// upto but not including 2
	s3 := slice[:2]
	// from 1 upto but not including 3
	s4 := slice[1:3]
	fmt.Println(s2, s3, s4)
}

func showMap() {
	fmt.Println("\nShow Map")
	// collection of string keys to value int
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])
	// can be modified
	m["foo"] = 47
	fmt.Println(m["foo"])
	// Remove key
	delete(m, "foo")
	fmt.Println(m)

	// Insert key
	m["foo"] = 43
	fmt.Println(m)
}

func showStruct() {
	fmt.Println("\nShow Struct")
	// Disperate types together, each field is defined at compile time (cannot be changed at run time)
	type user struct {
		// By default, each field is initialised to its zero value
		ID         int    // 0
		FirstName  string // empty
		SecondName string // empty
	}

	var user1 user
	fmt.Println("Uninitialised struct", user1)
	user1.ID = 1
	user1.FirstName = "Trevor"
	user1.SecondName = "Whinmill"
	fmt.Println(user1)

	// Initialised format
	user2 := user{ID: 2,
		FirstName:  "Elaine",
		SecondName: "Whinmill", // Needs a comma on the last line
	}
	fmt.Println(user2)
}
