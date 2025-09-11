package json_rpc

import "encoding/json"

type JsonRpcRequest struct {
	JsonRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int64       `json:"id"`
}
type JsonRpcResponse struct {
	JsonRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *JsonRpcError   `json:"error,omitempty"`
	ID      int64           `json:"id"`
}
type JsonRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
