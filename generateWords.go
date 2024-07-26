package main

import (
	"errors"
	"fmt"

	"github.com/tjarratt/babble"
)

func generateWords(n int, sep string) (string, error) {
	if n < 2 {
		return "", errors.New("[ERR] -n must be greater than or equal to 2")
	}
	// instanciate the babbler
	babbler := babble.NewBabbler()

	// optionally set your own separater
	if sep != "" {
		babbler.Separator = sep
	} else {
		babbler.Separator = "-"
	}

	// control how about password it will generated
	babbler.Count = n

	return fmt.Sprintln(babbler.Babble()), nil
}
