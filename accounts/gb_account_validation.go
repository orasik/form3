package form3

import (
	"errors"
	"fmt"
)

const GBBankIDCode = "GBDSC"

var ErrorGBBankIDShouldBe6Characters = errors.New("GB bank account should have bank ID with 6 characters")
var ErrorGBBICIsRequired = errors.New("GB bank account BIC is required")
var ErrorGBBankIDCode = errors.New(fmt.Sprintf("GB bank account BankID code should be %s", GBBankIDCode))

func (acc *Account) ValidateGBAccount() (bool, error) {
	if len(acc.AccountDetails.BankID) != 6 {
		return false, ErrorGBBankIDShouldBe6Characters
	}

	if acc.AccountDetails.Bic == "" {
		return false, ErrorGBBICIsRequired
	}

	if acc.AccountDetails.BankIDCode != GBBankIDCode {
		return false, ErrorGBBankIDCode
	}

	return true, nil
}
