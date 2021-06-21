package obj

import (
	"log"

	"github.com/Live4dreamCH/SoftDev_Backend/db"
)

type User struct {
	Id   int    //`json:"id"`
	Psw  string `json:"Psw" binding:"required"`
	Name string `json:"UserName" binding:"required"`
}

func (u *User) SignUp(name string, psw string) (suss bool, u_id int) {
	var dbu db.DB_user

	u_id, err := dbu.Search(name)
	if err == nil {
		return
	}

	u_id, err = dbu.Create(name, psw)
	if err != nil {
		return
	}
	suss = true
	log.Println("username", name, "sign up suss")
	return
}

func (u *User) LogIn(name string, psw string) (suss bool, valid_name bool, u_id int) {
	var dbu db.DB_user

	u_id, err := dbu.Search(name)
	if err != nil {
		return
	}
	valid_name = true

	suss, u_id, err = dbu.LoginQuery(name, psw)
	if err != nil {
		log.Println("username=", name, "psw=", psw, "login fail")
		return
	}
	suss = true
	return
}

func (u *User) PartAct(uid int, ActID int, PartPeriods []string) (suss bool, msg string) {
	var dbu db.DB_user
	var dba db.DB_act
	act_stop, _, err := dba.ActivityQuery(ActID) //act_len
	if err != nil {
		log.Println("ActID=", ActID, "ActID not found")
		msg = "ActID" //"ActID not found"
		return
	}
	if act_stop {
		log.Println("ActID=", ActID, "ActID stop!")
		msg = "Stopped" //此活动已停止投票
		return
	}

	// 1. 检查period参数的正确性,不正确则退出
	// 2. 用period换PartPeriodID
	PartPeriodID := make([]int, len(PartPeriods))
	i := 0
	for _, p := range PartPeriods {
		pid, err := dba.GetPeriodID(ActID, p)
		if err != nil {
			log.Println("period", p, "from PartAct.PartPeriods not in ActID", ActID, "err=", err)
			msg = "Periods"
			return
		}
		PartPeriodID[i] = pid
		i++
	}

	//请求列出的活动时间都在组织者活动时间列表中
	//插入参与活动的记录
	err = dbu.CreateVote(uid, ActID, PartPeriodID)
	if err != nil {
		log.Println("ActID", ActID, "cannot add voter", uid)
		msg = "Parted"
		return
	}
	suss = true
	return
}

func (u *User) GetName(uid int) (name string) {
	var dbu db.DB_user
	name, err := dbu.GetName(uid)
	if err != nil {
		log.Println(err)
		name = "unknown"
	}
	return
}
