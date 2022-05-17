package db

import (
	"dapp/afterend/eth"
	"database/sql" // 标准库，用于操作数据库
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 全局数据库对象
var db *sql.DB

//初始化数据库
func Init(dataBase string) (err error) {
	// 打开mysql数据库
	db, err = sql.Open("mysql", dataBase)
	if err != nil {
		panic(err.Error())
	}

	db.SetMaxOpenConns(1000)

	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		panic(err.Error())
	}

	return
}
//插入用户数据
func Reg(username,pwd string)(ac string,err error){
	//创建新账户
	account,err := eth.NewAcc(pwd)
	if err !=nil {
		fmt.Println("",err)
	}
	_, err = db.Exec("INSERT INTO table_user(table_user.username,table_user.`password`,table_user.account,table_user.isJoin) VALUES(?,?,?,?)",username,pwd,account,0)
	if err != nil {
		fmt.Println(err)
	}
	return account,nil
}

//插入是否创建学校
func IsJoin(username string)(err error)  {
	_,err = db.Exec("update table_user set table_user.isJoin=1 WHERE table_user.username =?",username)
	if err != nil{
		fmt.Println(err)
		return nil
	}
	return
}
//查询用户密码
func Passwd(username string)(users []USER,err error){
	rows,err := db.Query("SELECT table_user.`password` FROM table_user WHERE table_user.username=?",username)
	if err != nil{
		fmt.Println(err)
	}
	for rows.Next(){
		var user USER
		err = rows.Scan(&user.UserPassword)
		if err != nil{
			fmt.Println(err)
		}
		users = append(users,user)

	}
	return
}
//查询学校是否存在
func IsExitSchool(username string)(users []USER,err error){
	rows,err := db.Query("SELECT table_user.username FROM table_user WHERE table_user.username = ?",username)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next(){
		var user USER
		err = rows.Scan(&user.UserName)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users,user)
	}
	return
}

//查询是否创建学校
func IsCreateSchool(username string)(users []USER,err error)  {
	rows,err := db.Query("SELECT table_user.isJoin,table_user.account,table_user.password,table_user.username FROM table_user WHERE table_user.username=?",username)
	if err != nil{
		fmt.Println(err)
	}
	for rows.Next() {
		var user USER
		err = rows.Scan(&user.UserIsJoin,&user.UserAccount,&user.UserPassword,&user.UserName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
		fmt.Println("创建学校查询数据的数据",user)
	}
	return
}
//查询该用户名是否被注册
func IsRegUser(username string)(users []USER, err error)  {
	rows,err := db.Query("SELECT table_user.username FROM table_user WHERE table_user.username = ?",username)
	if err != nil {
		 fmt.Println(err)
	}
	for rows.Next(){
		var user USER
		err = rows.Scan(&user.UserName)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users,user)
		fmt.Println("用户名是否存在数据库查询数据返回",users)
	}
	return
}
//查询用户是否存在
func IsExitUser(schoolName string)(users []USER,err error)  {
	rows,err := db.Query("SELECT table_user.account,table_user.password FROM table_user WHERE table_user.username=?",schoolName)
	if err !=nil {
		return nil, err
	}
	// 按行读取数据到结构体成员中
	for rows.Next() {
		var user USER
		err = rows.Scan(&user.UserAccount,&user.UserPassword)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
		fmt.Println(users[0].UserPassword,users[0].UserAccount)
	}
	return
}
//登录
func Login(username string)(users []USER,err error)  {
	rows,err := db.Query("SELECT table_user.`password` FROM table_user WHERE table_user.username = ?",username)
	if err !=nil{
		return nil,err
	}
	for rows.Next(){
		var user USER
		err = rows.Scan(&user.UserPassword)
		if err != nil{
			return nil,err
		}
		users = append(users,user)

	}
	return
}

