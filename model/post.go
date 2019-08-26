package model

import (
	"errors"
	"fmt"

	"github.com/gamebody/goweb/initdb"
)

// Post 用户相关字段
type Post struct {
	Author     int64  `form:"author"`
	Title      string `form:"title"`
	Content    string `form:"content"`
	PV         int    `form:"pv"`
	CreateTime int    `form:"create_time"`
}

// Save 储存用户
func (p *Post) Save() error {
	sqlstr := "insert into blog.post (author, title, content) values (?,?,?)"

	_, err := initdb.Db.Exec(sqlstr, p.Author, p.Title, p.Content)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

// GetPosts 根据用户获取全部的文章
func GetPosts(author int64) []Post {
	var posts []Post

	stmtOut, err := initdb.Db.Prepare("SELECT author,title,content,pv,create_time FROM blog.post WHERE author = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(author)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		post := &Post{}
		rows.Scan(&post.Author, &post.Title, &post.Content, &post.PV, &post.CreateTime)

		fmt.Println(*post)

		posts = append(posts, *post)
	}

	return posts
}

// Check 校验post表单
func (p *Post) Check() error {
	if len(p.Title) == 0 {
		return errors.New("请填写标题")
	}
	if len(p.Content) == 0 {
		return errors.New("请填写内容")
	}
	return nil
}
