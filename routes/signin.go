package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"lib"
	"middleware"
	"model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func checkLoginForm(user model.User) error {
	if len(user.Name) == 0 {
		return errors.New("请填写用户名")
	}
	if len(user.Name) < 6 || len(user.Name) > 20 {
		return errors.New("用户名6-20位字符")
	}
	if len(user.Password) == 0 {
		return errors.New("请填写密码")
	}
	if len(user.Password) < 6 {
		return errors.New("密码不能少于6位")
	}
	return nil
}

// Signin  登录的路由
func Signin(router *gin.Engine) {
	// data := make(map[string]interface{})

	signin := router.Group("signin")

	signin.Use(middleware.NotLoginPass())

	signin.GET("", func(c *gin.Context) {
		session := sessions.Default(c)
		msgs := session.Flashes()
		session.Save()

		c.HTML(http.StatusOK, "signin.html", gin.H{
			"title": c.MustGet("hello"),
			"flash": msgs,
		})
	})

	signin.POST("", func(c *gin.Context) {
		session := sessions.Default(c)
		var user model.User
		var err error

		user.Name = c.PostForm("name")
		user.Password = c.PostForm("password")

		password := c.PostForm("password")

		if err := checkLoginForm(user); err != nil {
			session.AddFlash(err.Error())
			session.Save()
			c.Redirect(http.StatusFound, "/signin")
			return
		}
		if user, err = user.SearchUser(); err != nil {
			session.AddFlash(err.Error())
			session.Save()
			c.Redirect(http.StatusFound, "/signin")
			return
		}

		if user.Password != lib.Scrypt(password) {
			session.AddFlash("用户名或密码错误")
			session.Save()
			c.Redirect(http.StatusFound, "/signin")
			return
		}
		// 储存session,保持登录状态
		data, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
		}

		session.AddFlash("登陆成功！")
		session.Set("user", data)
		session.Save()

		c.Redirect(http.StatusFound, "/")
		return
	})

}
