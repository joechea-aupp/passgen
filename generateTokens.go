package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func generateTokens(n int) (string, error) {
	if n < 5 {
		return "", errors.New("[ERR] -n must be greater than or equal to 5")
	}
	token := make([]rune, n)

	for i := 0; i < n; i++ {

		specialChar := rand.Intn(15) + 33
		uppercaseLetter := rand.Intn(26) + 'A'
		lowercaseLetter := rand.Intn(26) + 'a'

		characters := []rune{rune(specialChar), rune(uppercaseLetter), rune(lowercaseLetter)}
		token[i] = characters[rand.Intn(len(characters))]
	}

	return fmt.Sprintf("%s", string(token)), nil
}
