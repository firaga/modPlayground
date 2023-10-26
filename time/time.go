package main

import (
	"fmt"
	"time"
)

func main() {
	var t time.Time
	//t = time.Unix(0, 0)
	fmt.Printf("%v\n", t)
	fmt.Println(t.Format("2006-01-02 15:04:05.000 Mon Jan"))
	fmt.Println(t.IsZero())
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}
