package main

import (
	"net/http"

	"github.com/Live4dreamCH/SoftDev_Backend/obj"
	"github.com/gin-gonic/gin"
)

type Getsection_json struct { //Participant
	Sid string `json:"SessionID" binding:"required"`
}

func GetActs(c *gin.Context) {
	var a obj.Act
	var getsect Getsection_json
	if err := c.ShouldBindJSON(&getsect); err != nil {
		c.JSON(400, gin.H{"Res": "NO", "Reason": "wrong json format!"})
		return
	}
	uid, err := sid_manager.get(getsect.Sid)
	if err != nil {
		c.JSON(400, gin.H{"Res": "NO", "Reason": "SessionID"}) //msg = "SessionID" //"SessionID invalid"
		return
	}
	suss, actlist := a.GetActs(uid)
	if !suss {
		c.JSON(http.StatusOK, gin.H{"Res": "NO", "Reason": "SessionID"})
	} else {
		c.JSON(http.StatusOK, gin.H{"Res": "OK", "Acts": actlist})
	}
}
