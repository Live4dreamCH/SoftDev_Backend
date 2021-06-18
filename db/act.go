package db

import (
	"database/sql"
	"log"
)

// 数据库
//var dbp *sql.DB
//dbp =new (sql.DB)
// 预编译语句
var (
	//a_search_u  *sql.Stmt
	//a_search    *sql.Stmt
	//a_create    *sql.Stmt
	//a_update    *sql.Stmt
	a_searchall *sql.Stmt
	a_searchact *sql.Stmt
	//a_searchperiod *sql.Stmt
)

func init() {
	var err error
	/*a_search, err = dbp.Prepare(
		`select count(*)
		from act
		where act_id = ?`)
	check(err)
	a_create, err = dbp.Prepare(
		`insert into act(act_id, org_id,act_name,act_len,act_des,act_stop)
		values (?, ?, ?, ?, ?, ?)`)
	check(err)
	a_update, _ = dbp.Prepare(
		`update act
		set act_stop=? ,act_final=?
		where act_id = ? `)*/
	//a_searchall = new(sql.Stmt)
	a_searchall, _ = dbp.Prepare(
		`select act_stop,act_len,org_id,act_des,act_name
		from act
		where act.act_id = ?`)
	a_searchact, err = dbp.Prepare(`select distinct act_id
		from act_period,vote
		where u_id = ? and act_period.period_id = vote.period_id`)
	check(err)
}

type DB_act struct{}

func (a *DB_act) ActinfoQuery(ActID int) (act_stop bool, act_len int, org_id int, act_des string, act_name string, err error) {
	err = a_searchall.QueryRow(ActID).Scan(&act_stop, &act_len, &org_id, &act_des, &act_name)
	if err != nil {
		log.Println(err)
		return act_stop, act_len, org_id, act_des, act_name, err
	}
	return
}
func (a *DB_act) OnepersonActQuery(uid int) (actidlist []int, err error) {
	//var num int = 0
	rows, newerr := a_searchact.Query(uid)
	if newerr != nil {
		log.Println(newerr)
		err = newerr
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
		//Periodid=append(Periodid,period_id)
		//actidlist[num] = act_id
		//num++
	}
	return
}

//
