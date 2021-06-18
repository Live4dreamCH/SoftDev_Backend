package db

import (
	"database/sql"
	"log"
)

// 预编译语句
var (
	u_search       *sql.Stmt
	u_create       *sql.Stmt
	u_login        *sql.Stmt
	u_findact      *sql.Stmt
	u_Createvote   *sql.Stmt
	u_Selectperiod *sql.Stmt
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
	u_login, _ = dbp.Prepare(
		`select u_id
		from user_info
		where u_name = ? and u_psw = ?`)
	u_findact, _ = dbp.Prepare(
		`select act_stop,act_len
		from act
		where act_id = ? `)
	u_Createvote, err = dbp.Prepare(
		`insert into vote(period_id, u_id)
		values (?, ?)`)
	u_Selectperiod, _ = dbp.Prepare(
		`select org_period,period_id
		from act_period
		where act_id = ?`)
	check(err)
}

type DB_user struct{}

func (u *DB_user) Search(name string) (u_id int, err error) {
	err = u_search.QueryRow(name).Scan(&u_id)
	if err != nil {
		return
	}
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

func (u *DB_user) ActivityQuery(ActID int) (act_stop bool, act_len int, err error) {
	err = u_findact.QueryRow(ActID).Scan(&act_stop, &act_len)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

//Createvote(uid,PartPeriods)
func (u *DB_user) Createvote(uid int, ActID int, Periodid []int) (err error) {
	//u_Createvote.Exec(,)
	for i := 0; i < len(Periodid); i++ {
		_, err := u_Createvote.Exec(Periodid[i], uid)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return
}

func (u *DB_user) PeriodQuery(ActID int) (num int, period []string, Periodid []int, err error) {
	//num = 0
	rows, err := u_Selectperiod.Query(ActID)
	if err != nil {
		log.Println(err)
		return num, period, Periodid, err
	}
	var org_period string
	var period_id int
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&org_period, &period_id)
		if err != nil {
			log.Println(err)
			return num, period, Periodid, err
		}
		//period[num] = org_period
		//Periodid[num] = period_id
		period = append(period, org_period)
		Periodid = append(Periodid, period_id)
		//num++
	}
	return
}
