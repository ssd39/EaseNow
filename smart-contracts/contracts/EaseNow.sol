// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract EaseNow is ERC20, Ownable {
    uint256 initialSupply = 1000 ether;
    string public seedProof;
    //gets updated by TEE seed server at every minuite else if price ratio changed  more then threshold amount
    // ENT/ETH
    uint256 public priceRatio = 100;

    // in this test version repayment cycle / locktime is 2min to easily test things out
    // in real world it can be 2weeks / 1month etc
    struct Borrower {
        uint256 creditLimit;
        bytes32 privateData;
        uint256 debt;
        uint256 lastPaymentTime;
        bool isDefaulted;
    }

    struct Lender {
        uint256 lockTime;
        uint256 amount;
    }

    mapping(address => Borrower) borrowerMap;
    mapping(address => Lender) lenderMap;

    event LenderWithdraw(address indexed lender, uint256 amount);
    event LenderDeposit(
        address indexed lender,
        uint256 amount,
        uint256 locktime
    );
    event LenderRewarded(address indexed lender, uint256 rewardAmount);
    event UserRegistred(
        address indexed userAddress,
        bytes32 privateData,
        uint256 creditLimit
    );
    event AmountRepaid(
        address indexed userAddress,
        uint256 amount,
        uint256 remaingDebt
    );
    event UserDefaulted(
        address indexed userAddress,
        uint256 defaultAmount,
        bytes32 privareData
    );
    event AmountBorrowed(
        address indexed userAddress,
        uint256 amount,
        uint256 remainingLimit,
        address merchent,
        bool isContract
    );

    constructor() Ownable(msg.sender) ERC20("EaseToken", "ENT") {
        _mint(msg.sender, initialSupply);
    }

    receive() external payable {}

    function init(string memory seedProof_) external {
        require(bytes(seedProof).length == 0, "Already initialised!");
        seedProof = seedProof_;
        _transferOwnership(msg.sender);
    }

    function faucet(address account) external onlyOwner {
        //For mvp and testnet only
        _mint(account, 100 ether);
    }

    function updatePriceRatio(uint256 newPriceRatio) external onlyOwner {
        require(bytes(seedProof).length != 0, "Core not intialised yet!");
        priceRatio = newPriceRatio;
    }

    function deposit() external payable {
        Lender storage lender = lenderMap[msg.sender];
        // first settle down unpaid rewards
        uint256 reward = getRewardAmount(
            lender.lockTime,
            lender.amount,
            block.number,
            priceRatio
        );
        if (reward != 0) {
            _mint(msg.sender, reward);
            emit LenderRewarded(msg.sender, reward);
        }
        lender.amount += msg.value;
        lender.lockTime = block.number;
        emit LenderDeposit(msg.sender, msg.value, lender.lockTime);
    }

    function withdraw(uint256 amount) external {
        Lender storage lender = lenderMap[msg.sender];
        require(
            lender.amount >= amount,
            "Can't withdraw more then your balance!"
        );
        require(
            lender.lockTime + 60 <= block.number,
            "Withdrawing amount before  unlock cycle!"
        );
        // first settle down unpaid rewards
        uint256 reward = getRewardAmount(
            lender.lockTime,
            lender.amount,
            block.number,
            priceRatio
        );
        if (reward != 0) {
            _mint(msg.sender, reward);
            emit LenderRewarded(msg.sender, reward);
        }
        lender.amount -= amount;
        emit LenderWithdraw(msg.sender, amount);
        if (address(this).balance < amount) {
            uint256 missingAmount = amount - address(this).balance;
            _mint(msg.sender, missingAmount * priceRatio);
            amount = address(this).balance;
        }
        if (amount > 0) {
            (bool sent, bytes memory data) = msg.sender.call{value: amount}("");
            require(sent, "Failed to withdraw!");
        }
    }

    function borrow(
        uint256 amount,
        address userAddress,
        string memory signature,
        address merchent,
        bool isContract,
        bytes memory contractCalldata
    ) external {
        Borrower storage borrower = borrowerMap[userAddress];
        require(borrower.privateData != bytes32(0), "User not found!");
        require(
            borrower.creditLimit - borrower.debt >= amount,
            "Spending more then your credit limit!"
        );
        require(!borrower.isDefaulted, "Your account is defaulted!");

        //TODO: verfy the signature that its really core authenticated merchent and transaction data
        borrower.debt += amount;
        emit AmountBorrowed(
            userAddress,
            amount,
            borrower.creditLimit - borrower.debt,
            merchent,
            isContract
        );
        if (!isContract) {
            (bool sent, bytes memory data) = merchent.call{value: amount}("");
            require(sent, "Failed to send Ether");
        } else {
            (bool success, bytes memory data) = merchent.call{
                value: amount
            }(contractCalldata);
            require(success, "Transaction to merchent contract failed!");
        }
    }

    function repay() external payable {
        Borrower storage borrower = borrowerMap[msg.sender];
        require(
            msg.value >= borrower.debt,
            "You are paying more then your debt amount"
        );
        uint256 fees = getRewardAmount(0, msg.value, 60, priceRatio);
        require(balanceOf(msg.sender) >= fees, "Not enough ENT to pay fees");
        _burn(msg.sender, fees);
        borrower.debt -= msg.value;
        emit AmountRepaid(msg.sender, msg.value, borrower.debt);
    }

    function reportDefault(address userAddress) external {
        Borrower storage borrower = borrowerMap[msg.sender];
        require(!borrower.isDefaulted, "User already defaulted!");
        require(
            borrower.debt > 0 && borrower.lastPaymentTime + 90 <= block.number,
            "Defaulted user criteria not matched!"
        );
        borrower.isDefaulted = true;
        emit UserDefaulted(userAddress, borrower.debt, borrower.privateData);
    }

    function registerUser(
        address userAddress,
        uint256 creditLimit,
        bytes32 privateData
    ) external onlyOwner {
        Borrower storage borrower = borrowerMap[userAddress];
        require(borrower.privateData == bytes32(0), "User already registered!");
        borrower.privateData = privateData;
        borrower.creditLimit = creditLimit;
        emit UserRegistred(userAddress, privateData, creditLimit);
    }

    function getRewardAmount(
        uint256 lockTime,
        uint256 amount,
        uint256 currentBlock,
        uint256 currentPrice
    ) private pure returns (uint256) {
        // currently charges are fixed to 1.5% of amount but in future it will auto adjust based on treasury
        // for example if some one defaults something to recover it smart contract will auto adjust
        return
            (((amount * currentPrice * 15) / 1000) *
                (currentBlock - lockTime)) / 60;
    }
}
