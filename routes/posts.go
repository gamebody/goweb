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

// Posts  文章相关路由
func Posts(router *gin.Engine) {
	// data := make(map[string]interface{})

	posts := router.Group("posts")

	posts.GET("/", middleware.LoginPass(), func(c *gin.Context) {
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

	posts.GET("/info/:postID", middleware.LoginPass(), func(c *gin.Context) {
		var post model.Post
		var comments []model.Comment

		if postID, err := strconv.Atoi(c.Param("postID")); err != nil {
			panic(err.Error())
		} else {
			post = model.GetPostByID(postID)
			comments = model.GetCommentsByPostID(postID)
		}

		if err := post.IncPV(); err != nil {
			panic(err.Error())
		}

		var user model.User
		session := sessions.Default(c)
		msgs := session.Flashes()
		data := session.Get("user").([]byte)
		json.Unmarshal(data, &user)
		fmt.Println(user.Avatar)
		session.Save()

		c.HTML(http.StatusOK, "post.html", gin.H{
			"title":    user.Name,
			"flash":    msgs,
			"user":     user,
			"post":     post,
			"comments": comments,
		})
	})

	posts.GET("/edit/:postID", middleware.LoginPass(), func(c *gin.Context) {
		var post model.Post
		postID := c.Param("postID")

		if postID, err := strconv.Atoi(postID); err != nil {
			panic(err.Error())
		} else {
			post = model.GetPostByID(postID)
		}

		fmt.Println(post)

		var user model.User
		session := sessions.Default(c)
		msgs := session.Flashes()
		data := session.Get("user").([]byte)
		json.Unmarshal(data, &user)
		fmt.Println(user.Avatar)
		session.Save()

		c.HTML(http.StatusOK, "edit.html", gin.H{
			"title": user.Name,
			"flash": msgs,
			"user":  user,
			"post":  post,
		})
	})

	posts.POST("/edit/:postID", middleware.LoginPass(), func(c *gin.Context) {
		postID := c.Param("postID")

		var user model.User
		session := sessions.Default(c)
		data := session.Get("user").([]byte)
		json.Unmarshal(data, &user)

		var post model.Post
		post.Author = user.ID
		post.Content = c.PostForm("content")
		post.Title = c.PostForm("title")

		if postID, err := strconv.Atoi(postID); err != nil {
			panic(err.Error())
		} else {
			post.ID = postID
		}

		if err := post.Check(); err != nil {
			session.AddFlash(err.Error())
			session.Save()
			c.Redirect(http.StatusFound, "/posts/edit/"+postID)
			return
		}

		if err := post.Edit(); err != nil {
			panic(err.Error())
		}

		session.AddFlash("更新成功！")
		session.Save()

		c.Redirect(http.StatusFound, "/posts/edit/"+postID)
	})

	posts.GET("/remove/:postID", middleware.LoginPass(), func(c *gin.Context) {
		postID := c.Param("postID")

		var user model.User
		session := sessions.Default(c)
		data := session.Get("user").([]byte)
		json.Unmarshal(data, &user)

		var post model.Post
		post.Author = user.ID
		post.Content = c.PostForm("content")
		post.Title = c.PostForm("title")

		if postID, err := strconv.Atoi(postID); err != nil {
			panic(err.Error())
		} else {
			post.ID = postID
		}

		if err := post.Delete(); err != nil {
			session.AddFlash(err.Error())
			session.Save()
			c.Redirect(http.StatusFound, "/posts")
			return
		}

		session.AddFlash("删除成功！")
		session.Save()

		c.Redirect(http.StatusFound, "/posts")
	})
}
