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
func (a *Act) Search_vote(aid int, uid int) (suss bool, period []string, vote []int) {
	var dba db.DB_act
	count, err := dba.Search_aid_uid(aid, uid)
	if err != nil {
		return
	}
	if count == 0 {
		log.Println("Aid Uid mismatching ")
	}
	period, vote, err = dba.Search_vote(aid)
	if err != nil {
		return
	}
	suss = true
	return
}

//活动发起者敲定最终时间
func (a *Act) Stop_act(aid int, uid int, Final string) (suss bool, aid_return int) {
	var dba db.DB_act
	aid_return = aid
	count, err := dba.Search_period(aid, uid, Final)
	if err != nil {
		return
	}
	if count == 0 {
		log.Println("Aid Uid or Period mismatching ")
		return
	}
	suss, err = dba.Update_final(aid, Final)
	if err != nil {
		return
	}
	suss = true
	return
}