func RegGet(username string)(users []USER,err error)  {
	rows,err := db.Query("SELECT table_user.account,table_user.isJoin FROM `table_user` WHERE table_user.username = ?",username)
	if err !=nil{
		return nil,err
	}
	for rows.Next(){
		var user USER
		err = rows.Scan(&user.UserAccount,&user.UserIsJoin)
		if err != nil{
			return nil,err
		}
		users = append(users,user)
	}
	return
}
//查询学校账号密码
func SchoolAP(username string)(users []USER,err error)  {
	rows,err := db.Query("SELECT table_user.account,table_user.password FROM `table_user` WHERE table_user.username = ?",username)
	if err !=nil{
		return nil,err
	}
	for rows.Next(){
		var user USER
		err = rows.Scan(&user.UserAccount,&user.UserPassword)
		if err != nil{
			return nil,err
		}
		users = append(users,user)
	}
	return
}
//插入学院信息
func InfoCollage(collageName, collageAccount, collagePasswd, collageSchoolName string,collagePid int) (err error) {
	_, err = db.Exec("INSERT INTO table_collage(table_collage.collageAccount,table_collage.collageName,table_collage.collagePasswd,table_collage.collageSchoolName,table_collage.collagePid)VALUES(?,?,?,?,?)", collageAccount,collageName,collagePasswd,collageSchoolName,collagePid)
	if err !=nil{
		fmt.Println(err)
		return err
	}
	return nil
}
//查询学院账号
func SeCollage(schoolName,collageName string) (collages []COLLEGE,err error) {
	rows,err := db.Query("SELECT * FROM `table_collage` WHERE table_collage.collageName=? and table_collage.collageSchoolName=?",collageName,schoolName)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next(){
		var collage COLLEGE
		err = rows.Scan(&collage.CollageId,&collage.CollegeName,&collage.CollegeAccount,&collage.CollegePassword,&collage.CollagePid,&collage.SchoolName)
		if err != nil {
			fmt.Println(err)
		}
		collages = append(collages,collage)
		fmt.Println("数据库查询到的数据",collages)
	}
	return
}
//添加科目
func AddSubject(subjectName string)(err error){
	_, err = db.Exec("INSERT INTO `tab_subject`(tab_subject.subjectName) VALUES(?)", subjectName)
	if err != nil {
		log.Panicln("添加科目插入数据库的错误",err)
	}
	return
}
//添加学生
func AddStudent(studentNe,studentAc,studentSchool,studentSpecial string,studentPid,studentId int) (err error) {
	_,err = db.Exec("INSERT INTO table_student(studentName,StudentSchool,StudentAccount,StudentPid,StudentsfId,StudentSpecial) VALUES(?,?,?,?,?,?)",studentNe,studentSchool,studentAc,studentPid,studentId,studentSpecial)
	if err != nil {
		fmt.Println("数据库获取到插入学生的数据",studentPid,studentAc,studentNe)
		fmt.Println("把学生数据插入数据库错误",err)
	}
	return
}
//查询学生
func SeStudent(studnetName string) (students []STUDENTS,err error) {
	rows,err := db.Query("SELECT table_student.studentName,table_student.StudentAccount FROM table_student WHERE table_student.studentName = ?",studnetName)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next(){
		var student STUDENTS
		err = rows.Scan(&student.StudentName,&student.StudentAccount)
		if err != nil {
			fmt.Println(err)
		}
		students = append(students,student)
		fmt.Println("数据库查询到的数据",students)
	}
	return
}
func SeStudentAc(studentAc string) (students []STUDENTS,err error) {
	rows,err := db.Query("SELECT table_student.studentName,table_student.StudentAccount,table_student.StudentSchool FROM table_student WHERE table_student.StudentAccount = ?",studentAc)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next(){
		var student STUDENTS
		err = rows.Scan(&student.StudentName,&student.StudentAccount,&student.StudentSchool)
		if err != nil {
			fmt.Println(err)
		}
		students = append(students,student)
		fmt.Println("数据库查询到的数据",students)
	}
	return
}
//添加成绩
func AddScore(studentName,studentSubject,studentAccount string,score,scorePid int)(err error){
	_,err = db.Exec("INSERT INTO tab_score(tab_score.studentName,tab_score.studentSubject,tab_score.studentScore,tab_score.scorePid,tab_score.studentAccount) VALUES(?,?,?,?,?)",studentName,studentSubject,score,scorePid,studentAccount)
	if err != nil{
		fmt.Println(err)
	}
	return
}
 //查询学生的成绩单
 func SeScore(stdentAccount string)(students []SUBJECT,err error){
 	rows,err := db.Query("SELECT tab_score.studentSubject,tab_score.studentScore FROM `tab_score` WHERE tab_score.studentAccount=?",stdentAccount)
 	if err != nil {
 		fmt.Println(err)
	}
	for rows.Next(){
		var student SUBJECT
		err = rows.Scan(&student.SubjectName,&student.Score)
		if err != nil{
			fmt.Println(err)
		}
		students = append(students,student)
	}
	return
 }
