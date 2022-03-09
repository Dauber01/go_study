package ch7_test

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
)

type Post struct { //#A
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func TestXml(t *testing.T) {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}
	var post Post
	xml.Unmarshal(xmlData, &post)
	t.Log(post)
}
