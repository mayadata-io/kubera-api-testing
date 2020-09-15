package authentication

import (
	"testing"
)

// TestCreateNewUser used to create a new kubera user
func TestCreateNewUser(t *testing.T) {
	var tests = map[string]struct {
		FirstName string
		LastName  string
		Email     string
		Password  string
		IsError   bool
	}{
		"No firstName": {
			LastName: "testF",
			Email:    "test@test.io",
			Password: "P@ssw0rd",
			IsError:  true,
		},
		"No lastName": {
			FirstName: "test",
			Email:     "nolastname@test.io",
			Password:  "P@ssw0rd",
			IsError:   false,
		},
		"No Email": {
			FirstName: "test",
			LastName:  "testF",
			Password:  "P@ssw0rd",
			IsError:   true,
		},
		"With No password": {
			FirstName: "test",
			LastName:  "testF",
			Email:     "test@test.io",
			IsError:   true,
		},
		"With firstName, lastName, email & password": {
			FirstName: "test",
			LastName:  "testF",
			Email:     "testing123@test.io", // TODO : Generate random string
			Password:  "P@ssw0rd",
		},
		"With invalid email": {
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
				UserCreation: UserCreation{
					FirstName: mock.FirstName,
					LastName:  mock.LastName,
					Credential: UserCredential{
						Email:    mock.Email,
						Password: mock.Password,
					},
				},
			})
			got, err := creator.Create()
			if !mock.IsError && err != nil {
				t.Fatalf("Expected no error got=%s: \nCreator=%s: \nToken=%s", err, creator, got)
			}
			if mock.IsError && err == nil {
				t.Fatalf("Expected error got none:  \nCreator=%s: \nToken=%s", creator, got)
			}
		})
	}
}
