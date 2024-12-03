package actio

import (
	"encoding/xml"
	"fmt"
)

// EncodeXML encodes a map of string key-value pairs into an XML string with the specified root element.
//
// Example Usage:
//
//	data := map[string]string{"name": "John", "age": "30"}
//	xmlBytes := EncodeXML(data, "person")
//	fmt.Println(string(xmlBytes)) // <person><name>John</name><age>30</age></person>
func EncodeXML(data map[string]string, rootElement string) []byte {
	xmlString := fmt.Sprintf("<%v>", rootElement)
	for k, v := range data {
		xmlString += fmt.Sprintf("<%s>%s</%s>", k, v, k)
	}
	xmlString += fmt.Sprintf("</%v>", rootElement)
	return []byte(xmlString)
}

// DecodeXML decodes an XML string into a map of string key-value pairs.
//
// Example Usage:
//
//	xmlString := "<person><name>John</name><age>30</age></person>"
//	data := DecodeXML([]byte(xmlString))
//	fmt.Println(data) // map[age:30 name:John]
func DecodeXML(xmlBytes []byte) map[string]string {
	type Entry struct {
		XMLName xml.Name
		Value   string `xml:",chardata"`
	}

	type Root struct {
		Entries []Entry `xml:",any"`
	}

	var root Root
	if err := xml.Unmarshal(xmlBytes, &root); err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return nil
	}

	data := make(map[string]string)
	for _, entry := range root.Entries {
		data[entry.XMLName.Local] = entry.Value
	}

	return data
}
