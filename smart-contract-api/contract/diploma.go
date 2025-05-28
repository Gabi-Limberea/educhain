// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "bce36fc6a91bee4c5901573e6ca32329b7",
}

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	abi abi.ABI
}

// NewContext creates a new instance of Context.
func NewContext() *Context {
	parsed, err := ContextMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Context{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Context) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// DiplomaMetaData contains all meta data concerning the Diploma contract.
var DiplomaMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"safeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "31f447555d8fd40d52220ea41050ccd9e5",
	Bin: "0x60806040525f600755348015610013575f80fd5b506040518060400160405280600781526020017f4469706c6f6d61000000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f44504c0000000000000000000000000000000000000000000000000000000000815250815f908161008e9190610320565b50806001908161009e9190610320565b5050503360085f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506103ef565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061016157607f821691505b6020821081036101745761017361011d565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101d67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261019b565b6101e0868361019b565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61022461021f61021a846101f8565b610201565b6101f8565b9050919050565b5f819050919050565b61023d8361020a565b6102516102498261022b565b8484546101a7565b825550505050565b5f90565b610265610259565b610270818484610234565b505050565b5b81811015610293576102885f8261025d565b600181019050610276565b5050565b601f8211156102d8576102a98161017a565b6102b28461018c565b810160208510156102c1578190505b6102d56102cd8561018c565b830182610275565b50505b505050565b5f82821c905092915050565b5f6102f85f19846008026102dd565b1980831691505092915050565b5f61031083836102e9565b9150826002028217905092915050565b610329826100e6565b67ffffffffffffffff811115610342576103416100f0565b5b61034c825461014a565b610357828285610297565b5f60209050601f831160018114610388575f8415610376578287015190505b6103808582610305565b8655506103e7565b601f1984166103968661017a565b5f5b828110156103bd57848901518255600182019150602085019450602081019050610398565b868310156103da57848901516103d6601f8916826102e9565b8355505b6001600288020188555050505b505050505050565b612644806103fc5f395ff3fe608060405234801561000f575f80fd5b5060043610610109575f3560e01c806370a08231116100a0578063c87b56dd1161006f578063c87b56dd146102b1578063d082e381146102e1578063d204c45e146102ff578063e985e9c51461031b578063f851a4401461034b57610109565b806370a082311461022b57806395d89b411461025b578063a22cb46514610279578063b88d4fde1461029557610109565b806323b872dd116100dc57806323b872dd146101a757806342842e0e146101c357806342966c68146101df5780636352211e146101fb57610109565b806301ffc9a71461010d57806306fdde031461013d578063081812fc1461015b578063095ea7b31461018b575b5f80fd5b61012760048036038101906101229190611a49565b610369565b6040516101349190611a8e565b60405180910390f35b61014561037a565b6040516101529190611b31565b60405180910390f35b61017560048036038101906101709190611b84565b610409565b6040516101829190611bee565b60405180910390f35b6101a560048036038101906101a09190611c31565b610424565b005b6101c160048036038101906101bc9190611c6f565b61043a565b005b6101dd60048036038101906101d89190611c6f565b610539565b005b6101f960048036038101906101f49190611b84565b610558565b005b61021560048036038101906102109190611b84565b61056e565b6040516102229190611bee565b60405180910390f35b61024560048036038101906102409190611cbf565b61057f565b6040516102529190611cf9565b60405180910390f35b610263610635565b6040516102709190611b31565b60405180910390f35b610293600480360381019061028e9190611d3c565b6106c5565b005b6102af60048036038101906102aa9190611ea6565b6106db565b005b6102cb60048036038101906102c69190611b84565b610700565b6040516102d89190611b31565b60405180910390f35b6102e9610712565b6040516102f69190611cf9565b60405180910390f35b61031960048036038101906103149190611fc4565b610718565b005b6103356004803603810190610330919061201e565b610840565b6040516103429190611a8e565b60405180910390f35b6103536108ce565b6040516103609190611bee565b60405180910390f35b5f610373826108f3565b9050919050565b60605f805461038890612089565b80601f01602080910402602001604051908101604052809291908181526020018280546103b490612089565b80156103ff5780601f106103d6576101008083540402835291602001916103ff565b820191905f5260205f20905b8154815290600101906020018083116103e257829003601f168201915b5050505050905090565b5f61041382610953565b5061041d826109d9565b9050919050565b6104368282610431610a12565b610a19565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036104aa575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016104a19190611bee565b60405180910390fd5b5f6104bd83836104b8610a12565b610a2b565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610533578382826040517f64283d7b00000000000000000000000000000000000000000000000000000000815260040161052a939291906120b9565b60405180910390fd5b50505050565b61055383838360405180602001604052805f8152506106db565b505050565b61056a5f82610565610a12565b610a2b565b5050565b5f61057882610953565b9050919050565b5f8073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036105f0575f6040517f89c62b640000000000000000000000000000000000000000000000000000000081526004016105e79190611bee565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b60606001805461064490612089565b80601f016020809104026020016040519081016040528092919081815260200182805461067090612089565b80156106bb5780601f10610692576101008083540402835291602001916106bb565b820191905f5260205f20905b81548152906001019060200180831161069e57829003601f168201915b5050505050905090565b6106d76106d0610a12565b8383610acf565b5050565b6106e684848461043a565b6106fa6106f1610a12565b85858585610c38565b50505050565b606061070b82610de4565b9050919050565b60075481565b60085f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079e9061215e565b60405180910390fd5b60075f8154809291906107b9906121a9565b91905055505f60075490506107ce8382610eef565b6107d88183610f0c565b5f15156108068460085f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610840565b15150361083b5761083a8360085f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166001610acf565b5b505050565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b60085f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f634906490660e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061094c575061094b82610f66565b5b9050919050565b5f8061095e83611047565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036109d057826040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016109c79190611cf9565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b610a268383836001611080565b505050565b5f60085f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610abb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ab29061215e565b60405180910390fd5b610ac684848461123f565b90509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610b3f57816040517f5b08ba18000000000000000000000000000000000000000000000000000000008152600401610b369190611bee565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051610c2b9190611a8e565b60405180910390a3505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b1115610ddd578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b8152600401610c969493929190612242565b6020604051808303815f875af1925050508015610cd157506040513d601f19601f82011682018060405250810190610cce91906122a0565b60015b610d52573d805f8114610cff576040519150601f19603f3d011682016040523d82523d5f602084013e610d04565b606091505b505f815103610d4a57836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401610d419190611bee565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610ddb57836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401610dd29190611bee565b60405180910390fd5b505b5050505050565b6060610def82610953565b505f60065f8481526020019081526020015f208054610e0d90612089565b80601f0160208091040260200160405190810160405280929190818152602001828054610e3990612089565b8015610e845780601f10610e5b57610100808354040283529160200191610e84565b820191905f5260205f20905b815481529060010190602001808311610e6757829003601f168201915b505050505090505f610e9461144a565b90505f815103610ea8578192505050610eea565b5f82511115610edc578082604051602001610ec4929190612305565b60405160208183030381529060405292505050610eea565b610ee584611460565b925050505b919050565b610f08828260405180602001604052805f8152506114c6565b5050565b8060065f8481526020019081526020015f209081610f2a91906124c5565b507ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce782604051610f5a9190611cf9565b60405180910390a15050565b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061103057507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80611040575061103f826114e9565b5b9050919050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b80806110b857505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b156111ea575f6110c784610953565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561113157508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b801561114457506111428184610840565b155b1561118657826040517fa9fbf51f00000000000000000000000000000000000000000000000000000000815260040161117d9190611bee565b60405180910390fd5b81156111e857838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b5f8061124a84611047565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161461128b5761128a818486611552565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611316576112ca5f855f80611080565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161461139557600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b606060405180602001604052805f815250905090565b606061146b82610953565b505f61147561144a565b90505f8151116114935760405180602001604052805f8152506114be565b8061149d84611615565b6040516020016114ae929190612305565b6040516020818303038152906040525b915050919050565b6114d083836116df565b6114e46114db610a12565b5f858585610c38565b505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b61155d8383836117d2565b611610575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036115d157806040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016115c89190611cf9565b60405180910390fd5b81816040517f177e802f000000000000000000000000000000000000000000000000000000008152600401611607929190612594565b60405180910390fd5b505050565b60605f600161162384611892565b0190505f8167ffffffffffffffff81111561164157611640611d82565b5b6040519080825280601f01601f1916602001820160405280156116735781602001600182028036833780820191505090505b5090505f82602001820190505b6001156116d4578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85816116c9576116c86125bb565b5b0494505f8503611680575b819350505050919050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361174f575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016117469190611bee565b60405180910390fd5b5f61175b83835f610a2b565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146117cd575f6040517f73c6ac6e0000000000000000000000000000000000000000000000000000000081526004016117c49190611bee565b60405180910390fd5b505050565b5f8073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561188957508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16148061184a57506118498484610840565b5b8061188857508273ffffffffffffffffffffffffffffffffffffffff16611870836109d9565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b5f805f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106118ee577a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083816118e4576118e36125bb565b5b0492506040810190505b6d04ee2d6d415b85acef8100000000831061192b576d04ee2d6d415b85acef81000000008381611921576119206125bb565b5b0492506020810190505b662386f26fc10000831061195a57662386f26fc1000083816119505761194f6125bb565b5b0492506010810190505b6305f5e1008310611983576305f5e1008381611979576119786125bb565b5b0492506008810190505b61271083106119a857612710838161199e5761199d6125bb565b5b0492506004810190505b606483106119cb57606483816119c1576119c06125bb565b5b0492506002810190505b600a83106119da576001810190505b80915050919050565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611a28816119f4565b8114611a32575f80fd5b50565b5f81359050611a4381611a1f565b92915050565b5f60208284031215611a5e57611a5d6119ec565b5b5f611a6b84828501611a35565b91505092915050565b5f8115159050919050565b611a8881611a74565b82525050565b5f602082019050611aa15f830184611a7f565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015611ade578082015181840152602081019050611ac3565b5f8484015250505050565b5f601f19601f8301169050919050565b5f611b0382611aa7565b611b0d8185611ab1565b9350611b1d818560208601611ac1565b611b2681611ae9565b840191505092915050565b5f6020820190508181035f830152611b498184611af9565b905092915050565b5f819050919050565b611b6381611b51565b8114611b6d575f80fd5b50565b5f81359050611b7e81611b5a565b92915050565b5f60208284031215611b9957611b986119ec565b5b5f611ba684828501611b70565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611bd882611baf565b9050919050565b611be881611bce565b82525050565b5f602082019050611c015f830184611bdf565b92915050565b611c1081611bce565b8114611c1a575f80fd5b50565b5f81359050611c2b81611c07565b92915050565b5f8060408385031215611c4757611c466119ec565b5b5f611c5485828601611c1d565b9250506020611c6585828601611b70565b9150509250929050565b5f805f60608486031215611c8657611c856119ec565b5b5f611c9386828701611c1d565b9350506020611ca486828701611c1d565b9250506040611cb586828701611b70565b9150509250925092565b5f60208284031215611cd457611cd36119ec565b5b5f611ce184828501611c1d565b91505092915050565b611cf381611b51565b82525050565b5f602082019050611d0c5f830184611cea565b92915050565b611d1b81611a74565b8114611d25575f80fd5b50565b5f81359050611d3681611d12565b92915050565b5f8060408385031215611d5257611d516119ec565b5b5f611d5f85828601611c1d565b9250506020611d7085828601611d28565b9150509250929050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611db882611ae9565b810181811067ffffffffffffffff82111715611dd757611dd6611d82565b5b80604052505050565b5f611de96119e3565b9050611df58282611daf565b919050565b5f67ffffffffffffffff821115611e1457611e13611d82565b5b611e1d82611ae9565b9050602081019050919050565b828183375f83830152505050565b5f611e4a611e4584611dfa565b611de0565b905082815260208101848484011115611e6657611e65611d7e565b5b611e71848285611e2a565b509392505050565b5f82601f830112611e8d57611e8c611d7a565b5b8135611e9d848260208601611e38565b91505092915050565b5f805f8060808587031215611ebe57611ebd6119ec565b5b5f611ecb87828801611c1d565b9450506020611edc87828801611c1d565b9350506040611eed87828801611b70565b925050606085013567ffffffffffffffff811115611f0e57611f0d6119f0565b5b611f1a87828801611e79565b91505092959194509250565b5f67ffffffffffffffff821115611f4057611f3f611d82565b5b611f4982611ae9565b9050602081019050919050565b5f611f68611f6384611f26565b611de0565b905082815260208101848484011115611f8457611f83611d7e565b5b611f8f848285611e2a565b509392505050565b5f82601f830112611fab57611faa611d7a565b5b8135611fbb848260208601611f56565b91505092915050565b5f8060408385031215611fda57611fd96119ec565b5b5f611fe785828601611c1d565b925050602083013567ffffffffffffffff811115612008576120076119f0565b5b61201485828601611f97565b9150509250929050565b5f8060408385031215612034576120336119ec565b5b5f61204185828601611c1d565b925050602061205285828601611c1d565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806120a057607f821691505b6020821081036120b3576120b261205c565b5b50919050565b5f6060820190506120cc5f830186611bdf565b6120d96020830185611cea565b6120e66040830184611bdf565b949350505050565b7f4f6e6c7920636f6e74726163742061646d696e2063616e206d696e74206469705f8201527f6c6f6d6100000000000000000000000000000000000000000000000000000000602082015250565b5f612148602483611ab1565b9150612153826120ee565b604082019050919050565b5f6020820190508181035f8301526121758161213c565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6121b382611b51565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036121e5576121e461217c565b5b600182019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f612214826121f0565b61221e81856121fa565b935061222e818560208601611ac1565b61223781611ae9565b840191505092915050565b5f6080820190506122555f830187611bdf565b6122626020830186611bdf565b61226f6040830185611cea565b8181036060830152612281818461220a565b905095945050505050565b5f8151905061229a81611a1f565b92915050565b5f602082840312156122b5576122b46119ec565b5b5f6122c28482850161228c565b91505092915050565b5f81905092915050565b5f6122df82611aa7565b6122e981856122cb565b93506122f9818560208601611ac1565b80840191505092915050565b5f61231082856122d5565b915061231c82846122d5565b91508190509392505050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026123847fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82612349565b61238e8683612349565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6123c96123c46123bf84611b51565b6123a6565b611b51565b9050919050565b5f819050919050565b6123e2836123af565b6123f66123ee826123d0565b848454612355565b825550505050565b5f90565b61240a6123fe565b6124158184846123d9565b505050565b5b818110156124385761242d5f82612402565b60018101905061241b565b5050565b601f82111561247d5761244e81612328565b6124578461233a565b81016020851015612466578190505b61247a6124728561233a565b83018261241a565b50505b505050565b5f82821c905092915050565b5f61249d5f1984600802612482565b1980831691505092915050565b5f6124b5838361248e565b9150826002028217905092915050565b6124ce82611aa7565b67ffffffffffffffff8111156124e7576124e6611d82565b5b6124f18254612089565b6124fc82828561243c565b5f60209050601f83116001811461252d575f841561251b578287015190505b61252585826124aa565b86555061258c565b601f19841661253b86612328565b5f5b828110156125625784890151825560018201915060208501945060208101905061253d565b8683101561257f578489015161257b601f89168261248e565b8355505b6001600288020188555050505b505050505050565b5f6040820190506125a75f830185611bdf565b6125b46020830184611cea565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffdfea2646970667358221220b3104a163e047a2fcb733a802a9b990d38395c2c27f354533b07c35424da6d7764736f6c637828302e382e32352d646576656c6f702e323032342e322e32342b636f6d6d69742e64626137353465630059",
}

