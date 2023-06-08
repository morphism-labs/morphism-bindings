// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morphism-labs/morphism-bindings/solc"
)

const L1CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1002,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1003,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_1_0_1600\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1020_storage\"},{\"astId\":1004,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_51_0_20\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_52_0_1568\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1019_storage\"},{\"astId\":1006,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_101_0_1\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1007,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_102_0_1568\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)1019_storage\"},{\"astId\":1008,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_151_0_32\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_152_0_1568\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)1019_storage\"},{\"astId\":1010,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1011,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1012,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1013,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"receiveNonce\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_uint256\"},{\"astId\":1014,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_address\"},{\"astId\":1015,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_uint240\"},{\"astId\":1016,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1017,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"208\",\"type\":\"t_array(t_uint256)1018_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1018_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[42]\",\"numberOfBytes\":\"1344\"},\"t_array(t_uint256)1019_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1020_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L1CrossDomainMessengerDeployedBin = "0x60806040526004361061015f5760003560e01c80638129fc1c116100c0578063b1b1b20911610074578063d764ad0b11610059578063d764ad0b146103aa578063ecc70428146103bd578063f9a94d7f1461042257600080fd5b8063b1b1b2091461035a578063b28ade251461038a57600080fd5b80638cbeeef2116100a55780638cbeeef2146102575780639fce812c146102e6578063a4e7f8bd1461031a57600080fd5b80638129fc1c146102ba57806383a74074146102cf57600080fd5b80633f827a5a1161011757806354fd4d50116100fc57806354fd4d501461026d5780635644cfdf1461028f5780636e296e45146102a557600080fd5b80633f827a5a1461022f5780634c1d6a691461025757600080fd5b80630ff754ea116101485780630ff754ea146101ac5780632828d7e8146102055780633dbb202b1461021a57600080fd5b8063028f85f7146101645780630c56849814610197575b600080fd5b34801561017057600080fd5b50610179601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b3480156101a357600080fd5b50610179603f81565b3480156101b857600080fd5b506101e07f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161018e565b34801561021157600080fd5b50610179604081565b61022d6102283660046119d7565b610438565b005b34801561023b57600080fd5b50610244600181565b60405161ffff909116815260200161018e565b34801561026357600080fd5b50610179619c4081565b34801561027957600080fd5b5061028261069c565b60405161018e9190611ab8565b34801561029b57600080fd5b5061017961138881565b3480156102b157600080fd5b506101e061073f565b3480156102c657600080fd5b5061022d61082b565b3480156102db57600080fd5b5061017962030d4081565b3480156102f257600080fd5b506101e07f000000000000000000000000000000000000000000000000000000000000000081565b34801561032657600080fd5b5061034a610335366004611ad2565b60cf6020526000908152604090205460ff1681565b604051901515815260200161018e565b34801561036657600080fd5b5061034a610375366004611ad2565b60cb6020526000908152604090205460ff1681565b34801561039657600080fd5b506101796103a5366004611aeb565b610a28565b61022d6103b8366004611b3f565b610a96565b3480156103c957600080fd5b5061041460ce547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b60405190815260200161018e565b34801561042e57600080fd5b5061041460cc5481565b6105717f0000000000000000000000000000000000000000000000000000000000000000610467858585610a28565b347fd764ad0b000000000000000000000000000000000000000000000000000000006104d360ce547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c6040516024016104ef9796959493929190611c0e565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611343565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a3385856105f660ce547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b86604051610608959493929190611c6d565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060ce80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b60606106c77f00000000000000000000000000000000000000000000000000000000000000006113f8565b6106f07f00000000000000000000000000000000000000000000000000000000000000006113f8565b6107197f00000000000000000000000000000000000000000000000000000000000000006113f8565b60405160200161072b93929190611cbb565b604051602081830303815290604052905090565b60cd5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff21530161080e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084015b60405180910390fd5b5060cd5473ffffffffffffffffffffffffffffffffffffffff1690565b6000547501000000000000000000000000000000000000000000900460ff1615808015610876575060005460017401000000000000000000000000000000000000000090910460ff16105b806108a85750303b1580156108a8575060005474010000000000000000000000000000000000000000900460ff166001145b610934576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610805565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000017905580156109ba57600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790555b6109c261152d565b8015610a2557600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6000611388619c4080603f610a44604063ffffffff8816611d60565b610a4e9190611dbf565b610a59601088611d60565b610a669062030d40611de6565b610a709190611de6565b610a7a9190611de6565b610a849190611de6565b610a8e9190611de6565b949350505050565b60f087901c60028110610b51576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a401610805565b8061ffff16600003610c46576000610ba2878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f9250611606915050565b600081815260cb602052604090205490915060ff1615610c44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c617965640000000000000000006064820152608401610805565b505b6000610c8c898989898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061162592505050565b9050610c96611648565b15610cce57853414610caa57610caa611e12565b600081815260cf602052604090205460ff1615610cc957610cc9611e12565b610e20565b3415610d82576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a401610805565b600081815260cf602052604090205460ff16610e20576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c61796564000000000000000000000000000000006064820152608401610805565b610e298761176c565b15610edc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a401610805565b600081815260cb602052604090205460ff1615610f7b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c61796564000000000000000000006064820152608401610805565b610f9c85610f8d611388619c40611de6565b67ffffffffffffffff166117e3565b1580610fc2575060cd5473ffffffffffffffffffffffffffffffffffffffff1661dead14155b156110db57600081815260cf602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555182917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016110d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d657373616765000000000000000000000000000000000000006064820152608401610805565b5050611319565b60cd80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a16179055600061116c88619c405a61112f9190611e41565b8988888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061180192505050565b60cd80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead1790559050801561120357600082815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a2611310565b600082815260cf602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff3201611310576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d657373616765000000000000000000000000000000000000006064820152608401610805565b50505060cc8790555b50505050505050565b905090565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6040517fe9e05c4200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063e9e05c429084906113c0908890839089906000908990600401611e58565b6000604051808303818588803b1580156113d957600080fd5b505af11580156113ed573d6000803e3d6000fd5b505050505050505050565b60608160000361143b57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611465578061144f81611eb0565b915061145e9050600a83611ee8565b915061143f565b60008167ffffffffffffffff81111561148057611480611efc565b6040519080825280601f01601f1916602001820160405280156114aa576020820181803683370190505b5090505b8415610a8e576114bf600183611e41565b91506114cc600a86611f2b565b6114d7906030611f3f565b60f81b8183815181106114ec576114ec611f57565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611526600a86611ee8565b94506114ae565b6000547501000000000000000000000000000000000000000000900460ff166115d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610805565b60cd80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055565b60006116148585858561181b565b805190602001209050949350505050565b60006116358787878787876118b4565b8051906020012090509695505050505050565b60003373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561132257507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639bf62d826040518163ffffffff1660e01b8152600401602060405180830381865afa15801561172c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117509190611f86565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b600073ffffffffffffffffffffffffffffffffffffffff82163014806117dd57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b92915050565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b6060848484846040516024016118349493929190611fa3565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b60608686868686866040516024016118d196959493929190611fed565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610a2557600080fd5b60008083601f84011261198757600080fd5b50813567ffffffffffffffff81111561199f57600080fd5b6020830191508360208285010111156119b757600080fd5b9250929050565b803563ffffffff811681146119d257600080fd5b919050565b600080600080606085870312156119ed57600080fd5b84356119f881611953565b9350602085013567ffffffffffffffff811115611a1457600080fd5b611a2087828801611975565b9094509250611a339050604086016119be565b905092959194509250565b60005b83811015611a59578181015183820152602001611a41565b83811115611a68576000848401525b50505050565b60008151808452611a86816020860160208601611a3e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611acb6020830184611a6e565b9392505050565b600060208284031215611ae457600080fd5b5035919050565b600080600060408486031215611b0057600080fd5b833567ffffffffffffffff811115611b1757600080fd5b611b2386828701611975565b9094509250611b369050602085016119be565b90509250925092565b600080600080600080600060c0888a031215611b5a57600080fd5b873596506020880135611b6c81611953565b95506040880135611b7c81611953565b9450606088013593506080880135925060a088013567ffffffffffffffff811115611ba657600080fd5b611bb28a828b01611975565b989b979a50959850939692959293505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a0830152611c6060c083018486611bc5565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff86168152608060208201526000611c9d608083018688611bc5565b905083604083015263ffffffff831660608301529695505050505050565b60008451611ccd818460208901611a3e565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611d09816001850160208a01611a3e565b60019201918201528351611d24816002840160208801611a3e565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681851681830481118215151615611d8757611d87611d31565b02949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff80841680611dda57611dda611d90565b92169190910492915050565b600067ffffffffffffffff808316818516808303821115611e0957611e09611d31565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600082821015611e5357611e53611d31565b500390565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015267ffffffffffffffff84166040820152821515606082015260a060808201526000611ea560a0830184611a6e565b979650505050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611ee157611ee1611d31565b5060010190565b600082611ef757611ef7611d90565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082611f3a57611f3a611d90565b500690565b60008219821115611f5257611f52611d31565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215611f9857600080fd5b8151611acb81611953565b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525060806040830152611fdc6080830185611a6e565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a083015261203860c0830184611a6e565b9897505050505050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1CrossDomainMessengerStorageLayoutJSON), L1CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1CrossDomainMessenger"] = L1CrossDomainMessengerStorageLayout
	deployedBytecodes["L1CrossDomainMessenger"] = L1CrossDomainMessengerDeployedBin
}
