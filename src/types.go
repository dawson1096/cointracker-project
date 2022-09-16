package src

type Ledger struct {
	Address string       `json:"address"`
	Balance int          `json:"balance"`
	Txns    []TxResponse `json:"txns"`
}

type SyncResponse struct {
	Hash160        string       `json:"hash160"`
	Address        string       `json:"address"`
	N_tx           int          `json:"n_tx"`
	N_unredeemed   int          `json:"n_unredeemed"`
	Total_received int          `json:"total_received"`
	Total_sent     int          `json:"total_sent"`
	Final_balance  int          `json:"final_balance"`
	Txs            []TxResponse `json:"txs"`
}

type LatestBlock struct {
	Hash        string `json:"hash"`
	Time        int    `json:"time"`
	Block_index int    `json:"block_index"`
	Height      int    `json:"height"`
	TxIndexes   []int  `json:"txIndexes"`
}

type BlockResponse struct {
	Hash        string       `json:"hash"`
	Ver         int          `json:"ver"`
	Prev_block  string       `json:"prev_block"`
	Mrkl_root   string       `json:"mrkl_root"`
	Time        int          `json:"time"`
	Bits        int          `json:"bits"`
	Next_block  []string     `json:"next_block"`
	Fee         int          `json:"fee"`
	Nonce       int          `json:"nonce"`
	N_tx        int          `json:"n_tx"`
	Size        int          `json:"size"`
	Block_index int          `json:"block_index"`
	Main_chain  bool         `json:"main_chain"`
	Height      int          `json:"height"`
	Weight      int          `json:"weight"`
	Tx          []TxResponse `json:"tx"`
}

type TxResponse struct {
	Hash         string          `json:"hash"`
	Ver          int             `json:"ver"`
	Vin_sz       int             `json:"vin_sz"`
	Vout_sz      int             `json:"vout_sz"`
	Size         int             `json:"size"`
	Weight       int             `json:"weight"`
	Fee          int             `json:"fee"`
	Relayed_by   string          `json:"relayed_by"`
	Lock_time    int             `json:"lock_time"`
	Tx_index     int             `json:"tx_index"`
	Double_spend bool            `json:"double_spend"`
	Time         int             `json:"time"`
	Block_index  int             `json:"block_index"`
	Block_height int             `json:"block_height"`
	Inputs       []InputResponse `json:"inputs"`
	Out          []OutResponse   `json:"out"`
	Result       int             `json:"result"`
	Balance      int             `json:"balance"`
}

type InputResponse struct {
	Sequence int         `json:"sequence"`
	Witness  string      `json:"witness"`
	Script   string      `json:"script"`
	Index    int         `json:"index"`
	Prev_out OutResponse `json:"prev_out"`
}

type OutResponse struct {
	Type               int                 `json:"type"`
	Spent              bool                `json:"spent"`
	Value              int                 `json:"value"`
	Spending_outpoints []SpendingOutpoints `json:"spending_outpoints"`
	N                  int                 `json:"n"`
	Tx_index           int                 `json:"tx_index"`
	Script             string              `json:"script"`
	Addr               string              `json:"addr"`
}

type SpendingOutpoints struct {
	Tx_index int `json:"tx_index"`
	N        int `json:"n"`
}
