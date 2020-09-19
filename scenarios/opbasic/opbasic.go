package opbasic

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	auth "github.com/mayadata-io/kubera-api-testing/kubera/authentication"
	common "github.com/mayadata-io/kubera-api-testing/kubera/common"
)

// template required to install openebs
type OpenebsTemplate struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	BaseType string `json:"baseType"`
	State    string `json:"state"`
	Kind     string `json:"kind"`
	Version  string `json:"version"`
}

// New instance of KuberaFetching
var newKuberaFetching common.KuberaFetching

// New instance of KuberaID
var newKuberaID common.KuberaID

// FetchOpenebsTemplate will return an array of templates
// available for installing OpenEBS via Kubera
func FetchOpenebsTemplate(newKuberaFetching common.KuberaFetching, newKuberaID common.KuberaID) ([]OpenebsTemplate, error) {
	var newTemplate []OpenebsTemplate
	client := resty.New()
	kErr := auth.KuberaError{}
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Cookie",
			fmt.Sprintf(
				"token=%s;authProvider=localAuthConfig",
				newKuberaFetching.Token,
			),
		).
		SetError(&kErr).
		Get(
			fmt.Sprintf(
				"%s/v3/groups/%s/openebstemplates",
				newKuberaFetching.Hostname,
				newKuberaID.GroupID,
			),
		)

	if err != nil || kErr.Message != "" || !resp.IsSuccess() {
		return newTemplate, &auth.APIError{
			Err:        err,
			Message:    kErr.Message,
			StatusCode: resp.StatusCode(),
			Status:     resp.Status(),
		}
	}

	json.Unmarshal([]byte(resp.Body()), &newTemplate)
	return newTemplate, nil

}
