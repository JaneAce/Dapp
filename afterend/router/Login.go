package router

import (
	"dapp/afterend/db"
	"fmt"
	_ "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)
//登入操作
func LoginHandler(c *gin.Context)  {
	//session := sessions.Default(c)
	schoolName := c.PostForm("schoolName")
	passwd := c.PostForm("schoolPasswd")
	users,err := db.IsExitUser(schoolName)
	if err!=nil{
		fmt.Println("登录错误",err)
		return
	}
	dbSchoolName,_ := db.IsRegUser(schoolName)
	if dbSchoolName == nil{
		respOK(c,"该学校未注册请先注册")
		return
	}
	if passwd != users[0].UserPassword{
		respOK(c,"密码错误")
	}
	respOK(c,"success")
	//session.Set("schoolAccount",users[0].UserAccount)
	//session.Save()
}
func CompanyLogin(c *gin.Context)  {
	companyName := c.PostForm("companyName")
	passwd := c.PostForm("companyPasswd")
	users,err := db.CompanyPasswd(companyName)
	if err!=nil{
		fmt.Println("登录错误",err)
		return
	}
	dbSchoolName,_ := db.IsCompany(companyName)
	if dbSchoolName == nil{
		respOK(c,"该公司未注册请先注册")
		return
	}
	if passwd != users[0].CompanyPw{
		respOK(c,"密码错误")
	}
	respOK(c,"success")
}
func StudentLogin(c *gin.Context)  {
	studentAccount:=c.PostForm("studentAccount")
	user,err := db.StudentLogin(studentAccount)
	if err != nil {
		fmt.Println(err)
	}
	if user == nil{
		respOK(c,"该账户未注册")
	}
	respOK(c,"登入成功")
}
