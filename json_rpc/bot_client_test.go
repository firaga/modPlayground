package json_rpc

import (
	"context"
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
