package db

import (
	"database/sql"
	"log"
)

// 预编译语句
var (
	u_search     *sql.Stmt
	u_create     *sql.Stmt
	u_login      *sql.Stmt
	u_Createvote *sql.Stmt
	u_GetName    *sql.Stmt
)

func init() {
	var err error
	u_search, err = dbp.Prepare(
		`select u_id
		from user_info
		where u_name = ?`)
	check(err)
	u_create, err = dbp.Prepare(
		`insert into user_info(u_name, u_psw)
		values (?, ?)`)
	check(err)
	u_login, err = dbp.Prepare(
		`select u_id
		from user_info
		where u_name = ? and u_psw = ?`)
	check(err)
	u_Createvote, err = dbp.Prepare(
		`insert into vote(period_id, u_id)
		values (?, ?)`)
	check(err)
	u_GetName, err = dbp.Prepare(
		`select u_name
		from user_info
		where u_id = ?`)
	check(err)
}

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
