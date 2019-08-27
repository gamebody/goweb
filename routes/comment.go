package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gamebody/goweb/middleware"
	"github.com/gamebody/goweb/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Comment  评论相关路由
func Comment(router *gin.Engine) {
	// data := make(map[string]interface{})

	comment := router.Group("comment")

	comment.GET("/", middleware.LoginPass(), func(c *gin.Context) {
		var user model.User
		session := sessions.Default(c)
		msgs := session.Flashes()
		data := session.Get("user").([]byte)
		json.Unmarshal(data, &user)

		posts := model.GetPosts(user.ID)

		session.Save()

		c.HTML(http.StatusOK, "posts.html", gin.H{
			"title": c.MustGet("hello"),
			"flash": msgs,
			"user":  user,
			"posts": posts,
		})
	})

	comment.POST("/", middleware.LoginPass(), func(c *gin.Context) {
		var user model.User
		session := sessions.Default(c)
		data := session.Get("user").([]byte)
		json.Unmarshal(data, &user)

		var comment model.Comment
		if postID, err := strconv.Atoi(c.PostForm("postid")); err != nil {
			panic(err.Error())
		} else {
			comment.PostID = postID
		}
		if author, err := strconv.Atoi(c.PostForm("author")); err != nil {
			panic(err.Error())
		} else {
			comment.Author = author
		}
		comment.Content = c.PostForm("content")

		// if err := post.Check(); err != nil {
		// 	session.AddFlash(err.Error())
		// 	session.Save()
		// 	c.Redirect(http.StatusFound, "/posts/create")
		// 	return
		// }

		if err := comment.Save(); err != nil {
			fmt.Println(err.Error())
		}

		session.AddFlash("恭喜您，评论成功！")
		session.Save()
		c.Redirect(http.StatusFound, "/posts/info/"+c.PostForm("postid"))
		return

	})

}
