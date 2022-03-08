package ch6

import (
	"encoding/csv"
	"os"
	"strconv"
	"testing"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func TestCsv(t *testing.T) {
	csvFile, err := os.Create("post.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{1, "哈哈", "mark town"},
		Post{2, "呵呵", "harminwei"},
		Post{3, "嘿嘿", "jorden"},
	}

	writer := csv.NewWriter(csvFile)

	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	file, err := os.Open("post.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	postNew := []Post{}
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		po := Post{Id: int(id), Content: item[1], Author: item[2]}
		postNew = append(postNew, po)
	}
	t.Log(postNew[0].Id)
	t.Log(postNew[0].Content)
	t.Log(postNew[0].Author)
}
