package xml

import (
	"strings"
)

// parseTagXMLElement parses the struct field tag from a special field:
//
//	struct {
//		XMLElement xml.Meta `xml:"name,empty"`
//		
//		// ...
//	}
func parseTagXMLElement(value string) (name string, empty bool) {
	var parts []string = strings.Split(value, ",")

	if len(parts) <= 0 {
		return "", false
	}

	name = strings.TrimSpace(parts[0])

	for _, part := range parts[1:] {
		part = strings.TrimSpace(part)

		switch part {
		case "empty":
			empty = true
		}
	}

	return
}
