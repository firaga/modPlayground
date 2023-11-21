package main

import (
	"fmt"
	"github.com/wuqinqiang/easyfsm"
)

var (
	// 业务
	businessName easyfsm.BusinessName = "order"

	// 对应状态
	initState easyfsm.State = 1 // 初始化
	paidState easyfsm.State = 2 // 已付款
	canceled  easyfsm.State = 3 // 已取消

	//对应事件
	paymentOrderEventName easyfsm.EventName = "paymentOrderEventName"
	cancelOrderEventName  easyfsm.EventName = "cancelOrderEventName"
)

type (
	orderParam struct {
		OrderNo string
	}
)

func init() {
	// 支付订单事件
	entity := easyfsm.NewEventEntity(paymentOrderEventName,
		func(opt *easyfsm.Param) (easyfsm.State, error) {
			param, ok := opt.Data.(orderParam)
			if !ok {
				panic("param err")
			}
			fmt.Printf("param:%+v\n", param)
			// 处理核心业务
			return paidState, nil
		})

	// 取消订单事件
	cancelEntity := easyfsm.NewEventEntity(cancelOrderEventName,
		func(opt *easyfsm.Param) (easyfsm.State, error) {
			// 处理核心业务
			param, ok := opt.Data.(orderParam)
			if !ok {
				panic("param err")
			}
			fmt.Printf("param:%+v\n", param)
			return canceled, nil
		})

	// 注册订单状态机
	easyfsm.RegisterStateMachine(businessName,
		initState,
		entity, cancelEntity)
}

func main() {

	// 正常操作

	// 第一步根据业务，以及当前状态生成fsm
	fsm := easyfsm.NewFSM(businessName, initState)

	// 第二步 调用具体
	currentState, err := fsm.Call(cancelOrderEventName,
		easyfsm.WithData(orderParam{OrderNo: "wuqinqiang050@gmail.com"}))

	fmt.Printf("[Success]call cancelOrderEventName err:%v\n", err)
	fmt.Printf("[Success]call cancelOrderEventName state:%v\n", currentState)

	//异常情况1，没有定义goods业务
	fsm = easyfsm.NewFSM("goods", paidState)
	currentState, err = fsm.Call(cancelOrderEventName,
		easyfsm.WithData(orderParam{OrderNo: "wuqinqiang050@gmail.com"}))
	fmt.Printf("[UnKnowBusiness]faild :%v\n", err)
	fmt.Printf("[UnKnowBusiness]faild state:%v\n", currentState)

	//异常情况1,没有定义状态:2
	fsm = easyfsm.NewFSM(businessName, easyfsm.State(2))
	currentState, err = fsm.Call(cancelOrderEventName,
		easyfsm.WithData(orderParam{OrderNo: "wuqinqiang050@gmail.com"}))
	fmt.Printf("[UnKnowState]faild :%v\n", err)
	fmt.Printf("[UnKnowState]faild state:%v\n", currentState)

	//异常情况2:没有定义状态1对应的发货事件
	fsm = easyfsm.NewFSM(businessName, initState)
	currentState, err = fsm.Call("shippingEvent",
		easyfsm.WithData(orderParam{OrderNo: "wuqinqiang050@gmail.com"}))
	fmt.Printf("[UnKnowEvent]faild :%v\n", err)
	fmt.Printf("[UnKnowEvent]faild state:%v\n", currentState)

}
