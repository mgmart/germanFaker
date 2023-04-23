package germanFaker

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
)

type SalutationInfo struct {
	Salutation string
	Gender     string
}

type CityInfo struct {
	Name     string
	AreaCode string
	PostCode string
}

type Person struct {
	Salutation string
	FirstName  string
	LastName   string
}

type Address struct {
	City     string
	PostCode string
	Phone    string
}

// Sample ...
type Sample struct {
	UUID    string  `faker:"uuid_digit"`
	Company string  `faker:"company"`
	Person  Person  `faker:"person"`
	Address Address `faker:"address"`
	Mobile  string  `faker:"germanmobilenumber"`
}

// CustomGenerator ...`
func CustomGenerator() {

	// Provides city, postcode and areacode
	_ = faker.AddProvider("address", func(v reflect.Value) (interface{}, error) {

		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		cityInfo := cityInfos[r1.Intn(len(cityInfos))]
		subscriberNumber := r1.Intn(10000000) // Generates a random number between 0 and 9999999
		phone := fmt.Sprintf("+49 %s %07d", cityInfo.AreaCode, subscriberNumber)

		return Address{
			Phone:    phone,
			City:     cityInfo.Name,
			PostCode: cityInfo.PostCode,
		}, nil
	})

	// Proivides Salutation, First- and Lastname
	_ = faker.AddProvider("person", func(v reflect.Value) (interface{}, error) {

		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		salutations := []SalutationInfo{
			{Salutation: "Herr", Gender: "male"},
			{Salutation: "Frau", Gender: "female"},
			{Salutation: "Dr.", Gender: "neutral"},
			{Salutation: "Graf", Gender: "male"},
			{Salutation: "Gräfin", Gender: "female"},
		}

		salut := salutations[r1.Intn(len(salutations))]
		var firstName string
		if salut.Gender == "male" {
			firstName = germanMaleFirstNames[r1.Intn(len(germanMaleFirstNames))]
		} else if salut.Gender == "female" {
			firstName = germanFemaleFirstNames[r1.Intn(len(germanFemaleFirstNames))]
		} else {
			gender := rand.Intn(2)
			if gender == 0 {
				firstName = germanMaleFirstNames[r1.Intn(len(germanMaleFirstNames))]
			} else {
				firstName = germanFemaleFirstNames[r1.Intn(len(germanFemaleFirstNames))]
			}
		}

		lastName := germanLastNames[r1.Intn(len(germanLastNames))]
		if salut.Salutation == "Graf" || salut.Salutation == "Gräfin" {
			lastName = "von " + lastName
		}
		ret := Person{
			Salutation: salut.Salutation,
			FirstName:  firstName,
			LastName:   lastName,
		}
		return ret, nil
	})

	// Provides Randon Company Name
	_ = faker.AddProvider("company", func(v reflect.Value) (interface{}, error) {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		ret := strings.Title(adjectives[r1.Intn(len(adjectives)-1)]) + "e " + substantives[r1.Intn(len(substantives)-1)] + " " + companyForm[r1.Intn(len(companyForm)-1)]
		return ret, nil
	})

	// Provides Random German mobile-phone numbers
	_ = faker.AddProvider("germanmobilenumber", func(v reflect.Value) (interface{}, error) {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		mobileNetworkPrefixes := []string{"151", "152", "157", "159", "160", "162", "163",
			"170", "171", "172", "173", "174", "175", "176", "177", "178", "179"}
		prefix := mobileNetworkPrefixes[r1.Intn(len(mobileNetworkPrefixes))]
		subscriberNumber := r1.Intn(10000000) // Generates a random number between 0 and 9999999
		return fmt.Sprintf("+49 %s %07d", prefix, subscriberNumber), nil
	})
}
