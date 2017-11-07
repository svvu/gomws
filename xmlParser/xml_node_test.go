package xmlParser

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func XMLNodeTestExample() []byte {
	response, ferr := ioutil.ReadFile("./examples/XMLNodeTest.xml")

	if ferr != nil {
		fmt.Println(ferr)
		return []byte{}
	}

	return response
}

func Test_GenerateXMLNode(t *testing.T) {
	Convey("Create xml node success", t, func() {
		xNode, err := GenerateXMLNode(XMLNodeTestExample())

		Convey("Err is nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Node value is not empty", func() {
			So(xNode.Value, ShouldNotBeEmpty)
		})
	})

	Convey("Create xml node fail", t, func() {
		exampleString := []byte("Not xml format, so should fail.")
		xNode, err := GenerateXMLNode(exampleString)

		Convey("Err is not nil", func() {
			So(err, ShouldNotBeNil)
		})

		Convey("Node value should be empty", func() {
			So(xNode.Value, ShouldBeEmpty)
		})
	})
}

func Test_CurrentKey(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())

	Convey("Root's key should be empty", t, func() {
		key := xNode.CurrentKey()
		So(key, ShouldBeZeroValue)
	})

	Convey("Message tag's key should be Message", t, func() {
		msgNodes := xNode.FindByKey("Message")
		msgNode := msgNodes[0]
		key := msgNode.CurrentKey()
		So(key, ShouldEqual, "Message")
	})
}

func Test_FindByKey(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())

	Convey("Find by key 'Status'", t, func() {
		fNodes := xNode.FindByKey("Status")

		Convey("Should have 1 nodes found", func() {
			So(fNodes, ShouldHaveLength, 1)
		})

		Convey("Node found's value is string 'GREEN_I'", func() {
			v, _ := fNodes[0].ToString()
			So(v, ShouldEqual, "GREEN_I")
		})
	})

	Convey("Find by key 'Message'", t, func() {
		fNodes := xNode.FindByKey("Message")

		Convey("Should have 4 nodes found", func() {
			So(fNodes, ShouldHaveLength, 4)
		})
	})

	Convey("Find by key 'NotExistKey'", t, func() {
		fNodes := xNode.FindByKey("NotExistKey")

		Convey("Should have 0 nodes found", func() {
			So(fNodes, ShouldHaveLength, 0)
		})
	})

	Convey("Node's value is not a valid map", t, func() {
		fNodes := xNode.FindByKey("Status")
		fNodes = fNodes[0].FindByKey("Status")

		Convey("Should have 0 nodes found", func() {
			So(fNodes, ShouldHaveLength, 0)
		})
	})
}

func Test_FindByKeys(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())

	Convey("Find by key 'Message2', 'Text'", t, func() {
		fNodes := xNode.FindByKeys("Message2", "Text")

		Convey("Should have 1 nodes found", func() {
			So(fNodes, ShouldHaveLength, 1)
		})

		Convey("Node found's value is string 'Error message 5'", func() {
			v, _ := fNodes[0].ToString()
			So(v, ShouldEqual, "Error message 5")
		})
	})

	Convey("Find by keys 'Message', 'Text'", t, func() {
		// Text under message2 will be ignored
		fNodes := xNode.FindByKeys("Message", "Text")

		Convey("Should have 4 nodes found", func() {
			So(fNodes, ShouldHaveLength, 4)
		})
	})

	Convey("Find by key 'Message', 'Messages'", t, func() {
		fNodes := xNode.FindByKeys("Message", "Messages")

		// Messages is not child node of Message.
		Convey("Should have 0 nodes found", func() {
			So(fNodes, ShouldHaveLength, 0)
		})
	})
}

func Test_FindByPath(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())

	Convey("Find by path Message.Locale", t, func() {
		fNodes := xNode.FindByKey("Messages")
		fNodes = fNodes[0].FindByPath("Message.Locale")

		Convey("Should have 4 nodes found", func() {
			So(fNodes, ShouldHaveLength, 4)
		})
	})

	Convey("When query path is not direct path of current node", t, func() {
		fNodes := xNode.FindByPath("Messages.Message.Locale")

		Convey("Should have 0 nodes found", func() {
			So(fNodes, ShouldHaveLength, 0)
		})
	})

	Convey("Node's value is not a valid map", t, func() {
		fNodes := xNode.FindByPath(
			"GetServiceStatusResponse.GetServiceStatusResult.Status",
		)
		fNodes = fNodes[0].FindByPath("Status")

		Convey("Should have 0 nodes found", func() {
			So(fNodes, ShouldHaveLength, 0)
		})
	})
}

