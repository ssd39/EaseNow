// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
package EaseNow

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
	_ = abi.ConvertType
)

// EaseNowMetaData contains all meta data concerning the EaseNow contract.
var EaseNowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaingDebt\",\"type\":\"uint256\"}],\"name\":\"AmountRepaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"locktime\",\"type\":\"uint256\"}],\"name\":\"LenderDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"LenderRewarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LenderWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"privareData\",\"type\":\"bytes32\"}],\"name\":\"UserDefaulted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"privateData\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creditLimit\",\"type\":\"uint256\"}],\"name\":\"UserRegistred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"purchaseData\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"seedProof_\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"creditLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"privateData\",\"type\":\"bytes32\"}],\"name\":\"registerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"reportDefault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPriceRatio\",\"type\":\"uint256\"}],\"name\":\"updatePriceRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052612710600655600160085534801561001a575f80fd5b50336040518060400160405280600981526020017f45617365546f6b656e00000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f454e540000000000000000000000000000000000000000000000000000000000815250816003908161009791906106dc565b5080600490816100a791906106dc565b5050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361011a575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161011191906107ea565b60405180910390fd5b6101298161014160201b60201c565b5061013c3360065461020460201b60201c565b6108c0565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610274575f6040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161026b91906107ea565b60405180910390fd5b6102855f838361028960201b60201c565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036102d9578060025f8282546102cd9190610830565b925050819055506103a7565b5f805f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610362578381836040517fe450d38c00000000000000000000000000000000000000000000000000000000815260040161035993929190610872565b60405180910390fd5b8181035f808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036103ee578060025f8282540392505081905550610438565b805f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161049591906108a7565b60405180910390a3505050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061051d57607f821691505b6020821081036105305761052f6104d9565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026105927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610557565b61059c8683610557565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6105e06105db6105d6846105b4565b6105bd565b6105b4565b9050919050565b5f819050919050565b6105f9836105c6565b61060d610605826105e7565b848454610563565b825550505050565b5f90565b610621610615565b61062c8184846105f0565b505050565b5b8181101561064f576106445f82610619565b600181019050610632565b5050565b601f8211156106945761066581610536565b61066e84610548565b8101602085101561067d578190505b61069161068985610548565b830182610631565b50505b505050565b5f82821c905092915050565b5f6106b45f1984600802610699565b1980831691505092915050565b5f6106cc83836106a5565b9150826002028217905092915050565b6106e5826104a2565b67ffffffffffffffff8111156106fe576106fd6104ac565b5b6107088254610506565b610713828285610653565b5f60209050601f831160018114610744575f8415610732578287015190505b61073c85826106c1565b8655506107a3565b601f19841661075286610536565b5f5b8281101561077957848901518255600182019150602085019450602081019050610754565b868310156107965784890151610792601f8916826106a5565b8355505b6001600288020188555050505b505050505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6107d4826107ab565b9050919050565b6107e4816107ca565b82525050565b5f6020820190506107fd5f8301846107db565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61083a826105b4565b9150610845836105b4565b925082820190508082111561085d5761085c610803565b5b92915050565b61086c816105b4565b82525050565b5f6060820190506108855f8301866107db565b6108926020830185610863565b61089f6040830184610863565b949350505050565b5f6020820190506108ba5f830184610863565b92915050565b6122f1806108cd5f395ff3fe608060405260043610610122575f3560e01c806370a082311161009f578063a9059cbb11610063578063a9059cbb1461039b578063d0e30db0146103d7578063d8bd9260146103e1578063dd62ed3e14610409578063f2fde38b1461044557610129565b806370a08231146102cd578063715018a6146103095780637be0732e1461031f5780638da5cb5b1461034757806395d89b411461037157610129565b80632e1a7d4d116100e65780632e1a7d4d14610221578063313ce567146102495780633a7d5115146102735780633b6631951461029b578063402d8883146102c357610129565b806306fdde031461012d578063095ea7b31461015757806318160ddd1461019357806323b872dd146101bd5780632caabfef146101f957610129565b3661012957005b5f80fd5b348015610138575f80fd5b5061014161046d565b60405161014e91906117e7565b60405180910390f35b348015610162575f80fd5b5061017d60048036038101906101789190611898565b6104fd565b60405161018a91906118f0565b60405180910390f35b34801561019e575f80fd5b506101a761051f565b6040516101b49190611918565b60405180910390f35b3480156101c8575f80fd5b506101e360048036038101906101de9190611931565b610528565b6040516101f091906118f0565b60405180910390f35b348015610204575f80fd5b5061021f600480360381019061021a91906119b4565b610556565b005b34801561022c575f80fd5b5061024760048036038101906102429190611a04565b610638565b005b348015610254575f80fd5b5061025d6108b4565b60405161026a9190611a4a565b60405180910390f35b34801561027e575f80fd5b5061029960048036038101906102949190611a63565b6108bc565b005b3480156102a6575f80fd5b506102c160048036038101906102bc9190611a8e565b610a12565b005b6102cb610a6b565b005b3480156102d8575f80fd5b506102f360048036038101906102ee9190611a63565b610bb6565b6040516103009190611918565b60405180910390f35b348015610314575f80fd5b5061031d610bfb565b005b34801561032a575f80fd5b5061034560048036038101906103409190611ab9565b610c0e565b005b348015610352575f80fd5b5061035b610c1b565b6040516103689190611b18565b60405180910390f35b34801561037c575f80fd5b50610385610c43565b60405161039291906117e7565b60405180910390f35b3480156103a6575f80fd5b506103c160048036038101906103bc9190611898565b610cd3565b6040516103ce91906118f0565b60405180910390f35b6103df610cf5565b005b3480156103ec575f80fd5b5061040760048036038101906104029190611a04565b610dfc565b005b348015610414575f80fd5b5061042f600480360381019061042a9190611b31565b610e54565b60405161043c9190611918565b60405180910390f35b348015610450575f80fd5b5061046b60048036038101906104669190611a63565b610ed6565b005b60606003805461047c90611b9c565b80601f01602080910402602001604051908101604052809291908181526020018280546104a890611b9c565b80156104f35780601f106104ca576101008083540402835291602001916104f3565b820191905f5260205f20905b8154815290600101906020018083116104d657829003601f168201915b5050505050905090565b5f80610507610f5a565b9050610514818585610f61565b600191505092915050565b5f600254905090565b5f80610532610f5a565b905061053f858285610f73565b61054a858585611005565b60019150509392505050565b61055e6110f5565b5f60095f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f801b8160010154146105e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105dd90611c16565b60405180910390fd5b81816001018190555082815f01819055507f7338f2f0fce044e8ff9f8257e4426c8912aa5dc780cbf77f446d4b10dd11a7e084838560405161062a93929190611c43565b60405180910390a150505050565b5f600a5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20905081816001015410156106bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b690611ce8565b60405180910390fd5b43603c825f01546106d09190611d33565b1115610711576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070890611dd6565b60405180910390fd5b5f610727825f015483600101544360085461117c565b90505f81146107745761073a33826111ce565b7f38af265c68180096307acf94a0034cafb4c264beb9c216b7f8b8776b57e82110338260405161076b929190611df4565b60405180910390a15b82826001015f8282546107879190611e1b565b925050819055507ff1efcc5bb6f7c762ce11c94881410b55191722a008338c0c9a88203c74604aea33846040516107bf929190611df4565b60405180910390a1824710156107fa575f47846107dc9190611e1b565b90506107f533600854836107f09190611e4e565b6111ce565b479350505b5f8311156108af575f803373ffffffffffffffffffffffffffffffffffffffff168560405161082890611ebc565b5f6040518083038185875af1925050503d805f8114610862576040519150601f19603f3d011682016040523d82523d5f602084013e610867565b606091505b5091509150816108ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108a390611f1a565b60405180910390fd5b50505b505050565b5f6012905090565b5f60095f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f209050806004015f9054906101000a900460ff161561094d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094490611f82565b60405180910390fd5b5f8160020154118015610970575043605a826003015461096d9190611d33565b11155b6109af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a690612010565b60405180910390fd5b6001816004015f6101000a81548160ff0219169083151502179055507ff08fd0dc9b5a0e09baf039a4d7a8897a6ab538c6de1e8e4c40c82560090b38758282600201548360010154604051610a069392919061202e565b60405180910390a15050565b5f801b60075414610a58576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4f906120ad565b60405180910390fd5b80600781905550610a683361124d565b50565b5f60095f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090508060020154341015610af2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ae99061213b565b60405180910390fd5b5f610b025f34603c60085461117c565b905080610b0e33610bb6565b1015610b4f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b46906121a3565b60405180910390fd5b610b593382611310565b34826002015f828254610b6c9190611e1b565b925050819055507f0f2afdea007852c85c572fd5ec20aede9297fb7614b1cbcd5594de764d3328ff33348460020154604051610baa939291906121c1565b60405180910390a15050565b5f805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b610c036110f5565b610c0c5f61124d565b565b610c166110f5565b505050565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060048054610c5290611b9c565b80601f0160208091040260200160405190810160405280929190818152602001828054610c7e90611b9c565b8015610cc95780601f10610ca057610100808354040283529160200191610cc9565b820191905f5260205f20905b815481529060010190602001808311610cac57829003601f168201915b5050505050905090565b5f80610cdd610f5a565b9050610cea818585611005565b600191505092915050565b5f600a5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f610d4b825f015483600101544360085461117c565b90505f8114610d9857610d5e33826111ce565b7f38af265c68180096307acf94a0034cafb4c264beb9c216b7f8b8776b57e821103382604051610d8f929190611df4565b60405180910390a15b34826001015f828254610dab9190611d33565b9250508190555043825f01819055507f4d0e8513a525a0e8d883d0fb5bf24c215136a9a954918fbce6d9f03d9bfacd133334845f0154604051610df0939291906121c1565b60405180910390a15050565b610e046110f5565b5f801b60075403610e4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4190612240565b60405180910390fd5b8060088190555050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b610ede6110f5565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610f4e575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610f459190611b18565b60405180910390fd5b610f578161124d565b50565b5f33905090565b610f6e838383600161138f565b505050565b5f610f7e8484610e54565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610fff5781811015610ff0578281836040517ffb8f41b2000000000000000000000000000000000000000000000000000000008152600401610fe7939291906121c1565b60405180910390fd5b610ffe84848484035f61138f565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603611075575f6040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260040161106c9190611b18565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036110e5575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016110dc9190611b18565b60405180910390fd5b6110f083838361155e565b505050565b6110fd610f5a565b73ffffffffffffffffffffffffffffffffffffffff1661111b610c1b565b73ffffffffffffffffffffffffffffffffffffffff161461117a5761113e610f5a565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016111719190611b18565b60405180910390fd5b565b5f603c858461118b9190611e1b565b6103e8600f858861119c9190611e4e565b6111a69190611e4e565b6111b0919061228b565b6111ba9190611e4e565b6111c4919061228b565b9050949350505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361123e575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016112359190611b18565b60405180910390fd5b6112495f838361155e565b5050565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611380575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016113779190611b18565b60405180910390fd5b61138b825f8361155e565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16036113ff575f6040517fe602df050000000000000000000000000000000000000000000000000000000081526004016113f69190611b18565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361146f575f6040517f94280d620000000000000000000000000000000000000000000000000000000081526004016114669190611b18565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015611558578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161154f9190611918565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036115ae578060025f8282546115a29190611d33565b9250508190555061167c565b5f805f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015611637578381836040517fe450d38c00000000000000000000000000000000000000000000000000000000815260040161162e939291906121c1565b60405180910390fd5b8181035f808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036116c3578060025f828254039250508190555061170d565b805f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161176a9190611918565b60405180910390a3505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6117b982611777565b6117c38185611781565b93506117d3818560208601611791565b6117dc8161179f565b840191505092915050565b5f6020820190508181035f8301526117ff81846117af565b905092915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6118348261180b565b9050919050565b6118448161182a565b811461184e575f80fd5b50565b5f8135905061185f8161183b565b92915050565b5f819050919050565b61187781611865565b8114611881575f80fd5b50565b5f813590506118928161186e565b92915050565b5f80604083850312156118ae576118ad611807565b5b5f6118bb85828601611851565b92505060206118cc85828601611884565b9150509250929050565b5f8115159050919050565b6118ea816118d6565b82525050565b5f6020820190506119035f8301846118e1565b92915050565b61191281611865565b82525050565b5f60208201905061192b5f830184611909565b92915050565b5f805f6060848603121561194857611947611807565b5b5f61195586828701611851565b935050602061196686828701611851565b925050604061197786828701611884565b9150509250925092565b5f819050919050565b61199381611981565b811461199d575f80fd5b50565b5f813590506119ae8161198a565b92915050565b5f805f606084860312156119cb576119ca611807565b5b5f6119d886828701611851565b93505060206119e986828701611884565b92505060406119fa868287016119a0565b9150509250925092565b5f60208284031215611a1957611a18611807565b5b5f611a2684828501611884565b91505092915050565b5f60ff82169050919050565b611a4481611a2f565b82525050565b5f602082019050611a5d5f830184611a3b565b92915050565b5f60208284031215611a7857611a77611807565b5b5f611a8584828501611851565b91505092915050565b5f60208284031215611aa357611aa2611807565b5b5f611ab0848285016119a0565b91505092915050565b5f805f60608486031215611ad057611acf611807565b5b5f611add86828701611884565b9350506020611aee868287016119a0565b9250506040611aff86828701611851565b9150509250925092565b611b128161182a565b82525050565b5f602082019050611b2b5f830184611b09565b92915050565b5f8060408385031215611b4757611b46611807565b5b5f611b5485828601611851565b9250506020611b6585828601611851565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611bb357607f821691505b602082108103611bc657611bc5611b6f565b5b50919050565b7f5573657220616c726561647920726567697374657265642100000000000000005f82015250565b5f611c00601883611781565b9150611c0b82611bcc565b602082019050919050565b5f6020820190508181035f830152611c2d81611bf4565b9050919050565b611c3d81611981565b82525050565b5f606082019050611c565f830186611b09565b611c636020830185611c34565b611c706040830184611909565b949350505050565b7f43616e2774207769746864726177206d6f7265207468656e20796f75722062615f8201527f6c616e6365210000000000000000000000000000000000000000000000000000602082015250565b5f611cd2602683611781565b9150611cdd82611c78565b604082019050919050565b5f6020820190508181035f830152611cff81611cc6565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611d3d82611865565b9150611d4883611865565b9250828201905080821115611d6057611d5f611d06565b5b92915050565b7f5769746864726177696e6720616d6f756e74206265666f72652020756e6c6f635f8201527f6b206379636c6521000000000000000000000000000000000000000000000000602082015250565b5f611dc0602883611781565b9150611dcb82611d66565b604082019050919050565b5f6020820190508181035f830152611ded81611db4565b9050919050565b5f604082019050611e075f830185611b09565b611e146020830184611909565b9392505050565b5f611e2582611865565b9150611e3083611865565b9250828203905081811115611e4857611e47611d06565b5b92915050565b5f611e5882611865565b9150611e6383611865565b9250828202611e7181611865565b91508282048414831517611e8857611e87611d06565b5b5092915050565b5f81905092915050565b50565b5f611ea75f83611e8f565b9150611eb282611e99565b5f82019050919050565b5f611ec682611e9c565b9150819050919050565b7f4661696c656420746f20776974686472617721000000000000000000000000005f82015250565b5f611f04601383611781565b9150611f0f82611ed0565b602082019050919050565b5f6020820190508181035f830152611f3181611ef8565b9050919050565b7f5573657220616c72656164792064656661756c746564210000000000000000005f82015250565b5f611f6c601783611781565b9150611f7782611f38565b602082019050919050565b5f6020820190508181035f830152611f9981611f60565b9050919050565b7f44656661756c7465642075736572206372697465726961206e6f74206d6174635f8201527f6865642100000000000000000000000000000000000000000000000000000000602082015250565b5f611ffa602483611781565b915061200582611fa0565b604082019050919050565b5f6020820190508181035f83015261202781611fee565b9050919050565b5f6060820190506120415f830186611b09565b61204e6020830185611909565b61205b6040830184611c34565b949350505050565b7f416c726561647920696e697469616c69736564210000000000000000000000005f82015250565b5f612097601483611781565b91506120a282612063565b602082019050919050565b5f6020820190508181035f8301526120c48161208b565b9050919050565b7f596f752061726520706179696e67206d6f7265207468656e20796f75722064655f8201527f627420616d6f756e740000000000000000000000000000000000000000000000602082015250565b5f612125602983611781565b9150612130826120cb565b604082019050919050565b5f6020820190508181035f83015261215281612119565b9050919050565b7f4e6f7420656e6f75676820454e5420746f2070617920666565730000000000005f82015250565b5f61218d601a83611781565b915061219882612159565b602082019050919050565b5f6020820190508181035f8301526121ba81612181565b9050919050565b5f6060820190506121d45f830186611b09565b6121e16020830185611909565b6121ee6040830184611909565b949350505050565b7f436f7265206e6f7420696e7469616c69736564207965742100000000000000005f82015250565b5f61222a601883611781565b9150612235826121f6565b602082019050919050565b5f6020820190508181035f8301526122578161221e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61229582611865565b91506122a083611865565b9250826122b0576122af61225e565b5b82820490509291505056fea264697066735822122009400ed2057b5975b364f5939c8f75aeb565de305cb064c88a87dd13325aecbe64736f6c634300081a0033",
}

