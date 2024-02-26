package derivation

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/morph-l2/bindings/bindings"
	"github.com/morph-l2/node/types"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind/backends"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/core"
	"github.com/scroll-tech/go-ethereum/core/rawdb"
	"github.com/scroll-tech/go-ethereum/eth"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/ethclient/authclient"
	"github.com/scroll-tech/go-ethereum/ethdb"
	"github.com/scroll-tech/go-ethereum/rpc"
	"github.com/stretchr/testify/require"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

func TestCompareBlock(t *testing.T) {
	eClient, err := ethclient.Dial("http://localhost:7545")
	require.NoError(t, err)
	l2Client, err := ethclient.Dial("http://localhost:8545")
	blockNumber, err := eClient.BlockNumber(context.Background())
	require.NoError(t, err)
	for i := 0; i < int(blockNumber); i++ {
		block, err := l2Client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		require.NoError(t, err)
		dBlock, err := eClient.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		require.True(t, reflect.DeepEqual(block.Header(), dBlock.Header()))
	}
}

func testNewDerivationClient(t *testing.T) *Derivation {
	ctx := context.Background()
	l1Client, err := ethclient.Dial("http://localhost:9545")
	addr := common.HexToAddress("0x6900000000000000000000000000000000000010")
	require.NoError(t, err)
	var secret [32]byte
	jwtSecret := common.FromHex(strings.TrimSpace("688f5d737bad920bdfb2fc2f488d6b6209eebda1dae949a8de91398d932c517a"))
	require.True(t, len(jwtSecret) == 32)
	copy(secret[:], jwtSecret)
	aClient, err := authclient.DialContext(context.Background(), "http://localhost:7551", secret)
	require.NoError(t, err)
	eClient, err := ethclient.Dial("http://localhost:7545")
	require.NoError(t, err)
	d := Derivation{
		ctx:                   ctx,
		l1Client:              l1Client,
		RollupContractAddress: addr,
		confirmations:         rpc.BlockNumber(5),
		l2Client:              types.NewRetryableClient(aClient, eClient, tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout))),
		validator:             nil,
		latestDerivation:      9,
		fetchBlockRange:       100,
		pollInterval:          1,
	}
	return &d
}

func TestFetchRollupData(t *testing.T) {
	d := testNewDerivationClient(t)
	_, err := d.fetchRollupDataByTxHash(common.HexToHash("0x27deefd6883567dec94fa3eca918de6a4a6d330e6f225b6282a5ab7e3c1a50dc"), 390)
	require.NoError(t, err)
}

func newSimulatedBackend(key *ecdsa.PrivateKey) (*backends.SimulatedBackend, ethdb.Database) {
	var gasLimit uint64 = 9_000_000
	auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	genAlloc := make(core.GenesisAlloc)
	genAlloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(9223372036854775807)}
	db := rawdb.NewMemoryDatabase()
	sim := backends.NewSimulatedBackendWithDatabase(db, genAlloc, gasLimit)
	return sim, db
}

