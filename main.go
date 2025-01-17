package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

var version = "0.0.1"

func main() {
	pass, error := func() (string, error) {
		cmdFlag := flag.NewFlagSet("cmdFlag", flag.ExitOnError)

		sep := cmdFlag.String("sep", "-", "a separator to use between words")
		n := cmdFlag.Int("n", 3, "the number of words (passphrase) to generate, must be greater than or equal to 2")
		phrases := cmdFlag.Bool("phrases", true, "generate a phrase instead of a password")
		c := cmdFlag.Int("c", 8, "the number of characters (token) to generate, must be greater than or equal to 5")
		v := cmdFlag.Bool("version", false, "print the version")

		cmdFlag.Parse(os.Args[1:])

		if *v {
			fmt.Println("Version:", version)
			os.Exit(0)
		}

		if *phrases {
			passgen, error := generateWords(*n, *sep)

			if error != nil {
				return "", error
			} else {
				return passgen, nil
			}
		} else {
			passgen, error := generateTokens(*c)

			if error != nil {
				return "", error
			} else {
				return passgen, nil
			}
		}
	}()

	defer func() {
		if error == nil {
			fmt.Println(pass)
			clipboard.WriteAll(pass)
			os.Exit(0)
		}
		os.Exit(1)
	}()
}
