package form3_test

import (
	form3 "github.com/orasik/form3/accounts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Account_Invalid_Country(t *testing.T) {
	acc := &form3.Account{}
	acc.Country = "a"
	err := form3.NewAccount(acc)
	assert.Equal(t, form3.ErrorCountryCodeShouldBe2Characters, err)
}

func Test_Account_More_Than_2_Characters_For_Country(t *testing.T) {
	acc := &form3.Account{}
	acc.Country = "United Kingdom"
	err := form3.NewAccount(acc)
	assert.Equal(t, form3.ErrorCountryCodeShouldBe2Characters, err)
}

func Test_Account_Invalid_Classification(t *testing.T) {
	acc := &form3.Account{}
	acc.Country = "GB"
	acc.AccountClassification = "Something"
	err := form3.NewAccount(acc)
	assert.Equal(t, form3.ErrorInvalidAccountClassification, err)
}

func Test_Account_Valid_Country_And_Classification(t *testing.T) {
	acc := &form3.Account{}
	acc.Country = "gb"
	acc.AccountClassification = "Personal"
	err := form3.NewAccount(acc)
	assert.Equal(t, nil, err)
	assert.Equal(t, "GB", acc.Country)
}

func Test_Account_GB_Validation(t *testing.T) {
	acc := &form3.Account{}
	acc.Country = "gb"
	acc.BankID = "123456"
	acc.BankIDCode = "GBDSC"
	acc.Bic = "Bic"
	acc.AccountClassification = "Personal"
	err := form3.NewAccount(acc)
	assert.Equal(t, nil, err)
	assert.Equal(t, "GB", acc.Country)
	ok, err := acc.Validate()

	assert.Equal(t, true, ok)
	assert.Equal(t, nil, err)
}

func Test_Account_AU_Validation(t *testing.T) {
	acc := &form3.Account{}
	acc.Country = "Au"
	acc.BankID = "123456"
	acc.BankIDCode = "AUBSB"
	acc.Bic = "Bic"
	acc.AccountClassification = "Business"
	err := form3.NewAccount(acc)
	assert.Equal(t, nil, err)
	assert.Equal(t, "AU", acc.Country)
	ok, err := acc.Validate()

	assert.Equal(t, true, ok)
	assert.Equal(t, nil, err)
}

func Test_Account_BE_Validation(t *testing.T) {
	acc := &form3.Account{}
	acc.Country = "bE"
	acc.BankID = "123"
	acc.BankIDCode = "BE"
	acc.AccountNumber = "1234567"
	acc.AccountClassification = "Business"
	err := form3.NewAccount(acc)
	assert.Equal(t, nil, err)
	assert.Equal(t, "BE", acc.Country)
	ok, err := acc.Validate()

	assert.Equal(t, true, ok)
	assert.Equal(t, nil, err)
}

func Test_Account_CA_Validation(t *testing.T) {
	acc := &form3.Account{}
	acc.Country = "ca"
	acc.BankID = "019184743"
	acc.BankIDCode = "CACPA"
	acc.Bic = "BIC"
	acc.AccountClassification = "Personal"
	err := form3.NewAccount(acc)
	assert.Equal(t, nil, err)
	assert.Equal(t, "CA", acc.Country)
	ok, err := acc.Validate()

	assert.Equal(t, true, ok)
	assert.Equal(t, nil, err)
}

func Test_Account_Validation_With_Unsupported_Country(t *testing.T) {
	acc := &form3.Account{}
	acc.Country = "cc"
	acc.AccountClassification = "Personal"
	err := form3.NewAccount(acc)
	assert.Equal(t, nil, err)
	assert.Equal(t, "CC", acc.Country)
	ok, err := acc.Validate()

	assert.Equal(t, false, ok)
	assert.Equal(t, form3.ErrorUnsupportedCountry, err)
}