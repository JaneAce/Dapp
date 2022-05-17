package router

import (
	"dapp/afterend/db"
	"dapp/afterend/eth"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateSchool(c *gin.Context)  {
	schoolName := c.PostForm("schoolName")
	ac := c.PostForm("account")
	user,err := db.Passwd(schoolName)
	Passwd := user[0].UserPassword
	if err != nil {
		fmt.Println(err)
	}
	eth.UnlockAc(ac,Passwd)
	Txopts :=eth.GetTxopts(ac,"F:\\Dappgeth\\dappdata\\keystore",Passwd)
	transfermessage, err :=eth.CreateSchool(schoolName,Txopts)
	if err != nil{
		fmt.Println("router创建学校的错误",err)
	}
	fmt.Println("创建学校放回的值",transfermessage)
	err1 :=db.IsJoin(schoolName)
	if err1 != nil{
		fmt.Println(err1)
	}
	var users []db.USER
	users,err = db.IsCreateSchool(schoolName)
	if err != nil{
		fmt.Println(err)
	}
	respOK(c,users)
}
