package main

import (
	"net/http"

	"github.com/Live4dreamCH/SoftDev_Backend/obj"
	"github.com/gin-gonic/gin"
)

type Getact_json struct { //Participant
	ActID int `json:"Actid" binding:"required"`
}

func GetAct(c *gin.Context) {
	var a obj.Act
	var gtemp Getact_json
	if err := c.ShouldBindJSON(&gtemp); err != nil {
		c.JSON(400, gin.H{"Res": "NO", "Reason": "wrong json format!"})
		return
	}
	suss, act := a.GetAct(gtemp.ActID)
	if !suss {
		c.JSON(http.StatusOK, gin.H{"Res": "NO", "Reason": "ActID"})
	} else {
		var resact ResAct
		resact.Act2ResAct(act)
		c.JSON(http.StatusOK, gin.H{"Res": "OK", "Act": resact})
	}
}
