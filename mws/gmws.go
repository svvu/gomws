package mws

import (
	"fmt"
	"os"
	"time"

	"github.com/kr/pretty"
)

const (
	envAccessKey = "AWS_ACCESS_KEY"
	envSecretKey = "AWS_SECRET_KEY"

	iso8061Format = time.RFC3339 // "2006-01-02T15:04:05Z07:00"
)

// Current timestamp in iso8061 format.
var now = func() string {
	return time.Now().UTC().Format(iso8061Format)
}

// APIClient the interface for API clients.
type APIClient interface {
	Version() string
	Name() string
	NewClient(config Config) (APIClient, error)
	GetServiceStatus() (*Response, error)
}

// Config is configuraton to create the gomws base.
// AccessKey and SecretKey are optional, bette to set them in evn variables.
type Config struct {
	SellerId  string
	AuthToken string
	Region    string
	AccessKey string
	SecretKey string
}

// Credential return credential either from value set in config or load from env variables.
func (config Config) Credential() Credential {
	if config.AccessKey == "" || config.SecretKey == "" {
		return GetCredential()
	}

	return Credential{
		AccessKey: config.AccessKey,
		SecretKey: config.SecretKey,
	}
}

// Credential the credential to access the API.
type Credential struct {
	AccessKey string
	SecretKey string
}

// GetCredential get the credential from evn variables.
func GetCredential() Credential {
	return Credential{
		AccessKey: os.Getenv(envAccessKey),
		SecretKey: os.Getenv(envSecretKey),
	}
}

// Error represents the error message from the API.
type Error struct {
	// Error type. Values: Sender, Server.
	Type string `json:"Type"`
	// Amazon error code.
	Code string `json:"Code"`
	// Text explain the error.
	Message string `json:"Message"`
	// Detail about the error.
	Detail string `json:"Detail"`
}

// Inspect print out the value in a user friendly way.
func Inspect(value interface{}) {
	fmt.Printf("%# v", pretty.Formatter(value))
}
