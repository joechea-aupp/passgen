package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	cmdFlag := flag.NewFlagSet("cmdFlag", flag.ExitOnError)

	sep := cmdFlag.String("sep", "-", "a separator to use between words")
	n := cmdFlag.Int("n", 3, "the number of words to generate, must be greater than or equal to 2")

	cmdFlag.Parse(os.Args[1:])

	passgen, error := generateWords(*n, *sep)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(passgen)
	}
}
