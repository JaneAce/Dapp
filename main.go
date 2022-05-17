package main

import (
	"dapp/afterend/db"
	_ "dapp/afterend/middleware"
	"dapp/afterend/router"
	_ "dapp/afterend/router"
	"fmt"
)

func main() {

	webPath := "127.0.0.1:8000"
	//数据库连接
	dbPath := "root:password@tcp(127.0.0.1:3306)/dapp?charset=utf8"

	err := db.Init(dbPath)

	if err != nil {
		fmt.Printf("初始化数据库错误:%v\n", err)
		return
	}

	err = router.Start(webPath)
}


