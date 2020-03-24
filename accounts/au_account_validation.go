package form3

import (
	"errors"
	"fmt"
)

const AUBankIDCode = "AUBSB"

var ErrorAUBankIDShouldBe6Characters = errors.New("AU bank ID should be either empty or 6 characters")
var ErrorAUBICIsRequired = errors.New("AU bank account BIC is required")
var ErrorAUBankIDCode = errors.New(fmt.Sprintf("AU bank account BankID code should be %s", AUBankIDCode))

func (ac *AccountDetails) ValidateAUAccount() (bool, error) {
	if ac.BankID != "" && len(ac.BankID) != 6 {
		return false, ErrorAUBankIDShouldBe6Characters
	}

	if ac.Bic == "" {
		return false, ErrorAUBICIsRequired
	}

	if ac.BankIDCode != AUBankIDCode {
		return false, ErrorAUBankIDCode
	}

	return true, nil
}
