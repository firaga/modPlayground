package main

import "fmt"

/**
声明人类这个接口，有吃饭“eat”的方法
*/
type computer struct {
	logo string
	//name string
	// 不能在结构体中声明方法
	// ShowLogo()
}

//如果一个struct嵌入另一个匿名结构体，就可以直接访问匿名结构体的字段或方法，从而实现继承
type dell struct {
	// 匿名结构体
	computer
	//logo string
}

//如果一个struct嵌套了另外一个有名结构体，是UML中的组合或者聚合（根据实际情况来定）
type apple struct {
	pc computer
	//name string
}

/**
声明了父类computer的showLogo方法，接收computer做对象
*/
func (ptr computer) ShowLogo() {
	fmt.Println("当前电脑的logo是：" + ptr.logo)
}
func (ptr computer) ShowName() {
	//fmt.Println("当前电脑的name 是：" + ptr.name)
}

func (ptr dell) ShowName() {
	//fmt.Println("当前电脑的name 是：" + ptr.computer.name)
	//fmt.Println("当前电脑的name 是：" + ptr.name)
}

func (ptr apple) ShowName() {
	//fmt.Println("当前电脑的name 是：" + ptr.pc.name)
	//fmt.Println("当前电脑的name 是：" + ptr.name)
}

func (ptr dell) ShowLogo() {
	fmt.Println("当前电脑的logo是：" + ptr.logo)
}

func (ptr apple) ShowLogo() {
	fmt.Println("当前电脑的logo是：" + ptr.pc.logo)
}

func main() {
	// 因为go只有伪继承, 其本质是一种语法糖, 无法实现继承中动态绑定等特性. 且构造函数, 连伪继承都达不到.
	//所以一下代码编译会报错类型不匹配 Cannot use 'dell{computer{logo: "dell"}}' (type dell) as type computer
	//var pc computer = dell{computer{logo: "dell"}}
	//showLogo(pc)
	//pc = apple{computer{logo: "apple"}}
	//showLogo(pc)

	// 伪继承可以直接调用，但是不能使用父类指针指向子类对象
	var d dell = dell{computer{logo: "computer"}}
	d.ShowLogo()
	d.computer.ShowLogo()
	// 组合更不能了。
	var a apple = apple{computer{logo: "apple"}}
	a.pc.ShowLogo()
	a.ShowLogo()
}

func showLogo(pc computer) {
	pc.ShowLogo()
}
