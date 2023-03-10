// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fnft1155

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

// TokenFractional1155MetaData contains all meta data concerning the TokenFractional1155 contract.
var TokenFractional1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Compose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erctokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Fractional\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721tokenId\",\"type\":\"uint256\"}],\"name\":\"NFTOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"compose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"fractional\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040805160208101909152600081526200002c8162000033565b50620001b6565b6002620000418282620000ea565b5050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200007057607f821691505b6020821081036200009157634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620000e557600081815260208120601f850160051c81016020861015620000c05750805b601f850160051c820191505b81811015620000e157828155600101620000cc565b5050505b505050565b81516001600160401b0381111562000106576200010662000045565b6200011e816200011784546200005b565b8462000097565b602080601f8311600181146200015657600084156200013d5750858301515b600019600386901b1c1916600185901b178555620000e1565b600085815260208120601f198616915b82811015620001875788860151825594840194600190910190840162000166565b5085821015620001a65787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b611d6d80620001c66000396000f3fe608060405234801561001057600080fd5b50600436106100ce5760003560e01c806367c8eb101161008c578063bd85b03911610066578063bd85b039146101bd578063e985e9c5146101d0578063ea78803f1461020c578063f242432a1461023e57600080fd5b806367c8eb1014610184578063a22cb46514610197578063aa3552e3146101aa57600080fd5b8062fdd58e146100d357806301ffc9a7146100f95780630e89341c1461011c5780632eb2c2d61461013c5780634e1273f4146101515780635af7714014610171575b600080fd5b6100e66100e136600461146a565b610251565b6040519081526020015b60405180910390f35b61010c6101073660046114ac565b6102ea565b60405190151581526020016100f0565b61012f61012a3660046114d0565b61033a565b6040516100f0919061152f565b61014f61014a36600461168e565b6103ce565b005b61016461015f36600461173c565b61041a565b6040516100f09190611844565b61014f61017f36600461146a565b610544565b6100e661019236600461146a565b61079b565b61014f6101a5366004611857565b610828565b61014f6101b8366004611895565b610837565b6100e66101cb3660046114d0565b610a9b565b61010c6101de3660046118f8565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205460ff1690565b61021f61021a3660046114d0565b610adf565b604080516001600160a01b0390931683526020830191909152016100f0565b61014f61024c366004611926565b610b39565b60006001600160a01b0383166102c15760405162461bcd60e51b815260206004820152602a60248201527f455243313135353a2061646472657373207a65726f206973206e6f742061207660448201526930b634b21037bbb732b960b11b60648201526084015b60405180910390fd5b506000818152602081815260408083206001600160a01b03861684529091529020545b92915050565b60006001600160e01b03198216636cdb3d1360e11b148061031b57506001600160e01b031982166303a24d0760e21b145b806102e457506301ffc9a760e01b6001600160e01b03198316146102e4565b6060600280546103499061198f565b80601f01602080910402602001604051908101604052809291908181526020018280546103759061198f565b80156103c25780601f10610397576101008083540402835291602001916103c2565b820191906000526020600020905b8154815290600101906020018083116103a557829003601f168201915b50505050509050919050565b6001600160a01b0385163314806103ea57506103ea85336101de565b6104065760405162461bcd60e51b81526004016102b8906119c9565b6104138585858585610b7e565b5050505050565b6060815183511461047f5760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b60648201526084016102b8565b6000835167ffffffffffffffff81111561049b5761049b611542565b6040519080825280602002602001820160405280156104c4578160200160208202803683370190505b50905060005b845181101561053c5761050f8582815181106104e8576104e8611a17565b602002602001015185838151811061050257610502611a17565b6020026020010151610251565b82828151811061052157610521611a17565b602090810291909101015261053581611a43565b90506104ca565b509392505050565b6000818152600460205260409020600201546105725760405162461bcd60e51b81526004016102b890611a5c565b60003360405163bd85b03960e01b815260048101849052909150600090309063bd85b03990602401602060405180830381865afa1580156105b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105db9190611a83565b604051627eeac760e11b81526001600160a01b0384166004820152602481018590529091508190309062fdd58e90604401602060405180830381865afa158015610629573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064d9190611a83565b1461069a5760405162461bcd60e51b815260206004820152601960248201527f4e6f74206f776e696e6720616c6c2074686520746f6b656e730000000000000060448201526064016102b8565b600083815260046020819052604091829020805460019091015492516323b872dd60e01b815230928101929092526001600160a01b03878116602484015260448301939093526101009004909116906323b872dd90606401600060405180830381600087803b15801561070c57600080fd5b505af1158015610720573d6000803e3d6000fd5b5050505061072f828483610d5b565b600083815260046020526040808220805460ff1916815560020191909155516001600160a01b0380861691908416907fb36eff7b5633a7f06feb2b1be43ab1fb34648ea59b5d5bd0303350c4c899f8349061078d9087815260200190565b60405180910390a350505050565b6001600160a01b038216600090815260056020908152604080832084845290915281205460ff166107fc5760405162461bcd60e51b815260206004820152600b60248201526a125b9d985b1a590813919560aa1b60448201526064016102b8565b506001600160a01b03919091166000908152600560209081526040808320938352929052206001015490565b610833338383610edc565b5050565b600082116108875760405162461bcd60e51b815260206004820152601d60248201527f616d6f756e74206d7573742067726561746572207468616e207a65726f00000060448201526064016102b8565b600061089260035490565b90506108a2600380546001019055565b846001600160a01b0381166323b872dd336040516001600160e01b031960e084901b1681526001600160a01b03909116600482015230602482015260448101889052606401600060405180830381600087803b15801561090157600080fd5b505af1158015610915573d6000803e3d6000fd5b50506040516331a9108f60e11b8152600481018890523092506001600160a01b0384169150636352211e90602401602060405180830381865afa158015610960573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109849190611a9c565b6001600160a01b0316146109c95760405162461bcd60e51b815260206004820152600c60248201526b18da1958dac819985a5b195960a21b60448201526064016102b8565b6000828152600460209081526040808320805460018083018b905560ff196001600160a01b038881166101000282166001600160a81b0319909416939093178217845560029093018a9055908b168552600584528285208a865290935290832080549091168217815501839055339050610a4581848787610fbc565b6040805187815260208101859052808201879052905133916001600160a01b038a16917f6e74140de0afe89e44cc22522a7dc848c82c4dca3923a26e2637dcc7c6b30bc59181900360600190a350505050505050565b60008181526004602052604081205460ff16610ac95760405162461bcd60e51b81526004016102b890611a5c565b5060009081526004602052604090206002015490565b600081815260046020526040812054819060ff16610b0f5760405162461bcd60e51b81526004016102b890611a5c565b5050600090815260046020526040902080546001909101546101009091046001600160a01b031691565b6001600160a01b038516331480610b555750610b5585336101de565b610b715760405162461bcd60e51b81526004016102b8906119c9565b61041385858585856110c7565b8151835114610be05760405162461bcd60e51b815260206004820152602860248201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206044820152670dad2e6dac2e8c6d60c31b60648201526084016102b8565b6001600160a01b038416610c065760405162461bcd60e51b81526004016102b890611ab9565b3360005b8451811015610ced576000858281518110610c2757610c27611a17565b602002602001015190506000858381518110610c4557610c45611a17565b602090810291909101810151600084815280835260408082206001600160a01b038e168352909352919091205490915081811015610c955760405162461bcd60e51b81526004016102b890611afe565b6000838152602081815260408083206001600160a01b038e8116855292528083208585039055908b16825281208054849290610cd2908490611b48565b9250508190555050505080610ce690611a43565b9050610c0a565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051610d3d929190611b5b565b60405180910390a4610d538187878787876111f1565b505050505050565b6001600160a01b038316610dbd5760405162461bcd60e51b815260206004820152602360248201527f455243313135353a206275726e2066726f6d20746865207a65726f206164647260448201526265737360e81b60648201526084016102b8565b336000610dc98461134c565b90506000610dd68461134c565b60408051602080820183526000918290528882528181528282206001600160a01b038b1683529052205490915084811015610e5f5760405162461bcd60e51b8152602060048201526024808201527f455243313135353a206275726e20616d6f756e7420657863656564732062616c604482015263616e636560e01b60648201526084016102b8565b6000868152602081815260408083206001600160a01b038b81168086529184528285208a8703905582518b81529384018a90529092908816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a46040805160208101909152600090525b50505050505050565b816001600160a01b0316836001600160a01b031603610f4f5760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b60648201526084016102b8565b6001600160a01b03838116600081815260016020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b03841661101c5760405162461bcd60e51b815260206004820152602160248201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736044820152607360f81b60648201526084016102b8565b3360006110288561134c565b905060006110358561134c565b90506000868152602081815260408083206001600160a01b038b16845290915281208054879290611067908490611b48565b909155505060408051878152602081018790526001600160a01b03808a1692600092918716917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4610ed383600089898989611397565b6001600160a01b0384166110ed5760405162461bcd60e51b81526004016102b890611ab9565b3360006110f98561134c565b905060006111068561134c565b90506000868152602081815260408083206001600160a01b038c168452909152902054858110156111495760405162461bcd60e51b81526004016102b890611afe565b6000878152602081815260408083206001600160a01b038d8116855292528083208985039055908a16825281208054889290611186908490611b48565b909155505060408051888152602081018890526001600160a01b03808b16928c821692918816917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a46111e6848a8a8a8a8a611397565b505050505050505050565b6001600160a01b0384163b15610d535760405163bc197c8160e01b81526001600160a01b0385169063bc197c81906112359089908990889088908890600401611b89565b6020604051808303816000875af1925050508015611270575060408051601f3d908101601f1916820190925261126d91810190611be7565b60015b61131c5761127c611c04565b806308c379a0036112b55750611290611c20565b8061129b57506112b7565b8060405162461bcd60e51b81526004016102b8919061152f565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e2d455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b60648201526084016102b8565b6001600160e01b0319811663bc197c8160e01b14610ed35760405162461bcd60e51b81526004016102b890611caa565b6040805160018082528183019092526060916000919060208083019080368337019050509050828160008151811061138657611386611a17565b602090810291909101015292915050565b6001600160a01b0384163b15610d535760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e61906113db9089908990889088908890600401611cf2565b6020604051808303816000875af1925050508015611416575060408051601f3d908101601f1916820190925261141391810190611be7565b60015b6114225761127c611c04565b6001600160e01b0319811663f23a6e6160e01b14610ed35760405162461bcd60e51b81526004016102b890611caa565b6001600160a01b038116811461146757600080fd5b50565b6000806040838503121561147d57600080fd5b823561148881611452565b946020939093013593505050565b6001600160e01b03198116811461146757600080fd5b6000602082840312156114be57600080fd5b81356114c981611496565b9392505050565b6000602082840312156114e257600080fd5b5035919050565b6000815180845260005b8181101561150f576020818501810151868301820152016114f3565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006114c960208301846114e9565b634e487b7160e01b600052604160045260246000fd5b601f8201601f1916810167ffffffffffffffff8111828210171561157e5761157e611542565b6040525050565b600067ffffffffffffffff82111561159f5761159f611542565b5060051b60200190565b600082601f8301126115ba57600080fd5b813560206115c782611585565b6040516115d48282611558565b83815260059390931b85018201928281019150868411156115f457600080fd5b8286015b8481101561160f57803583529183019183016115f8565b509695505050505050565b600082601f83011261162b57600080fd5b813567ffffffffffffffff81111561164557611645611542565b60405161165c601f8301601f191660200182611558565b81815284602083860101111561167157600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600060a086880312156116a657600080fd5b85356116b181611452565b945060208601356116c181611452565b9350604086013567ffffffffffffffff808211156116de57600080fd5b6116ea89838a016115a9565b9450606088013591508082111561170057600080fd5b61170c89838a016115a9565b9350608088013591508082111561172257600080fd5b5061172f8882890161161a565b9150509295509295909350565b6000806040838503121561174f57600080fd5b823567ffffffffffffffff8082111561176757600080fd5b818501915085601f83011261177b57600080fd5b8135602061178882611585565b6040516117958282611558565b83815260059390931b85018201928281019150898411156117b557600080fd5b948201945b838610156117dc5785356117cd81611452565b825294820194908201906117ba565b965050860135925050808211156117f257600080fd5b506117ff858286016115a9565b9150509250929050565b600081518084526020808501945080840160005b838110156118395781518752958201959082019060010161181d565b509495945050505050565b6020815260006114c96020830184611809565b6000806040838503121561186a57600080fd5b823561187581611452565b91506020830135801515811461188a57600080fd5b809150509250929050565b600080600080608085870312156118ab57600080fd5b84356118b681611452565b93506020850135925060408501359150606085013567ffffffffffffffff8111156118e057600080fd5b6118ec8782880161161a565b91505092959194509250565b6000806040838503121561190b57600080fd5b823561191681611452565b9150602083013561188a81611452565b600080600080600060a0868803121561193e57600080fd5b853561194981611452565b9450602086013561195981611452565b93506040860135925060608601359150608086013567ffffffffffffffff81111561198357600080fd5b61172f8882890161161a565b600181811c908216806119a357607f821691505b6020821081036119c357634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602e908201527f455243313135353a2063616c6c6572206973206e6f7420746f6b656e206f776e60408201526d195c881bdc88185c1c1c9bdd995960921b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201611a5557611a55611a2d565b5060010190565b6020808252600d908201526c24b73b30b634b2103a37b5b2b760991b604082015260600190565b600060208284031215611a9557600080fd5b5051919050565b600060208284031215611aae57600080fd5b81516114c981611452565b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b808201808211156102e4576102e4611a2d565b604081526000611b6e6040830185611809565b8281036020840152611b808185611809565b95945050505050565b6001600160a01b0386811682528516602082015260a060408201819052600090611bb590830186611809565b8281036060840152611bc78186611809565b90508281036080840152611bdb81856114e9565b98975050505050505050565b600060208284031215611bf957600080fd5b81516114c981611496565b600060033d1115611c1d5760046000803e5060005160e01c5b90565b600060443d1015611c2e5790565b6040516003193d81016004833e81513d67ffffffffffffffff8160248401118184111715611c5e57505050505090565b8285019150815181811115611c765750505050505090565b843d8701016020828501011115611c905750505050505090565b611c9f60208286010187611558565b509095945050505050565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b6001600160a01b03868116825285166020820152604081018490526060810183905260a060808201819052600090611d2c908301846114e9565b97965050505050505056fea2646970667358221220bb09c6e1a5599b5ae2f19697545fc91327b9fc4c62807c4e8aec7b29119c6c1164736f6c63430008110033",
}

