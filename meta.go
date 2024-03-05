package xml

import (
	"reflect"
)

// Meta is used to add meta-data to the struct (rather than just a struct field).
//
// For example:
//
//	type Link struct {
//		XMLElement xml.Meta `xml:"link,empty"` // <----------
//		
//		Rel  string `xml:"rel,attr"`
//		HRef string `xml:"href,attr"`
//	}
type Meta struct{}

var (
	reflectedTypeMeta reflect.Type = reflect.TypeOf(Meta{})
)
