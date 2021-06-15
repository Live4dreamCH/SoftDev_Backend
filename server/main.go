//服务器主程序
package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// 全局变量
var sid_manager SessIDManager

// 初始化全局变量
func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	sid_manager.m = make(map[string]int)
	//print()
}

func main() {
	// rand.Seed(1)
	r := gin.Default()
	r.POST("/SignUp", SignUp)
	r.POST("/LogIn", LogIn)
	err := r.Run(":8140")
	if err != nil {
		panic(err)
	}
}
