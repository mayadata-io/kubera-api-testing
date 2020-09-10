package authentication

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	conn "github.com/mayadata-io/kubera-api-testing/kubera/connection"
)

// UserCreation contains structure of the input data to create NewUser
type UserCreation struct {
	FirstName   string   `json:"firstName"`
	LastName    string   `json:"lastName"`
	Credentials UserCred `json:"password"`
}

// UserCred contains structure of the user secrets input
type UserCred struct {
	Email    string `json:"publicValue"`
	Password string `json:"secretValue"`
}

// String print UserCreation in pretty format
func (t UserCreation) String() string {
	if t.Credentials.Password != "" {
		t.Credentials.Password = "XXXX"
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
		Credentials: UserCred{
			Email:    config.Credentials.Email,
			Password: config.Credentials.Password,
		},
	}
}

// Create is  a method , It creates new user in kubera by making -
//  post reqest to API server
func (t *UserCreation) Create() (Token, error) {
	client := resty.New()
	url := conn.Get().HostName + "v3/localauthconfig/"
	token := Token{}
	kErr := KuberaError{}
	credJSON, err := json.Marshal(t)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(credJSON).
		//SetError method automatic unmarshalling if response status code >= 399
		// and content type either JSON or XML.
		SetError(&kErr).
		//SetResult automatic unmarshalling for the request,
		// if response status code is between 200 and 299
		SetResult(&token).
		Post(url)

	if err != nil || kErr.Message != "" || !resp.IsSuccess() {
		return Token{}, &APIError{
			Err:        err,
			Message:    kErr.Message,
			StatusCode: resp.StatusCode(),
			Status:     resp.Status(),
		}
	}

	return token, nil
}
