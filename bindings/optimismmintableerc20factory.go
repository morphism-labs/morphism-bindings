// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OptimismMintableERC20FactoryMetaData contains all meta data concerning the OptimismMintableERC20Factory contract.
var OptimismMintableERC20FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"}],\"name\":\"OptimismMintableERC20Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"StandardL2TokenCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"createOptimismMintableERC20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"createStandardL2Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61010060405234801561001157600080fd5b5060405161225938038061225983398101604081905261003091610050565b6001608081905260a052600060c0526001600160a01b031660e052610080565b60006020828403121561006257600080fd5b81516001600160a01b038116811461007957600080fd5b9392505050565b60805160a05160c05160e05161219b6100be6000396000818160d3015261026501526000610153015260006101280152600060fd015261219b6000f3fe60806040523480156200001157600080fd5b5060043610620000525760003560e01c806354fd4d501462000057578063896f93d11462000079578063ce5ac90f14620000b6578063ee9a31a214620000cd575b600080fd5b62000061620000f5565b604051620000709190620005b2565b60405180910390f35b620000906200008a366004620006b0565b620001a0565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200162000070565b62000090620000c7366004620006b0565b620001b7565b620000907f000000000000000000000000000000000000000000000000000000000000000081565b6060620001227f000000000000000000000000000000000000000000000000000000000000000062000376565b6200014d7f000000000000000000000000000000000000000000000000000000000000000062000376565b620001787f000000000000000000000000000000000000000000000000000000000000000062000376565b6040516020016200018c9392919062000747565b604051602081830303815290604052905090565b6000620001af848484620001b7565b949350505050565b600073ffffffffffffffffffffffffffffffffffffffff841662000261576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4f7074696d69736d4d696e7461626c654552433230466163746f72793a206d7560448201527f73742070726f766964652072656d6f746520746f6b656e206164647265737300606482015260840160405180910390fd5b60007f0000000000000000000000000000000000000000000000000000000000000000858585604051620002959062000525565b620002a49493929190620007c3565b604051809103906000f080158015620002c1573d6000803e3d6000fd5b5090508073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf60405160405180910390a360405133815273ffffffffffffffffffffffffffffffffffffffff80871691908316907f52fe89dd5930f343d25650b62fd367bae47088bcddffd2a88350a6ecdd620cdb9060200160405180910390a3949350505050565b6060600062000385836200043b565b600101905060008167ffffffffffffffff811115620003a857620003a8620005ce565b6040519080825280601f01601f191660200182016040528015620003d3576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084620003dd57509392505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831062000485577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310620004b2576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310620004d157662386f26fc10000830492506010015b6305f5e1008310620004ea576305f5e100830492506008015b6127108310620004ff57612710830492506004015b6064831062000512576064830492506002015b600a83106200051f576001015b92915050565b611971806200081e83390190565b60005b838110156200055057818101518382015260200162000536565b8381111562000560576000848401525b50505050565b600081518084526200058081602086016020860162000533565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000620005c7602083018462000566565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126200060f57600080fd5b813567ffffffffffffffff808211156200062d576200062d620005ce565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715620006765762000676620005ce565b816040528381528660208588010111156200069057600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600060608486031215620006c657600080fd5b833573ffffffffffffffffffffffffffffffffffffffff81168114620006eb57600080fd5b9250602084013567ffffffffffffffff808211156200070957600080fd5b6200071787838801620005fd565b935060408601359150808211156200072e57600080fd5b506200073d86828701620005fd565b9150509250925092565b600084516200075b81846020890162000533565b80830190507f2e00000000000000000000000000000000000000000000000000000000000000808252855162000799816001850160208a0162000533565b60019201918201528351620007b681600284016020880162000533565b0160020195945050505050565b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525060806040830152620007fe608083018562000566565b828103606084015262000812818562000566565b97965050505050505056fe6101206040523480156200001257600080fd5b50604051620019713803806200197183398101604081905262000035916200016d565b6001600080848460036200004a83826200028c565b5060046200005982826200028c565b50505060809290925260a05260c05250506001600160a01b0390811660e052166101005262000358565b80516001600160a01b03811681146200009b57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000c857600080fd5b81516001600160401b0380821115620000e557620000e5620000a0565b604051601f8301601f19908116603f01168101908282118183101715620001105762000110620000a0565b816040528381526020925086838588010111156200012d57600080fd5b600091505b8382101562000151578582018301518183018401529082019062000132565b83821115620001635760008385830101525b9695505050505050565b600080600080608085870312156200018457600080fd5b6200018f8562000083565b93506200019f6020860162000083565b60408601519093506001600160401b0380821115620001bd57600080fd5b620001cb88838901620000b6565b93506060870151915080821115620001e257600080fd5b50620001f187828801620000b6565b91505092959194509250565b600181811c908216806200021257607f821691505b6020821081036200023357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200028757600081815260208120601f850160051c81016020861015620002625750805b601f850160051c820191505b8181101562000283578281556001016200026e565b5050505b505050565b81516001600160401b03811115620002a857620002a8620000a0565b620002c081620002b98454620001fd565b8462000239565b602080601f831160018114620002f85760008415620002df5750858301515b600019600386901b1c1916600185901b17855562000283565b600085815260208120601f198616915b82811015620003295788860151825594840194600190910190840162000308565b5085821015620003485787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c05160e051610100516115b8620003b9600039600081816102f50152818161038a015281816105cf01526107a90152600081816101a9015261031b015260006107380152600061070f015260006106e601526115b86000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063ae1f6aaf1161008c578063dd62ed3e11610066578063dd62ed3e1461033f578063e78cea92146102f3578063ee9a31a21461038557600080fd5b8063ae1f6aaf146102f3578063c01e1bd614610319578063d6c0b2c41461031957600080fd5b80639dc29fac116100bd5780639dc29fac146102ba578063a457c2d7146102cd578063a9059cbb146102e057600080fd5b806370a082311461027c57806395d89b41146102b257600080fd5b806323b872dd1161012f5780633950935111610114578063395093511461024c57806340c10f191461025f57806354fd4d501461027457600080fd5b806323b872dd1461022a578063313ce5671461023d57600080fd5b806306fdde031161016057806306fdde03146101f0578063095ea7b31461020557806318160ddd1461021857600080fd5b806301ffc9a71461017c578063033964be146101a4575b600080fd5b61018f61018a3660046112d1565b6103ac565b60405190151581526020015b60405180910390f35b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f861049d565b60405161019b9190611346565b61018f6102133660046113c0565b61052f565b6002545b60405190815260200161019b565b61018f6102383660046113ea565b610547565b6040516012815260200161019b565b61018f61025a3660046113c0565b61056b565b61027261026d3660046113c0565b6105b7565b005b6101f86106df565b61021c61028a366004611426565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101f8610782565b6102726102c83660046113c0565b610791565b61018f6102db3660046113c0565b6108a8565b61018f6102ee3660046113c0565b610979565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b61021c61034d366004611441565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f1d1d8b63000000000000000000000000000000000000000000000000000000007fec4fc8e3000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000851683148061046557507fffffffff00000000000000000000000000000000000000000000000000000000858116908316145b8061049457507fffffffff00000000000000000000000000000000000000000000000000000000858116908216145b95945050505050565b6060600380546104ac90611474565b80601f01602080910402602001604051908101604052809291908181526020018280546104d890611474565b80156105255780601f106104fa57610100808354040283529160200191610525565b820191906000526020600020905b81548152906001019060200180831161050857829003601f168201915b5050505050905090565b60003361053d818585610987565b5060019392505050565b600033610555858285610b3b565b610560858585610c12565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061053d90829086906105b29087906114c7565b610987565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610681576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460448201527f67652063616e206d696e7420616e64206275726e00000000000000000000000060648201526084015b60405180910390fd5b61068b8282610e81565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885826040516106d391815260200190565b60405180910390a25050565b606061070a7f0000000000000000000000000000000000000000000000000000000000000000610f74565b6107337f0000000000000000000000000000000000000000000000000000000000000000610f74565b61075c7f0000000000000000000000000000000000000000000000000000000000000000610f74565b60405160200161076e93929190611506565b604051602081830303815290604052905090565b6060600480546104ac90611474565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610856576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460448201527f67652063616e206d696e7420616e64206275726e0000000000000000000000006064820152608401610678565b6108608282611032565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5826040516106d391815260200190565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091908381101561096c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610678565b6105608286868403610987565b60003361053d818585610c12565b73ffffffffffffffffffffffffffffffffffffffff8316610a29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff8216610acc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610c0c5781811015610bff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610678565b610c0c8484848403610987565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610cb5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff8216610d58576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610e0e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610c0c565b73ffffffffffffffffffffffffffffffffffffffff8216610efe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610678565b8060026000828254610f1091906114c7565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b60606000610f81836111ee565b600101905060008167ffffffffffffffff811115610fa157610fa161157c565b6040519080825280601f01601f191660200182016040528015610fcb576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084610fd557509392505050565b73ffffffffffffffffffffffffffffffffffffffff82166110d5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff82166000908152602081905260409020548181101561118b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610b2e565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310611237577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310611263576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061128157662386f26fc10000830492506010015b6305f5e1008310611299576305f5e100830492506008015b61271083106112ad57612710830492506004015b606483106112bf576064830492506002015b600a83106112cb576001015b92915050565b6000602082840312156112e357600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461131357600080fd5b9392505050565b60005b8381101561133557818101518382015260200161131d565b83811115610c0c5750506000910152565b602081526000825180602084015261136581604085016020870161131a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146113bb57600080fd5b919050565b600080604083850312156113d357600080fd5b6113dc83611397565b946020939093013593505050565b6000806000606084860312156113ff57600080fd5b61140884611397565b925061141660208501611397565b9150604084013590509250925092565b60006020828403121561143857600080fd5b61131382611397565b6000806040838503121561145457600080fd5b61145d83611397565b915061146b60208401611397565b90509250929050565b600181811c9082168061148857607f821691505b6020821081036114c1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60008219821115611501577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500190565b6000845161151881846020890161131a565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611554816001850160208a0161131a565b6001920191820152835161156f81600284016020880161131a565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c634300080f000aa164736f6c634300080f000a",
}

