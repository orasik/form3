package form3

import (
	"errors"
	"fmt"
)

const CABankIDCode = "CACPA"

var ErrorCABankIDShouldBe9Characters = errors.New("CA bank account (if available) should be 9 characters starting with 0")
var ErrorCABICIsRequired = errors.New("CA bank account BIC is required")
var ErrorCABankIDCode = errors.New(fmt.Sprintf("CA bank account BankID code should be %s", CABankIDCode))

func (acc *Account) ValidateCAAccount() (bool, error) {

	if len(acc.AccountDetails.BankID) != 0 {
		if len(acc.AccountDetails.BankID) != 9 ||
			string(acc.AccountDetails.BankID[0]) != "0" {
			return false, ErrorCABankIDShouldBe9Characters
		}
	}

	if acc.AccountDetails.Bic == "" {
		return false, ErrorCABICIsRequired
	}

	if acc.AccountDetails.BankIDCode != CABankIDCode {
		return false, ErrorCABankIDCode
	}

	return true, nil
}
