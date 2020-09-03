package authentication

import (
	"testing"

	conn "github.com/mayadata-io/kubera-api-testing/kubera/connection"
)

func TestFetchToken(t *testing.T) {
	connection := conn.Get()
	var tests = map[string]struct {
		Hostname string
		Username string
		Password string
		IsError  bool
	}{
		"No host": {
			Username: connection.UserName,
			Password: connection.Password,
			IsError:  true,
		},
		"No username": {
			Hostname: connection.HostName,
			Password: connection.Password,
			IsError:  true,
		},
		"No password": {
			Hostname: connection.HostName,
			Username: connection.UserName,
			IsError:  true,
		},
		"With host, username & password": {
			Hostname: connection.HostName,
			Username: connection.UserName,
			Password: connection.Password,
		},
		"With host, username & invalid password": {
			Hostname: connection.HostName,
			Username: connection.UserName,
			Password: "invalid",
			IsError:  true,
		},
		"With invalid host, username & password": {
			Hostname: "invalid",
			Username: connection.UserName,
			Password: connection.Password,
			IsError:  true,
		},
		"With host, invalid username & password": {
			Hostname: connection.HostName,
			Username: "invalid",
			Password: connection.Password,
			IsError:  true,
		},
	}
	for name, mock := range tests {
		name := name // Pin it
		mock := mock // Pin it
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			fetcher := NewTokenFetcher(TokenFetchingConfig{
				Connection: conn.Connection{
					HostName: mock.Hostname,
					UserName: mock.Username,
					Password: mock.Password,
				},
			})
			tkn, err := fetcher.Fetch()
			if !mock.IsError && err != nil {
				t.Fatalf("Expected no error got=%s: \nFetcher=%s: \nToken=%s", err, fetcher, tkn)
			}
			if mock.IsError && err == nil {
				t.Fatalf("Expected error got none:  \nFetcher=%s: \nToken=%s", fetcher, tkn)
			}
		})
	}
}

// TestCreateNewUser is to create a new user
func TestCreateNewUser(t *testing.T) {
	connection := conn.Get()
	var tests = map[string]struct {
		HostName  string
		FirstName string
		LastName  string
		Email     string
		Password  string
		IsError   bool
	}{
		"No host": {
			FirstName: "test",
			LastName:  "testF",
			Email:     "test@test.io",
			Password:  "P@ssw0rd",
			IsError:   true,
		},
		"No firstName": {
			HostName: connection.HostName,
			LastName: "testF",
			Email:    "test@test.io",
			Password: "P@ssw0rd",
			IsError:  true,
		},
		"No lastName": {
			HostName:  connection.HostName,
			FirstName: "test",
			Email:     "test@test.io",
			Password:  "P@ssw0rd",
		},
		"No Email": {
			HostName:  connection.HostName,
			FirstName: "test",
			LastName:  "testF",
			Password:  "P@ssw0rd",
			IsError:   true,
		},
		"With No password": {
			HostName:  connection.HostName,
			FirstName: "test",
			LastName:  "testF",
			Email:     "test@test.io",
			IsError:   true,
		},
		"With host, firstName, lastName, email & password": {
			HostName:  connection.HostName,
			FirstName: "test",
			LastName:  "testF",
			Email:     "test@test.io", // TODO : Generate random string
			Password:  "P@ssw0rd",
		},
		"With invalid email": {
			HostName:  connection.HostName,
			FirstName: "test",
			LastName:  "testF",
			Email:     "testtest.io",
			Password:  "P@ssw0rd",
			IsError:   true,
		},
	}

	for name, mock := range tests {
		name := name
		mock := mock
		t.Run(name, func(t *testing.T) {
			creator := NewUserCreator(UserCreateConfig{
				UserForm: conn.UserForm{
					Host:      mock.HostName,
					FirstName: mock.FirstName,
					LastName:  mock.LastName,
					Email:     mock.Email,
					Password:  mock.Password,
				},
			})
			createUser, err := creator.Create()
			if !mock.IsError && err != nil {
				t.Fatalf("Expected no error got=%s: \nCreator=%s: \nToken=%s", err, creator, createUser)
			}
			if mock.IsError && err == nil {
				t.Fatalf("Expected error got none:  \nCreator=%s: \nToken=%s", creator, createUser)
			}
		})
	}
}
