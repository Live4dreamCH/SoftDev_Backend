package main

import (
	"log"
	"net/http"

	"github.com/Live4dreamCH/SoftDev_Backend/obj"
	"github.com/gin-gonic/gin"
)

type Getsection_json struct { //Participant
	Sid string `json:"SessionID" binding:"required"`
}

type ResAct struct {
	ActID   int      `json:"ActID"`
	OrgName string   `json:"OrgName"`
	ActName string   `json:"ActName"`
	Len     int      `json:"Length"`
	Des     string   `json:"Description"`
	Stop    int      `json:"Stopped"`
	Period  []string `json:"OrgPeriods"`
}

func (r *ResAct) Act2ResAct(a obj.Act) {
	var u obj.User
	r.ActID = a.Aid
	r.OrgName = u.GetName(a.Uid)
	r.ActName = a.Name
	r.Len = a.Len
	r.Des = a.Des
	if a.Stop {
		r.Stop = 1
	} else {
		r.Stop = 0
	}
	r.Period = a.Period
	for i := range r.Period {
		backfix := r.Period[i][len(r.Period[i])-3 : len(r.Period[i])]
		if backfix == ":00" {
			r.Period[i] = r.Period[i][:len(r.Period[i])-3]
		} else {
			log.Println("strip wrong! origin time str = ", r.Period[i], "; not ended with ':00'")
		}
	}
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
		c.JSON(http.StatusOK, gin.H{"Res": "NO", "Reason": "SessionID"}) //msg = "SessionID" //"SessionID invalid"
		return
	}

	suss, actlist := a.GetActs(uid)
	if !suss {
		c.JSON(http.StatusOK, gin.H{"Res": "NO", "Reason": "GetActs error"})
	} else {
		resactlist := make([]ResAct, len(actlist))
		for i := 0; i < len(actlist); i++ {
			resactlist[i].Act2ResAct(actlist[i])
		}
		c.JSON(http.StatusOK, gin.H{"Res": "OK", "Acts": resactlist})
	}
}
