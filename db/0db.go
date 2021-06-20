// 对后端事物的数据库操作建模；使用嵌入式sql, 与数据库交换数据
package db

import (
	"database/sql"

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

func init() {
	var err error
	dbp, err = sql.Open("mysql", "root:jsj86_mhq_lch@tcp(114.116.234.101:3306)/app")
	check(err)
	err = dbp.Ping()
	check(err)
}
