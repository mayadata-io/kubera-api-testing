package authentication

import (
	"testing"

	conn "github.com/mayadata-io/kubera-api-testing/kubera/connection"
)

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
			creator := NewUserCreator(UserCreationConfig{
				UserForm: conn.Account{
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
