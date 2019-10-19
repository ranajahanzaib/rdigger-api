package main

import (
	"fmt"

	nspx "github.com/nspx/core"
)

func main() {
	// Init message
	fmt.Println("Booting up rdigger...")

	// See if the nspx is imported.
	nspx.Hello()
}
