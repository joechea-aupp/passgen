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
	phrases := cmdFlag.Bool("phrases", true, "generate a phrase instead of a password")
	c := cmdFlag.Int("c", 8, "the number of characters to generate, must be greater than or equal to 5")

	cmdFlag.Parse(os.Args[1:])

	if *phrases {
		passgen, error := generateWords(*n, *sep)

		if error != nil {
			fmt.Println(error)
			os.Exit(1)
		} else {
			fmt.Println(passgen)
			os.Exit(0)
		}
	} else {
		passgen, error := generateTokens(*c)

		if error != nil {
			fmt.Println(error)
			os.Exit(1)
		} else {
			fmt.Println(passgen)
			os.Exit(0)
		}
	}
}
