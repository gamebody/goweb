package model

import (
	"errors"
	"fmt"

	"github.com/gamebody/goweb/lib"

	"github.com/gamebody/goweb/initdb"
)

// User 用户相关字段
type User struct {
	ID         int    `form:"id"`
	Name       string `form:"name" binding:"required,min=1,max=10"`
	Gender     string `form:"gender" binding:"required"`
	Bio        string `form:"bio" binding:"required"`
	Password   string `form:"password" binding:"required"`
	Repassword string `form:"repassword" binding:"required"`
	Avatar     string
}

// Save 储存用户
func (user *User) Save() error {
	password := lib.Scrypt(user.Password)

	sqlstr := "insert into blog.user (name, gender, bio, password, avatar) values (?,?,?,?,?)"

	_, errr := initdb.Db.Exec(sqlstr, user.Name, user.Gender, user.Bio, password, user.Avatar)
	if errr != nil {
		fmt.Println("数据库出错了")
		return errr
	}
	return nil

	// id, err := res.LastInsertId()
	// if err != nil {
	// 	return err
	// }
}

// CheckUserInTable 根据用户名查询
func (user *User) CheckUserInTable(name string) error {
	var n int
	sqlstr := "SELECT COUNT(*) AS n FROM blog.user WHERE name = ?"
	stmtOut, err := initdb.Db.Prepare(sqlstr)
	if err != nil {
		return err
	}
	defer stmtOut.Close()

	err = stmtOut.QueryRow(name).Scan(&n)
	if err != nil {
		return err
	}

	if n != 0 {
		return errors.New("用户已存在")
	}
	return nil
}

// SearchUser 根据name查询
func (user *User) SearchUser() (User, error) {
	var user2 User

	stmtOut, err := initdb.Db.Prepare("SELECT id,name,password,bio,avatar,gender FROM blog.user WHERE name = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	if err := stmtOut.QueryRow(user.Name).Scan(&user2.ID, &user2.Name, &user2.Password, &user2.Bio, &user2.Avatar, &user2.Gender); err != nil {
		return user2, err
	}

	return user2, nil
}

// Check 校验字段
func (user *User) Check() error {
	if user.Name == "" || len(user.Name) > 10 {
		return errors.New("请输入正确的用户名")
	}

	if user.Password == "" || len(user.Password) < 6 {
		return errors.New("密码不能小于6位")
	}

	if user.Bio == "" || len(user.Bio) > 30 {
		return errors.New("个人简介字符在1~30之间")
	}

	if user.Password != user.Repassword {
		return errors.New("两次输入密码不一致")
	}

	if user.Gender == "" {
		return errors.New("请输入性别")
	}

	if user.Avatar == "" {
		return errors.New("请选择头像")
	}

	return nil
}
