package json_rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCommon(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := os.Getenv("JSON_RPC_URL")
	if baseUrl == "" {
		t.Fatal("JSON_RPC_URL is empty")
	}
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
	}
	res, err := c.Request(context.Background(), "getCleanupPlan", params)
	assert.Nil(t, err)
	spew.Dump(res)
}

func TestGetTokenList(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := os.Getenv("JSON_RPC_URL")
	if baseUrl == "" {
		t.Fatal("JSON_RPC_URL is empty")
	}
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
	}
	res, err := c.Request(context.Background(), "getTokenList", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestGetTokenInfo(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := os.Getenv("JSON_RPC_URL")
	if baseUrl == "" {
		t.Fatal("JSON_RPC_URL is empty")
	}
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
		"tokenId":  "57QjGXRHGcDQjFrHNQcFXmL5yXCPdAX2RW75DqKYjsAz",
	}
	res, err := c.Request(context.Background(), "getTokenInfo", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}
func TestGetLimitOrderList(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := os.Getenv("JSON_RPC_URL")
	if baseUrl == "" {
		t.Fatal("JSON_RPC_URL is empty")
	}
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
		"tokenId":  "57QjGXRHGcDQjFrHNQcFXmL5yXCPdAX2RW75DqKYjsAz",
	}
	res, err := c.Request(context.Background(), "getLimitOrderList", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestGetSniperTasks(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := os.Getenv("JSON_RPC_URL")
	if baseUrl == "" {
		t.Fatal("JSON_RPC_URL is empty")
	}
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
	}
	res, err := c.Request(context.Background(), "getSniperTasks", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestGetCopyTradingList(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := os.Getenv("JSON_RPC_URL")
	if baseUrl == "" {
		t.Fatal("JSON_RPC_URL is empty")
	}
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		//"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
	}
	res, err := c.Request(context.Background(), "getCopyTradingList", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}
