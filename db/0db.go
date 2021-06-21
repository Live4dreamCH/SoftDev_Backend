// 对后端事物的数据库操作建模；使用嵌入式sql, 与数据库交换数据
package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// 严厉检查，让问题在启动时得以发现
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// 数据库
var dbp *sql.DB

// 预编译语句
var (
	a_search         *sql.Stmt
	a_search_aid     *sql.Stmt
	a_search_aid_uid *sql.Stmt
	a_search_vote    *sql.Stmt
	a_search_period  *sql.Stmt
	a_create         *sql.Stmt
	a_create_period  *sql.Stmt
	a_update         *sql.Stmt
	a_searchall      *sql.Stmt
	a_searchact      *sql.Stmt
	a_findact        *sql.Stmt
	a_Selectperiod   *sql.Stmt
	a_GetPeriodID    *sql.Stmt
)
var (
	u_search     *sql.Stmt
	u_create     *sql.Stmt
	u_login      *sql.Stmt
	u_Createvote *sql.Stmt
	u_GetName    *sql.Stmt
)

func init() {
	b, err := os.ReadFile("../pwd/local_mysql.txt")
	check(err)
	psw := string(b)
	if psw[len(psw)-1] == '\n' {
		psw = psw[:len(psw)-2]
	}

	dbp, err = sql.Open("mysql", "root:"+psw+"@/app?charset=utf8")
	check(err)
	err = dbp.Ping()
	check(err)

	a_search, err = dbp.Prepare(
		`select count(*)
		from act
		where act_id = ?`)
	check(err)
	a_searchall, err = dbp.Prepare(
		`select act_stop,act_len,org_id,act_des,act_name, act_final
		from act
		where act.act_id = ?`)
	check(err)
	a_search_aid, err = dbp.Prepare(
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
