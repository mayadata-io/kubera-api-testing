package connection

import (
	"encoding/json"
	"os"
	"sync"
)

const (
	kuberaHost = "KUBERA_HOST" //KUBERA_HOST is Env variable contains value of Kubera's Domain/IP
	kuberaUser = "KUBERA_USER" // Kubera Username
	kuberaPass = "KUBERA_PASS" //Kubera password
)

type Account struct {
	Host      string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

// Connection defines the structure of userCerdentials
type Connection struct {
	HostName string
	UserName string
	Password string
}

func (c Connection) String() string {
	if c.Password != "" {
		c.Password = "XXXX"
	}
	raw, err := json.MarshalIndent(
		c,
		" ",
		".",
	)
	if err != nil {
		panic(err)
	}
	return string(raw)
}

var connection Connection

// Get return the credentials of kubera user ,pass and host
func Get() Connection {

	if connection.HostName != "" {
		return connection
	}
	var once sync.Once

	once.Do(func() {
		connection = Connection{
			HostName: os.Getenv(kuberaHost),
			UserName: os.Getenv(kuberaUser),
			Password: os.Getenv(kuberaPass),
		}
	})
	return connection

}
