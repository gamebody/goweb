package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Init 注册路由
func Init() *gin.Engine {

	router := gin.Default()

	router.Static("/public", "./public")
	router.LoadHTMLGlob("templates/*.html")

	// session
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.Use(func(c *gin.Context) {
		c.Set("hello", "小云云不小")
		c.Next()
	})

	router.GET("", func(c *gin.Context) {
		session := sessions.Default(c)
		msgs := session.Flashes()
		session.Save()

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello，world！",
			"flash": msgs,
		})
	})

	Signup(router)  // 注册相关
	Signin(router)  // 登录相关
	Signout(router) // 登出相关
	Posts(router)   // 登出相关

	return router
}
