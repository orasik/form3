package form3_test

import (
	"fmt"
	"github.com/gojektech/heimdall/httpclient"
	form3 "github.com/orasik/form3/accounts"
	client "github.com/orasik/form3/client"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// Disable logs during unit test
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}
func Test_Create_With_Wrong_Country_Code(t *testing.T) {
	acc, _ := makeAccount()
	acc.Country = "Some Random Country Name"

	c := client.AccountHttpClient{
		HttpClient: httpclient.NewClient(httpclient.WithHTTPTimeout(5 * time.Second)),
		BaseURL:    "url",
		EndPoint:   "",
	}

	_, err := c.Create(acc)

	assert.Equal(t, err, form3.ErrorCountryCodeShouldBe2Characters)
}

func Test_Create_With_Unexpected_Response(t *testing.T) {
	acc, _ := makeAccount()
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/", r.RequestURI)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{ "response": "ok" }`))
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	c := client.AccountHttpClient{
		HttpClient: httpclient.NewClient(httpclient.WithHTTPTimeout(5 * time.Second)),
		BaseURL:    server.URL,
		EndPoint:   "/v1/",
	}

	_, err := c.Create(acc)

	assert.Equal(t, client.ErrorUnmarshallingCreateResponseToAccountStruct, err)
}

func Test_Create_With_Error_Response(t *testing.T) {
	acc, _ := makeAccount()
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/", r.RequestURI)

		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{ "error_message": "something went wrong" }`))
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	c := client.AccountHttpClient{
		HttpClient: httpclient.NewClient(httpclient.WithHTTPTimeout(5 * time.Second)),
		BaseURL:    server.URL,
		EndPoint:   "/v1/",
	}

	_, err := c.Create(acc)

	assert.Equal(t, "can not create a new account something went wrong", err.Error())
}

func Test_Create_With_Corrupted_Error_Response(t *testing.T) {
	acc, _ := makeAccount()
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/", r.RequestURI)

		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{ "blabla": "blabla" }`))
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	c := client.AccountHttpClient{
		HttpClient: httpclient.NewClient(httpclient.WithHTTPTimeout(5 * time.Second)),
		BaseURL:    server.URL,
		EndPoint:   "/v1/",
	}

	_, err := c.Create(acc)

	assert.Equal(t, client.ErrorUnmarshallingCreateErrorResponse, err)
}

func Test_Fetch_With_TimeOut_Error(t *testing.T) {
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/accountId", r.RequestURI)
		time.Sleep(10 * time.Millisecond)
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	c := client.AccountHttpClient{
		HttpClient: httpclient.NewClient(httpclient.WithHTTPTimeout(1 * time.Millisecond)),
		BaseURL:    server.URL,
		EndPoint:   "/v1/",
	}

	_, err := c.Fetch("accountId")

	assert.Equal(t, fmt.Sprintf("error preparing fetch request to uri %s%s", server.URL, "/v1/accountId"), err.Error())
}

func Test_Fetch_With_Response_Error(t *testing.T) {
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/accountId", r.RequestURI)

		w.Write([]byte(`hello`))
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	c := client.AccountHttpClient{
		HttpClient: httpclient.NewClient(httpclient.WithHTTPTimeout(5 * time.Millisecond)),
		BaseURL:    server.URL,
		EndPoint:   "/v1/",
	}

	_, err := c.Fetch("accountId")

	assert.Equal(t, "error in unmarshalling Fetch json response", err.Error())
}