func TestDecodeBatch(t *testing.T) {
	abi, err := bindings.RollupMetaData.GetAbi()
	require.NoError(t, err)
	hexData := "0x16b799c9000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000002201d4ceff4b2335970615354f0a03e2124745d1c193d509e9109e4910b869448ce0f199dc7fb956c94a69ba951785dae12d9d6c7ed3073dd5a5453151d9996a25127ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d75700000000000000000000000000000000000000000000000000000000000002600000000000000000000000000000000000000000000000000000000000000059000000000000000000000000000000000000000000000000008cb3decea512e30f962b50492fee707a926cd3465d2ac9b2b9655553a915578d00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000003d010000000000000001000000006570805b0000000000000000000000000000000000000000000000000000000000000000000000000098968000010001000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000005bc55b7eb2a19d020541768ca2c88e9feed6df90edc34e7aac67e1c2ae0a2ae72e2b381671f72bf2067d74ab0c11a91000000000000000000000000000000000fbdc060a4fd3fb5923c5ff9a430acdba1fbc17136e76cb85af00b274dfeaed8ce982afd68ffea340dd71c1391e31ad3"
	hexData = "0x16b799c900000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000007e00598741c60fade057ffcf8325b9293d8c4cb3050100ffa3eae839ddcf9c543fc03d2784f29d69c292e52e737cd3e2da355adb93988694ad3a3e506a44a88993727ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d7570000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000007900000000000000000100000000000000010000000000000001e3bf30a9d601ef9e5180c2e6baf65b3bb603de775f58a1792d79d2e4e0daf30aa191d6f404af264276151c6c5e4eca85d5545cd56327d76230185fcd2bf8efe60000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000005d1140000000000000002000000006579793d00000000000000000000000000000000000000000000000000000000000000000000000000989680000000000000000000000003000000006579793f0000000000000000000000000000000000000000000000000000000000000000000000000098968000000000000000000000000400000000657979400000000000000000000000000000000000000000000000000000000000000000000000000098968000010000000000000000000500000000657979420000000000000000000000000000000000000000000000000000000000000000000000000098968000000000000000000000000600000000657979440000000000000000000000000000000000000000000000000000000000000000000000000098968000000000000000000000000700000000657979450000000000000000000000000000000000000000000000000000000000000000000000000098968000010000000000000000000800000000657979470000000000000000000000000000000000000000000000000000000000000000000000000098968000000000000000000000000900000000657979480000000000000000000000000000000000000000000000000000000000000000000000000098968000000000000000000000000a00000000657979490000000000000000000000000000000000000000000000000000000000000000000000000098968000000000000000000000000b000000006579794a0000000000000000000000000000000000000000000000000000000000000000000000000098968000000000000000000000000c000000006579794c0000000000000000000000000000000000000000000000000000000000000000000000000098968000000000000000000000000d000000006579794d0000000000000000000000000000000000000000000000000000000000000000000000000098968000000000000000000000000e000000006579794e0000000000000000000000000000000000000000000000000000000000000000000000000098968000000000000000000000000f000000006579795000000000000000000000000000000000000000000000000000000000000000000000000000989680000000000000000000000010000000006579795100000000000000000000000000000000000000000000000000000000000000000000000000989680000000000000000000000011000000006579795200000000000000000000000000000000000000000000000000000000000000000000000000989680000000000000000000000012000000006579795400000000000000000000000000000000000000000000000000000000000000000000000000989680000000000000000000000013000000006579795500000000000000000000000000000000000000000000000000000000000000000000000000989680000000000000000000000014000000006579795600000000000000000000000000000000000000000000000000000000000000000000000000989680000000000000000000000015000000006579795800000000000000000000000000000000000000000000000000000000000000000000000000989680000000000000008cf88a808405f5e10082881694530000000000000000000000000000000000000f80a4bede39b5000000000000000000000000000000000000000000000000000000000000000783019ecda019ed2e6515399be155b690b6d99e98c95ce5b5848ff68b47b7f385f11aada9cda072168cc84fce7af487029f739ec046fa65de2e5ad8164ec7e75e7f6e7f01fe530000008cf88a018405f5e10082885494530000000000000000000000000000000000000f80a47046559700000000000000000000000000000000000000000000000001fb87d0c13b5c5083019ecea0897b453b86033c4623dc8644ddc4630fc7425fb9252f62312005f40c034a6edda0055cb46f77326f19e9e100558d5d0c985349e27d950a222efdb66f2ff6f7d8750000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000a4f1b116b8464509f38377dc36f9617cda99c87a81473f62c4a847629a6a3a547ea2361065002a46bbf658fe70f78f30000000000000000000000000000000016e1d66dc0d2bf949049c481128980f9fb26177a6396447704080cb97e2270565b2f64a32633558510b43067324e44dd"
	txData, err := hexutil.Decode(hexData)
	require.NoError(t, err)
	require.NoError(t, err)
	args, err := abi.Methods["commitBatch"].Inputs.Unpack(txData[4:])
	rollupBatchData := args[0].(struct {
		Version                uint8     "json:\"version\""
		ParentBatchHeader      []uint8   "json:\"parentBatchHeader\""
		Chunks                 [][]uint8 "json:\"chunks\""
		SkippedL1MessageBitmap []uint8   "json:\"skippedL1MessageBitmap\""
		PrevStateRoot          [32]uint8 "json:\"prevStateRoot\""
		PostStateRoot          [32]uint8 "json:\"postStateRoot\""
		WithdrawalRoot         [32]uint8 "json:\"withdrawalRoot\""
		Signature              struct {
			Version   *big.Int   "json:\"version\""
			Signers   []*big.Int "json:\"signers\""
			Signature []uint8    "json:\"signature\""
		} "json:\"signature\""
	})

	var chunks []hexutil.Bytes
	for _, chunk := range rollupBatchData.Chunks {
		chunks = append(chunks, chunk)
	}
	batch := eth.RPCRollupBatch{
		Version:                uint(rollupBatchData.Version),
		ParentBatchHeader:      rollupBatchData.ParentBatchHeader,
		Chunks:                 chunks,
		SkippedL1MessageBitmap: rollupBatchData.SkippedL1MessageBitmap,
		PrevStateRoot:          common.BytesToHash(rollupBatchData.PrevStateRoot[:]),
		PostStateRoot:          common.BytesToHash(rollupBatchData.PostStateRoot[:]),
		WithdrawRoot:           common.BytesToHash(rollupBatchData.WithdrawalRoot[:]),
	}
	_, err = ParseBatch(batch)
	require.NoError(t, err)
}

func DialEthClientWithTimeout(ctx context.Context, url string) (*ethclient.Client, *rpc.Client, error) {
	ctxt, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	c, err := rpc.DialContext(ctxt, url)
	if err != nil {
		return nil, nil, err
	}
	return ethclient.NewClient(c), c, nil
}

func TestCompareBlockHash(t *testing.T) {
	sClient, err := ethclient.Dial("http://localhost:8545")
	require.NoError(t, err)
	dClient, err := ethclient.Dial("http://localhost:7545")
	require.NoError(t, err)
	latest, err := dClient.BlockNumber(context.Background())
	for i := 0; i <= int(latest); i++ {
		sBlock, err := sClient.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		require.NoError(t, err)
		dBlock, err := dClient.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		require.NoError(t, err)
		require.Equal(t, sBlock.Hash(), dBlock.Hash())
	}
}
