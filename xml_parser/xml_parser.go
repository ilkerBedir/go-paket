package xml_parser
import (
	"encoding/xml"
	"log"
)
func ParseXMLFile[T any](values []byte,mapper T) any{
	err := xml.Unmarshal(values,&mapper)
    if err != nil {
		log.Fatal(err)
	}
	return mapper
}