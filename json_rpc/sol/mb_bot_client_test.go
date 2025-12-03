package sol

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"solbot_client"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetTokenListMultiWallet(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := json_rpc.GetConfig().Url
	json_rpc.InitBotClient(baseUrl, logger)
	c := json_rpc.GetBotClient()
	params := map[string]interface{}{
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
		"address":  json_rpc.GetConfig().Address,
	}
	res, err := c.Request(context.Background(), "upsertCopyTrading", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestGetTokenInfoMultiWallet(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := json_rpc.GetConfig().Url
	json_rpc.InitBotClient(baseUrl, logger)
	c := json_rpc.GetBotClient()
	params := map[string]interface{}{
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
		"tokenId":  json_rpc.GetConfig().Token,
	}
	res, err := c.Request(context.Background(), "getTokenInfoMultiWallet", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestGetTokenInfoWithPublicKey(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := json_rpc.GetConfig().Url
	json_rpc.InitBotClient(baseUrl, logger)
	c := json_rpc.GetBotClient()
	params := map[string]interface{}{
		"userId":    os.Getenv("JSON_RPC_USERID"),
		"username":  os.Getenv("JSON_RPC_USERNAME"),
		"tokenId":   json_rpc.GetConfig().Token,
		"publicKey": json_rpc.GetConfig().Address,
	}
	res, err := c.Request(context.Background(), "getTokenInfo", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}
func TestGetLimitOrderListWithPublicKey(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := json_rpc.GetConfig().Url
	json_rpc.InitBotClient(baseUrl, logger)
	c := json_rpc.GetBotClient()
	params := map[string]interface{}{
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
		"tokenId":  "",
		"publicKeys": []string{
			"",
		},
	}
	res, err := c.Request(context.Background(), "getLimitOrderList", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}

func TestGetSniperTasksWithPublicKey(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := json_rpc.GetConfig().Url
	json_rpc.InitBotClient(baseUrl, logger)
	c := json_rpc.GetBotClient()
	params := map[string]interface{}{
		"userId":    os.Getenv("JSON_RPC_USERID"),
		"username":  os.Getenv("JSON_RPC_USERNAME"),
		"publicKey": "",
	}
	res, err := c.Request(context.Background(), "getSniperTasks", params)
	assert.Nil(t, err)
	//spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}
