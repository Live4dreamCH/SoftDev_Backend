package db

import (
	"database/sql"
)

// 预编译语句
var (
	a_search_u *sql.Stmt
	a_search   *sql.Stmt
	a_create   *sql.Stmt
	a_update   *sql.Stmt
)

func init() {
	var err error

	a_search, err = dbp.Prepare(
		`select count(*)
		from act
		where act_id = ?`)
	check(err)
	a_create, err = dbp.Prepare(
		`insert into act(act_id, org_id,act_name,act_len,act_des,act_stop)
		values (?, ?, ?, ?, ?, ?)`)
	check(err)
	a_update, err = dbp.Prepare(
		`update act
		set act_stop=? ,act_final=?
		where act_id = ? `)
	check(err)
}
