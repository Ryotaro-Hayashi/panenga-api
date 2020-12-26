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
	// ポート
	port := os.Getenv("MYSQL_PORT")

	var err error

	if os.Getenv("PRODUCTION") == "true" {
    
		Conn, err = sql.Open(driverName, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database))


		if err != nil {
			fmt.Printf("open mysql, %s\n", err)
		}

		fmt.Println("* ===== open mysql success ===== *")

		_, err = Conn.Exec("CREATE DATABASE IF NOT EXISTS heroku_292c3038acfa34f")
		if err != nil {
			fmt.Printf("create database, %s\n", err)
		}
		fmt.Println("* ===== create database success ===== *")

		_, err = Conn.Exec("USE heroku_292c3038acfa34f")
		if err != nil {
			fmt.Printf("use database, %s\n", err)
		}
		fmt.Println("* ===== use database success ===== *")

		_, err = Conn.Exec("CREATE TABLE panels(id INTEGER PRIMARY KEY AUTO_INCREMENT, title VARCHAR(64) NOT NULL, panel_image VARCHAR(255) NOT NULL UNIQUE, created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)")
		if err != nil {
			fmt.Printf("create table, %s\n", err)
		}
		fmt.Println("* ===== create table success ===== *")

		_, err = Conn.Exec("INSERT INTO panels (title, panel_image) VALUES ('三井アウトレットパーク札幌北広島', 'https://file24-d.kuku.lu/files/20201226-0812_c12fae45310a831a7a39834c9ecd9dcb.jpg')")
		if err != nil {
			fmt.Printf("insert test data 1, %s\n", err)
		}
		fmt.Println("* ===== insert test data 1 success ===== *")

		_, err = Conn.Exec("INSERT INTO panels (title, panel_image) VALUES ('青森ねぶた祭り', 'https://file24-d.kuku.lu/files/20201226-0814_1e834f3d94c0baf3e5e914dd7135181e.jpg')")
		if err != nil {
			fmt.Printf("insert test data 2, %s\n", err)
		}
		fmt.Println("* ===== insert test data 2 success ===== *")

		_, err = Conn.Exec("INSERT INTO panels (title, panel_image) VALUES ('函館湯川温泉', 'https://file14-d.kuku.lu/files/20201226-1653_dccf804db713b5a6ee66256fdcffb68b.jpg')")
		if err != nil {
			fmt.Printf("insert test data 3, %s\n", err)
		}
		fmt.Println("* ===== insert test data 3 success ===== *")

		_, err = Conn.Exec("INSERT INTO panels (title, panel_image) VALUES ('函館五稜郭タワー', 'https://file14-d.kuku.lu/files/20201226-1651_537299245553a7d06578ebe54fa18b59.png')")
		if err != nil {
			fmt.Printf("insert test data 4, %s\n", err)
		}
		fmt.Println("* ===== insert test data 4 success ===== *")

		_, err = Conn.Exec("INSERT INTO panels (title, panel_image) VALUES ('北海道新幹線・新函館北斗駅', 'https://file14-d.kuku.lu/files/20201226-1651_988537a094b4e60d128b7ecc64ec7d70.jpg')")
		if err != nil {
			fmt.Printf("insert test data 5, %s\n", err)
		}
		fmt.Println("* ===== insert test data 5 success ===== *")

		_, err = Conn.Exec("INSERT INTO panels (title, panel_image) VALUES ('福岡のどこか', 'https://file14-d.kuku.lu/files/20201226-1650_0f3624c43d357b05a98dc7529db909d7.jpeg')")
		if err != nil {
			fmt.Printf("insert test data 6, %s\n", err)
		}
		fmt.Println("* ===== insert test data 6 success ===== *")

	} else {
		Conn, err = sql.Open(driverName,
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
	}
	if err != nil {
		fmt.Printf("open mysql, %s\n", err)
	}
	fmt.Println("* ===== open mysql success ===== *")
}
