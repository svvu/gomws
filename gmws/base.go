package gmws

import (
	. "../http_client"
)

type MwsBase struct {
	SellerId      string
	AuthToken     string
	Region        string
	MarketPlaceId string
	Host          string
	Version       string
	Name          string
	accessKey     string
	secretKey     string
}

// sellerId, authToken, region string
func NewMwsBase(config MwsConfig, version, name string) *MwsBase {
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
	if mError != nil {
		// TODO
	}

	base := MwsBase{
		SellerId:      config.SellerId,
		AuthToken:     config.AuthToken,
		Region:        region,
		MarketPlaceId: marketPlace.Id,
		Host:          marketPlace.EndPoint,
		Version:       version,
		Name:          name,
		accessKey:     config.AccessKey,
		secretKey:     config.SecretKey,
	}
	return &base
}

func (base MwsBase) Path() string {
	return "/" + base.Name + "/" + base.Version
}

func (base MwsBase) EndPoint() string {
	return base.Host + "/" + base.Path()
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
		"MWSAuthToken":     base.AuthToken,
		"SignatureMethod":  base.SignatureMethod(),
		"SignatureVersion": base.SignatureVersion(),
		"AWSAccessKeyId":   base.getCredential().AccessKey,
		"Version":          base.Version,
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

func (base MwsBase) HttpClient(params NormalizedParameters) *MwsHttpClient {
	httpClient := MwsHttpClient{
		EndPoint:   base.EndPoint(),
		Parameters: params,
	}

	httpClient.AugmentParameters(base.paramsToAugment())
	httpClient.SignQuery(base.getCredential().SecretKey)
	return &httpClient
}
