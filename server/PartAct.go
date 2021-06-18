package main

import (
	"net/http"

	"github.com/Live4dreamCH/SoftDev_Backend/obj"
	"github.com/gin-gonic/gin"
)

type PartGetAct_json struct { //Participant
	Sid         string   `json:"SessionID" binding:"required"`
	ActID       int      `json:"ActID" binding:"required"`
	PartPeriods []string `json:"PartPeriods" binding:"required"`
}

func PartAct(c *gin.Context) {
	var u obj.User
	var uid,actid int
	var period []string
	var ptemp PartGetAct_json
	var err error
	if err = c.ShouldBindJSON(&ptemp); err != nil {
		c.JSON(400, gin.H{"Res": "NO", "Reason": "Sid/ActID/PartPeriods format wrong!"})
		return
	}
	uid, err = sid_manager.get(ptemp.Sid)
	if err != nil {
		return
	}
	actid=ptemp.ActID
	period=ptemp.PartPeriods
	suss, msg := u.PartAct(uid, actid, period)
	if !suss {
		c.JSON(http.StatusOK, gin.H{"Res": "NO", "Reason": msg})
	} else {
		c.JSON(http.StatusOK, gin.H{"Res": "OK"})
	}
	return
}
