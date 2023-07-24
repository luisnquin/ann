package main

import (
	"math/rand"
	"time"
)

func getRandomQuote() string {
	quotes := getQuotes()

	rand.Seed(time.Now().Unix())

	return quotes[rand.Intn(len(quotes))]
}
