package db

import (
	"log"
)

type DB_user struct{}

func (u *DB_user) Search(name string) (u_id int, err error) {
	err = u_search.QueryRow(name).Scan(&u_id)
	return
}

func (u *DB_user) LoginQuery(name string, psw string) (suss bool, u_id int, err error) {
	err = u_login.QueryRow(name, psw).Scan(&u_id)
	if err != nil {
		log.Println(err)
		return
	}
	suss = true
	return
}

func (u *DB_user) Create(name string, psw string) (u_id int, err error) {
	res, err := u_create.Exec(name, psw)
	if err != nil {
		log.Println(err)
		return
	}
	i, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return
	}
	u_id = int(i)
	return
}

func (u *DB_user) CreateVote(uid int, ActID int, Periodid []int) (err error) {
	for i := 0; i < len(Periodid); i++ {
		_, err = u_Createvote.Exec(Periodid[i], uid)
		if err != nil {
			log.Println(err)
			return
		}
	}
	return
}

func (u *DB_user) GetName(uid int) (name string, err error) {
	err = u_GetName.QueryRow(uid).Scan(&name)
	return
}
