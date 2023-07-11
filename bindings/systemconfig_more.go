// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morphism-labs/morphism-bindings/solc"
)

const SystemConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1011_storage\"},{\"astId\":1003,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1010_storage\"},{\"astId\":1005,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_bytes32\"},{\"astId\":1008,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"gasLimit\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint64\"},{\"astId\":1009,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_resourceConfig\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_struct(ResourceConfig)1012_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1010_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1011_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_struct(ResourceConfig)1012_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ResourceMetering.ResourceConfig\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SystemConfigStorageLayout = new(solc.StorageLayout)

var SystemConfigDeployedBin = "0x608060405234801561001057600080fd5b50600436106101515760003560e01c8063b40a817c116100cd578063f2fde38b11610081578063f68016b711610066578063f68016b7146103f7578063f975e9251461040b578063ffa1ad741461041e57600080fd5b8063f2fde38b146103db578063f45e65d8146103ee57600080fd5b8063c9b26f61116100b2578063c9b26f611461028b578063cc731b021461029e578063e81b2c6d146103d257600080fd5b8063b40a817c14610265578063c71973f61461027857600080fd5b80634f16540b11610124578063715018a611610109578063715018a61461022c5780638da5cb5b14610234578063935f029e1461025257600080fd5b80634f16540b146101f057806354fd4d501461021757600080fd5b80630c18c1621461015657806318d13918146101725780631fd19ee1146101875780634add321d146101cf575b600080fd5b61015f60655481565b6040519081526020015b60405180910390f35b610185610180366004611370565b610426565b005b7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c08545b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610169565b6101d76104ea565b60405167ffffffffffffffff9091168152602001610169565b61015f7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0881565b61021f610515565b604051610169919061140c565b6101856105b8565b60335473ffffffffffffffffffffffffffffffffffffffff166101aa565b61018561026036600461141f565b6105cc565b610185610273366004611459565b610665565b6101856102863660046115b1565b610750565b6101856102993660046115cd565b610764565b6103626040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260695463ffffffff8082168352640100000000820460ff9081166020850152650100000000008304169383019390935266010000000000008104831660608301526a0100000000000000000000810490921660808201526e0100000000000000000000000000009091046fffffffffffffffffffffffffffffffff1660a082015290565b6040516101699190600060c08201905063ffffffff80845116835260ff602085015116602084015260ff6040850151166040840152806060850151166060840152806080850151166080840152506fffffffffffffffffffffffffffffffff60a08401511660a083015292915050565b61015f60675481565b6101856103e9366004611370565b610794565b61015f60665481565b6068546101d79067ffffffffffffffff1681565b6101856104193660046115e6565b610848565b61015f600081565b61042e610afb565b610456817f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0855565b6040805173ffffffffffffffffffffffffffffffffffffffff8316602082015260009101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052905060035b60007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be836040516104de919061140c565b60405180910390a35050565b6069546000906105109063ffffffff6a0100000000000000000000820481169116611688565b905090565b60606105407f0000000000000000000000000000000000000000000000000000000000000000610b7c565b6105697f0000000000000000000000000000000000000000000000000000000000000000610b7c565b6105927f0000000000000000000000000000000000000000000000000000000000000000610b7c565b6040516020016105a4939291906116b4565b604051602081830303815290604052905090565b6105c0610afb565b6105ca6000610c3a565b565b6105d4610afb565b606582905560668190556040805160208101849052908101829052600090606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529050600160007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be83604051610658919061140c565b60405180910390a3505050565b61066d610afb565b6106756104ea565b67ffffffffffffffff168167ffffffffffffffff1610156106f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064015b60405180910390fd5b606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff831690811790915560408051602080820193909352815180820390930183528101905260026104ad565b610758610afb565b61076181610cb1565b50565b61076c610afb565b60678190556040805160208082018490528251808303909101815290820190915260006104ad565b61079c610afb565b73ffffffffffffffffffffffffffffffffffffffff811661083f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016106ee565b61076181610c3a565b600054610100900460ff16158080156108685750600054600160ff909116105b806108825750303b158015610882575060005460ff166001145b61090e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016106ee565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561096c57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610974611125565b61097d88610794565b606587905560668690556067859055606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff86161790557f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c088390556109ed82610cb1565b6109f56104ea565b67ffffffffffffffff168467ffffffffffffffff161015610a72576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016106ee565b8015610ad557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b60335473ffffffffffffffffffffffffffffffffffffffff1633146105ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106ee565b60606000610b89836111c4565b600101905060008167ffffffffffffffff811115610ba957610ba9611474565b6040519080825280601f01601f191660200182016040528015610bd3576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084610bdd57509392505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b8060a001516fffffffffffffffffffffffffffffffff16816060015163ffffffff161115610d61576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f53797374656d436f6e6669673a206d696e206261736520666565206d7573742060448201527f6265206c657373207468616e206d61782062617365000000000000000000000060648201526084016106ee565b6001816040015160ff1611610df8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a2064656e6f6d696e61746f72206d757374206260448201527f65206c6172676572207468616e2031000000000000000000000000000000000060648201526084016106ee565b6068546080820151825167ffffffffffffffff90921691610e19919061172a565b63ffffffff161115610e87576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016106ee565b6000816020015160ff1611610f1e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a20656c6173746963697479206d756c7469706c60448201527f6965722063616e6e6f742062652030000000000000000000000000000000000060648201526084016106ee565b8051602082015163ffffffff82169160ff90911690610f3e908290611749565b610f489190611793565b63ffffffff1614610fdb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f53797374656d436f6e6669673a20707265636973696f6e206c6f73732077697460448201527f6820746172676574207265736f75726365206c696d697400000000000000000060648201526084016106ee565b805160698054602084015160408501516060860151608087015160a09097015163ffffffff9687167fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009095169490941764010000000060ff94851602177fffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffffff166501000000000093909216929092027fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1617660100000000000091851691909102177fffff0000000000000000000000000000000000000000ffffffffffffffffffff166a010000000000000000000093909416929092027fffff00000000000000000000000000000000ffffffffffffffffffffffffffff16929092176e0100000000000000000000000000006fffffffffffffffffffffffffffffffff90921691909102179055565b600054610100900460ff166111bc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106ee565b6105ca6112a7565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061120d577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310611239576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061125757662386f26fc10000830492506010015b6305f5e100831061126f576305f5e100830492506008015b612710831061128357612710830492506004015b60648310611295576064830492506002015b600a83106112a1576001015b92915050565b600054610100900460ff1661133e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106ee565b6105ca33610c3a565b803573ffffffffffffffffffffffffffffffffffffffff8116811461136b57600080fd5b919050565b60006020828403121561138257600080fd5b61138b82611347565b9392505050565b60005b838110156113ad578181015183820152602001611395565b838111156113bc576000848401525b50505050565b600081518084526113da816020860160208601611392565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061138b60208301846113c2565b6000806040838503121561143257600080fd5b50508035926020909101359150565b803567ffffffffffffffff8116811461136b57600080fd5b60006020828403121561146b57600080fd5b61138b82611441565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b803563ffffffff8116811461136b57600080fd5b803560ff8116811461136b57600080fd5b80356fffffffffffffffffffffffffffffffff8116811461136b57600080fd5b600060c082840312156114fa57600080fd5b60405160c0810181811067ffffffffffffffff82111715611544577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604052905080611553836114a3565b8152611561602084016114b7565b6020820152611572604084016114b7565b6040820152611583606084016114a3565b6060820152611594608084016114a3565b60808201526115a560a084016114c8565b60a08201525092915050565b600060c082840312156115c357600080fd5b61138b83836114e8565b6000602082840312156115df57600080fd5b5035919050565b6000806000806000806000610180888a03121561160257600080fd5b61160b88611347565b965060208801359550604088013594506060880135935061162e60808901611441565b925061163c60a08901611347565b915061164b8960c08a016114e8565b905092959891949750929550565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff8083168185168083038211156116ab576116ab611659565b01949350505050565b600084516116c6818460208901611392565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611702816001850160208a01611392565b6001920191820152835161171d816002840160208801611392565b0160020195945050505050565b600063ffffffff8083168185168083038211156116ab576116ab611659565b600063ffffffff80841680611787577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600063ffffffff808316818516818304811182151516156117b6576117b6611659565b0294935050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(SystemConfigStorageLayoutJSON), SystemConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SystemConfig"] = SystemConfigStorageLayout
	deployedBytecodes["SystemConfig"] = SystemConfigDeployedBin
}
