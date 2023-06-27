// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morphism-labs/morphism-bindings/solc"
)

const OptimismPortalStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/OptimismPortal.sol:OptimismPortal\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/OptimismPortal.sol:OptimismPortal\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/OptimismPortal.sol:OptimismPortal\",\"label\":\"params\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_struct(ResourceParams)1012_storage\"},{\"astId\":1003,\"contract\":\"contracts/L1/OptimismPortal.sol:OptimismPortal\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_array(t_uint256)1009_storage\"},{\"astId\":1004,\"contract\":\"contracts/L1/OptimismPortal.sol:OptimismPortal\",\"label\":\"l2Sender\",\"offset\":0,\"slot\":\"50\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"contracts/L1/OptimismPortal.sol:OptimismPortal\",\"label\":\"finalizedWithdrawals\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1006,\"contract\":\"contracts/L1/OptimismPortal.sol:OptimismPortal\",\"label\":\"provenWithdrawals\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_mapping(t_bytes32,t_struct(ProvenWithdrawal)1011_storage)\"},{\"astId\":1007,\"contract\":\"contracts/L1/OptimismPortal.sol:OptimismPortal\",\"label\":\"paused\",\"offset\":0,\"slot\":\"53\",\"type\":\"t_bool\"},{\"astId\":1008,\"contract\":\"contracts/L1/OptimismPortal.sol:OptimismPortal\",\"label\":\"zkevm\",\"offset\":1,\"slot\":\"53\",\"type\":\"t_contract(ZKEVM)1010\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1009_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[48]\",\"numberOfBytes\":\"1536\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_contract(ZKEVM)1010\":{\"encoding\":\"inplace\",\"label\":\"contract ZKEVM\",\"numberOfBytes\":\"20\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_mapping(t_bytes32,t_struct(ProvenWithdrawal)1011_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e struct OptimismPortal.ProvenWithdrawal)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_struct(ProvenWithdrawal)1011_storage\"},\"t_struct(ProvenWithdrawal)1011_storage\":{\"encoding\":\"inplace\",\"label\":\"struct OptimismPortal.ProvenWithdrawal\",\"numberOfBytes\":\"64\"},\"t_struct(ResourceParams)1012_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ResourceMetering.ResourceParams\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var OptimismPortalStorageLayout = new(solc.StorageLayout)

