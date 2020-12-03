package main

import (
	"fmt"
	pm "pdpoll/pd_manager"
)

var mydb pm.MyDB

func init() {

	mydb.New(20)
}

func main() {

	defer mydb.Close()

	sql := "select row_to_json(t.*) from (select * from imagetab where user_id=17) t"
	results := mydb.ReadMany(sql)
	fmt.Println(results)
}
