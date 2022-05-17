package db

type USER struct {
	// 结构体成员tag，包含3种，一种是数据库，一种是json，一种是form表单
	UserID      int        `sql:"user_id" json:"userid" form:"userid"`                // 唯一标识的
	UserName string `sql:"user_name",json:"user_name"，from:"user_name"`
	UserPassword string `sql:"user_pwd" json:"user_password" form:"user_password"` // 用户密码
	UserAccount string `sql:"account" json:"user_account" form:"user_account"`
	UserIsJoin string `sql:"isJoin" json:"user_isJoin" form:"user_isJoin"`
}
type LOGINREG struct {
	User string `from:"username" json:"user" uri:"user"`
	Pssword string `form:"password" json:"password" uri:"password"`
}
type COLLEGE struct {
	CollageId string `sql:"collageId",json:"collageId",form:"collageId"`
	CollegeName string `sql:"collageName",json:"collegeName",form:"collegeName"`
	CollegeAccount string `sql:"collageAccount",json:"college_account",form:"collegeAccount"`
	CollegePassword string `sql:"collagePasswd",json:"college_password",form:"collegePasswd"`
	SchoolName string `sql:"collageSchoolName",json:"school_name",form:"schoolName"'`
	CollagePid string `sql:"collagePid",json:"collage_pid",form:"collagePid"`
}
type SUBJECT struct {
	SubjectName string `json:"subject_name",from:"subjectName"`
	Pid int `json:"pid",from:"pid"`
	Score int `json:"score" from:"score"`
}
type STUDENTS struct {
	StudentDbId int
	StudentName string `sql:"studentName" json:"student_name"`
	StudentSchool string
	StudentCollege string
	StudentAccount string `sql:"studentAccount" json:"student_account"`
	StudentSex string
	StudentId string
	StudentSpecial string
	StudentLevel string
	StudentSystem int
	StudentStarTime int
	Subject SUBJECT
	StudentPid int `sql:"studentPid" json:"student_pid"`
}
type CERTS struct {
	Students STUDENTS
	CertHash string
	Subject []SUBJECT
}
type COMPANY struct {
	CompanyId int
	CompanyName string
	CompanyAccount string
	CompanyIsJoin string
	CompanyPw string
}