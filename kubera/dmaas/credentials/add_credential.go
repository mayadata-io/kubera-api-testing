package credentials

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/mayadata-io/kubera-api-testing/kubera/authentication"
	"github.com/mayadata-io/kubera-api-testing/kubera/connection"
	"github.com/tidwall/gjson"
)

// CloudCredential contains all parameters which should be passed to add cloud credential
type CloudCredential struct {
	Name       string   `json:"name"`
	ProjectID  string   `json:"projectId"`
	ProviderID string   `json:"providerId"`
	Cred       CredType `json:"credential"`
}

// CredType is cloud provider access id and secret key
type CredType struct {
	AccessID  string `json:"access_id"`
	SecretKey string `json:"secret_key"`
}

// CredentialAdding represent the fields required to add a credential
type CredentialAdding struct {
	CloudName string
	CredName  string
	Username  string
	Password  string

	jwtToken   string
	hostName   string
	groupID    string
	projectID  string
	providerID string
}

// NewCredentialsAdder return fields value required to add a credential
func NewCredentialsAdder(cred CredentialAdding) *CredentialAdding {
	return &CredentialAdding{
		CloudName: cred.CloudName,
		CredName:  cred.CredName,
		Username:  cred.Username,
		Password:  cred.Password,
	}
}

// SetKuberaDetails sets JWT token and Kubera url
func (a *CredentialAdding) SetKuberaDetails() error {
	conn := connection.Get()
	fetcher := authentication.NewTokenFetcher(
		authentication.TokenFetchingConfig{
			Connection: conn,
		})
	tkn, err := fetcher.Fetch()
	if err != nil {
		return err
	}
	// setting JWT token and Kubera hostname
	a.jwtToken = tkn.JWT
	a.hostName = connection.GetHostName()
	return nil
}

// FetchGroupProjectID method sets group and project ID
func (a *CredentialAdding) FetchGroupProjectID() error {
	client := resty.New()
	kErr := authentication.KuberaError{}
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Cookie",
			fmt.Sprintf(
				"token=%s;authProvider=localAuthConfig",
				a.jwtToken,
			),
		).
		SetError(&kErr).
		Get(
			fmt.Sprintf(
				"%s/v3/groups",
				a.hostName,
			),
		)

	if err != nil || kErr.Message != "" || !resp.IsSuccess() {
		return &authentication.APIError{
			Err:        err,
			Message:    kErr.Message,
			StatusCode: resp.StatusCode(),
			Status:     resp.Status(),
		}
	}
	// groupIDPath denotes the path of group id by using ProjectAccount as filter.
	groupIDPath := fmt.Sprintf("data.#(name==%s).id", "ProjectAccount")
	// projectIDPath denotes the path of project id by using ProjectAccount as filter.
	projectIDPath := fmt.Sprintf("data.#(name==%s).projectId", "ProjectAccount")

	// set group ID and project ID
	a.groupID = gjson.GetBytes(resp.Body(), groupIDPath).String()
	a.projectID = gjson.GetBytes(resp.Body(), projectIDPath).String()

	return nil
}

// FetchBackupProviderID method sets backup provider ID
func (a *CredentialAdding) FetchBackupProviderID() error {
	client := resty.New()
	kErr := authentication.KuberaError{}
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Cookie",
			fmt.Sprintf(
				"token=%s;authProvider=localAuthConfig",
				a.jwtToken,
			),
		).
		SetError(&kErr).
		Get(
			fmt.Sprintf("%s/v3/backupproviders",
				a.hostName,
			),
		)
	if err != nil || kErr.Message != "" || !resp.IsSuccess() {
		return &authentication.APIError{
			Err:        err,
			Message:    kErr.Message,
			StatusCode: resp.StatusCode(),
			Status:     resp.Status(),
		}
	}
	// path of backup provider id by using name as filter.
	providerIDPath := fmt.Sprintf("data.#(name==%s).id", a.CloudName)
	// set backup provider id
	a.providerID = gjson.GetBytes(resp.Body(), providerIDPath).String()
	return nil
}

// AddCredential method will add cloud provider credential
func (a *CredentialAdding) AddCredential() error {
	client := resty.New()
	kErr := authentication.KuberaError{}
	// body of POST request
	body := CloudCredential{
		Name:       a.CredName,
		ProjectID:  a.projectID,
		ProviderID: a.providerID,
		Cred: CredType{
			AccessID:  a.Username,
			SecretKey: a.Password,
		},
	}
	// body of POST request in JSON format
	bodyRaw, err := json.Marshal(body)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(bodyRaw).
		SetHeader("Cookie",
			fmt.Sprintf(
				"token=%s;authProvider=localAuthConfig",
				a.jwtToken,
			),
		).
		//SetError method automatic unmarshalling if response status code >= 399
		// and content type either JSON or XML.
		SetError(&kErr).
		Post(
			fmt.Sprintf(
				"%s/v3/groups/%s/providercredentials",
				connection.GetHostName(),
				a.groupID,
			),
		)

	if err != nil || kErr.Message != "" || !resp.IsSuccess() {
		return &authentication.APIError{
			Err:        err,
			Message:    kErr.Message,
			StatusCode: resp.StatusCode(),
			Status:     resp.Status(),
		}
	}
	return nil
}

// AddCloudCredential method add backup providers credential
func (a *CredentialAdding) AddCloudCredential() error {
	fns := []func() error{
		a.SetKuberaDetails,
		a.FetchGroupProjectID,
		a.FetchBackupProviderID,
		a.AddCredential,
	}
	for _, fn := range fns {
		err := fn()
		if err != nil {
			return err
		}
	}
	return nil
}
