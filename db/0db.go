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
}
