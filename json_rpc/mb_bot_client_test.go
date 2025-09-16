package json_rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetTokenListWithPublicKeys(t *testing.T) {
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
		"publicKeys": []string{
			"3xuhHkXmw7VAs6LoNTRgN7TxaWTRfwoy53dvPdmVmpFf",
		},
	}
	res, err := c.Request(context.Background(), "getTokenListWithPublicKeys", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestGetTokenInfoWithPublicKey(t *testing.T) {
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
		//"publicKey": "3xuhHkXmw7VAs6LoNTRgN7TxaWTRfwoy53dvPdmVmpFf",
		"publicKey": "23KZzQwjzLFnMwwJ622eJmJcWsHXPbFYRcDgXVU8EfPC",
	}
	res, err := c.Request(context.Background(), "getTokenInfo", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}
func TestGetLimitOrderListWithPublicKey(t *testing.T) {
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
		//"publicKeys": []string{
		//	"3xuhHkXmw7VAs6LoNTRgN7TxaWTRfwoy53dvPdmVmpFf",
		//},
		//"publicKey": "3xuhHkXmw7VAs6LoNTRgN7TxaWTRfwoy53dvPdmVmpFf",
	}
	res, err := c.Request(context.Background(), "getLimitOrderList", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestGetSniperTasksWithPublicKey(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := os.Getenv("JSON_RPC_URL")
	if baseUrl == "" {
		t.Fatal("JSON_RPC_URL is empty")
	}
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		"userId":    os.Getenv("JSON_RPC_USERID"),
		"username":  os.Getenv("JSON_RPC_USERNAME"),
		"publicKey": "3xuhHkXmw7VAs6LoNTRgN7TxaWTRfwoy53dvPdmVmpFf",
	}
	res, err := c.Request(context.Background(), "getSniperTasks", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}
