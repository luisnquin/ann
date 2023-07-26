package main

import (
	"math/rand"
)

func getRandomProgramQuote() string {
	quotes := getProgramQuotes()

	return quotes[rand.Intn(len(quotes))]
}
