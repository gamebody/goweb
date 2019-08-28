package model

import (
	"fmt"

	"github.com/gamebody/goweb/initdb"
)

// Comment 评论相关字段
type Comment struct {
	ID         int
	PostID     int
	Author     int
	Content    string
	CreateTime string
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

// GetCommentsByPostID 根据文章id获取评论
func GetCommentsByPostID(postID int) (comments []Comment) {

	stmtOut, err := initdb.Db.Prepare("SELECT id,author,postid,content,create_time FROM blog.comment WHERE postid = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(postID)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		comment := &Comment{}
		rows.Scan(&comment.ID, &comment.Author, &comment.PostID, &comment.Content, &comment.CreateTime)

		fmt.Println(*comment)

		comments = append(comments, *comment)
	}

	return
}