// OptimismMintableERC20FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use OptimismMintableERC20FactoryMetaData.ABI instead.
var OptimismMintableERC20FactoryABI = OptimismMintableERC20FactoryMetaData.ABI

// OptimismMintableERC20FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OptimismMintableERC20FactoryMetaData.Bin instead.
var OptimismMintableERC20FactoryBin = OptimismMintableERC20FactoryMetaData.Bin

// DeployOptimismMintableERC20Factory deploys a new Ethereum contract, binding an instance of OptimismMintableERC20Factory to it.
func DeployOptimismMintableERC20Factory(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *OptimismMintableERC20Factory, error) {
	parsed, err := OptimismMintableERC20FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptimismMintableERC20FactoryBin), backend, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptimismMintableERC20Factory{OptimismMintableERC20FactoryCaller: OptimismMintableERC20FactoryCaller{contract: contract}, OptimismMintableERC20FactoryTransactor: OptimismMintableERC20FactoryTransactor{contract: contract}, OptimismMintableERC20FactoryFilterer: OptimismMintableERC20FactoryFilterer{contract: contract}}, nil
}

// OptimismMintableERC20Factory is an auto generated Go binding around an Ethereum contract.
type OptimismMintableERC20Factory struct {
	OptimismMintableERC20FactoryCaller     // Read-only binding to the contract
	OptimismMintableERC20FactoryTransactor // Write-only binding to the contract
	OptimismMintableERC20FactoryFilterer   // Log filterer for contract events
}