// TokenFractional1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenFractional1155MetaData.ABI instead.
var TokenFractional1155ABI = TokenFractional1155MetaData.ABI

// TokenFractional1155Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenFractional1155MetaData.Bin instead.
var TokenFractional1155Bin = TokenFractional1155MetaData.Bin

// DeployTokenFractional1155 deploys a new Ethereum contract, binding an instance of TokenFractional1155 to it.
func DeployTokenFractional1155(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenFractional1155, error) {
	parsed, err := TokenFractional1155MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenFractional1155Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenFractional1155{TokenFractional1155Caller: TokenFractional1155Caller{contract: contract}, TokenFractional1155Transactor: TokenFractional1155Transactor{contract: contract}, TokenFractional1155Filterer: TokenFractional1155Filterer{contract: contract}}, nil
}

// TokenFractional1155 is an auto generated Go binding around an Ethereum contract.
type TokenFractional1155 struct {
	TokenFractional1155Caller     // Read-only binding to the contract
	TokenFractional1155Transactor // Write-only binding to the contract
	TokenFractional1155Filterer   // Log filterer for contract events
}

// TokenFractional1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type TokenFractional1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFractional1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenFractional1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFractional1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFractional1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFractional1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenFractional1155Session struct {
	Contract     *TokenFractional1155 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TokenFractional1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenFractional1155CallerSession struct {
	Contract *TokenFractional1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TokenFractional1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenFractional1155TransactorSession struct {
	Contract     *TokenFractional1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TokenFractional1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type TokenFractional1155Raw struct {
	Contract *TokenFractional1155 // Generic contract binding to access the raw methods on
}

// TokenFractional1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenFractional1155CallerRaw struct {
	Contract *TokenFractional1155Caller // Generic read-only contract binding to access the raw methods on
}

// TokenFractional1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenFractional1155TransactorRaw struct {
	Contract *TokenFractional1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenFractional1155 creates a new instance of TokenFractional1155, bound to a specific deployed contract.
func NewTokenFractional1155(address common.Address, backend bind.ContractBackend) (*TokenFractional1155, error) {
	contract, err := bindTokenFractional1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenFractional1155{TokenFractional1155Caller: TokenFractional1155Caller{contract: contract}, TokenFractional1155Transactor: TokenFractional1155Transactor{contract: contract}, TokenFractional1155Filterer: TokenFractional1155Filterer{contract: contract}}, nil
}

// NewTokenFractional1155Caller creates a new read-only instance of TokenFractional1155, bound to a specific deployed contract.
func NewTokenFractional1155Caller(address common.Address, caller bind.ContractCaller) (*TokenFractional1155Caller, error) {
	contract, err := bindTokenFractional1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFractional1155Caller{contract: contract}, nil
}

// NewTokenFractional1155Transactor creates a new write-only instance of TokenFractional1155, bound to a specific deployed contract.
func NewTokenFractional1155Transactor(address common.Address, transactor bind.ContractTransactor) (*TokenFractional1155Transactor, error) {
	contract, err := bindTokenFractional1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFractional1155Transactor{contract: contract}, nil
}

// NewTokenFractional1155Filterer creates a new log filterer instance of TokenFractional1155, bound to a specific deployed contract.
func NewTokenFractional1155Filterer(address common.Address, filterer bind.ContractFilterer) (*TokenFractional1155Filterer, error) {
	contract, err := bindTokenFractional1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFractional1155Filterer{contract: contract}, nil
}

// bindTokenFractional1155 binds a generic wrapper to an already deployed contract.
func bindTokenFractional1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenFractional1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFractional1155 *TokenFractional1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenFractional1155.Contract.TokenFractional1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFractional1155 *TokenFractional1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFractional1155.Contract.TokenFractional1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFractional1155 *TokenFractional1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFractional1155.Contract.TokenFractional1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFractional1155 *TokenFractional1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenFractional1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFractional1155 *TokenFractional1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFractional1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFractional1155 *TokenFractional1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFractional1155.Contract.contract.Transact(opts, method, params...)
}

