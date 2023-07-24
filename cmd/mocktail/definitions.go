package main

import "github.com/luisnquin/mocktail/internal/faker"

type generatorItem struct {
	name      string
	generator func() string
}

func getGenerators() []generatorItem {
	return []generatorItem{
		{
			name:      "UUID",
			generator: faker.UUID,
		},
		{
			name:      "Nano ID",
			generator: faker.NanoID,
		},
		{
			name:      "Date time (UTC)",
			generator: faker.DateTime,
		},
		{
			name:      "Date",
			generator: faker.Date,
		},
		{
			name:      "Email",
			generator: faker.Email,
		},
		{
			name:      "Full name",
			generator: faker.FullName,
		},
		{
			name:      "Username",
			generator: faker.Username,
		},
		{
			name:      "Phone number",
			generator: faker.PhoneNumber,
		},
		{
			name:      "Credit card",
			generator: faker.CreditCardNumber,
		},
		{
			name:      "Lorem sentence",
			generator: faker.LoremSentence,
		},
		{
			name:      "Postal code",
			generator: faker.PostalCode,
		},
		{
			name:      "City",
			generator: faker.City,
		},
		{
			name:      "Address",
			generator: faker.Address,
		},
		{
			name:      "Hexadecimal color",
			generator: faker.HexColor,
		},
	}
}

func getQuotes() []string {
	return []string{
		"Infinite stars, infinite possibilities.",
		"In the vastness lies the unknown.",
		"Life's beauty is in its mysteries.",
		"Knowledge begins with wonder.",
		"Seek truth, question everything.",
		"The mind's depth is boundless.",
		"Doubt, the cradle of wisdom.",
		"Within chaos, harmony resides.",
		"Embrace the enigma of existence.",
		"Wisdom is born from introspection.",
		"In the void, consciousness blooms.",
		"Perceive, therefore you are.",
		"Through darkness, light emerges.",
		"The heart seeks meaning beyond reason.",
		"The journey is the destination.",
		"Infinite perspectives, one reality.",
		"Truth whispers amidst the noise.",
		"Time reveals the essence of things.",
		"Infinite paths, finite choices.",
		"Simplicity hides profound complexity.",
		"Infinite questions, finite answers.",
		"The universe mirrors the soul.",
		"Essence transcends the material.",
		"In silence, wisdom speaks.",
		"Within nothingness lies potential.",
		"The unknown beckons the curious.",
		"Love is the canvas of existence.",
		"Perception shapes our reality.",
		"Through doubt, we find certainty.",
		"Ninym Ralei the best girl",
		"The mind is the universe's greatest enigma.",
		"Within the depths of uncertainty, lies truth.",
		"Time dances to the melody of eternity.",
		"The heart's language surpasses words.",
		"Through empathy, we glimpse the cosmos.",
		"Consciousness: the universe experiencing itself.",
		"The search for meaning is the journey of life.",
		"Existence is a symphony of interconnectedness.",
		"In the dance of atoms, life emerges.",
		"The soul's home is in the boundless.",
		"Infinite love expands the finite heart.",
		"In the mirror of nature, we find ourselves.",
		"The paradoxes of life reveal its depth.",
		"Thoughts are whispers of the soul.",
		"The universe breathes through every being.",
		"Knowledge humbles, wisdom uplifts.",
		"Through suffering, resilience is born.",
		"In the embrace of solitude, truth arises.",
		"The stars remind us of our smallness and greatness.",
		"The tapestry of reality weaves between worlds.",
		"The observer shapes the observed.",
		"In the dance of opposites, harmony is found.",
		"The cosmos speaks in the language of symbols.",
		"The search for purpose shapes our destiny.",
		"Infinite potential resides in the present moment.",
		"The universe: a library of untold stories.",
		"In the void of nothingness, all is possible.",
		"Through change, we glimpse eternity.",
		"The stars carry the dreams of humanity.",
	}
}