// Diploma is an auto generated Go binding around an Ethereum contract.
type Diploma struct {
	abi abi.ABI
}

// NewDiploma creates a new instance of Diploma.
func NewDiploma() *Diploma {
	parsed, err := DiplomaMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Diploma{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Diploma) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (diploma *Diploma) PackAdmin() []byte {
	enc, err := diploma.abi.Pack("admin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (diploma *Diploma) UnpackAdmin(data []byte) (common.Address, error) {
	out, err := diploma.abi.Unpack("admin", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (diploma *Diploma) PackApprove(to common.Address, tokenId *big.Int) []byte {
	enc, err := diploma.abi.Pack("approve", to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (diploma *Diploma) PackBalanceOf(owner common.Address) []byte {
	enc, err := diploma.abi.Pack("balanceOf", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (diploma *Diploma) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := diploma.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (diploma *Diploma) PackBurn(tokenId *big.Int) []byte {
	enc, err := diploma.abi.Pack("burn", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetApproved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (diploma *Diploma) PackGetApproved(tokenId *big.Int) []byte {
	enc, err := diploma.abi.Pack("getApproved", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetApproved is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (diploma *Diploma) UnpackGetApproved(data []byte) (common.Address, error) {
	out, err := diploma.abi.Unpack("getApproved", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (diploma *Diploma) PackIsApprovedForAll(owner common.Address, operator common.Address) []byte {
	enc, err := diploma.abi.Pack("isApprovedForAll", owner, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (diploma *Diploma) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := diploma.abi.Unpack("isApprovedForAll", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (diploma *Diploma) PackName() []byte {
	enc, err := diploma.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (diploma *Diploma) UnpackName(data []byte) (string, error) {
	out, err := diploma.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackOwnerOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (diploma *Diploma) PackOwnerOf(tokenId *big.Int) []byte {
	enc, err := diploma.abi.Pack("ownerOf", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwnerOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (diploma *Diploma) UnpackOwnerOf(data []byte) (common.Address, error) {
	out, err := diploma.abi.Unpack("ownerOf", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackSafeMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd204c45e.
//
// Solidity: function safeMint(address to, string uri) returns()
func (diploma *Diploma) PackSafeMint(to common.Address, uri string) []byte {
	enc, err := diploma.abi.Pack("safeMint", to, uri)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (diploma *Diploma) PackSafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := diploma.abi.Pack("safeTransferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeTransferFrom0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (diploma *Diploma) PackSafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) []byte {
	enc, err := diploma.abi.Pack("safeTransferFrom0", from, to, tokenId, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (diploma *Diploma) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := diploma.abi.Pack("setApprovalForAll", operator, approved)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (diploma *Diploma) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := diploma.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (diploma *Diploma) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := diploma.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (diploma *Diploma) PackSymbol() []byte {
	enc, err := diploma.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (diploma *Diploma) UnpackSymbol(data []byte) (string, error) {
	out, err := diploma.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTokenCounter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (diploma *Diploma) PackTokenCounter() []byte {
	enc, err := diploma.abi.Pack("tokenCounter")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenCounter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (diploma *Diploma) UnpackTokenCounter(data []byte) (*big.Int, error) {
	out, err := diploma.abi.Unpack("tokenCounter", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackTokenURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (diploma *Diploma) PackTokenURI(tokenId *big.Int) []byte {
	enc, err := diploma.abi.Pack("tokenURI", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenURI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (diploma *Diploma) UnpackTokenURI(data []byte) (string, error) {
	out, err := diploma.abi.Unpack("tokenURI", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (diploma *Diploma) PackTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := diploma.abi.Pack("transferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// DiplomaApproval represents a Approval event raised by the Diploma contract.
type DiplomaApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const DiplomaApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (DiplomaApproval) ContractEventName() string {
	return DiplomaApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (diploma *Diploma) UnpackApprovalEvent(log *types.Log) (*DiplomaApproval, error) {
	event := "Approval"
	if log.Topics[0] != diploma.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DiplomaApproval)
	if len(log.Data) > 0 {
		if err := diploma.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diploma.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// DiplomaApprovalForAll represents a ApprovalForAll event raised by the Diploma contract.
type DiplomaApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const DiplomaApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (DiplomaApprovalForAll) ContractEventName() string {
	return DiplomaApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (diploma *Diploma) UnpackApprovalForAllEvent(log *types.Log) (*DiplomaApprovalForAll, error) {
	event := "ApprovalForAll"
	if log.Topics[0] != diploma.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DiplomaApprovalForAll)
	if len(log.Data) > 0 {
		if err := diploma.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diploma.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// DiplomaBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Diploma contract.
type DiplomaBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const DiplomaBatchMetadataUpdateEventName = "BatchMetadataUpdate"

// ContractEventName returns the user-defined event name.
func (DiplomaBatchMetadataUpdate) ContractEventName() string {
	return DiplomaBatchMetadataUpdateEventName
}

// UnpackBatchMetadataUpdateEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (diploma *Diploma) UnpackBatchMetadataUpdateEvent(log *types.Log) (*DiplomaBatchMetadataUpdate, error) {
	event := "BatchMetadataUpdate"
	if log.Topics[0] != diploma.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DiplomaBatchMetadataUpdate)
	if len(log.Data) > 0 {
		if err := diploma.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diploma.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// DiplomaMetadataUpdate represents a MetadataUpdate event raised by the Diploma contract.
type DiplomaMetadataUpdate struct {
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const DiplomaMetadataUpdateEventName = "MetadataUpdate"

// ContractEventName returns the user-defined event name.
func (DiplomaMetadataUpdate) ContractEventName() string {
	return DiplomaMetadataUpdateEventName
}

// UnpackMetadataUpdateEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (diploma *Diploma) UnpackMetadataUpdateEvent(log *types.Log) (*DiplomaMetadataUpdate, error) {
	event := "MetadataUpdate"
	if log.Topics[0] != diploma.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DiplomaMetadataUpdate)
	if len(log.Data) > 0 {
		if err := diploma.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diploma.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// DiplomaTransfer represents a Transfer event raised by the Diploma contract.
type DiplomaTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const DiplomaTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (DiplomaTransfer) ContractEventName() string {
	return DiplomaTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (diploma *Diploma) UnpackTransferEvent(log *types.Log) (*DiplomaTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != diploma.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DiplomaTransfer)
	if len(log.Data) > 0 {
		if err := diploma.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diploma.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (diploma *Diploma) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], diploma.abi.Errors["ERC721IncorrectOwner"].ID.Bytes()[:4]) {
		return diploma.UnpackERC721IncorrectOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], diploma.abi.Errors["ERC721InsufficientApproval"].ID.Bytes()[:4]) {
		return diploma.UnpackERC721InsufficientApprovalError(raw[4:])
	}
	if bytes.Equal(raw[:4], diploma.abi.Errors["ERC721InvalidApprover"].ID.Bytes()[:4]) {
		return diploma.UnpackERC721InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], diploma.abi.Errors["ERC721InvalidOperator"].ID.Bytes()[:4]) {
		return diploma.UnpackERC721InvalidOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], diploma.abi.Errors["ERC721InvalidOwner"].ID.Bytes()[:4]) {
		return diploma.UnpackERC721InvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], diploma.abi.Errors["ERC721InvalidReceiver"].ID.Bytes()[:4]) {
		return diploma.UnpackERC721InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], diploma.abi.Errors["ERC721InvalidSender"].ID.Bytes()[:4]) {
		return diploma.UnpackERC721InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], diploma.abi.Errors["ERC721NonexistentToken"].ID.Bytes()[:4]) {
		return diploma.UnpackERC721NonexistentTokenError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// DiplomaERC721IncorrectOwner represents a ERC721IncorrectOwner error raised by the Diploma contract.
type DiplomaERC721IncorrectOwner struct {
	Sender  common.Address
	TokenId *big.Int
	Owner   common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721IncorrectOwner(address sender, uint256 tokenId, address owner)
func DiplomaERC721IncorrectOwnerErrorID() common.Hash {
	return common.HexToHash("0x64283d7b313c8117c125f736876fa2b4e90ea3831a4716dfdb87d2f540e26289")
}

// UnpackERC721IncorrectOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721IncorrectOwner(address sender, uint256 tokenId, address owner)
func (diploma *Diploma) UnpackERC721IncorrectOwnerError(raw []byte) (*DiplomaERC721IncorrectOwner, error) {
	out := new(DiplomaERC721IncorrectOwner)
	if err := diploma.abi.UnpackIntoInterface(out, "ERC721IncorrectOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiplomaERC721InsufficientApproval represents a ERC721InsufficientApproval error raised by the Diploma contract.
type DiplomaERC721InsufficientApproval struct {
	Operator common.Address
	TokenId  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InsufficientApproval(address operator, uint256 tokenId)
func DiplomaERC721InsufficientApprovalErrorID() common.Hash {
	return common.HexToHash("0x177e802f6f313bc89797ecace66d6d29ab4719cbaaacbb87367264048b1eb861")
}

// UnpackERC721InsufficientApprovalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InsufficientApproval(address operator, uint256 tokenId)
func (diploma *Diploma) UnpackERC721InsufficientApprovalError(raw []byte) (*DiplomaERC721InsufficientApproval, error) {
	out := new(DiplomaERC721InsufficientApproval)
	if err := diploma.abi.UnpackIntoInterface(out, "ERC721InsufficientApproval", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiplomaERC721InvalidApprover represents a ERC721InvalidApprover error raised by the Diploma contract.
type DiplomaERC721InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidApprover(address approver)
func DiplomaERC721InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xa9fbf51f86b8e03595d59dc726bb10c329bb24f62589be276d8dd193ca0b69ea")
}

// UnpackERC721InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidApprover(address approver)
func (diploma *Diploma) UnpackERC721InvalidApproverError(raw []byte) (*DiplomaERC721InvalidApprover, error) {
	out := new(DiplomaERC721InvalidApprover)
	if err := diploma.abi.UnpackIntoInterface(out, "ERC721InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiplomaERC721InvalidOperator represents a ERC721InvalidOperator error raised by the Diploma contract.
type DiplomaERC721InvalidOperator struct {
	Operator common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidOperator(address operator)
func DiplomaERC721InvalidOperatorErrorID() common.Hash {
	return common.HexToHash("0x5b08ba185e8f577075361f3a3555a6580a227ce22734dcc979c1aeadf894658b")
}

// UnpackERC721InvalidOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidOperator(address operator)
func (diploma *Diploma) UnpackERC721InvalidOperatorError(raw []byte) (*DiplomaERC721InvalidOperator, error) {
	out := new(DiplomaERC721InvalidOperator)
	if err := diploma.abi.UnpackIntoInterface(out, "ERC721InvalidOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiplomaERC721InvalidOwner represents a ERC721InvalidOwner error raised by the Diploma contract.
type DiplomaERC721InvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidOwner(address owner)
func DiplomaERC721InvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x89c62b6479af2e623826dcc39c5133061d35b66d72de92833401dd2fd6567480")
}

// UnpackERC721InvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidOwner(address owner)
func (diploma *Diploma) UnpackERC721InvalidOwnerError(raw []byte) (*DiplomaERC721InvalidOwner, error) {
	out := new(DiplomaERC721InvalidOwner)
	if err := diploma.abi.UnpackIntoInterface(out, "ERC721InvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiplomaERC721InvalidReceiver represents a ERC721InvalidReceiver error raised by the Diploma contract.
type DiplomaERC721InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidReceiver(address receiver)
func DiplomaERC721InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0x64a0ae9278f805eaf991dcd18ca78756d280b7508b764ef1b255c55845c11df9")
}

// UnpackERC721InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidReceiver(address receiver)
func (diploma *Diploma) UnpackERC721InvalidReceiverError(raw []byte) (*DiplomaERC721InvalidReceiver, error) {
	out := new(DiplomaERC721InvalidReceiver)
	if err := diploma.abi.UnpackIntoInterface(out, "ERC721InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiplomaERC721InvalidSender represents a ERC721InvalidSender error raised by the Diploma contract.
type DiplomaERC721InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidSender(address sender)
func DiplomaERC721InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x73c6ac6e10798e95d99e1f130d923eb40193ecb8d094ec3dce93292564eb3b17")
}

// UnpackERC721InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidSender(address sender)
func (diploma *Diploma) UnpackERC721InvalidSenderError(raw []byte) (*DiplomaERC721InvalidSender, error) {
	out := new(DiplomaERC721InvalidSender)
	if err := diploma.abi.UnpackIntoInterface(out, "ERC721InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiplomaERC721NonexistentToken represents a ERC721NonexistentToken error raised by the Diploma contract.
type DiplomaERC721NonexistentToken struct {
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721NonexistentToken(uint256 tokenId)
func DiplomaERC721NonexistentTokenErrorID() common.Hash {
	return common.HexToHash("0x7e273289a3a9ef6670f06df7dca227856fc925e956db96980692764a8bc734d7")
}

// UnpackERC721NonexistentTokenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721NonexistentToken(uint256 tokenId)
func (diploma *Diploma) UnpackERC721NonexistentTokenError(raw []byte) (*DiplomaERC721NonexistentToken, error) {
	out := new(DiplomaERC721NonexistentToken)
	if err := diploma.abi.UnpackIntoInterface(out, "ERC721NonexistentToken", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC165MetaData contains all meta data concerning the ERC165 contract.
var ERC165MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "fbe6207e789596d822f8d6c84d99c25e2c",
}

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	abi abi.ABI
}

// NewERC165 creates a new instance of ERC165.
func NewERC165() *ERC165 {
	parsed, err := ERC165MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC165{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC165) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC165 *ERC165) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC165.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC165 *ERC165) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC165.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// ERC721MetaData contains all meta data concerning the ERC721 contract.
var ERC721MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "0f963b4e355adba23f537209afc38a19f8",
}

// ERC721 is an auto generated Go binding around an Ethereum contract.
type ERC721 struct {
	abi abi.ABI
}

// NewERC721 creates a new instance of ERC721.
func NewERC721() *ERC721 {
	parsed, err := ERC721MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC721{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC721) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (eRC721 *ERC721) PackApprove(to common.Address, tokenId *big.Int) []byte {
	enc, err := eRC721.abi.Pack("approve", to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (eRC721 *ERC721) PackBalanceOf(owner common.Address) []byte {
	enc, err := eRC721.abi.Pack("balanceOf", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (eRC721 *ERC721) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC721.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetApproved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (eRC721 *ERC721) PackGetApproved(tokenId *big.Int) []byte {
	enc, err := eRC721.abi.Pack("getApproved", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetApproved is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (eRC721 *ERC721) UnpackGetApproved(data []byte) (common.Address, error) {
	out, err := eRC721.abi.Unpack("getApproved", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (eRC721 *ERC721) PackIsApprovedForAll(owner common.Address, operator common.Address) []byte {
	enc, err := eRC721.abi.Pack("isApprovedForAll", owner, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (eRC721 *ERC721) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := eRC721.abi.Unpack("isApprovedForAll", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC721 *ERC721) PackName() []byte {
	enc, err := eRC721.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC721 *ERC721) UnpackName(data []byte) (string, error) {
	out, err := eRC721.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackOwnerOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (eRC721 *ERC721) PackOwnerOf(tokenId *big.Int) []byte {
	enc, err := eRC721.abi.Pack("ownerOf", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwnerOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (eRC721 *ERC721) UnpackOwnerOf(data []byte) (common.Address, error) {
	out, err := eRC721.abi.Unpack("ownerOf", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (eRC721 *ERC721) PackSafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := eRC721.abi.Pack("safeTransferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeTransferFrom0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (eRC721 *ERC721) PackSafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) []byte {
	enc, err := eRC721.abi.Pack("safeTransferFrom0", from, to, tokenId, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (eRC721 *ERC721) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := eRC721.abi.Pack("setApprovalForAll", operator, approved)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC721 *ERC721) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC721.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC721 *ERC721) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC721.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC721 *ERC721) PackSymbol() []byte {
	enc, err := eRC721.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC721 *ERC721) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC721.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTokenURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (eRC721 *ERC721) PackTokenURI(tokenId *big.Int) []byte {
	enc, err := eRC721.abi.Pack("tokenURI", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenURI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (eRC721 *ERC721) UnpackTokenURI(data []byte) (string, error) {
	out, err := eRC721.abi.Unpack("tokenURI", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (eRC721 *ERC721) PackTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := eRC721.abi.Pack("transferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// ERC721Approval represents a Approval event raised by the ERC721 contract.
type ERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const ERC721ApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC721Approval) ContractEventName() string {
	return ERC721ApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (eRC721 *ERC721) UnpackApprovalEvent(log *types.Log) (*ERC721Approval, error) {
	event := "Approval"
	if log.Topics[0] != eRC721.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721Approval)
	if len(log.Data) > 0 {
		if err := eRC721.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC721ApprovalForAll represents a ApprovalForAll event raised by the ERC721 contract.
type ERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const ERC721ApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (ERC721ApprovalForAll) ContractEventName() string {
	return ERC721ApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (eRC721 *ERC721) UnpackApprovalForAllEvent(log *types.Log) (*ERC721ApprovalForAll, error) {
	event := "ApprovalForAll"
	if log.Topics[0] != eRC721.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721ApprovalForAll)
	if len(log.Data) > 0 {
		if err := eRC721.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC721Transfer represents a Transfer event raised by the ERC721 contract.
type ERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC721TransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC721Transfer) ContractEventName() string {
	return ERC721TransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (eRC721 *ERC721) UnpackTransferEvent(log *types.Log) (*ERC721Transfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC721.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721Transfer)
	if len(log.Data) > 0 {
		if err := eRC721.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (eRC721 *ERC721) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC721.abi.Errors["ERC721IncorrectOwner"].ID.Bytes()[:4]) {
		return eRC721.UnpackERC721IncorrectOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721.abi.Errors["ERC721InsufficientApproval"].ID.Bytes()[:4]) {
		return eRC721.UnpackERC721InsufficientApprovalError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721.abi.Errors["ERC721InvalidApprover"].ID.Bytes()[:4]) {
		return eRC721.UnpackERC721InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721.abi.Errors["ERC721InvalidOperator"].ID.Bytes()[:4]) {
		return eRC721.UnpackERC721InvalidOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721.abi.Errors["ERC721InvalidOwner"].ID.Bytes()[:4]) {
		return eRC721.UnpackERC721InvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721.abi.Errors["ERC721InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC721.UnpackERC721InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721.abi.Errors["ERC721InvalidSender"].ID.Bytes()[:4]) {
		return eRC721.UnpackERC721InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721.abi.Errors["ERC721NonexistentToken"].ID.Bytes()[:4]) {
		return eRC721.UnpackERC721NonexistentTokenError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC721ERC721IncorrectOwner represents a ERC721IncorrectOwner error raised by the ERC721 contract.
type ERC721ERC721IncorrectOwner struct {
	Sender  common.Address
	TokenId *big.Int
	Owner   common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721IncorrectOwner(address sender, uint256 tokenId, address owner)
func ERC721ERC721IncorrectOwnerErrorID() common.Hash {
	return common.HexToHash("0x64283d7b313c8117c125f736876fa2b4e90ea3831a4716dfdb87d2f540e26289")
}

// UnpackERC721IncorrectOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721IncorrectOwner(address sender, uint256 tokenId, address owner)
func (eRC721 *ERC721) UnpackERC721IncorrectOwnerError(raw []byte) (*ERC721ERC721IncorrectOwner, error) {
	out := new(ERC721ERC721IncorrectOwner)
	if err := eRC721.abi.UnpackIntoInterface(out, "ERC721IncorrectOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ERC721InsufficientApproval represents a ERC721InsufficientApproval error raised by the ERC721 contract.
type ERC721ERC721InsufficientApproval struct {
	Operator common.Address
	TokenId  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InsufficientApproval(address operator, uint256 tokenId)
func ERC721ERC721InsufficientApprovalErrorID() common.Hash {
	return common.HexToHash("0x177e802f6f313bc89797ecace66d6d29ab4719cbaaacbb87367264048b1eb861")
}

// UnpackERC721InsufficientApprovalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InsufficientApproval(address operator, uint256 tokenId)
func (eRC721 *ERC721) UnpackERC721InsufficientApprovalError(raw []byte) (*ERC721ERC721InsufficientApproval, error) {
	out := new(ERC721ERC721InsufficientApproval)
	if err := eRC721.abi.UnpackIntoInterface(out, "ERC721InsufficientApproval", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ERC721InvalidApprover represents a ERC721InvalidApprover error raised by the ERC721 contract.
type ERC721ERC721InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidApprover(address approver)
func ERC721ERC721InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xa9fbf51f86b8e03595d59dc726bb10c329bb24f62589be276d8dd193ca0b69ea")
}

// UnpackERC721InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidApprover(address approver)
func (eRC721 *ERC721) UnpackERC721InvalidApproverError(raw []byte) (*ERC721ERC721InvalidApprover, error) {
	out := new(ERC721ERC721InvalidApprover)
	if err := eRC721.abi.UnpackIntoInterface(out, "ERC721InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ERC721InvalidOperator represents a ERC721InvalidOperator error raised by the ERC721 contract.
type ERC721ERC721InvalidOperator struct {
	Operator common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidOperator(address operator)
func ERC721ERC721InvalidOperatorErrorID() common.Hash {
	return common.HexToHash("0x5b08ba185e8f577075361f3a3555a6580a227ce22734dcc979c1aeadf894658b")
}

// UnpackERC721InvalidOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidOperator(address operator)
func (eRC721 *ERC721) UnpackERC721InvalidOperatorError(raw []byte) (*ERC721ERC721InvalidOperator, error) {
	out := new(ERC721ERC721InvalidOperator)
	if err := eRC721.abi.UnpackIntoInterface(out, "ERC721InvalidOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ERC721InvalidOwner represents a ERC721InvalidOwner error raised by the ERC721 contract.
type ERC721ERC721InvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidOwner(address owner)
func ERC721ERC721InvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x89c62b6479af2e623826dcc39c5133061d35b66d72de92833401dd2fd6567480")
}

// UnpackERC721InvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidOwner(address owner)
func (eRC721 *ERC721) UnpackERC721InvalidOwnerError(raw []byte) (*ERC721ERC721InvalidOwner, error) {
	out := new(ERC721ERC721InvalidOwner)
	if err := eRC721.abi.UnpackIntoInterface(out, "ERC721InvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ERC721InvalidReceiver represents a ERC721InvalidReceiver error raised by the ERC721 contract.
type ERC721ERC721InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidReceiver(address receiver)
func ERC721ERC721InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0x64a0ae9278f805eaf991dcd18ca78756d280b7508b764ef1b255c55845c11df9")
}

// UnpackERC721InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidReceiver(address receiver)
func (eRC721 *ERC721) UnpackERC721InvalidReceiverError(raw []byte) (*ERC721ERC721InvalidReceiver, error) {
	out := new(ERC721ERC721InvalidReceiver)
	if err := eRC721.abi.UnpackIntoInterface(out, "ERC721InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ERC721InvalidSender represents a ERC721InvalidSender error raised by the ERC721 contract.
type ERC721ERC721InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidSender(address sender)
func ERC721ERC721InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x73c6ac6e10798e95d99e1f130d923eb40193ecb8d094ec3dce93292564eb3b17")
}

// UnpackERC721InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidSender(address sender)
func (eRC721 *ERC721) UnpackERC721InvalidSenderError(raw []byte) (*ERC721ERC721InvalidSender, error) {
	out := new(ERC721ERC721InvalidSender)
	if err := eRC721.abi.UnpackIntoInterface(out, "ERC721InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721ERC721NonexistentToken represents a ERC721NonexistentToken error raised by the ERC721 contract.
type ERC721ERC721NonexistentToken struct {
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721NonexistentToken(uint256 tokenId)
func ERC721ERC721NonexistentTokenErrorID() common.Hash {
	return common.HexToHash("0x7e273289a3a9ef6670f06df7dca227856fc925e956db96980692764a8bc734d7")
}

// UnpackERC721NonexistentTokenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721NonexistentToken(uint256 tokenId)
func (eRC721 *ERC721) UnpackERC721NonexistentTokenError(raw []byte) (*ERC721ERC721NonexistentToken, error) {
	out := new(ERC721ERC721NonexistentToken)
	if err := eRC721.abi.UnpackIntoInterface(out, "ERC721NonexistentToken", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721BurnableMetaData contains all meta data concerning the ERC721Burnable contract.
var ERC721BurnableMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "3b74845f0397cba3d60ef1578dfc820204",
}

// ERC721Burnable is an auto generated Go binding around an Ethereum contract.
type ERC721Burnable struct {
	abi abi.ABI
}

// NewERC721Burnable creates a new instance of ERC721Burnable.
func NewERC721Burnable() *ERC721Burnable {
	parsed, err := ERC721BurnableMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC721Burnable{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC721Burnable) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (eRC721Burnable *ERC721Burnable) PackApprove(to common.Address, tokenId *big.Int) []byte {
	enc, err := eRC721Burnable.abi.Pack("approve", to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (eRC721Burnable *ERC721Burnable) PackBalanceOf(owner common.Address) []byte {
	enc, err := eRC721Burnable.abi.Pack("balanceOf", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (eRC721Burnable *ERC721Burnable) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC721Burnable.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (eRC721Burnable *ERC721Burnable) PackBurn(tokenId *big.Int) []byte {
	enc, err := eRC721Burnable.abi.Pack("burn", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetApproved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (eRC721Burnable *ERC721Burnable) PackGetApproved(tokenId *big.Int) []byte {
	enc, err := eRC721Burnable.abi.Pack("getApproved", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetApproved is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (eRC721Burnable *ERC721Burnable) UnpackGetApproved(data []byte) (common.Address, error) {
	out, err := eRC721Burnable.abi.Unpack("getApproved", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (eRC721Burnable *ERC721Burnable) PackIsApprovedForAll(owner common.Address, operator common.Address) []byte {
	enc, err := eRC721Burnable.abi.Pack("isApprovedForAll", owner, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (eRC721Burnable *ERC721Burnable) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := eRC721Burnable.abi.Unpack("isApprovedForAll", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC721Burnable *ERC721Burnable) PackName() []byte {
	enc, err := eRC721Burnable.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC721Burnable *ERC721Burnable) UnpackName(data []byte) (string, error) {
	out, err := eRC721Burnable.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackOwnerOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (eRC721Burnable *ERC721Burnable) PackOwnerOf(tokenId *big.Int) []byte {
	enc, err := eRC721Burnable.abi.Pack("ownerOf", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwnerOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (eRC721Burnable *ERC721Burnable) UnpackOwnerOf(data []byte) (common.Address, error) {
	out, err := eRC721Burnable.abi.Unpack("ownerOf", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (eRC721Burnable *ERC721Burnable) PackSafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := eRC721Burnable.abi.Pack("safeTransferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeTransferFrom0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (eRC721Burnable *ERC721Burnable) PackSafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) []byte {
	enc, err := eRC721Burnable.abi.Pack("safeTransferFrom0", from, to, tokenId, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (eRC721Burnable *ERC721Burnable) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := eRC721Burnable.abi.Pack("setApprovalForAll", operator, approved)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC721Burnable *ERC721Burnable) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC721Burnable.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC721Burnable *ERC721Burnable) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC721Burnable.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC721Burnable *ERC721Burnable) PackSymbol() []byte {
	enc, err := eRC721Burnable.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC721Burnable *ERC721Burnable) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC721Burnable.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTokenURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (eRC721Burnable *ERC721Burnable) PackTokenURI(tokenId *big.Int) []byte {
	enc, err := eRC721Burnable.abi.Pack("tokenURI", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenURI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (eRC721Burnable *ERC721Burnable) UnpackTokenURI(data []byte) (string, error) {
	out, err := eRC721Burnable.abi.Unpack("tokenURI", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (eRC721Burnable *ERC721Burnable) PackTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := eRC721Burnable.abi.Pack("transferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// ERC721BurnableApproval represents a Approval event raised by the ERC721Burnable contract.
type ERC721BurnableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const ERC721BurnableApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC721BurnableApproval) ContractEventName() string {
	return ERC721BurnableApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (eRC721Burnable *ERC721Burnable) UnpackApprovalEvent(log *types.Log) (*ERC721BurnableApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC721Burnable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721BurnableApproval)
	if len(log.Data) > 0 {
		if err := eRC721Burnable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721Burnable.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC721BurnableApprovalForAll represents a ApprovalForAll event raised by the ERC721Burnable contract.
type ERC721BurnableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const ERC721BurnableApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (ERC721BurnableApprovalForAll) ContractEventName() string {
	return ERC721BurnableApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (eRC721Burnable *ERC721Burnable) UnpackApprovalForAllEvent(log *types.Log) (*ERC721BurnableApprovalForAll, error) {
	event := "ApprovalForAll"
	if log.Topics[0] != eRC721Burnable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721BurnableApprovalForAll)
	if len(log.Data) > 0 {
		if err := eRC721Burnable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721Burnable.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC721BurnableTransfer represents a Transfer event raised by the ERC721Burnable contract.
type ERC721BurnableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC721BurnableTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC721BurnableTransfer) ContractEventName() string {
	return ERC721BurnableTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (eRC721Burnable *ERC721Burnable) UnpackTransferEvent(log *types.Log) (*ERC721BurnableTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC721Burnable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721BurnableTransfer)
	if len(log.Data) > 0 {
		if err := eRC721Burnable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721Burnable.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (eRC721Burnable *ERC721Burnable) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC721Burnable.abi.Errors["ERC721IncorrectOwner"].ID.Bytes()[:4]) {
		return eRC721Burnable.UnpackERC721IncorrectOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Burnable.abi.Errors["ERC721InsufficientApproval"].ID.Bytes()[:4]) {
		return eRC721Burnable.UnpackERC721InsufficientApprovalError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Burnable.abi.Errors["ERC721InvalidApprover"].ID.Bytes()[:4]) {
		return eRC721Burnable.UnpackERC721InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Burnable.abi.Errors["ERC721InvalidOperator"].ID.Bytes()[:4]) {
		return eRC721Burnable.UnpackERC721InvalidOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Burnable.abi.Errors["ERC721InvalidOwner"].ID.Bytes()[:4]) {
		return eRC721Burnable.UnpackERC721InvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Burnable.abi.Errors["ERC721InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC721Burnable.UnpackERC721InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Burnable.abi.Errors["ERC721InvalidSender"].ID.Bytes()[:4]) {
		return eRC721Burnable.UnpackERC721InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Burnable.abi.Errors["ERC721NonexistentToken"].ID.Bytes()[:4]) {
		return eRC721Burnable.UnpackERC721NonexistentTokenError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC721BurnableERC721IncorrectOwner represents a ERC721IncorrectOwner error raised by the ERC721Burnable contract.
type ERC721BurnableERC721IncorrectOwner struct {
	Sender  common.Address
	TokenId *big.Int
	Owner   common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721IncorrectOwner(address sender, uint256 tokenId, address owner)
func ERC721BurnableERC721IncorrectOwnerErrorID() common.Hash {
	return common.HexToHash("0x64283d7b313c8117c125f736876fa2b4e90ea3831a4716dfdb87d2f540e26289")
}

// UnpackERC721IncorrectOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721IncorrectOwner(address sender, uint256 tokenId, address owner)
func (eRC721Burnable *ERC721Burnable) UnpackERC721IncorrectOwnerError(raw []byte) (*ERC721BurnableERC721IncorrectOwner, error) {
	out := new(ERC721BurnableERC721IncorrectOwner)
	if err := eRC721Burnable.abi.UnpackIntoInterface(out, "ERC721IncorrectOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721BurnableERC721InsufficientApproval represents a ERC721InsufficientApproval error raised by the ERC721Burnable contract.
type ERC721BurnableERC721InsufficientApproval struct {
	Operator common.Address
	TokenId  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InsufficientApproval(address operator, uint256 tokenId)
func ERC721BurnableERC721InsufficientApprovalErrorID() common.Hash {
	return common.HexToHash("0x177e802f6f313bc89797ecace66d6d29ab4719cbaaacbb87367264048b1eb861")
}

// UnpackERC721InsufficientApprovalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InsufficientApproval(address operator, uint256 tokenId)
func (eRC721Burnable *ERC721Burnable) UnpackERC721InsufficientApprovalError(raw []byte) (*ERC721BurnableERC721InsufficientApproval, error) {
	out := new(ERC721BurnableERC721InsufficientApproval)
	if err := eRC721Burnable.abi.UnpackIntoInterface(out, "ERC721InsufficientApproval", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721BurnableERC721InvalidApprover represents a ERC721InvalidApprover error raised by the ERC721Burnable contract.
type ERC721BurnableERC721InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidApprover(address approver)
func ERC721BurnableERC721InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xa9fbf51f86b8e03595d59dc726bb10c329bb24f62589be276d8dd193ca0b69ea")
}

// UnpackERC721InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidApprover(address approver)
func (eRC721Burnable *ERC721Burnable) UnpackERC721InvalidApproverError(raw []byte) (*ERC721BurnableERC721InvalidApprover, error) {
	out := new(ERC721BurnableERC721InvalidApprover)
	if err := eRC721Burnable.abi.UnpackIntoInterface(out, "ERC721InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721BurnableERC721InvalidOperator represents a ERC721InvalidOperator error raised by the ERC721Burnable contract.
type ERC721BurnableERC721InvalidOperator struct {
	Operator common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidOperator(address operator)
func ERC721BurnableERC721InvalidOperatorErrorID() common.Hash {
	return common.HexToHash("0x5b08ba185e8f577075361f3a3555a6580a227ce22734dcc979c1aeadf894658b")
}

// UnpackERC721InvalidOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidOperator(address operator)
func (eRC721Burnable *ERC721Burnable) UnpackERC721InvalidOperatorError(raw []byte) (*ERC721BurnableERC721InvalidOperator, error) {
	out := new(ERC721BurnableERC721InvalidOperator)
	if err := eRC721Burnable.abi.UnpackIntoInterface(out, "ERC721InvalidOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721BurnableERC721InvalidOwner represents a ERC721InvalidOwner error raised by the ERC721Burnable contract.
type ERC721BurnableERC721InvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidOwner(address owner)
func ERC721BurnableERC721InvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x89c62b6479af2e623826dcc39c5133061d35b66d72de92833401dd2fd6567480")
}

// UnpackERC721InvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidOwner(address owner)
func (eRC721Burnable *ERC721Burnable) UnpackERC721InvalidOwnerError(raw []byte) (*ERC721BurnableERC721InvalidOwner, error) {
	out := new(ERC721BurnableERC721InvalidOwner)
	if err := eRC721Burnable.abi.UnpackIntoInterface(out, "ERC721InvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721BurnableERC721InvalidReceiver represents a ERC721InvalidReceiver error raised by the ERC721Burnable contract.
type ERC721BurnableERC721InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidReceiver(address receiver)
func ERC721BurnableERC721InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0x64a0ae9278f805eaf991dcd18ca78756d280b7508b764ef1b255c55845c11df9")
}

// UnpackERC721InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidReceiver(address receiver)
func (eRC721Burnable *ERC721Burnable) UnpackERC721InvalidReceiverError(raw []byte) (*ERC721BurnableERC721InvalidReceiver, error) {
	out := new(ERC721BurnableERC721InvalidReceiver)
	if err := eRC721Burnable.abi.UnpackIntoInterface(out, "ERC721InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721BurnableERC721InvalidSender represents a ERC721InvalidSender error raised by the ERC721Burnable contract.
type ERC721BurnableERC721InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidSender(address sender)
func ERC721BurnableERC721InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x73c6ac6e10798e95d99e1f130d923eb40193ecb8d094ec3dce93292564eb3b17")
}

// UnpackERC721InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidSender(address sender)
func (eRC721Burnable *ERC721Burnable) UnpackERC721InvalidSenderError(raw []byte) (*ERC721BurnableERC721InvalidSender, error) {
	out := new(ERC721BurnableERC721InvalidSender)
	if err := eRC721Burnable.abi.UnpackIntoInterface(out, "ERC721InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721BurnableERC721NonexistentToken represents a ERC721NonexistentToken error raised by the ERC721Burnable contract.
type ERC721BurnableERC721NonexistentToken struct {
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721NonexistentToken(uint256 tokenId)
func ERC721BurnableERC721NonexistentTokenErrorID() common.Hash {
	return common.HexToHash("0x7e273289a3a9ef6670f06df7dca227856fc925e956db96980692764a8bc734d7")
}

// UnpackERC721NonexistentTokenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721NonexistentToken(uint256 tokenId)
func (eRC721Burnable *ERC721Burnable) UnpackERC721NonexistentTokenError(raw []byte) (*ERC721BurnableERC721NonexistentToken, error) {
	out := new(ERC721BurnableERC721NonexistentToken)
	if err := eRC721Burnable.abi.UnpackIntoInterface(out, "ERC721NonexistentToken", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721URIStorageMetaData contains all meta data concerning the ERC721URIStorage contract.
var ERC721URIStorageMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "9054044919cb68bf60ed5b514853fa4e21",
}

// ERC721URIStorage is an auto generated Go binding around an Ethereum contract.
type ERC721URIStorage struct {
	abi abi.ABI
}

// NewERC721URIStorage creates a new instance of ERC721URIStorage.
func NewERC721URIStorage() *ERC721URIStorage {
	parsed, err := ERC721URIStorageMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC721URIStorage{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC721URIStorage) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (eRC721URIStorage *ERC721URIStorage) PackApprove(to common.Address, tokenId *big.Int) []byte {
	enc, err := eRC721URIStorage.abi.Pack("approve", to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (eRC721URIStorage *ERC721URIStorage) PackBalanceOf(owner common.Address) []byte {
	enc, err := eRC721URIStorage.abi.Pack("balanceOf", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (eRC721URIStorage *ERC721URIStorage) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC721URIStorage.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetApproved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (eRC721URIStorage *ERC721URIStorage) PackGetApproved(tokenId *big.Int) []byte {
	enc, err := eRC721URIStorage.abi.Pack("getApproved", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetApproved is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (eRC721URIStorage *ERC721URIStorage) UnpackGetApproved(data []byte) (common.Address, error) {
	out, err := eRC721URIStorage.abi.Unpack("getApproved", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (eRC721URIStorage *ERC721URIStorage) PackIsApprovedForAll(owner common.Address, operator common.Address) []byte {
	enc, err := eRC721URIStorage.abi.Pack("isApprovedForAll", owner, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (eRC721URIStorage *ERC721URIStorage) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := eRC721URIStorage.abi.Unpack("isApprovedForAll", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC721URIStorage *ERC721URIStorage) PackName() []byte {
	enc, err := eRC721URIStorage.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC721URIStorage *ERC721URIStorage) UnpackName(data []byte) (string, error) {
	out, err := eRC721URIStorage.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackOwnerOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (eRC721URIStorage *ERC721URIStorage) PackOwnerOf(tokenId *big.Int) []byte {
	enc, err := eRC721URIStorage.abi.Pack("ownerOf", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwnerOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (eRC721URIStorage *ERC721URIStorage) UnpackOwnerOf(data []byte) (common.Address, error) {
	out, err := eRC721URIStorage.abi.Unpack("ownerOf", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (eRC721URIStorage *ERC721URIStorage) PackSafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := eRC721URIStorage.abi.Pack("safeTransferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeTransferFrom0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (eRC721URIStorage *ERC721URIStorage) PackSafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) []byte {
	enc, err := eRC721URIStorage.abi.Pack("safeTransferFrom0", from, to, tokenId, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (eRC721URIStorage *ERC721URIStorage) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := eRC721URIStorage.abi.Pack("setApprovalForAll", operator, approved)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC721URIStorage *ERC721URIStorage) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC721URIStorage.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC721URIStorage *ERC721URIStorage) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC721URIStorage.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC721URIStorage *ERC721URIStorage) PackSymbol() []byte {
	enc, err := eRC721URIStorage.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC721URIStorage *ERC721URIStorage) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC721URIStorage.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTokenURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (eRC721URIStorage *ERC721URIStorage) PackTokenURI(tokenId *big.Int) []byte {
	enc, err := eRC721URIStorage.abi.Pack("tokenURI", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenURI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (eRC721URIStorage *ERC721URIStorage) UnpackTokenURI(data []byte) (string, error) {
	out, err := eRC721URIStorage.abi.Unpack("tokenURI", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (eRC721URIStorage *ERC721URIStorage) PackTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := eRC721URIStorage.abi.Pack("transferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// ERC721URIStorageApproval represents a Approval event raised by the ERC721URIStorage contract.
type ERC721URIStorageApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const ERC721URIStorageApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC721URIStorageApproval) ContractEventName() string {
	return ERC721URIStorageApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (eRC721URIStorage *ERC721URIStorage) UnpackApprovalEvent(log *types.Log) (*ERC721URIStorageApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC721URIStorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721URIStorageApproval)
	if len(log.Data) > 0 {
		if err := eRC721URIStorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721URIStorage.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC721URIStorageApprovalForAll represents a ApprovalForAll event raised by the ERC721URIStorage contract.
type ERC721URIStorageApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const ERC721URIStorageApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (ERC721URIStorageApprovalForAll) ContractEventName() string {
	return ERC721URIStorageApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (eRC721URIStorage *ERC721URIStorage) UnpackApprovalForAllEvent(log *types.Log) (*ERC721URIStorageApprovalForAll, error) {
	event := "ApprovalForAll"
	if log.Topics[0] != eRC721URIStorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721URIStorageApprovalForAll)
	if len(log.Data) > 0 {
		if err := eRC721URIStorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721URIStorage.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC721URIStorageBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the ERC721URIStorage contract.
type ERC721URIStorageBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const ERC721URIStorageBatchMetadataUpdateEventName = "BatchMetadataUpdate"

// ContractEventName returns the user-defined event name.
func (ERC721URIStorageBatchMetadataUpdate) ContractEventName() string {
	return ERC721URIStorageBatchMetadataUpdateEventName
}

// UnpackBatchMetadataUpdateEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (eRC721URIStorage *ERC721URIStorage) UnpackBatchMetadataUpdateEvent(log *types.Log) (*ERC721URIStorageBatchMetadataUpdate, error) {
	event := "BatchMetadataUpdate"
	if log.Topics[0] != eRC721URIStorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721URIStorageBatchMetadataUpdate)
	if len(log.Data) > 0 {
		if err := eRC721URIStorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721URIStorage.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC721URIStorageMetadataUpdate represents a MetadataUpdate event raised by the ERC721URIStorage contract.
type ERC721URIStorageMetadataUpdate struct {
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC721URIStorageMetadataUpdateEventName = "MetadataUpdate"

// ContractEventName returns the user-defined event name.
func (ERC721URIStorageMetadataUpdate) ContractEventName() string {
	return ERC721URIStorageMetadataUpdateEventName
}

// UnpackMetadataUpdateEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (eRC721URIStorage *ERC721URIStorage) UnpackMetadataUpdateEvent(log *types.Log) (*ERC721URIStorageMetadataUpdate, error) {
	event := "MetadataUpdate"
	if log.Topics[0] != eRC721URIStorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721URIStorageMetadataUpdate)
	if len(log.Data) > 0 {
		if err := eRC721URIStorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721URIStorage.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC721URIStorageTransfer represents a Transfer event raised by the ERC721URIStorage contract.
type ERC721URIStorageTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC721URIStorageTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC721URIStorageTransfer) ContractEventName() string {
	return ERC721URIStorageTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (eRC721URIStorage *ERC721URIStorage) UnpackTransferEvent(log *types.Log) (*ERC721URIStorageTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC721URIStorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721URIStorageTransfer)
	if len(log.Data) > 0 {
		if err := eRC721URIStorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721URIStorage.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (eRC721URIStorage *ERC721URIStorage) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC721URIStorage.abi.Errors["ERC721IncorrectOwner"].ID.Bytes()[:4]) {
		return eRC721URIStorage.UnpackERC721IncorrectOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721URIStorage.abi.Errors["ERC721InsufficientApproval"].ID.Bytes()[:4]) {
		return eRC721URIStorage.UnpackERC721InsufficientApprovalError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721URIStorage.abi.Errors["ERC721InvalidApprover"].ID.Bytes()[:4]) {
		return eRC721URIStorage.UnpackERC721InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721URIStorage.abi.Errors["ERC721InvalidOperator"].ID.Bytes()[:4]) {
		return eRC721URIStorage.UnpackERC721InvalidOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721URIStorage.abi.Errors["ERC721InvalidOwner"].ID.Bytes()[:4]) {
		return eRC721URIStorage.UnpackERC721InvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721URIStorage.abi.Errors["ERC721InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC721URIStorage.UnpackERC721InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721URIStorage.abi.Errors["ERC721InvalidSender"].ID.Bytes()[:4]) {
		return eRC721URIStorage.UnpackERC721InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721URIStorage.abi.Errors["ERC721NonexistentToken"].ID.Bytes()[:4]) {
		return eRC721URIStorage.UnpackERC721NonexistentTokenError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC721URIStorageERC721IncorrectOwner represents a ERC721IncorrectOwner error raised by the ERC721URIStorage contract.
type ERC721URIStorageERC721IncorrectOwner struct {
	Sender  common.Address
	TokenId *big.Int
	Owner   common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721IncorrectOwner(address sender, uint256 tokenId, address owner)
func ERC721URIStorageERC721IncorrectOwnerErrorID() common.Hash {
	return common.HexToHash("0x64283d7b313c8117c125f736876fa2b4e90ea3831a4716dfdb87d2f540e26289")
}

// UnpackERC721IncorrectOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721IncorrectOwner(address sender, uint256 tokenId, address owner)
func (eRC721URIStorage *ERC721URIStorage) UnpackERC721IncorrectOwnerError(raw []byte) (*ERC721URIStorageERC721IncorrectOwner, error) {
	out := new(ERC721URIStorageERC721IncorrectOwner)
	if err := eRC721URIStorage.abi.UnpackIntoInterface(out, "ERC721IncorrectOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721URIStorageERC721InsufficientApproval represents a ERC721InsufficientApproval error raised by the ERC721URIStorage contract.
type ERC721URIStorageERC721InsufficientApproval struct {
	Operator common.Address
	TokenId  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InsufficientApproval(address operator, uint256 tokenId)
func ERC721URIStorageERC721InsufficientApprovalErrorID() common.Hash {
	return common.HexToHash("0x177e802f6f313bc89797ecace66d6d29ab4719cbaaacbb87367264048b1eb861")
}

// UnpackERC721InsufficientApprovalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InsufficientApproval(address operator, uint256 tokenId)
func (eRC721URIStorage *ERC721URIStorage) UnpackERC721InsufficientApprovalError(raw []byte) (*ERC721URIStorageERC721InsufficientApproval, error) {
	out := new(ERC721URIStorageERC721InsufficientApproval)
	if err := eRC721URIStorage.abi.UnpackIntoInterface(out, "ERC721InsufficientApproval", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721URIStorageERC721InvalidApprover represents a ERC721InvalidApprover error raised by the ERC721URIStorage contract.
type ERC721URIStorageERC721InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidApprover(address approver)
func ERC721URIStorageERC721InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xa9fbf51f86b8e03595d59dc726bb10c329bb24f62589be276d8dd193ca0b69ea")
}

// UnpackERC721InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidApprover(address approver)
func (eRC721URIStorage *ERC721URIStorage) UnpackERC721InvalidApproverError(raw []byte) (*ERC721URIStorageERC721InvalidApprover, error) {
	out := new(ERC721URIStorageERC721InvalidApprover)
	if err := eRC721URIStorage.abi.UnpackIntoInterface(out, "ERC721InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721URIStorageERC721InvalidOperator represents a ERC721InvalidOperator error raised by the ERC721URIStorage contract.
type ERC721URIStorageERC721InvalidOperator struct {
	Operator common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidOperator(address operator)
func ERC721URIStorageERC721InvalidOperatorErrorID() common.Hash {
	return common.HexToHash("0x5b08ba185e8f577075361f3a3555a6580a227ce22734dcc979c1aeadf894658b")
}

// UnpackERC721InvalidOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidOperator(address operator)
func (eRC721URIStorage *ERC721URIStorage) UnpackERC721InvalidOperatorError(raw []byte) (*ERC721URIStorageERC721InvalidOperator, error) {
	out := new(ERC721URIStorageERC721InvalidOperator)
	if err := eRC721URIStorage.abi.UnpackIntoInterface(out, "ERC721InvalidOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721URIStorageERC721InvalidOwner represents a ERC721InvalidOwner error raised by the ERC721URIStorage contract.
type ERC721URIStorageERC721InvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidOwner(address owner)
func ERC721URIStorageERC721InvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x89c62b6479af2e623826dcc39c5133061d35b66d72de92833401dd2fd6567480")
}

// UnpackERC721InvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidOwner(address owner)
func (eRC721URIStorage *ERC721URIStorage) UnpackERC721InvalidOwnerError(raw []byte) (*ERC721URIStorageERC721InvalidOwner, error) {
	out := new(ERC721URIStorageERC721InvalidOwner)
	if err := eRC721URIStorage.abi.UnpackIntoInterface(out, "ERC721InvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721URIStorageERC721InvalidReceiver represents a ERC721InvalidReceiver error raised by the ERC721URIStorage contract.
type ERC721URIStorageERC721InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidReceiver(address receiver)
func ERC721URIStorageERC721InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0x64a0ae9278f805eaf991dcd18ca78756d280b7508b764ef1b255c55845c11df9")
}

// UnpackERC721InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidReceiver(address receiver)
func (eRC721URIStorage *ERC721URIStorage) UnpackERC721InvalidReceiverError(raw []byte) (*ERC721URIStorageERC721InvalidReceiver, error) {
	out := new(ERC721URIStorageERC721InvalidReceiver)
	if err := eRC721URIStorage.abi.UnpackIntoInterface(out, "ERC721InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721URIStorageERC721InvalidSender represents a ERC721InvalidSender error raised by the ERC721URIStorage contract.
type ERC721URIStorageERC721InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidSender(address sender)
func ERC721URIStorageERC721InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x73c6ac6e10798e95d99e1f130d923eb40193ecb8d094ec3dce93292564eb3b17")
}

// UnpackERC721InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidSender(address sender)
func (eRC721URIStorage *ERC721URIStorage) UnpackERC721InvalidSenderError(raw []byte) (*ERC721URIStorageERC721InvalidSender, error) {
	out := new(ERC721URIStorageERC721InvalidSender)
	if err := eRC721URIStorage.abi.UnpackIntoInterface(out, "ERC721InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721URIStorageERC721NonexistentToken represents a ERC721NonexistentToken error raised by the ERC721URIStorage contract.
type ERC721URIStorageERC721NonexistentToken struct {
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721NonexistentToken(uint256 tokenId)
func ERC721URIStorageERC721NonexistentTokenErrorID() common.Hash {
	return common.HexToHash("0x7e273289a3a9ef6670f06df7dca227856fc925e956db96980692764a8bc734d7")
}

// UnpackERC721NonexistentTokenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721NonexistentToken(uint256 tokenId)
func (eRC721URIStorage *ERC721URIStorage) UnpackERC721NonexistentTokenError(raw []byte) (*ERC721URIStorageERC721NonexistentToken, error) {
	out := new(ERC721URIStorageERC721NonexistentToken)
	if err := eRC721URIStorage.abi.UnpackIntoInterface(out, "ERC721NonexistentToken", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721UtilsMetaData contains all meta data concerning the ERC721Utils contract.
var ERC721UtilsMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "69a0e1855b20228ef72f47b7d5ea481fe2",
	Bin: "0x607b604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea2646970667358221220a4dcf96872b19a4714fa987433521294c93b7baee43112ebf7f6a803bc048f5064736f6c637828302e382e32352d646576656c6f702e323032342e322e32342b636f6d6d69742e64626137353465630059",
}

// ERC721Utils is an auto generated Go binding around an Ethereum contract.
type ERC721Utils struct {
	abi abi.ABI
}

// NewERC721Utils creates a new instance of ERC721Utils.
func NewERC721Utils() *ERC721Utils {
	parsed, err := ERC721UtilsMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC721Utils{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC721Utils) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// IERC1155ErrorsMetaData contains all meta data concerning the IERC1155Errors contract.
var IERC1155ErrorsMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"}]",
	ID:  "49e98a0c2402d1f94368da16504e36ede1",
}

// IERC1155Errors is an auto generated Go binding around an Ethereum contract.
type IERC1155Errors struct {
	abi abi.ABI
}

// NewIERC1155Errors creates a new instance of IERC1155Errors.
func NewIERC1155Errors() *IERC1155Errors {
	parsed, err := IERC1155ErrorsMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IERC1155Errors{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IERC1155Errors) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (iERC1155Errors *IERC1155Errors) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], iERC1155Errors.abi.Errors["ERC1155InsufficientBalance"].ID.Bytes()[:4]) {
		return iERC1155Errors.UnpackERC1155InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC1155Errors.abi.Errors["ERC1155InvalidApprover"].ID.Bytes()[:4]) {
		return iERC1155Errors.UnpackERC1155InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC1155Errors.abi.Errors["ERC1155InvalidArrayLength"].ID.Bytes()[:4]) {
		return iERC1155Errors.UnpackERC1155InvalidArrayLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC1155Errors.abi.Errors["ERC1155InvalidOperator"].ID.Bytes()[:4]) {
		return iERC1155Errors.UnpackERC1155InvalidOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC1155Errors.abi.Errors["ERC1155InvalidReceiver"].ID.Bytes()[:4]) {
		return iERC1155Errors.UnpackERC1155InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC1155Errors.abi.Errors["ERC1155InvalidSender"].ID.Bytes()[:4]) {
		return iERC1155Errors.UnpackERC1155InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC1155Errors.abi.Errors["ERC1155MissingApprovalForAll"].ID.Bytes()[:4]) {
		return iERC1155Errors.UnpackERC1155MissingApprovalForAllError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// IERC1155ErrorsERC1155InsufficientBalance represents a ERC1155InsufficientBalance error raised by the IERC1155Errors contract.
type IERC1155ErrorsERC1155InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InsufficientBalance(address sender, uint256 balance, uint256 needed, uint256 tokenId)
func IERC1155ErrorsERC1155InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0x03dee4c573c982787b5f3537d6323ffaca9d864448aa6bd828ada9e5d0837036")
}

// UnpackERC1155InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InsufficientBalance(address sender, uint256 balance, uint256 needed, uint256 tokenId)
func (iERC1155Errors *IERC1155Errors) UnpackERC1155InsufficientBalanceError(raw []byte) (*IERC1155ErrorsERC1155InsufficientBalance, error) {
	out := new(IERC1155ErrorsERC1155InsufficientBalance)
	if err := iERC1155Errors.abi.UnpackIntoInterface(out, "ERC1155InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC1155ErrorsERC1155InvalidApprover represents a ERC1155InvalidApprover error raised by the IERC1155Errors contract.
type IERC1155ErrorsERC1155InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidApprover(address approver)
func IERC1155ErrorsERC1155InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0x3e31884e33c33ce0039d1905e3c252950ae3b95240f36d4fff81f5ff6752ef99")
}

// UnpackERC1155InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidApprover(address approver)
func (iERC1155Errors *IERC1155Errors) UnpackERC1155InvalidApproverError(raw []byte) (*IERC1155ErrorsERC1155InvalidApprover, error) {
	out := new(IERC1155ErrorsERC1155InvalidApprover)
	if err := iERC1155Errors.abi.UnpackIntoInterface(out, "ERC1155InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC1155ErrorsERC1155InvalidArrayLength represents a ERC1155InvalidArrayLength error raised by the IERC1155Errors contract.
type IERC1155ErrorsERC1155InvalidArrayLength struct {
	IdsLength    *big.Int
	ValuesLength *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidArrayLength(uint256 idsLength, uint256 valuesLength)
func IERC1155ErrorsERC1155InvalidArrayLengthErrorID() common.Hash {
	return common.HexToHash("0x5b0599913619cfa5633692652638ed25cafcd079c9beae8c251b12c23dcc83f2")
}

// UnpackERC1155InvalidArrayLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidArrayLength(uint256 idsLength, uint256 valuesLength)
func (iERC1155Errors *IERC1155Errors) UnpackERC1155InvalidArrayLengthError(raw []byte) (*IERC1155ErrorsERC1155InvalidArrayLength, error) {
	out := new(IERC1155ErrorsERC1155InvalidArrayLength)
	if err := iERC1155Errors.abi.UnpackIntoInterface(out, "ERC1155InvalidArrayLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC1155ErrorsERC1155InvalidOperator represents a ERC1155InvalidOperator error raised by the IERC1155Errors contract.
type IERC1155ErrorsERC1155InvalidOperator struct {
	Operator common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidOperator(address operator)
func IERC1155ErrorsERC1155InvalidOperatorErrorID() common.Hash {
	return common.HexToHash("0xced3e10010b9d2aa24827119d0db4a8feec73aea48b4b3e470d8a9f3ff723569")
}

// UnpackERC1155InvalidOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidOperator(address operator)
func (iERC1155Errors *IERC1155Errors) UnpackERC1155InvalidOperatorError(raw []byte) (*IERC1155ErrorsERC1155InvalidOperator, error) {
	out := new(IERC1155ErrorsERC1155InvalidOperator)
	if err := iERC1155Errors.abi.UnpackIntoInterface(out, "ERC1155InvalidOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC1155ErrorsERC1155InvalidReceiver represents a ERC1155InvalidReceiver error raised by the IERC1155Errors contract.
type IERC1155ErrorsERC1155InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidReceiver(address receiver)
func IERC1155ErrorsERC1155InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0x57f447ceed621d9e134e26de5772c88799abb7322ce2a87f95dce247d47105c6")
}

// UnpackERC1155InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidReceiver(address receiver)
func (iERC1155Errors *IERC1155Errors) UnpackERC1155InvalidReceiverError(raw []byte) (*IERC1155ErrorsERC1155InvalidReceiver, error) {
	out := new(IERC1155ErrorsERC1155InvalidReceiver)
	if err := iERC1155Errors.abi.UnpackIntoInterface(out, "ERC1155InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC1155ErrorsERC1155InvalidSender represents a ERC1155InvalidSender error raised by the IERC1155Errors contract.
type IERC1155ErrorsERC1155InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidSender(address sender)
func IERC1155ErrorsERC1155InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x01a83514e94b34009110b75cac6742ba33bd7c62f18a3616bafea52855d3b175")
}

// UnpackERC1155InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidSender(address sender)
func (iERC1155Errors *IERC1155Errors) UnpackERC1155InvalidSenderError(raw []byte) (*IERC1155ErrorsERC1155InvalidSender, error) {
	out := new(IERC1155ErrorsERC1155InvalidSender)
	if err := iERC1155Errors.abi.UnpackIntoInterface(out, "ERC1155InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC1155ErrorsERC1155MissingApprovalForAll represents a ERC1155MissingApprovalForAll error raised by the IERC1155Errors contract.
type IERC1155ErrorsERC1155MissingApprovalForAll struct {
	Operator common.Address
	Owner    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155MissingApprovalForAll(address operator, address owner)
func IERC1155ErrorsERC1155MissingApprovalForAllErrorID() common.Hash {
	return common.HexToHash("0xe237d922be9fac42efeaaaffb42cc6b57e0ff95d94a1b74daeff69adc7657754")
}

// UnpackERC1155MissingApprovalForAllError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155MissingApprovalForAll(address operator, address owner)
func (iERC1155Errors *IERC1155Errors) UnpackERC1155MissingApprovalForAllError(raw []byte) (*IERC1155ErrorsERC1155MissingApprovalForAll, error) {
	out := new(IERC1155ErrorsERC1155MissingApprovalForAll)
	if err := iERC1155Errors.abi.UnpackIntoInterface(out, "ERC1155MissingApprovalForAll", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "0c4176746b15005f34f50bb7b0ca8f54b2",
}

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	abi abi.ABI
}

// NewIERC165 creates a new instance of IERC165.
func NewIERC165() *IERC165 {
	parsed, err := IERC165MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IERC165{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IERC165) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (iERC165 *IERC165) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := iERC165.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (iERC165 *IERC165) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := iERC165.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// IERC20ErrorsMetaData contains all meta data concerning the IERC20Errors contract.
var IERC20ErrorsMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"}]",
	ID:  "efe47d2545a61930635f46c4176bc45e9d",
}

// IERC20Errors is an auto generated Go binding around an Ethereum contract.
type IERC20Errors struct {
	abi abi.ABI
}

// NewIERC20Errors creates a new instance of IERC20Errors.
func NewIERC20Errors() *IERC20Errors {
	parsed, err := IERC20ErrorsMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IERC20Errors{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IERC20Errors) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (iERC20Errors *IERC20Errors) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], iERC20Errors.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return iERC20Errors.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC20Errors.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return iERC20Errors.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC20Errors.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return iERC20Errors.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC20Errors.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return iERC20Errors.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC20Errors.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return iERC20Errors.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC20Errors.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return iERC20Errors.UnpackERC20InvalidSpenderError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// IERC20ErrorsERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the IERC20Errors contract.
type IERC20ErrorsERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func IERC20ErrorsERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (iERC20Errors *IERC20Errors) UnpackERC20InsufficientAllowanceError(raw []byte) (*IERC20ErrorsERC20InsufficientAllowance, error) {
	out := new(IERC20ErrorsERC20InsufficientAllowance)
	if err := iERC20Errors.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC20ErrorsERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the IERC20Errors contract.
type IERC20ErrorsERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func IERC20ErrorsERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (iERC20Errors *IERC20Errors) UnpackERC20InsufficientBalanceError(raw []byte) (*IERC20ErrorsERC20InsufficientBalance, error) {
	out := new(IERC20ErrorsERC20InsufficientBalance)
	if err := iERC20Errors.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC20ErrorsERC20InvalidApprover represents a ERC20InvalidApprover error raised by the IERC20Errors contract.
type IERC20ErrorsERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func IERC20ErrorsERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (iERC20Errors *IERC20Errors) UnpackERC20InvalidApproverError(raw []byte) (*IERC20ErrorsERC20InvalidApprover, error) {
	out := new(IERC20ErrorsERC20InvalidApprover)
	if err := iERC20Errors.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC20ErrorsERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the IERC20Errors contract.
type IERC20ErrorsERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func IERC20ErrorsERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (iERC20Errors *IERC20Errors) UnpackERC20InvalidReceiverError(raw []byte) (*IERC20ErrorsERC20InvalidReceiver, error) {
	out := new(IERC20ErrorsERC20InvalidReceiver)
	if err := iERC20Errors.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC20ErrorsERC20InvalidSender represents a ERC20InvalidSender error raised by the IERC20Errors contract.
type IERC20ErrorsERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func IERC20ErrorsERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (iERC20Errors *IERC20Errors) UnpackERC20InvalidSenderError(raw []byte) (*IERC20ErrorsERC20InvalidSender, error) {
	out := new(IERC20ErrorsERC20InvalidSender)
	if err := iERC20Errors.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC20ErrorsERC20InvalidSpender represents a ERC20InvalidSpender error raised by the IERC20Errors contract.
type IERC20ErrorsERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func IERC20ErrorsERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (iERC20Errors *IERC20Errors) UnpackERC20InvalidSpenderError(raw []byte) (*IERC20ErrorsERC20InvalidSpender, error) {
	out := new(IERC20ErrorsERC20InvalidSpender)
	if err := iERC20Errors.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC4906MetaData contains all meta data concerning the IERC4906 contract.
var IERC4906MetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "63f5416ce8714e50c7519a9098a49b1515",
}

// IERC4906 is an auto generated Go binding around an Ethereum contract.
type IERC4906 struct {
	abi abi.ABI
}

// NewIERC4906 creates a new instance of IERC4906.
func NewIERC4906() *IERC4906 {
	parsed, err := IERC4906MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IERC4906{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IERC4906) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (iERC4906 *IERC4906) PackApprove(to common.Address, tokenId *big.Int) []byte {
	enc, err := iERC4906.abi.Pack("approve", to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (iERC4906 *IERC4906) PackBalanceOf(owner common.Address) []byte {
	enc, err := iERC4906.abi.Pack("balanceOf", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (iERC4906 *IERC4906) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := iERC4906.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetApproved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (iERC4906 *IERC4906) PackGetApproved(tokenId *big.Int) []byte {
	enc, err := iERC4906.abi.Pack("getApproved", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetApproved is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (iERC4906 *IERC4906) UnpackGetApproved(data []byte) (common.Address, error) {
	out, err := iERC4906.abi.Unpack("getApproved", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (iERC4906 *IERC4906) PackIsApprovedForAll(owner common.Address, operator common.Address) []byte {
	enc, err := iERC4906.abi.Pack("isApprovedForAll", owner, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (iERC4906 *IERC4906) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := iERC4906.abi.Unpack("isApprovedForAll", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackOwnerOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (iERC4906 *IERC4906) PackOwnerOf(tokenId *big.Int) []byte {
	enc, err := iERC4906.abi.Pack("ownerOf", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwnerOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (iERC4906 *IERC4906) UnpackOwnerOf(data []byte) (common.Address, error) {
	out, err := iERC4906.abi.Unpack("ownerOf", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (iERC4906 *IERC4906) PackSafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := iERC4906.abi.Pack("safeTransferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeTransferFrom0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (iERC4906 *IERC4906) PackSafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) []byte {
	enc, err := iERC4906.abi.Pack("safeTransferFrom0", from, to, tokenId, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (iERC4906 *IERC4906) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := iERC4906.abi.Pack("setApprovalForAll", operator, approved)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (iERC4906 *IERC4906) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := iERC4906.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (iERC4906 *IERC4906) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := iERC4906.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (iERC4906 *IERC4906) PackTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := iERC4906.abi.Pack("transferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// IERC4906Approval represents a Approval event raised by the IERC4906 contract.
type IERC4906Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const IERC4906ApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (IERC4906Approval) ContractEventName() string {
	return IERC4906ApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (iERC4906 *IERC4906) UnpackApprovalEvent(log *types.Log) (*IERC4906Approval, error) {
	event := "Approval"
	if log.Topics[0] != iERC4906.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC4906Approval)
	if len(log.Data) > 0 {
		if err := iERC4906.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC4906.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IERC4906ApprovalForAll represents a ApprovalForAll event raised by the IERC4906 contract.
type IERC4906ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const IERC4906ApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (IERC4906ApprovalForAll) ContractEventName() string {
	return IERC4906ApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (iERC4906 *IERC4906) UnpackApprovalForAllEvent(log *types.Log) (*IERC4906ApprovalForAll, error) {
	event := "ApprovalForAll"
	if log.Topics[0] != iERC4906.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC4906ApprovalForAll)
	if len(log.Data) > 0 {
		if err := iERC4906.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC4906.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IERC4906BatchMetadataUpdate represents a BatchMetadataUpdate event raised by the IERC4906 contract.
type IERC4906BatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const IERC4906BatchMetadataUpdateEventName = "BatchMetadataUpdate"

// ContractEventName returns the user-defined event name.
func (IERC4906BatchMetadataUpdate) ContractEventName() string {
	return IERC4906BatchMetadataUpdateEventName
}

// UnpackBatchMetadataUpdateEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (iERC4906 *IERC4906) UnpackBatchMetadataUpdateEvent(log *types.Log) (*IERC4906BatchMetadataUpdate, error) {
	event := "BatchMetadataUpdate"
	if log.Topics[0] != iERC4906.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC4906BatchMetadataUpdate)
	if len(log.Data) > 0 {
		if err := iERC4906.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC4906.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IERC4906MetadataUpdate represents a MetadataUpdate event raised by the IERC4906 contract.
type IERC4906MetadataUpdate struct {
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const IERC4906MetadataUpdateEventName = "MetadataUpdate"

// ContractEventName returns the user-defined event name.
func (IERC4906MetadataUpdate) ContractEventName() string {
	return IERC4906MetadataUpdateEventName
}

// UnpackMetadataUpdateEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (iERC4906 *IERC4906) UnpackMetadataUpdateEvent(log *types.Log) (*IERC4906MetadataUpdate, error) {
	event := "MetadataUpdate"
	if log.Topics[0] != iERC4906.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC4906MetadataUpdate)
	if len(log.Data) > 0 {
		if err := iERC4906.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC4906.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IERC4906Transfer represents a Transfer event raised by the IERC4906 contract.
type IERC4906Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const IERC4906TransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (IERC4906Transfer) ContractEventName() string {
	return IERC4906TransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (iERC4906 *IERC4906) UnpackTransferEvent(log *types.Log) (*IERC4906Transfer, error) {
	event := "Transfer"
	if log.Topics[0] != iERC4906.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC4906Transfer)
	if len(log.Data) > 0 {
		if err := iERC4906.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC4906.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IERC721MetaData contains all meta data concerning the IERC721 contract.
var IERC721MetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "9f29074339aed8185519b3052179829b9d",
}

// IERC721 is an auto generated Go binding around an Ethereum contract.
type IERC721 struct {
	abi abi.ABI
}

// NewIERC721 creates a new instance of IERC721.
func NewIERC721() *IERC721 {
	parsed, err := IERC721MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IERC721{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IERC721) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (iERC721 *IERC721) PackApprove(to common.Address, tokenId *big.Int) []byte {
	enc, err := iERC721.abi.Pack("approve", to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (iERC721 *IERC721) PackBalanceOf(owner common.Address) []byte {
	enc, err := iERC721.abi.Pack("balanceOf", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (iERC721 *IERC721) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := iERC721.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetApproved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (iERC721 *IERC721) PackGetApproved(tokenId *big.Int) []byte {
	enc, err := iERC721.abi.Pack("getApproved", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetApproved is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (iERC721 *IERC721) UnpackGetApproved(data []byte) (common.Address, error) {
	out, err := iERC721.abi.Unpack("getApproved", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (iERC721 *IERC721) PackIsApprovedForAll(owner common.Address, operator common.Address) []byte {
	enc, err := iERC721.abi.Pack("isApprovedForAll", owner, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (iERC721 *IERC721) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := iERC721.abi.Unpack("isApprovedForAll", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackOwnerOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (iERC721 *IERC721) PackOwnerOf(tokenId *big.Int) []byte {
	enc, err := iERC721.abi.Pack("ownerOf", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwnerOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (iERC721 *IERC721) UnpackOwnerOf(data []byte) (common.Address, error) {
	out, err := iERC721.abi.Unpack("ownerOf", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (iERC721 *IERC721) PackSafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := iERC721.abi.Pack("safeTransferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeTransferFrom0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (iERC721 *IERC721) PackSafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) []byte {
	enc, err := iERC721.abi.Pack("safeTransferFrom0", from, to, tokenId, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (iERC721 *IERC721) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := iERC721.abi.Pack("setApprovalForAll", operator, approved)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (iERC721 *IERC721) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := iERC721.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (iERC721 *IERC721) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := iERC721.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (iERC721 *IERC721) PackTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := iERC721.abi.Pack("transferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// IERC721Approval represents a Approval event raised by the IERC721 contract.
type IERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const IERC721ApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (IERC721Approval) ContractEventName() string {
	return IERC721ApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (iERC721 *IERC721) UnpackApprovalEvent(log *types.Log) (*IERC721Approval, error) {
	event := "Approval"
	if log.Topics[0] != iERC721.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC721Approval)
	if len(log.Data) > 0 {
		if err := iERC721.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC721.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IERC721ApprovalForAll represents a ApprovalForAll event raised by the IERC721 contract.
type IERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const IERC721ApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (IERC721ApprovalForAll) ContractEventName() string {
	return IERC721ApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (iERC721 *IERC721) UnpackApprovalForAllEvent(log *types.Log) (*IERC721ApprovalForAll, error) {
	event := "ApprovalForAll"
	if log.Topics[0] != iERC721.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC721ApprovalForAll)
	if len(log.Data) > 0 {
		if err := iERC721.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC721.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IERC721Transfer represents a Transfer event raised by the IERC721 contract.
type IERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const IERC721TransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (IERC721Transfer) ContractEventName() string {
	return IERC721TransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (iERC721 *IERC721) UnpackTransferEvent(log *types.Log) (*IERC721Transfer, error) {
	event := "Transfer"
	if log.Topics[0] != iERC721.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC721Transfer)
	if len(log.Data) > 0 {
		if err := iERC721.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC721.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IERC721ErrorsMetaData contains all meta data concerning the IERC721Errors contract.
var IERC721ErrorsMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"}]",
	ID:  "e778ab2217b8df25a12b9c8141f0edb134",
}

// IERC721Errors is an auto generated Go binding around an Ethereum contract.
type IERC721Errors struct {
	abi abi.ABI
}

// NewIERC721Errors creates a new instance of IERC721Errors.
func NewIERC721Errors() *IERC721Errors {
	parsed, err := IERC721ErrorsMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IERC721Errors{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IERC721Errors) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (iERC721Errors *IERC721Errors) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], iERC721Errors.abi.Errors["ERC721IncorrectOwner"].ID.Bytes()[:4]) {
		return iERC721Errors.UnpackERC721IncorrectOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC721Errors.abi.Errors["ERC721InsufficientApproval"].ID.Bytes()[:4]) {
		return iERC721Errors.UnpackERC721InsufficientApprovalError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC721Errors.abi.Errors["ERC721InvalidApprover"].ID.Bytes()[:4]) {
		return iERC721Errors.UnpackERC721InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC721Errors.abi.Errors["ERC721InvalidOperator"].ID.Bytes()[:4]) {
		return iERC721Errors.UnpackERC721InvalidOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC721Errors.abi.Errors["ERC721InvalidOwner"].ID.Bytes()[:4]) {
		return iERC721Errors.UnpackERC721InvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC721Errors.abi.Errors["ERC721InvalidReceiver"].ID.Bytes()[:4]) {
		return iERC721Errors.UnpackERC721InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC721Errors.abi.Errors["ERC721InvalidSender"].ID.Bytes()[:4]) {
		return iERC721Errors.UnpackERC721InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], iERC721Errors.abi.Errors["ERC721NonexistentToken"].ID.Bytes()[:4]) {
		return iERC721Errors.UnpackERC721NonexistentTokenError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// IERC721ErrorsERC721IncorrectOwner represents a ERC721IncorrectOwner error raised by the IERC721Errors contract.
type IERC721ErrorsERC721IncorrectOwner struct {
	Sender  common.Address
	TokenId *big.Int
	Owner   common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721IncorrectOwner(address sender, uint256 tokenId, address owner)
func IERC721ErrorsERC721IncorrectOwnerErrorID() common.Hash {
	return common.HexToHash("0x64283d7b313c8117c125f736876fa2b4e90ea3831a4716dfdb87d2f540e26289")
}

// UnpackERC721IncorrectOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721IncorrectOwner(address sender, uint256 tokenId, address owner)
func (iERC721Errors *IERC721Errors) UnpackERC721IncorrectOwnerError(raw []byte) (*IERC721ErrorsERC721IncorrectOwner, error) {
	out := new(IERC721ErrorsERC721IncorrectOwner)
	if err := iERC721Errors.abi.UnpackIntoInterface(out, "ERC721IncorrectOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC721ErrorsERC721InsufficientApproval represents a ERC721InsufficientApproval error raised by the IERC721Errors contract.
type IERC721ErrorsERC721InsufficientApproval struct {
	Operator common.Address
	TokenId  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InsufficientApproval(address operator, uint256 tokenId)
func IERC721ErrorsERC721InsufficientApprovalErrorID() common.Hash {
	return common.HexToHash("0x177e802f6f313bc89797ecace66d6d29ab4719cbaaacbb87367264048b1eb861")
}

// UnpackERC721InsufficientApprovalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InsufficientApproval(address operator, uint256 tokenId)
func (iERC721Errors *IERC721Errors) UnpackERC721InsufficientApprovalError(raw []byte) (*IERC721ErrorsERC721InsufficientApproval, error) {
	out := new(IERC721ErrorsERC721InsufficientApproval)
	if err := iERC721Errors.abi.UnpackIntoInterface(out, "ERC721InsufficientApproval", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC721ErrorsERC721InvalidApprover represents a ERC721InvalidApprover error raised by the IERC721Errors contract.
type IERC721ErrorsERC721InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidApprover(address approver)
func IERC721ErrorsERC721InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xa9fbf51f86b8e03595d59dc726bb10c329bb24f62589be276d8dd193ca0b69ea")
}

// UnpackERC721InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidApprover(address approver)
func (iERC721Errors *IERC721Errors) UnpackERC721InvalidApproverError(raw []byte) (*IERC721ErrorsERC721InvalidApprover, error) {
	out := new(IERC721ErrorsERC721InvalidApprover)
	if err := iERC721Errors.abi.UnpackIntoInterface(out, "ERC721InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC721ErrorsERC721InvalidOperator represents a ERC721InvalidOperator error raised by the IERC721Errors contract.
type IERC721ErrorsERC721InvalidOperator struct {
	Operator common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidOperator(address operator)
func IERC721ErrorsERC721InvalidOperatorErrorID() common.Hash {
	return common.HexToHash("0x5b08ba185e8f577075361f3a3555a6580a227ce22734dcc979c1aeadf894658b")
}

// UnpackERC721InvalidOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidOperator(address operator)
func (iERC721Errors *IERC721Errors) UnpackERC721InvalidOperatorError(raw []byte) (*IERC721ErrorsERC721InvalidOperator, error) {
	out := new(IERC721ErrorsERC721InvalidOperator)
	if err := iERC721Errors.abi.UnpackIntoInterface(out, "ERC721InvalidOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC721ErrorsERC721InvalidOwner represents a ERC721InvalidOwner error raised by the IERC721Errors contract.
type IERC721ErrorsERC721InvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidOwner(address owner)
func IERC721ErrorsERC721InvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x89c62b6479af2e623826dcc39c5133061d35b66d72de92833401dd2fd6567480")
}

// UnpackERC721InvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidOwner(address owner)
func (iERC721Errors *IERC721Errors) UnpackERC721InvalidOwnerError(raw []byte) (*IERC721ErrorsERC721InvalidOwner, error) {
	out := new(IERC721ErrorsERC721InvalidOwner)
	if err := iERC721Errors.abi.UnpackIntoInterface(out, "ERC721InvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC721ErrorsERC721InvalidReceiver represents a ERC721InvalidReceiver error raised by the IERC721Errors contract.
type IERC721ErrorsERC721InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidReceiver(address receiver)
func IERC721ErrorsERC721InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0x64a0ae9278f805eaf991dcd18ca78756d280b7508b764ef1b255c55845c11df9")
}

// UnpackERC721InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidReceiver(address receiver)
func (iERC721Errors *IERC721Errors) UnpackERC721InvalidReceiverError(raw []byte) (*IERC721ErrorsERC721InvalidReceiver, error) {
	out := new(IERC721ErrorsERC721InvalidReceiver)
	if err := iERC721Errors.abi.UnpackIntoInterface(out, "ERC721InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC721ErrorsERC721InvalidSender represents a ERC721InvalidSender error raised by the IERC721Errors contract.
type IERC721ErrorsERC721InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidSender(address sender)
func IERC721ErrorsERC721InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x73c6ac6e10798e95d99e1f130d923eb40193ecb8d094ec3dce93292564eb3b17")
}

// UnpackERC721InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidSender(address sender)
func (iERC721Errors *IERC721Errors) UnpackERC721InvalidSenderError(raw []byte) (*IERC721ErrorsERC721InvalidSender, error) {
	out := new(IERC721ErrorsERC721InvalidSender)
	if err := iERC721Errors.abi.UnpackIntoInterface(out, "ERC721InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC721ErrorsERC721NonexistentToken represents a ERC721NonexistentToken error raised by the IERC721Errors contract.
type IERC721ErrorsERC721NonexistentToken struct {
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721NonexistentToken(uint256 tokenId)
func IERC721ErrorsERC721NonexistentTokenErrorID() common.Hash {
	return common.HexToHash("0x7e273289a3a9ef6670f06df7dca227856fc925e956db96980692764a8bc734d7")
}

// UnpackERC721NonexistentTokenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721NonexistentToken(uint256 tokenId)
func (iERC721Errors *IERC721Errors) UnpackERC721NonexistentTokenError(raw []byte) (*IERC721ErrorsERC721NonexistentToken, error) {
	out := new(IERC721ErrorsERC721NonexistentToken)
	if err := iERC721Errors.abi.UnpackIntoInterface(out, "ERC721NonexistentToken", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IERC721MetadataMetaData contains all meta data concerning the IERC721Metadata contract.
var IERC721MetadataMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "61fe0e250065542921b665f5c5901d6954",
}

// IERC721Metadata is an auto generated Go binding around an Ethereum contract.
type IERC721Metadata struct {
	abi abi.ABI
}

// NewIERC721Metadata creates a new instance of IERC721Metadata.
func NewIERC721Metadata() *IERC721Metadata {
	parsed, err := IERC721MetadataMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IERC721Metadata{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IERC721Metadata) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (iERC721Metadata *IERC721Metadata) PackApprove(to common.Address, tokenId *big.Int) []byte {
	enc, err := iERC721Metadata.abi.Pack("approve", to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (iERC721Metadata *IERC721Metadata) PackBalanceOf(owner common.Address) []byte {
	enc, err := iERC721Metadata.abi.Pack("balanceOf", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (iERC721Metadata *IERC721Metadata) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := iERC721Metadata.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetApproved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (iERC721Metadata *IERC721Metadata) PackGetApproved(tokenId *big.Int) []byte {
	enc, err := iERC721Metadata.abi.Pack("getApproved", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetApproved is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (iERC721Metadata *IERC721Metadata) UnpackGetApproved(data []byte) (common.Address, error) {
	out, err := iERC721Metadata.abi.Unpack("getApproved", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (iERC721Metadata *IERC721Metadata) PackIsApprovedForAll(owner common.Address, operator common.Address) []byte {
	enc, err := iERC721Metadata.abi.Pack("isApprovedForAll", owner, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (iERC721Metadata *IERC721Metadata) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := iERC721Metadata.abi.Unpack("isApprovedForAll", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (iERC721Metadata *IERC721Metadata) PackName() []byte {
	enc, err := iERC721Metadata.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (iERC721Metadata *IERC721Metadata) UnpackName(data []byte) (string, error) {
	out, err := iERC721Metadata.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackOwnerOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (iERC721Metadata *IERC721Metadata) PackOwnerOf(tokenId *big.Int) []byte {
	enc, err := iERC721Metadata.abi.Pack("ownerOf", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwnerOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (iERC721Metadata *IERC721Metadata) UnpackOwnerOf(data []byte) (common.Address, error) {
	out, err := iERC721Metadata.abi.Unpack("ownerOf", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (iERC721Metadata *IERC721Metadata) PackSafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := iERC721Metadata.abi.Pack("safeTransferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeTransferFrom0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (iERC721Metadata *IERC721Metadata) PackSafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) []byte {
	enc, err := iERC721Metadata.abi.Pack("safeTransferFrom0", from, to, tokenId, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (iERC721Metadata *IERC721Metadata) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := iERC721Metadata.abi.Pack("setApprovalForAll", operator, approved)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (iERC721Metadata *IERC721Metadata) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := iERC721Metadata.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (iERC721Metadata *IERC721Metadata) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := iERC721Metadata.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (iERC721Metadata *IERC721Metadata) PackSymbol() []byte {
	enc, err := iERC721Metadata.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (iERC721Metadata *IERC721Metadata) UnpackSymbol(data []byte) (string, error) {
	out, err := iERC721Metadata.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTokenURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (iERC721Metadata *IERC721Metadata) PackTokenURI(tokenId *big.Int) []byte {
	enc, err := iERC721Metadata.abi.Pack("tokenURI", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenURI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (iERC721Metadata *IERC721Metadata) UnpackTokenURI(data []byte) (string, error) {
	out, err := iERC721Metadata.abi.Unpack("tokenURI", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (iERC721Metadata *IERC721Metadata) PackTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := iERC721Metadata.abi.Pack("transferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// IERC721MetadataApproval represents a Approval event raised by the IERC721Metadata contract.
type IERC721MetadataApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const IERC721MetadataApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (IERC721MetadataApproval) ContractEventName() string {
	return IERC721MetadataApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (iERC721Metadata *IERC721Metadata) UnpackApprovalEvent(log *types.Log) (*IERC721MetadataApproval, error) {
	event := "Approval"
	if log.Topics[0] != iERC721Metadata.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC721MetadataApproval)
	if len(log.Data) > 0 {
		if err := iERC721Metadata.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC721Metadata.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IERC721MetadataApprovalForAll represents a ApprovalForAll event raised by the IERC721Metadata contract.
type IERC721MetadataApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const IERC721MetadataApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (IERC721MetadataApprovalForAll) ContractEventName() string {
	return IERC721MetadataApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (iERC721Metadata *IERC721Metadata) UnpackApprovalForAllEvent(log *types.Log) (*IERC721MetadataApprovalForAll, error) {
	event := "ApprovalForAll"
	if log.Topics[0] != iERC721Metadata.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC721MetadataApprovalForAll)
	if len(log.Data) > 0 {
		if err := iERC721Metadata.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC721Metadata.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IERC721MetadataTransfer represents a Transfer event raised by the IERC721Metadata contract.
type IERC721MetadataTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const IERC721MetadataTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (IERC721MetadataTransfer) ContractEventName() string {
	return IERC721MetadataTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (iERC721Metadata *IERC721Metadata) UnpackTransferEvent(log *types.Log) (*IERC721MetadataTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != iERC721Metadata.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC721MetadataTransfer)
	if len(log.Data) > 0 {
		if err := iERC721Metadata.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC721Metadata.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// IERC721ReceiverMetaData contains all meta data concerning the IERC721Receiver contract.
var IERC721ReceiverMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "08a7ce7cc0edd7faa65ff9c8d26f106fa1",
}

// IERC721Receiver is an auto generated Go binding around an Ethereum contract.
type IERC721Receiver struct {
	abi abi.ABI
}

// NewIERC721Receiver creates a new instance of IERC721Receiver.
func NewIERC721Receiver() *IERC721Receiver {
	parsed, err := IERC721ReceiverMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IERC721Receiver{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IERC721Receiver) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackOnERC721Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (iERC721Receiver *IERC721Receiver) PackOnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) []byte {
	enc, err := iERC721Receiver.abi.Pack("onERC721Received", operator, from, tokenId, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOnERC721Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (iERC721Receiver *IERC721Receiver) UnpackOnERC721Received(data []byte) ([4]byte, error) {
	out, err := iERC721Receiver.abi.Unpack("onERC721Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, err
}

// MathMetaData contains all meta data concerning the Math contract.
var MathMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "0d8105693c8c1739c1e21e9a98e55557b4",
	Bin: "0x607b604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea2646970667358221220725ba174858da06e71b5664d9b1029336ee1c632b8bef65c125dedabd3261c1c64736f6c637828302e382e32352d646576656c6f702e323032342e322e32342b636f6d6d69742e64626137353465630059",
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	abi abi.ABI
}

// NewMath creates a new instance of Math.
func NewMath() *Math {
	parsed, err := MathMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Math{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Math) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PanicMetaData contains all meta data concerning the Panic contract.
var PanicMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "cd2bb527256802d98b7a039f463a5af59c",
	Bin: "0x607b604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea2646970667358221220f870ff88fc09deadfb178f0a2d68de14549cb6fc988c8868fc886e1a3254a4e364736f6c637828302e382e32352d646576656c6f702e323032342e322e32342b636f6d6d69742e64626137353465630059",
}

// Panic is an auto generated Go binding around an Ethereum contract.
type Panic struct {
	abi abi.ABI
}

// NewPanic creates a new instance of Panic.
func NewPanic() *Panic {
	parsed, err := PanicMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Panic{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Panic) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// SafeCastMetaData contains all meta data concerning the SafeCast contract.
var SafeCastMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntToUint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintToInt\",\"type\":\"error\"}]",
	ID:  "d461b3d9e4efa750d1b2ac716ed17737df",
	Bin: "0x607b604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea2646970667358221220d5b84cc8291846509793ffe4c7d4a48e107042ce5bc98442bf996f10381edf7164736f6c637828302e382e32352d646576656c6f702e323032342e322e32342b636f6d6d69742e64626137353465630059",
}

// SafeCast is an auto generated Go binding around an Ethereum contract.
type SafeCast struct {
	abi abi.ABI
}

// NewSafeCast creates a new instance of SafeCast.
func NewSafeCast() *SafeCast {
	parsed, err := SafeCastMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &SafeCast{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *SafeCast) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (safeCast *SafeCast) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], safeCast.abi.Errors["SafeCastOverflowedIntDowncast"].ID.Bytes()[:4]) {
		return safeCast.UnpackSafeCastOverflowedIntDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], safeCast.abi.Errors["SafeCastOverflowedIntToUint"].ID.Bytes()[:4]) {
		return safeCast.UnpackSafeCastOverflowedIntToUintError(raw[4:])
	}
	if bytes.Equal(raw[:4], safeCast.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return safeCast.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], safeCast.abi.Errors["SafeCastOverflowedUintToInt"].ID.Bytes()[:4]) {
		return safeCast.UnpackSafeCastOverflowedUintToIntError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// SafeCastSafeCastOverflowedIntDowncast represents a SafeCastOverflowedIntDowncast error raised by the SafeCast contract.
type SafeCastSafeCastOverflowedIntDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedIntDowncast(uint8 bits, int256 value)
func SafeCastSafeCastOverflowedIntDowncastErrorID() common.Hash {
	return common.HexToHash("0x327269a7f29c3c5436f42eeed1c1adf0d4d525f36360483f4e83ab79e98f9089")
}

// UnpackSafeCastOverflowedIntDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedIntDowncast(uint8 bits, int256 value)
func (safeCast *SafeCast) UnpackSafeCastOverflowedIntDowncastError(raw []byte) (*SafeCastSafeCastOverflowedIntDowncast, error) {
	out := new(SafeCastSafeCastOverflowedIntDowncast)
	if err := safeCast.abi.UnpackIntoInterface(out, "SafeCastOverflowedIntDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SafeCastSafeCastOverflowedIntToUint represents a SafeCastOverflowedIntToUint error raised by the SafeCast contract.
type SafeCastSafeCastOverflowedIntToUint struct {
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedIntToUint(int256 value)
func SafeCastSafeCastOverflowedIntToUintErrorID() common.Hash {
	return common.HexToHash("0xa8ce4432b175c373e5f41aba830358e5361584f628450fd436c066323ad91ac2")
}

// UnpackSafeCastOverflowedIntToUintError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedIntToUint(int256 value)
func (safeCast *SafeCast) UnpackSafeCastOverflowedIntToUintError(raw []byte) (*SafeCastSafeCastOverflowedIntToUint, error) {
	out := new(SafeCastSafeCastOverflowedIntToUint)
	if err := safeCast.abi.UnpackIntoInterface(out, "SafeCastOverflowedIntToUint", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SafeCastSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the SafeCast contract.
type SafeCastSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func SafeCastSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (safeCast *SafeCast) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*SafeCastSafeCastOverflowedUintDowncast, error) {
	out := new(SafeCastSafeCastOverflowedUintDowncast)
	if err := safeCast.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SafeCastSafeCastOverflowedUintToInt represents a SafeCastOverflowedUintToInt error raised by the SafeCast contract.
type SafeCastSafeCastOverflowedUintToInt struct {
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintToInt(uint256 value)
func SafeCastSafeCastOverflowedUintToIntErrorID() common.Hash {
	return common.HexToHash("0x24775e0629ae69d78c11bae050651b81820407f300ff750ff2be51e4ce75c37f")
}

// UnpackSafeCastOverflowedUintToIntError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintToInt(uint256 value)
func (safeCast *SafeCast) UnpackSafeCastOverflowedUintToIntError(raw []byte) (*SafeCastSafeCastOverflowedUintToInt, error) {
	out := new(SafeCastSafeCastOverflowedUintToInt)
	if err := safeCast.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintToInt", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SignedMathMetaData contains all meta data concerning the SignedMath contract.
var SignedMathMetaData = bind.MetaData{
	ABI: "[]",
	ID:  "c41c813baf04492ab4ca21c0670e754261",
	Bin: "0x607b604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea26469706673582212207a3f11312eabde78932b98eb8d6889f5156c915edd558785b54cae5adf01e83b64736f6c637828302e382e32352d646576656c6f702e323032342e322e32342b636f6d6d69742e64626137353465630059",
}

// SignedMath is an auto generated Go binding around an Ethereum contract.
type SignedMath struct {
	abi abi.ABI
}

// NewSignedMath creates a new instance of SignedMath.
func NewSignedMath() *SignedMath {
	parsed, err := SignedMathMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &SignedMath{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *SignedMath) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StringsInvalidAddressFormat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StringsInvalidChar\",\"type\":\"error\"}]",
	ID:  "9391e99410a643a46d8efb21dcfbe14aa3",
	Bin: "0x607b604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea264697066735822122099d0e5c2b2977e7e8fa964070e17fc7ad45f2e512183f3d03c41d2faba5645db64736f6c637828302e382e32352d646576656c6f702e323032342e322e32342b636f6d6d69742e64626137353465630059",
}

// Strings is an auto generated Go binding around an Ethereum contract.
type Strings struct {
	abi abi.ABI
}

// NewStrings creates a new instance of Strings.
func NewStrings() *Strings {
	parsed, err := StringsMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Strings{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Strings) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (strings *Strings) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], strings.abi.Errors["StringsInsufficientHexLength"].ID.Bytes()[:4]) {
		return strings.UnpackStringsInsufficientHexLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], strings.abi.Errors["StringsInvalidAddressFormat"].ID.Bytes()[:4]) {
		return strings.UnpackStringsInvalidAddressFormatError(raw[4:])
	}
	if bytes.Equal(raw[:4], strings.abi.Errors["StringsInvalidChar"].ID.Bytes()[:4]) {
		return strings.UnpackStringsInvalidCharError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// StringsStringsInsufficientHexLength represents a StringsInsufficientHexLength error raised by the Strings contract.
type StringsStringsInsufficientHexLength struct {
	Value  *big.Int
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringsInsufficientHexLength(uint256 value, uint256 length)
func StringsStringsInsufficientHexLengthErrorID() common.Hash {
	return common.HexToHash("0xe22e27eb9593f9947dc34771120a3349dd201c662753f0b60502fc1d8e422233")
}

// UnpackStringsInsufficientHexLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringsInsufficientHexLength(uint256 value, uint256 length)
func (strings *Strings) UnpackStringsInsufficientHexLengthError(raw []byte) (*StringsStringsInsufficientHexLength, error) {
	out := new(StringsStringsInsufficientHexLength)
	if err := strings.abi.UnpackIntoInterface(out, "StringsInsufficientHexLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StringsStringsInvalidAddressFormat represents a StringsInvalidAddressFormat error raised by the Strings contract.
type StringsStringsInvalidAddressFormat struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringsInvalidAddressFormat()
func StringsStringsInvalidAddressFormatErrorID() common.Hash {
	return common.HexToHash("0x1d15ae44cf5601ace2ebdfa1451a862de1975935c0c8ba6788ae598e5e29a6dd")
}

// UnpackStringsInvalidAddressFormatError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringsInvalidAddressFormat()
func (strings *Strings) UnpackStringsInvalidAddressFormatError(raw []byte) (*StringsStringsInvalidAddressFormat, error) {
	out := new(StringsStringsInvalidAddressFormat)
	if err := strings.abi.UnpackIntoInterface(out, "StringsInvalidAddressFormat", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// StringsStringsInvalidChar represents a StringsInvalidChar error raised by the Strings contract.
type StringsStringsInvalidChar struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringsInvalidChar()
func StringsStringsInvalidCharErrorID() common.Hash {
	return common.HexToHash("0x94e2737eaa44cfdde863f17909ef1e0595339c0eb63c5453ecb27449d2dcd443")
}

// UnpackStringsInvalidCharError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringsInvalidChar()
func (strings *Strings) UnpackStringsInvalidCharError(raw []byte) (*StringsStringsInvalidChar, error) {
	out := new(StringsStringsInvalidChar)
	if err := strings.abi.UnpackIntoInterface(out, "StringsInvalidChar", raw); err != nil {
		return nil, err
	}
	return out, nil
}
