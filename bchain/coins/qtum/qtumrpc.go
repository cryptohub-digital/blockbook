package qtum

import (
	"encoding/json"
	"github.com/core-coin/go-core/xcbclient"
	"github.com/cryptohub-digital/blockbook/contracts"
	"math/big"

	"github.com/cryptohub-digital/blockbook/bchain"
	"github.com/cryptohub-digital/blockbook/bchain/coins/btc"
	"github.com/golang/glog"
)

// QtumRPC is an interface to JSON-RPC bitcoind service.
type QtumRPC struct {
	*btc.BitcoinRPC
	minFeeRate *big.Int // satoshi per kb
}

// NewQtumRPC returns new QtumRPC instance.
func NewQtumRPC(config json.RawMessage, pushHandler func(bchain.NotificationType)) (bchain.BlockChain, error) {
	b, err := btc.NewBitcoinRPC(config, pushHandler)
	if err != nil {
		return nil, err
	}

	s := &QtumRPC{
		b.(*btc.BitcoinRPC),
		big.NewInt(400000),
	}
	s.RPCMarshaler = btc.JSONMarshalerV1{}
	s.ChainConfig.SupportsEstimateSmartFee = true

	return s, nil
}

func (b *QtumRPC) GetRPCClient() *xcbclient.Client {
	return nil
}

func (b *QtumRPC) GetSmartContracts() (*contracts.ChequableToken, *contracts.BountiableToken) {
	return nil, nil
}

// Initialize initializes QtumRPC instance.
func (b *QtumRPC) Initialize() error {
	ci, err := b.GetChainInfo()
	if err != nil {
		return err
	}
	chainName := ci.Chain

	params := GetChainParams(chainName)

	// always create parser
	b.Parser = NewQtumParser(params, b.ChainConfig)

	// parameters for getInfo request
	if params.Net == MainnetMagic {
		b.Testnet = false
		b.Network = "livenet"
	} else {
		b.Testnet = true
		b.Network = "testnet"
	}

	glog.Info("rpc: block chain ", params.Name)

	return nil
}

// GetTransactionForMempool returns a transaction by the transaction ID
// It could be optimized for mempool, i.e. without block time and confirmations
func (b *QtumRPC) GetTransactionForMempool(txid string) (*bchain.Tx, error) {
	return b.GetTransaction(txid)
}

// EstimateSmartFee returns fee estimation
func (b *QtumRPC) EstimateSmartFee(blocks int, conservative bool) (big.Int, error) {
	feeRate, err := b.BitcoinRPC.EstimateSmartFee(blocks, conservative)
	if err != nil {
		if b.minFeeRate.Cmp(&feeRate) == 1 {
			feeRate = *b.minFeeRate
		}
	}
	return feeRate, err
}
