# gomws
Amazon mws API in Go

[![Build Status](https://travis-ci.org/svvu/gomws.svg?branch=master)](https://travis-ci.org/svvu/gomws)

# Usage
Example usage can be found in main.go

Import the packages
```go
import(
  "github.com/svvu/gomws/mws"
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
If AccessKey and SecretKey not find in the pass in configuration, then it will try to retrieve them from env variables (**AWS_ACCESS_KEY** and **AWS_SECRET_KEY**).

Create the client
```go
productsClient, err := products.NewClient(config)
```

Call the operations, the response is a struct contains result xml string and error if operation fail
```go
fmt.Println("------GetMatchingProduct------")
response, err := productsClient.GetMatchingProduct([]string{"ASIN"})
// Check http error.
if err != nil {
  fmt.Println(err.Error())
}
defer response.Body.Close()
// Check whether or not the response return MWS errors.
if response.Error != nil {
  fmt.Println(response.Error.Error())
}
```

Use XMLNode parser to get the data from response.

Create the xmlNode parser.
```go
/*
  xmlNode in fact is a map[string]interface{}.
  xmlNode is a wrapper over mxj map.

  All attributes will become a node with key '-attributesName'.
  Tags with attributes, their value will become a node with key '#text'.

  Ex:
    <ProductName sku="ABC">
      This will become node also.
    </ProductName>

  Will become:
    map[string]interface{
      "-sku": "ABC",
      "#text": "This will become node also.",
    }
*/
parser, err := response.ResultParser()
```
View the response in xml format.
```go
parser.PrintXML()
```
Get the data by key (xml tag name)
```go
// products is a slice of XMLNode, means individual one can be used to retrieve
// data by methods provided by XMLNode, ex: FindByKey
products := parser.FindByKey("Product")
```
Many methods can be used to traverse the xml tree. For more info, refer to the [godoc](https://godoc.org/github.com/svvu/gomws/gmws)
```go
// FindByKey get the nodes in any place of the tree.
productNodes := parser.FindByKey("Product")

// FindByKeys can used to retrieve nodes which are children of other nodes.
productNameNodes := parser.FindByKeys("Product", "Title")

// FindByPath get the nodes with specify tree path.
// Keys in the path are separated by '.'.
// Note: the first key must be a direct child of current node, and each subsequential
//  key must be direct child of previous key.
productNameNodes := parser.FindByPath("Product.AttributeSets.ItemAttributes.Title")
```
To get the value out of node, use the corresponding type methods
```go
xmlNode.ToString()
xmlNode.ToInt()
xmlNode.ToFloat()
xmlNode.ToBool()
xmlNode.ToTime()

// Ex:
productNameNodes := parser.FindByKeys("Product", "Title")
name, err := productNameNodes[0].ToString()
```
To unmarshall the data to a struct, use method
```go
xmlNode.ToStruct(structPointer)
```
The struct use json format tags.
```go
// Ex:
// To unmarshal the tag:
//  <Message>
//    <Locale>en_US</Locale>
//    <Text>Error message 1</Text>
//  </Message>
// Can use struct:
msg := struct {
  Locale string `json:"Locale"`
  Text   string `json:"Text"`
}{}
err := parser.FindByKey("Message")[0].ToStruct(&msg)

// To unmarshal the attributes, use -attributeName tag.
// To unmarshal the value of tags with attributes, use #text tag.
// Ex:
// To unmarshal the tag:
//  <MessageId MarketplaceID="ATVPDKDDIKX0D" SKU="24478624">
//		173964729
//  </MessageId>
// Can use struct:
type msgID struct {
  MarketplaceID string `json:"-MarketplaceID"`
  SKU           string `json:"-SKU"`
  ID            string `json:"#text"`
}
msgid := msgID{}
err := parser.FindByKey("MessageId")[0].ToStruct(&msgid)
```

Other usefull methods
```go
// Get the current tag name of the node.
xmlNode.CurrentKey()
// Get the direct children's node name.
xmlNode.Elements()
// Check whether or not the current node is the leaf, which means can't traverse deeper.
xmlNode.IsLeaf()
// Get a list of path to all the leaves.
xmlNode.LeafPaths()
// Get a list of nodes which are leaves.
xmlNode.LeafNodes()
```

# APIs

## Products
The Products API helps to get information to match your products to existing product listings on Amazon Marketplace websites.

The Products API returns product attributes, current Marketplace pricing information, and a variety of other product and listing information.

## Orders
The Orders API helps to retrieve orders information on Amazon Marketplace.

The Orders API returns orders list, items info in the order, and a variety of other orders information.

# TODO
* Add support for other APIs
* Record api request to test the endpoint methods
