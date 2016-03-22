package gmws

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/clbanning/mxj"
)

// XMLNode is wrapper to mxj map.
// It can traverse the xml to get the data.
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

// Key return the key (tag) for the current node.
// Root node has empty key.
func (xn *XMLNode) Key() string {
	keys := strings.Split(xn.Path, ".")
	return keys[len(keys)-1]
}

// FindByKey get the element data by any key (tag) in any depth.
// The method return a list XMLNode, each node represents all the sub-elements
// 	of the match key. The nodes then can be use to traverse deeper individually.
func (xn *XMLNode) FindByKey(key string) ([]XMLNode, error) {
	xnode, err := xn.ToMap()
	if err != nil {
		return nil, errors.New("Node can't traverse deeper.")
	}

	paths := xnode.PathsForKey(key)
	valuesMap := []XMLNode{}
	for _, p := range paths {
		nodes, err := xn.FindByPath(p)
		if err != nil {
			return nil, err
		}
		valuesMap = append(valuesMap, nodes...)
	}

	return valuesMap, nil
}

// FindByPath get the element data by path string.
// Path is relative to the current XMLNode.
// Path is separated by '.', ex: "Tag1.Tag2.Tag3".
// The method return a list XMLNode, each node represents all the sub-elements
// 	of the match key. The nodes then can be use to traverse deeper individually.
func (xn *XMLNode) FindByPath(path string) ([]XMLNode, error) {
	xnode, err := xn.ToMap()
	if err != nil {
		return nil, errors.New("Node can't traverse deeper.")
	}

	values, err := xnode.ValuesForPath(path)
	if err != nil {
		return nil, err
	}

	valuesMap := make([]XMLNode, len(values))
	for i, m := range values {
		node := XMLNode{Value: m, Path: path}
		valuesMap[i] = node
	}

	return valuesMap, nil
}

// FindByFullPath get the element data by absolute path.
// Path is separated by '.', ex: "Tag1.Tag2.Tag3".
// If current node have path "A.B.C", and query path is "A.B.C.D.E",
// 	then the method will search elements in current node with path "D.E".
// The method return a list XMLNode, each node represents all the sub-elements
// 	of the match key. The nodes then can be use to traverse deeper individually.
func (xn *XMLNode) FindByFullPath(path string) ([]XMLNode, error) {
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

// ToMap convert the node value to a mxj map.
// If fail to convert, an error will be returned.
func (xn *XMLNode) ToMap() (mxj.Map, error) {
	if xn.ValueType() == reflect.Map {
		return mxj.Map(xn.Value.(map[string]interface{})), nil
	}

	return nil, errors.New("Value is not a map.")
}

// ToString convert the node value to string.
// If value is not valid string, an error will be returned.
func (xn *XMLNode) ToString() (string, error) {
	if xn.ValueType() == reflect.String {
		return xn.Value.(string), nil
	}

	return "", errors.New("Value is not a valid string.")
}

// ToInt convert the node value to int.
// If value is not valid int, an error will be returned.
func (xn *XMLNode) ToInt() (int, error) {
	value, err := xn.ToString()
	if err != nil {
		return 0, errors.New("Value is not a valid int.")
	}
	i, err := strconv.Atoi(value)
	return i, err
}

// ToFloat convert the node value to float64.
// If value is not valid float64, an error will be returned.
func (xn *XMLNode) ToFloat() (float64, error) {
	value, err := xn.ToString()
	if err != nil {
		return 0, errors.New("Value is not a valid float.")
	}
	f, err := strconv.ParseFloat(value, 64)
	return f, err
}

// ToBool convert the node value to bool.
// If value is not valid bool, an error will be returned.
func (xn *XMLNode) ToBool() (bool, error) {
	value, err := xn.ToString()
	if err != nil {
		return false, errors.New("Value is not a valid bool.")
	}
	b, err := strconv.ParseBool(value)
	return b, err
}

// ToTime convert the node value to timestamp.
// If value is not valid timestamp, an error will be returned.
func (xn *XMLNode) ToTime() (time.Time, error) {
	value, err := xn.ToString()
	if err != nil {
		return time.Time{}, errors.New("Value is not a valid time.")
	}
	t, err := time.Parse(time.RFC3339, value)
	return t, err
}

// ToStruct

// XML return the raw xml data.
func (xn *XMLNode) XML() ([]byte, error) {
	xmap, err := xn.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return xmap.Xml()
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
