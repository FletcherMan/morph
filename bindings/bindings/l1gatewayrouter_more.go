// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L1GatewayRouterStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/gateways/L1GatewayRouter.sol:L1GatewayRouter\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/gateways/L1GatewayRouter.sol:L1GatewayRouter\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/gateways/L1GatewayRouter.sol:L1GatewayRouter\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1010_storage\"},{\"astId\":1003,\"contract\":\"contracts/L1/gateways/L1GatewayRouter.sol:L1GatewayRouter\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/L1/gateways/L1GatewayRouter.sol:L1GatewayRouter\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1009_storage\"},{\"astId\":1005,\"contract\":\"contracts/L1/gateways/L1GatewayRouter.sol:L1GatewayRouter\",\"label\":\"ethGateway\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_address\"},{\"astId\":1006,\"contract\":\"contracts/L1/gateways/L1GatewayRouter.sol:L1GatewayRouter\",\"label\":\"defaultERC20Gateway\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_address\"},{\"astId\":1007,\"contract\":\"contracts/L1/gateways/L1GatewayRouter.sol:L1GatewayRouter\",\"label\":\"ERC20Gateway\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_mapping(t_address,t_address)\"},{\"astId\":1008,\"contract\":\"contracts/L1/gateways/L1GatewayRouter.sol:L1GatewayRouter\",\"label\":\"gatewayInContext\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1009_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1010_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_address)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e address)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_address\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1GatewayRouterStorageLayout = new(solc.StorageLayout)

