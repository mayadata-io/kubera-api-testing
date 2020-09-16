package authentication

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/mayadata-io/kubera-api-testing/kubera/connection"
)

// UserCreation contains structure of the input data to create NewUser
type UserCreation struct {
	FirstName  string         `json:"firstName"`
	LastName   string         `json:"lastName"`
	Credential UserCredential `json:"password"`
}

// UserCredential contains structure of the user secrets input
type UserCredential struct {
	Email    string `json:"publicValue"`
	Password string `json:"secretValue"`
}

// String print UserCreation in pretty format
func (t UserCreation) String() string {
	if t.Credential.Password != "" {
		t.Credential.Password = "XXXX"
	}
	raw, err := json.MarshalIndent(
		t,
		" ",
		".",
	)
	if err != nil {
		panic(err)
	}
	return string(raw)
}

// UserCreationConfig defines the structure of user input fields
type UserCreationConfig struct {
	UserCreation
}

// NewUserCreator returns a new instance of UserCreation
func NewUserCreator(config UserCreationConfig) *UserCreation {
	return &UserCreation{
		FirstName: config.FirstName,
		LastName:  config.LastName,
		Credential: UserCredential{
			Email:    config.Credential.Email,
			Password: config.Credential.Password,
		},
	}
}

// Create will creates new kubera user in kubera by making post request
func (t *UserCreation) Create() (Token, error) {
	client := resty.New()
	url := fmt.Sprintf("%sv3/localauthconfig/", connection.Get().HostName)
	response := Token{}
	kErr := KuberaError{}
	body, err := json.Marshal(t)
	if err != nil {
		return Token{}, &APIError{
			Err: err,
		}
	}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		//SetError invoke if response status code >= 399 and content type JSON
		SetError(&kErr).
		//SetResult invoke if response status code is between 200 and 299
		SetResult(&response).
		Post(url)

	if err != nil || kErr.Message != "" || !resp.IsSuccess() {
		return Token{}, &APIError{
			Err:        err,
			Message:    kErr.Message,
			StatusCode: resp.StatusCode(),
			Status:     resp.Status(),
		}
	}

	return response, nil
}
