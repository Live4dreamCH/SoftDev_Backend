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
	}
	return
}
