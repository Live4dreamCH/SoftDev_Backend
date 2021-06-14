package main

import (
	"log"
	"net/http"

	"github.com/Live4dreamCH/SoftDev_Backend/obj"
	"github.com/gin-gonic/gin"
)

func LogIn(c *gin.Context) {
	var u obj.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, gin.H{"Res": "NO", "Reason": "wrong json format!"})
		return
	}

	suss, valid_name, u_id := u.LogIn(u.Name, u.Psw)
	if !suss {
		res := gin.H{"Res": "NO", "Reason": ""}
		if !valid_name {
			res["Reason"] = "UserName"
		} else {
			res["Reason"] = "Psw"
		}
		c.JSON(http.StatusOK, res)
	} else {
		sid := sid_manager.set(u_id)
		c.JSON(http.StatusOK, gin.H{"Res": "OK", "SessionID": sid})
		log.Println("login suss: uid=", u_id, "SessionID=", sid)
	}
}