// EaseNowABI is the input ABI used to generate the binding from.
// Deprecated: Use EaseNowMetaData.ABI instead.
var EaseNowABI = EaseNowMetaData.ABI

// EaseNowBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EaseNowMetaData.Bin instead.
var EaseNowBin = EaseNowMetaData.Bin

// DeployEaseNow deploys a new Ethereum contract, binding an instance of EaseNow to it.
func DeployEaseNow(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EaseNow, error) {
	parsed, err := EaseNowMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EaseNowBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EaseNow{EaseNowCaller: EaseNowCaller{contract: contract}, EaseNowTransactor: EaseNowTransactor{contract: contract}, EaseNowFilterer: EaseNowFilterer{contract: contract}}, nil
}

// EaseNow is an auto generated Go binding around an Ethereum contract.
type EaseNow struct {
	EaseNowCaller     // Read-only binding to the contract
	EaseNowTransactor // Write-only binding to the contract
	EaseNowFilterer   // Log filterer for contract events
}

// EaseNowCaller is an auto generated read-only Go binding around an Ethereum contract.
type EaseNowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EaseNowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EaseNowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EaseNowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EaseNowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EaseNowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EaseNowSession struct {
	Contract     *EaseNow          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EaseNowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EaseNowCallerSession struct {
	Contract *EaseNowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EaseNowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EaseNowTransactorSession struct {
	Contract     *EaseNowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EaseNowRaw is an auto generated low-level Go binding around an Ethereum contract.
type EaseNowRaw struct {
	Contract *EaseNow // Generic contract binding to access the raw methods on
}

// EaseNowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EaseNowCallerRaw struct {
	Contract *EaseNowCaller // Generic read-only contract binding to access the raw methods on
}

// EaseNowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EaseNowTransactorRaw struct {
	Contract *EaseNowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEaseNow creates a new instance of EaseNow, bound to a specific deployed contract.
func NewEaseNow(address common.Address, backend bind.ContractBackend) (*EaseNow, error) {
	contract, err := bindEaseNow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EaseNow{EaseNowCaller: EaseNowCaller{contract: contract}, EaseNowTransactor: EaseNowTransactor{contract: contract}, EaseNowFilterer: EaseNowFilterer{contract: contract}}, nil
}

// NewEaseNowCaller creates a new read-only instance of EaseNow, bound to a specific deployed contract.
func NewEaseNowCaller(address common.Address, caller bind.ContractCaller) (*EaseNowCaller, error) {
	contract, err := bindEaseNow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EaseNowCaller{contract: contract}, nil
}

// NewEaseNowTransactor creates a new write-only instance of EaseNow, bound to a specific deployed contract.
func NewEaseNowTransactor(address common.Address, transactor bind.ContractTransactor) (*EaseNowTransactor, error) {
	contract, err := bindEaseNow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EaseNowTransactor{contract: contract}, nil
}

// NewEaseNowFilterer creates a new log filterer instance of EaseNow, bound to a specific deployed contract.
func NewEaseNowFilterer(address common.Address, filterer bind.ContractFilterer) (*EaseNowFilterer, error) {
	contract, err := bindEaseNow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EaseNowFilterer{contract: contract}, nil
}

// bindEaseNow binds a generic wrapper to an already deployed contract.
func bindEaseNow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EaseNowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EaseNow *EaseNowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EaseNow.Contract.EaseNowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EaseNow *EaseNowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EaseNow.Contract.EaseNowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EaseNow *EaseNowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EaseNow.Contract.EaseNowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EaseNow *EaseNowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EaseNow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EaseNow *EaseNowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EaseNow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EaseNow *EaseNowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EaseNow.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_EaseNow *EaseNowCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EaseNow.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_EaseNow *EaseNowSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _EaseNow.Contract.Allowance(&_EaseNow.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_EaseNow *EaseNowCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _EaseNow.Contract.Allowance(&_EaseNow.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EaseNow *EaseNowCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EaseNow.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EaseNow *EaseNowSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EaseNow.Contract.BalanceOf(&_EaseNow.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EaseNow *EaseNowCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EaseNow.Contract.BalanceOf(&_EaseNow.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EaseNow *EaseNowCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EaseNow.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EaseNow *EaseNowSession) Decimals() (uint8, error) {
	return _EaseNow.Contract.Decimals(&_EaseNow.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EaseNow *EaseNowCallerSession) Decimals() (uint8, error) {
	return _EaseNow.Contract.Decimals(&_EaseNow.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EaseNow *EaseNowCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EaseNow.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EaseNow *EaseNowSession) Name() (string, error) {
	return _EaseNow.Contract.Name(&_EaseNow.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EaseNow *EaseNowCallerSession) Name() (string, error) {
	return _EaseNow.Contract.Name(&_EaseNow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EaseNow *EaseNowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EaseNow.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EaseNow *EaseNowSession) Owner() (common.Address, error) {
	return _EaseNow.Contract.Owner(&_EaseNow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EaseNow *EaseNowCallerSession) Owner() (common.Address, error) {
	return _EaseNow.Contract.Owner(&_EaseNow.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EaseNow *EaseNowCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EaseNow.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EaseNow *EaseNowSession) Symbol() (string, error) {
	return _EaseNow.Contract.Symbol(&_EaseNow.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EaseNow *EaseNowCallerSession) Symbol() (string, error) {
	return _EaseNow.Contract.Symbol(&_EaseNow.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EaseNow *EaseNowCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EaseNow.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EaseNow *EaseNowSession) TotalSupply() (*big.Int, error) {
	return _EaseNow.Contract.TotalSupply(&_EaseNow.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EaseNow *EaseNowCallerSession) TotalSupply() (*big.Int, error) {
	return _EaseNow.Contract.TotalSupply(&_EaseNow.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_EaseNow *EaseNowTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _EaseNow.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_EaseNow *EaseNowSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _EaseNow.Contract.Approve(&_EaseNow.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_EaseNow *EaseNowTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _EaseNow.Contract.Approve(&_EaseNow.TransactOpts, spender, value)
}

// Borrow is a paid mutator transaction binding the contract method 0x7be0732e.
//
// Solidity: function borrow(uint256 amount, bytes32 purchaseData, address user) returns()
func (_EaseNow *EaseNowTransactor) Borrow(opts *bind.TransactOpts, amount *big.Int, purchaseData [32]byte, user common.Address) (*types.Transaction, error) {
	return _EaseNow.contract.Transact(opts, "borrow", amount, purchaseData, user)
}

// Borrow is a paid mutator transaction binding the contract method 0x7be0732e.
//
// Solidity: function borrow(uint256 amount, bytes32 purchaseData, address user) returns()
func (_EaseNow *EaseNowSession) Borrow(amount *big.Int, purchaseData [32]byte, user common.Address) (*types.Transaction, error) {
	return _EaseNow.Contract.Borrow(&_EaseNow.TransactOpts, amount, purchaseData, user)
}

// Borrow is a paid mutator transaction binding the contract method 0x7be0732e.
//
// Solidity: function borrow(uint256 amount, bytes32 purchaseData, address user) returns()
func (_EaseNow *EaseNowTransactorSession) Borrow(amount *big.Int, purchaseData [32]byte, user common.Address) (*types.Transaction, error) {
	return _EaseNow.Contract.Borrow(&_EaseNow.TransactOpts, amount, purchaseData, user)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_EaseNow *EaseNowTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EaseNow.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_EaseNow *EaseNowSession) Deposit() (*types.Transaction, error) {
	return _EaseNow.Contract.Deposit(&_EaseNow.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_EaseNow *EaseNowTransactorSession) Deposit() (*types.Transaction, error) {
	return _EaseNow.Contract.Deposit(&_EaseNow.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 seedProof_) returns()
func (_EaseNow *EaseNowTransactor) Init(opts *bind.TransactOpts, seedProof_ [32]byte) (*types.Transaction, error) {
	return _EaseNow.contract.Transact(opts, "init", seedProof_)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 seedProof_) returns()
func (_EaseNow *EaseNowSession) Init(seedProof_ [32]byte) (*types.Transaction, error) {
	return _EaseNow.Contract.Init(&_EaseNow.TransactOpts, seedProof_)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 seedProof_) returns()
func (_EaseNow *EaseNowTransactorSession) Init(seedProof_ [32]byte) (*types.Transaction, error) {
	return _EaseNow.Contract.Init(&_EaseNow.TransactOpts, seedProof_)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x2caabfef.
//
// Solidity: function registerUser(address userAddress, uint256 creditLimit, bytes32 privateData) returns()
func (_EaseNow *EaseNowTransactor) RegisterUser(opts *bind.TransactOpts, userAddress common.Address, creditLimit *big.Int, privateData [32]byte) (*types.Transaction, error) {
	return _EaseNow.contract.Transact(opts, "registerUser", userAddress, creditLimit, privateData)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x2caabfef.
//
// Solidity: function registerUser(address userAddress, uint256 creditLimit, bytes32 privateData) returns()
func (_EaseNow *EaseNowSession) RegisterUser(userAddress common.Address, creditLimit *big.Int, privateData [32]byte) (*types.Transaction, error) {
	return _EaseNow.Contract.RegisterUser(&_EaseNow.TransactOpts, userAddress, creditLimit, privateData)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x2caabfef.
//
// Solidity: function registerUser(address userAddress, uint256 creditLimit, bytes32 privateData) returns()
func (_EaseNow *EaseNowTransactorSession) RegisterUser(userAddress common.Address, creditLimit *big.Int, privateData [32]byte) (*types.Transaction, error) {
	return _EaseNow.Contract.RegisterUser(&_EaseNow.TransactOpts, userAddress, creditLimit, privateData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EaseNow *EaseNowTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EaseNow.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EaseNow *EaseNowSession) RenounceOwnership() (*types.Transaction, error) {
	return _EaseNow.Contract.RenounceOwnership(&_EaseNow.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EaseNow *EaseNowTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EaseNow.Contract.RenounceOwnership(&_EaseNow.TransactOpts)
}

// Repay is a paid mutator transaction binding the contract method 0x402d8883.
//
// Solidity: function repay() payable returns()
func (_EaseNow *EaseNowTransactor) Repay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EaseNow.contract.Transact(opts, "repay")
}

// Repay is a paid mutator transaction binding the contract method 0x402d8883.
//
// Solidity: function repay() payable returns()
func (_EaseNow *EaseNowSession) Repay() (*types.Transaction, error) {
	return _EaseNow.Contract.Repay(&_EaseNow.TransactOpts)
}

// Repay is a paid mutator transaction binding the contract method 0x402d8883.
//
// Solidity: function repay() payable returns()
func (_EaseNow *EaseNowTransactorSession) Repay() (*types.Transaction, error) {
	return _EaseNow.Contract.Repay(&_EaseNow.TransactOpts)
}

// ReportDefault is a paid mutator transaction binding the contract method 0x3a7d5115.
//
// Solidity: function reportDefault(address userAddress) returns()
func (_EaseNow *EaseNowTransactor) ReportDefault(opts *bind.TransactOpts, userAddress common.Address) (*types.Transaction, error) {
	return _EaseNow.contract.Transact(opts, "reportDefault", userAddress)
}

// ReportDefault is a paid mutator transaction binding the contract method 0x3a7d5115.
//
// Solidity: function reportDefault(address userAddress) returns()
func (_EaseNow *EaseNowSession) ReportDefault(userAddress common.Address) (*types.Transaction, error) {
	return _EaseNow.Contract.ReportDefault(&_EaseNow.TransactOpts, userAddress)
}

// ReportDefault is a paid mutator transaction binding the contract method 0x3a7d5115.
//
// Solidity: function reportDefault(address userAddress) returns()
func (_EaseNow *EaseNowTransactorSession) ReportDefault(userAddress common.Address) (*types.Transaction, error) {
	return _EaseNow.Contract.ReportDefault(&_EaseNow.TransactOpts, userAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_EaseNow *EaseNowTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EaseNow.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_EaseNow *EaseNowSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EaseNow.Contract.Transfer(&_EaseNow.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_EaseNow *EaseNowTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EaseNow.Contract.Transfer(&_EaseNow.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_EaseNow *EaseNowTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EaseNow.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_EaseNow *EaseNowSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EaseNow.Contract.TransferFrom(&_EaseNow.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_EaseNow *EaseNowTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _EaseNow.Contract.TransferFrom(&_EaseNow.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EaseNow *EaseNowTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EaseNow.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EaseNow *EaseNowSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EaseNow.Contract.TransferOwnership(&_EaseNow.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EaseNow *EaseNowTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EaseNow.Contract.TransferOwnership(&_EaseNow.TransactOpts, newOwner)
}

// UpdatePriceRatio is a paid mutator transaction binding the contract method 0xd8bd9260.
//
// Solidity: function updatePriceRatio(uint256 newPriceRatio) returns()
func (_EaseNow *EaseNowTransactor) UpdatePriceRatio(opts *bind.TransactOpts, newPriceRatio *big.Int) (*types.Transaction, error) {
	return _EaseNow.contract.Transact(opts, "updatePriceRatio", newPriceRatio)
}

// UpdatePriceRatio is a paid mutator transaction binding the contract method 0xd8bd9260.
//
// Solidity: function updatePriceRatio(uint256 newPriceRatio) returns()
func (_EaseNow *EaseNowSession) UpdatePriceRatio(newPriceRatio *big.Int) (*types.Transaction, error) {
	return _EaseNow.Contract.UpdatePriceRatio(&_EaseNow.TransactOpts, newPriceRatio)
}

// UpdatePriceRatio is a paid mutator transaction binding the contract method 0xd8bd9260.
//
// Solidity: function updatePriceRatio(uint256 newPriceRatio) returns()
func (_EaseNow *EaseNowTransactorSession) UpdatePriceRatio(newPriceRatio *big.Int) (*types.Transaction, error) {
	return _EaseNow.Contract.UpdatePriceRatio(&_EaseNow.TransactOpts, newPriceRatio)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_EaseNow *EaseNowTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _EaseNow.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_EaseNow *EaseNowSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _EaseNow.Contract.Withdraw(&_EaseNow.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_EaseNow *EaseNowTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _EaseNow.Contract.Withdraw(&_EaseNow.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EaseNow *EaseNowTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EaseNow.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EaseNow *EaseNowSession) Receive() (*types.Transaction, error) {
	return _EaseNow.Contract.Receive(&_EaseNow.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EaseNow *EaseNowTransactorSession) Receive() (*types.Transaction, error) {
	return _EaseNow.Contract.Receive(&_EaseNow.TransactOpts)
}

// EaseNowAmountRepaidIterator is returned from FilterAmountRepaid and is used to iterate over the raw logs and unpacked data for AmountRepaid events raised by the EaseNow contract.
type EaseNowAmountRepaidIterator struct {
	Event *EaseNowAmountRepaid // Event containing the contract specifics and raw log

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
func (it *EaseNowAmountRepaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EaseNowAmountRepaid)
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
		it.Event = new(EaseNowAmountRepaid)
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
func (it *EaseNowAmountRepaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EaseNowAmountRepaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EaseNowAmountRepaid represents a AmountRepaid event raised by the EaseNow contract.
type EaseNowAmountRepaid struct {
	UserAddress common.Address
	Amount      *big.Int
	RemaingDebt *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAmountRepaid is a free log retrieval operation binding the contract event 0x0f2afdea007852c85c572fd5ec20aede9297fb7614b1cbcd5594de764d3328ff.
//
// Solidity: event AmountRepaid(address userAddress, uint256 amount, uint256 remaingDebt)
func (_EaseNow *EaseNowFilterer) FilterAmountRepaid(opts *bind.FilterOpts) (*EaseNowAmountRepaidIterator, error) {

	logs, sub, err := _EaseNow.contract.FilterLogs(opts, "AmountRepaid")
	if err != nil {
		return nil, err
	}
	return &EaseNowAmountRepaidIterator{contract: _EaseNow.contract, event: "AmountRepaid", logs: logs, sub: sub}, nil
}

// WatchAmountRepaid is a free log subscription operation binding the contract event 0x0f2afdea007852c85c572fd5ec20aede9297fb7614b1cbcd5594de764d3328ff.
//
// Solidity: event AmountRepaid(address userAddress, uint256 amount, uint256 remaingDebt)
func (_EaseNow *EaseNowFilterer) WatchAmountRepaid(opts *bind.WatchOpts, sink chan<- *EaseNowAmountRepaid) (event.Subscription, error) {

	logs, sub, err := _EaseNow.contract.WatchLogs(opts, "AmountRepaid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EaseNowAmountRepaid)
				if err := _EaseNow.contract.UnpackLog(event, "AmountRepaid", log); err != nil {
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

// ParseAmountRepaid is a log parse operation binding the contract event 0x0f2afdea007852c85c572fd5ec20aede9297fb7614b1cbcd5594de764d3328ff.
//
// Solidity: event AmountRepaid(address userAddress, uint256 amount, uint256 remaingDebt)
func (_EaseNow *EaseNowFilterer) ParseAmountRepaid(log types.Log) (*EaseNowAmountRepaid, error) {
	event := new(EaseNowAmountRepaid)
	if err := _EaseNow.contract.UnpackLog(event, "AmountRepaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EaseNowApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EaseNow contract.
type EaseNowApprovalIterator struct {
	Event *EaseNowApproval // Event containing the contract specifics and raw log

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
func (it *EaseNowApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EaseNowApproval)
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
		it.Event = new(EaseNowApproval)
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
func (it *EaseNowApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EaseNowApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EaseNowApproval represents a Approval event raised by the EaseNow contract.
type EaseNowApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EaseNow *EaseNowFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EaseNowApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EaseNow.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EaseNowApprovalIterator{contract: _EaseNow.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EaseNow *EaseNowFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EaseNowApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EaseNow.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EaseNowApproval)
				if err := _EaseNow.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EaseNow *EaseNowFilterer) ParseApproval(log types.Log) (*EaseNowApproval, error) {
	event := new(EaseNowApproval)
	if err := _EaseNow.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EaseNowLenderDepositIterator is returned from FilterLenderDeposit and is used to iterate over the raw logs and unpacked data for LenderDeposit events raised by the EaseNow contract.
type EaseNowLenderDepositIterator struct {
	Event *EaseNowLenderDeposit // Event containing the contract specifics and raw log

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
func (it *EaseNowLenderDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EaseNowLenderDeposit)
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
		it.Event = new(EaseNowLenderDeposit)
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
func (it *EaseNowLenderDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EaseNowLenderDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EaseNowLenderDeposit represents a LenderDeposit event raised by the EaseNow contract.
type EaseNowLenderDeposit struct {
	Lender   common.Address
	Amount   *big.Int
	Locktime *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLenderDeposit is a free log retrieval operation binding the contract event 0x4d0e8513a525a0e8d883d0fb5bf24c215136a9a954918fbce6d9f03d9bfacd13.
//
// Solidity: event LenderDeposit(address lender, uint256 amount, uint256 locktime)
func (_EaseNow *EaseNowFilterer) FilterLenderDeposit(opts *bind.FilterOpts) (*EaseNowLenderDepositIterator, error) {

	logs, sub, err := _EaseNow.contract.FilterLogs(opts, "LenderDeposit")
	if err != nil {
		return nil, err
	}
	return &EaseNowLenderDepositIterator{contract: _EaseNow.contract, event: "LenderDeposit", logs: logs, sub: sub}, nil
}

// WatchLenderDeposit is a free log subscription operation binding the contract event 0x4d0e8513a525a0e8d883d0fb5bf24c215136a9a954918fbce6d9f03d9bfacd13.
//
// Solidity: event LenderDeposit(address lender, uint256 amount, uint256 locktime)
func (_EaseNow *EaseNowFilterer) WatchLenderDeposit(opts *bind.WatchOpts, sink chan<- *EaseNowLenderDeposit) (event.Subscription, error) {

	logs, sub, err := _EaseNow.contract.WatchLogs(opts, "LenderDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EaseNowLenderDeposit)
				if err := _EaseNow.contract.UnpackLog(event, "LenderDeposit", log); err != nil {
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

// ParseLenderDeposit is a log parse operation binding the contract event 0x4d0e8513a525a0e8d883d0fb5bf24c215136a9a954918fbce6d9f03d9bfacd13.
//
// Solidity: event LenderDeposit(address lender, uint256 amount, uint256 locktime)
func (_EaseNow *EaseNowFilterer) ParseLenderDeposit(log types.Log) (*EaseNowLenderDeposit, error) {
	event := new(EaseNowLenderDeposit)
	if err := _EaseNow.contract.UnpackLog(event, "LenderDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EaseNowLenderRewardedIterator is returned from FilterLenderRewarded and is used to iterate over the raw logs and unpacked data for LenderRewarded events raised by the EaseNow contract.
type EaseNowLenderRewardedIterator struct {
	Event *EaseNowLenderRewarded // Event containing the contract specifics and raw log

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
func (it *EaseNowLenderRewardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EaseNowLenderRewarded)
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
		it.Event = new(EaseNowLenderRewarded)
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
func (it *EaseNowLenderRewardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EaseNowLenderRewardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EaseNowLenderRewarded represents a LenderRewarded event raised by the EaseNow contract.
type EaseNowLenderRewarded struct {
	Lender       common.Address
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLenderRewarded is a free log retrieval operation binding the contract event 0x38af265c68180096307acf94a0034cafb4c264beb9c216b7f8b8776b57e82110.
//
// Solidity: event LenderRewarded(address lender, uint256 rewardAmount)
func (_EaseNow *EaseNowFilterer) FilterLenderRewarded(opts *bind.FilterOpts) (*EaseNowLenderRewardedIterator, error) {

	logs, sub, err := _EaseNow.contract.FilterLogs(opts, "LenderRewarded")
	if err != nil {
		return nil, err
	}
	return &EaseNowLenderRewardedIterator{contract: _EaseNow.contract, event: "LenderRewarded", logs: logs, sub: sub}, nil
}

// WatchLenderRewarded is a free log subscription operation binding the contract event 0x38af265c68180096307acf94a0034cafb4c264beb9c216b7f8b8776b57e82110.
//
// Solidity: event LenderRewarded(address lender, uint256 rewardAmount)
func (_EaseNow *EaseNowFilterer) WatchLenderRewarded(opts *bind.WatchOpts, sink chan<- *EaseNowLenderRewarded) (event.Subscription, error) {

	logs, sub, err := _EaseNow.contract.WatchLogs(opts, "LenderRewarded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EaseNowLenderRewarded)
				if err := _EaseNow.contract.UnpackLog(event, "LenderRewarded", log); err != nil {
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

// ParseLenderRewarded is a log parse operation binding the contract event 0x38af265c68180096307acf94a0034cafb4c264beb9c216b7f8b8776b57e82110.
//
// Solidity: event LenderRewarded(address lender, uint256 rewardAmount)
func (_EaseNow *EaseNowFilterer) ParseLenderRewarded(log types.Log) (*EaseNowLenderRewarded, error) {
	event := new(EaseNowLenderRewarded)
	if err := _EaseNow.contract.UnpackLog(event, "LenderRewarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EaseNowLenderWithdrawIterator is returned from FilterLenderWithdraw and is used to iterate over the raw logs and unpacked data for LenderWithdraw events raised by the EaseNow contract.
type EaseNowLenderWithdrawIterator struct {
	Event *EaseNowLenderWithdraw // Event containing the contract specifics and raw log

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
func (it *EaseNowLenderWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EaseNowLenderWithdraw)
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
		it.Event = new(EaseNowLenderWithdraw)
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
func (it *EaseNowLenderWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EaseNowLenderWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EaseNowLenderWithdraw represents a LenderWithdraw event raised by the EaseNow contract.
type EaseNowLenderWithdraw struct {
	Lender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLenderWithdraw is a free log retrieval operation binding the contract event 0xf1efcc5bb6f7c762ce11c94881410b55191722a008338c0c9a88203c74604aea.
//
// Solidity: event LenderWithdraw(address lender, uint256 amount)
func (_EaseNow *EaseNowFilterer) FilterLenderWithdraw(opts *bind.FilterOpts) (*EaseNowLenderWithdrawIterator, error) {

	logs, sub, err := _EaseNow.contract.FilterLogs(opts, "LenderWithdraw")
	if err != nil {
		return nil, err
	}
	return &EaseNowLenderWithdrawIterator{contract: _EaseNow.contract, event: "LenderWithdraw", logs: logs, sub: sub}, nil
}

// WatchLenderWithdraw is a free log subscription operation binding the contract event 0xf1efcc5bb6f7c762ce11c94881410b55191722a008338c0c9a88203c74604aea.
//
// Solidity: event LenderWithdraw(address lender, uint256 amount)
func (_EaseNow *EaseNowFilterer) WatchLenderWithdraw(opts *bind.WatchOpts, sink chan<- *EaseNowLenderWithdraw) (event.Subscription, error) {

	logs, sub, err := _EaseNow.contract.WatchLogs(opts, "LenderWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EaseNowLenderWithdraw)
				if err := _EaseNow.contract.UnpackLog(event, "LenderWithdraw", log); err != nil {
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

// ParseLenderWithdraw is a log parse operation binding the contract event 0xf1efcc5bb6f7c762ce11c94881410b55191722a008338c0c9a88203c74604aea.
//
// Solidity: event LenderWithdraw(address lender, uint256 amount)
func (_EaseNow *EaseNowFilterer) ParseLenderWithdraw(log types.Log) (*EaseNowLenderWithdraw, error) {
	event := new(EaseNowLenderWithdraw)
	if err := _EaseNow.contract.UnpackLog(event, "LenderWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EaseNowOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EaseNow contract.
type EaseNowOwnershipTransferredIterator struct {
	Event *EaseNowOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EaseNowOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EaseNowOwnershipTransferred)
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
		it.Event = new(EaseNowOwnershipTransferred)
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
func (it *EaseNowOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EaseNowOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EaseNowOwnershipTransferred represents a OwnershipTransferred event raised by the EaseNow contract.
type EaseNowOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EaseNow *EaseNowFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EaseNowOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EaseNow.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EaseNowOwnershipTransferredIterator{contract: _EaseNow.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EaseNow *EaseNowFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EaseNowOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EaseNow.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EaseNowOwnershipTransferred)
				if err := _EaseNow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EaseNow *EaseNowFilterer) ParseOwnershipTransferred(log types.Log) (*EaseNowOwnershipTransferred, error) {
	event := new(EaseNowOwnershipTransferred)
	if err := _EaseNow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EaseNowTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EaseNow contract.
type EaseNowTransferIterator struct {
	Event *EaseNowTransfer // Event containing the contract specifics and raw log

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
func (it *EaseNowTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EaseNowTransfer)
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
		it.Event = new(EaseNowTransfer)
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
func (it *EaseNowTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EaseNowTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EaseNowTransfer represents a Transfer event raised by the EaseNow contract.
type EaseNowTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EaseNow *EaseNowFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EaseNowTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EaseNow.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EaseNowTransferIterator{contract: _EaseNow.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EaseNow *EaseNowFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EaseNowTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EaseNow.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EaseNowTransfer)
				if err := _EaseNow.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EaseNow *EaseNowFilterer) ParseTransfer(log types.Log) (*EaseNowTransfer, error) {
	event := new(EaseNowTransfer)
	if err := _EaseNow.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EaseNowUserDefaultedIterator is returned from FilterUserDefaulted and is used to iterate over the raw logs and unpacked data for UserDefaulted events raised by the EaseNow contract.
type EaseNowUserDefaultedIterator struct {
	Event *EaseNowUserDefaulted // Event containing the contract specifics and raw log

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
func (it *EaseNowUserDefaultedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EaseNowUserDefaulted)
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
		it.Event = new(EaseNowUserDefaulted)
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
func (it *EaseNowUserDefaultedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EaseNowUserDefaultedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EaseNowUserDefaulted represents a UserDefaulted event raised by the EaseNow contract.
type EaseNowUserDefaulted struct {
	UserAddress   common.Address
	DefaultAmount *big.Int
	PrivareData   [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserDefaulted is a free log retrieval operation binding the contract event 0xf08fd0dc9b5a0e09baf039a4d7a8897a6ab538c6de1e8e4c40c82560090b3875.
//
// Solidity: event UserDefaulted(address userAddress, uint256 defaultAmount, bytes32 privareData)
func (_EaseNow *EaseNowFilterer) FilterUserDefaulted(opts *bind.FilterOpts) (*EaseNowUserDefaultedIterator, error) {

	logs, sub, err := _EaseNow.contract.FilterLogs(opts, "UserDefaulted")
	if err != nil {
		return nil, err
	}
	return &EaseNowUserDefaultedIterator{contract: _EaseNow.contract, event: "UserDefaulted", logs: logs, sub: sub}, nil
}

// WatchUserDefaulted is a free log subscription operation binding the contract event 0xf08fd0dc9b5a0e09baf039a4d7a8897a6ab538c6de1e8e4c40c82560090b3875.
//
// Solidity: event UserDefaulted(address userAddress, uint256 defaultAmount, bytes32 privareData)
func (_EaseNow *EaseNowFilterer) WatchUserDefaulted(opts *bind.WatchOpts, sink chan<- *EaseNowUserDefaulted) (event.Subscription, error) {

	logs, sub, err := _EaseNow.contract.WatchLogs(opts, "UserDefaulted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EaseNowUserDefaulted)
				if err := _EaseNow.contract.UnpackLog(event, "UserDefaulted", log); err != nil {
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

// ParseUserDefaulted is a log parse operation binding the contract event 0xf08fd0dc9b5a0e09baf039a4d7a8897a6ab538c6de1e8e4c40c82560090b3875.
//
// Solidity: event UserDefaulted(address userAddress, uint256 defaultAmount, bytes32 privareData)
func (_EaseNow *EaseNowFilterer) ParseUserDefaulted(log types.Log) (*EaseNowUserDefaulted, error) {
	event := new(EaseNowUserDefaulted)
	if err := _EaseNow.contract.UnpackLog(event, "UserDefaulted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EaseNowUserRegistredIterator is returned from FilterUserRegistred and is used to iterate over the raw logs and unpacked data for UserRegistred events raised by the EaseNow contract.
type EaseNowUserRegistredIterator struct {
	Event *EaseNowUserRegistred // Event containing the contract specifics and raw log

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
func (it *EaseNowUserRegistredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EaseNowUserRegistred)
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
		it.Event = new(EaseNowUserRegistred)
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
func (it *EaseNowUserRegistredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EaseNowUserRegistredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EaseNowUserRegistred represents a UserRegistred event raised by the EaseNow contract.
type EaseNowUserRegistred struct {
	UserAddress common.Address
	PrivateData [32]byte
	CreditLimit *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUserRegistred is a free log retrieval operation binding the contract event 0x7338f2f0fce044e8ff9f8257e4426c8912aa5dc780cbf77f446d4b10dd11a7e0.
//
// Solidity: event UserRegistred(address userAddress, bytes32 privateData, uint256 creditLimit)
func (_EaseNow *EaseNowFilterer) FilterUserRegistred(opts *bind.FilterOpts) (*EaseNowUserRegistredIterator, error) {

	logs, sub, err := _EaseNow.contract.FilterLogs(opts, "UserRegistred")
	if err != nil {
		return nil, err
	}
	return &EaseNowUserRegistredIterator{contract: _EaseNow.contract, event: "UserRegistred", logs: logs, sub: sub}, nil
}

// WatchUserRegistred is a free log subscription operation binding the contract event 0x7338f2f0fce044e8ff9f8257e4426c8912aa5dc780cbf77f446d4b10dd11a7e0.
//
// Solidity: event UserRegistred(address userAddress, bytes32 privateData, uint256 creditLimit)
func (_EaseNow *EaseNowFilterer) WatchUserRegistred(opts *bind.WatchOpts, sink chan<- *EaseNowUserRegistred) (event.Subscription, error) {

	logs, sub, err := _EaseNow.contract.WatchLogs(opts, "UserRegistred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EaseNowUserRegistred)
				if err := _EaseNow.contract.UnpackLog(event, "UserRegistred", log); err != nil {
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

// ParseUserRegistred is a log parse operation binding the contract event 0x7338f2f0fce044e8ff9f8257e4426c8912aa5dc780cbf77f446d4b10dd11a7e0.
//
// Solidity: event UserRegistred(address userAddress, bytes32 privateData, uint256 creditLimit)
func (_EaseNow *EaseNowFilterer) ParseUserRegistred(log types.Log) (*EaseNowUserRegistred, error) {
	event := new(EaseNowUserRegistred)
	if err := _EaseNow.contract.UnpackLog(event, "UserRegistred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
