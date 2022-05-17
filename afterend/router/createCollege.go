package router

import (
	"dapp/afterend/db"
	"dapp/afterend/eth"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateCollege(c *gin.Context)  {
	collegeName := c.PostForm("collegeName")
	schoolName := c.PostForm("schoolName")
	Passwd := c.PostForm("Passwd")
	user,err := db.SchoolAP(schoolName)
	if err != nil {
		fmt.Println(err)
	}
	schoolAc := user[0].UserAccount
	schoolpw := user[0].UserPassword
	account,err := eth.NewAcc(Passwd)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("创建学院查询到的学校密码和账号",schoolAc,schoolpw)
	Txopts :=eth.GetTxopts(schoolAc,"F:\\Dappgeth\\dappdata\\keystore",schoolpw)
	err,transfermessage:=eth.CreateCollege(Passwd,collegeName,common.HexToAddress(account),Txopts)
	if err != nil{
		fmt.Println("router创建学院的错误",err)
	}
	fmt.Println("创建学院的返回值",transfermessage)
	collegePid := eth.GetCollagePid()
	pid := collegePid.String()//转成string
	num, err := strconv.Atoi(pid)
	err1 := db.InfoCollage(collegeName,account,Passwd,schoolName,num)
	if err1 != nil {
		fmt.Println(err)
	}

	var users db.COLLEGE
	users.CollegeAccount = account
	users.CollegeName = collegeName
	respOK(c,users)
}
