package main

import (
	"fmt"
	"github.com/wuqinqiang/easyfsm"
	"testing"
	"time"
)

func TestOrder(t *testing.T) {
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

func TestOrderFail(t *testing.T) {
	// 第一步根据业务，以及当前状态生成fsm
	fsm := easyfsm.NewFSM(businessName, defaultState)

	currentState, err := fsm.Call(sendProductEventName,
		easyfsm.WithData(productParam{OrderNo: "zhou.jichen@gmail.com", ProductId: "1111100", ReceiverAddress: "beijing"}))

	fmt.Printf("[Success]call paymentOrderEventName err:%v\n", err)
	fmt.Printf("[Success]call paymentOrderEventName state:%v\n", currentState)
	time.Sleep(2 * time.Second)
}

func TestOrderCallback(t *testing.T) {
	// 第一步根据业务，以及当前状态生成fsm
	fsm := easyfsm.NewFSM(businessName, sentState)

	currentState, err := fsm.Call(sendProductConfirmEventName,
		easyfsm.WithData(productParam{OrderNo: "zhou.jichen@gmail.com", ProductId: "1111100", ReceiverAddress: "beijing"}))

	fmt.Printf("[Success]call paymentOrderEventName err:%v\n", err)
	fmt.Printf("[Success]call paymentOrderEventName state:%v\n", currentState)
	time.Sleep(2 * time.Second)

	currentState, err = fsm.Call(receiveProductConfirmEventName,
		easyfsm.WithData(productParam{OrderNo: "zhou.jichen@gmail.com", ProductId: "1111100", ReceiverAddress: "beijing"}))

	fmt.Printf("[Success]call paymentOrderEventName err:%v\n", err)
	fmt.Printf("[Success]call paymentOrderEventName state:%v\n", currentState)
	time.Sleep(2 * time.Second)

	//不成功,已完成订单不能取消
	currentState, err = fsm.Call(receiverCancelEventName,
		easyfsm.WithData(productParam{OrderNo: "zhou.jichen@gmail.com", ProductId: "1111100", ReceiverAddress: "beijing"}))

	fmt.Printf("[Success]call paymentOrderEventName err:%v\n", err)
	fmt.Printf("[Success]call paymentOrderEventName state:%v\n", currentState)
	time.Sleep(2 * time.Second)
}

func TestOrderCallbackCancel(t *testing.T) {
	// 第一步根据业务，以及当前状态生成fsm
	fsm := easyfsm.NewFSM(businessName, sentState)

	currentState, err := fsm.Call(sendProductConfirmEventName,
		easyfsm.WithData(productParam{OrderNo: "zhou.jichen@gmail.com", ProductId: "1111100", ReceiverAddress: "beijing"}))

	fmt.Printf("[Success]call paymentOrderEventName err:%v\n", err)
	fmt.Printf("[Success]call paymentOrderEventName state:%v\n", currentState)
	time.Sleep(2 * time.Second)

	currentState, err = fsm.Call(receiverCancelEventName,
		easyfsm.WithData(productParam{OrderNo: "zhou.jichen@gmail.com", ProductId: "1111100", ReceiverAddress: "beijing"}))

	fmt.Printf("[Success]call paymentOrderEventName err:%v\n", err)
	fmt.Printf("[Success]call paymentOrderEventName state:%v\n", currentState)
	time.Sleep(2 * time.Second)
}
