package rest

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// SuccessRequest for ,if request success then to get token
type SuccessRequest struct {
	Token string `json:"jwt"`
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

type PostResponse struct {
	Resp       *resty.Response
	Token      string
	Message    string
	Error      error
	StatusCode int
	// Body    []byte
}

type HttpPost interface {
	getToken() PostResponse
}

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
	}
	return &sendResp

}
