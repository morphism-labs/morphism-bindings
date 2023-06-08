// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morphism-labs/morphism-bindings/solc"
)

const ZKEVMStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/ZKEVM.sol:ZKEVM\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/ZKEVM.sol:ZKEVM\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/ZKEVM.sol:ZKEVM\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1016_storage\"},{\"astId\":1003,\"contract\":\"contracts/L1/ZKEVM.sol:ZKEVM\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/L1/ZKEVM.sol:ZKEVM\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1015_storage\"},{\"astId\":1005,\"contract\":\"contracts/L1/ZKEVM.sol:ZKEVM\",\"label\":\"PORTAL\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_contract(OptimismPortal)1017\"},{\"astId\":1006,\"contract\":\"contracts/L1/ZKEVM.sol:ZKEVM\",\"label\":\"submitter\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_address\"},{\"astId\":1007,\"contract\":\"contracts/L1/ZKEVM.sol:ZKEVM\",\"label\":\"challenger\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_address\"},{\"astId\":1008,\"contract\":\"contracts/L1/ZKEVM.sol:ZKEVM\",\"label\":\"lastBatchSequenced\",\"offset\":20,\"slot\":\"103\",\"type\":\"t_uint64\"},{\"astId\":1009,\"contract\":\"contracts/L1/ZKEVM.sol:ZKEVM\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1010,\"contract\":\"contracts/L1/ZKEVM.sol:ZKEVM\",\"label\":\"commitments\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_mapping(t_uint64,t_bytes32)\"},{\"astId\":1011,\"contract\":\"contracts/L1/ZKEVM.sol:ZKEVM\",\"label\":\"stateRoots\",\"offset\":0,\"slot\":\"106\",\"type\":\"t_mapping(t_uint64,t_bytes32)\"},{\"astId\":1012,\"contract\":\"contracts/L1/ZKEVM.sol:ZKEVM\",\"label\":\"originTimestamps\",\"offset\":0,\"slot\":\"107\",\"type\":\"t_mapping(t_uint64,t_uint256)\"},{\"astId\":1013,\"contract\":\"contracts/L1/ZKEVM.sol:ZKEVM\",\"label\":\"confirmBatchIndex\",\"offset\":0,\"slot\":\"108\",\"type\":\"t_mapping(t_uint64,t_bool)\"},{\"astId\":1014,\"contract\":\"contracts/L1/ZKEVM.sol:ZKEVM\",\"label\":\"challenges\",\"offset\":0,\"slot\":\"109\",\"type\":\"t_mapping(t_uint64,t_struct(BatchChallenge)1018_storage)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1015_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1016_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_contract(OptimismPortal)1017\":{\"encoding\":\"inplace\",\"label\":\"contract OptimismPortal\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_mapping(t_uint64,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint64 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint64\",\"value\":\"t_bool\"},\"t_mapping(t_uint64,t_bytes32)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint64 =\u003e bytes32)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint64\",\"value\":\"t_bytes32\"},\"t_mapping(t_uint64,t_struct(BatchChallenge)1018_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint64 =\u003e struct ZKEVM.BatchChallenge)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint64\",\"value\":\"t_struct(BatchChallenge)1018_storage\"},\"t_mapping(t_uint64,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint64 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint64\",\"value\":\"t_uint256\"},\"t_struct(BatchChallenge)1018_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ZKEVM.BatchChallenge\",\"numberOfBytes\":\"128\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var ZKEVMStorageLayout = new(solc.StorageLayout)

