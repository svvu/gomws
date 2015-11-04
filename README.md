# gomws
Amazon mws API in Go

# Usage
Import the packages
```go
import(
  "github.com/svvu/gomws/gmws"
  "github.com/svvu/gomws/mws/products"
)
```
Setup the configuration
```go
config := gmws.MwsConfig{
  SellerId:  "SellerId",
  AuthToken: "AuthToken",
  Region:    "US",

  // Optional if already set in env variable
  AccessKey: "AKey",
  SecretKey: "SKey",
}
```
If AccessKey and SecretKey not find in the pass in configuration, then it will try to retrieve them from env variables (**AWS_ACCESS_KEY** and **AWS_SecretKey**).

Create the client
```go
productsClient, err := products.NewClient(config)
```

Call the operations, the response is a struct contains result xml string and error if operation fail
```go
fmt.Println("------GetMatchingProduct------")
response := productsClient.GetMatchingProduct([]string{"ASIN"})
if response.Error != nil {
	fmt.Println(response.Error.Error())
}
// result is xml response in string
fmt.Println(response.Result)
```

# APIs
Currently only support Products

## Products
The Products API helps to get information to match your products to existing product listings on Amazon Marketplace websites.

The Products API returns product attributes, current Marketplace pricing information, and a variety of other product and listing information.

# TODO
* Add support for other APIs
* Add feature for response to convert result xml string to struct/json