func Test_FindByFullPath(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())
	fNode := xNode.FindByPath(
		"GetServiceStatusResponse.GetServiceStatusResult.Messages",
	)

	Convey("Find by path Message.Locale", t, func() {
		fNodes := fNode[0].FindByFullPath(
			"GetServiceStatusResponse.GetServiceStatusResult.Messages.Message.Locale",
		)

		Convey("Should have 4 nodes found", func() {
			So(fNodes, ShouldHaveLength, 4)
		})
	})
}

func Test_Elements(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())

	Convey("Elements of Messages", t, func() {
		fNode := xNode.FindByKey("Messages")
		elements := fNode[0].Elements()

		Convey("Should have 2 key", func() {
			So(elements, ShouldHaveLength, 2)
		})

		Convey("Should be Message", func() {
			So(elements[0], ShouldEqual, "Message")
		})
	})

	Convey("Elements of Message", t, func() {
		fNode := xNode.FindByKey("Message")
		elements := fNode[0].Elements()

		Convey("Should have 2 keys", func() {
			So(elements, ShouldHaveLength, 2)
		})

		Convey("First key should be Locale", func() {
			So(elements[0], ShouldEqual, "Locale")
		})

		Convey("Second key should be Text", func() {
			So(elements[1], ShouldEqual, "Text")
		})
	})
}

func Test_IsLeaf(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())

	Convey("Messages tag is not leaf", t, func() {
		fNode := xNode.FindByKey("Messages")

		So(fNode[0].IsLeaf(), ShouldBeFalse)
	})

	Convey("Text tag is leaf", t, func() {
		fNode := xNode.FindByKey("Text")

		So(fNode[0].IsLeaf(), ShouldBeTrue)
	})
}

func Test_LeafPaths(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())

	Convey("Root node", t, func() {
		leafPaths := xNode.LeafPaths()

		// 15 tags and 2 attributes.
		Convey("Should have 19 leaves", func() {
			So(leafPaths, ShouldHaveLength, 19)
		})
	})

	Convey("Leaf node", t, func() {
		fNode := xNode.FindByKey("Text")
		leafPaths := fNode[0].LeafPaths()

		Convey("Should have no leaf", func() {
			So(leafPaths, ShouldHaveLength, 0)
		})
	})
}

func Test_LeafNodes(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())

	Convey("Root node", t, func() {
		leafNodes := xNode.LeafNodes()

		// 15 tags and 2 attributes.
		Convey("Should have 19 leaf nodes", func() {
			So(leafNodes, ShouldHaveLength, 19)
		})
	})

	Convey("Leaf node", t, func() {
		fNode := xNode.FindByKey("Text")
		leafNodes := fNode[0].LeafNodes()

		Convey("Should have no leaf", func() {
			So(leafNodes, ShouldHaveLength, 0)
		})
	})
}

func Test_ValueType(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())

	Convey("Value is a map", t, func() {
		So(xNode.ValueType(), ShouldEqual, reflect.Map)
	})

	Convey("Value is string", t, func() {
		fNode := xNode.FindByKey("Locale")
		So(fNode[0].ValueType(), ShouldEqual, reflect.String)
	})
}

func Test_ToMap(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())

	Convey("Value is a map", t, func() {
		node, err := xNode.ToMap()

		Convey("Error is nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("A map is returned", func() {
			So(reflect.TypeOf(node).Kind(), ShouldEqual, reflect.Map)
		})
	})

	Convey("Value is not a map", t, func() {
		fNode := xNode.FindByKey("Locale")

		node, err := fNode[0].ToMap()

		Convey("Error is not nil", func() {
			So(err, ShouldNotBeNil)
		})

		Convey("A nil map is returned", func() {
			So(node, ShouldBeNil)
		})
	})
}

func Test_ToString(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())

	Convey("Value is a string", t, func() {
		fNode := xNode.FindByKey("Locale")

		s, err := fNode[0].ToString()

		Convey("Error is nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("A string 'en_US' is returned", func() {
			So(s, ShouldEqual, "en_US")
		})
	})

	Convey("Value is not a string", t, func() {
		s, err := xNode.ToString()

		Convey("Error is not nil", func() {
			So(err, ShouldNotBeNil)
		})

		Convey("An empty string is returned", func() {
			So(s, ShouldEqual, "")
		})
	})
}

