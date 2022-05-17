package router

import (
	"dapp/afterend/db"
	"dapp/afterend/eth"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
	"strconv"
)

func AddScore(c *gin.Context)  {
	score := c.PostForm("score")
	score_,_ :=strconv.ParseInt(score,10,64)
	scores := big.NewInt(score_)
	subject := c.PostForm("subject")
	StudentAc := c.PostForm("studentAc")
	collegeName :=c.PostForm("collageName")
	schoolName := c.PostForm("schoolName")
	collage,err :=db.SeCollage(schoolName,collegeName)
	if err != nil {
		fmt.Println(err)
	}
	eth.UnlockAc(collage[0].CollegeAccount,collage[0].CollegePassword)
	fmt.Println("添加成绩查询到学院账号和密码",collage[0].CollegeAccount,collage[0].CollegePassword)
	fmt.Println("添加成绩前端传过来的数据",collage[0].CollegeAccount,collage[0].CollegePassword)
	fmt.Println("添加的科目名称和成绩",subject,score)
	Txopts :=eth.GetTxopts(collage[0].CollegeAccount,"F:\\Dappgeth\\dappdata\\keystore",collage[0].CollegePassword)
	transfermessage, err := eth.AddScore(common.HexToAddress(StudentAc),subject,scores,Txopts)
	if err != nil{
		fmt.Println("router添加学车成绩的错误",err)
	}
	scorePid := eth.GetStudentPid()
	pid := scorePid.String()//转成string
	num, _ := strconv.Atoi(pid)
	student,err := db.SeStudentAc(StudentAc)
	if err != nil {
		fmt.Println(err)
	}
	db.AddScore(student[0].StudentName,subject,StudentAc,int(score_),num)
	fmt.Println("添加学生成绩的值",transfermessage)
	var subjects db.SUBJECT
	subjects.SubjectName = subject
	subjects.Score = int(score_)
	respOK(c,subjects)
}

//获取添加成绩页面数据
func AddScoreMg(c *gin.Context){
	studentAc := c.PostForm("studentAccount")
	studentMg,err := db.SeScore(studentAc)
	fmt.Println("前端传来的学生账号",studentAc)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("添加成绩页面数据库查询的数据",studentMg[0])
	respOK(c,studentMg)
}

