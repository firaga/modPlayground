package main

import (
	"context"
	"easyfsm/fsm"
	"fmt"
)

type testDoneFsmParam struct {
	Name string
}

func main() {
	// new一个fsm实例

	var Testing, CanOnline fsm.State
	Testing = 1
	CanOnline = 2
	// 设置状态
	sm := fsm.NewFsm(activityBase)
	fsm.Register(activityBase, int64(Testing), int64(CanOnline), testDone)

	// 设置参数

	arg := &testDoneFsmParam{
		Name: "hello",
	}

	// 执行状态转移操作

	sm.SetState(int64(Testing), int64(CanOnline))
	fmt.Println(sm.CurrentState)
	resp, err := sm.Transfer(context.Background(), arg)
	fmt.Println(resp, err)
	fmt.Println(sm.CurrentState)
}

const activityBase fsm.BusinessName = "activity"

// 其他文件中:

func testDone(ctx context.Context, req interface{}) (interface{}, error) {

	// 检查参数类型

	request, ok := req.(*testDoneFsmParam)

	if !ok {
		return 0, fsm.ParamConvertInvalidError
	}
	// 执行业务逻辑
	request.Name = "world"
	return request, nil
}
