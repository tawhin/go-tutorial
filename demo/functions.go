package demo

import "fmt"

func RunFuncsDemo() {
	// we are only interested in any errors returned.
	// _ is a write only var, which hides the returned port
	_, err := startWebServer(3000, 3)
	fmt.Println(err)
}

// comma delimited parameters have the same type
// returns an int and an error.
func startWebServer(port, retries int) (int, error) {
	var serverErr error = nil
	fmt.Println("Starting web server...")
	// do something

	fmt.Println("Web server started on port", port)
	return port, serverErr
}
