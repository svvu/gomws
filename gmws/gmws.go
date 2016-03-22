package gmws

import (
	"fmt"
	"os"

	"github.com/kr/pretty"
	"github.com/svvu/gomws/mwsHttps"
)

// MwsConfig is configuraton to create the gomws base.
// AccessKey and SecretKey are optional, bette to set them in evn variables.
type MwsConfig struct {
	SellerId  string
	AuthToken string
	Region    string
	AccessKey string
	SecretKey string
}

// MwsClient the interface for API clients.
type MwsClient interface {
	Version() string
	Name() string
	NewClient(config MwsConfig) (MwsClient, error)
	GetServiceStatus() (mwsHttps.Response, error)
}

const (
	envAccessKey = "AWS_ACCESS_KEY"
	envSecretKey = "AWS_SECRET_KEY"
)

// Credential the credential to access the API.
type Credential struct {
	AccessKey string
	SecretKey string
}

// GetCredential get the credential from evn variables.
func GetCredential() Credential {
	credential := Credential{}
	credential.AccessKey = os.Getenv(envAccessKey)
	credential.SecretKey = os.Getenv(envSecretKey)

	return credential
}

// Inspect print out the value in a user friendly way.
func Inspect(value interface{}) {
	fmt.Printf("%# v", pretty.Formatter(value))
}
