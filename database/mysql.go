package database

import (
	// *sql.DB对象

	//配置，主要是数据库用户、密码等
	"demo/conf"
	"fmt"

	//GORM
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义一个全局对象db
var DB *gorm.DB

// 定义一个初始化数据库的函数
func InitDB() (err error) {
	// 读取配置文件
	var c conf.Conf
	conf := c.GetConf()
	// DSN:Data Source Name
	dsn := conf.User + ":" + conf.Pwd + "@tcp(" + conf.Host + ")/server?charset=utf8mb4&parseTime=True"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("\n\n!!!!sql open err", err)
		return err
	}

	db, err := DB.DB()
	err = db.Ping()
	if err != nil {
		fmt.Printf("\n\n!!!!!!!!!!db open err!!!!")
		return err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(100)
	// 尝试与数据库建立连接（校验dsn是否正确）
	// err = DB.Ping()
	if err != nil {
		fmt.Printf("\n\n!!!!!!!!!!db open err!!!!")
		return err
	}
	fmt.Println("\n\n****db connected successfully!!")
	return nil
}
