package form3

import (
	"errors"
	"fmt"
)

const BEBankIDCode = "BE"

var ErrorBEBankIDShouldBe3Characters = errors.New("BE bank ID should be 3 characters")
var ErrorBEBankIDCode = errors.New(fmt.Sprintf("BE bank account BankID code should be %s", BEBankIDCode))
var ErrorBEAccountNumberHasToBe7Characters = errors.New("BE account number should be 7 characters")

func (ac *AccountDetails) ValidateBEAccount() (bool, error) {
	if len(ac.BankID) != 3 {
		return false, ErrorBEBankIDShouldBe3Characters
	}

	if ac.BankIDCode != BEBankIDCode {
		return false, ErrorBEBankIDCode
	}

	if len(ac.AccountNumber) != 7 {
		return false, ErrorBEAccountNumberHasToBe7Characters
	}

	return true, nil
}
