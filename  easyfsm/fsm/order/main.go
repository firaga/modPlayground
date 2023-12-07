package main

import (
	"fmt"
	"github.com/wuqinqiang/easyfsm"
	"time"
)

var (
	businessName     easyfsm.BusinessName = "order"
	defaultState     easyfsm.State        = 0 //初始状态
	waitSendState    easyfsm.State        = 1 //等待卖家发货
	sentState        easyfsm.State        = 2 //卖家已发货
	waitReceiveState easyfsm.State        = 3 //等待买家收货
	completeState    easyfsm.State        = 4 //完成订单
	cancelState      easyfsm.State        = 5 //取消订单
)

type LogHook struct {
}

func (h LogHook) Before(opt *easyfsm.Param) {
	fmt.Println("事件执行前")
}

func (h LogHook) After(opt easyfsm.Param, state easyfsm.State, err error) {
	fmt.Println("事件执行后")
}

type SentHook struct {
}

func (h SentHook) Before(opt *easyfsm.Param) {
	fmt.Println("卖家已发货 - 发货前")
}

func (h SentHook) After(opt easyfsm.Param, state easyfsm.State, err error) {
	fmt.Println("卖家已发货 - 发货后")
}

type (
	orderParam struct {
		OrderNo string
	}

	productParam struct {
		OrderNo         string
		ProductId       string
		ReceiverAddress string
	}
)

var (
	createOrderEventName           easyfsm.EventName = "createOrderEventName"
	sendProductEventName           easyfsm.EventName = "sendProductEventName"
	sendProductConfirmEventName    easyfsm.EventName = "sendProductConfirmEventName"
	receiveProductConfirmEventName easyfsm.EventName = "receiveProductConfirmEventName"
	receiverCancelEventName        easyfsm.EventName = "receiverCancelEventName"
)

func init() {
	//创建订单
	createOrder := easyfsm.NewEventEntity(createOrderEventName,
		func(opt *easyfsm.Param) (easyfsm.State, error) {
			param, ok := opt.Data.(orderParam)
			if !ok {
				panic("param err")
			}
			fmt.Printf("param:%+v\n", param)
			// 处理核心业务
			return waitSendState, nil
		}, easyfsm.WithHook(LogHook{}))
	//卖家发货
	sendProduct := easyfsm.NewEventEntity(sendProductEventName,
		func(opt *easyfsm.Param) (easyfsm.State, error) {
			param, ok := opt.Data.(productParam)
			if !ok {
				panic("param err")
			}
			fmt.Printf("param:%+v\n", param)
			// 处理核心业务
			return sentState, nil
		}, easyfsm.WithHook(LogHook{}))
	//确认已发货
	confirmSentProduct := easyfsm.NewEventEntity(sendProductConfirmEventName,
		func(opt *easyfsm.Param) (easyfsm.State, error) {
			param, ok := opt.Data.(productParam)
			if !ok {
				panic("param err")
			}
			fmt.Printf("param:%+v\n", param)
			// 处理核心业务
			return waitReceiveState, nil
		}, easyfsm.WithHook(LogHook{}))
	//确认收货
	confirmReceiveProduct := easyfsm.NewEventEntity(receiveProductConfirmEventName,
		func(opt *easyfsm.Param) (easyfsm.State, error) {
			param, ok := opt.Data.(productParam)
			if !ok {
				panic("param err")
			}
			fmt.Printf("param:%+v\n", param)
			// 处理核心业务
			return completeState, nil
		}, easyfsm.WithHook(LogHook{}))
	//(系统,卖家,买家)取消订单
	cancelOrder := easyfsm.NewEventEntity(receiverCancelEventName,
		func(opt *easyfsm.Param) (easyfsm.State, error) {
			param, ok := opt.Data.(productParam)
			if !ok {
				panic("param err")
			}
			fmt.Printf("param:%+v\n", param)
			// 处理核心业务
			return cancelState, nil
		}, easyfsm.WithHook(LogHook{}))
	// 注册订单状态机
	easyfsm.RegisterStateMachine(businessName,
		defaultState,
		createOrder,
	)
	easyfsm.RegisterStateMachine(businessName,
		waitSendState,
		sendProduct,
		cancelOrder,
	)
	easyfsm.RegisterStateMachine(businessName,
		sentState,
		confirmSentProduct, //等待steam回调状态"等待买家收货",确认卖家确实已经发货了
		cancelOrder,
	)
	easyfsm.RegisterStateMachine(businessName,
		waitReceiveState,
		confirmReceiveProduct,
		cancelOrder,
	)
	easyfsm.RegisterStateMachine(businessName,
		completeState,
	)
	easyfsm.RegisterStateMachine(businessName,
		cancelState,
	)
}

func main() {
	// 第一步根据业务，以及当前状态生成fsm
	fsm := easyfsm.NewFSM(businessName, defaultState)

	currentState, err := fsm.Call(createOrderEventName,
		easyfsm.WithData(orderParam{OrderNo: "zhou.jichen@gmail.com"}))

	fmt.Printf("[Success]call paymentOrderEventName err:%v\n", err)
	fmt.Printf("[Success]call paymentOrderEventName state:%v\n", currentState)
	time.Sleep(2 * time.Second)

	currentState, err = fsm.Call(sendProductEventName,
		easyfsm.WithData(productParam{OrderNo: "zhou.jichen@gmail.com", ProductId: "1111100", ReceiverAddress: "beijing"}))

	fmt.Printf("[Success]call paymentOrderEventName err:%v\n", err)
	fmt.Printf("[Success]call paymentOrderEventName state:%v\n", currentState)
	time.Sleep(2 * time.Second)
}
