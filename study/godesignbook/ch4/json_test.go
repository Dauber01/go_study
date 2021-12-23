package ch4

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"` //omitempty 表示如果空值,不现实属性
	Actors []string
}

func TestJson(t *testing.T) {
	data := jsonFormat()
	jsonToStru(data)
}

func jsonToStru(data []byte) {
	var title []struct{ Title string }
	if err := json.Unmarshal(data, &title); err != nil {
		log.Fatalf("json unmarshal fail %s", err)
	}
	//对应的 .NewDecoder().Decode(&) 为流式解码器,在此处无法使用
	//json.NewDecoder(data).Decode(&title)
	fmt.Println(title)
}

func jsonFormat() []byte {
	var movie = []Movie{
		{Title: "haha", Year: 1990, Color: false, Actors: []string{"a", "b"}},
		{Title: "hehe", Year: 1999, Color: true, Actors: []string{"c", "d"}},
	}
	//data, err := json.Marshal(movie)
	//格式化输出,行开头的标记和缩进符
	data, err := json.MarshalIndent(movie, "", "	")
	if err != nil {
		log.Fatalf("json unmarshal fail %s", err)
	}
	fmt.Printf("%s\n", data)
	fmt.Printf("%T\n", data)
	return data
}
