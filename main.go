package main

import (
	"fmt"

	"github.com/svvu/gomws/mws"
	"github.com/svvu/gomws/mws/products"
)

func main() {
	config := mws.Config{
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

	// Example 1
	fmt.Println("------GetServiceStatus------")
	statusResponse, err := productsClient.GetServiceStatus()
	// Check http client error.
	if err != nil {
		fmt.Println(err.Error())
	}
	defer statusResponse.Body.Close()
	// Check whether or not the API return errors.
	if statusResponse.Error != nil {
		fmt.Println(statusResponse.Error.Error())
	} else {
		xmlNode, _ := statusResponse.ResultParser()
		xmlNode.PrintXML() // Print the xml response with indention.
	}

	// Example 2
	fmt.Println("------GetMatchingProduct------")
	proResponse, err := productsClient.GetMatchingProduct([]string{"B00ON8R5EO", "B000EVOSE4"})
	if err != nil {
		fmt.Println(err.Error())
	}
	defer proResponse.Body.Close()
	if proResponse.Error != nil {
		fmt.Println(proResponse.Error.Error())
		return
	}

	// Create a result parser for the response.
	parser, _ := proResponse.ResultParser()

	// Get the first product from response.
	productOne := parser.FindByKey("Product")[0]

	// Find the title node.
	productNameNode := productOne.FindByKey("Title")
	// Get the name value.
	name, err := productNameNode[0].ToString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Product name: %v \n", name)

	// Find the height for package dimensions.
	heightNode := productOne.FindByKeys("PackageDimensions", "Height")
	// Inspect the heightNode map.
	mws.Inspect(heightNode)
}
