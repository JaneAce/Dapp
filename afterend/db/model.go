package db

type USER struct {
	// 结构体成员tag，包含3种，一种是数据库，一种是json，一种是form表单
	UserID      int        `sql:"user_id" json:"userid" form:"userid"`                // 唯一标识的
	UserEmail   string     `sql:"user_email" json:"user_email" form:"user_email"`       // 用户邮箱
	UserPassword string `sql:"user_pwd" json:"user_password" form:"user_password"` // 用户密码
}

