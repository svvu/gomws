package main

import (
	. "./gmws"
	. "./mws"
	"fmt"
)

func main() {
	config := MwsConfig{
		SellerId:  "",
		AuthToken: "",
		AccessKey: "",
		SecretKey: "",
	}
	products := NewProductsClient(config)
	result, err := products.GetMatchingProductForId("ASIN", []string{"B000EVOSE4"})
	fmt.Println(result)
	fmt.Println(err)
}
