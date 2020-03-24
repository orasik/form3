package form3

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gojektech/heimdall/httpclient"
	"github.com/orasik/form3/accounts"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type Response struct {
	*http.Response
}

type AccountHttpClient struct {
	HttpClient *httpclient.Client
	BaseURL    string
	EndPoint   string
}

var (
	// Create endpoint error messages
	ErrorAccountIsInvalid                           = errors.New("account data is invalid")
	ErrorCanNotMarshalJsonForCreateRequest          = errors.New("can not marshal JSON for Create request")
	ErrorUnmarshallingCreateResponseToAccountStruct = errors.New("Error unmarshalling Create response to Account Struct")
	ErrorReadingCreateResponseBosy                  = errors.New("error reading Create response body")
	ErrorUnmarshallingCreateErrorResponse           = errors.New("Error unmarshalling Create Error response")
)

// Create will create an account object and validate its data before sending the
// request. This validation is based on bank's country as mentioned here:
// https://api-docs.form3.tech/api.html#organisation-accounts
func (c *AccountHttpClient) Create(acc *form3.Account) (*form3.Account, error) {
	err := form3.NewAccount(acc)

	if err != nil {
		return nil, err
	}

	_, err = acc.Validate()

	if err != nil {
		log.Errorf("can not create a new account %s", err.Error())
		return nil, ErrorAccountIsInvalid
	}

	request := &form3.ClientBody{Account: acc}

	jsonRequest, err := json.Marshal(request)

	if err != nil {
		log.Error("Can not marshal JSON")
		return nil, ErrorCanNotMarshalJsonForCreateRequest
	}

	log.Debugf("Create Request %s", string(jsonRequest))
	r := bytes.NewReader(jsonRequest)
	res, err := c.HttpClient.Post(c.BaseURL+c.EndPoint, r, nil)
	if err != nil {
		log.Errorf("can not create account %s", err.Error())
		return nil, err
	}
	if res.StatusCode == 201 {
		body, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			log.Error(ErrorReadingCreateResponseBosy.Error())
			return nil, ErrorReadingCreateResponseBosy
		}
		log.Debugf("Response Code %d", res.StatusCode)
		obj := &form3.ClientBody{}
		err = json.Unmarshal(body, obj)
		if err != nil || obj.Account == nil {
			log.Error(ErrorUnmarshallingCreateResponseToAccountStruct.Error())
			return nil, ErrorUnmarshallingCreateResponseToAccountStruct
		}

		return obj.Account, nil
	}

	// If status code is not 201
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	log.Debugf("Response Code %d", res.StatusCode)
	obj := &form3.ErrorResponse{}
	err = json.Unmarshal(body, obj)
	if err != nil || obj.Message == "" {
		log.Error(ErrorUnmarshallingCreateErrorResponse.Error())
		return nil, ErrorUnmarshallingCreateErrorResponse
	}

	return nil, errors.New(fmt.Sprintf("can not create a new account %s", obj.Message))
}

// List will return all accounts in an organisation
func (c *AccountHttpClient) List(pageNumber uint8, pageSize uint8) ([]*form3.Account, error) {
	var errorMessage string
	if pageSize == 0 {
		pageSize = 100
	}
	uri := fmt.Sprintf("%s%s?page[number]=%d&page[size]=%d", c.BaseURL, c.EndPoint, pageNumber, pageSize)
	res, err := c.HttpClient.Get(uri, nil)
	if err != nil {
		errorMessage = fmt.Sprintf("error in sending List request to uri %s", uri)
		log.Error(errorMessage)

		return nil, errors.New(errorMessage)
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode == 200 {
		if err != nil {
			errorMessage = "error in List request"
			log.Error(errorMessage)

			return nil, errors.New(errorMessage)
		}
		obj := &form3.ListResponse{}
		err = json.Unmarshal(body, obj)
		if err != nil {
			errorMessage = "error in unmarshalling List json response"
			log.Error(errorMessage)

			return nil, errors.New(errorMessage)
		}

		return obj.Arr, nil
	}

	log.Errorf("List status code %d", res.StatusCode)
	log.Errorf("List Response %s", string(body))
	return nil, errors.New("error listing accounts")
}

// Fetch will fetch account details based on account uuid
func (c *AccountHttpClient) Fetch(accountID string) (*form3.Account, error) {
	uri := c.BaseURL + c.EndPoint + accountID
	res, err := c.HttpClient.Get(uri, nil)
	if err != nil {
		errorMessage := fmt.Sprintf("error preparing fetch request to uri %s", uri)
		log.Error(errorMessage)

		return nil, errors.New(errorMessage)
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		errorMessage := "error in Fetch request"
		log.Error(errorMessage)

		return nil, errors.New(errorMessage)
	}

	obj := &form3.ClientBody{}
	err = json.Unmarshal(body, obj)
	if err != nil || obj.Account == nil {
		errorMessage := "error in unmarshalling Fetch json response"
		log.Error(errorMessage)

		return nil, errors.New(errorMessage)
	}

	return obj.Account, nil
}

// Delete will remove the account from organisation (soft-delete)
func (c *AccountHttpClient) Delete(accountID string, version int8) (bool, error) {
	var errorMessage string
	uri := fmt.Sprintf("%s%s%s?version=%d", c.BaseURL, c.EndPoint, accountID, version)
	res, err := c.HttpClient.Delete(uri, nil)
	if err != nil {
		errorMessage = fmt.Sprintf("error in sedning Delete request to uri %s", uri)
		log.Error(errorMessage)
		return false, errors.New(errorMessage)
	}

	switch res.StatusCode {
	case 204:
		log.Infof("account %s has been deleted successfully", accountID)

		return true, nil
	case 404:
		errorMessage = fmt.Sprintf("error in deleting account %s %s", accountID, "Specified resource does not exist")
	case 409:
		errorMessage = fmt.Sprintf("error in deleting account %s %s %d", accountID, "Specified version incorrect", version)
	default:
		errorMessage = fmt.Sprintf("unknown response code %d", res.StatusCode)
	}

	return false, errors.New(errorMessage)
}
