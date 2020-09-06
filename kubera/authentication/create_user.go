package authentication

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	conn "github.com/mayadata-io/kubera-api-testing/kubera/connection"
)

type User struct {
	FirstName   string   `json:"firstName"`
	LastName    string   `json:"lastName"`
	Credentials UserCred `json:"password"`
}

type UserCred struct {
	Email    string `json:"publicValue"`
	Password string `json:"secretValue"`
}

// UserCreation Contains the user form structure
type UserCreation struct {
	conn.Account `json:"connection"`
}

// String print UserCreation in pretty format
func (t UserCreation) String() string {
	if t.Password != "" {
		t.Password = "XXXX"
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

type UserCreationConfig struct {
	UserForm conn.Account
}

// NewUserCreator returns a new instance of UserCreation
func NewUserCreator(config UserCreationConfig) *UserCreation {
	return &UserCreation{
		Account: config.UserForm,
	}
}

// Create returns authenticated token
func (u *UserCreation) Create() (Token, error) {
	client := resty.New()
	url := u.Host + "v3/localauthconfig/"
	form := User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Credentials: UserCred{
			Email:    u.Email,
			Password: u.Password,
		},
	}
	token := Token{}
	kErr := KuberaError{}
	credJSON, err := json.Marshal(form)
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
