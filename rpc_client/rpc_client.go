package rpc_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type RPCClient struct {
	url      string
	username string
	password string
}

func NewRPCClient(url, username, password string) (*RPCClient, error) {
	if url == "" || username == "" || password == "" {
		return nil, errors.New("invalid RPC client parameters")
	}
	return &RPCClient{
		url:      url,
		username: username,
		password: password,
	}, nil
}

func (c *RPCClient) call(method string, params []interface{}) (json.RawMessage, error) {
	type rpcRequest struct {
		JSONRPC string        `json:"jsonrpc"`
		Method  string        `json:"method"`
		Params  []interface{} `json:"params"`
		ID      int           `json:"id"`
	}

	type rpcResponse struct {
		Result json.RawMessage `json:"result"`
		Error  *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
		ID int `json:"id"`
	}

	requestBody, err := json.Marshal(
		rpcRequest{
			JSONRPC: "1.0",
			Method:  method,
			Params:  params,
			ID:      1,
		},
	)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("RPC server returned HTTP status %s", resp.Status)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rpcResp rpcResponse
	if err := json.Unmarshal(responseBody, &rpcResp); err != nil {
		return nil, err
	}

	if rpcResp.Error != nil {
		return nil, fmt.Errorf("RPC error %d: %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}

	return rpcResp.Result, nil
}

func (c *RPCClient) GetBlockchainInfo() (map[string]interface{}, error) {
	response, err := c.call("getblockchaininfo", nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *RPCClient) GetBalance() (float64, error) {
	response, err := c.call("getbalance", nil)
	if err != nil {
		return 0, err
	}

	var result float64
	if err := json.Unmarshal(response, &result); err != nil {
		return 0, err
	}

	return result, nil
}

func (c *RPCClient) SendTransaction(from, to string, amount float64) (string, error) {
	params := []interface{}{from, map[string]interface{}{to: amount}}
	response, err := c.call("sendtoaddress", params)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %w", err)
	}

	var txID string
	if err := json.Unmarshal(response, &txID); err != nil {
		return "", fmt.Errorf("failed to parse transaction ID: %w", err)
	}

	return txID, nil
}
