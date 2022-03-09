package ch7_test

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateXml(t *testing.T) {
	post := Post{
		Id:      "1",
		Content: "Hello word",
		Author: Author{
			Id:   "1",
			Name: "ahbama",
		},
	}
	//output, err := xml.Marshal(&post)
	//使用MarshalIndent方法可以格式化输出文件格式,第二参表示每行以什么开头,第三参表示锁进内容
	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		panic(err)
	}
	//err = ioutil.WriteFile("post_create.xm", output, 0644)
	//格式化方法输出的文件
	//err = ioutil.WriteFile("post_create_fomart.xml", output, 0644)
	//增加xml头信息
	err = ioutil.WriteFile("post_create_fomart.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		panic(err)
	}
}

//使用encoder的方式创建xml文件
func TestEncoderCreateXml(t *testing.T) {
	post := Post{
		Id:      "1",
		Content: "Hello word",
		Author: Author{
			Id:   "1",
			Name: "ahbama",
		},
	}
	xmlFile, err := os.Create("post_encoder.xml")
	if err != nil {
		panic(err)
	}
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		panic(err)
	}
}
