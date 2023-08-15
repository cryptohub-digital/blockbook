package dbtestdata

import (
	"encoding/hex"

	"github.com/cryptohub-digital/blockbook/bchain"
)

const (
	XcbTx1Packed         = "08e8dd870210a6a6f0db051a6d08ece40212050430e234001888a40122081bc0159d530e60003220998d535fb50fc55eafc591c20acf9ae13cebb96676fe90fcd136ea1f941135203a16cb79fbc0290a1a3cf017f702e604ba234568533110af4216cb656dadee521bea601692312454a655a0f49051ddc9480a22070a025208120101"
	XcbTx1FailedPacked   = "08e8dd870210a6a6f0db051a6d08ece40212050430e234001888a40122081bc0159d530e60003220998d535fb50fc55eafc591c20acf9ae13cebb96676fe90fcd136ea1f941135203a16cb79fbc0290a1a3cf017f702e604ba234568533110af4216cb656dadee521bea601692312454a655a0f49051ddc9480a22040a025208"
	XcbTx1NoStatusPacked = "08e8dd870210a6a6f0db051a6d08ece40212050430e234001888a40122081bc0159d530e60003220998d535fb50fc55eafc591c20acf9ae13cebb96676fe90fcd136ea1f941135203a16cb79fbc0290a1a3cf017f702e604ba234568533110af4216cb656dadee521bea601692312454a655a0f49051ddc9480a22070a025208120155"
	XcbTx2Packed         = "08a9f2a20210a6a6f0db051aa60208a00712043b9aca00189aa8022ac401e86e7c5f00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000100000000000000000000ab27b691efe91718cb73207207d92dbd175e6b10c7560000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000002b5e3af16b188000032206fc698f1f6037551826fd86fa1b77c27a16c62f8916f9fe9942cd89b2fc8118a3a16ab98e5e2ba00469ce51440c22d4d4b79a56da712297f4216ab44de35413ee2b672d322938e2fcc932d5c0cf8ec8822aa010a02941a1201011aa0010a16ab98e5e2ba00469ce51440c22d4d4b79a56da712297f1220000000000000000000000000000000000000000000000002b5e3af16b18800001a20c17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a11a2000000000000000000000ab44de35413ee2b672d322938e2fcc932d5c0cf8ec881a2000000000000000000000ab27b691efe91718cb73207207d92dbd175e6b10c756"
	XcbTx3Packed         = "08c782a30210a6a6f0db051aa401080312043b9aca00188099022a444b40e90100000000000000000000ab094a15c3dc43095c7450c59bf56263e9827065f3060000000000000000000000000000000000000000000000000de0b6b3a764000032204f65e846f570bb121b959bd37fbe57f4a6a61598095cbc4c6eaaa66aed7f66bd3a16ab98e5e2ba00469ce51440c22d4d4b79a56da712297f4216ab228a4d4263e067df56b1dd226acb939f532ff7ab5b22aa010a028c801201011aa0010a16ab98e5e2ba00469ce51440c22d4d4b79a56da712297f12200000000000000000000000000000000000000000000000000de0b6b3a76400001a20c17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a11a2000000000000000000000ab228a4d4263e067df56b1dd226acb939f532ff7ab5b1a2000000000000000000000ab094a15c3dc43095c7450c59bf56263e9827065f306"
)

// GetTestCoreCoinTypeBlock1 returns block #1
func GetTestCoreCoinTypeBlock1(parser bchain.BlockChainParser) *bchain.Block {
	return &bchain.Block{
		BlockHeader: bchain.BlockHeader{
			Height:        4767873,
			Hash:          "0x35553ec376267a47c7c2657d5df1d5bb788fc6ca2396dd0070ca61ef45d4edec",
			Size:          2362,
			Time:          1534858022,
			Confirmations: 179,
		},
		Txs: unpackCoreCoinTxs([]string{XcbTx2Packed, XcbTx3Packed}, parser),
	}
}

func unpackCoreCoinTxs(packed []string, parser bchain.BlockChainParser) []bchain.Tx {
	r := make([]bchain.Tx, len(packed))
	for i, p := range packed {
		b, err := hex.DecodeString(p)
		if err != nil {
			panic(err)
		}
		tx, _, err := parser.UnpackTx(b)
		if err != nil {
			panic(err)
		}
		r[i] = *tx
	}
	return r
}
