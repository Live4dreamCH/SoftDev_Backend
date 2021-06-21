package obj

import (
	"errors"
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

//create act,向act表里插入
func (a *Act) Create(uid int, name string, len int, des string, period []string) (suss bool, aid int) {
	var dba db.DB_act
	//获取一个不存在的aid
	for {
		aid = get_aid()
		count, err := dba.Search(aid)
		if err != nil {
			log.Println("search aid err", err)
			return
		}
		if count == 0 {
			break
		}
	}
	_, err := dba.Create(aid, uid, name, len, des)
	if err != nil {
		return
	}
	_, err = dba.Create_period(aid, period)
	if err != nil {
		return
	}
	suss = true
	log.Println("aid", aid, "create act suss")
	return
}

// 每次生成[0,10^9)区间内的int整数，类似腾讯会议号
func get_aid() int {
	return rand.Intn(1e9)
}

//查询投票数
func (a *Act) Search_vote(aid int, uid int) (period []string, vote []int, err error) {
	var dba db.DB_act
	period = make([]string, 0)
	vote = make([]int, 0)

	count, err := dba.Search_aid(aid)
	if err != nil {
		log.Println(err)
		return
	}
	if count == 0 {
		log.Println("Aid lost", aid)
		err = errors.New("ActID")
		return
	}

	count, err = dba.Search_aid_uid(aid, uid)
	if err != nil {
		log.Println(err)
		return
	}
	if count == 0 {
		log.Println("Aid Uid mismatching ", aid, uid)
		err = errors.New("Auth")
		return
	}

	period, vote, err = dba.Search_vote(aid)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

//活动发起者敲定最终时间
func (a *Act) Stop_act(aid int, uid int, Final string) (err error) {
	var dba db.DB_act

	count, err := dba.Search_aid(aid)
	if err != nil {
		log.Println(err)
		return
	}
	if count == 0 {
		log.Println("Aid lost", aid)
		err = errors.New("ActID")
		return
	}

	count, err = dba.Search_aid_uid(aid, uid)
	if err != nil {
		log.Println(err)
		return
	}
	if count == 0 {
		log.Println("Aid Uid mismatching ", aid, uid)
		err = errors.New("Auth")
		return
	}

	count, err = dba.Search_period(aid, uid, Final)
	if err != nil {
		log.Println(err)
		return
	}
	if count == 0 {
		log.Println("Period mismatching ", Final)
		err = errors.New("Periods")
		return
	}

	_, err = dba.Update_final(aid, Final)
	if err != nil {
		log.Println("Update_final ", err)
		return
	}
	log.Println("stop act,the ActID is:", aid)
	return
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
