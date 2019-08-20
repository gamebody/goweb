package model

import (
	"errors"

	"github.com/gamebody/goweb/initdb"
)

// Post 用户相关字段
type Post struct {
	Author  int64  `form:"author"`
	Title   string `form:"title"`
	Content string `form:"content"`
	PV      int    `form:"pv"`
}

// Save 储存用户
func (p *Post) Save() error {
	sqlstr := "insert into blog.post (author, title, content) values (?,?,?)"

	_, err := initdb.Db.Exec(sqlstr, p.Author, p.Title, p.Content)
	if err != nil {
		// panic(err.Error())
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
