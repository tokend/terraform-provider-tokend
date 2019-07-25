package derive

import "github.com/btcsuite/btcd/chaincfg"

func init() {
	chaincfg.Register(DashMainnetChainParams)
	chaincfg.Register(DashTestnetChainParams)
}

type NetworkType int32

const (
	NetworkTypeBTCMainnet NetworkType = iota + 1
	NetworkTypeBTCTestnet
	NetworkTypeDashMainnet
	NetworkTypeDashTestnet
	NetworkTypeETHMainnet
	NetworkTypeETHTestnet
)

var (
	// https://github.com/dashevo/dashcore-lib/blob/master/docs/networks.md
	DashMainnetChainParams = &chaincfg.Params{
		// Address encoding magics
		PubKeyHashAddrID: 0x4c,
		ScriptHashAddrID: 0x10,
		PrivateKeyID:     0xcc,
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xad, 0xe4},
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xb2, 0x1e},
	}
	// https://github.com/dashevo/dashcore-lib/blob/master/docs/networks.md
	DashTestnetChainParams = &chaincfg.Params{
		PubKeyHashAddrID: 0x8c,
		ScriptHashAddrID: 0x13,
		PrivateKeyID:     0xef,
		HDPrivateKeyID:   [4]byte{0x04, 0x35, 0x83, 0x94},
		HDPublicKeyID:    [4]byte{0x04, 0x35, 0x87, 0xcf},
	}
)

func NetworkParams(network NetworkType) *chaincfg.Params {
	switch network {
	case NetworkTypeBTCMainnet, NetworkTypeETHMainnet:
		return &chaincfg.MainNetParams
	case NetworkTypeBTCTestnet, NetworkTypeETHTestnet:
		return &chaincfg.TestNet3Params
	case NetworkTypeDashMainnet:
		return DashMainnetChainParams
	case NetworkTypeDashTestnet:
		return DashTestnetChainParams
	default:
		panic("unknown network type")
	}
}
