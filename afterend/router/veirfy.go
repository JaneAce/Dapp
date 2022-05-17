package router

import (
	"dapp/afterend/db"
	"dapp/afterend/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Veirfy(c *gin.Context)  {
	certPath := c.PostForm("certsName")
	resultHash,resultAc := middleware.Recongntion(certPath)
	dbHash,_ :=db.SeCerts(resultAc)
	fmt.Println("数据库",dbHash[0].CertHash)
	fmt.Println("证书哈希值",resultHash)
	if dbHash[0].CertHash==resultHash{
		respOK(c,"true")
	}
	respOK(c,"false")
}
