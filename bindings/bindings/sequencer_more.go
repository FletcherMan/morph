// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const SequencerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/staking/Sequencer.sol:Sequencer\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L2/staking/Sequencer.sol:Sequencer\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L2/staking/Sequencer.sol:Sequencer\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1013_storage\"},{\"astId\":1003,\"contract\":\"contracts/L2/staking/Sequencer.sol:Sequencer\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/L2/staking/Sequencer.sol:Sequencer\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1012_storage\"},{\"astId\":1005,\"contract\":\"contracts/L2/staking/Sequencer.sol:Sequencer\",\"label\":\"sequencerSetVerifyHash\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bytes32\"},{\"astId\":1006,\"contract\":\"contracts/L2/staking/Sequencer.sol:Sequencer\",\"label\":\"blockHeight0\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"contracts/L2/staking/Sequencer.sol:Sequencer\",\"label\":\"sequencerSet0\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_array(t_address)dyn_storage\"},{\"astId\":1008,\"contract\":\"contracts/L2/staking/Sequencer.sol:Sequencer\",\"label\":\"blockHeight1\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"contracts/L2/staking/Sequencer.sol:Sequencer\",\"label\":\"sequencerSet1\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_array(t_address)dyn_storage\"},{\"astId\":1010,\"contract\":\"contracts/L2/staking/Sequencer.sol:Sequencer\",\"label\":\"blockHeight2\",\"offset\":0,\"slot\":\"106\",\"type\":\"t_uint256\"},{\"astId\":1011,\"contract\":\"contracts/L2/staking/Sequencer.sol:Sequencer\",\"label\":\"sequencerSet2\",\"offset\":0,\"slot\":\"107\",\"type\":\"t_array(t_address)dyn_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_address)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"address[]\",\"numberOfBytes\":\"32\"},\"t_array(t_uint256)1012_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1013_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SequencerStorageLayout = new(solc.StorageLayout)

