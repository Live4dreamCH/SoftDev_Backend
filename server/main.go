//服务器主程序
package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

// 全局变量
var sid_manager SessIDManager

// 初始化全局变量
func init() {
	rand.Seed(time.Now().UnixNano())

	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	sid_manager.m = make(map[string]int)
}

func main() {
	r := gin.Default()
	r.POST("/SignUp", SignUp)
	r.POST("/LogIn", LogIn)
	r.POST("/CreateAct", CreateAct)
	err := r.Run(":8140")
	if err != nil {
		panic(err)
	}
}