//插入证书相关信息
func Cert(studentName,studentAccount,certHash string,graduation,cet4,credit,certPid int){
	_,err := db.Exec("INSERT INTO tab_cert(tab_cert.studentName,tab_cert.studentAccount,tab_cert.graduation,tab_cert.cet4,tab_cert.credit,tab_cert.certhash,tab_cert.certpid)VALUES(?,?,?,?,?,?,?)",studentName,studentAccount,graduation,cet4,credit,certHash,certPid)
	if err != nil{
		fmt.Println(err)
	}
	return
}
func SeCerts(studentAc string)(certs []CERTS,err error)  {
	rows,err := db.Query("SELECT tab_cert.certhash FROM `tab_cert` WHERE tab_cert.studentAccount = ?",studentAc)
	if err !=nil {
		fmt.Println(err)
	}
	for rows.Next(){
		var cert CERTS
		err = rows.Scan(&cert.CertHash)
		if err != nil {
			fmt.Println(err)
		}
		certs = append(certs,cert)
		fmt.Println("查询证书的hash值",certs)
	}
	return
}
//添加公司
func InfoCompany(companyName,companyPw string) (account string,err error){
	//创建新账户
	account,err = eth.NewAcc(companyPw)
	if err !=nil {
		fmt.Println("",err)
	}
	_, err = db.Exec("INSERT INTO tab_company(tab_company.company_name,tab_company.company_pw,tab_company.company_Account,tab_company.company_isJoin) VALUES(?,?,?,?)",companyName,companyPw,account,0)
	if err != nil {
		fmt.Println("注册公司数据库错误",err)
	}
	return account,nil
}
//查询该用户名是否被注册
func IsRegCompany(username string)(users []COMPANY, err error)  {
	rows,err := db.Query("SELECT tab_company.company_name FROM tab_company WHERE tab_company.company_name = ?",username)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next(){
		var user COMPANY
		err = rows.Scan(&user.CompanyName)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users,user)
		fmt.Println("公司是否存在数据库查询数据返回",users)
	}
	return
}
//查询公司信息
func SeCompany(companyName string)(company []COMPANY,err error){
	rows,err := db.Query("SELECT tab_company.company_name,tab_company.company_Account,tab_company.company_isJoin FROM `tab_company` WHERE tab_company.company_name = ?",companyName)
	if err !=nil{
		return nil,err
	}
	for rows.Next(){
		var user COMPANY
		err = rows.Scan(&user.CompanyName,&user.CompanyAccount,&user.CompanyIsJoin)
		if err != nil{
			return nil,err
		}
		company = append(company,user)
	}
	return
}
func CompanyPasswd(username string)(users []COMPANY,err error){
	rows,err := db.Query("SELECT tab_company.company_pw FROM tab_company WHERE tab_company.company_name = ?",username)
	if err != nil{
		fmt.Println(err)
	}
	for rows.Next(){
		var user COMPANY
		err = rows.Scan(&user.CompanyPw)
		if err != nil{
			fmt.Println(err)
		}
		users = append(users,user)

	}
	return
}

//插入是否创建学校
func IsJoinCompany(username string)(err error)  {
	_,err = db.Exec("update tab_company set tab_company.company_isJoin=1 WHERE tab_company.company_name =?",username)
	if err != nil{
		fmt.Println(err)
		return nil
	}
	return
}

//查询该公司是否被注册
func IsCompany(username string)(users []COMPANY, err error)  {
	rows,err := db.Query("SELECT tab_company.company_name FROM tab_company WHERE tab_company.company_name = ?",username)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next(){
		var user COMPANY
		err = rows.Scan(&user.CompanyName)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users,user)
		fmt.Println("用户名是否存在数据库查询数据返回",users)
	}
	return
}
func StudentLogin(studentAccount string)(students []STUDENTS,err error)  {
 rows,err := db.Query("SELECT table_student.StudentAccount FROM `table_student` WHERE table_student.StudentAccount=?",studentAccount)
 if err != nil {
 	fmt.Println(err)
 }
	for rows.Next() {
		var student STUDENTS
		err = rows.Scan(&student.StudentAccount)
		if err != nil {
			fmt.Println(err)
		}
		students = append(students,student)
	}
	return
}

//后台数据展示
func BackSchool() (schools []USER,err error) {
	rows,err := db.Query("SELECT * FROM `table_user`")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next(){
		var school USER
		err = rows.Scan(&school.UserID,&school.UserPassword,&school.UserName,&school.UserAccount,&school.UserIsJoin)
		if err != nil {
			fmt.Println(err)
		}
		schools =append(schools,school)
	}
	return

}
func BackStudent() (students []STUDENTS,err error) {
	rows,err := db.Query("SELECT * FROM `table_student`")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next(){
		var student STUDENTS
		err = rows.Scan(&student.StudentDbId,&student.StudentName,&student.StudentAccount,&student.StudentPid,&student.StudentSchool,&student.StudentId,&student.StudentSpecial)
		if err != nil {
			fmt.Println(err)
		}
		students = append(students,student)
	}
	return
}
func BackCollage() (collages []COLLEGE,err error) {
	rows,err := db.Query("SELECT * FROM `table_collage`")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next(){
		var collage COLLEGE
		err = rows.Scan(&collage.CollageId,&collage.CollegeName,&collage.CollegeAccount,&collage.CollegePassword,&collage.CollagePid,&collage.SchoolName)
		if err != nil {
			fmt.Println(err)
		}
		collages = append(collages,collage)
	}
	return
}
func BackCompany() (companys []COMPANY,err error) {
	rows,err := db.Query("SELECT * FROM `tab_company`")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next(){
		var company COMPANY
		err = rows.Scan(&company.CompanyId,&company.CompanyName,&company.CompanyAccount,&company.CompanyPw,&company.CompanyIsJoin)
		if err != nil {
			fmt.Println(err)
		}
		companys = append(companys,company)
	}
	return
}