var SequencerDeployedBin = "0x608060405234801561000f575f80fd5b506004361061018f575f3560e01c806377d7dffb116100dd578063a224cee711610088578063b1bdeab311610063578063b1bdeab314610325578063eae5b1e31461032d578063f2fde38b14610335575f80fd5b8063a224cee7146102f7578063a2e53a941461030a578063a384c12e1461031d575f80fd5b806389609d74116100b857806389609d74146102b35780638da5cb5b146102c65780639b8201a4146102e4575f80fd5b806377d7dffb146102715780637d99e8ac14610279578063807de4431461028c575f80fd5b806338871fac1161013d5780636d46e987116101185780636d46e9871461023c5780636d7f64db1461025f578063715018a614610267575f80fd5b806338871fac14610216578063480265c91461021e57806365fd0f8d14610233575f80fd5b806328d1357a1161016d57806328d1357a146101cc57806329025fcb146101d55780632aae60bd146101de575f80fd5b80630d78725b146101935780630e06ede8146101af57806317f24c2d146101b7575b5f80fd5b61019c60655481565b6040519081526020015b60405180910390f35b606b5461019c565b6101bf610348565b6040516101a691906110b2565b61019c606a5481565b61019c60685481565b6101f16101ec36600461110b565b610499565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a6565b60675461019c565b6102266104ce565b6040516101a69190611122565b61019c60665481565b61024f61024a3660046111b4565b610508565b60405190151581526020016101a6565b6101bf61057e565b61026f6105e9565b005b6101bf6105fc565b6101f161028736600461110b565b610667565b6101f17f000000000000000000000000000000000000000000000000000000000000000081565b6101f16102c136600461110b565b610676565b60335473ffffffffffffffffffffffffffffffffffffffff166101f1565b61026f6102f23660046111d4565b610685565b61026f6103053660046111d4565b61083f565b61024f6103183660046111b4565b610a96565b61019c610bf0565b60695461019c565b6101bf610c18565b61026f6103433660046111b4565b610c83565b6060606a5443106103be57606b8054806020026020016040519081016040528092919081815260200182805480156103b457602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610389575b5050505050905090565b60685443106104305760698054806020026020016040519081016040528092919081815260200182805480156103b457602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610389575050505050905090565b60678054806020026020016040519081016040528092919081815260200182805480156103b457602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610389575050505050905090565b606781815481106104a8575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b606060665460676068546069606a54606b6040516020016104f496959493929190611296565b604051602081830303815290604052905090565b5f610578606b80548060200260200160405190810160405280929190818152602001828054801561056d57602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610542575b505050505083610d3a565b92915050565b606060678054806020026020016040519081016040528092919081815260200182805480156103b457602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610389575050505050905090565b6105f1610dad565b6105fa5f610e2e565b565b6060606b8054806020026020016040519081016040528092919081815260200182805480156103b457602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610389575050505050905090565b606b81815481106104a8575f80fd5b606981815481106104a8575f80fd5b337f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1614610729576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79204c325374616b696e6720636f6e747261637400000000000000000060448201526064015b60405180910390fd5b606a546107374360026112ed565b111561078b5760688054606655606a5490556107544360026112ed565b606a556069805461076791606791610fdc565b50606b805461077891606991610fdc565b50610785606b8383611028565b50610799565b610797606b8383611028565b505b60665460676068546069606a54606b6040516020016107bd96959493929190611296565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101206065557f7083eed0a633eebfb4ad275c34bdd165d2df2c83d7415e880220b116fb46bc6282826108244360026112ed565b6040516108339392919061136c565b60405180910390a15050565b5f54610100900460ff161580801561085d57505f54600160ff909116105b806108765750303b15801561087657505f5460ff166001145b610902576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610720565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561095e575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b816109c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f696e76616c69642073657175656e6365722073657400000000000000000000006044820152606401610720565b6109cd610ea4565b6109d960678484611028565b506109e660698484611028565b506109f3606b8484611028565b507f7083eed0a633eebfb4ad275c34bdd165d2df2c83d7415e880220b116fb46bc6283835f604051610a279392919061136c565b60405180910390a18015610a91575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b5f606a544310610b0d57610578606b80548060200260200160405190810160405280929190818152602001828054801561056d57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161054257505050505083610d3a565b6068544310610b8357610578606980548060200260200160405190810160405280929190818152602001828054801561056d57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161054257505050505083610d3a565b610578606780548060200260200160405190810160405280929190818152602001828054801561056d57602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161054257505050505083610d3a565b5f606a544310610c015750606b5490565b6068544310610c11575060695490565b5060675490565b606060698054806020026020016040519081016040528092919081815260200182805480156103b457602002820191905f5260205f2090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610389575050505050905090565b610c8b610dad565b73ffffffffffffffffffffffffffffffffffffffff8116610d2e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610720565b610d3781610e2e565b50565b5f805b8351811015610da457838181518110610d5857610d5861138f565b602002602001015173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610d9c576001915050610578565b600101610d3d565b505f9392505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146105fa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610720565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16610f3a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610720565b6105fa5f54610100900460ff16610fd3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610720565b6105fa33610e2e565b828054828255905f5260205f20908101928215611018575f5260205f209182015b82811115611018578254825591600101919060010190610ffd565b5061102492915061109e565b5090565b828054828255905f5260205f20908101928215611018579160200282015b828111156110185781547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff843516178255602090920191600190910190611046565b5b80821115611024575f815560010161109f565b602080825282518282018190525f9190848201906040850190845b818110156110ff57835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016110cd565b50909695505050505050565b5f6020828403121561111b575f80fd5b5035919050565b5f602080835283518060208501525f5b8181101561114e57858101830151858201604001528201611132565b505f6040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146111af575f80fd5b919050565b5f602082840312156111c4575f80fd5b6111cd8261118c565b9392505050565b5f80602083850312156111e5575f80fd5b823567ffffffffffffffff808211156111fc575f80fd5b818501915085601f83011261120f575f80fd5b81358181111561121d575f80fd5b8660208260051b8501011115611231575f80fd5b60209290920196919550909350505050565b5f815480845260208085019450835f5260205f205f5b8381101561128b57815473ffffffffffffffffffffffffffffffffffffffff1687529582019560019182019101611259565b509495945050505050565b86815260c060208201525f6112ae60c0830188611243565b86604084015282810360608401526112c68187611243565b905084608084015282810360a08401526112e08185611243565b9998505050505050505050565b80820180821115610578577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8183525f60208085019450825f5b8581101561128b5773ffffffffffffffffffffffffffffffffffffffff6113598361118c565b1687529582019590820190600101611333565b604081525f61137f604083018587611325565b9050826020830152949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(SequencerStorageLayoutJSON), SequencerStorageLayout); err != nil {
		panic(err)
	}

	layouts["Sequencer"] = SequencerStorageLayout
	deployedBytecodes["Sequencer"] = SequencerDeployedBin
}