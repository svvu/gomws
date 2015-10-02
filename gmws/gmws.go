package gmws

import (
	"../mwsHttps"
	"os"
	"strings"
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

func OptionalParams(acceptKeys []string, ops []mwsHttps.Parameters) mwsHttps.Parameters {
	param := mwsHttps.Parameters{}
	op := mwsHttps.Parameters{}

	if len(ops) == 0 {
		return param
	} else {
		for _, p := range ops {
			op.Merge(p)
		}
	}

	for _, key := range acceptKeys {
		value, ok := op[key]
		if ok {
			param[strings.Title(key)] = value
			delete(op, key)
		}
	}
	return param
}
