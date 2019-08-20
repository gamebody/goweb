package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginPass 登录通过
func LoginPass() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		data := session.Get("user")
		if data == nil {
			session.AddFlash("请登录")
			session.Save()
			c.Redirect(http.StatusFound, "/signin")
			c.Abort()
			return
		}
		c.Next()
	}
}

// NotLoginPass 没登录通过
func NotLoginPass() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		data := session.Get("user")
		if data != nil {
			session.AddFlash("已登录，切换账号请登出")
			session.Save()
			backURL := c.GetHeader("referer")
			c.Redirect(http.StatusFound, backURL)
			c.Abort()
			return
		}
		c.Next()
	}
}