// NFTOf is a free data retrieval call binding the contract method 0x67c8eb10.
//
// Solidity: function NFTOf(address erc721address, uint256 erc721tokenId) view returns(uint256)
func (_TokenFractional1155 *TokenFractional1155Caller) NFTOf(opts *bind.CallOpts, erc721address common.Address, erc721tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenFractional1155.contract.Call(opts, &out, "NFTOf", erc721address, erc721tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NFTOf is a free data retrieval call binding the contract method 0x67c8eb10.
//
// Solidity: function NFTOf(address erc721address, uint256 erc721tokenId) view returns(uint256)
func (_TokenFractional1155 *TokenFractional1155Session) NFTOf(erc721address common.Address, erc721tokenId *big.Int) (*big.Int, error) {
	return _TokenFractional1155.Contract.NFTOf(&_TokenFractional1155.CallOpts, erc721address, erc721tokenId)
}

// NFTOf is a free data retrieval call binding the contract method 0x67c8eb10.
//
// Solidity: function NFTOf(address erc721address, uint256 erc721tokenId) view returns(uint256)
func (_TokenFractional1155 *TokenFractional1155CallerSession) NFTOf(erc721address common.Address, erc721tokenId *big.Int) (*big.Int, error) {
	return _TokenFractional1155.Contract.NFTOf(&_TokenFractional1155.CallOpts, erc721address, erc721tokenId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_TokenFractional1155 *TokenFractional1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenFractional1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_TokenFractional1155 *TokenFractional1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _TokenFractional1155.Contract.BalanceOf(&_TokenFractional1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_TokenFractional1155 *TokenFractional1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _TokenFractional1155.Contract.BalanceOf(&_TokenFractional1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_TokenFractional1155 *TokenFractional1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _TokenFractional1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_TokenFractional1155 *TokenFractional1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _TokenFractional1155.Contract.BalanceOfBatch(&_TokenFractional1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_TokenFractional1155 *TokenFractional1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _TokenFractional1155.Contract.BalanceOfBatch(&_TokenFractional1155.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_TokenFractional1155 *TokenFractional1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TokenFractional1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_TokenFractional1155 *TokenFractional1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _TokenFractional1155.Contract.IsApprovedForAll(&_TokenFractional1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_TokenFractional1155 *TokenFractional1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _TokenFractional1155.Contract.IsApprovedForAll(&_TokenFractional1155.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenFractional1155 *TokenFractional1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TokenFractional1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenFractional1155 *TokenFractional1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenFractional1155.Contract.SupportsInterface(&_TokenFractional1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenFractional1155 *TokenFractional1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenFractional1155.Contract.SupportsInterface(&_TokenFractional1155.CallOpts, interfaceId)
}

// TokenOf is a free data retrieval call binding the contract method 0xea78803f.
//
// Solidity: function tokenOf(uint256 tokenId) view returns(address, uint256)
func (_TokenFractional1155 *TokenFractional1155Caller) TokenOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _TokenFractional1155.contract.Call(opts, &out, "tokenOf", tokenId)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// TokenOf is a free data retrieval call binding the contract method 0xea78803f.
//
// Solidity: function tokenOf(uint256 tokenId) view returns(address, uint256)
func (_TokenFractional1155 *TokenFractional1155Session) TokenOf(tokenId *big.Int) (common.Address, *big.Int, error) {
	return _TokenFractional1155.Contract.TokenOf(&_TokenFractional1155.CallOpts, tokenId)
}

// TokenOf is a free data retrieval call binding the contract method 0xea78803f.
//
// Solidity: function tokenOf(uint256 tokenId) view returns(address, uint256)
func (_TokenFractional1155 *TokenFractional1155CallerSession) TokenOf(tokenId *big.Int) (common.Address, *big.Int, error) {
	return _TokenFractional1155.Contract.TokenOf(&_TokenFractional1155.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_TokenFractional1155 *TokenFractional1155Caller) TotalSupply(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenFractional1155.contract.Call(opts, &out, "totalSupply", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_TokenFractional1155 *TokenFractional1155Session) TotalSupply(tokenId *big.Int) (*big.Int, error) {
	return _TokenFractional1155.Contract.TotalSupply(&_TokenFractional1155.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_TokenFractional1155 *TokenFractional1155CallerSession) TotalSupply(tokenId *big.Int) (*big.Int, error) {
	return _TokenFractional1155.Contract.TotalSupply(&_TokenFractional1155.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_TokenFractional1155 *TokenFractional1155Caller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _TokenFractional1155.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_TokenFractional1155 *TokenFractional1155Session) Uri(arg0 *big.Int) (string, error) {
	return _TokenFractional1155.Contract.Uri(&_TokenFractional1155.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_TokenFractional1155 *TokenFractional1155CallerSession) Uri(arg0 *big.Int) (string, error) {
	return _TokenFractional1155.Contract.Uri(&_TokenFractional1155.CallOpts, arg0)
}

// Compose is a paid mutator transaction binding the contract method 0x5af77140.
//
// Solidity: function compose(address to, uint256 tokenId) returns()
func (_TokenFractional1155 *TokenFractional1155Transactor) Compose(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenFractional1155.contract.Transact(opts, "compose", to, tokenId)
}

// Compose is a paid mutator transaction binding the contract method 0x5af77140.
//
// Solidity: function compose(address to, uint256 tokenId) returns()
func (_TokenFractional1155 *TokenFractional1155Session) Compose(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenFractional1155.Contract.Compose(&_TokenFractional1155.TransactOpts, to, tokenId)
}

// Compose is a paid mutator transaction binding the contract method 0x5af77140.
//
// Solidity: function compose(address to, uint256 tokenId) returns()
func (_TokenFractional1155 *TokenFractional1155TransactorSession) Compose(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenFractional1155.Contract.Compose(&_TokenFractional1155.TransactOpts, to, tokenId)
}

// Fractional is a paid mutator transaction binding the contract method 0xaa3552e3.
//
// Solidity: function fractional(address erc721address, uint256 erc721tokenId, uint256 amount, bytes data) returns()
func (_TokenFractional1155 *TokenFractional1155Transactor) Fractional(opts *bind.TransactOpts, erc721address common.Address, erc721tokenId *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenFractional1155.contract.Transact(opts, "fractional", erc721address, erc721tokenId, amount, data)
}

// Fractional is a paid mutator transaction binding the contract method 0xaa3552e3.
//
// Solidity: function fractional(address erc721address, uint256 erc721tokenId, uint256 amount, bytes data) returns()
func (_TokenFractional1155 *TokenFractional1155Session) Fractional(erc721address common.Address, erc721tokenId *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenFractional1155.Contract.Fractional(&_TokenFractional1155.TransactOpts, erc721address, erc721tokenId, amount, data)
}

// Fractional is a paid mutator transaction binding the contract method 0xaa3552e3.
//
// Solidity: function fractional(address erc721address, uint256 erc721tokenId, uint256 amount, bytes data) returns()
func (_TokenFractional1155 *TokenFractional1155TransactorSession) Fractional(erc721address common.Address, erc721tokenId *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenFractional1155.Contract.Fractional(&_TokenFractional1155.TransactOpts, erc721address, erc721tokenId, amount, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_TokenFractional1155 *TokenFractional1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _TokenFractional1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_TokenFractional1155 *TokenFractional1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _TokenFractional1155.Contract.SafeBatchTransferFrom(&_TokenFractional1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_TokenFractional1155 *TokenFractional1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _TokenFractional1155.Contract.SafeBatchTransferFrom(&_TokenFractional1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_TokenFractional1155 *TokenFractional1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenFractional1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_TokenFractional1155 *TokenFractional1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenFractional1155.Contract.SafeTransferFrom(&_TokenFractional1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_TokenFractional1155 *TokenFractional1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenFractional1155.Contract.SafeTransferFrom(&_TokenFractional1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TokenFractional1155 *TokenFractional1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _TokenFractional1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TokenFractional1155 *TokenFractional1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TokenFractional1155.Contract.SetApprovalForAll(&_TokenFractional1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TokenFractional1155 *TokenFractional1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TokenFractional1155.Contract.SetApprovalForAll(&_TokenFractional1155.TransactOpts, operator, approved)
}

// TokenFractional1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TokenFractional1155 contract.
type TokenFractional1155ApprovalForAllIterator struct {
	Event *TokenFractional1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TokenFractional1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFractional1155ApprovalForAll)
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
		it.Event = new(TokenFractional1155ApprovalForAll)
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
func (it *TokenFractional1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFractional1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFractional1155ApprovalForAll represents a ApprovalForAll event raised by the TokenFractional1155 contract.
type TokenFractional1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_TokenFractional1155 *TokenFractional1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*TokenFractional1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TokenFractional1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TokenFractional1155ApprovalForAllIterator{contract: _TokenFractional1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_TokenFractional1155 *TokenFractional1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TokenFractional1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TokenFractional1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFractional1155ApprovalForAll)
				if err := _TokenFractional1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_TokenFractional1155 *TokenFractional1155Filterer) ParseApprovalForAll(log types.Log) (*TokenFractional1155ApprovalForAll, error) {
	event := new(TokenFractional1155ApprovalForAll)
	if err := _TokenFractional1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFractional1155ComposeIterator is returned from FilterCompose and is used to iterate over the raw logs and unpacked data for Compose events raised by the TokenFractional1155 contract.
type TokenFractional1155ComposeIterator struct {
	Event *TokenFractional1155Compose // Event containing the contract specifics and raw log

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
func (it *TokenFractional1155ComposeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFractional1155Compose)
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
		it.Event = new(TokenFractional1155Compose)
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
func (it *TokenFractional1155ComposeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFractional1155ComposeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFractional1155Compose represents a Compose event raised by the TokenFractional1155 contract.
type TokenFractional1155Compose struct {
	Sender  common.Address
	TokenId *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCompose is a free log retrieval operation binding the contract event 0xb36eff7b5633a7f06feb2b1be43ab1fb34648ea59b5d5bd0303350c4c899f834.
//
// Solidity: event Compose(address indexed sender, uint256 tokenId, address indexed to)
func (_TokenFractional1155 *TokenFractional1155Filterer) FilterCompose(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*TokenFractional1155ComposeIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenFractional1155.contract.FilterLogs(opts, "Compose", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenFractional1155ComposeIterator{contract: _TokenFractional1155.contract, event: "Compose", logs: logs, sub: sub}, nil
}

// WatchCompose is a free log subscription operation binding the contract event 0xb36eff7b5633a7f06feb2b1be43ab1fb34648ea59b5d5bd0303350c4c899f834.
//
// Solidity: event Compose(address indexed sender, uint256 tokenId, address indexed to)
func (_TokenFractional1155 *TokenFractional1155Filterer) WatchCompose(opts *bind.WatchOpts, sink chan<- *TokenFractional1155Compose, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenFractional1155.contract.WatchLogs(opts, "Compose", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFractional1155Compose)
				if err := _TokenFractional1155.contract.UnpackLog(event, "Compose", log); err != nil {
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

// ParseCompose is a log parse operation binding the contract event 0xb36eff7b5633a7f06feb2b1be43ab1fb34648ea59b5d5bd0303350c4c899f834.
//
// Solidity: event Compose(address indexed sender, uint256 tokenId, address indexed to)
func (_TokenFractional1155 *TokenFractional1155Filterer) ParseCompose(log types.Log) (*TokenFractional1155Compose, error) {
	event := new(TokenFractional1155Compose)
	if err := _TokenFractional1155.contract.UnpackLog(event, "Compose", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFractional1155FractionalIterator is returned from FilterFractional and is used to iterate over the raw logs and unpacked data for Fractional events raised by the TokenFractional1155 contract.
type TokenFractional1155FractionalIterator struct {
	Event *TokenFractional1155Fractional // Event containing the contract specifics and raw log

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
func (it *TokenFractional1155FractionalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFractional1155Fractional)
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
		it.Event = new(TokenFractional1155Fractional)
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
func (it *TokenFractional1155FractionalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFractional1155FractionalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFractional1155Fractional represents a Fractional event raised by the TokenFractional1155 contract.
type TokenFractional1155Fractional struct {
	Erc721     common.Address
	Spender    common.Address
	ErctokenId *big.Int
	TokenId    *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFractional is a free log retrieval operation binding the contract event 0x6e74140de0afe89e44cc22522a7dc848c82c4dca3923a26e2637dcc7c6b30bc5.
//
// Solidity: event Fractional(address indexed erc721, address indexed spender, uint256 erctokenId, uint256 tokenId, uint256 amount)
func (_TokenFractional1155 *TokenFractional1155Filterer) FilterFractional(opts *bind.FilterOpts, erc721 []common.Address, spender []common.Address) (*TokenFractional1155FractionalIterator, error) {

	var erc721Rule []interface{}
	for _, erc721Item := range erc721 {
		erc721Rule = append(erc721Rule, erc721Item)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenFractional1155.contract.FilterLogs(opts, "Fractional", erc721Rule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenFractional1155FractionalIterator{contract: _TokenFractional1155.contract, event: "Fractional", logs: logs, sub: sub}, nil
}

// WatchFractional is a free log subscription operation binding the contract event 0x6e74140de0afe89e44cc22522a7dc848c82c4dca3923a26e2637dcc7c6b30bc5.
//
// Solidity: event Fractional(address indexed erc721, address indexed spender, uint256 erctokenId, uint256 tokenId, uint256 amount)
func (_TokenFractional1155 *TokenFractional1155Filterer) WatchFractional(opts *bind.WatchOpts, sink chan<- *TokenFractional1155Fractional, erc721 []common.Address, spender []common.Address) (event.Subscription, error) {

	var erc721Rule []interface{}
	for _, erc721Item := range erc721 {
		erc721Rule = append(erc721Rule, erc721Item)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenFractional1155.contract.WatchLogs(opts, "Fractional", erc721Rule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFractional1155Fractional)
				if err := _TokenFractional1155.contract.UnpackLog(event, "Fractional", log); err != nil {
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

// ParseFractional is a log parse operation binding the contract event 0x6e74140de0afe89e44cc22522a7dc848c82c4dca3923a26e2637dcc7c6b30bc5.
//
// Solidity: event Fractional(address indexed erc721, address indexed spender, uint256 erctokenId, uint256 tokenId, uint256 amount)
func (_TokenFractional1155 *TokenFractional1155Filterer) ParseFractional(log types.Log) (*TokenFractional1155Fractional, error) {
	event := new(TokenFractional1155Fractional)
	if err := _TokenFractional1155.contract.UnpackLog(event, "Fractional", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFractional1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the TokenFractional1155 contract.
type TokenFractional1155TransferBatchIterator struct {
	Event *TokenFractional1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *TokenFractional1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFractional1155TransferBatch)
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
		it.Event = new(TokenFractional1155TransferBatch)
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
func (it *TokenFractional1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFractional1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFractional1155TransferBatch represents a TransferBatch event raised by the TokenFractional1155 contract.
type TokenFractional1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_TokenFractional1155 *TokenFractional1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*TokenFractional1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenFractional1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenFractional1155TransferBatchIterator{contract: _TokenFractional1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_TokenFractional1155 *TokenFractional1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *TokenFractional1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenFractional1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFractional1155TransferBatch)
				if err := _TokenFractional1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_TokenFractional1155 *TokenFractional1155Filterer) ParseTransferBatch(log types.Log) (*TokenFractional1155TransferBatch, error) {
	event := new(TokenFractional1155TransferBatch)
	if err := _TokenFractional1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFractional1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the TokenFractional1155 contract.
type TokenFractional1155TransferSingleIterator struct {
	Event *TokenFractional1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *TokenFractional1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFractional1155TransferSingle)
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
		it.Event = new(TokenFractional1155TransferSingle)
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
func (it *TokenFractional1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFractional1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFractional1155TransferSingle represents a TransferSingle event raised by the TokenFractional1155 contract.
type TokenFractional1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_TokenFractional1155 *TokenFractional1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*TokenFractional1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenFractional1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenFractional1155TransferSingleIterator{contract: _TokenFractional1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_TokenFractional1155 *TokenFractional1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *TokenFractional1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenFractional1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFractional1155TransferSingle)
				if err := _TokenFractional1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_TokenFractional1155 *TokenFractional1155Filterer) ParseTransferSingle(log types.Log) (*TokenFractional1155TransferSingle, error) {
	event := new(TokenFractional1155TransferSingle)
	if err := _TokenFractional1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFractional1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the TokenFractional1155 contract.
type TokenFractional1155URIIterator struct {
	Event *TokenFractional1155URI // Event containing the contract specifics and raw log

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
func (it *TokenFractional1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFractional1155URI)
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
		it.Event = new(TokenFractional1155URI)
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
func (it *TokenFractional1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFractional1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFractional1155URI represents a URI event raised by the TokenFractional1155 contract.
type TokenFractional1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_TokenFractional1155 *TokenFractional1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*TokenFractional1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TokenFractional1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &TokenFractional1155URIIterator{contract: _TokenFractional1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_TokenFractional1155 *TokenFractional1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *TokenFractional1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TokenFractional1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFractional1155URI)
				if err := _TokenFractional1155.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_TokenFractional1155 *TokenFractional1155Filterer) ParseURI(log types.Log) (*TokenFractional1155URI, error) {
	event := new(TokenFractional1155URI)
	if err := _TokenFractional1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
