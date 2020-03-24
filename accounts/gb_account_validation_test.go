package form3_test

import (
	form3 "github.com/orasik/form3/accounts"
	"testing"
)

func Test_Missing_GB_BankID(t *testing.T) {
	account := &form3.Account{
		AccountDetails: form3.AccountDetails{
			Country: "GB",
			BankID:  "123456",
			Bic:     "BIC",
		},
	}
	ok, err := account.ValidateGBAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorGBBankIDCode {
		t.Errorf("Should raise Missing BankID code")
	}

}

func Test_GB_BankID_Larger_Than_6_Characters(t *testing.T) {
	account := &form3.Account{
		AccountDetails: form3.AccountDetails{
			Country: "GB",
			BankID:  "1234567",
			Bic:     "BIC",
		},
	}

	ok, err := account.ValidateGBAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorGBBankIDShouldBe6Characters {
		t.Errorf("Should raise BankID should be 6 characters only")
	}

}

func Test_GB_BankID_Less_Than_6_Characters(t *testing.T) {
	account := &form3.Account{
		AccountDetails: form3.AccountDetails{
			Country: "GB",
			BankID:  "12345",
			Bic:     "BIC",
		},
	}

	ok, err := account.ValidateGBAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorGBBankIDShouldBe6Characters {
		t.Errorf("Should raise BankID must be 6 characters")
	}

}

func Test_GB_Missing_BIC(t *testing.T) {
	account := &form3.Account{
		AccountDetails: form3.AccountDetails{
			Country: "GB",
			BankID:  "123456",
		},
	}

	ok, err := account.ValidateGBAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorGBBICIsRequired {
		t.Errorf("Should raise Missing BIC")
	}

}

func Test_GB_Missing_Bank_Code(t *testing.T) {
	account := &form3.Account{
		AccountDetails: form3.AccountDetails{
			Country: "GB",
			BankID:  "123456",
			Bic:     "BIC",
		},
	}

	ok, err := account.ValidateGBAccount()

	if ok != false {
		t.Errorf("Should be invalid")
	}
	if err != form3.ErrorGBBankIDCode {
		t.Errorf("Should raise missing Bank ID Code")
	}

}

func Test_Valid_GBA_account(t *testing.T) {
	account := &form3.Account{
		AccountDetails: form3.AccountDetails{
			Country:    "GB",
			BankID:     "123456",
			BankIDCode: "GBDSC",
			Bic:        "BIC",
		},
	}

	ok, err := account.ValidateGBAccount()

	if ok != true {
		t.Errorf("Should be valid")
	}
	if err != nil {
		t.Errorf("Should not raise an error")
		t.Error(err)
	}

}