func Test_Valid_Fetch(t *testing.T) {
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/e5b590eb-2f58-4094-aa44-b52e875df159", r.RequestURI)

		w.Write([]byte(`{
    "data": {
        "attributes": {
            "account_classification": "Personal",
            "account_matching_opt_out": false,
            "account_number": "49R3IKD7",
            "alternative_bank_account_names": null,
            "bank_id": "123456",
            "bank_id_code": "GBDSC",
            "base_currency": "GBP",
            "bic": "NWBKGB22",
            "country": "GB",
            "joint_account": false
        },
        "created_on": "2020-03-22T11:21:48.084Z",
        "id": "e5b590eb-2f58-4094-aa44-b52e875df159",
        "modified_on": "2020-03-22T11:21:48.084Z",
        "organisation_id": "cd9fa404-aa64-4eb6-b155-5d156f3c383b",
        "type": "accounts",
        "version": 0
    },
    "links": {
        "self": "/v1/organisation/accounts/e5b590eb-2f58-4094-aa44-b52e875df159"
    }
}`))
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	c := client.AccountHttpClient{
		HttpClient: httpclient.NewClient(httpclient.WithHTTPTimeout(5 * time.Millisecond)),
		BaseURL:    server.URL,
		EndPoint:   "/v1/",
	}

	account, err := c.Fetch("e5b590eb-2f58-4094-aa44-b52e875df159")

	assert.Equal(t, nil, err)
	assert.Equal(t, "e5b590eb-2f58-4094-aa44-b52e875df159", account.ID)
	assert.Equal(t, "Personal", account.AccountClassification)
	assert.Equal(t, false, account.AccountMatchingOptOut)
	assert.Equal(t, "49R3IKD7", account.AccountNumber)
	assert.Equal(t, "123456", account.BankID)
	assert.Equal(t, "GBDSC", account.BankIDCode)
	assert.Equal(t, "GBP", account.BaseCurrency)
	assert.Equal(t, "NWBKGB22", account.Bic)
	assert.Equal(t, "GB", account.Country)
	assert.Equal(t, false, account.JointAccount)
	assert.Equal(t, "cd9fa404-aa64-4eb6-b155-5d156f3c383b", account.OrganisationID)
	assert.Equal(t, "accounts", account.Type)
}

func Test_Delete_With_TimeOut_Error(t *testing.T) {
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/v1/accountId?version=5", r.RequestURI)
		time.Sleep(10 * time.Millisecond)
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	c := client.AccountHttpClient{
		HttpClient: httpclient.NewClient(httpclient.WithHTTPTimeout(1 * time.Millisecond)),
		BaseURL:    server.URL,
		EndPoint:   "/v1/",
	}

	_, err := c.Delete("accountId", 5)

	assert.Equal(t, fmt.Sprintf("error in sedning Delete request to uri %s%s", server.URL, "/v1/accountId?version=5"), err.Error())
}

func Test_Delete_With_204_Response(t *testing.T) {
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/v1/accountId?version=0", r.RequestURI)

		w.WriteHeader(http.StatusNoContent)
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	c := client.AccountHttpClient{
		HttpClient: httpclient.NewClient(httpclient.WithHTTPTimeout(1 * time.Millisecond)),
		BaseURL:    server.URL,
		EndPoint:   "/v1/",
	}

	ok, err := c.Delete("accountId", 0)

	assert.Equal(t, true, ok)
	assert.Equal(t, nil, err)
}

func Test_Delete_With_404_Response(t *testing.T) {
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/v1/accountId?version=0", r.RequestURI)

		w.WriteHeader(http.StatusNotFound)
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	c := client.AccountHttpClient{
		HttpClient: httpclient.NewClient(httpclient.WithHTTPTimeout(1 * time.Millisecond)),
		BaseURL:    server.URL,
		EndPoint:   "/v1/",
	}

	ok, err := c.Delete("accountId", 0)

	assert.Equal(t, false, ok)
	assert.Equal(t, "error in deleting account accountId Specified resource does not exist", err.Error())
}