// OptimismMintableERC20FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptimismMintableERC20FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismMintableERC20FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptimismMintableERC20FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismMintableERC20FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptimismMintableERC20FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismMintableERC20FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptimismMintableERC20FactorySession struct {
	Contract     *OptimismMintableERC20Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OptimismMintableERC20FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptimismMintableERC20FactoryCallerSession struct {
	Contract *OptimismMintableERC20FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// OptimismMintableERC20FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptimismMintableERC20FactoryTransactorSession struct {
	Contract     *OptimismMintableERC20FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// OptimismMintableERC20FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptimismMintableERC20FactoryRaw struct {
	Contract *OptimismMintableERC20Factory // Generic contract binding to access the raw methods on
}

// OptimismMintableERC20FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptimismMintableERC20FactoryCallerRaw struct {
	Contract *OptimismMintableERC20FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// OptimismMintableERC20FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptimismMintableERC20FactoryTransactorRaw struct {
	Contract *OptimismMintableERC20FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptimismMintableERC20Factory creates a new instance of OptimismMintableERC20Factory, bound to a specific deployed contract.
func NewOptimismMintableERC20Factory(address common.Address, backend bind.ContractBackend) (*OptimismMintableERC20Factory, error) {
	contract, err := bindOptimismMintableERC20Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptimismMintableERC20Factory{OptimismMintableERC20FactoryCaller: OptimismMintableERC20FactoryCaller{contract: contract}, OptimismMintableERC20FactoryTransactor: OptimismMintableERC20FactoryTransactor{contract: contract}, OptimismMintableERC20FactoryFilterer: OptimismMintableERC20FactoryFilterer{contract: contract}}, nil
}

