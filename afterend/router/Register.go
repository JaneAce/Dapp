package router

import (
	"dapp/afterend/db"
	_ "dapp/afterend/db"
	"github.com/gin-gonic/gin"
)

func RegHandler(c *gin.Context)  {
	userPassword := c.PostForm("user_password")
	userEmail := c.PostForm("user_email")
	result,_ := db.IsExitUser(userEmail)
	if result == nil {
		respError(c,"该用户已被注册")
	}
	if len(userEmail) == 0 {
		respError(c,"用户邮箱不能为空")
		return
	}
	if len(userPassword) == 0 {
		respError(c,"用户密码不能为空")
		return
	}
	if err := db.Reg(userEmail,userPassword); err != nil {
		respError(c,err)
		return
	}
	//c.HTML(http.StatusOK,"login.html",nil)
}

