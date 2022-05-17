package router

import (
	"dapp/afterend/db"
	"dapp/afterend/eth"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddSubject(c *gin.Context)  {
	subjectName := c.PostForm("subjectName")
	schoolName := c.PostForm("schoolName")
	collageName :=c.PostForm("collageName")
	collage,err := db.SeCollage(schoolName,collageName)
	collageAc := collage[0].CollegeAccount
	collagePw := collage[0].CollegePassword
	fmt.Println("添加科目查询到学院的账号和密码",collageAc,collagePw,subjectName)
	Txops:= eth.GetTxopts(collageAc,"F:\\Dappgeth\\dappdata\\keystore",collagePw)
	transformMg,err := eth.AddSubject(subjectName,Txops)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("创建科目的返回值",transformMg)
	db.AddSubject(subjectName)
	var subjectMg db.SUBJECT
	subjectMg.SubjectName = subjectName
	fmt.Println("添加科目的数据",subjectMg)
	respOK(c,subjectMg)
}
