package gomws

import (
	"../mwsHttps"
	"os"
	"strings"
)

// The configuraton to create the gomws base.
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
	NewClient() (MwsClient, error)
	GetServiceStatus() (mwsHttps.Result, error)
}

const (
	EnvAccessKey = "AWS_ACCESS_KEY"
	EnvSecretKey = "AWS_SecretKey"
)

type Credential struct {
	AccessKey string
	SecretKey string
}

// GetCredential get the credential from evn variables.
func GetCredential() Credential {
	credential := Credential{}
	credential.AccessKey = os.Getenv(EnvAccessKey)
	credential.SecretKey = os.Getenv(EnvSecretKey)

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
