package eth

type SCHOOL struct {
	// 结构体成员tag，包含3种，一种是json，一种是form表单
	SchoolName     int        `json:"userid" form:"userid"`
	SchoolAddress  string     ` json:"user_email" form:"user_email"`
}
