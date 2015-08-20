package gmws

import (
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
	Path() string
	Endpoint() string
}

type MwsBase struct {
	SellerId      string
	AuthToken     string
	Region        string
	MarketPlaceId string
	Host          string
	accessKey     string
	secretKey     string
}

// sellerId, authToken, region string
func NewMwsBase(config MwsConfig) *MwsBase {
	if config.SellerId == "" {
		// Log
		return nil
	}

	if config.AuthToken == "" {
		// Log
		return nil
	}

	region := config.Region
	if region == "" {
		region = "US"
	}

	marketPlace, mError := NewMarketPlace(region)
	if eError {
		// TODO
	}

	base := MwsBase{
		SellerId:      config.SellerId,
		AuthToken:     config.AuthToken,
		Region:        region,
		MarketPlaceId: marketPlace.Id,
		Host:          marketPlace.endpoint,
		accessKey:     config.AccessKey,
		secretKey:     config.SecretKey,
	}
	return &base
}

func (base MwsBase) SignatureMethod() string {
	return "HmacSHA256"
}

func (base MwsBase) SignatureVersion() string {
	return "2"
}

func (base MwsBase) paramsToAugment() map[string]string {
	clientInfo := map[string]string{
		"SellerId":         base.SellerId,
		"MWSAuthToken":     base.authToken,
		"SignatureMethod":  base.SignatureMethod(),
		"SignatureVersion": base.SignatureVersion(),
		"AWSAccessKeyId":   base.getCredential().AccessKey,
	}
	return clientInfo
}

func (base MwsBase) getCredential() Credential {
	if base.accessKey != "" && base.secretKey != "" {
		return Credential{base.accessKey, base.secretKey}
	} else {
		return GetCredential()
	}
}

type Credential struct {
	AccessKey string
	SecretKey string
}

func GetCredential() Credential {
	credential := Credential{}
	credential.AccessKey = os.Getenv(AWSAccessKeyId)
	credential.SecretKey = os.Getenv(SecretAccessKey)

	return credential
}
