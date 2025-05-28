package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"time"
)

// checkErr Method
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// initDatabase Method 从SQL文件读取并执行初始化
func initDatabase(db *sql.DB) error {
	// 读取sql文件
	content, err := os.ReadFile("./api_sqlite/init.sql")
	checkErr(err)

	// 执行sql语句
	_, err = db.Exec(string(content))
	return err
}

func main() {
	currentDir, err := os.Getwd()
	checkErr(err)
	fmt.Println("Current Directory: ", currentDir)

	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	// 初始化sqlite
	err = initDatabase(db)

	// 插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username,department,created) values (?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("eastwind", "996", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	// 更新数据
	stmt, err = db.Prepare("update userinfo set username = ? where uid = ?")
	checkErr(err)

	res, err = stmt.Exec("eastwind666", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}

	// 删除数据
	//stmt, err = db.Prepare("delete from userinfo where uid = ?")
	//checkErr(err)
	//
	//res, err = stmt.Exec(id)
	//checkErr(err)
	//
	//affect, err = res.RowsAffected()
	//fmt.Println(affect)

	// 关闭数据库
	err = db.Close()
	checkErr(err)
}
