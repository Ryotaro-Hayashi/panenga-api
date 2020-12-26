package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

// Driver名
const driverName = "mysql"

// DB接続(Connection)情報
var Conn *sql.DB

func init() {
	/* ===== データベースへ接続する. ===== */
	// 接続先データベース
	database := os.Getenv("MYSQL_DATABASE")
	// ユーザ
	user := os.Getenv("MYSQL_USER")
	// パスワード
	password := os.Getenv("MYSQL_PASSWORD")
	// 接続先ホスト
	host := os.Getenv("MYSQL_HOST")
	// 接続先ポート
	port := os.Getenv("MYSQL_PORT")

	var err error

	if os.Getenv("PRODUCTION") == "true" {
		dataSource := os.Getenv("DATABASE_URL")
		Conn, err = sql.Open(driverName, dataSource)
	} else {
		Conn, err = sql.Open(driverName,
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
	}
	if err != nil {
		fmt.Errorf("open mysql, %s", err)
	}
}
