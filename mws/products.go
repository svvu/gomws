package mws

import (
	. "../gmws"
	. "../mwsHttpClient"
)

type Products struct {
	*MwsBase
}

func NewProductsClient(config MwsConfig) *Products {
	prodcuts := new(Products)
	base := NewMwsBase(config, prodcuts.Version(), prodcuts.Name())
	prodcuts.MwsBase = base
	return prodcuts
}

func (p Products) Version() string {
	return "2011-10-01"
}

func (p Products) Name() string {
	return "Products"
}

func (p Products) GetMatchingProductForId(idType string, idList []string) (Result, error) {
	params := Parameters{
		"Action":        "GetMatchingProductForId",
		"IdType":        idType,
		"IdList":        idList,
		"MarketplaceId": p.MarketPlaceId,
	}
	structedParams, err := params.StructureKeys("IdList", "Id").Normalize()

	if err != nil {
		return "", err
	}

	httpClient := p.HttpClient(structedParams)
	return httpClient.Request()
}
