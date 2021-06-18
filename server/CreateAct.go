package main

import (
	"log"
	"net/http"

	"github.com/Live4dreamCH/SoftDev_Backend/obj"
	"github.com/gin-gonic/gin"
)

type CreateAct_json struct {
	Sid    string   `json:"SessionID" binding:"required"`
	Name   string   `json:"ActName" binding:"required"`
	Len    int      `json:"Length" binding:"required"`
	Des    string   `json:"Description" binding:"required"`
	Period []string `json:"OrgPeriods" binding:"required"`
}

func (j *CreateAct_json) json2Act() (a obj.Act, err error) {
	uid, err := sid_manager.get(j.Sid)
	if err != nil {
		return
	}
	a.Uid = uid
	a.Name = j.Name
	a.Len = j.Len
	a.Des = j.Des
	a.Period = j.Period
	return
}

func CreateAct(c *gin.Context) {
	var j CreateAct_json
	if err := c.ShouldBindJSON(&j); err != nil {
		c.JSON(400, gin.H{"Res": "NO", "Reason": "wrong json format!"})
		return
	}

	a, err := j.json2Act()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Res": "NO", "Reason": "SessionID"})
		return
	}

	suss, aid := a.Create(a.Uid, a.Name, a.Len, a.Des, a.Period)
	if !suss {
		res := gin.H{"Res": "NO", "Reason": "Data type does not match"}
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, gin.H{"Res": "OK", "ActID": aid})
		log.Println("create suss: uid=", aid, "ActID=", aid)
	}
}
