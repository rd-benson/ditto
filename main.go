package main

import (
	"github.com/rd-benson/ditto/pigeon"
)

// ditto!
func main() {
	// Using project config for testing
	pigeon.Start("./config/pigeon.yaml")
}
