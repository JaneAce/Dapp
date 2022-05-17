package router

import (
	"dapp/afterend/db"
	"dapp/afterend/eth"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func CreateCert(c *gin.Context)  {
	studentName := c.PostForm("certsStudentName")
	certsStudentAccount := c.PostForm("certsStudentAccount")
	studentAccount := common.HexToAddress(certsStudentAccount)
	certsStudentCET4 := c.PostForm("certsStudentETC4")
	cet4,_ := strconv.ParseInt(certsStudentCET4,10,64)
	certsStudentall := c.PostForm("certsStudentall")
	point,_ := strconv.ParseInt(certsStudentall,10,64)
	certStudentGraduation := c.PostForm("certStudentGraduation")
	graduation,_ := strconv.ParseInt(certStudentGraduation,10,64)
	schoolName := c.PostForm("schoolName")
	user,err := db.SchoolAP(schoolName)
	if err != nil {
		log.Panicln(err)
	}
	schoolAccount := user[0].UserAccount
	schoolPw := user[0].UserPassword
	eth.UnlockAc(schoolAccount,schoolPw)
	Txopts :=eth.GetTxopts(schoolAccount,"F:\\Dappgeth\\dappdata\\keystore",schoolPw)
	certTran,err := eth.CreateCerts(studentName,studentAccount,cet4,point,graduation,Txopts)
	if err != nil {
		log.Fatal("生成证书",err)
	}
	certHash := certTran.Hash().Hex()
	certPid := eth.GetCertPid()
	pid := certPid.String()//转成string
	num, err := strconv.Atoi(pid)

	db.Cert(studentName,certsStudentAccount,certHash,int(graduation),int(cet4),int(point),num)
	studentMg,err := db.SeScore(certsStudentAccount)
	var certs db.CERTS
	certs.CertHash = certTran.Hash().Hex()
		certs.Subject = studentMg
	respOK(c,certs)
}
func ShowCerts(c *gin.Context)  {
	var certs db.CERTS
	studentAccount := c.PostForm("studentAccount")
	studentMg,err := db.SeScore(studentAccount)
	if err != nil {
		fmt.Println(err)
	}
	certHash,_ := db.SeCerts(studentAccount)
	student,_ := db.SeStudentAc(studentAccount)
	certs.Students.StudentSchool = student[0].StudentSchool
	certs.Students.StudentName = student[0].StudentName
	certs.Subject = studentMg
	certs.CertHash = certHash[0].CertHash
	certs.Students.StudentAccount = studentAccount
	fmt.Println("查询证书的哈希值",certHash[0].CertHash)
	respOK(c,certs)
}
