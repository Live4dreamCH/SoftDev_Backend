package obj

import (
	"log"

	"github.com/Live4dreamCH/SoftDev_Backend/db"
	//"db" //github.com/Live4dreamCH/SoftDev_Backend/
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
	suss=true
	return
}

//调用时：suss, msg := u.PartAct(uid, ptemp.ActID, ptemp.PartPeriods)
func (u *User) PartAct(uid int, ActID int, PartPeriods []string) (suss bool, msg string) {
	var dbu db.DB_user
	act_stop, _, err := dbu.ActivityQuery(ActID) //act_len
	if err != nil {
		log.Println("ActID=", ActID, "ActID not found")
		msg = "ActID" //"ActID not found"
		suss = false
		return
	}
	if act_stop {
		log.Println("ActID=", ActID, "ActID stop!")
		msg = "Stopped" //此活动已停止投票
		suss = false
		return
	}
	//var dbu2 db.DB_user
	//PeriodQuery(ActID int) (num int, period []string, err error)
	_, period, Periodid, err := dbu.PeriodQuery(ActID) //num
	if err != nil {
		log.Println("ActID=", ActID, "ActID does not have periods")
		msg = "ActID" //"ActID not found"
		suss = false
		return suss, msg
	}
	var i int
	for i = 0; i < len(PartPeriods); i++ { //判别请求列出的活动时间是否合理
		for j := 0; j < len(period); j++ {
			if PartPeriods[i] == period[j] {
				break
			}
			if j == len(period) {
				break
			}
		}
	}

	if i < len(PartPeriods) {
		suss = false
		msg = "Periods"
		return suss, msg
	}
	//请求列出的活动时间都在组织者活动时间列表中
	//插入参与活动的记录
	err = dbu.Createvote(uid, ActID, Periodid)
	if err != nil {
		log.Println("ActID=", ActID, "ActID cannot add vote")
		msg = "Periods" //"ActID not found"
		suss = false
		return
	}
	suss = true
	return
}
