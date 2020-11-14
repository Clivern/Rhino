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

	for i := 0; i < len(types); i++ {
		if strings.HasPrefix(types[i], "@fake(:anyof[") {
			item := strings.TrimPrefix(types[i], "@fake(:anyof[")
			item = strings.TrimSuffix(item, "])")
			items := strings.Split(item, "||")

			data = strings.Replace(data, types[i], items[rand.Intn(len(items))], -1)
		} else if types[i] == "@fake(:lat)" {
			data = strings.Replace(data, types[i], fmt.Sprintf("%f", f.Latitude), -1)
		} else if types[i] == "@fake(:long)" {
			data = strings.Replace(data, types[i], fmt.Sprintf("%f", f.Longitude), -1)
		} else if types[i] == "@fake(:cc_number)" {
			data = strings.Replace(data, types[i], f.CreditCardNumber, -1)
		} else if types[i] == "@fake(:cc_type)" {
			data = strings.Replace(data, types[i], f.CreditCardType, -1)
		} else if types[i] == "@fake(:email)" {
			data = strings.Replace(data, types[i], f.Email, -1)
		} else if types[i] == "@fake(:domain_name)" {
			data = strings.Replace(data, types[i], f.DomainName, -1)
		} else if types[i] == "@fake(:ipv4)" {
			data = strings.Replace(data, types[i], f.IPV4, -1)
		} else if types[i] == "@fake(:ipv6)" {
			data = strings.Replace(data, types[i], f.IPV6, -1)
		} else if types[i] == "@fake(:password)" {
			data = strings.Replace(data, types[i], f.Password, -1)
		} else if types[i] == "@fake(:phone_number)" {
			data = strings.Replace(data, types[i], f.PhoneNumber, -1)
		} else if types[i] == "@fake(:mac_address)" {
			data = strings.Replace(data, types[i], f.MacAddress, -1)
		} else if types[i] == "@fake(:url)" {
			data = strings.Replace(data, types[i], f.URL, -1)
		} else if types[i] == "@fake(:username)" {
			data = strings.Replace(data, types[i], f.UserName, -1)
		} else if types[i] == "@fake(:toll_free_number)" {
			data = strings.Replace(data, types[i], f.TollFreeNumber, -1)
		} else if types[i] == "@fake(:e_164_phone_number)" {
			data = strings.Replace(data, types[i], f.E164PhoneNumber, -1)
		} else if types[i] == "@fake(:title_male)" {
			data = strings.Replace(data, types[i], f.TitleMale, -1)
		} else if types[i] == "@fake(:title_female)" {
			data = strings.Replace(data, types[i], f.TitleFemale, -1)
		} else if types[i] == "@fake(:first_name)" {
			data = strings.Replace(data, types[i], f.FirstName, -1)
		} else if types[i] == "@fake(:first_name_male)" {
			data = strings.Replace(data, types[i], f.FirstNameMale, -1)
		} else if types[i] == "@fake(:first_name_female)" {
			data = strings.Replace(data, types[i], f.FirstNameFemale, -1)
		} else if types[i] == "@fake(:last_name)" {
			data = strings.Replace(data, types[i], f.LastName, -1)
		} else if types[i] == "@fake(:name)" {
			data = strings.Replace(data, types[i], f.Name, -1)
		} else if types[i] == "@fake(:unix_time)" {
			data = strings.Replace(data, types[i], fmt.Sprintf("%d", f.UnixTime), -1)
		} else if types[i] == "@fake(:date)" {
			data = strings.Replace(data, types[i], f.Date, -1)
		} else if types[i] == "@fake(:time)" {
			data = strings.Replace(data, types[i], f.Time, -1)
		} else if types[i] == "@fake(:month_name)" {
			data = strings.Replace(data, types[i], f.MonthName, -1)
		} else if types[i] == "@fake(:year)" {
			data = strings.Replace(data, types[i], f.Year, -1)
		} else if types[i] == "@fake(:day_of_week)" {
			data = strings.Replace(data, types[i], f.DayOfWeek, -1)
		} else if types[i] == "@fake(:day_of_month)" {
			data = strings.Replace(data, types[i], f.DayOfMonth, -1)
		} else if types[i] == "@fake(:timestamp)" {
			data = strings.Replace(data, types[i], f.Timestamp, -1)
		} else if types[i] == "@fake(:century)" {
			data = strings.Replace(data, types[i], f.Century, -1)
		} else if types[i] == "@fake(:timezone)" {
			data = strings.Replace(data, types[i], f.TimeZone, -1)
		} else if types[i] == "@fake(:time_period)" {
			data = strings.Replace(data, types[i], f.TimePeriod, -1)
		} else if types[i] == "@fake(:word)" {
			data = strings.Replace(data, types[i], f.Word, -1)
		} else if types[i] == "@fake(:sentence)" {
			data = strings.Replace(data, types[i], f.Sentence, -1)
		} else if types[i] == "@fake(:paragraph)" {
			data = strings.Replace(data, types[i], f.Paragraph, -1)
		} else if types[i] == "@fake(:currency)" {
			data = strings.Replace(data, types[i], f.Currency, -1)
		} else if types[i] == "@fake(:amount)" {
			data = strings.Replace(data, types[i], fmt.Sprintf("%f", f.Amount), -1)
		} else if types[i] == "@fake(:amount_with_currency)" {
			data = strings.Replace(data, types[i], f.AmountWithCurrency, -1)
		} else if types[i] == "@fake(:uuid_hyphenated)" {
			data = strings.Replace(data, types[i], f.UUIDHypenated, -1)
		} else if types[i] == "@fake(:uuid_digit)" {
			data = strings.Replace(data, types[i], f.UUID, -1)
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
