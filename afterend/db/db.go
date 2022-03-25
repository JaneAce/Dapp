package db

import (
	"database/sql" // 标准库，用于操作数据库
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
func Reg(email,pwd string)(err error){
	_, err = db.Exec("INSERT INTO tab_user(tab_user.user_email,tab_user.usr_pwd) VALUES(?,?)", email, pwd)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return
}
//查询用户是否存在
func IsExitUser(email string)(users []USER,err error)  {
	rows,err := db.Query("SELECT tab_user.user_email FROM tab_user WHERE tab_user.user_email=?",email)
	if err !=nil {
		return nil, err
	}
	// 按行读取数据到结构体成员中
	for rows.Next() {
		var user USER
		err = rows.Scan(&user.UserEmail)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
		fmt.Println(users[0].UserEmail,rows)
	}
	return
}
//登入
func Login(email string)(users []USER,err error)  {
	rows,err := db.Query("SELECT tab_user.user_email,tab_user.usr_pwd FROM tab_user WHERE tab_user.user_email=?",email)
	if err !=nil{
		return nil,err
	}
	for rows.Next(){
		var user USER
		err = rows.Scan(&user.UserEmail,&user.UserPassword)
		if err != nil{
			return nil,err
		}
		users = append(users,user)

	}
	return
}


