package authentication

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	conn "github.com/mayadata-io/kubera-api-testing/kubera/connection"
)

type NewUserForm struct {
	Fname       string `json:"firstName"`
	Lname       string `json:"lastName"`
	Credentials Cred   `json:"password"`
}

type Cred struct {
	Email    string `json:"publicValue"`
	Password string `json:"secretValue"`
}

// UserCreate Contains the user form structure
type UserCreate struct {
	conn.UserForm `json:"connection"`
}

// String print UserCreate in pretty format
func (t UserCreate) String() string {
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

type UserCreateConfig struct {
	UserForm conn.UserForm
}

// NewUserCreator returns a new instance of UserCreate
func NewUserCreator(config UserCreateConfig) *UserCreate {
	return &UserCreate{
		UserForm: config.UserForm,
	}
}

// Create returns authenticated token
func (d *UserCreate) Create() (Token, error) {
	client := resty.New()
	url := d.Host + "v3/localauthconfig/"
	form := NewUserForm{
		Fname: d.FirstName,
		Lname: d.LastName,
		Credentials: Cred{
			Email:    d.Email,
			Password: d.Password,
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
