package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/go-sql-driver/mysql"
)

// Driver名
const driverName = "mysql"

// Conn 各repositoryで利用するDB接続(Connection)情報
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

	Conn, err := sql.Open(driverName,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
	if err != nil {
		log.Print(err)
	}
}