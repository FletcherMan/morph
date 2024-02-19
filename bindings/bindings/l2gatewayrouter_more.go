// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L2GatewayRouterStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/gateways/L2GatewayRouter.sol:L2GatewayRouter\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L2/gateways/L2GatewayRouter.sol:L2GatewayRouter\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L2/gateways/L2GatewayRouter.sol:L2GatewayRouter\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1009_storage\"},{\"astId\":1003,\"contract\":\"contracts/L2/gateways/L2GatewayRouter.sol:L2GatewayRouter\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/L2/gateways/L2GatewayRouter.sol:L2GatewayRouter\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1008_storage\"},{\"astId\":1005,\"contract\":\"contracts/L2/gateways/L2GatewayRouter.sol:L2GatewayRouter\",\"label\":\"ethGateway\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_address\"},{\"astId\":1006,\"contract\":\"contracts/L2/gateways/L2GatewayRouter.sol:L2GatewayRouter\",\"label\":\"defaultERC20Gateway\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_address\"},{\"astId\":1007,\"contract\":\"contracts/L2/gateways/L2GatewayRouter.sol:L2GatewayRouter\",\"label\":\"ERC20Gateway\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_mapping(t_address,t_address)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1008_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1009_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_address)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e address)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_address\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2GatewayRouterStorageLayout = new(solc.StorageLayout)

