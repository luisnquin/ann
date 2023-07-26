package main

import (
	"math/rand"
)

func getRandomQuote() string {
	quotes := getQuotes()

	return quotes[rand.Intn(len(quotes))]
}
