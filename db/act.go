package db

import (
	"database/sql"
	"log"
)

// 预编译语句
var (
	a_search_u       *sql.Stmt
	a_search         *sql.Stmt
	a_search_aid_uid *sql.Stmt
	a_search_vote    *sql.Stmt
	a_search_period  *sql.Stmt
	a_create         *sql.Stmt
	a_create_period  *sql.Stmt
	a_update         *sql.Stmt
)

func init() {
	var err error

	a_search, err = dbp.Prepare(
		`select count(*)
		from act
		where act_id = ?`)
	check(err)
	a_search_aid_uid, err = dbp.Prepare(
		`select count(*)
		from act
		where act_id = ? and org_id=?`)
	check(err)
	a_search_vote, err = dbp.Prepare(
		`select org_period, count(u_id)
		from act_period,vote
		where act_id=? and act_period.period_id=vote.period_id
		group by org_period`)
	check(err)

	a_search_period, err = dbp.Prepare(
		`select count(*)
		from act,act_period
		where act.act_id=? and org_id=? and act_period.org_period=? and act.act_id=act_period.act_id`)
	check(err)

	a_create, err = dbp.Prepare(
		`insert into act(act_id, org_id,act_name,act_len,act_des)
		values (?, ?, ?, ?, ?)`)
	check(err)
	a_create_period, err = dbp.Prepare(
		`insert into act_period(act_id, org_period)
		values (?, ?)`)
	check(err)
	a_update, err = dbp.Prepare(
		`update act
		set act_stop=? ,act_final=?
		where act_id = ? `)
	check(err)
}

type DB_act struct{}

func (a *DB_act) Search(act_id int) (count int, err error) {
	err = a_search.QueryRow(act_id).Scan(&count)
	return
}

func (a *DB_act) Search_aid_uid(act_id int, org_id int) (count int, err error) {
	err = a_search_aid_uid.QueryRow(act_id, org_id).Scan(&count)
	if err != nil {
		return
	}
	return
}

func (a *DB_act) Search_vote(act_id int) (period []string, vote []int, err error) {
	var count int
	var period_temp string
	rows, err := a_search_vote.Query(act_id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&period_temp, &count)
		if err != nil {
			return
		}
		period = append(period, period_temp)
		vote = append(vote, count)
	}
	return
}

//查找aid,uid,period对应的记录在不在
func (a *DB_act) Search_period(act_id int, org_id int, org_period string) (count int, err error) {
	err = a_search_period.QueryRow(act_id, org_id, org_period).Scan(&count)
	if err != nil {
		return
	}
	return
}

func (a *DB_act) Create(act_id int, org_id int, act_name string, act_len int, act_des string) (suss bool, err error) {
	_, err = a_create.Exec(act_id, org_id, act_name, act_len, act_des)
	if err != nil {
		log.Println(err)
		return
	}
	suss = true
	return
}

func (a *DB_act) Create_period(act_id int, org_period []string) (suss bool, err error) {
	i := 0
	for i < len(org_period) {
		_, err = a_create_period.Exec(act_id, org_period[i])
		i = i + 1
		if err != nil {
			log.Println(err)
			return
		}
	}
	suss = true
	return
}

func (a *DB_act) Update_final(act_id int, act_final string) (suss bool, err error) {
	act_stop := true
	_, err = a_update.Exec(act_stop, act_final, act_id)
	if err != nil {
		log.Println(err)
		return
	}
	suss = true
	return
}
