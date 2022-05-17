package router

import (
	"dapp/afterend/db"
	_ "dapp/afterend/db"
	"fmt"
	_ "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)
//注册账号
func RegHandler(c *gin.Context)  {
	schoolName := c.PostForm("schoolName")
	passwd := c.PostForm("schoolPasswd")
	fmt.Println("学校名称",schoolName)
	dbSchoolName,_ := db.IsRegUser(schoolName)
	if dbSchoolName != nil{
		fmt.Println("注册路由的数据",dbSchoolName[0].UserName)
		respOK(c,"该学校已被注册")
		return
	}
	_,err1 := db.Reg(schoolName,passwd)
	if err1!= nil {
		fmt.Println("注册", err1)
	}
	users,err := db.IsCreateSchool(schoolName)
	if err != nil{
		fmt.Println(err)
	}
	respOK(c,users)
	}
func RegGet(c *gin.Context)  {
	// 声明接收的变量
	var  users db.USER
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	userurl:=c.Query("user")

	user,err := db.RegGet(userurl)
	users.UserIsJoin = user[0].UserIsJoin
	users.UserAccount= user[0].UserAccount
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("url",userurl)
	fmt.Println("regGet数据库获取的值",users)
	respOK(c,users)
}





