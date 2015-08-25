package gmws

import (
	"os"
)

type MwsConfig struct {
	SellerId  string
	AuthToken string
	Region    string
	AccessKey string
	SecretKey string
}

type MwsClient interface {
	Version() string
	Name() string
}

const (
	EnvAccessKey = "AWS_ACCESS_KEY"
	EnvSecretKey = "AWS_SecretKey"
)

type Credential struct {
	AccessKey string
	SecretKey string
}

func GetCredential() Credential {
	credential := Credential{}
	credential.AccessKey = os.Getenv(EnvAccessKey)
	credential.SecretKey = os.Getenv(EnvSecretKey)

	return credential
}
