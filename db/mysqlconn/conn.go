package mysqlconn

import (
	"database/sql"
	"fmt"
	"github.com/skydrive/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)


var mysqldb *sql.DB

func init() {
	var error error
	//mysqldb, error := sql.Open("mysql", config.GetDataSourceName())
	//上面这种会给全局的变量覆盖，报runtime error: invalid memory address or nil pointer dereference
	mysqldb, error = sql.Open("mysql", config.GetDataSourceName())
	if error != nil {
		fmt.Println(" connect mysql Error", error)
		return
	}
	//最大链接数
	mysqldb.SetMaxOpenConns(50)
	//空闲连接数
	mysqldb.SetMaxIdleConns(10)
	//最大连接周期
	mysqldb.SetConnMaxLifetime(1 * time.Second)
	if pingerror := mysqldb.Ping(); pingerror != nil {
		fmt.Println(" connect mysql failed", pingerror)
		return
	}
	fmt.Println("connect mysql success")
}

func DbConnect() *sql.DB {
	return mysqldb
}
