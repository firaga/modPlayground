package sol

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"solbot_client"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCommon(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := json_rpc.GetConfig().Url
	json_rpc.InitBotClient(baseUrl, logger)
	c := json_rpc.GetBotClient()
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
	baseUrl := json_rpc.GetConfig().Url
	json_rpc.InitBotClient(baseUrl, logger)
	c := json_rpc.GetBotClient()
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
	baseUrl := json_rpc.GetConfig().Url
	json_rpc.InitBotClient(baseUrl, logger)
	c := json_rpc.GetBotClient()
	params := map[string]interface{}{
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
		"tokenId":  json_rpc.GetConfig().Token,
	}
	res, err := c.Request(context.Background(), "getTokenInfo", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}
func TestGetLimitOrderList(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := json_rpc.GetConfig().Url
	json_rpc.InitBotClient(baseUrl, logger)
	c := json_rpc.GetBotClient()
	params := map[string]interface{}{
		"userId": "",
		//"username": os.Getenv("JSON_RPC_USERNAME"),
	}
	res, err := c.Request(context.Background(), "getLimitOrderList", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestGetSniperTasks(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := json_rpc.GetConfig().Url
	json_rpc.InitBotClient(baseUrl, logger)
	c := json_rpc.GetBotClient()
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
	baseUrl := json_rpc.GetConfig().Url
	json_rpc.InitBotClient(baseUrl, logger)
	c := json_rpc.GetBotClient()
	params := map[string]interface{}{
		//"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
	}
	res, err := c.Request(context.Background(), "getCopyTradingList", params)
	assert.Nil(t, err)
	spew.Dump(res)
}

func TestListTradesSummaryHandler(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := json_rpc.GetConfig().Url
	json_rpc.InitBotClient(baseUrl, logger)
	c := json_rpc.GetBotClient()
	params := map[string]interface{}{
		//"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
		"address":  json_rpc.GetConfig().Address,
		"day":      7,
		"page":     1,
		"pageSize": 10,
	}
	res, err := c.Request(context.Background(), "listTradesSummary", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}
