package obj

import (
	"log"
	"math/rand"

	"github.com/Live4dreamCH/SoftDev_Backend/db"
)

type Act struct {
	Aid    int      //`json:"Actid"`
	Uid    int      //"userid"
	Name   string   //`json:"ActName" binding:"required"`
	Len    int      //`json:"Length" binding:"required"`
	Des    string   //`json:"Description" binding:"required"`
	Stop   bool     //"0:not stop,1:stop"
	Final  string   //the final time
	Period []string //`json:"OrgPeriods" binding:"required"`
}

func (a *Act) Create(uid int, name string, len int, des string) (suss bool, aid int) {
	get_aid()
	return
}

// 每次生成[0,10^9)区间内的int整数，类似腾讯会议号
func get_aid() int {
	return rand.Intn(1e9)
}

func (a *Act) GetAct(act_id int) (suss bool, act Act) {
	var dba db.DB_act
	act_stop, act_len, org_id, act_des, act_name, act_final, err := dba.ActinfoQuery(act_id)
	if err != nil {
		return
	}
	act.Aid = act_id
	act.Des = act_des
	act.Len = act_len
	act.Uid = org_id
	act.Name = act_name
	act.Stop = act_stop
	if act_stop {
		act.Period = append(act.Period, act_final)
		suss = true
		return
	} else {
		period, err := dba.PeriodQuery(act_id)
		if err != nil {
			return
		}
		act.Period = make([]string, len(period))
		i := 0
		for p := range period {
			act.Period[i] = p
			i++
		}
		suss = true
		return
	}
}

func (a *Act) GetActs(uid int) (suss bool, actlist []Act) {
	var i int
	var actidlist []int
	var tempact Act
	var dbu db.DB_act
	actidlist, err := dbu.OnepersonActQuery(uid) //act_len
	if err != nil {
		log.Println("uid=", uid, "uid doesn't have activity")
		suss = false
		return
	}
	for i = 0; i < len(actidlist); i++ {
		suss, tempact = a.GetAct(actidlist[i])
		if !suss {
			return
		}
		actlist = append(actlist, tempact)
	}
	suss = true
	return
}
