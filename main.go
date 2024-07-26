package main

import (
	"fmt"

	"github.com/tjarratt/babble"
)

func main() {
	babbler := babble.NewBabbler()
	fmt.Println(babbler.Babble())

	// optionally set your own separater
	babbler.Separator = "~"
	fmt.Println(babbler.Babble())
}
