package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gamebody/goweb/middleware"
	"github.com/gamebody/goweb/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Posts  文章相关路由
func Posts(router *gin.Engine) {
	// data := make(map[string]interface{})

	posts := router.Group("posts")

	posts.GET("/create", middleware.LoginPass(), func(c *gin.Context) {
		var user model.User
		session := sessions.Default(c)
		msgs := session.Flashes()
		data := session.Get("user").([]byte)
		json.Unmarshal(data, &user)
		fmt.Println(user.Avatar)
		session.Save()

		c.HTML(http.StatusOK, "create.html", gin.H{
			"title": c.MustGet("hello"),
			"flash": msgs,
			"user":  user,
		})
	})

	posts.POST("/create", middleware.LoginPass(), func(c *gin.Context) {
		var user model.User
		session := sessions.Default(c)
		data := session.Get("user").([]byte)
		json.Unmarshal(data, &user)

		var post model.Post
		post.Author = user.ID
		post.Content = c.PostForm("content")
		post.Title = c.PostForm("title")

		if err := post.Check(); err != nil {
			session.AddFlash(err.Error())
			session.Save()
			c.Redirect(http.StatusFound, "/posts/create")
			return
		}

		if err := post.Save(); err != nil {
			fmt.Println(err.Error())
		}

		session.AddFlash("恭喜您，发表成功！")
		session.Save()
		id := string(user.ID)
		c.Redirect(http.StatusFound, "/posts/"+id)
		return

	})

	posts.GET("/:id", middleware.LoginPass(), func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println(id)

		var user model.User
		session := sessions.Default(c)
		msgs := session.Flashes()
		data := session.Get("user").([]byte)
		json.Unmarshal(data, &user)
		fmt.Println(user.Avatar)
		session.Save()

		c.HTML(http.StatusOK, "posts.html", gin.H{
			"title": user.Name,
			"flash": msgs,
			"user":  user,
		})
	})
}
