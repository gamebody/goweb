package routes

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gamebody/goweb/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type flash struct {
	Name    string
	Message string
}

// Signup 注册页面的路由
func Signup(router *gin.Engine) {
	// data := make(map[string]interface{})

	signup := router.Group("signup")

	signup.GET("", func(c *gin.Context) {
		session := sessions.Default(c)
		msgs := session.Flashes()
		session.Clear()
		session.Save()

		c.HTML(http.StatusOK, "signup.html", gin.H{
			"title": c.MustGet("hello"),
			"flash": msgs,
		})
	})

	signup.POST("", func(c *gin.Context) {
		session := sessions.Default(c)
		var user model.User
		var saveFileName string

		user.Name = c.PostForm("name")
		user.Gender = c.PostForm("gender")
		user.Bio = c.PostForm("bio")
		user.Password = c.PostForm("password")
		user.Repassword = c.PostForm("repassword")

		file, err := c.FormFile("avatar")
		if err != nil {
			fmt.Println(err.Error())
		}
		if file != nil {
			// 上传文件
			saveFileName = string(time.Now().Format("20060102150405")) + "_" + file.Filename
			user.Avatar = saveFileName
			err = c.SaveUploadedFile(file, "./public/imgs/"+saveFileName)
		}

		if err != nil {
			panic(err.Error())
		}

		if err := user.Check(); err != nil {
			// 校验失败，删除上传的文件
			if err := os.Remove("./public/imgs/" + saveFileName); err != nil {
				fmt.Println(err.Error())
			}

			// session flash 提示
			session.AddFlash(err.Error())
			session.Save()

			c.Redirect(http.StatusFound, "/signup")
			return
		}

		if err := user.CheckUserInTable(user.Name); err != nil {
			fmt.Println(err)
			session.AddFlash("用户名已存在")
			session.Save()
			c.Redirect(http.StatusFound, "/signup")
			return
		}

		if err := user.Save(); err != nil {
			if err := os.Remove("./public/imgs/" + saveFileName); err != nil {
				fmt.Println(err.Error())
			}
			session.AddFlash(err.Error())
			session.Save()
			c.Redirect(http.StatusFound, "/signup")
			return
		}

		session.AddFlash("恭喜您，注册成功！")
		session.Save()
		c.Redirect(http.StatusFound, "/signin")
		return
	})

}
