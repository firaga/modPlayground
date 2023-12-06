// 作者：程序员麻辣烫
// 链接：https://zhuanlan.zhihu.com/p/518225089
// 来源：知乎
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
package main

import (
	"context"
	"fmt"
	"github.com/looplab/fsm"
)

type Door struct {
	To  string
	FSM *fsm.FSM
}

func NewDoor(to string) *Door {
	d := &Door{
		To: to,
	}

	d.FSM = fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{
			//指定状态
			"leave_closed": func(ctx context.Context, e *fsm.Event) { d.leaveClose(e) },
			"before_open":  func(ctx context.Context, e *fsm.Event) { d.beforeOpen(e) },
			"enter_open":   func(ctx context.Context, e *fsm.Event) { d.enterOpen(e) },
			"after_open":   func(ctx context.Context, e *fsm.Event) { d.afterOpen(e) },
			//通用状态
			"enter_state": func(ctx context.Context, e *fsm.Event) { d.enterState(e) },
		},
	)
	return d
}

func (d *Door) beforeOpen(e *fsm.Event) {
	fmt.Printf("beforeOpen, The door to %s is %s\n", d.To, e.Dst)
}

func (d *Door) enterOpen(e *fsm.Event) {
	fmt.Printf("enterOpen, The door to %s is %s\n", d.To, e.Dst)
}

func (d *Door) afterOpen(e *fsm.Event) {
	fmt.Printf("afterOpen, The door to %s is %s\n", d.To, e.Dst)
}

func (d *Door) leaveOpen(e *fsm.Event) {
	fmt.Printf("leaveOpen, The door to %s is %s\n", d.To, e.Dst)
}

func (d *Door) leaveClose(e *fsm.Event) {
	fmt.Printf("leaveClose, The door to %s is %s\n", d.To, e.Dst)
}

func (d *Door) enterState(e *fsm.Event) {
	fmt.Printf("enterState, The door to %s is %s\n", d.To, e.Dst)
}

func main() {
	ctx := context.Background()
	door := NewDoor("heaven")
	fmt.Println("当前状态为：" + door.FSM.Current())

	err := door.FSM.Event(ctx, "open")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("当前状态为：" + door.FSM.Current())
	err = door.FSM.Event(ctx, "close")
	if err != nil {
		fmt.Println(err)
	}
}
