package xml

import (
	"strings"
)

// parseTagStructField parses the struct field tag from a field in a struct.
// For example:
//
//	struct {
//		// ...
//		
//		Banana string `xml:"banana,attr"`
//		
//		// ...
//	}
func parseTagStructField(value string) (name string, attr bool) {
	var parts []string = strings.Split(value, ",")

	if len(parts) <= 0 {
		return "", false
	}

	name = strings.TrimSpace(parts[0])

	for _, part := range parts[1:] {
		part = strings.TrimSpace(part)

		switch part {
		case "attr":
			attr = true
		}
	}

	return
}
