package client

import (
	"errors"
	"fmt"
	"handshake/primitives"
)

type RESTClient struct {
	*Client
}

func NewREST(host string, opts ...Opt) *RESTClient {
	c := &RESTClient{
		Client: &Client{
			network: primitives.NetworkMainnet,
			host:    host,
		},
	}
	for _, opt := range opts {
		opt(c.Client)
	}
	return c
}

func (n *RESTClient) GetInfo() (*NodeInfo, error) {
	info := new(NodeInfo)
	if err := n.getJSON("", info); err != nil {
		return nil, err
	}
	return info, nil
}

func (n *RESTClient) GetMempoolSnapshot() ([]string, error) {
	var out []string
	if err := n.getJSON("mempool", out); err != nil {
		return nil, err
	}
	return out, nil
}

func (n *RESTClient) GetMempoolRejectsFilter() (*MempoolRejectsFilterInfo, error) {
	info := new(MempoolRejectsFilterInfo)
	if err := n.getJSON("mempool/invalid", info); err != nil {
		return nil, err
	}
	return info, nil
}

func (n *RESTClient) TestMempoolRejectsFilter(hash string) (bool, error) {
	info := new(struct {
		Invalid bool `json:"invalid"`
	})
	if err := n.getJSON(fmt.Sprintf("mempool/invalid/%s", hash), info); err != nil {
		return false, err
	}
	return info.Invalid, nil
}

func (n *RESTClient) GetBlockByHash(hash string) (*RESTBlock, error) {
	block := new(RESTBlock)
	if err := n.getJSON(fmt.Sprintf("block/%s", hash), block); err != nil {
		return nil, err
	}
	return block, nil
}

func (n *RESTClient) GetBlockByHeight(height int) (*RESTBlock, error) {
	if height < 0 {
		return nil, errors.New("cannot set a negative height")
	}
	block := new(RESTBlock)
	if err := n.getJSON(fmt.Sprintf("block/%d", height), block); err != nil {
		return nil, err
	}
	return block, nil
}

func (n *RESTClient) BroadcastTransaction(tx string) error {
	body := struct {
		Tx string `json:"tx"`
	}{
		tx,
	}
	res := new(struct {
		Success bool `json:"success"`
	})
	if err := n.postJSON("broadcast", body, res); err != nil {
		return err
	}
	if !res.Success {
		return errors.New("error broadcasting transaction, check HSD logs")
	}
	return nil
}

func (n *RESTClient) BroadcastClaim(claim string) error {
	body := struct {
		Claim string `json:"claim"`
	}{
		claim,
	}
	res := new(struct {
		Success bool `json:"success"`
	})
	if err := n.postJSON("claim", body, res); err != nil {
		return err
	}
	if !res.Success {
		return errors.New("error broadcasting claim, check HSD logs")
	}
	return nil
}

func (n *RESTClient) EstimateFee(blocks int) (uint64, error) {
	if blocks < 0 {
		return 0, errors.New("blocks cannot be negative")
	}
	res := new(struct {
		Rate uint64 `json:"rate"`
	})
	if err := n.getJSON(fmt.Sprintf("fee?blocks=%d", blocks), res); err != nil {
		return 0, err
	}
	return res.Rate, nil
}

func (n *RESTClient) ResetBlockchain(height int) error {
	if height < 0 {
		return errors.New("cannot set a zero height")
	}
	body := struct {
		Height int `json:"height"`
	}{
		height,
	}
	res := new(struct {
		Success bool `json:"success"`
	})
	if err := n.postJSON("reset", body, res); err != nil {
		return err
	}
	if !res.Success {
		return errors.New("error resetting chain, check HSD logs")
	}
	return nil
}

func (n *RESTClient) GetCoinByOutpoint(hash string, index int) (*Coin, error) {
	res := new(Coin)
	if err := n.getJSON(fmt.Sprintf("coin/%s/%d", hash, index), res); err != nil {
		return nil, err
	}
	return res, nil
}

func (n *RESTClient) GetCoinsByAddress(address string) ([]*Coin, error) {
	var res []*Coin
	if err := n.getJSON(fmt.Sprintf("coins/address/%s", address), res); err != nil {
		return nil, err
	}
	return res, nil
}

func (n *RESTClient) GetTransactionByHash(hash string) (*RESTTransaction, error) {
	res := new(RESTTransaction)
	if err := n.getJSON(fmt.Sprintf("tx/%s", hash), res); err != nil {
		return nil, err
	}
	return res, nil
}

func (n *RESTClient) GetTransactionsByAddress(address string) ([]*RESTTransaction, error) {
	var res []*RESTTransaction
	if err := n.getJSON(fmt.Sprintf("tx/address/%s", address), res); err != nil {
		return nil, err
	}
	return res, nil
}

func (n *RESTClient) getJSON(path string, res interface{}) error {
	return getJSON(n.Client, n.makeURL(path), res)
}

func (n *RESTClient) postJSON(path string, body interface{}, res interface{}) error {
	return postJSON(n.Client, n.makeURL(path), body, res)
}

func (n *RESTClient) makeURL(path string) string {
	var port int
	if n.port != 0 {
		port = n.port
	} else {
		port = n.network.RPCPort()
	}
	return fmt.Sprintf("%s:%d/%s", n.host, port, path)
}