var OptimismPortalDeployedBin = "0x6080604052600436106101625760003560e01c80638c3152e9116100c0578063cff0ab9611610074578063e965084c11610059578063e965084c1461049f578063e9e05c421461052b578063f04987501461053e57600080fd5b8063cff0ab96146103de578063d53a822f1461047f57600080fd5b8063a14238e7116100a5578063a14238e714610343578063a35d99df14610373578063b8ded9c7146103ac57600080fd5b80638c3152e9146102f65780639bf62d821461031657600080fd5b80635c975abb11610117578063724c184c116100fc578063724c184c146102ad5780638456cb59146102e15780638b4c40b01461018757600080fd5b80635c975abb146102735780636dbffb781461028d57600080fd5b80633f4ba83a116101485780633f4ba83a1461021c57806353f3fd1a1461023157806354fd4d501461025157600080fd5b80621c2ff61461018e578063340735f7146101ec57600080fd5b36610189576101873334620186a0600060405180602001604052806000815250610572565b005b600080fd5b34801561019a57600080fd5b506101c27f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101f857600080fd5b5061020c610207366004612afd565b61080d565b60405190151581526020016101e3565b34801561022857600080fd5b506101876108e1565b34801561023d57600080fd5b5061018761024c366004612ce4565b610a04565b34801561025d57600080fd5b50610266610f3c565b6040516101e39190612df2565b34801561027f57600080fd5b5060355461020c9060ff1681565b34801561029957600080fd5b5061020c6102a8366004612e05565b610fdf565b3480156102b957600080fd5b506101c27f000000000000000000000000000000000000000000000000000000000000000081565b3480156102ed57600080fd5b506101876110b6565b34801561030257600080fd5b50610187610311366004612e1e565b6111d6565b34801561032257600080fd5b506032546101c29073ffffffffffffffffffffffffffffffffffffffff1681565b34801561034f57600080fd5b5061020c61035e366004612e05565b60336020526000908152604090205460ff1681565b34801561037f57600080fd5b5061039361038e366004612e6b565b611ab1565b60405167ffffffffffffffff90911681526020016101e3565b3480156103b857600080fd5b506035546101c290610100900473ffffffffffffffffffffffffffffffffffffffff1681565b3480156103ea57600080fd5b50600154610446906fffffffffffffffffffffffffffffffff81169067ffffffffffffffff7001000000000000000000000000000000008204811691780100000000000000000000000000000000000000000000000090041683565b604080516fffffffffffffffffffffffffffffffff909416845267ffffffffffffffff92831660208501529116908201526060016101e3565b34801561048b57600080fd5b5061018761049a366004612e96565b611aca565b3480156104ab57600080fd5b506104fd6104ba366004612e05565b603460205260009081526040902080546001909101546fffffffffffffffffffffffffffffffff8082169170010000000000000000000000000000000090041683565b604080519384526fffffffffffffffffffffffffffffffff92831660208501529116908201526060016101e3565b610187610539366004612eb1565b610572565b34801561054a57600080fd5b506101c27f000000000000000000000000000000000000000000000000000000000000000081565b8260005a905083156106295773ffffffffffffffffffffffffffffffffffffffff87161561062957604080517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260248101919091527f4f7074696d69736d506f7274616c3a206d7573742073656e6420746f2061646460448201527f72657373283029207768656e206372656174696e67206120636f6e747261637460648201526084015b60405180910390fd5b6106338351611ab1565b67ffffffffffffffff168567ffffffffffffffff1610156106d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f4f7074696d69736d506f7274616c3a20676173206c696d697420746f6f20736d60448201527f616c6c00000000000000000000000000000000000000000000000000000000006064820152608401610620565b6201d4c083511115610744576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4f7074696d69736d506f7274616c3a206461746120746f6f206c6172676500006044820152606401610620565b33328114610765575033731111000000000000000000000000000000001111015b60003488888888604051602001610780959493929190612f2a565b604051602081830303815290604052905060008973ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32846040516107f09190612df2565b60405180910390a450506108048282611cd3565b50505050505050565b600084815b60208110156108d5578085901c6001166001036108785785816020811061083b5761083b612f8f565b60200201358260405160200161085b929190918252602082015260400190565b6040516020818303038152906040528051906020012091506108c3565b8186826020811061088b5761088b612f8f565b60200201356040516020016108aa929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b806108cd81612fed565b915050610812565b50909114949350505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146109a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4f7074696d69736d506f7274616c3a206f6e6c7920677561726469616e20636160448201527f6e20756e706175736500000000000000000000000000000000000000000000006064820152608401610620565b603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040513381527f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa906020015b60405180910390a1565b60355460ff1615610a71576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f7074696d69736d506f7274616c3a20706175736564000000000000000000006044820152606401610620565b3073ffffffffffffffffffffffffffffffffffffffff16846040015173ffffffffffffffffffffffffffffffffffffffff1603610b30576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4f7074696d69736d506f7274616c3a20796f752063616e6e6f742073656e642060448201527f6d6573736167657320746f2074686520706f7274616c20636f6e7472616374006064820152608401610620565b6040517fa25ae557000000000000000000000000000000000000000000000000000000008152600481018490526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a25ae55790602401606060405180830381865afa158015610bbe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610be29190613045565b519050610bfc610bf7368590038501856130aa565b612000565b8114610c8a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4f7074696d69736d506f7274616c3a20696e76616c6964206f7574707574207260448201527f6f6f742070726f6f6600000000000000000000000000000000000000000000006064820152608401610620565b6000610c958661205c565b6000818152603460209081526040918290208251606081018452815481526001909101546fffffffffffffffffffffffffffffffff8082169383018490527001000000000000000000000000000000009091041692810192909252919250901580610dc75750805160408083015190517fa25ae5570000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff90911660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a25ae55790602401606060405180830381865afa158015610d9f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc39190613045565b5114155b610e53576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f4f7074696d69736d506f7274616c3a207769746864726177616c20686173682060448201527f68617320616c7265616479206265656e2070726f76656e0000000000000000006064820152608401610620565b60408051602080820185905260008284015282518083038401815260609092019092528051910120604080516060810182528581526fffffffffffffffffffffffffffffffff42811660208084019182528b831684860190815260008981526034835286812095518655925190518416700100000000000000000000000000000000029316929092176001909301929092558a830151908b0151925173ffffffffffffffffffffffffffffffffffffffff918216939091169186917f67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f629190a45050505050505050565b6060610f677f000000000000000000000000000000000000000000000000000000000000000061208c565b610f907f000000000000000000000000000000000000000000000000000000000000000061208c565b610fb97f000000000000000000000000000000000000000000000000000000000000000061208c565b604051602001610fcb93929190613110565b604051602081830303815290604052905090565b6040517fa25ae557000000000000000000000000000000000000000000000000000000008152600481018290526000906110b09073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063a25ae55790602401606060405180830381865afa158015611071573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110959190613045565b602001516fffffffffffffffffffffffffffffffff166121c9565b92915050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461117b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f4f7074696d69736d506f7274616c3a206f6e6c7920677561726469616e20636160448201527f6e207061757365000000000000000000000000000000000000000000000000006064820152608401610620565b603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040513381527f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258906020016109fa565b60355460ff1615611243576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f7074696d69736d506f7274616c3a20706175736564000000000000000000006044820152606401610620565b60325473ffffffffffffffffffffffffffffffffffffffff1661dead146112ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4f7074696d69736d506f7274616c3a2063616e206f6e6c79207472696767657260448201527f206f6e65207769746864726177616c20706572207472616e73616374696f6e006064820152608401610620565b60006112f78261205c565b60008181526034602090815260408083208151606081018352815481526001909101546fffffffffffffffffffffffffffffffff808216948301859052700100000000000000000000000000000000909104169181019190915292935090036113e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4f7074696d69736d506f7274616c3a207769746864726177616c20686173206e60448201527f6f74206265656e2070726f76656e2079657400000000000000000000000000006064820152608401610620565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663887862726040518163ffffffff1660e01b8152600401602060405180830381865afa15801561144d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114719190613186565b81602001516fffffffffffffffffffffffffffffffff16101561153c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604b60248201527f4f7074696d69736d506f7274616c3a207769746864726177616c2074696d657360448201527f74616d70206c657373207468616e204c32204f7261636c65207374617274696e60648201527f672074696d657374616d70000000000000000000000000000000000000000000608482015260a401610620565b61155b81602001516fffffffffffffffffffffffffffffffff166121c9565b61160d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604560248201527f4f7074696d69736d506f7274616c3a2070726f76656e2077697468647261776160448201527f6c2066696e616c697a6174696f6e20706572696f6420686173206e6f7420656c60648201527f6170736564000000000000000000000000000000000000000000000000000000608482015260a401610620565b60408181015190517fa25ae5570000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff90911660048201526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a25ae55790602401606060405180830381865afa1580156116b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116d69190613045565b8251815191925014611790576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604960248201527f4f7074696d69736d506f7274616c3a206f757470757420726f6f742070726f7660448201527f656e206973206e6f74207468652073616d652061732063757272656e74206f7560648201527f7470757420726f6f740000000000000000000000000000000000000000000000608482015260a401610620565b6117af81602001516fffffffffffffffffffffffffffffffff166121c9565b611861576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f4f7074696d69736d506f7274616c3a206f75747075742070726f706f73616c2060448201527f66696e616c697a6174696f6e20706572696f6420686173206e6f7420656c617060648201527f7365640000000000000000000000000000000000000000000000000000000000608482015260a401610620565b60008381526033602052604090205460ff1615611900576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f4f7074696d69736d506f7274616c3a207769746864726177616c20686173206160448201527f6c7265616479206265656e2066696e616c697a656400000000000000000000006064820152608401610620565b600083815260336020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055908601516032805473ffffffffffffffffffffffffffffffffffffffff9092167fffffffffffffffffffffffff00000000000000000000000000000000000000009092169190911790558501516080860151606087015160a08801516119a29392919061226c565b603280547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead17905560405190915084907fdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b90611a0790841515815260200190565b60405180910390a280158015611a1d5750326001145b15611aaa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f4f7074696d69736d506f7274616c3a207769746864726177616c206661696c6560448201527f64000000000000000000000000000000000000000000000000000000000000006064820152608401610620565b5050505050565b6000611abe82601061319f565b6110b0906152086131cf565b600054610100900460ff1615808015611aea5750600054600160ff909116105b80611b045750303b158015611b04575060005460ff166001145b611b90576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610620565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015611bee57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b603280547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055603580548315157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909116179055611c506122ca565b8015611cb357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b600154600090611d09907801000000000000000000000000000000000000000000000000900467ffffffffffffffff16436131fb565b90506000611d156123ad565b90506000816020015160ff16826000015163ffffffff16611d369190613241565b90508215611e6d57600154600090611d6d908390700100000000000000000000000000000000900467ffffffffffffffff166132a9565b90506000836040015160ff1683611d84919061331d565b600154611da49084906fffffffffffffffffffffffffffffffff1661331d565b611dae9190613241565b600154909150600090611dff90611dd89084906fffffffffffffffffffffffffffffffff166133d9565b866060015163ffffffff168760a001516fffffffffffffffffffffffffffffffff16612473565b90506001861115611e2e57611e2b611dd882876040015160ff1660018a611e2691906131fb565b612488565b90505b6fffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff4316021760015550505b60018054869190601090611ea0908490700100000000000000000000000000000000900467ffffffffffffffff166131cf565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000015163ffffffff16600160000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff161315611f83576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f5265736f757263654d65746572696e673a2063616e6e6f7420627579206d6f7260448201527f6520676173207468616e20617661696c61626c6520676173206c696d697400006064820152608401610620565b600154600090611faf906fffffffffffffffffffffffffffffffff1667ffffffffffffffff881661344d565b90506000611fc148633b9aca006124dd565b611fcb908361348a565b905060005a611fda90886131fb565b905080821115611ff657611ff6611ff182846131fb565b6124f6565b5050505050505050565b6000816000015182602001518360400151846060015160405160200161203f949392919093845260208401929092526040830152606082015260800190565b604051602081830303815290604052805190602001209050919050565b80516020808301516040808501516060860151608087015160a0880151935160009761203f97909695910161349e565b6060816000036120cf57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156120f957806120e381612fed565b91506120f29050600a8361348a565b91506120d3565b60008167ffffffffffffffff81111561211457612114612b3c565b6040519080825280601f01601f19166020018201604052801561213e576020820181803683370190505b5090505b84156121c1576121536001836131fb565b9150612160600a866134f5565b61216b906030613509565b60f81b81838151811061218057612180612f8f565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506121ba600a8661348a565b9450612142565b949350505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f4daa2916040518163ffffffff1660e01b8152600401602060405180830381865afa158015612236573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061225a9190613186565b6122649083613509565b421192915050565b600080600061227c866000612524565b9050806122b2576308c379a06000526020805278185361666543616c6c3a204e6f7420656e6f756768206761736058526064601cfd5b600080855160208701888b5af1979650505050505050565b600054610100900460ff16612361576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610620565b60408051606081018252633b9aca00808252600060208301524367ffffffffffffffff169190920181905278010000000000000000000000000000000000000000000000000217600155565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a08101919091527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663cc731b026040518163ffffffff1660e01b815260040160c060405180830381865afa15801561244a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061246e9190613546565b905090565b60006121c16124828585612542565b83612552565b6000670de0b6b3a76400006124c96124a08583613241565b6124b290670de0b6b3a76400006132a9565b6124c485670de0b6b3a764000061331d565b612561565b6124d3908661331d565b6121c19190613241565b6000818310156124ed57816124ef565b825b9392505050565b6000805a90505b825a61250990836131fb565b101561251f5761251882612fed565b91506124fd565b505050565b600080603f83619c4001026040850201603f5a021015949350505050565b6000818312156124ed57816124ef565b60008183126124ed57816124ef565b60006124ef670de0b6b3a76400008361257986612592565b612583919061331d565b61258d9190613241565b6127d6565b60008082136125fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f554e444546494e454400000000000000000000000000000000000000000000006044820152606401610620565b6000606061260a84612a15565b03609f8181039490941b90931c6c465772b2bbbb5f824b15207a3081018102606090811d6d0388eaa27412d5aca026815d636e018202811d6d0df99ac502031bf953eff472fdcc018202811d6d13cdffb29d51d99322bdff5f2211018202811d6d0a0f742023def783a307a986912e018202811d6d01920d8043ca89b5239253284e42018202811d6c0b7a86d7375468fac667a0a527016c29508e458543d8aa4df2abee7883018302821d6d0139601a2efabe717e604cbb4894018302821d6d02247f7a7b6594320649aa03aba1018302821d7fffffffffffffffffffffffffffffffffffffff73c0c716a594e00d54e3c4cbc9018302821d7ffffffffffffffffffffffffffffffffffffffdc7b88c420e53a9890533129f6f01830290911d7fffffffffffffffffffffffffffffffffffffff465fda27eb4d63ded474e5f832019091027ffffffffffffffff5f6af8f7b3396644f18e157960000000000000000000000000105711340daa0d5f769dba1915cef59f0815a5506027d0267a36c0c95b3975ab3ee5b203a7614a3f75373f047d803ae7b6687f2b393909302929092017d57115e47018c7177eebf7cd370a3356a1b7863008a5ae8028c72b88642840160ae1d92915050565b60007ffffffffffffffffffffffffffffffffffffffffffffffffdb731c958f34d94c1821361280757506000919050565b680755bf798b4a1bf1e58212612879576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4558505f4f564552464c4f5700000000000000000000000000000000000000006044820152606401610620565b6503782dace9d9604e83901b059150600060606bb17217f7d1cf79abc9e3b39884821b056b80000000000000000000000001901d6bb17217f7d1cf79abc9e3b39881029093037fffffffffffffffffffffffffffffffffffffffdbf3ccf1604d263450f02a550481018102606090811d6d0277594991cfc85f6e2461837cd9018202811d7fffffffffffffffffffffffffffffffffffffe5adedaa1cb095af9e4da10e363c018202811d6db1bbb201f443cf962f1a1d3db4a5018202811d7ffffffffffffffffffffffffffffffffffffd38dc772608b0ae56cce01296c0eb018202811d6e05180bb14799ab47a8a8cb2a527d57016d02d16720577bd19bf614176fe9ea6c10fe68e7fd37d0007b713f765084018402831d9081019084017ffffffffffffffffffffffffffffffffffffffe2c69812cf03b0763fd454a8f7e010290911d6e0587f503bb6ea29d25fcb7401964500190910279d835ebba824c98fb31b83b2ca45c000000000000000000000000010574029d9dc38563c32e5c2f6dc192ee70ef65f9978af30260c3939093039290921c92915050565b6000808211612a80576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f554e444546494e454400000000000000000000000000000000000000000000006044820152606401610620565b5060016fffffffffffffffffffffffffffffffff821160071b82811c67ffffffffffffffff1060061b1782811c63ffffffff1060051b1782811c61ffff1060041b1782811c60ff10600390811b90911783811c600f1060021b1783811c909110821b1791821c111790565b8061040081018310156110b057600080fd5b6000806000806104608587031215612b1457600080fd5b84359350612b258660208701612aeb565b939693955050505061042082013591610440013590565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b803573ffffffffffffffffffffffffffffffffffffffff81168114612b8f57600080fd5b919050565b600082601f830112612ba557600080fd5b813567ffffffffffffffff80821115612bc057612bc0612b3c565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715612c0657612c06612b3c565b81604052838152866020858801011115612c1f57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060c08284031215612c5157600080fd5b60405160c0810167ffffffffffffffff8282108183111715612c7557612c75612b3c565b8160405282935084358352612c8c60208601612b6b565b6020840152612c9d60408601612b6b565b6040840152606085013560608401526080850135608084015260a0850135915080821115612cca57600080fd5b50612cd785828601612b94565b60a0830152505092915050565b6000806000808486036104c0811215612cfc57600080fd5b853567ffffffffffffffff811115612d1357600080fd5b612d1f88828901612c3f565b9550506020860135935060807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082011215612d5957600080fd5b50604085019150612d6d8660c08701612aeb565b905092959194509250565b60005b83811015612d93578181015183820152602001612d7b565b83811115612da2576000848401525b50505050565b60008151808452612dc0816020860160208601612d78565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006124ef6020830184612da8565b600060208284031215612e1757600080fd5b5035919050565b600060208284031215612e3057600080fd5b813567ffffffffffffffff811115612e4757600080fd5b6121c184828501612c3f565b803567ffffffffffffffff81168114612b8f57600080fd5b600060208284031215612e7d57600080fd5b6124ef82612e53565b80358015158114612b8f57600080fd5b600060208284031215612ea857600080fd5b6124ef82612e86565b600080600080600060a08688031215612ec957600080fd5b612ed286612b6b565b945060208601359350612ee760408701612e53565b9250612ef560608701612e86565b9150608086013567ffffffffffffffff811115612f1157600080fd5b612f1d88828901612b94565b9150509295509295909350565b8581528460208201527fffffffffffffffff0000000000000000000000000000000000000000000000008460c01b16604082015282151560f81b604882015260008251612f7e816049850160208701612d78565b919091016049019695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361301e5761301e612fbe565b5060010190565b80516fffffffffffffffffffffffffffffffff81168114612b8f57600080fd5b60006060828403121561305757600080fd5b6040516060810181811067ffffffffffffffff8211171561307a5761307a612b3c565b6040528251815261308d60208401613025565b602082015261309e60408401613025565b60408201529392505050565b6000608082840312156130bc57600080fd5b6040516080810181811067ffffffffffffffff821117156130df576130df612b3c565b8060405250823581526020830135602082015260408301356040820152606083013560608201528091505092915050565b60008451613122818460208901612d78565b80830190507f2e00000000000000000000000000000000000000000000000000000000000000808252855161315e816001850160208a01612d78565b60019201918201528351613179816002840160208801612d78565b0160020195945050505050565b60006020828403121561319857600080fd5b5051919050565b600067ffffffffffffffff808316818516818304811182151516156131c6576131c6612fbe565b02949350505050565b600067ffffffffffffffff8083168185168083038211156131f2576131f2612fbe565b01949350505050565b60008282101561320d5761320d612fbe565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261325057613250613212565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83147f8000000000000000000000000000000000000000000000000000000000000000831416156132a4576132a4612fbe565b500590565b6000808312837f8000000000000000000000000000000000000000000000000000000000000000018312811516156132e3576132e3612fbe565b837f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01831381161561331757613317612fbe565b50500390565b60007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60008413600084138583048511828216161561335e5761335e612fbe565b7f8000000000000000000000000000000000000000000000000000000000000000600087128682058812818416161561339957613399612fbe565b600087129250878205871284841616156133b5576133b5612fbe565b878505871281841616156133cb576133cb612fbe565b505050929093029392505050565b6000808212827f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0384138115161561341357613413612fbe565b827f800000000000000000000000000000000000000000000000000000000000000003841281161561344757613447612fbe565b50500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561348557613485612fbe565b500290565b60008261349957613499613212565b500490565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a08301526134e960c0830184612da8565b98975050505050505050565b60008261350457613504613212565b500690565b6000821982111561351c5761351c612fbe565b500190565b805163ffffffff81168114612b8f57600080fd5b805160ff81168114612b8f57600080fd5b600060c0828403121561355857600080fd5b60405160c0810181811067ffffffffffffffff8211171561357b5761357b612b3c565b60405261358783613521565b815261359560208401613535565b60208201526135a660408401613535565b60408201526135b760608401613521565b60608201526135c860808401613521565b60808201526135d960a08401613025565b60a0820152939250505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(OptimismPortalStorageLayoutJSON), OptimismPortalStorageLayout); err != nil {
		panic(err)
	}

	layouts["OptimismPortal"] = OptimismPortalStorageLayout
	deployedBytecodes["OptimismPortal"] = OptimismPortalDeployedBin
}
