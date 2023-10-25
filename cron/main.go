package main

import (
	"fmt"
	"github.com/robfig/cron"
)

// 主函数
func main() {
	cron2 := cron.New() //创建一个cron实例

	//执行定时任务（每5秒执行一次）
	err := cron2.AddFunc("*/5 * * * * *", print5)
	if err != nil {
		fmt.Println(err)
	}

	//启动/关闭
	cron2.Start()
	defer cron2.Stop()
	select {
	//查询语句，保持程序运行，在这里等同于for{}
	}
}

// 执行函数
func print5() {
	fmt.Println("每5s执行一次cron")
}
