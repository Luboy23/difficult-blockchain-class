package genesis

import "time"


type Genesis struct {
	Date time.Time		`json:"date"`
	ChainID	uint16		`json:"date"`
	TransPerBlock uint16	`json:"trans_per_block"`
	Difficulty uint16		`json:"difficulty"`
	MiningReward	uint64		`json:"mining_reward"`
	GasPrice	uint64			`json:"gas_price"`
	Balances	map[string]uint64		`json:"balances"`
}

