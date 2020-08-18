package rest

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// SuccessRequest for ,if request success then to get token
type SuccessRequest struct {
	Token     string `json:"jwt"`
	AccountID string `json:"accountId"`
}

// FailedRequest fail return error
type FailedRequest struct {
	Msg string `json:"message"`
}

// UserData contains the details of credentials
type UserData struct {
	UserName string
	Password string
	Host     string
}

// PostResponse for defining the response of postmethod
type PostResponse struct {
	Resp       *resty.Response
	Token      string
	Message    string
	Error      error
	StatusCode int
	AccountID  string
	// Body    []byte
}

// type HttpPost interface {
// 	getToken() PostResponse
// }

func (u UserData) getToken() *PostResponse {
	client := resty.New()
	url := u.Host + "v3/token"
	token := SuccessRequest{}
	message := FailedRequest{}
	bodyData := fmt.Sprintf(`{"code": "%s:%s", "authProvider": "localAuthConfig"}`, u.UserName, u.Password)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Cookie", "PL=DirectorOnline").
		SetBody(bodyData).
		SetError(&message).
		SetResult(&token). // or SetResult(AuthSuccess{}).
		Post(url)

	sendResp := PostResponse{
		Error:   err,
		Token:   token.Token,
		Message: message.Msg,
		Resp:    resp,
		// Body:    resp.Body(),
		StatusCode: resp.StatusCode(),
		AccountID:  token.AccountID,
	}
	return &sendResp

}

// GetReqResponse structure of GET response
type GetReqResponse struct {
	Error      error
	Resp       *resty.Response
	StatusCode int
}

// getRequest func is to make GET method requests
func (p PostResponse) getRequest(api string) *GetReqResponse {
	fmt.Println("api-----", api)
	client := resty.New()

	url := api + p.AccountID
	header := fmt.Sprintf("token= %s; authProvider=localAuthConfig", p.Token)

	resp, err := client.R().
		SetHeader("Cookie", header).
		Get(url)

	printOutput(resp, err) // To Check the respose
	return &GetReqResponse{
		Error:      err,
		Resp:       resp,
		StatusCode: resp.StatusCode(),
	}
}
func printOutput(resp *resty.Response, err error) {
	fmt.Println(resp, err)
}
