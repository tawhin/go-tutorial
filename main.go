package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pluralsight/webservice/controllers"
	"github.com/pluralsight/webservice/demo"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "demo" {
		fmt.Println("Run pluralsight demo code")
		demo.RunFuncsDemo()
		demo.RunTypesDemo()
		demo.RunLoopsDemo()
		demo.RunSwitchDemo()
	} else {
		fmt.Println("Run webserver localhost:3000")
		controllers.RegisterControllers()
		http.ListenAndServe(":3000", nil)
	}
}
