package router

import (
	"dapp/afterend/middleware"
	"fmt"
	_ "github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin" // 导入web服务框架
	_ "net/http"
)

func Start(webPath string)(err error)  {
	// 1.创建路由
	r := gin.Default()
	r.Use(middleware.Cors())//默认跨域
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/ceshi", func(c *gin.Context) {
		message := c.Query("username" )
		fmt.Println(message)
	})
	r.GET("/",RegHandler)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(webPath)
	return
}
// 封装函数，用于向客户端返回正确信息
func respOK(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}

// 封装函数，用于向客户端返回错误消息
func respError(c *gin.Context,msg interface{}) {
	c.JSON(200, gin.H{
		"code": 1,
		"message": msg,
	})
}
