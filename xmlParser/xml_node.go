package xmlParser

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/clbanning/mxj"
)

/*
XMLNode is wrapper to mxj map.
It can traverse the xml to get the data.
NOTE:
	All attributes will also become a node with key '-attributesName'.
	And tags with attributes, their value will become a node with key '#text'.
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
type XMLNode struct {
	Value interface{}
	Path  string
}

// GenerateXMLNode generate an XMLNode object from the xml response body.
func GenerateXMLNode(xmlBuffer []byte) (*XMLNode, error) {
	m, err := mxj.NewMapXml(xmlBuffer)
	xNode := XMLNode{Value: m.Old()}

	return &xNode, err
}

// CurrentKey return the key (tag) for the current node.
// Root node has empty key.
func (xn *XMLNode) CurrentKey() string {
	keys := strings.Split(xn.Path, ".")
	return keys[len(keys)-1]
}

// FindByKey get the element data by any key (tag) in any depth.
// The method return a list XMLNode, each node represents all the sub-elements
// 	of the match key. The nodes then can be use to traverse deeper individually.
func (xn *XMLNode) FindByKey(key string) []XMLNode {
	valuesMap := []XMLNode{}

	xnode, err := xn.ToMap()
	if err != nil {
		return valuesMap
	}

	paths := xnode.PathsForKey(key)
	for _, p := range paths {
		nodes := xn.FindByPath(p)
		valuesMap = append(valuesMap, nodes...)
	}

	return valuesMap
}

// FindByKeys get the element data by any keys (tag) in any depth.
// Subsequential key need to be child of previous.
// The method return a list XMLNode, each node represents all the sub-elements
// 	of the match key. The nodes then can be use to traverse deeper individually.
// Ex:
//	If current node have path "A.B.C.D.E" and "A.B2.C2.D2.E",
// 	then node.FindByKeys("B", "E") will return nodes E under first path.
func (xn *XMLNode) FindByKeys(keys ...string) []XMLNode {
	valuesMap := []XMLNode{}

	if len(keys) <= 0 {
		return valuesMap
	}

	keysRegexp := regexp.MustCompile(
		`(^|\.)` + strings.Join(keys, `(\.(\w+\.)*)`) + `($|\.)`,
	)

	nodes := xn.FindByKey(keys[len(keys)-1])
	for _, node := range nodes {
		if keysRegexp.MatchString(node.Path) {
			valuesMap = append(valuesMap, node)
		}
	}

	return valuesMap
}

// FindByPath get the element data by path string.
// Path is relative to the current XMLNode.
// Path need to be start with direct sub node of current node.
// Path is separated by '.', ex: "Tag1.Tag2.Tag3".
// If current node is "A", and has sub path "B.C", then query on "B.C" will
// 	return node C. But Query on "C" will return empty nodes.
// The method return a list XMLNode, each node represents all the sub-elements
// 	of the match key. The nodes then can be use to traverse deeper individually.
func (xn *XMLNode) FindByPath(path string) []XMLNode {
	valuesMap := []XMLNode{}

	xnode, err := xn.ToMap()
	if err != nil {
		return valuesMap
	}

	values, err := xnode.ValuesForPath(path)
	if err != nil {
		return valuesMap
	}

	for _, m := range values {
		node := XMLNode{Value: m, Path: path}
		valuesMap = append(valuesMap, node)
	}

	return valuesMap
}

// FindByFullPath get the element data by absolute path.
// Path is separated by '.', ex: "Tag1.Tag2.Tag3".
// If current node have path "A.B.C", and query path is "A.B.C.D.E",
// 	then the method will search elements in current node with path "D.E".
// The method return a list XMLNode, each node represents all the sub-elements
// 	of the match key. The nodes then can be use to traverse deeper individually.
func (xn *XMLNode) FindByFullPath(path string) []XMLNode {
	subPath := strings.Replace(path, xn.Path+".", "", 1)
	return xn.FindByPath(subPath)
}

// Elements return the keys of immediate sub-elements of the current node.
func (xn *XMLNode) Elements() []string {
	xnode, err := xn.ToMap()
	if err != nil {
		return []string{}
	}

	elements, err := xnode.Elements("")
	if err != nil {
		return []string{}
	}

	return elements
}

// IsLeaf check whether or not the current node is a leaf node.
func (xn *XMLNode) IsLeaf() bool {
	return reflect.TypeOf(xn.Value).Kind() != reflect.Map
}

// LeafPaths return the path to the leaf nodes.
func (xn *XMLNode) LeafPaths() []string {
	xnode, err := xn.ToMap()
	if err != nil {
		return []string{}
	}

	return xnode.LeafPaths()
}

// LeafNodes return all the leaf nodes of current node.
func (xn *XMLNode) LeafNodes() []XMLNode {
	xnode, err := xn.ToMap()
	if err != nil {
		return []XMLNode{}
	}
	nodes := xnode.LeafNodes()
	lnodes := make([]XMLNode, len(nodes))
	for i, node := range nodes {
		lnodes[i] = XMLNode{
			Value: node.Value,
			Path:  node.Path,
		}
	}

	return lnodes
}

// ValueType return the type of node value.
func (xn *XMLNode) ValueType() reflect.Kind {
	return reflect.TypeOf(xn.Value).Kind()
}

/*
ToMap convert the node value to a mxj map.
If fail to convert, an error will be returned.
Tags have no sub tag, but have attributes is also a map.
Attributes of the tag has key '-attributesName'.
Tags' value has key '#test'.
Ex:
	<MessageId MarketplaceID="ATVPDKDDIKX0D" SKU="24478624">
		173964729
	</MessageId>
After to map,
 	map[string]string{
 		"-MarketplaceID": "ATVPDKDDIKX0D",
 		"-SKU": "24478624",
 		"#text": "173964729",
	}
*/
func (xn *XMLNode) ToMap() (mxj.Map, error) {
	if xn.ValueType() == reflect.Map {
		return mxj.Map(xn.Value.(map[string]interface{})), nil
	}

	return nil, errors.New("value is not a map")
}

