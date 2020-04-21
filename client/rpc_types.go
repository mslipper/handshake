package client

type InfoResult struct {
	Version         string  `json:"version"`
	Protocolversion int     `json:"protocolversion"`
	Walletversion   int     `json:"walletversion"`
	Balance         int     `json:"balance"`
	Blocks          int     `json:"blocks"`
	Timeoffset      int     `json:"timeoffset"`
	Connections     int     `json:"connections"`
	Proxy           string  `json:"proxy"`
	Difficulty      float64 `json:"difficulty"`
	Testnet         bool    `json:"testnet"`
	Keypoololdest   int     `json:"keypoololdest"`
	Keypoolsize     int     `json:"keypoolsize"`
	UnlockedUntil   int     `json:"unlocked_until"`
	Paytxfee        float64 `json:"paytxfee"`
	Relayfee        float64 `json:"relayfee"`
	Errors          string  `json:"errors"`
}

type MemoryInfoResult struct {
	Total       int `json:"total"`
	JsHeap      int `json:"jsHeap"`
	JsHeapTotal int `json:"jsHeapTotal"`
	NativeHeap  int `json:"nativeHeap"`
	External    int `json:"external"`
}

type ValidateAddressResult struct {
	IsValid        bool   `json:"isvalid"`
	Address        string `json:"address"`
	IsScript       bool   `json:"isscript"`
	IsSpendable    bool   `json:"isspendable"`
	WitnessVersion int    `json:"witness_version"`
	WitnessProgram string `json:"witness_program"`
}

type CreateMultisigResponse struct {
	Address      string `json:"address"`
	RedeemScript string `json:"redeemScript"`
}
