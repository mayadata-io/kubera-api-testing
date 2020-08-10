package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/yalp/jsonpath"
)

func getToken(usrDetail Cred) (*Token, error) {
	url := usrDetail.Host + "/v3/token"
	method := "POST"

	bodyContent := fmt.Sprintf("{\"code\": \"%s:%s\", \"authProvider\": \"localAuthConfig\"}", usrDetail.UserName, usrDetail.Password)
	Body := strings.NewReader(bodyContent)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, Body)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "PL=DirectorOnline")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	if res.StatusCode != 201 {
		fmt.Println("Credentilas provided is not correct ", res.Status)
		// return "", "", nil
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	var data interface{}
	json.Unmarshal(body, &data)

	TOKEN, err := jsonpath.Read(data, "$.jwt")
	if err != nil {
		panic(err)
	}
	accountID, err := jsonpath.Read(data, "$.accountId")
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprint(accountID))
	resp := Token{
		Token:     fmt.Sprint(TOKEN),
		AccountID: fmt.Sprint(accountID),
	}
	return &resp, nil
}
func getAccountDetails(host, token, accountID string) {

	url := host + "/v3/account/" + accountID
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Cookie", "token="+token+"; authProvider=localAuthConfig")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)

}
