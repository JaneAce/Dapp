package router

import (
	"dapp/afterend/db"
	_ "dapp/afterend/middleware"
	"dapp/certsPhoto"
	"fmt"
	"github.com/gin-gonic/gin"
)
func DownLoadCert(c *gin.Context)  {
	certsName := c.PostForm("certsName")
	student,err := db.SeStudentAc(certsName)
	if err != nil {
		fmt.Println(err)
	}
	studentName := student[0].StudentName
	certsPhoto.Screenshot(studentName)
}
