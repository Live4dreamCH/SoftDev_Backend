package main

import (
	"log"
	"net/http"

	"github.com/Live4dreamCH/SoftDev_Backend/obj"
	"github.com/gin-gonic/gin"
)

type StopAct_json struct {
	Sid          string `json:"SessionID" binding:"required"`
	ActID        int    `json:"Actid" binding:"required"`
	FinalPeriods string `json:"FinalPeriods" binding:"required"`
}

func (j *StopAct_json) json3Act() (a obj.Act, err error) {
	uid, err := sid_manager.get(j.Sid)
	if err != nil {
		return
	}
	a.Aid = j.ActID
	a.Uid = uid
	a.Final = j.FinalPeriods
	return
}

func StoptAct(c *gin.Context) {
	var j StopAct_json
	var a obj.Act
	if err := c.ShouldBindJSON(&j); err != nil {
		c.JSON(400, gin.H{"Res": "NO", "Reason": "wrong json format!"})
		return
	}
	a, _ = j.json3Act()
	suss, aid := a.Stop_act(a.Aid, a.Uid, a.Final)
	if !suss {
		res := gin.H{"Res": "NO", "Reason": "SessionID/ActID/Periods/Auth has problem"}
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, gin.H{"Res": "OK"})
		log.Println("stop act,the Actid is:", aid)
	}
}