func Test_Delete_With_409_Response(t *testing.T) {
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/v1/accountId?version=0", r.RequestURI)

		w.WriteHeader(http.StatusConflict)
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	c := client.AccountHttpClient{
		HttpClient: httpclient.NewClient(httpclient.WithHTTPTimeout(1 * time.Millisecond)),
		BaseURL:    server.URL,
		EndPoint:   "/v1/",
	}

	ok, err := c.Delete("accountId", 0)

	assert.Equal(t, false, ok)
	assert.Equal(t, "error in deleting account accountId Specified version incorrect 0", err.Error())
}

func Test_Delete_With_208_Response(t *testing.T) {
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/v1/accountId?version=0", r.RequestURI)

		w.WriteHeader(http.StatusAlreadyReported)
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	c := client.AccountHttpClient{
		HttpClient: httpclient.NewClient(httpclient.WithHTTPTimeout(100 * time.Millisecond)),
		BaseURL:    server.URL,
		EndPoint:   "/v1/",
	}

	ok, err := c.Delete("accountId", 0)

	assert.Equal(t, false, ok)
	assert.Equal(t, "unknown response code 208", err.Error())
}

func Test_List_With_TimeOut_Error(t *testing.T) {
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1?page[number]=0&page[size]=100", r.RequestURI)
		time.Sleep(10 * time.Millisecond)
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	c := client.AccountHttpClient{
		HttpClient: httpclient.NewClient(httpclient.WithHTTPTimeout(1 * time.Millisecond)),
		BaseURL:    server.URL,
		EndPoint:   "/v1",
	}

	_, err := c.List(0, 0)

	assert.Equal(t, fmt.Sprintf("error in sending List request to uri %s%s", server.URL, "/v1?page[number]=0&page[size]=100"), err.Error())
}

func Test_List_With_Corrupted_Response(t *testing.T) {
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1?page[number]=0&page[size]=100", r.RequestURI)

		w.Write([]byte(`hello`))
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	c := client.AccountHttpClient{
		HttpClient: httpclient.NewClient(httpclient.WithHTTPTimeout(100 * time.Millisecond)),
		BaseURL:    server.URL,
		EndPoint:   "/v1",
	}

	_, err := c.List(0, 0)

	assert.Equal(t, "error in unmarshalling List json response", err.Error())
}

func Test_List_With_Invalid_Page_Number(t *testing.T) {
	// This is to mimic sending a negative number or a string as page number
	// the API will return server error with 500 status
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1?page[number]=0&page[size]=50", r.RequestURI)

		w.WriteHeader(http.StatusInternalServerError)
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	c := client.AccountHttpClient{
		HttpClient: httpclient.NewClient(httpclient.WithHTTPTimeout(100 * time.Millisecond)),
		BaseURL:    server.URL,
		EndPoint:   "/v1",
	}

	_, err := c.List(0, 50)

	assert.Equal(t, fmt.Sprintf("error in sending List request to uri %s%s", server.URL, "/v1?page[number]=0&page[size]=50"), err.Error())
}

// helper functions
func makeAccount() (*form3.Account, error) {
	accID := uuid.NewV4()
	orgID := uuid.NewV4()

	acc := &form3.Account{
		ID:             accID.String(),
		Type:           "accounts",
		OrganisationID: orgID.String(),
		AccountDetails: form3.AccountDetails{
			Country:               "GB",
			BaseCurrency:          "GBP",
			AccountNumber:         randomString(8),
			BankID:                "123456",
			BankIDCode:            "GBDSC",
			Bic:                   "NWBKGB22",
			AccountClassification: form3.AccountClassificationPersonal,
			JointAccount:          false,
			AccountMatchingOptOut: false,
			PrivateIdentification: form3.PrivateIdentification{
				Title:          "Mr",
				FirstName:      "Golang",
				LastName:       "Developer",
				BirthDate:      "2009-11-10",
				BirthCountry:   "US",
				DocumentNumber: "1234567",
				Address:        "[Robert Griesemer Rob Pike Ken Thompson]",
				City:           "Mountain View, California",
				Country:        "US",
			},
		},
	}

	return acc, nil
}

func randomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
