package main

import (
	"fmt"

	"github.com/svvu/gomws/gmws"
	"github.com/svvu/gomws/mws/products"
)

func main() {
	config := gmws.MwsConfig{
		SellerId:  "SellerId",
		AuthToken: "AuthToken",
		Region:    "US",

		// Optional if set in env variable
		AccessKey: "AKey",
		SecretKey: "SKey",
	}
	productsClient, err := products.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("------GetServiceStatus------")
	response := productsClient.GetServiceStatus()
	if response.Error != nil {
		fmt.Println(response.Error.Error())
	}
	xmlParser := gmws.NewXMLParser(response)
	xmlParser.PrettyPrint()

	fmt.Println("------GetMatchingProduct------")
	response = productsClient.GetMatchingProduct([]string{"B00ON8R5EO", "B000EVOSE4"})
	// Check http response error
	if response.Error != nil {
		fmt.Println(response.Error.Error())
	}

	xmlParser = gmws.NewXMLParser(response)
	// Check whether or not API send back error message
	if xmlParser.HasError() {
		fmt.Println(xmlParser.GetError())
	}

	gmp := products.GetMatchingProductResult{}
	xmlParser.Parse(&gmp)
	// Individual result might have error
	for _, r := range gmp.Results {
		if r.Error != nil {
			fmt.Println(r.Error)
		} else {
			fmt.Println(r.Products)
		}
	}
}
