package handler

import (
	database "demo/database"
	model "demo/model/user"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

//查数据
func Test(c *gin.Context) {
	//model数据

	c.JSON(200, gin.H{
		"MSG": "test",
	})

	fmt.Println("getData test")
}

// 查询所有
func Get(c *gin.Context) {
	var user []model.User
	ret := database.DB.Table("user").Find(&user)
	if ret.Error != nil {
		panic(ret.Error)
	}
	c.JSON(200, gin.H{
		"status": 200, "message": user,
	})
}

func Add(c *gin.Context) {
	//model数据
	name := c.PostForm("name")
	age := c.PostForm("age")
	intAge, err := strconv.Atoi(age)

	if err != nil {
		fmt.Printf("convert age from string to int failed, err:%v\n", err)
		c.JSON(200, gin.H{
			"status": 500, "message": "get lastinsert ID failed", "log": err,
		})
		return
	}

	user := model.User{Name: name, Age: intAge}
	//创建/插入
	ret := database.DB.Table("user").Create(&user)

	// 新插入数据的id
	if ret.Error != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", ret.Error)
		c.JSON(200, gin.H{
			"status": 500, "message": "get lastinsert ID failed", "log": ret.Error,
		})
		return
	}
	fmt.Println("insert: ", name, age)
	c.JSON(200, gin.H{
		"status": 200, "message": "insert successfully", "PrimaryKey": ret.RowsAffected,
	})

}

func UpdateData(c *gin.Context) {

	name := c.Query("name")

	ret := database.DB.Table("user").Where("name = ?", name).Update("age", 0)

	// 新插入数据的id
	if ret.Error != nil {
		fmt.Printf("update ID failed, err:%v\n", ret.Error)
		c.JSON(200, gin.H{
			"status": 500, "message": "update lastinsert ID failed", "log": ret.Error,
		})
		return
	}
	// 受影响数为0时
	if ret.RowsAffected == 0 {
		c.JSON(200, gin.H{
			"status": 100, "message": "update db err", "log": "0 rows affected",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": 200, "message": "update successfully", "PrimaryKey": ret.RowsAffected,
	})

}

func DelData(c *gin.Context) {
	var user []model.User

	name := c.Query("name")

	ret := database.DB.Table("user").Where("name = ?", name).Delete(&user)
	if ret.Error != nil {
		fmt.Printf("delete failed, err:%v\n", ret.Error)
		c.JSON(200, gin.H{
			"status": 500, "message": "delete failed", "log": ret.Error,
		})
		return
	}

	// 受影响数为0时
	if ret.RowsAffected == 0 {
		c.JSON(200, gin.H{
			"status": 100, "message": "del db err", "log": "0 rows affected",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": 200, "message": "del successfully", "PrimaryKey": ret.RowsAffected,
	})

}

func PatchData(c *gin.Context) {

}

func OptionData(c *gin.Context) {

}
