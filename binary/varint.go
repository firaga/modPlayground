package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	buf := make([]byte, binary.MaxVarintLen64)
	for _, x := range []int64{-65, 1, 2, 127, 128, 255, 256, 1732} {
		n := binary.PutVarint(buf, x)
		fmt.Print(x, "输出的可变长度为：", n, "，十六进制为：")
		fmt.Printf("%b\n", buf[:n])
	}
}

// 输出
//-65输出的可变长度为：2，十六进制为：8101
//1输出的可变长度为：1，十六进制为：02
//2输出的可变长度为：1，十六进制为：04
//127输出的可变长度为：2，十六进制为：fe01
//128输出的可变长度为：2，十六进制为：8002
//255输出的可变长度为：2，十六进制为：fe03
//256输出的可变长度为：2，十六进制为：8004
