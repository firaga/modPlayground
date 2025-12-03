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

func TestGetWalletList(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	baseUrl := GetConfig().Url
	InitBotClient(baseUrl, logger)
	c := GetBotClient()
	params := map[string]interface{}{
		//"userId": WJT,
		"userId":   os.Getenv("JSON_RPC_USERID"),
		"username": os.Getenv("JSON_RPC_USERNAME"),
	}
	res, err := c.Request(context.Background(), "getWalletList", params)
	assert.Nil(t, err)
	spew.Dump(res)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}
