package xml_test

import (
	"fmt"
	"strings"

	"github.com/reiver/go-xml"
)

func ExampleEncode_emptyElement() {

	type Link struct {
		XMLElement xml.Meta `xml:"link,empty"`
		Rel string  `xml:"rel,attr"`
		HRef string `xml:"href,attr"`
	}

	var link Link = Link{
		Rel: "enclosure",
		HRef: "http://example.com/path/to/video.php",
	}

	var buffer strings.Builder

	err := xml.Encode(&buffer, link)
	if nil != err {
		fmt.Println("ERROR: problem encoding to XML:", err)
		return
	}

	fmt.Println(buffer.String())

	// Output:
	//
	// <link rel="enclosure" href="http://example.com/path/to/video.php" />
}
