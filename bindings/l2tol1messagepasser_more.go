// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morphism-labs/morphism-bindings/solc"
)

const L2ToL1MessagePasserStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser\",\"label\":\"_branch\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_array(t_bytes1006)32_storage\"},{\"astId\":1001,\"contract\":\"contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser\",\"label\":\"leafNodesCount\",\"offset\":0,\"slot\":\"32\",\"type\":\"t_uint256\"},{\"astId\":1002,\"contract\":\"contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser\",\"label\":\"_gap\",\"offset\":0,\"slot\":\"33\",\"type\":\"t_array(t_uint256)1007_storage\"},{\"astId\":1003,\"contract\":\"contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser\",\"label\":\"sentMessages\",\"offset\":0,\"slot\":\"43\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1004,\"contract\":\"contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"44\",\"type\":\"t_uint240\"},{\"astId\":1005,\"contract\":\"contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser\",\"label\":\"messageRoot\",\"offset\":0,\"slot\":\"45\",\"type\":\"t_bytes32\"}],\"types\":{\"t_array(t_bytes1006)32_storage\":{\"encoding\":\"inplace\",\"label\":\"bytes32[32]\",\"numberOfBytes\":\"1024\"},\"t_array(t_uint256)1007_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[10]\",\"numberOfBytes\":\"320\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L2ToL1MessagePasserStorageLayout = new(solc.StorageLayout)

var L2ToL1MessagePasserDeployedBin = "0x6080604052600436106100b55760003560e01c806389c09d3811610069578063c2b3e5ac1161004e578063c2b3e5ac146101db578063d4b9f4fa146101ee578063ecc704281461020457600080fd5b806389c09d38146101a2578063b58343bb146101c557600080fd5b806344df8e701161009a57806344df8e701461013b57806354fd4d501461015057806382e3702d1461017257600080fd5b8063340735f7146100de5780633f827a5a1461011357600080fd5b366100d9576100d733620186a060405180602001604052806000815250610219565b005b600080fd5b3480156100ea57600080fd5b506100fe6100f9366004610931565b61033b565b60405190151581526020015b60405180910390f35b34801561011f57600080fd5b50610128600181565b60405161ffff909116815260200161010a565b34801561014757600080fd5b506100d761040f565b34801561015c57600080fd5b50610165610447565b60405161010a91906109ef565b34801561017e57600080fd5b506100fe61018d366004610a09565b602b6020526000908152604090205460ff1681565b3480156101ae57600080fd5b506101b76104ea565b60405190815260200161010a565b3480156101d157600080fd5b506101b760205481565b6100d76101e9366004610a51565b610219565b3480156101fa57600080fd5b506101b7602d5481565b34801561021057600080fd5b506101b76105c7565b600061026d6040518060c001604052806102316105c7565b815233602082015273ffffffffffffffffffffffffffffffffffffffff871660408201523460608201526080810186905260a0018490526105fd565b90506102788161064a565b6102806104ea565b602d5573ffffffffffffffffffffffffffffffffffffffff8416336102a36105c7565b7f3034a173cd04c8b47937b2f95d00794df516ce57df60210b71de7ecf6d2fe4f334878787602d546040516102dc959493929190610b55565b60405180910390a45050602c80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b600084815b6020811015610403578085901c6001166001036103a65785816020811061036957610369610b88565b602002013582604051602001610389929190918252602082015260400190565b6040516020818303038152906040528051906020012091506103f1565b818682602081106103b9576103b9610b88565b60200201356040516020016103d8929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b806103fb81610be6565b915050610340565b50909114949350505050565b476104198161074d565b60405181907f7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f90600090a250565b60606104727f0000000000000000000000000000000000000000000000000000000000000000610777565b61049b7f0000000000000000000000000000000000000000000000000000000000000000610777565b6104c47f0000000000000000000000000000000000000000000000000000000000000000610777565b6040516020016104d693929190610c1e565b604051602081830303815290604052905090565b602054600090819081805b60208110156105be578083901c600116600103610552576000816020811061051f5761051f610b88565b0154604080516020810192909252810185905260600160405160208183030381529060405280519060200120935061057f565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b604080516020810184905290810183905260600160405160208183030381529060405280519060200120915080806105b690610be6565b9150506104f5565b50919392505050565b602c546000906105f8907dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff166001610835565b905090565b80516020808301516040808501516060860151608087015160a0880151935160009761062d979096959101610c94565b604051602081830303815290604052805190602001209050919050565b80600161065960206002610e0b565b6106639190610e17565b6020541061069d576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020600081546106ae90610be6565b9182905550905060005b602081101561073f578082901c6001166001036106eb5782600082602081106106e3576106e3610b88565b015550505050565b600081602081106106fe576106fe610b88565b01546040805160208101929092528101849052606001604051602081830303815290604052805190602001209250808061073790610be6565b9150506106b8565b50610748610e2e565b505050565b8060405161075a90610925565b6040518091039082f0905080158015610748573d6000803e3d6000fd5b6060600061078483610843565b600101905060008167ffffffffffffffff8111156107a4576107a4610a22565b6040519080825280601f01601f1916602001820160405280156107ce576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846107d857509392505050565b60f081901b82175b92915050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061088c577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef810000000083106108b8576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106108d657662386f26fc10000830492506010015b6305f5e10083106108ee576305f5e100830492506008015b612710831061090257612710830492506004015b60648310610914576064830492506002015b600a831061083d5760010192915050565b600880610e5e83390190565b600080600080610460858703121561094857600080fd5b8435935061042085018681111561095e57600080fd5b939660208601965093359461044001359392505050565b60005b83811015610990578181015183820152602001610978565b8381111561099f576000848401525b50505050565b600081518084526109bd816020860160208601610975565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610a0260208301846109a5565b9392505050565b600060208284031215610a1b57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080600060608486031215610a6657600080fd5b833573ffffffffffffffffffffffffffffffffffffffff81168114610a8a57600080fd5b925060208401359150604084013567ffffffffffffffff80821115610aae57600080fd5b818601915086601f830112610ac257600080fd5b813581811115610ad457610ad4610a22565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610b1a57610b1a610a22565b81604052828152896020848701011115610b3357600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b85815284602082015260a060408201526000610b7460a08301866109a5565b606083019490945250608001529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610c1757610c17610bb7565b5060010190565b60008451610c30818460208901610975565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610c6c816001850160208a01610975565b60019201918201528351610c87816002840160208801610975565b0160020195945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a0830152610cdf60c08301846109a5565b98975050505050505050565b600181815b80851115610d4457817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610d2a57610d2a610bb7565b80851615610d3757918102915b93841c9390800290610cf0565b509250929050565b600082610d5b5750600161083d565b81610d685750600061083d565b8160018114610d7e5760028114610d8857610da4565b600191505061083d565b60ff841115610d9957610d99610bb7565b50506001821b61083d565b5060208310610133831016604e8410600b8410161715610dc7575081810a61083d565b610dd18383610ceb565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610e0357610e03610bb7565b029392505050565b6000610a028383610d4c565b600082821015610e2957610e29610bb7565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fdfe608060405230fffea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2ToL1MessagePasserStorageLayoutJSON), L2ToL1MessagePasserStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2ToL1MessagePasser"] = L2ToL1MessagePasserStorageLayout
	deployedBytecodes["L2ToL1MessagePasser"] = L2ToL1MessagePasserDeployedBin
}
