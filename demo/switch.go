package demo

import "fmt"

func switchOnString(input string) string {
	var output string
	switch input {
	case "1":
		output = "Hit case 1"
	case "2":
		output = "Hit case 2 and fall through to "
		fallthrough
	case "3":
		output += "Hit case 3"
	default:
		output = "Unhandled input"
	}

	return output
}

func RunSwitchDemo() {
	fmt.Println("\nRun Switch")
	fmt.Println(switchOnString("1"))
	fmt.Println(switchOnString("2"))
	fmt.Println(switchOnString("3"))
	fmt.Println(switchOnString("4"))
}