var L1GatewayRouterDeployedBin = "0x60806040526004361061016d575f3560e01c80638c00ce73116100c6578063c52a3bbc1161007c578063ce8c3e0611610057578063ce8c3e06146103db578063f219fa6614610407578063f2fde38b1461041a575f80fd5b8063c52a3bbc1461037c578063c676ad29146103a9578063ce0b63ce146103c8575f80fd5b80638eaac8a3116100ac5780638eaac8a3146103485780639f8420b314610356578063aac476f814610369575f80fd5b80638c00ce73146102f25780638da5cb5b1461031e575f80fd5b8063485cc95511610126578063705b05b811610101578063705b05b81461028a578063715018a6146102cb57806384bd13b0146102df575f80fd5b8063485cc9551461022d5780635dfd5b9a1461024c578063635c86371461026b575f80fd5b80633a9a7b20116101565780633a9a7b20146101995780633d1d31c7146101ef57806343c667411461020e575f80fd5b80630aea8c261461017157806321425ee014610186575b5f80fd5b61018461017f366004611a17565b610439565b005b610184610194366004611a86565b610684565b3480156101a4575f80fd5b506068546101c59073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101fa575f80fd5b50610184610209366004611ab8565b6106c2565b348015610219575f80fd5b506101c5610228366004611ab8565b610740565b348015610238575f80fd5b50610184610247366004611ad3565b61078f565b348015610257575f80fd5b50610184610266366004611ab8565b610a30565b348015610276575f80fd5b50610184610285366004611b85565b610aae565b348015610295575f80fd5b506101c56102a4366004611ab8565b60676020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b3480156102d6575f80fd5b50610184610cfe565b6101846102ed366004611c2a565b610d11565b3480156102fd575f80fd5b506065546101c59073ffffffffffffffffffffffffffffffffffffffff1681565b348015610329575f80fd5b5060335473ffffffffffffffffffffffffffffffffffffffff166101c5565b6101846102ed366004611cbc565b610184610364366004611d2a565b610d73565b610184610377366004611d4a565b610daf565b348015610387575f80fd5b5061039b610396366004611da7565b610f81565b6040519081526020016101e6565b3480156103b4575f80fd5b506101c56103c3366004611ab8565b611178565b6101846103d6366004611a86565b61123d565b3480156103e6575f80fd5b506066546101c59073ffffffffffffffffffffffffffffffffffffffff1681565b610184610415366004611de5565b611248565b348015610425575f80fd5b50610184610434366004611ab8565b61125a565b60685473ffffffffffffffffffffffffffffffffffffffff16156104be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f4f6e6c79206e6f7420696e20636f6e746578740000000000000000000000000060448201526064015b60405180910390fd5b5f6104c886610740565b905073ffffffffffffffffffffffffffffffffffffffff8116610547576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f206761746577617920617661696c61626c6500000000000000000000000060448201526064016104b5565b606880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555f338460405160200161059b929190611e93565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f0aea8c26000000000000000000000000000000000000000000000000000000008252915073ffffffffffffffffffffffffffffffffffffffff831690630aea8c26903490610625908b908b908b9088908b90600401611ec1565b5f604051808303818588803b15801561063c575f80fd5b505af115801561064e573d5f803e3d5ffd5b5050606880547fffffffffffffffffffffffff000000000000000000000000000000000000000016905550505050505050505050565b6106bd8333845f5b6040519080825280601f01601f1916602001820160405280156106b6576020820181803683370190505b5085610439565b505050565b6106ca611311565b6065805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907fa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a905f90a35050565b73ffffffffffffffffffffffffffffffffffffffff8082165f9081526067602052604081205490911680610789575060665473ffffffffffffffffffffffffffffffffffffffff165b92915050565b5f54610100900460ff16158080156107ad57505f54600160ff909116105b806107c65750303b1580156107c657505f5460ff166001145b610852576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104b5565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156108ae575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6108b6611392565b73ffffffffffffffffffffffffffffffffffffffff82161561094057606680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091556040515f907f2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1908290a35b73ffffffffffffffffffffffffffffffffffffffff8316156109ca57606580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091556040515f907fa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a908290a35b80156106bd575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b610a38611311565b6066805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1905f90a35050565b610ab6611311565b8051825114610b21576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f6c656e677468206d69736d61746368000000000000000000000000000000000060448201526064016104b5565b5f5b82518110156106bd575f60675f858481518110610b4257610b42611f11565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050828281518110610bb757610bb7611f11565b602002602001015160675f868581518110610bd457610bd4611f11565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828281518110610c6457610c64611f11565b602002602001015173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16858481518110610cab57610cab611f11565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167f0ead4808404683f66d413d788a768219ea9785c97889221193103841a5841eaf60405160405180910390a450600101610b23565b610d06611311565b610d0f5f611430565b565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f73686f756c64206e657665722062652063616c6c65640000000000000000000060448201526064016104b5565b610dab33835f5b6040519080825280601f01601f191660200182016040528015610da4576020820181803683370190505b5084610daf565b5050565b60685473ffffffffffffffffffffffffffffffffffffffff1615610e2f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f4f6e6c79206e6f7420696e20636f6e746578740000000000000000000000000060448201526064016104b5565b60655473ffffffffffffffffffffffffffffffffffffffff1680610eaf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f657468206761746577617920617661696c61626c65000000000000000000000060448201526064016104b5565b5f3384604051602001610ec3929190611e93565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527faac476f8000000000000000000000000000000000000000000000000000000008252915073ffffffffffffffffffffffffffffffffffffffff83169063aac476f8903490610f4b908a908a9087908a90600401611f3e565b5f604051808303818588803b158015610f62575f80fd5b505af1158015610f74573d5f803e3d5ffd5b5050505050505050505050565b6068545f9073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461101a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f4f6e6c7920696e206465706f73697420636f6e7465787400000000000000000060448201526064016104b5565b5f336040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301529192505f918616906370a0823190602401602060405180830381865afa15801561108b573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110af9190611f83565b90506110d373ffffffffffffffffffffffffffffffffffffffff86168784876114a6565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301528291908716906370a0823190602401602060405180830381865afa158015611140573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111649190611f83565b61116e9190611f9a565b9695505050505050565b5f8061118383610740565b905073ffffffffffffffffffffffffffffffffffffffff81166111a857505f92915050565b6040517fc676ad2900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015282169063c676ad2990602401602060405180830381865afa158015611212573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112369190611fd2565b9392505050565b6106bd83835f610d7a565b6112548484845f61068c565b50505050565b611262611311565b73ffffffffffffffffffffffffffffffffffffffff8116611305576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104b5565b61130e81611430565b50565b60335473ffffffffffffffffffffffffffffffffffffffff163314610d0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104b5565b5f54610100900460ff16611428576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104b5565b610d0f61153b565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526112549085906115da565b5f54610100900460ff166115d1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104b5565b610d0f33611430565b5f61163b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166116e79092919063ffffffff16565b905080515f148061165b57508080602001905181019061165b9190611fed565b6106bd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016104b5565b60606116f584845f856116fd565b949350505050565b60608247101561178f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016104b5565b5f808673ffffffffffffffffffffffffffffffffffffffff1685876040516117b7919061200c565b5f6040518083038185875af1925050503d805f81146117f1576040519150601f19603f3d011682016040523d82523d5f602084013e6117f6565b606091505b509150915061180787838387611812565b979650505050505050565b606083156118a75782515f036118a05773ffffffffffffffffffffffffffffffffffffffff85163b6118a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016104b5565b50816116f5565b6116f583838151156118bc5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b59190612027565b73ffffffffffffffffffffffffffffffffffffffff8116811461130e575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561198557611985611911565b604052919050565b5f82601f83011261199c575f80fd5b813567ffffffffffffffff8111156119b6576119b6611911565b6119e760207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161193e565b8181528460208386010111156119fb575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f60a08688031215611a2b575f80fd5b8535611a36816118f0565b94506020860135611a46816118f0565b935060408601359250606086013567ffffffffffffffff811115611a68575f80fd5b611a748882890161198d565b95989497509295608001359392505050565b5f805f60608486031215611a98575f80fd5b8335611aa3816118f0565b95602085013595506040909401359392505050565b5f60208284031215611ac8575f80fd5b8135611236816118f0565b5f8060408385031215611ae4575f80fd5b8235611aef816118f0565b91506020830135611aff816118f0565b809150509250929050565b5f82601f830112611b19575f80fd5b8135602067ffffffffffffffff821115611b3557611b35611911565b8160051b611b4482820161193e565b9283528481018201928281019087851115611b5d575f80fd5b83870192505b84831015611807578235611b76816118f0565b82529183019190830190611b63565b5f8060408385031215611b96575f80fd5b823567ffffffffffffffff80821115611bad575f80fd5b611bb986838701611b0a565b93506020850135915080821115611bce575f80fd5b50611bdb85828601611b0a565b9150509250929050565b5f8083601f840112611bf5575f80fd5b50813567ffffffffffffffff811115611c0c575f80fd5b602083019150836020828501011115611c23575f80fd5b9250929050565b5f805f805f805f60c0888a031215611c40575f80fd5b8735611c4b816118f0565b96506020880135611c5b816118f0565b95506040880135611c6b816118f0565b94506060880135611c7b816118f0565b93506080880135925060a088013567ffffffffffffffff811115611c9d575f80fd5b611ca98a828b01611be5565b989b979a50959850939692959293505050565b5f805f805f60808688031215611cd0575f80fd5b8535611cdb816118f0565b94506020860135611ceb816118f0565b935060408601359250606086013567ffffffffffffffff811115611d0d575f80fd5b611d1988828901611be5565b969995985093965092949392505050565b5f8060408385031215611d3b575f80fd5b50508035926020909101359150565b5f805f8060808587031215611d5d575f80fd5b8435611d68816118f0565b935060208501359250604085013567ffffffffffffffff811115611d8a575f80fd5b611d968782880161198d565b949793965093946060013593505050565b5f805f60608486031215611db9575f80fd5b8335611dc4816118f0565b92506020840135611dd4816118f0565b929592945050506040919091013590565b5f805f8060808587031215611df8575f80fd5b8435611e03816118f0565b93506020850135611e13816118f0565b93969395505050506040820135916060013590565b5f5b83811015611e42578181015183820152602001611e2a565b50505f910152565b5f8151808452611e61816020860160208601611e28565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f6116f56040830184611e4a565b5f73ffffffffffffffffffffffffffffffffffffffff808816835280871660208401525084604083015260a06060830152611eff60a0830185611e4a565b90508260808301529695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152608060408201525f611f726080830185611e4a565b905082606083015295945050505050565b5f60208284031215611f93575f80fd5b5051919050565b81810381811115610789577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f60208284031215611fe2575f80fd5b8151611236816118f0565b5f60208284031215611ffd575f80fd5b81518015158114611236575f80fd5b5f825161201d818460208701611e28565b9190910192915050565b602081525f6112366020830184611e4a56fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(L1GatewayRouterStorageLayoutJSON), L1GatewayRouterStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1GatewayRouter"] = L1GatewayRouterStorageLayout
	deployedBytecodes["L1GatewayRouter"] = L1GatewayRouterDeployedBin
}