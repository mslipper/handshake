package client

import (
	"encoding/json"
	"fmt"
	"github.com/mslipper/handshake/primitives"
	"net/http"
	"strconv"
	"sync/atomic"
)

type RPCClient struct {
	*Client
	rpcID int64
}

func NewRPC(host string, opts ...Opt) *RPCClient {
	c := &RPCClient{
		Client: &Client{
			network: primitives.NetworkMainnet,
			host:    host,
			c: http.DefaultClient,
		},
	}
	for _, opt := range opts {
		opt(c.Client)
	}
	return c
}

func (r *RPCClient) Stop() error {
	var res json.RawMessage
	if err := r.executeRPC("stop", res); err != nil {
		return err
	}
	return nil
}

func (r *RPCClient) GetInfo() (*InfoResult, error) {
	res := new(InfoResult)
	if err := r.executeRPC("getinfo", res); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RPCClient) GetMemoryInfo() (*MemoryInfoResult, error) {
	res := new(MemoryInfoResult)
	if err := r.executeRPC("getmemoryinfo", res); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RPCClient) SetLogLevel(level string) error {
	if err := r.executeRPC("setloglevel", nil, level); err != nil {
		return err
	}
	return nil
}

func (r *RPCClient) ValidateAddress(address string) (*ValidateAddressResult, error) {
	res := new(ValidateAddressResult)
	if err := r.executeRPC("validateaddress", res, address); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RPCClient) CreateMultisig(numRequired int, keys []string) (*CreateMultisigResponse, error) {
	res := new(CreateMultisigResponse)
	if err := r.executeRPC("createmultisig", res, numRequired, keys); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RPCClient) SignMessageWithPrivkey(privKey string, message string) (string, error) {
	var res string
	if err := r.executeRPC("signmessagewithprivkey", &res, privKey, message); err != nil {
		return "", err
	}
	return res, nil
}

func (r *RPCClient) VerifyMessage(address string, signature string, message string) (bool, error) {
	var res bool
	if err := r.executeRPC("verifymessage", res, address, signature, message); err != nil {
		return false, err
	}
	return res, nil
}

func (r *RPCClient) SetMockTime(timestamp int) error {
	if err := r.executeRPC("setmocktime", nil, strconv.Itoa(timestamp)); err != nil {
		return err
	}
	return nil
}

func (r *RPCClient) PruneBlockchain() error {
	if err := r.executeRPC("pruneblockchain", nil); err != nil {
		return err
	}
	return nil
}

func (r *RPCClient) InvalidateBlock(blockHash string) error {
	if err := r.executeRPC("invalidateblock", nil, blockHash); err != nil {
		return err
	}
	return nil
}

func (r *RPCClient) ReconsiderBlock(blockHash string) error {
	if err := r.executeRPC("reconsiderblock", nil, blockHash); err != nil {
		return err
	}
	return nil
}

func (r *RPCClient) GetBlockchainInfo() (*GetBlockchainInfoResponse, error) {
	res := new(GetBlockchainInfoResponse)
	if err := r.executeRPC("verifymessage", res); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RPCClient) GetBestBlockHash() (string, error) {
	var res string
	if err := r.executeRPC("getbestblockhash", &res); err != nil {
		return "", err
	}
	return res, nil
}

func (r *RPCClient) GetBlockCount() (int, error) {
	var res int
	if err := r.executeRPC("getblockcount", &res); err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RPCClient) GetBlockByHashWithoutTxs(blockHash string) (*BlockWithoutTxsResponse, error) {
	res := new(BlockWithoutTxsResponse)
	if err := r.executeRPC("getblock", res, blockHash, true, false); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RPCClient) GetBlockByHashWithTxs(blockHash string) (*BlockWithTxsResponse, error) {
	res := new(BlockWithTxsResponse)
	if err := r.executeRPC("getblock", res, blockHash, true, true); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RPCClient) GetBlockHexByHash(blockHash string) (string, error) {
	var res string
	if err := r.executeRPC("getblock", &res, blockHash, false, false); err != nil {
		return "", err
	}
	return res, nil
}

func (r *RPCClient) GetBlockByHeightWithoutTxs(blockHeight int) (*BlockWithoutTxsResponse, error) {
	res := new(BlockWithoutTxsResponse)
	if err := r.executeRPC("getblockbyheight", res, blockHeight, true, false); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RPCClient) GetBlockByHeightWithTxs(blockHeight int) (*BlockWithTxsResponse, error) {
	res := new(BlockWithTxsResponse)
	if err := r.executeRPC("getblockbyheight", res, blockHeight, true, true); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RPCClient) GetBlockHexByHeight(blockHeight int) (string, error) {
	var res string
	if err := r.executeRPC("getblockbyheight", &res, blockHeight, false, false); err != nil {
		return "", err
	}
	return res, nil
}

func (r *RPCClient) GetBlockHashByHeight(blockHeight int) (string, error) {
	var res string
	if err := r.executeRPC("getblockhash", &res, blockHeight); err != nil {
		return "", err
	}
	return res, nil
}

func (r *RPCClient) GetBlockHeaderByHash(blockHash string) (*BlockHeaderResponse, error) {
	res := new(BlockHeaderResponse)
	if err := r.executeRPC("getblockheader", res, blockHash, true); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RPCClient) GetBlockHeaderHexByHash(blockHash string) (string, error) {
	var res string
	if err := r.executeRPC("getblockheader", &res, blockHash, false); err != nil {
		return "", err
	}
	return res, nil
}

func (r *RPCClient) GetChainTips() ([]*ChainTipResponse, error) {
	res := make([]*ChainTipResponse, 0)
	if err := r.executeRPC("getchaintips", res); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RPCClient) GetDifficulty() (float64, error) {
	var res float64
	if err := r.executeRPC("getdifficulty", &res); err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RPCClient) executeRPC(method string, res interface{}, args ...interface{}) error {
	return executeRPC(r.Client, r.makeURL(), atomic.AddInt64(&r.rpcID, 1), method, res, args...)
}

func (r *RPCClient) makeURL() string {
	var port int
	if r.port != 0 {
		port = r.port
	} else {
		port = r.network.RPCPort()
	}
	return fmt.Sprintf("%s:%d", r.host, port)
}