var ZKEVMDeployedBin = "0x6080604052600436106101a15760003560e01c80638d644bb7116100e1578063c7ab20391161008a578063f2fde38b11610064578063f2fde38b146105b0578063f4daa291146105d0578063f71b51f3146105e7578063fc7e286d1461061757600080fd5b8063c7ab2039146104c3578063e1e158a51461057f578063eb1ec18f1461059b57600080fd5b8063956d7191116100bb578063956d719114610463578063acc1245a14610490578063c0c53b8b146104b057600080fd5b80638d644bb7146103f85780638da5cb5b1461040b5780638dc45d9a1461043657600080fd5b80632e1a7d4d1161014e578063534db0e211610128578063534db0e21461035157806359ef11201461037e5780636177fd181461039e578063715018a6146103e357600080fd5b80632e1a7d4d146102d75780633a4b66f1146102f7578063423fa856146102ff57600080fd5b80631b4082c91161017f5780631b4082c91461023f57806323ec85181461027a5780632b259441146102aa57600080fd5b806301117700146101a6578063032ecb38146101c85780630ff754ea146101e8575b600080fd5b3480156101b257600080fd5b506101c66101c1366004612062565b610644565b005b3480156101d457600080fd5b506101c66101e33660046120f4565b610990565b3480156101f457600080fd5b506065546102159073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561024b57600080fd5b5061026c61025a3660046120f4565b606b6020526000908152604090205481565b604051908152602001610236565b34801561028657600080fd5b5061029a610295366004612131565b610aa2565b6040519015158152602001610236565b3480156102b657600080fd5b5061026c6102c53660046120f4565b606a6020526000908152604090205481565b3480156102e357600080fd5b506101c66102f236600461214e565b610c1d565b6101c6610e21565b34801561030b57600080fd5b506067546103389074010000000000000000000000000000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610236565b34801561035d57600080fd5b506067546102159073ffffffffffffffffffffffffffffffffffffffff1681565b34801561038a57600080fd5b506101c6610399366004612167565b611004565b3480156103aa57600080fd5b5061029a6103b9366004612131565b73ffffffffffffffffffffffffffffffffffffffff16600090815260686020526040902054151590565b3480156103ef57600080fd5b506101c66112c3565b6101c66104063660046120f4565b6112d7565b34801561041757600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff16610215565b34801561044257600080fd5b506066546102159073ffffffffffffffffffffffffffffffffffffffff1681565b34801561046f57600080fd5b5061026c61047e3660046120f4565b60696020526000908152604090205481565b34801561049c57600080fd5b5061029a6104ab3660046120f4565b611723565b6101c66104be3660046121ea565b611790565b3480156104cf57600080fd5b506105356104de3660046120f4565b606d60205260009081526040902080546001820154600283015460039093015467ffffffffffffffff8316936801000000000000000090930473ffffffffffffffffffffffffffffffffffffffff16929060ff1685565b6040805167ffffffffffffffff909616865273ffffffffffffffffffffffffffffffffffffffff90941660208601529284019190915260608301521515608082015260a001610236565b34801561058b57600080fd5b5061026c670de0b6b3a764000081565b3480156105a757600080fd5b5061026c606481565b3480156105bc57600080fd5b506101c66105cb366004612131565b6119cb565b3480156105dc57600080fd5b5061026c620186a081565b3480156105f357600080fd5b5061029a6106023660046120f4565b606c6020526000908152604090205460ff1681565b34801561062357600080fd5b5061026c610632366004612131565b60686020526000908152604090205481565b60665473ffffffffffffffffffffffffffffffffffffffff166106c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f5375626d69747465722063616e6e6f742062652061646472657373283029000060448201526064015b60405180910390fd5b60665473ffffffffffffffffffffffffffffffffffffffff163314610749576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f43616c6c6572206e6f74207375626d697474657200000000000000000000000060448201526064016106bf565b60665473ffffffffffffffffffffffffffffffffffffffff16600090815260686020526040902054670de0b6b3a764000011156107e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f496e73756666696369656e74207374616b696e6720616d6f756e74000000000060448201526064016106bf565b606754819074010000000000000000000000000000000000000000900467ffffffffffffffff1660005b8281101561090a57606361f61860008061082583611a9e565b91509150600061086a8383878d8d8b81811061084357610843612235565b90506020028101906108559190612264565b61085f90806122a2565b506060949350505050565b805160209081028183012067ffffffffffffffff8a166000908152606990925260409091208190559091508a8a888181106108a7576108a7612235565b90506020028101906108b99190612264565b67ffffffffffffffff89166000908152606a6020908152604080832093820135909355606b905220429055876108ee8161233d565b985050505050505050808061090290612364565b91505061080c565b50606780547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000067ffffffffffffffff84811682029290921792839055604051920416907f85d226f7b6c307388b0488cb75509c840f0154cbed044f4b3c4e52af7056911990600090a250505050565b61099981611723565b156109a357600080fd5b67ffffffffffffffff81166000908152606b602052604081205442906109cd90620186a09061239c565b1190508015610a5e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43616e6e6f7420636f6e6669726d20626174636820696e73696465207468652060448201527f6368616c6c656e67652077696e646f770000000000000000000000000000000060648201526084016106bf565b5067ffffffffffffffff166000908152606c6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b60675460009074010000000000000000000000000000000000000000900467ffffffffffffffff168103610ad857506000919050565b60665473ffffffffffffffffffffffffffffffffffffffff90811690831603610b5757606754606c90600090610b329060019074010000000000000000000000000000000000000000900467ffffffffffffffff166123b4565b67ffffffffffffffff16815260208101919091526040016000205460ff161592915050565b60005b60675467ffffffffffffffff7401000000000000000000000000000000000000000090910481169082161015610c145767ffffffffffffffff81166000908152606d602052604090205473ffffffffffffffffffffffffffffffffffffffff8481166801000000000000000090920416148015610bf4575067ffffffffffffffff81166000908152606d602052604090206003015460ff16155b15610c025750600192915050565b80610c0c8161233d565b915050610b5a565b50506000919050565b610c2633610aa2565b15610c8d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5573657220697320696e206368616c6c656e676500000000000000000000000060448201526064016106bf565b3360009081526068602052604090205481610ca757600080fd5b8181610cbb670de0b6b3a76400008361239c565b1115610de25760665473ffffffffffffffffffffffffffffffffffffffff1633148015610d0b575060675474010000000000000000000000000000000000000000900467ffffffffffffffff1615155b15610de25760675474010000000000000000000000000000000000000000900467ffffffffffffffff166000908152606b60205260409020544290610d5490620186a09061239c565b1115610de2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f7375626d69747465722073686f756c64207761697420626174636820746f206260448201527f6520636f6e6669726d000000000000000000000000000000000000000000000060648201526084016106bf565b81831115610ded5750805b3360009081526068602052604081208054859290610e0c9084906123dd565b90915550610e1c90503382611b31565b505050565b60665473ffffffffffffffffffffffffffffffffffffffff16610ea0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f5375626d69747465722063616e6e6f742062652061646472657373283029000060448201526064016106bf565b60665473ffffffffffffffffffffffffffffffffffffffff163314610f21576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f43616c6c6572206e6f74207375626d697474657200000000000000000000000060448201526064016106bf565b60665473ffffffffffffffffffffffffffffffffffffffff16600090815260686020526040902054670de0b6b3a764000090610f5e90349061239c565b1015610fc6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f446f206e6f74206861766520656e6f756768206465706f73697400000000000060448201526064016106bf565b60665473ffffffffffffffffffffffffffffffffffffffff1660009081526068602052604081208054349290610ffd90849061239c565b9091555050565b67ffffffffffffffff83166000908152606d602052604090205468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1661104957600080fd5b67ffffffffffffffff83166000908152606d602052604090206003015460ff16156110d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f4368616c6c656e676520616c72656164792066696e697368656400000000000060448201526064016106bf565b67ffffffffffffffff83166000908152606d602052604081206002015442906110fb9060649061239c565b1190508061114757611142846040518060400160405280600781526020017f74696d656f757400000000000000000000000000000000000000000000000000815250611bdd565b61127a565b816111ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f496e76616c69642070726f6f660000000000000000000000000000000000000060448201526064016106bf565b67ffffffffffffffff84166000908152606960205260409020546111d59084908490611cdb565b61123b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f50726f7665206661696c6564000000000000000000000000000000000000000060448201526064016106bf565b61127a846040518060400160405280600d81526020017f70726f7665207375636365737300000000000000000000000000000000000000815250611cf7565b50505067ffffffffffffffff166000908152606d6020526040902060030180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b6112cb611e11565b6112d56000611e92565b565b60675473ffffffffffffffffffffffffffffffffffffffff16611356576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4368616c6c656e6765722063616e6e6f7420626520616464726573732830290060448201526064016106bf565b60675473ffffffffffffffffffffffffffffffffffffffff1633146113d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f43616c6c6572206e6f74206368616c6c656e676572000000000000000000000060448201526064016106bf565b67ffffffffffffffff81166000908152606b60205260408120549003611459576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4261746368206e6f74206578697374000000000000000000000000000000000060448201526064016106bf565b67ffffffffffffffff81166000908152606d602052604090205468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff16156114fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f416c726561647920686173206368616c6c656e6765000000000000000000000060448201526064016106bf565b67ffffffffffffffff81166000908152606b6020526040812054429061152690620186a09061239c565b119050806115b6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f43616e6e6f74206368616c6c656e6765206261746368206f757473696465207460448201527f6865206368616c6c656e67652077696e646f770000000000000000000000000060648201526084016106bf565b670de0b6b3a76400003410156115cb57600080fd5b60675473ffffffffffffffffffffffffffffffffffffffff166000908152606860205260408120805434929061160290849061239c565b90915550506040805160a08101825267ffffffffffffffff848116808352336020808501828152348688018181524260608901908152600060808a01818152888252606d8752908b902099518a54955173ffffffffffffffffffffffffffffffffffffffff1668010000000000000000027fffffffff00000000000000000000000000000000000000000000000000000000909616991698909817939093178855516001880155905160028701559351600390950180549515157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0090961695909517909455845190815292830191909152917fff4bd2fd14b0670e131544fdc1e7c53f06371ccac4a3a8e0445a93d02d68fb05910160405180910390a25050565b67ffffffffffffffff81166000908152606d602052604081205468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff161580159061178a575067ffffffffffffffff82166000908152606d602052604090206003015460ff16155b92915050565b600054610100900460ff16158080156117b05750600054600160ff909116105b806117ca5750303b1580156117ca575060005460ff166001145b611856576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016106bf565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156118b457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6118bc611f09565b606780546065805473ffffffffffffffffffffffffffffffffffffffff8089167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255606680548884169216821790559085167fffffffff00000000000000000000000000000000000000000000000000000000909216919091179091556000908152606860205260408120805434929061195d90849061239c565b909155505080156119c557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b6119d3611e11565b73ffffffffffffffffffffffffffffffffffffffff8116611a76576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016106bf565b611a7f81611e92565b50565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b60008061f6188311611ab7575060039261290492509050565b620493e08311611acf5750600e926201107692509050565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f434952435549545f434f4e46494700000000000000000000000000000000000060448201526064016106bf565b6000611b4e835a8460405180602001604052806000815250611fa8565b905080610e1c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c6564000000000000000000000000000000000000000000000000000000000060648201526084016106bf565b67ffffffffffffffff82166000908152606d602090815260408083205460665473ffffffffffffffffffffffffffffffffffffffff90811685526068909352908320805468010000000000000000909204909216929091670de0b6b3a764000091611c4883856123dd565b909155505073ffffffffffffffffffffffffffffffffffffffff821660009081526068602052604081208054670de0b6b3a76400009290611c8a90849061239c565b925050819055508367ffffffffffffffff167fe70d3820e244d5f71d1a6395db24f3460e8dca966edc1fd3625b6292880a877a8385604051611ccd9291906123f4565b60405180910390a250505050565b6000828103611cec57506000611cf0565b5060015b9392505050565b67ffffffffffffffff82166000908152606d60209081526040808320805460019091015460665473ffffffffffffffffffffffffffffffffffffffff908116865260689094528285205468010000000000000000909204909316808552918420805492949192849290611d6b9084906123dd565b909155505060665473ffffffffffffffffffffffffffffffffffffffff1660009081526068602052604081208054849290611da790849061239c565b909155505060665460405167ffffffffffffffff8716917fe70d3820e244d5f71d1a6395db24f3460e8dca966edc1fd3625b6292880a877a91611e029173ffffffffffffffffffffffffffffffffffffffff169088906123f4565b60405180910390a25050505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146112d5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106bf565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16611fa0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106bf565b6112d5611fc2565b600080600080845160208601878a8af19695505050505050565b600054610100900460ff16612059576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106bf565b6112d533611e92565b6000806020838503121561207557600080fd5b823567ffffffffffffffff8082111561208d57600080fd5b818501915085601f8301126120a157600080fd5b8135818111156120b057600080fd5b8660208260051b85010111156120c557600080fd5b60209290920196919550909350505050565b803567ffffffffffffffff811681146120ef57600080fd5b919050565b60006020828403121561210657600080fd5b611cf0826120d7565b73ffffffffffffffffffffffffffffffffffffffff81168114611a7f57600080fd5b60006020828403121561214357600080fd5b8135611cf08161210f565b60006020828403121561216057600080fd5b5035919050565b60008060006040848603121561217c57600080fd5b612185846120d7565b9250602084013567ffffffffffffffff808211156121a257600080fd5b818601915086601f8301126121b657600080fd5b8135818111156121c557600080fd5b8760208285010111156121d757600080fd5b6020830194508093505050509250925092565b6000806000606084860312156121ff57600080fd5b833561220a8161210f565b9250602084013561221a8161210f565b9150604084013561222a8161210f565b809150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6183360301811261229857600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126122d757600080fd5b83018035915067ffffffffffffffff8211156122f257600080fd5b60200191503681900382131561230757600080fd5b9250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681810361235a5761235a61230e565b6001019392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036123955761239561230e565b5060010190565b600082198211156123af576123af61230e565b500190565b600067ffffffffffffffff838116908316818110156123d5576123d561230e565b039392505050565b6000828210156123ef576123ef61230e565b500390565b73ffffffffffffffffffffffffffffffffffffffff8316815260006020604081840152835180604085015260005b8181101561243e57858101830151858201606001528201612422565b81811115612450576000606083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160600194935050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(ZKEVMStorageLayoutJSON), ZKEVMStorageLayout); err != nil {
		panic(err)
	}

	layouts["ZKEVM"] = ZKEVMStorageLayout
	deployedBytecodes["ZKEVM"] = ZKEVMDeployedBin
}