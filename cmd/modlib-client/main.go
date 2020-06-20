package main

import (
	"fmt"

	"github.com/prabhu1990/modlib"
	"github.com/prabhu1990/modlib/clientlib"
)

func main() {
	fmt.Println("Running Client")
	fmt.Println("Config: ", modlib.Config())
	fmt.Println(clientlib.Hello())
}
