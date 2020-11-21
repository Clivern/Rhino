// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"github.com/bxcodec/faker/v3"
)

// Faker type
type Faker struct {
	Latitude           float32 `faker:"lat"`
	Longitude          float32 `faker:"long"`
	CreditCardNumber   string  `faker:"cc_number"`
	CreditCardType     string  `faker:"cc_type"`
	Email              string  `faker:"email"`
	DomainName         string  `faker:"domain_name"`
	IPV4               string  `faker:"ipv4"`
	IPV6               string  `faker:"ipv6"`
	Password           string  `faker:"password"`
	PhoneNumber        string  `faker:"phone_number"`
	MacAddress         string  `faker:"mac_address"`
	URL                string  `faker:"url"`
	UserName           string  `faker:"username"`
	TollFreeNumber     string  `faker:"toll_free_number"`
	E164PhoneNumber    string  `faker:"e_164_phone_number"`
	TitleMale          string  `faker:"title_male"`
	TitleFemale        string  `faker:"title_female"`
	FirstName          string  `faker:"first_name"`
	FirstNameMale      string  `faker:"first_name_male"`
	FirstNameFemale    string  `faker:"first_name_female"`
	LastName           string  `faker:"last_name"`
	Name               string  `faker:"name"`
	UnixTime           int64   `faker:"unix_time"`
	Date               string  `faker:"date"`
	Time               string  `faker:"time"`
	MonthName          string  `faker:"month_name"`
	Year               string  `faker:"year"`
	DayOfWeek          string  `faker:"day_of_week"`
	DayOfMonth         string  `faker:"day_of_month"`
	Timestamp          string  `faker:"timestamp"`
	Century            string  `faker:"century"`
	TimeZone           string  `faker:"timezone"`
	TimePeriod         string  `faker:"time_period"`
	Word               string  `faker:"word"`
	Sentence           string  `faker:"sentence"`
	Paragraph          string  `faker:"paragraph"`
	Currency           string  `faker:"currency"`
	Amount             float64 `faker:"amount"`
	AmountWithCurrency string  `faker:"amount_with_currency"`
	UUIDHypenated      string  `faker:"uuid_hyphenated"`
	UUID               string  `faker:"uuid_digit"`
}

// Transform populate faked data
func (f *Faker) Transform(data string) (string, error) {
	types := f.GetTypesFound(data)

	err := faker.FakeData(f)

	if err != nil {
		return data, err
	}

	flags := f.GetFakeData()

	for i := 0; i < len(types); i++ {
		if strings.HasPrefix(types[i], "@fake(:anyof[") {
			item := strings.TrimPrefix(types[i], "@fake(:anyof[")
			item = strings.TrimSuffix(item, "])")
			items := strings.Split(item, "||")
			data = strings.Replace(
				data,
				types[i],
				items[rand.Intn(len(items))],
				-1,
			)
		} else {
			if val, ok := flags[types[i]]; ok {
				data = strings.Replace(data, types[i], val, -1)
			}
		}
	}

	return data, nil
}

// GetTypesFound grep all fake data tags
func (f *Faker) GetTypesFound(data string) []string {
	result := []string{}
	r := regexp.MustCompile(`@fake\((.)+?\)`)
	matches := r.FindAllStringIndex(data, -1)

	for n := 0; n < len(matches); n++ {
		result = append(
			result,
			data[matches[n][0]:matches[n][1]],
		)
	}

	return result
}

// GetFakeData gets a map of fake data
func (f *Faker) GetFakeData() map[string]string {
	flags := make(map[string]string)

	flags["@fake(:lat)"] = fmt.Sprintf("%f", f.Latitude)
	flags["@fake(:long)"] = fmt.Sprintf("%f", f.Longitude)
	flags["@fake(:amount)"] = fmt.Sprintf("%f", f.Amount)
	flags["@fake(:cc_number)"] = f.CreditCardNumber
	flags["@fake(:cc_type)"] = f.CreditCardType
	flags["@fake(:email)"] = f.Email
	flags["@fake(:domain_name)"] = f.DomainName
	flags["@fake(:ipv4)"] = f.IPV4
	flags["@fake(:ipv6)"] = f.IPV6
	flags["@fake(:password)"] = f.Password
	flags["@fake(:phone_number)"] = f.PhoneNumber
	flags["@fake(:mac_address)"] = f.MacAddress
	flags["@fake(:url)"] = f.URL
	flags["@fake(:username)"] = f.UserName
	flags["@fake(:toll_free_number)"] = f.TollFreeNumber
	flags["@fake(:e_164_phone_number)"] = f.E164PhoneNumber
	flags["@fake(:title_male)"] = f.TitleMale
	flags["@fake(:title_female)"] = f.TitleFemale
	flags["@fake(:first_name)"] = f.FirstName
	flags["@fake(:first_name_male)"] = f.FirstNameMale
	flags["@fake(:first_name_female)"] = f.FirstNameFemale
	flags["@fake(:last_name)"] = f.LastName
	flags["@fake(:name)"] = f.Name
	flags["@fake(:unix_time)"] = fmt.Sprintf("%d", f.UnixTime)
	flags["@fake(:date)"] = f.Date
	flags["@fake(:time)"] = f.Time
	flags["@fake(:month_name)"] = f.MonthName
	flags["@fake(:year)"] = f.Year
	flags["@fake(:day_of_week)"] = f.DayOfWeek
	flags["@fake(:day_of_month)"] = f.DayOfMonth
	flags["@fake(:timestamp)"] = f.Timestamp
	flags["@fake(:century)"] = f.Century
	flags["@fake(:timezone)"] = f.TimeZone
	flags["@fake(:time_period)"] = f.TimePeriod
	flags["@fake(:word)"] = f.Word
	flags["@fake(:sentence)"] = f.Sentence
	flags["@fake(:paragraph)"] = f.Paragraph
	flags["@fake(:currency)"] = f.Currency
	flags["@fake(:amount_with_currency)"] = f.AmountWithCurrency
	flags["@fake(:uuid_hyphenated)"] = f.UUIDHypenated
	flags["@fake(:uuid_digit)"] = f.UUID

	return flags
}
