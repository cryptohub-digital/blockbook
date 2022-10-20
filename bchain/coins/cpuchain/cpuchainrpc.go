package cpuchain

import (
	"encoding/json"
	"github.com/core-coin/go-core/xcbclient"
	"github.com/cryptohub-digital/blockbook/contracts"

	"github.com/cryptohub-digital/blockbook/bchain"
	"github.com/cryptohub-digital/blockbook/bchain/coins/btc"
	"github.com/golang/glog"
)

// CPUchainRPC is an interface to JSON-RPC bitcoind service.
type CPUchainRPC struct {
	*btc.BitcoinRPC
}

// NewCPUchainRPC returns new CPUchainRPC instance.
func NewCPUchainRPC(config json.RawMessage, pushHandler func(bchain.NotificationType)) (bchain.BlockChain, error) {
	b, err := btc.NewBitcoinRPC(config, pushHandler)
	if err != nil {
		return nil, err
	}

	s := &CPUchainRPC{
		b.(*btc.BitcoinRPC),
	}
	s.RPCMarshaler = btc.JSONMarshalerV2{}
	s.ChainConfig.SupportsEstimateFee = false

	return s, nil
}

func (b *CPUchainRPC) GetRPCClient() *xcbclient.Client {
	return nil
}

func (b *CPUchainRPC) GetSmartContracts() (*contracts.ChequableToken, *contracts.BountiableToken) {
	return nil, nil
}

// Initialize initializes CPUchainRPC instance.
func (b *CPUchainRPC) Initialize() error {
	ci, err := b.GetChainInfo()
	if err != nil {
		return err
	}
	chainName := ci.Chain

	glog.Info("Chain name ", chainName)
	params := GetChainParams(chainName)

	// always create parser
	b.Parser = NewCPUchainParser(params, b.ChainConfig)

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
