package mws

import (
	"github.com/svvu/gomws/xmlParser"
)

// ResultParser a wrapper of xml parser to add some helper methods.
type ResultParser struct {
	*xmlParser.XMLNode
}

// NewResultParser create a new result parser.
func NewResultParser(data []byte) (*ResultParser, error) {
	xmlNode, err := xmlParser.GenerateXMLNode(data)

	if err != nil {
		return nil, err
	}

	return &ResultParser{XMLNode: xmlNode}, nil
}

// HasErrorNodes will check whether or not the xml node tree has any erorr node.
// If it contains errors, true will be returned.
func (rp *ResultParser) HasErrorNodes() bool {
	errorNodes := rp.FindByKey("Error")
	if len(errorNodes) > 0 {
		return true
	}
	return false
}

// GetMWSErrors will return an array of Error struct from the xml node tree.
func (rp *ResultParser) GetMWSErrors() ([]Error, error) {
	errorNodes := rp.FindByKey("Error")

	errors := []Error{}
	for _, en := range errorNodes {
		apiError := Error{}
		err := en.ToStruct(&apiError)
		if err != nil {
			return errors, err
		}
		errors = append(errors, apiError)
	}
	return errors, nil
}
