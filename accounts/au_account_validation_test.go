package form3_test

import (
	form3 "github.com/orasik/form3/accounts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Missing_AU_BankID(t *testing.T) {
	account := &form3.AccountDetails{
		Country:    "AU",
		Bic:        "BIC",
		BankIDCode: "AUBSB",
	}

	ok, err := account.ValidateAUAccount()

	if ok != true {
		t.Errorf("Should be invalid")
	}
	if err != nil {
		t.Errorf("Should be valid")
	}

}

func Test_AU_BankID_Larger_Than_6_Characters(t *testing.T) {
	account := &form3.AccountDetails{
		Country: "AU",
		BankID:  "1234567",
	}

	ok, err := account.ValidateAUAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorAUBankIDShouldBe6Characters {
		t.Errorf("Should raise Missing BankID code")
	}

}

func Test_AU_BankID_Less_Than_6_Characters(t *testing.T) {
	account := &form3.AccountDetails{
		Country: "AU",
		BankID:  "123",
	}

	ok, err := account.ValidateAUAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorAUBankIDShouldBe6Characters {
		t.Errorf("Should raise Missing BankID code")
	}

}

func Test_Missing_AU_BIC(t *testing.T) {
	account := &form3.AccountDetails{
		Country: "AU",
		BankID:  "123456",
	}

	ok, err := account.ValidateAUAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorAUBICIsRequired {
		t.Errorf("Should raise Missing BIC")
	}

}

func Test_Missing_AU_Bank_Code(t *testing.T) {
	account := &form3.AccountDetails{
		Country: "AU",
		BankID:  "123456",
		Bic:     "BIC",
	}

	ok, err := account.ValidateAUAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorAUBankIDCode {
		t.Errorf("Should raise missing Bank ID Code")
	}

}

func Test_Valid_AU_Account(t *testing.T) {
	account := &form3.AccountDetails{
		Country:    "AU",
		BankID:     "123456",
		Bic:        "BIC",
		BankIDCode: "AUBSB",
	}

	ok, err := account.ValidateAUAccount()

	assert.Equal(t, ok, true)
	if ok != true {
		t.Errorf("Should be valid")
	}
	if err != nil {
		t.Errorf("Error should be nil")
		t.Error(err)
	}

}