// ToString convert the node value to string.
// If value is not valid string, an error will be returned.
func (xn *XMLNode) ToString() (string, error) {
	if xn.ValueType() == reflect.String {
		return xn.Value.(string), nil
	}

	return "", errors.New("value is not a valid string")
}

// ToInt convert the node value to int.
// If value is not valid int, an error will be returned.
func (xn *XMLNode) ToInt() (int, error) {
	switch xn.ValueType() {
	case reflect.String:
		value, err := xn.ToString()
		if err != nil {
			return 0, errors.New("can not convert value to int")
		}
		return strconv.Atoi(value)
	case reflect.Int:
		return xn.Value.(int), nil
	default:
		return 0, errors.New("can not convert value to int")
	}
}

// ToFloat64 convert the node value to float64.
// If value is not valid float64, an error will be returned.
func (xn *XMLNode) ToFloat64() (float64, error) {
	switch xn.ValueType() {
	case reflect.String:
		value, err := xn.ToString()
		if err != nil {
			return 0.0, errors.New("can not convert value to float64")
		}
		return strconv.ParseFloat(value, 64)
	case reflect.Float64:
		return xn.Value.(float64), nil
	default:
		return 0.0, errors.New("can not convert value to float64")
	}
}

// ToBool convert the node value to bool.
// If value is not valid bool, an error will be returned.
func (xn *XMLNode) ToBool() (bool, error) {
	switch xn.ValueType() {
	case reflect.String:
		value, err := xn.ToString()
		if err != nil {
			return false, errors.New("can not convert value to bool")
		}
		return strconv.ParseBool(value)
	case reflect.Bool:
		return xn.Value.(bool), nil
	default:
		return false, errors.New("can not convert value to bool")
	}
}

// ToTime convert the node value to timestamp.
// If value is not valid timestamp, an error will be returned.
func (xn *XMLNode) ToTime() (time.Time, error) {
	value, err := xn.ToString()
	if err != nil {
		return time.Time{}, errors.New("can not convert value to time")
	}
	t, err := time.Parse(time.RFC3339, value)
	return t, err
}

/*
ToStruct unmarshal the node value to struct.
If value can not be unmarshal, an error will returned.
ToStruct use json tag to unmarshal the map.
Ex:
To unmarshal the tag:
	<MessageId MarketplaceID="ATVPDKDDIKX0D" SKU="24478624">
		173964729
	</MessageId>
Can use struct:
	msgID := struct {
		MarketplaceID string `json:"-MarketplaceID"`
		SKU           string `json:"-SKU"`
		ID            string `json:"#text"`
	}{}
*/
func (xn *XMLNode) ToStruct(structPtr interface{}) error {
	xmap, err := xn.ToMap()
	if err != nil {
		return errors.New("value can not be unmarshal to struct")
	}
	return xmap.Struct(structPtr)
}

// XML return the raw xml data.
// If current node has key, use it as root node tag.
// If current node doest has key, and only have one child node, then the child
// 	node's key will become root node tag.
// If current node doest has key, and have more than one child node, then
// 	the a default tag <doc> will use as root tag.
func (xn *XMLNode) XML() ([]byte, error) {
	xmap, err := xn.ToMap()
	if err != nil {
		return []byte{}, err
	}
	rootTag := xn.CurrentKey()
	if rootTag == "" {
		return xmap.Xml()
	}
	return xmap.Xml(rootTag)
}

// PrintXML print the xml with two space indention.
func (xn *XMLNode) PrintXML() {
	xmap, err := xn.ToMap()
	if err != nil {
		fmt.Println("")
	}
	xml, _ := xmap.XmlIndent("", "  ")
	fmt.Println(string(xml))
}

func stringInSlice(s string, slice []string) bool {
	for _, str := range slice {
		if s == str {
			return true
		}
	}
	return false
}
