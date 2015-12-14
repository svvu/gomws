package gmws

import (
	"os"
	"strings"

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

// OptionalParams get the values from the pass in parameters.
// Only values for keys that are accepted will be returned.
//
// Note: The keys returned will be in title case.
//
// If the key appear in mulit parameters, later one will override the previous.
// Ex:
// 		ps := []mwsHttps.Parameters{
// 			{"key1": "value1", "key2": "value2"},
// 			{"key1": "newValue1", "key3": "value3"},
// 		}
// 		acceptKeys := []string{"key1", "key2"}
// 		resultParams := OptionalParams(acceptKeys, ps)
// result:
// 		resultParams -> {"Key1": "newValue1", "Key2": "value2"}
func OptionalParams(acceptKeys []string, ops []Parameters) Parameters {
	param := Parameters{}
	op := Parameters{}

	if len(ops) == 0 {
		return param
	}

	for _, p := range ops {
		op.Merge(p)
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