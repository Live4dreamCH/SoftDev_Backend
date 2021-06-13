package main

import (
	"net/http"

	"github.com/Live4dreamCH/SoftDev_Backend/obj"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var u obj.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, gin.H{"res": "NO", "reason": "wrong json format!"})
		return
	}

	suss, _ := u.SignUp(u.Name, u.Psw)
	if !suss {
		c.JSON(http.StatusOK, gin.H{"res": "NO"})
	} else {
		c.JSON(http.StatusOK, gin.H{"res": "OK"})
	}
}
