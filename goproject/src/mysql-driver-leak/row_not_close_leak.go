package main

import (
	//"encoding/json"

	"fmt"
	"grutils/grapps/config"
	"time"
	//"grutils/grmath"
)

func main() {
	// go test -v -test.run Test_Order

	config.Init("config.xml")

	adminConn, err := config.DataAdminGet(false)
	if err != nil {
		fmt.Println(err)
		return
	}
	//date_format(`create_time`, '%Y-%m-%d %H:%i:%s')

	for {

		time.Sleep(time.Second)

		sql := "SELECT * from wechatinfo"

		rows, err := adminConn.Query(sql)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(rows.ColumnTypes)

		//defer rows.Close()

		//for rows.Next() {

		//}
	}
	return

}
