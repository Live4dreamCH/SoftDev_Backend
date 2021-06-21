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
	//活动发起者
	r.POST("/SignUp", SignUp)
	r.POST("/LogIn", LogIn)
	r.POST("/CreateAct", CreateAct)
	r.POST("/OrgGetAct", OrgGetAct)
	r.POST("/StopAct", StoptAct)
	//活动参与者三个函数
	r.POST("/GetActs", GetActs)
	r.POST("/PartAct", PartAct)
	r.POST("/GetAct", GetAct)
	err := r.Run(":8140")
	if err != nil {
		panic(err)
	}
}
