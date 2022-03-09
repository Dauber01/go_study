package ch7_test

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type PostNew struct { //#A
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func TestNewXml(t *testing.T) {
	xmlFile, err := os.Open("post_new.xml")
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}
	var postnew PostNew
	xml.Unmarshal(xmlData, &postnew)
	t.Log(postnew)
}

//当需要处理流式的xml数据,或者当xml数据过大,不能够用unmarshal一次进行解析的时候,使用decoder手动进行解析
func TestDecoderXml(t *testing.T) {
	xmlFile, err := os.Open("post_new.xml")
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			//当没有元素的时候返回io.EOF
			break
		}
		if err != nil {
			panic(err)
		}
		switch se := t.(type) {
		case xml.StartElement:
			//在该实例中判断是否comment,只对comment进行解析
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se)
			}
		}
	}
}
