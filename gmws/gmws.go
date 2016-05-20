package gmws

import (
	"fmt"
	"os"

	"github.com/kr/pretty"
	"github.com/svvu/gomws/mwsHttps"
)

const (
	envAccessKey = "AWS_ACCESS_KEY"
	envSecretKey = "AWS_SECRET_KEY"
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

// HasErrors will check whether or not the xml node tree has any erorr node.
// If it contains errors, true will be returned.
func HasErrors(xmlNode *XMLNode) bool {
	errorNodes := xmlNode.FindByKey("Error")
	if len(errorNodes) > 0 {
		return true
	}
	return false
}

// GetErrors will return an array of Error struct from the xml node tree.
func GetErrors(xmlNode *XMLNode) ([]Error, error) {
	errorNodes := xmlNode.FindByKey("Error")

	errors := []Error{}
	for _, en := range errorNodes {
		error := Error{}
		err := en.ToStruct(&error)
		if err != nil {
			return errors, err
		}
		errors = append(errors, error)
	}
	return errors, nil
}
