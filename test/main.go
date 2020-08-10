package main

import (
	"flag"
	"fmt"
)

// Cred structure of project credentials
type Cred struct {
	UserName string
	Password string
	Host     string
}

// Token is structure of access token
type Token struct {
	Token     string
	AccountID string
}

func main() {

	userName := flag.String("username", "bhaskar@JB007.com", "a string")
	password := flag.String("password", "bhaskar@JB007.com", "a string")
	host := flag.String("host", "http://15.207.111.248", "a string")
	flag.Parse()
	userCred := Cred{
		UserName: *userName,
		Password: *password,
		Host:     *host,
	}

	//to aceess Kubera API need private token
	token, err := getToken(userCred)
	if err != nil {
		fmt.Println("error:: ", err)
	}
	// sample function to get userAccount Details
	getAccountDetails(userCred.Host, token.Token, token.AccountID)

}
