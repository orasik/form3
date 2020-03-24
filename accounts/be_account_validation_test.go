package form3_test

import (
	form3 "github.com/orasik/form3/accounts"
	"testing"
)

func Test_Missing_BE_BankID(t *testing.T) {
	account := &form3.AccountDetails{
		Country: "BE",
	}

	ok, err := account.ValidateBEAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorBEBankIDShouldBe3Characters {
		t.Errorf("Should raise Missing BankID code")
	}

}

func Test_BE_BankID_Larger_Than_3_Characters(t *testing.T) {
	account := &form3.AccountDetails{
		Country: "BE",
		BankID:  "1234",
	}

	ok, err := account.ValidateBEAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorBEBankIDShouldBe3Characters {
		t.Errorf("Should raise Missing BankID code")
	}

}

func Test_BE_BankID_Less_Than_3_Characters(t *testing.T) {
	account := &form3.AccountDetails{
		Country: "BE",
		BankID:  "12",
	}

	ok, err := account.ValidateBEAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorBEBankIDShouldBe3Characters {
		t.Errorf("Should raise BankID character length error")
	}

}

func Test_BE_Missing_Bank_Code(t *testing.T) {
	account := &form3.AccountDetails{
		Country: "BE",
		BankID:  "123",
	}

	ok, err := account.ValidateBEAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorBEBankIDCode {
		t.Errorf("Should raise missing Bank ID Code")
	}

}

func Test_Valid_BE_Account(t *testing.T) {
	account := &form3.AccountDetails{
		Country:       "BE",
		BankID:        "123",
		BankIDCode:    "BE",
		AccountNumber: "1234567",
	}

	ok, err := account.ValidateBEAccount()

	if ok != true {
		t.Errorf("Should be valid")
	}
	if err != nil {
		t.Errorf("Should be valid")
		t.Error(err)
	}

}
