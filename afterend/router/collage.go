package router

import (
	"dapp/afterend/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Collage(c *gin.Context)  {
	schoolName := c.PostForm("schoolName")
	collageName := c.PostForm("collageName")
	fmt.Println("前台传到后台的数据",schoolName,collageName)
	collage,err := db.SeCollage(schoolName,collageName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(collage)
	respOK(c,collage)
}