// NewOptimismMintableERC20FactoryCaller creates a new read-only instance of OptimismMintableERC20Factory, bound to a specific deployed contract.
func NewOptimismMintableERC20FactoryCaller(address common.Address, caller bind.ContractCaller) (*OptimismMintableERC20FactoryCaller, error) {
	contract, err := bindOptimismMintableERC20Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismMintableERC20FactoryCaller{contract: contract}, nil
}

// NewOptimismMintableERC20FactoryTransactor creates a new write-only instance of OptimismMintableERC20Factory, bound to a specific deployed contract.
func NewOptimismMintableERC20FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*OptimismMintableERC20FactoryTransactor, error) {
	contract, err := bindOptimismMintableERC20Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismMintableERC20FactoryTransactor{contract: contract}, nil
}

// NewOptimismMintableERC20FactoryFilterer creates a new log filterer instance of OptimismMintableERC20Factory, bound to a specific deployed contract.
func NewOptimismMintableERC20FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*OptimismMintableERC20FactoryFilterer, error) {
	contract, err := bindOptimismMintableERC20Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimismMintableERC20FactoryFilterer{contract: contract}, nil
}

// bindOptimismMintableERC20Factory binds a generic wrapper to an already deployed contract.
func bindOptimismMintableERC20Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OptimismMintableERC20FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismMintableERC20Factory.Contract.OptimismMintableERC20FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.OptimismMintableERC20FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.OptimismMintableERC20FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismMintableERC20Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.contract.Transact(opts, method, params...)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryCaller) BRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismMintableERC20Factory.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactorySession) BRIDGE() (common.Address, error) {
	return _OptimismMintableERC20Factory.Contract.BRIDGE(&_OptimismMintableERC20Factory.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryCallerSession) BRIDGE() (common.Address, error) {
	return _OptimismMintableERC20Factory.Contract.BRIDGE(&_OptimismMintableERC20Factory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OptimismMintableERC20Factory.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactorySession) Version() (string, error) {
	return _OptimismMintableERC20Factory.Contract.Version(&_OptimismMintableERC20Factory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryCallerSession) Version() (string, error) {
	return _OptimismMintableERC20Factory.Contract.Version(&_OptimismMintableERC20Factory.CallOpts)
}

// CreateOptimismMintableERC20 is a paid mutator transaction binding the contract method 0xce5ac90f.
//
// Solidity: function createOptimismMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryTransactor) CreateOptimismMintableERC20(opts *bind.TransactOpts, _remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.contract.Transact(opts, "createOptimismMintableERC20", _remoteToken, _name, _symbol)
}

// CreateOptimismMintableERC20 is a paid mutator transaction binding the contract method 0xce5ac90f.
//
// Solidity: function createOptimismMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactorySession) CreateOptimismMintableERC20(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.CreateOptimismMintableERC20(&_OptimismMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// CreateOptimismMintableERC20 is a paid mutator transaction binding the contract method 0xce5ac90f.
//
// Solidity: function createOptimismMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryTransactorSession) CreateOptimismMintableERC20(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.CreateOptimismMintableERC20(&_OptimismMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// CreateStandardL2Token is a paid mutator transaction binding the contract method 0x896f93d1.
//
// Solidity: function createStandardL2Token(address _remoteToken, string _name, string _symbol) returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryTransactor) CreateStandardL2Token(opts *bind.TransactOpts, _remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.contract.Transact(opts, "createStandardL2Token", _remoteToken, _name, _symbol)
}

// CreateStandardL2Token is a paid mutator transaction binding the contract method 0x896f93d1.
//
// Solidity: function createStandardL2Token(address _remoteToken, string _name, string _symbol) returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactorySession) CreateStandardL2Token(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.CreateStandardL2Token(&_OptimismMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// CreateStandardL2Token is a paid mutator transaction binding the contract method 0x896f93d1.
//
// Solidity: function createStandardL2Token(address _remoteToken, string _name, string _symbol) returns(address)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryTransactorSession) CreateStandardL2Token(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _OptimismMintableERC20Factory.Contract.CreateStandardL2Token(&_OptimismMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// OptimismMintableERC20FactoryOptimismMintableERC20CreatedIterator is returned from FilterOptimismMintableERC20Created and is used to iterate over the raw logs and unpacked data for OptimismMintableERC20Created events raised by the OptimismMintableERC20Factory contract.
type OptimismMintableERC20FactoryOptimismMintableERC20CreatedIterator struct {
	Event *OptimismMintableERC20FactoryOptimismMintableERC20Created // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OptimismMintableERC20FactoryOptimismMintableERC20CreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismMintableERC20FactoryOptimismMintableERC20Created)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OptimismMintableERC20FactoryOptimismMintableERC20Created)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OptimismMintableERC20FactoryOptimismMintableERC20CreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimismMintableERC20FactoryOptimismMintableERC20CreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimismMintableERC20FactoryOptimismMintableERC20Created represents a OptimismMintableERC20Created event raised by the OptimismMintableERC20Factory contract.
type OptimismMintableERC20FactoryOptimismMintableERC20Created struct {
	LocalToken  common.Address
	RemoteToken common.Address
	Deployer    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOptimismMintableERC20Created is a free log retrieval operation binding the contract event 0x52fe89dd5930f343d25650b62fd367bae47088bcddffd2a88350a6ecdd620cdb.
//
// Solidity: event OptimismMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryFilterer) FilterOptimismMintableERC20Created(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address) (*OptimismMintableERC20FactoryOptimismMintableERC20CreatedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _OptimismMintableERC20Factory.contract.FilterLogs(opts, "OptimismMintableERC20Created", localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &OptimismMintableERC20FactoryOptimismMintableERC20CreatedIterator{contract: _OptimismMintableERC20Factory.contract, event: "OptimismMintableERC20Created", logs: logs, sub: sub}, nil
}

// WatchOptimismMintableERC20Created is a free log subscription operation binding the contract event 0x52fe89dd5930f343d25650b62fd367bae47088bcddffd2a88350a6ecdd620cdb.
//
// Solidity: event OptimismMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryFilterer) WatchOptimismMintableERC20Created(opts *bind.WatchOpts, sink chan<- *OptimismMintableERC20FactoryOptimismMintableERC20Created, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _OptimismMintableERC20Factory.contract.WatchLogs(opts, "OptimismMintableERC20Created", localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimismMintableERC20FactoryOptimismMintableERC20Created)
				if err := _OptimismMintableERC20Factory.contract.UnpackLog(event, "OptimismMintableERC20Created", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOptimismMintableERC20Created is a log parse operation binding the contract event 0x52fe89dd5930f343d25650b62fd367bae47088bcddffd2a88350a6ecdd620cdb.
//
// Solidity: event OptimismMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryFilterer) ParseOptimismMintableERC20Created(log types.Log) (*OptimismMintableERC20FactoryOptimismMintableERC20Created, error) {
	event := new(OptimismMintableERC20FactoryOptimismMintableERC20Created)
	if err := _OptimismMintableERC20Factory.contract.UnpackLog(event, "OptimismMintableERC20Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OptimismMintableERC20FactoryStandardL2TokenCreatedIterator is returned from FilterStandardL2TokenCreated and is used to iterate over the raw logs and unpacked data for StandardL2TokenCreated events raised by the OptimismMintableERC20Factory contract.
type OptimismMintableERC20FactoryStandardL2TokenCreatedIterator struct {
	Event *OptimismMintableERC20FactoryStandardL2TokenCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OptimismMintableERC20FactoryStandardL2TokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismMintableERC20FactoryStandardL2TokenCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OptimismMintableERC20FactoryStandardL2TokenCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OptimismMintableERC20FactoryStandardL2TokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimismMintableERC20FactoryStandardL2TokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimismMintableERC20FactoryStandardL2TokenCreated represents a StandardL2TokenCreated event raised by the OptimismMintableERC20Factory contract.
type OptimismMintableERC20FactoryStandardL2TokenCreated struct {
	RemoteToken common.Address
	LocalToken  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStandardL2TokenCreated is a free log retrieval operation binding the contract event 0xceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf.
//
// Solidity: event StandardL2TokenCreated(address indexed remoteToken, address indexed localToken)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryFilterer) FilterStandardL2TokenCreated(opts *bind.FilterOpts, remoteToken []common.Address, localToken []common.Address) (*OptimismMintableERC20FactoryStandardL2TokenCreatedIterator, error) {

	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _OptimismMintableERC20Factory.contract.FilterLogs(opts, "StandardL2TokenCreated", remoteTokenRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &OptimismMintableERC20FactoryStandardL2TokenCreatedIterator{contract: _OptimismMintableERC20Factory.contract, event: "StandardL2TokenCreated", logs: logs, sub: sub}, nil
}

// WatchStandardL2TokenCreated is a free log subscription operation binding the contract event 0xceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf.
//
// Solidity: event StandardL2TokenCreated(address indexed remoteToken, address indexed localToken)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryFilterer) WatchStandardL2TokenCreated(opts *bind.WatchOpts, sink chan<- *OptimismMintableERC20FactoryStandardL2TokenCreated, remoteToken []common.Address, localToken []common.Address) (event.Subscription, error) {

	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _OptimismMintableERC20Factory.contract.WatchLogs(opts, "StandardL2TokenCreated", remoteTokenRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimismMintableERC20FactoryStandardL2TokenCreated)
				if err := _OptimismMintableERC20Factory.contract.UnpackLog(event, "StandardL2TokenCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStandardL2TokenCreated is a log parse operation binding the contract event 0xceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf.
//
// Solidity: event StandardL2TokenCreated(address indexed remoteToken, address indexed localToken)
func (_OptimismMintableERC20Factory *OptimismMintableERC20FactoryFilterer) ParseStandardL2TokenCreated(log types.Log) (*OptimismMintableERC20FactoryStandardL2TokenCreated, error) {
	event := new(OptimismMintableERC20FactoryStandardL2TokenCreated)
	if err := _OptimismMintableERC20Factory.contract.UnpackLog(event, "StandardL2TokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
