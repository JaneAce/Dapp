package router

import (
	"dapp/afterend/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Student(c *gin.Context)  {
	studentAccount := c.Query("ac")
	student,err := db.SeStudentAc(studentAccount)
	if err != nil{
		fmt.Println(err)
	}
	respOK(c,student)
}
