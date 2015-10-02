package main

import (
	. "./gmws"
	"./mws/products"
	// . "./mwsHttps"
	"fmt"
)

func main() {
	config := MwsConfig{
		SellerId:  "",
		AuthToken: "",
		AccessKey: "",
		SecretKey: "",
	}
	products := products.NewClient(config)
	fmt.Println("------GetServiceStatus------")
	result, err := products.GetServiceStatus()
	result.PrettyPrint()
	if err != nil {
		fmt.Println(err)
	}

}
