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

func TestUpsertCopyTrading(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := os.Getenv("JSON_RPC_URL")
	if baseUrl == "" {
		t.Fatal("JSON_RPC_URL is empty")
	}
	baseUrl = GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	setting := "{\"settings\": {\"1000\": 10000, \"-1000\": 10000}}"
	params := map[string]interface{}{
		"id":                         "b2f36265dcfb4eb4ae83554f6282799f",
		"userId":                     os.Getenv("JSON_RPC_USERID"),
		"username":                   os.Getenv("JSON_RPC_USERNAME"),
		"address":                    GetConfig().Address,
		"tag":                        "0x70",
		"target":                     GetConfig().Target,
		"slippage":                   2000,
		"slippageSell":               2000,
		"ratio":                      100,
		"priorityFee":                100000000,
		"priorityFeeSell":            100000000,
		"approveGasPrice":            100000000,
		"totalUpperLimit":            "100000000000000",
		"lowerLimitOfOneTransaction": -1,
		"upperLimitOfOneTransaction": -1,
		"onlySell":                   false,
		"enabled":                    true,
		"copyPancake":                false,
		"copyFourMeme":               false,
		"copySell":                   true,
		"autoSell":                   true,
		"autoSellParams":             setting,
	}
	res, err := c.Request(context.Background(), "upsertCopyTrading", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestDeleteCopyTrading(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := os.Getenv("JSON_RPC_URL")
	if baseUrl == "" {
		t.Fatal("JSON_RPC_URL is empty")
	}
	baseUrl = GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		"userId": os.Getenv("JSON_RPC_USERID"),
		//"userId":   "user1",
		"username": os.Getenv("JSON_RPC_USERNAME"),
		"id":       "1",
	}
	res, err := c.Request(context.Background(), "deleteCopyTrading", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestGetCopyTrading(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := os.Getenv("JSON_RPC_URL")
	if baseUrl == "" {
		t.Fatal("JSON_RPC_URL is empty")
	}
	baseUrl = GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		//"userId": os.Getenv("JSON_RPC_USERID"),
		"userId": "6470729468",
		//"username": os.Getenv("JSON_RPC_USERNAME"),
		"id": "e63f8da87d0249ed839c22edc626c25b",
	}
	res, err := c.Request(context.Background(), "getCopyTrading", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestGetCopyTradingList2(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := os.Getenv("JSON_RPC_URL")
	if baseUrl == "" {
		t.Fatal("JSON_RPC_URL is empty")
	}
	baseUrl = GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		"userId": os.Getenv("JSON_RPC_USERID"),
		//"userId":   "user1",
		"username": os.Getenv("JSON_RPC_USERNAME"),
		"id":       "2",
	}
	res, err := c.Request(context.Background(), "getCopyTradingList", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestSwitchCopyTrading(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := os.Getenv("JSON_RPC_URL")
	if baseUrl == "" {
		t.Fatal("JSON_RPC_URL is empty")
	}
	baseUrl = GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		"userId": os.Getenv("JSON_RPC_USERID"),
		//"userId":   "user1",
		"username": os.Getenv("JSON_RPC_USERNAME"),
		"id":       "2",
		"enabled":  true,
	}
	res, err := c.Request(context.Background(), "switchCopyTrading", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestUpsertAllOnlySell(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := os.Getenv("JSON_RPC_URL")
	if baseUrl == "" {
		t.Fatal("JSON_RPC_URL is empty")
	}
	baseUrl = GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		"userId": os.Getenv("JSON_RPC_USERID"),
		//"userId":   "user1",
		"username": os.Getenv("JSON_RPC_USERNAME"),
		"onlySell": true,
	}
	res, err := c.Request(context.Background(), "upsertAllOnlySell", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestEnableAllCopyTrading(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := os.Getenv("JSON_RPC_URL")
	if baseUrl == "" {
		t.Fatal("JSON_RPC_URL is empty")
	}
	baseUrl = GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
	}
	res, err := c.Request(context.Background(), "enableAllCopyTrading", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestDisableAllCopyTrading(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := os.Getenv("JSON_RPC_URL")
	if baseUrl == "" {
		t.Fatal("JSON_RPC_URL is empty")
	}
	baseUrl = GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
	}
	res, err := c.Request(context.Background(), "disableAllCopyTrading", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestSetAccountConfig(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := os.Getenv("JSON_RPC_URL")
	if baseUrl == "" {
		t.Fatal("JSON_RPC_URL is empty")
	}
	baseUrl = GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
		"config": map[string]interface{}{
			"copyTradingOnceBuySwitch": false,
			"copyTradingMsgLevel":      1,
		},
	}
	res, err := c.Request(context.Background(), "setAccountConfig", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestGetAccountConfig(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
	}
	res, err := c.Request(context.Background(), "getAccountInfo", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestGetGetCopyTradingFollowers(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
	}
	res, err := c.Request(context.Background(), "getGetCopyTradingFollowers", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestGetGetCopyTradingTradeStats(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	userId := GetConfig().UserId
	params := map[string]interface{}{
		"userId": userId,
		//"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
	}
	res, err := c.Request(context.Background(), "getGetCopyTradingTradeStats", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestAddExcludeToken(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		//"userId": WJT,
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
		"token":    "KGST7",
	}
	res, err := c.Request(context.Background(), "addExcludeToken", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestDelExcludeToken(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		//"userId": WJT,
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
		"symbol":   "KGST",
	}
	res, err := c.Request(context.Background(), "delExcludeToken", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestGetExcludeTokenList(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		//"userId": WJT,
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
	}
	res, err := c.Request(context.Background(), "getExcludeTokenList", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestGetToken(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		//"userId": WJT,
		"userId":       os.Getenv("JSON_RPC_USERID"),
		"username":     os.Getenv("JSON_RPC_USERNAME"),
		"tokenAddress": GetConfig().Token,
	}
	res, err := c.Request(context.Background(), "getToken", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}
