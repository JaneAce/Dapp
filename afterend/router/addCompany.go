package router

import (
	"dapp/afterend/db"
	"dapp/afterend/eth"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RegCompany(c *gin.Context)  {
	companyName := c.PostForm("companyName")
	passwd := c.PostForm("companyPasswd")
	fmt.Println("公司名称",companyName)
	dbCompangName,_ := db.IsRegCompany(companyName)
	if dbCompangName != nil{
		fmt.Println("注册路由的数据",dbCompangName[0].CompanyName)
		respOK(c,"该公司已被注册")
		return
	}
	_,err1 := db.InfoCompany(companyName,passwd)
	if err1!= nil {
		fmt.Println("注册", err1)
	}
	users,err := db.IsCreateSchool(companyName)
	if err != nil{
		fmt.Println(err)
	}
	respOK(c,users)
}
func isCreatCompany(c *gin.Context)  {
	// 声明接收的变量
	var  users db.USER
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	userurl:=c.Query("user")

	user,err := db.SeCompany(userurl)
	users.UserIsJoin = user[0].CompanyIsJoin
	users.UserAccount= user[0].CompanyAccount
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("url",userurl)
	fmt.Println("secompany数据库获取的值",users)
	respOK(c,users)
}
func CreateCompany(c *gin.Context)  {
	companyName := c.PostForm("companyName")
	ac := c.PostForm("account")
	user,err := db.CompanyPasswd(companyName)
	Passwd := user[0].CompanyPw
	if err != nil {
		fmt.Println(err)
	}
	eth.UnlockAc(ac,Passwd)
	Txopts :=eth.GetTxopts(ac,"F:\\Dappgeth\\dappdata\\keystore",Passwd)
	transfermessage, err :=eth.CreateSchool(companyName,Txopts)
	if err != nil{
		fmt.Println("router创建学校的错误",err)
	}
	fmt.Println("创建学校放回的值",transfermessage)
	err1 :=db.IsJoinCompany(companyName)
	if err1 != nil{
		fmt.Println(err1)
	}
	var users []db.COMPANY
	users,err = db.SeCompany(companyName)
	if err != nil{
		fmt.Println(err)
	}
	respOK(c,users)
}
