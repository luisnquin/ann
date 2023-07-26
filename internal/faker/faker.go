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
)

var (
	goFakeItFaker *gofakeit.Faker
	kolpaFaker    kolpa.Generator
)

func init() {
	rand.Seed(time.Now().Unix())
	goFakeItFaker = gofakeit.New(time.Now().Unix())
	kolpaFaker = kolpa.C("en_US")
}

func Username() string {
	return goFakeItFaker.Username()
}

func DateTime() string {
	return randatetime.BetweenYears(2020, 2023).UTC().Format(time.RFC3339)
}

func Date() string {
	return randatetime.BetweenYears(2020, 2023).UTC().Format(time.DateOnly)
}

func Email() string {
	return goFakeItFaker.Email()
}

func NanoID() string {
	return gonanoid.Must()
}

func UUID() string {
	return uuid.NewString()
}

func PhoneNumber() string {
	return kolpaFaker.Phone()
}

func FullName() string {
	return fmt.Sprintf("%s %s", goFakeItFaker.FirstName(), goFakeItFaker.LastName())
}

func LoremSentence() string {
	return kolpaFaker.LoremSentence()
}

func CreditCardNumber() string {
	return goFakeItFaker.CreditCardNumber(&gofakeit.CreditCardOptions{
		Types: []string{"visa", "mastercard"},
		Gaps:  true,
	})
}

func PostalCode() string {
	if false {
		return goFakeItFaker.Zip()
	}

	return kolpaFaker.Postcode()
}

func Address() string {
	address := goFakeItFaker.Address()

	return fmt.Sprintf("%s, %s, %s %s", address.Street, address.City, address.Country, address.Zip)
}

func City() string {
	return kolpaFaker.City()
}

func HexColor() string {
	return goFakeItFaker.HexColor()
}

func EmployeeCode() string {
	maxLen := 8

	const possibleLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const possibleNumbers = "123456789"

	var b strings.Builder

	for i := 0; i < maxLen; i++ {
		crumb := possibleNumbers[rand.Intn(len(possibleNumbers))]
		b.WriteByte(crumb)
	}

	finalCrumb := possibleLetters[rand.Intn(len(possibleLetters))]
	b.WriteByte(finalCrumb)

	return b.String()
}
