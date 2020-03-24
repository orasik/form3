package form3

import (
	"errors"
	"strings"
)

const AccountClassificationPersonal = "Personal"
const AccountClassificationBusiness = "Business"

var ErrorCountryCodeShouldBe2Characters = errors.New("country code can not be more than 2 characters")
var ErrorInvalidAccountClassification = errors.New("AccountClassification should be either Personal or Business")
var ErrorUnsupportedCountry = errors.New("unsupported country")

type ClientBody struct {
	*Account `json:"data"`
}

type ListResponse struct {
	Arr []*Account `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"error_message"`
}

type Account struct {
	ID             string `json:"id"`
	Type           string `json:"type"`
	OrganisationID string `json:"organisation_id"`
	AccountDetails `json:"attributes"`
	Version        int8 `json:"version"`
}

type AccountDetails struct {
	Country               string `json:"country"`
	BaseCurrency          string `json:"base_currency"`
	AccountNumber         string `json:"account_number"`
	BankID                string `json:"bank_id"`
	BankIDCode            string `json:"bank_id_code"`
	Bic                   string `json:"bic"`
	AccountClassification string `json:"account_classification"`
	JointAccount          bool   `json:"joint_account"`
	AccountMatchingOptOut bool   `json:"account_matching_opt_out"`
	PrivateIdentification `json:"private_identification"`
}

type PrivateIdentification struct {
	Title          string `json:"title"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	BirthDate      string `json:"birth_date"`
	BirthCountry   string `json:"birth_country"`
	DocumentNumber string `json:"document_number"`
	Address        string `json:"address"`
	City           string `json:"city"`
	Country        string `json:"country"`
}

func NewAccount(acc *Account) error {
	if len(acc.Country) != 2 {
		return ErrorCountryCodeShouldBe2Characters
	}

	if ok, err := validateAccountClassification(acc.AccountClassification); !ok {
		return err
	}

	acc.Country = strings.ToUpper(acc.Country)

	return nil
}

func (acc *Account) Validate() (bool, error) {

	switch acc.AccountDetails.Country {
	case "GB":
		return acc.ValidateGBAccount()
	case "CA":
		return acc.ValidateCAAccount()
	case "AU":
		return acc.ValidateAUAccount()
	case "BE":
		return acc.ValidateBEAccount()
	default:
		return false, ErrorUnsupportedCountry
	}
}

func validateAccountClassification(ac string) (bool, error) {
	if ac != AccountClassificationPersonal && ac != AccountClassificationBusiness {
		return false,ErrorInvalidAccountClassification
	}

	return true, nil
}
