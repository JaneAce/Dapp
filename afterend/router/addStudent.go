package router

import (
	"dapp/afterend/db"
	"dapp/afterend/eth"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"strconv"
)

func AddStudent(c *gin.Context)  {
	schoolName := c.PostForm("schoolName")
	studentLevel := c.PostForm("studentLevel")
	StudentName := c.PostForm("studentName")
	sfId := c.PostForm("StudentId")
	id,_ :=strconv.ParseInt(sfId,10,64)
	StudentId := big.NewInt(id)
	StudentSex := c.PostForm("StudentSex")
	system,_ :=strconv.ParseInt(c.PostForm("StudentSystem"),10,64)
	StudentSystem := big.NewInt(system)
	studentStartTime:=c.PostForm("studentStartTime")
	startTime,_ := strconv.ParseInt(studentStartTime,10,64)
	StudentStartTimes := big.NewInt(startTime)
	StudnetSpecial := c.PostForm("StudnetSpecial")
	collegeName := c.PostForm("collegeName")
	StudentSchool := c.PostForm("StudentSchool")
	accounts,err := eth.NewAcc("123456")
	user,err := db.SchoolAP(schoolName)
	if err != nil {
		log.Panicln(err)
	}
	schoolAccount := user[0].UserAccount
	schoolPw := user[0].UserPassword
	fmt.Println("创建学生查询的账号和密码",schoolAccount,schoolPw)
	Txopts :=eth.GetTxopts(schoolAccount,"F:\\Dappgeth\\dappdata\\keystore",schoolPw)
	transfermessage, err := eth.CreateStudent(schoolName,StudentName,StudentSex,StudnetSpecial,studentLevel,collegeName,common.HexToAddress(accounts),StudentId,StudentSystem,StudentStartTimes,Txopts)
	if err != nil{
		fmt.Println("router创建学校的错误",err)
	}
	studentPid := eth.GetStudentPid()
	pid := studentPid.String()//转成string
	num, _ := strconv.Atoi(pid)
	dbsfId,_ := strconv.Atoi(sfId)
	db.AddStudent(StudentName,accounts,StudentSchool,StudnetSpecial,num,dbsfId)
	fmt.Println("获取的学院名称",collegeName)
	fmt.Println("创建学生放回的值",transfermessage)
	var student db.STUDENTS
	student.StudentName = StudentName
	student.StudentStarTime = int(startTime)
	student.StudentAccount = accounts
	student.StudentCollege =collegeName
	student.StudentId = string(id)
	student.StudentLevel = studentLevel
	student.StudentSchool = schoolName
	student.StudentSpecial = StudnetSpecial
	student.StudentSex = StudentSex
	student.StudentSystem = int(system)
	student.StudentPid = num
	respOK(c,student)
}
func SelectStudnet(c *gin.Context)  {
	studnetName := c.PostForm("studentName")
	student,err := db.SeStudent(studnetName)
	if err != nil {
		fmt.Println(err)
	}
	respOK(c,student)
}


