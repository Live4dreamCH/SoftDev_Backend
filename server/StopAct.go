package main

import (
	"net/http"

	"github.com/Live4dreamCH/SoftDev_Backend/obj"
	"github.com/gin-gonic/gin"
)

type StopAct_json struct {
	Sid          string `json:"SessionID" binding:"required"`
	ActID        int    `json:"Actid" binding:"required"`
	FinalPeriods string `json:"FinalPeriods" binding:"required"`
}

func (j *StopAct_json) json2Act() (a obj.Act, err error) {
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
	if err := c.ShouldBindJSON(&j); err != nil {
		c.JSON(400, gin.H{"Res": "NO", "Reason": "wrong json format!"})
		return
	}

	a, err := j.json2Act()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Res": "NO", "Reason": "SessionID"})
		return
	}

	err = a.Stop_act(a.Aid, a.Uid, a.Final)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"Res": "OK"})
	} else {
		res := gin.H{"Res": "NO", "Reason": err.Error()}
		c.JSON(http.StatusOK, res)
	}
}
