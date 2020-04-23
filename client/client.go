package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mslipper/handshake/primitives"
	"io/ioutil"
	"net/http"
)

type Client struct {
	network primitives.Network
	apiKey  string
	host    string
	port    int
	c       *http.Client
}

type Opt func(c *Client)

type rpcRequest struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
	ID     int64         `json:"id"`
}

type rpcResponse struct {
	ID     int64 `json:"id"`
	Result json.RawMessage
	Error  *RPCError
}

type RPCError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (r *RPCError) Error() string {
	return r.Message
}

func WithAPIKey(apiKey string) Opt {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

func WithHTTPClient(client *http.Client) Opt {
	return func(c *Client) {
		c.c = client
	}
}

func WithNetwork(n primitives.Network) Opt {
	return func(c *Client) {
		c.network = n
	}
}

func WithPort(port int) Opt {
	return func(c *Client) {
		c.port = port
	}
}

func getJSON(client *Client, url string, result interface{}) error {
	res, err := client.c.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("non-200 status code: %d", res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, result); err != nil {
		return err
	}
	return nil
}

func postJSON(client *Client, url string, body interface{}, result interface{}) error {
	bodyB, err := json.Marshal(body)
	if err != nil {
		return err
	}
	res, err := client.c.Post(url, "application/json", bytes.NewReader(bodyB))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("non-200 status code: %d", res.StatusCode)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(resBody, result); err != nil {
		return err
	}
	return nil
}

func executeRPC(client *Client, url string, id int64, method string, resp interface{}, params ...interface{}) error {
	reqBody, err := json.Marshal(&rpcRequest{
		ID:     id,
		Method: method,
		Params: params,
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if client.apiKey != "" {
		req.SetBasicAuth("x", client.apiKey)
	}
	res, err := client.c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("non-200 status code: %d", res.StatusCode)
	}
	if resp == nil {
		return nil
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	resBody := new(rpcResponse)
	if err := json.Unmarshal(body, resBody); err != nil {
		return err
	}
	if resBody.Error != nil {
		return resBody.Error
	}
	if err := json.Unmarshal(resBody.Result, resp); err != nil {
		return err
	}
	return nil
}
