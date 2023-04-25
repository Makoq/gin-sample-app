package controllers

import (
	"demo/database"
	model "demo/model/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) Register(c *gin.Context) {
    // 处理注册逻辑
    name:=c.PostForm("name")
    pwd:=c.PostForm("pwd")

    account:=model.Account{Name: name,Pwd: pwd}
    ret:=database.DB.Table("account").Create(&account)
    // 新插入数据的id
	if ret.Error != nil {
		fmt.Printf("register failed, err:%v\n", ret.Error)
		c.JSON(200, gin.H{
			"status": 500, "message": "get lastinsert ID failed", "log": ret.Error,
		})
		return
	}
	fmt.Println("register account", name)
	c.JSON(200, gin.H{
		"status": 200, "message": "register successfully", "PrimaryKey": ret.RowsAffected,
	})

}

func (uc *UserController) Login(c *gin.Context) {
    // 处理登录逻辑

    name:=c.Query("name")
    pwd:=c.Query("pwd")

    ret:=database.DB.Table("account").Where("name= ? AND pwd= ?",name,pwd)

    if ret.Error!=nil{
        fmt.Printf("delete failed, err:%v\n", ret.Error)
		c.JSON(200, gin.H{
			"status": -1, "message": "err", "log": ret.Error,
		})
        return
    }
    
    c.JSON(200, gin.H{
        "status": 0, "message": "success",
    })
}