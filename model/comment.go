package model

import "github.com/gamebody/goweb/initdb"

// Comment 评论相关字段
type Comment struct {
	Author  int
	Content string
	PostID  int
}

// Save 保存评论
func (c *Comment) Save() error {
	sqlstr := "insert into blog.comment (author, content, postid) values (?,?,?)"

	_, err := initdb.Db.Exec(sqlstr, c.Author, c.Content, c.PostID)
	if err != nil {
		panic(err.Error())
	}
	return nil
}
