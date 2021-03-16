package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func mysql() {
	fmt.Printf("query:\n")
	db, e := sql.Open("mysql", "root:@/mysql?charset=utf8")
	//说明：第一个参数"mysql"表示打开一个MYSQL链接
	//第二个参数是一个DSN，格式是：
	//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&paramN=valueN]

	if e != nil {
		fmt.Printf("conn error.")
		return
	}
	rows, e := db.Query("select host,user from user")
	if e != nil {
		fmt.Printf("query error.")
		return
	}

	i := 0
	for rows.Next() {
		i++
		var ss, s2 string
		e = rows.Scan(&ss, &s2)
		if e == nil {
			fmt.Printf("Title: %s,%s \n", ss, s2)
		} else {
			fmt.Printf("error")
		}
	}
	fmt.Printf("Total: %d\n", i)
	db.Close()

}
