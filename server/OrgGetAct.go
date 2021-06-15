package main

import (
	"log"
	"net/http"

	"github.com/Live4dreamCH/SoftDev_Backend/obj"
	"github.com/gin-gonic/gin"
)

type GetAct_json struct {
	Sid   string `json:"SessionID" binding:"required"`
	ActID int    `json:"Actid" binding:"required"`
}

func (j *GetAct_json) json2Act() (a obj.Act, err error) {
	uid, err := sid_manager.get(j.Sid)
	if err != nil {
		return
	}
	a.Aid = j.ActID
	a.Uid = uid
	return
}

func OrgGetAct(c *gin.Context) {
	var j GetAct_json
	var a obj.Act
	var period []string
	var vote []int
	var suss bool
	if err := c.ShouldBindJSON(&j); err != nil {
		c.JSON(400, gin.H{"Res": "NO", "Reason": "wrong json format!"})
		return
	}
	a, _ = j.json2Act()
	suss, period, vote = a.Search_vote(a.Aid, a.Uid)
	if !suss {
		res := gin.H{"Res": "NO", "Reason": "Aid or Uid has problem"}
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, gin.H{"Res": "OK", "AvlbPeriods": period, "Votes": vote})
		log.Println("check vote suss")
	}
}
