package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/test?charset=utf8")
	checkErr(err)

	// 插入语句
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(err)

	// 解析参数到语句里,并插入
	res, err := stmt.Exec("ew", "研发部门", "2012-12-09")
	checkErr(err)

	// 获取插入的id
	id, err := res.LastInsertId()
	checkErr(err)

	// 打印插入的ID
	fmt.Println(id)

	// 更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid = ?")
	checkErr(err)

	res, err = stmt.Exec("test", id)
	checkErr(err)

	// 获取SQL语句执行后影响的行数
	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(strconv.Itoa(uid) + "," + username + "," + department + "," + created)
	}

	// 删除数据
	stmt, err = db.Prepare("delete from userinfo where uid = ?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// 关闭数据库连接
	db.Close()
}

// checkErr Method
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
