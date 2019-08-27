package model

import (
	"errors"
	"fmt"

	"github.com/gamebody/goweb/initdb"
)

// Post 用户相关字段
type Post struct {
	ID         int    `form:"id"`
	Author     int64  `form:"author"`
	Title      string `form:"title"`
	Content    string `form:"content"`
	PV         int    `form:"pv"`
	CreateTime string `form:"create_time"`
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

// Edit 编辑文章
func (p *Post) Edit() error {
	sqlstr := "UPDATE blog.post SET title=?,content=? WHERE id = ?"

	_, err := initdb.Db.Exec(sqlstr, p.Title, p.Content, p.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Post) Delete() error {
	sqlstr := "DELETE FROM blog.post WHERE id = ?"

	_, err := initdb.Db.Exec(sqlstr, p.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetPosts 根据用户获取全部的文章
func GetPosts(author int64) []Post {
	var posts []Post

	stmtOut, err := initdb.Db.Prepare("SELECT id,author,title,content,pv,create_time FROM blog.post WHERE author = ?")
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
		rows.Scan(&post.ID, &post.Author, &post.Title, &post.Content, &post.PV, &post.CreateTime)

		fmt.Println(*post)

		posts = append(posts, *post)
	}

	return posts
}

// GetPostByID 根据id获取文章详情
func GetPostByID(id int) (post Post) {
	stmtOut, err := initdb.Db.Prepare("SELECT id,author,title,content,pv,create_time FROM blog.post WHERE id = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	if err := stmtOut.QueryRow(id).Scan(&post.ID, &post.Author, &post.Title, &post.Content, &post.PV, &post.CreateTime); err != nil {
		panic(err.Error())
	}
	return
}

// IncPV 浏览量 + 1
func (p *Post) IncPV() error {
	sqlstr := "UPDATE blog.post SET pv=pv+1 WHERE id = ?"

	_, err := initdb.Db.Exec(sqlstr, p.ID)
	if err != nil {
		return err
	}
	return nil
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
