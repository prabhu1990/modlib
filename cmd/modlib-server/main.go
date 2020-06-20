package main

import (
	"fmt"

	"github.com/prabhu1990/internal/auth"
	"github.com/prabhu1990/modlib"
	"github.com/prabhu1990/modlib/serverlib"
)

func main() {
	fmt.Println("Running server")
	fmt.Println("Config: ", modlib.Config())
	fmt.Println("Auth: ", auth.Auth())
	fmt.Println(serverlib.Hello())
}
