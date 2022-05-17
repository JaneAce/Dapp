package router

import (
	"dapp/afterend/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UploadCert(c *gin.Context)  {
	certname := c.PostForm("certsName")
	result,hash:= middleware.Recongntion(certname)
	fmt.Println("识别到的结果",result,hash)
}

