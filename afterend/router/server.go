package router

import (
	"dapp/afterend/middleware"
	"github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	_ "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin" // 导入web服务框架
	"net/http"
	_ "net/http"
)

func Start(webPath string)(err error)  {
	// 1.创建路由
	r := gin.Default()
	r.Use(middleware.Cors())//默认跨域
	store := cookie.NewStore([]byte("secrect"))
	r.Use(sessions.Sessions("schoolAccount",store))
	r.POST("/reg", RegHandler)
	r.POST("/selectStudent",SelectStudnet)
	r.POST("/IsCreate",CreateSchool)
	r.POST("/createCollege",CreateCollege)
	r.POST("/collage",Collage)
	r.POST("/addSubject",AddSubject)
	r.GET("/reg",RegGet)
	r.POST("/creatSchool",RegGet)
	r.POST("/login",LoginHandler)
	r.POST("/newSchoolAc", RegHandler)
	r.POST("/addStudent",AddStudent)
	r.POST("/addScore",AddScore)
	r.POST("/createCerts",CreateCert)
	r.POST("/addScoreMg",AddScoreMg)
	r.POST("/certs",ShowCerts)
	r.POST("/regcom",RegCompany)
	r.GET("/company",isCreatCompany)
	r.POST("/createCompany",CreateCompany)
	r.POST("/companyLogin",CompanyLogin)
	r.POST("/downloadCert",DownLoadCert)
	r.POST("/uploadCert",UploadCert)
	r.POST("/loginStudent",StudentLogin)
	r.GET("/Student",Student)
	r.POST("/veirfy",Veirfy)
	r.GET("/background",BackSchoolHandler)
	r.GET("/backStudent",BackStudentHandler)
	r.GET("/BackCollage",BackCollageHandler)
	r.GET("/BackCompany",BackCompanyHandler)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(webPath)
	return
}
// 封装函数，用于向客户端返回正确信息
func respOK(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}

// 封装函数，用于向客户端返回错误消息
func respError(c *gin.Context,msg interface{}) {
	c.JSON(500, gin.H{
		"code": 1,
		"message": msg,
	})
}

func GetUrlArg(r *http.Request,name string)string{
	var arg string
	values := r.URL.Query()
	arg=values.Get(name)
	return arg
}
