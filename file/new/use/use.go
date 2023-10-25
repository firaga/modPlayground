package main

import "file/new/used"

func main() {
	a := &used.A{}
	a.One = 1

	//不能读取返回类型为接口的struct元素, 需要断言才能读取One
	//b := used.NewA()
	//b.One = 1
}
