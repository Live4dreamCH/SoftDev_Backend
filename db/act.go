package db

import (
	"database/sql"
	"log"
)

// 预编译语句
var (
	a_searchall    *sql.Stmt
	a_searchact    *sql.Stmt
	a_findact      *sql.Stmt
	a_Selectperiod *sql.Stmt
	a_GetPeriodID  *sql.Stmt
)

func init() {
	var err error
	a_searchall, err = dbp.Prepare(
		`select act_stop,act_len,org_id,act_des,act_name, act_final
		from act
		where act.act_id = ?`)
	check(err)
	a_searchact, err = dbp.Prepare(
		`select act_id
		from act_period,vote
		where u_id = ? and act_period.period_id = vote.period_id
		union
		select act_id
		from act
		where org_id = ?`)
	check(err)
	a_findact, err = dbp.Prepare(
		`select act_stop,act_len
		from act
		where act_id = ? `)
	check(err)
	a_Selectperiod, err = dbp.Prepare(
		`select org_period,period_id
		from act_period
		where act_id = ?`)
	check(err)
	a_GetPeriodID, err = dbp.Prepare(
		`select ap.period_id
		from act_period ap
		where ap.act_id = ? and ap.org_period = ?`)
	check(err)
}

type DB_act struct{}

func (a *DB_act) ActinfoQuery(ActID int) (act_stop bool, act_len int, org_id int, act_des string, act_name string, act_final string, err error) {
	var null_act_final sql.NullString
	err = a_searchall.QueryRow(ActID).Scan(&act_stop, &act_len, &org_id, &act_des, &act_name, &null_act_final)
	if act_stop {
		act_final = null_act_final.String
	}
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (a *DB_act) OnepersonActQuery(uid int) (actidlist []int, err error) {
	rows, err := a_searchact.Query(uid, uid)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var act_id int
		err = rows.Scan(&act_id)
		if err != nil {
			log.Println(err)
			return
		}
		actidlist = append(actidlist, act_id)
	}
	return
}

func (a *DB_act) ActivityQuery(ActID int) (act_stop bool, act_len int, err error) {
	err = a_findact.QueryRow(ActID).Scan(&act_stop, &act_len)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (a *DB_act) PeriodQuery(ActID int) (period map[string]int, err error) {
	period = make(map[string]int)
	rows, err := a_Selectperiod.Query(ActID)
	if err != nil {
		log.Println(err)
		return
	}
	var org_period string
	var period_id int
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&org_period, &period_id)
		if err != nil {
			log.Println(err)
			return
		}
		period[org_period] = period_id
	}
	return
}

func (a *DB_act) GetPeriodID(ActID int, period string) (pid int, err error) {
	err = a_GetPeriodID.QueryRow(ActID, period).Scan(&pid)
	return
}
