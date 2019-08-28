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
	Name       string
	Avatar     string
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

	stmtOut, err := initdb.Db.Prepare("SELECT c.id,author,postid,content,create_time,name,avatar FROM blog.comment c left join blog.user u ON c.author = u.id WHERE c.postid=?")
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
		rows.Scan(&comment.ID, &comment.Author, &comment.PostID, &comment.Content, &comment.CreateTime, &comment.Name, &comment.Avatar)

		fmt.Println(*comment)

		comments = append(comments, *comment)
	}

	return
}
