package database

import (
	"database/sql"
	"demo/conf"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var DB *sql.DB

// 定义一个初始化数据库的函数
func InitDB() (err error) {
	// 读取配置文件
	var c conf.Conf
	conf := c.GetConf()
	// DSN:Data Source Name
	dsn := conf.User + ":" + conf.Pwd + "@tcp(" + conf.Host + ")/server?charset=utf8mb4&parseTime=True"

	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	DB, err = sql.Open(conf.Dbname, dsn)
	if err != nil {
		fmt.Printf("\n\n!!!!sql open err", err)
		return err
	}
	// See "Important settings" section.
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = DB.Ping()
	if err != nil {
		fmt.Printf("\n\n!!!!!!!!!!db open err!!!!")
		return err
	}
	fmt.Println("\n\n****db connected successfully!!")
	return nil
}