func Test_ToInt(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())

	Convey("Value is a int", t, func() {
		fNode := xNode.FindByKey("-SKU")

		i, err := fNode[0].ToInt()

		Convey("Error is nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Value 24478624 is returned", func() {
			So(i, ShouldEqual, 24478624)
		})
	})

	Convey("Value is not a int", t, func() {
		i, err := xNode.ToInt()

		Convey("Error is not nil", func() {
			So(err, ShouldNotBeNil)
		})

		Convey("A zero value is returned", func() {
			So(i, ShouldEqual, 0)
		})
	})
}

func Test_ToFloat(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())

	Convey("Value is a float", t, func() {
		fNode := xNode.FindByKey("UpHour")

		i, err := fNode[0].ToFloat()

		Convey("Error is nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Value 2.3 is returned", func() {
			So(i, ShouldEqual, 2.3)
		})
	})

	Convey("Value is not a float", t, func() {
		i, err := xNode.ToFloat()

		Convey("Error is not nil", func() {
			So(err, ShouldNotBeNil)
		})

		Convey("A zero value is returned", func() {
			So(i, ShouldEqual, 0)
		})
	})
}

func Test_ToBool(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())

	Convey("Value is a bool", t, func() {
		fNode := xNode.FindByKey("UP")

		i, err := fNode[0].ToBool()

		Convey("Error is nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Value is true", func() {
			So(i, ShouldBeTrue)
		})
	})

	Convey("Value is not a float", t, func() {
		i, err := xNode.ToBool()

		Convey("Error is not nil", func() {
			So(err, ShouldNotBeNil)
		})

		Convey("Default false is returned", func() {
			So(i, ShouldBeFalse)
		})
	})
}

func Test_ToTime(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())

	Convey("Value is timestamp", t, func() {
		fNode := xNode.FindByKey("Timestamp")

		i, err := fNode[0].ToTime()

		Convey("Error is nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Value is 2013-09-05 18:12:21.83 +0000 UTC", func() {
			So(i.String(), ShouldEqual, "2013-09-05 18:12:21.83 +0000 UTC")
		})
	})

	Convey("Value is not a time", t, func() {
		i, err := xNode.ToTime()

		Convey("Error is not nil", func() {
			So(err, ShouldNotBeNil)
		})

		Convey("Zero time is returned", func() {
			So(i.String(), ShouldEqual, "0001-01-01 00:00:00 +0000 UTC")
		})
	})
}

func Test_ToStruct(t *testing.T) {
	msg := struct {
		Locale string `json:"Locale"`
		Text   string `json:"Text"`
	}{}
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())

	Convey("Value is not a map", t, func() {
		fNode := xNode.FindByKey("Timestamp")

		err := fNode[0].ToStruct(&msg)

		Convey("Error is not nil", func() {
			So(err, ShouldNotBeNil)
		})

		Convey("struct is returned", func() {
			So(msg, ShouldBeZeroValue)
		})
	})

	Convey("Value is a map", t, func() {
		fNode := xNode.FindByKey("Message")

		err := fNode[0].ToStruct(&msg)

		Convey("Error is nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("msg struct Local is en_US", func() {
			So(msg.Locale, ShouldEqual, "en_US")
		})

		Convey("msg struct Text is Error message 1", func() {
			So(msg.Text, ShouldEqual, "Error message 1")
		})
	})

	Convey("Value is not a map but with xml attributes", t, func() {
		msgID := struct {
			MarketplaceID string `json:"-MarketplaceID"`
			SKU           string `json:"-SKU"`
			ID            string `json:"#text"`
		}{}
		fNode := xNode.FindByKey("MessageId")

		err := fNode[0].ToStruct(&msgID)

		Convey("Error is nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("msgID struct MarketplaceID is ATVPDKDDIKX0D", func() {
			So(msgID.MarketplaceID, ShouldEqual, "ATVPDKDDIKX0D")
		})

		Convey("msgID struct SKU is 24478624", func() {
			So(msgID.SKU, ShouldEqual, "24478624")
		})

		Convey("msgID struct ID is 173964729", func() {
			So(msgID.ID, ShouldEqual, "173964729")
		})
	})
}

func Test_XML(t *testing.T) {
	xNode, _ := GenerateXMLNode(XMLNodeTestExample())
	fNode := xNode.FindByKey("Message")

	Convey("Value is not a map", t, func() {
		xml, err := fNode[0].XML()

		Convey("Error not nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("xml is returned", func() {
			expectedXML := "<Message><Locale>en_US</Locale><Text>Error message 1</Text></Message>"
			So(string(xml), ShouldEqual, expectedXML)
		})
	})
}
