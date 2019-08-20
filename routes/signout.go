package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Signout 退出页面的路由
func Signout(router *gin.Engine) {
	// data := make(map[string]interface{})

	signout := router.Group("signout")

	signout.GET("", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Delete("user")
		session.AddFlash("登出成功！")
		session.Save()

		c.Redirect(http.StatusFound, "/")
	})

}
