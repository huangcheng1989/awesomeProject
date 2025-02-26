package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type User struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Age      int        `json:"age"`
	Birthday *time.Time `json:"birthday"`
}

// MarshalJSON 重写json时间类型输出
func (u *User) MarshalJSON() ([]byte, error) {
	type Temp User
	// ToMarshal 匿名结构体
	type ToMarshal struct {
		*Temp
		Birthday string `json:"birthday"`
	}
	var to = ToMarshal{
		Temp: (*Temp)(u),
	}
	if u.Birthday != nil {
		to.Birthday = u.Birthday.Format("2006-01-02")
	}
	return json.Marshal(to)
}

func handle(c *gin.Context) {
	var user User
	c.ShouldBindJSON(&user)
	log.Printf("Birthday: %+v", user.Birthday)
	//u, _ := json.Marshal(user)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": &user,
	})
}

func main() {
	router := gin.Default()
	router.POST("/", handle)
	router.Run(":9090") // 默认会在0.0.0.0:8080监听
}
