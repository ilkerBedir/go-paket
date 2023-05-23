package xml_parser

import (
	"os"
	"log"
	"io"
	"encoding/xml"
	"fmt"
)
type Users struct {
    XMLName xml.Name `xml:"users"`
    Users   []User   `xml:"user"`
}
type User struct {
    XMLName xml.Name `xml:"user"`
    Type    string   `xml:"type,attr"`
    Name    string   `xml:"name"`
    Social  Social   `xml:"social"`
}
type Social struct {
    XMLName  xml.Name `xml:"social"`
    Facebook string   `xml:"facebook"`
    Twitter  string   `xml:"twitter"`
    Youtube  string   `xml:"youtube"`
}
func parser(){
	xmlFile,err:=os.Open("users.xml")
	if err!=nil{
        log.Fatal(err)
    }
	defer xmlFile.Close()
	byteValues,err:=io.ReadAll(xmlFile)
	if err!=nil{
		log.Fatal(err)
	}
	var users Users
	ParseXMLFile(byteValues,&users) 
	for i := 0; i < len(users.Users); i++ {
        fmt.Println("User Type: " + users.Users[i].Type)
        log.Println("User Name: " + users.Users[i].Name)
        fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
    }
}