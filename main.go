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
	xmlNode, _ := gmws.GenerateXMLNode(response.Body)
	// Print the xml response with indention.
	xmlNode.PrintXML()

	fmt.Println("------GetMatchingProduct------")
	response = productsClient.GetMatchingProduct([]string{"B00ON8R5EO", "B000EVOSE4"})
	// Check http response error
	if response.Error != nil {
		fmt.Println(response.Error.Error())
	}

	xmlNode, _ = gmws.GenerateXMLNode(response.Body)
	// Check whether or not API send back error message
	if gmws.HasErrors(xmlNode) {
		fmt.Println(gmws.GetErrors(xmlNode))
	}

	// Get the first product from response.
	productOne := xmlNode.FindByKey("Product")[0]

	// Find the title node.
	productNameNode := productOne.FindByKey("Title")
	// Get the name value.
	name, err := productNameNode[0].ToString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Product name: %v \n", name)
	productOne.PrintXML()

	// Find the height for package dimensions.
	heightNode := productOne.FindByKeys("PackageDimensions", "Height")
	// Inspect the heightNode map.
	gmws.Inspect(heightNode)
}