var L2GatewayRouterDeployedBin = "0x608060405260043610610162575f3560e01c80636dc24183116100c65780638da5cb5b1161007c578063c7cdea3711610057578063c7cdea3714610395578063ce8c3e06146103a8578063f2fde38b146103d4575f80fd5b80638da5cb5b14610339578063a93a4af914610363578063c676ad2914610376575f80fd5b8063715018a6116100ac578063715018a6146102eb5780638431f5c1146102ff5780638c00ce731461030d575f80fd5b80636dc2418314610297578063705b05b8146102aa575f80fd5b806354bbd59c1161011b5780635dfd5b9a116101015780635dfd5b9a14610246578063635c8637146102655780636c07ea4314610284575f80fd5b806354bbd59c14610214578063575361b614610233575f80fd5b80633d1d31c71161014b5780633d1d31c71461018e57806343c66741146101ad578063485cc955146101f5575f80fd5b8063232e8748146101665780632fcc29fa1461017b575b5f80fd5b61017961017436600461126a565b6103f3565b005b6101796101893660046112d8565b61045a565b348015610199575f80fd5b506101796101a836600461130a565b610497565b3480156101b8575f80fd5b506101cc6101c736600461130a565b610515565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b348015610200575f80fd5b5061017961020f366004611325565b610564565b34801561021f575f80fd5b506101cc61022e36600461130a565b610805565b610179610241366004611462565b6108ca565b348015610251575f80fd5b5061017961026036600461130a565b610a28565b348015610270575f80fd5b5061017961027f366004611557565b610aa6565b6101796102923660046112d8565b610cf6565b6101796102a53660046115b7565b610d2f565b3480156102b5575f80fd5b506101cc6102c436600461130a565b60676020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b3480156102f6575f80fd5b50610179610e81565b610179610174366004611614565b348015610318575f80fd5b506065546101cc9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610344575f80fd5b5060335473ffffffffffffffffffffffffffffffffffffffff166101cc565b6101796103713660046116a6565b610e94565b348015610381575f80fd5b506101cc61039036600461130a565b610ea6565b6101796103a33660046116e9565b610f0a565b3480156103b3575f80fd5b506066546101cc9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156103df575f80fd5b506101796103ee36600461130a565b610f19565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f73686f756c64206e657665722062652063616c6c65640000000000000000000060448201526064015b60405180910390fd5b61049283835f5b6040519080825280601f01601f19166020018201604052801561048b576020820181803683370190505b5084610d2f565b505050565b61049f610fd0565b6065805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907fa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a905f90a35050565b73ffffffffffffffffffffffffffffffffffffffff8082165f908152606760205260408120549091168061055e575060665473ffffffffffffffffffffffffffffffffffffffff165b92915050565b5f54610100900460ff161580801561058257505f54600160ff909116105b8061059b5750303b15801561059b57505f5460ff166001145b610627576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610451565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610683575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61068b611051565b73ffffffffffffffffffffffffffffffffffffffff82161561071557606680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091556040515f907f2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1908290a35b73ffffffffffffffffffffffffffffffffffffffff83161561079f57606580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091556040515f907fa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a908290a35b8015610492575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b5f8061081083610515565b905073ffffffffffffffffffffffffffffffffffffffff811661083557505f92915050565b6040517f54bbd59c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301528216906354bbd59c90602401602060405180830381865afa15801561089f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108c39190611709565b9392505050565b5f6108d486610515565b905073ffffffffffffffffffffffffffffffffffffffff8116610953576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f206761746577617920617661696c61626c650000000000000000000000006044820152606401610451565b5f3384604051602001610967929190611785565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f575361b6000000000000000000000000000000000000000000000000000000008252915073ffffffffffffffffffffffffffffffffffffffff83169063575361b69034906109f1908b908b908b9088908b906004016117bb565b5f604051808303818588803b158015610a08575f80fd5b505af1158015610a1a573d5f803e3d5ffd5b505050505050505050505050565b610a30610fd0565b6066805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1905f90a35050565b610aae610fd0565b8051825114610b19576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f6c656e677468206d69736d6174636800000000000000000000000000000000006044820152606401610451565b5f5b8251811015610492575f60675f858481518110610b3a57610b3a61180b565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050828281518110610baf57610baf61180b565b602002602001015160675f868581518110610bcc57610bcc61180b565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550828281518110610c5c57610c5c61180b565b602002602001015173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16858481518110610ca357610ca361180b565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167f0ead4808404683f66d413d788a768219ea9785c97889221193103841a5841eaf60405160405180910390a450600101610b1b565b6104928333845f5b6040519080825280601f01601f191660200182016040528015610d28576020820181803683370190505b50856108ca565b60655473ffffffffffffffffffffffffffffffffffffffff1680610daf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f657468206761746577617920617661696c61626c6500000000000000000000006044820152606401610451565b5f3384604051602001610dc3929190611785565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f6dc24183000000000000000000000000000000000000000000000000000000008252915073ffffffffffffffffffffffffffffffffffffffff831690636dc24183903490610e4b908a908a9087908a90600401611838565b5f604051808303818588803b158015610e62575f80fd5b505af1158015610e74573d5f803e3d5ffd5b5050505050505050505050565b610e89610fd0565b610e925f6110ef565b565b610ea08484845f610cfe565b50505050565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f756e737570706f7274656400000000000000000000000000000000000000000060448201525f90606401610451565b610f1533835f610461565b5050565b610f21610fd0565b73ffffffffffffffffffffffffffffffffffffffff8116610fc4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610451565b610fcd816110ef565b50565b60335473ffffffffffffffffffffffffffffffffffffffff163314610e92576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610451565b5f54610100900460ff166110e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610451565b610e92611165565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff166111fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610451565b610e92336110ef565b73ffffffffffffffffffffffffffffffffffffffff81168114610fcd575f80fd5b5f8083601f840112611235575f80fd5b50813567ffffffffffffffff81111561124c575f80fd5b602083019150836020828501011115611263575f80fd5b9250929050565b5f805f805f6080868803121561127e575f80fd5b853561128981611204565b9450602086013561129981611204565b935060408601359250606086013567ffffffffffffffff8111156112bb575f80fd5b6112c788828901611225565b969995985093965092949392505050565b5f805f606084860312156112ea575f80fd5b83356112f581611204565b95602085013595506040909401359392505050565b5f6020828403121561131a575f80fd5b81356108c381611204565b5f8060408385031215611336575f80fd5b823561134181611204565b9150602083013561135181611204565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156113d0576113d061135c565b604052919050565b5f82601f8301126113e7575f80fd5b813567ffffffffffffffff8111156114015761140161135c565b61143260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611389565b818152846020838601011115611446575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f60a08688031215611476575f80fd5b853561148181611204565b9450602086013561149181611204565b935060408601359250606086013567ffffffffffffffff8111156114b3575f80fd5b6114bf888289016113d8565b95989497509295608001359392505050565b5f82601f8301126114e0575f80fd5b8135602067ffffffffffffffff8211156114fc576114fc61135c565b8160051b61150b828201611389565b9283528481018201928281019087851115611524575f80fd5b83870192505b8483101561154c57823561153d81611204565b8252918301919083019061152a565b979650505050505050565b5f8060408385031215611568575f80fd5b823567ffffffffffffffff8082111561157f575f80fd5b61158b868387016114d1565b935060208501359150808211156115a0575f80fd5b506115ad858286016114d1565b9150509250929050565b5f805f80608085870312156115ca575f80fd5b84356115d581611204565b935060208501359250604085013567ffffffffffffffff8111156115f7575f80fd5b611603878288016113d8565b949793965093946060013593505050565b5f805f805f805f60c0888a03121561162a575f80fd5b873561163581611204565b9650602088013561164581611204565b9550604088013561165581611204565b9450606088013561166581611204565b93506080880135925060a088013567ffffffffffffffff811115611687575f80fd5b6116938a828b01611225565b989b979a50959850939692959293505050565b5f805f80608085870312156116b9575f80fd5b84356116c481611204565b935060208501356116d481611204565b93969395505050506040820135916060013590565b5f80604083850312156116fa575f80fd5b50508035926020909101359150565b5f60208284031215611719575f80fd5b81516108c381611204565b5f81518084525f5b818110156117485760208185018101518683018201520161172c565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f6117b36040830184611724565b949350505050565b5f73ffffffffffffffffffffffffffffffffffffffff808816835280871660208401525084604083015260a060608301526117f960a0830185611724565b90508260808301529695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152608060408201525f61186c6080830185611724565b90508260608301529594505050505056fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(L2GatewayRouterStorageLayoutJSON), L2GatewayRouterStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2GatewayRouter"] = L2GatewayRouterStorageLayout
	deployedBytecodes["L2GatewayRouter"] = L2GatewayRouterDeployedBin
}