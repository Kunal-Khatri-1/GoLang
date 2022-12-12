// ---------------------- I PICKED IT UP ONLINE, COULD'NT UNDERSTAND IT--------------------------

package main

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
)

// package for XML document query
// installing => go get github.com/antchfx/xmlquery

func main() {
	fmt.Println(" Dynamic XML parser without Struct in Go")

	s := `<?xml version="1.0" encoding="UTF-8" ?>
	<rss version="2.0">
	<channel>
	  <title>W3Schools Home Page</title>
	  <link>https://www.w3schools.com</link>
	  <description>Free web building tutorials</description>
	  <item>
		<title>RSS Tutorial</title>
		<link>https://www.w3schools.com/xml/xml_rss.asp</link>
		<description>New RSS tutorial on W3Schools</description>
	  </item>
	  <item>
		<title>XML Tutorial</title>
		<link>https://www.w3schools.com/xml</link>
		<description>New XML tutorial on W3Schools</description>
	  </item>
	</channel>
	</rss>`

	doc, err := xmlquery.Parse(strings.NewReader(s))
	checkNilError(err)

	fmt.Println(doc)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
