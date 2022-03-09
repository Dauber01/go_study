package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

//连接到数据库
func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp passwd=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

//获取指定的帖子
func retrieve(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, post.Content, post.Author)
	return
}

//创意一篇新帖子
func (post *Post) create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

//更新指定的帖子
func (post *Post) update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

//删除指定的帖子
func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}
