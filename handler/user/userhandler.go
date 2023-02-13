package handler

import (
	database "demo/database"
	model "demo/model/user"
	"fmt"
	
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
	rows, err := database.DB.Query("SELECT * FROM user")
	if err != nil {
		panic(err)
	}
	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {

		}

		users = append(users, user)
	}
	c.JSON(200, gin.H{
		"status": 200, "message": users,
	})
}

func Add(c *gin.Context) {
	//model数据
	name := c.PostForm("name")
	age := c.PostForm("age")

	sql := "insert into user(name,age) values(?,?)"
	ret, err := database.DB.Exec(sql, name, age)
	if err != nil {
		fmt.Printf("add db err", err)

		c.JSON(200, gin.H{
			"status": 500, "message": "add db err", "log": err,
		})
		return

	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		c.JSON(200, gin.H{
			"status": 500, "message": "get lastinsert ID failed", "log": err,
		})
		return
	}
	fmt.Println("insert: ", name, age)
	c.JSON(200, gin.H{
		"status": 200, "message": "insert successfully", "PrimaryKey": theID,
	})

}

func UpdateData(c *gin.Context) {

}

func DelData(c *gin.Context) {
	key := c.Query("key")
	sql := "DELETE FROM user WHERE id = ?"
	ret, err := database.DB.Exec(sql, key)
	fmt.Print("key", key)
	if err != nil {
		fmt.Printf("del db err", err)

		c.JSON(200, gin.H{
			"status": 500, "message": "del db err", "log": err,
		})
		return

	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Print("del err")
		c.JSON(200, gin.H{
			"status": 500, "message": "del db err", "log": err,
		})
		return
	}
	// 受影响数为0时
	if n == 0 {
		c.JSON(200, gin.H{
			"status": 500, "message": "del db err", "log": "0 rows affected",
		})
		return
	}

	fmt.Println("del rows %d: ", n)
	c.JSON(200, gin.H{
		"status": 200, "message": "delete successfully", "log": n,
	})

}

func PatchData(c *gin.Context) {

}

func OptionData(c *gin.Context) {

}
