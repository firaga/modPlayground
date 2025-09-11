package json_rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"
)

var (
	botClient *BotClientService
	once      sync.Once
)

func GetBotClient() *BotClientService {
	return botClient
}

func InitBotClient(baseUrl string, logger *zap.Logger) {
	once.Do(func() {
		// 创建自定义传输层
		transport := &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,              // 最大空闲连接数
			MaxIdleConnsPerHost:   100,              // 每个host的最大空闲连接数
			IdleConnTimeout:       90 * time.Second, // 空闲连接超时时间
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxConnsPerHost:       100, // 每个host的最大连接数
		}

		// 创建HTTP客户端
		client := &http.Client{
			Transport: transport,
			Timeout:   30 * time.Second,
		}

		botClient = &BotClientService{
			baseURL:    baseUrl,
			httpClient: client,
			logger:     logger,
		}
	})
}

type BotClientService struct {
	baseURL    string
	httpClient *http.Client
	logger     *zap.Logger
	userLocks  sync.Map // 用户ID到锁的映射，确保同一用户同时只能有一个请求
}

// sendRpcRequest 发送JSON-RPC请求的通用方法
func (b *BotClientService) sendRpcRequest(ctx context.Context, method string, params interface{}, result interface{}) error {
	// 创建JSON-RPC请求
	rpcReq := &JsonRpcRequest{
		JsonRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      time.Now().UnixNano(),
	}

	// 序列化请求
	reqBody, err := json.Marshal(rpcReq)
	if err != nil {
		b.logger.Error("Failed to marshal request",
			zap.Error(err),
			zap.String("method", method))
		return err
	}

	// 创建HTTP请求
	req, err := http.NewRequestWithContext(ctx, "POST", b.baseURL+"/rpc", bytes.NewReader(reqBody))
	if err != nil {
		b.logger.Error("Failed to create request",
			zap.Error(err),
			zap.String("method", method))
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")

	// 发送请求
	resp, err := b.httpClient.Do(req)
	if err != nil {
		b.logger.Error("Failed to send request",
			zap.Error(err),
			zap.String("method", method))
		return err
	}
	defer resp.Body.Close()

	// 解析响应
	var rpcResp JsonRpcResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		b.logger.Error("Failed to decode response",
			zap.Error(err),
			zap.String("method", method))
		return err
	}

	// 检查错误
	if rpcResp.Error != nil {
		b.logger.Error("RPC error",
			zap.Int("code", rpcResp.Error.Code),
			zap.String("message", rpcResp.Error.Message),
			zap.String("method", method))
		return fmt.Errorf("RPC error: %s", rpcResp.Error.Message)
	}

	// 解析结果
	if err := json.Unmarshal(rpcResp.Result, result); err != nil {
		b.logger.Error("Failed to unmarshal result",
			zap.Error(err),
			zap.String("method", method))
		return err
	}

	return nil
}

func (b *BotClientService) Request(ctx context.Context, method string, params interface{}) (interface{}, error) {
	var response struct {
		Data   interface{} `json:"data"`
		Error  string      `json:"error"`
		Status int         `json:"status"`
	}

	if err := b.sendRpcRequest(ctx, method, params, &response); err != nil {
		return response.Data, err
	}

	if response.Status != 0 {
		return response.Data, fmt.Errorf("%s", response.Error)
	}
	return response.Data, nil
}
