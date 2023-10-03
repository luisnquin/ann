package faker

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/luisnquin/randatetime"
	"github.com/malisit/kolpa"
	gonanoid "github.com/matoous/go-nanoid/v2"
	piozFaker "github.com/pioz/faker"
)

type Generator struct {
	rand   *rand.Rand
	faker2 *gofakeit.Faker
	faker3 kolpa.Generator
}

func NewGenerator() Generator {
	seed := time.Now().Unix()

	return Generator{
		rand:   rand.New(rand.NewSource(seed)),
		faker2: gofakeit.New(seed),
		faker3: kolpa.C("en_US"),
	}
}

func (g Generator) Username() string {
	return piozFaker.Username()
}

func (g Generator) DateTime() string {
	return randatetime.BetweenYears(2020, 2023).UTC().Format(time.RFC3339)
}

func (g Generator) Date() string {
	return randatetime.BetweenYears(2020, 2023).UTC().Format(time.DateOnly)
}

func (g Generator) Email() string {
	return piozFaker.SafeEmail()
}

func (g Generator) NanoID() string {
	id, _ := gonanoid.Generate("-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 14)
	if id[0] == '-' || id[len(id)-1] == '-' {
		return g.NanoID()
	}

	return id
}

func (g Generator) UUID() string {
	return uuid.NewString()
}

func (g Generator) PhoneNumber() string {
	return g.faker3.Phone()
}

func (g Generator) FullName() string {
	return fmt.Sprintf("%s %s", g.faker2.FirstName(), g.faker2.LastName())
}

func (g Generator) Sentence() string {
	return piozFaker.ParagraphWithSentenceCount(g.getRandNot0(3))
}

func (g Generator) CreditCardNumber() string {
	return g.faker2.CreditCardNumber(&gofakeit.CreditCardOptions{
		Types: []string{"visa", "mastercard"},
		Gaps:  true,
	})
}

func (g Generator) PostalCode() string {
	return g.faker3.Postcode()
}

func (g Generator) Address() string {
	address := g.faker2.Address()

	return fmt.Sprintf("%s, %s, %s %s", address.Street, address.City, address.Country, address.Zip)
}

func (g Generator) City() string {
	return piozFaker.AddressCity()
}

func (g Generator) HexColor() string {
	return g.faker2.HexColor()
}

func (g Generator) EmployeeCode() string {
	const (
		possibleLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		possibleNumbers = "123456789"

		maxLen = 8
	)

	var b strings.Builder

	for i := 0; i < maxLen; i++ {
		crumb := possibleNumbers[g.getRandNot0(len(possibleNumbers))]
		b.WriteByte(crumb)
	}

	finalCrumb := possibleLetters[g.getRandNot0(len(possibleLetters))]
	b.WriteByte(finalCrumb)

	return b.String()
}

func (g Generator) getRandNot0(maxN int) int {
	n := g.rand.Intn(maxN)
	if n == 0 {
		return g.getRandNot0(maxN)
	}

	return n
}
