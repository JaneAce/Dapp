package router

import (
	"dapp/afterend/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func BackSchoolHandler(c *gin.Context)  {
	school,err :=db.BackSchool()
	if err != nil {
		fmt.Println(err)
	}
	respOK(c,school)
}

func BackStudentHandler(c *gin.Context)  {
	student,err := db.BackStudent()
	if err != nil{
		fmt.Println(err)
	}
	respOK(c,student)
}
func BackCollageHandler(c *gin.Context)  {
	collage,err := db.BackCollage()
	if err != nil {
		fmt.Println(err)
	}
	respOK(c,collage)
}
func BackCompanyHandler(c *gin.Context)  {
	company,err := db.BackCompany()
	if err != nil {
		fmt.Println(err)
	}
	respOK(c,company)
}