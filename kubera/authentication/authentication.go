package authentication

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	conn "github.com/mayadata-io/kubera-api-testing/kubera/connection"
)

// Token struct of JWT token and AccountID
type Token struct {
	JWT       string `json:"jwt"`
	AccountID string `json:"accountId"`
}

func (t Token) String() string {
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

// APIError struct of error, error message, status code and status
type APIError struct {
	Err        error  `json:"err"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
}

func (e *APIError) Error() string {
	raw, err := json.MarshalIndent(
		e,
		" ",
		".",
	)
	if err != nil {
		panic(err)
	}
	return string(raw)
}

// KuberaError represent error message
type KuberaError struct {
	Message string `json:"message"`
}

// TokenFetching struct of token fetching connection
type TokenFetching struct {
	conn.Connection `json:"connection"`
}

func (t TokenFetching) String() string {
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

// TokenFetchingConfig is struct of connection
type TokenFetchingConfig struct {
	Connection conn.Connection
}

// NewTokenFetcher returns a new instance of TokenFetching
func NewTokenFetcher(config TokenFetchingConfig) *TokenFetching {
	return &TokenFetching{
		Connection: config.Connection,
	}
}

// Fetch returns authenticated token
func (t *TokenFetching) Fetch() (Token, error) {
	client := resty.New()
	token := Token{}
	kErr := KuberaError{}
	bodyData := map[string]interface{}{
		"code":         fmt.Sprintf("%s:%s", t.UserName, t.Password),
		"authProvider": "localAuthConfig",
	}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(bodyData).
		//SetError method automatic unmarshalling if response status code >= 399
		// and content type either JSON or XML.
		SetError(&kErr).
		//SetResult automatic unmarshalling for the request,
		// if response status code is between 200 and 299
		SetResult(&token).
		Post(
			fmt.Sprintf(
				"%s/v3/token",
				t.HostName,
			),
		)

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
