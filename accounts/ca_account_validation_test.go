package form3_test

import (
	form3 "github.com/orasik/form3/accounts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Empty_CA_BankID(t *testing.T) {
	account := &form3.Account{
		AccountDetails: form3.AccountDetails{
			Country:    "CA",
			BankIDCode: form3.CABankIDCode,
			Bic:        "BIC",
		},
	}
	ok, err := account.ValidateCAAccount()

	assert.Equal(t, ok, true)
	assert.Equal(t, err, nil)
}

func Test_CA_BankID_Should_Be_9_Characters_Starting_With_Zero(t *testing.T) {
	account := &form3.Account{
		AccountDetails: form3.AccountDetails{
			Country:    "CA",
			BankID:     "123456789",
			BankIDCode: form3.CABankIDCode,
			Bic:        "BIC",
		},
	}

	ok, err := account.ValidateCAAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorCABankIDShouldBe9Characters {
		t.Errorf("Should raise BankID should be 9 characters starting with 0")
	}

}

func Test_CA_BankID_Less_Than_9_Characters_Starting_With_Zero(t *testing.T) {
	account := &form3.Account{
		AccountDetails: form3.AccountDetails{
			Country:    "CA",
			BankID:     "0123",
			BankIDCode: form3.CABankIDCode,
			Bic:        "BIC",
		},
	}

	ok, err := account.ValidateCAAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorCABankIDShouldBe9Characters {
		t.Errorf("Should raise BankID should be 9 characters starting with 0")
	}

}
func Test_CA_Missing_BIC(t *testing.T) {
	account := &form3.Account{
		AccountDetails: form3.AccountDetails{
			Country: "GB",
			BankID:  "012345678",
		},
	}

	ok, err := account.ValidateCAAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorCABICIsRequired {
		t.Errorf("Should raise Missing BIC")
		t.Error(err)
	}

}

func Test_CA_Missing_Bank_Code(t *testing.T) {
	account := &form3.Account{
		AccountDetails: form3.AccountDetails{
			Country: "CA",
			BankID:  "012345678",
			Bic:     "BIC",
		},
	}

	ok, err := account.ValidateCAAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorCABankIDCode {
		t.Errorf("Should raise missing Bank ID Code")
	}

}

func Test_Valid_CA_Account(t *testing.T) {
	account := &form3.Account{
		AccountDetails: form3.AccountDetails{
			Country:    "CA",
			BankID:     "012345678",
			BankIDCode: form3.CABankIDCode,
			Bic:        "BIC",
		},
	}

	ok, err := account.ValidateCAAccount()
	assert.Equal(t, ok, true)
	assert.Equal(t, err, nil)

}
