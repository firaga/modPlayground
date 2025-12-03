package json_rpc

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	cfg := GetConfig()
	// 3. 使用配置变量
	fmt.Printf("数据库地址：%s/%s\n", cfg.UserId, cfg.Url)
}
