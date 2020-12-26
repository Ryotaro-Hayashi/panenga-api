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
		Conn, err = sql.Open(driverName, fmt.Sprintf("%s:%s@tcp(%s::3306)/%s?parseTime=true", user, password, host, database))

		if err != nil {
			fmt.Printf("open mysql, %s", err)
			return
		}

		fmt.Printf("open mysql success")

		_, err = Conn.Exec("CREATE DATABASE IF NOT EXISTS panenga_db")
		if err != nil {
			fmt.Printf("create database, %s", err)
			return
		}

		_, err = Conn.Exec("USE panenga_db")
		if err != nil {
			fmt.Printf("use database, %s", err)
			return
		}

		_, err = Conn.Exec("CREATE TABLE panels(id INTEGER PRIMARY KEY AUTO_INCREMENT, title VARCHAR(64) NOT NULL, panel_image VARCHAR(255) NOT NULL UNIQUE, created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP)")
		if err != nil {
			fmt.Printf("create table, %s", err)
			return
		}

		_, err = Conn.Exec("INSERT INTO panels (title, panel_image) VALUES (三井アウトレットパーク札幌北広島, https://file24-d.kuku.lu/files/20201226-0812_c12fae45310a831a7a39834c9ecd9dcb.jpg)")
		if err != nil {
			fmt.Printf("insert test data 1, %s", err)
			return
		}

		_, err = Conn.Exec("INSERT INTO panels (title, panel_image) VALUES (青森ねぶた祭り, https://file24-d.kuku.lu/files/20201226-0814_1e834f3d94c0baf3e5e914dd7135181e.jpg)")
		if err != nil {
			fmt.Printf("insert test data 2, %s", err)
			return
		}

		Conn.Close()
		
	} else {
		Conn, err = sql.Open(driverName,
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
	}
	if err != nil {
		fmt.Printf("open mysql, %s", err)
	}
}
