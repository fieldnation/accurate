package accurate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Candidate represents a candidate to be screened.
type Candidate struct {
	Resource     string     `json:"resource"`
	ID           string     `json:"id"`
	Created      *time.Time `json:"created"`
	Updated      *time.Time `json:"updated"`
	Revision     string     `json:"revision"`
	FirstName    string     `json:"firstName"`
	LastName     string     `json:"lastName"`
	MiddleName   string     `json:"middleName"`
	DateOfBirth  string     `json:"dateOfBirth"`
	SSN          string     `json:"ssn"`
	Email        string     `json:"email"`
	Phone        string     `json:"phone"`
	Address      string     `json:"address"`
	City         string     `json:"city"`
	Region       string     `json:"region"`
	Country      string     `json:"country"`
	PostalCode   string     `json:"postalCode"`
	GovernmentID struct {
		Country string `json:"country"`
		Type    string `json:"type"`
		Number  string `json:"number"`
	} `json:"governmentId"`
	Aliases      []string `json:"aliases"`
	Educations   []string `json:"educations"`
	PrevEmployed string   `json:"prevEmployed"`
	Employments  []string `json:"employments"`
	Licenses     []string `json:"licenses"`
	Convicted    string   `json:"convicted"`
	Convictions  []string `json:"convictions"`
	References   []string `json:"references"`
}

// Create sends a request to create a new Candidate.
func (c Candidate) Create() error {

	// required sanity check
	if c.FirstName == "" || c.LastName == "" || c.Email == "" {
		return fmt.Errorf(
			"missing one of these required fields firstname: %q, lastname: %q, email: %q",
			c.FirstName, c.LastName, c.Email,
		)
	}

	// marshal the candidate to buffered bytes representing JSON
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(b)

	// create a new request
	url := URL.String() + candidate
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}

	// set API key for authentication and authorization
	req.SetBasicAuth(clientID, clientSecret)

	// send the HTTP request with the default Go client
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// check the HTTP response status code is 201
	if resp.StatusCode != http.StatusCreated {

		// read the HTTP response body
		defer resp.Body.Close()
		b, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		// return the HTTP response body as an error
		return errors.New(string(b))
	}

	return nil
}
