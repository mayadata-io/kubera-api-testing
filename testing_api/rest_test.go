package rest

// Import resty into your code and refer it as `resty`.
import (
	"fmt"
	"testing"
)

//TestGetUserDetails is to get the Details of User
// API : TODO add structure of api
func TestGetUserDetails(t *testing.T) {
	//U contains the details which is useful to get access token
	// TODO need to get these values from env or flags
	u := UserData{
		Host:     "http://15.207.111.248/",
		UserName: "test1@e2e.io",
		Password: "Test@e2e.io",
	}
	// getToken() is a method to get token
	res := u.getToken()

	fmt.Println(res.Error)
	if res.StatusCode != 201 {
		t.Errorf(" Error : %s \n %s ", res.Message, res.Error)
	}

	t.Logf(res.Token) // Using token need to access API of userPage

	// Need to write a Get method  request to access kubera useraccess api page
}
