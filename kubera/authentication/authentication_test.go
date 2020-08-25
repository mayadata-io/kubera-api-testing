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
