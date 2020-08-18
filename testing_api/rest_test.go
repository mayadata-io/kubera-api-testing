package rest

// Import resty into your code and refer it as `resty`.
import (
	"fmt"
	"testing"
)

//TestGetUserDetails is to get the Details of User
// API : TODO add structure of api
func TestGetUserDetails(t *testing.T) {
	//usr contains the details which is useful to get access token
	// TODO need to get these values from env or flags
	usr := UserData{
		Host:     "http://15.207.111.248/",
		UserName: "test1@e2e.io",
		Password: "Test@e2e.io",
	}
	// getToken() is a method to get token
	res := usr.getToken()

	fmt.Println(res.Error)
	if res.StatusCode != 201 {
		t.Errorf(" Error : %s \n %s ", res.Message, res.Error)
	}
	t.Logf(res.Token) //by Using token need to access API of userPage
	getRes := res.getRequest(usr.Host + "v3/account/")
	if getRes.StatusCode != 200 {
		t.Errorf("Error : %s \n Expected StatusCode %d but, got %d", getRes.Error, 200, getRes.StatusCode)
	} else {
		t.Logf(" \n \n - - -  Login successful - - -")
	}
}

// //TestSignupUser  for create a new user
// func TestSignupUser(t *testing.T) {

// }
