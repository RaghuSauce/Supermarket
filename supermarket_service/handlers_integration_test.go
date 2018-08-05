package supermarket_service

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"flag"
)

//Routes
const (
	//Application Endpoints
	baseURL   = "http://localhost:8081"
	IndexURL  = baseURL + "/"
	FetchURL  = IndexURL + "fetch"
	AddURL    = IndexURL + "add"
	RemoveURL = IndexURL + "remove/code="
)

//Responses
const (
	IndexResponse         = "Supermarket-API:"
	FetchResponse         = `[{"producecode":"A12T-4GH7-QPL9-3N4M","name":"Lettuce","unitprice":"3.46"},{"producecode":"E5T6-9UI3-TH15-QR88","name":"Peach","unitprice":"2.99"},{"producecode":"YRT6-72AS-K736-L4AR","name":"Green Pepper","unitprice":"0.79"},{"producecode":"TQ4C-VV6T-75ZX-1RMR","name":"Gala Apple","unitprice":"3.59"}]`
	AddProduceItemPayload = `{"producecode":"A12T-4GH7-QPL9-3N4N","name":"Andy Sandberg","unitprice":"30.46"}`
	ValidResponse         = `Success`
	removeCodeURL         = RemoveURL + `A12T-4GH7-QPL9-3N4N`
)

var(
	runIntegration = flag.Bool("integration",false,"Flag for running integration tests")
)

func TestIndex(t *testing.T) {
	if ! *runIntegration{
		t.Skip("Skipping Integration Tests")
	}

	if response, err := http.Get(IndexURL); err != nil {
		t.Errorf("Error On Index GET\n%s", err)
	} else {
		defer response.Body.Close()
		if responseContent, err := ioutil.ReadAll(response.Body); err != nil {
			t.Errorf("Error Parsing the response body\n%s", err)
		} else if !strings.Contains(string(responseContent), IndexResponse) {
			t.Errorf("Expeted content to contain\n%s\nfound\n%s", IndexResponse, string(responseContent))
		}
	}
}

//Test the Fetch Method to get all Items in the database
func TestFetchProduceList(t *testing.T) {
	if ! *runIntegration{
		t.Skip("Skipping Integration Tests")
	}

	if response, err := http.Get(FetchURL); err != nil { //Check for error on HTTP GET
		t.Errorf("Error on %s \n Error:%s", FetchURL, err)
	} else {
		defer response.Body.Close()
		if responseContent, err := ioutil.ReadAll(response.Body); err != nil { //Check for error On reading the Response body
			t.Errorf("Error while parsing the reponse body \n %s", err)
		} else {
			//Check to see of if the result matches the initial set of the database
			//Have to remove trailing whitespaces before comparison
			if strings.TrimRight(string(responseContent), "\n") != FetchResponse {
				t.Errorf("Response Mismatch for fetch endpoint\n"+
					"Expected:%s\nGot:%s",
					FetchResponse, responseContent)

			}
		}
	}
}

//Test Adding to the Database
func TestAddProduceItem(t *testing.T) {
	if ! *runIntegration{
		t.Skip("Skipping Integration Tests")
	}

	postContent := []byte(AddProduceItemPayload)

	if response, err := http.Post(AddURL, "application/json", bytes.NewBuffer(postContent)); err != nil {
		t.Errorf("Error on POST to %s\n%s", AddURL, err)
	} else {
		defer response.Body.Close()
		if responseContent, err := ioutil.ReadAll(response.Body); err != nil {
			t.Errorf("Error Reading the contents of the response \n%s", err)
		} else {
			if strings.TrimRight(string(responseContent), "\n") != ValidResponse {
				t.Errorf("Response Mismatch for add endpoint\n"+
					"Expected:%s\nGot:%s",
					ValidResponse, responseContent)
			}
		}
	}
}

func TestRemoveProduceItem(t *testing.T) {
	if ! *runIntegration{
		t.Skip("Skipping Integration Tests")
	}

	if request, err := http.NewRequest("DELETE", removeCodeURL, nil); err != nil {
		t.Errorf("error creating request for DELETE \n%s", err)
	} else {
		if response, err := http.DefaultClient.Do(request); err != nil {
			t.Errorf("error sending request for DELETE\n%s", err)
		} else {
			defer response.Body.Close()

			if responseContent, err := ioutil.ReadAll(response.Body); err != nil {
				t.Errorf("Error parsing content from DELETE endpoint\n%s", err)
			} else if strings.TrimRight(string(responseContent), "\n") != ValidResponse {
				t.Errorf("Response did not match Expected Result for DELETE\nExpected:%s\nGot:%s", ValidResponse, responseContent)
			}
		}
	}
}
