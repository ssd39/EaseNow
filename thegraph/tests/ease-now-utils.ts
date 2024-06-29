import { newMockEvent } from "matchstick-as"
import { ethereum, Address, BigInt, Bytes } from "@graphprotocol/graph-ts"
import {
  AmountBorrowed,
  AmountRepaid,
  Approval,
  LenderDeposit,
  LenderRewarded,
  LenderWithdraw,
  OwnershipTransferred,
  Transfer,
  UserDefaulted,
  UserRegistred
} from "../generated/EaseNow/EaseNow"

export function createAmountBorrowedEvent(
  userAddress: Address,
  amount: BigInt,
  remainingLimit: BigInt,
  merchent: Address,
  isContract: boolean
): AmountBorrowed {
  let amountBorrowedEvent = changetype<AmountBorrowed>(newMockEvent())

  amountBorrowedEvent.parameters = new Array()

  amountBorrowedEvent.parameters.push(
    new ethereum.EventParam(
      "userAddress",
      ethereum.Value.fromAddress(userAddress)
    )
  )
  amountBorrowedEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )
  amountBorrowedEvent.parameters.push(
    new ethereum.EventParam(
      "remainingLimit",
      ethereum.Value.fromUnsignedBigInt(remainingLimit)
    )
  )
  amountBorrowedEvent.parameters.push(
    new ethereum.EventParam("merchent", ethereum.Value.fromAddress(merchent))
  )
  amountBorrowedEvent.parameters.push(
    new ethereum.EventParam(
      "isContract",
      ethereum.Value.fromBoolean(isContract)
    )
  )

  return amountBorrowedEvent
}

export function createAmountRepaidEvent(
  userAddress: Address,
  amount: BigInt,
  remaingDebt: BigInt
): AmountRepaid {
  let amountRepaidEvent = changetype<AmountRepaid>(newMockEvent())

  amountRepaidEvent.parameters = new Array()

  amountRepaidEvent.parameters.push(
    new ethereum.EventParam(
      "userAddress",
      ethereum.Value.fromAddress(userAddress)
    )
  )
  amountRepaidEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )
  amountRepaidEvent.parameters.push(
    new ethereum.EventParam(
      "remaingDebt",
      ethereum.Value.fromUnsignedBigInt(remaingDebt)
    )
  )

  return amountRepaidEvent
}

export function createApprovalEvent(
  owner: Address,
  spender: Address,
  value: BigInt
): Approval {
  let approvalEvent = changetype<Approval>(newMockEvent())

  approvalEvent.parameters = new Array()

  approvalEvent.parameters.push(
    new ethereum.EventParam("owner", ethereum.Value.fromAddress(owner))
  )
  approvalEvent.parameters.push(
    new ethereum.EventParam("spender", ethereum.Value.fromAddress(spender))
  )
  approvalEvent.parameters.push(
    new ethereum.EventParam("value", ethereum.Value.fromUnsignedBigInt(value))
  )

  return approvalEvent
}

export function createLenderDepositEvent(
  lender: Address,
  amount: BigInt,
  locktime: BigInt
): LenderDeposit {
  let lenderDepositEvent = changetype<LenderDeposit>(newMockEvent())

  lenderDepositEvent.parameters = new Array()

  lenderDepositEvent.parameters.push(
    new ethereum.EventParam("lender", ethereum.Value.fromAddress(lender))
  )
  lenderDepositEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )
  lenderDepositEvent.parameters.push(
    new ethereum.EventParam(
      "locktime",
      ethereum.Value.fromUnsignedBigInt(locktime)
    )
  )

  return lenderDepositEvent
}

export function createLenderRewardedEvent(
  lender: Address,
  rewardAmount: BigInt
): LenderRewarded {
  let lenderRewardedEvent = changetype<LenderRewarded>(newMockEvent())

  lenderRewardedEvent.parameters = new Array()

  lenderRewardedEvent.parameters.push(
    new ethereum.EventParam("lender", ethereum.Value.fromAddress(lender))
  )
  lenderRewardedEvent.parameters.push(
    new ethereum.EventParam(
      "rewardAmount",
      ethereum.Value.fromUnsignedBigInt(rewardAmount)
    )
  )

  return lenderRewardedEvent
}

export function createLenderWithdrawEvent(
  lender: Address,
  amount: BigInt
): LenderWithdraw {
  let lenderWithdrawEvent = changetype<LenderWithdraw>(newMockEvent())

  lenderWithdrawEvent.parameters = new Array()

  lenderWithdrawEvent.parameters.push(
    new ethereum.EventParam("lender", ethereum.Value.fromAddress(lender))
  )
  lenderWithdrawEvent.parameters.push(
    new ethereum.EventParam("amount", ethereum.Value.fromUnsignedBigInt(amount))
  )

  return lenderWithdrawEvent
}

export function createOwnershipTransferredEvent(
  previousOwner: Address,
  newOwner: Address
): OwnershipTransferred {
  let ownershipTransferredEvent = changetype<OwnershipTransferred>(
    newMockEvent()
  )

  ownershipTransferredEvent.parameters = new Array()

  ownershipTransferredEvent.parameters.push(
    new ethereum.EventParam(
      "previousOwner",
      ethereum.Value.fromAddress(previousOwner)
    )
  )
  ownershipTransferredEvent.parameters.push(
    new ethereum.EventParam("newOwner", ethereum.Value.fromAddress(newOwner))
  )

  return ownershipTransferredEvent
}

export function createTransferEvent(
  from: Address,
  to: Address,
  value: BigInt
): Transfer {
  let transferEvent = changetype<Transfer>(newMockEvent())

  transferEvent.parameters = new Array()

  transferEvent.parameters.push(
    new ethereum.EventParam("from", ethereum.Value.fromAddress(from))
  )
  transferEvent.parameters.push(
    new ethereum.EventParam("to", ethereum.Value.fromAddress(to))
  )
  transferEvent.parameters.push(
    new ethereum.EventParam("value", ethereum.Value.fromUnsignedBigInt(value))
  )

  return transferEvent
}

export function createUserDefaultedEvent(
  userAddress: Address,
  defaultAmount: BigInt,
  privareData: Bytes
): UserDefaulted {
  let userDefaultedEvent = changetype<UserDefaulted>(newMockEvent())

  userDefaultedEvent.parameters = new Array()

  userDefaultedEvent.parameters.push(
    new ethereum.EventParam(
      "userAddress",
      ethereum.Value.fromAddress(userAddress)
    )
  )
  userDefaultedEvent.parameters.push(
    new ethereum.EventParam(
      "defaultAmount",
      ethereum.Value.fromUnsignedBigInt(defaultAmount)
    )
  )
  userDefaultedEvent.parameters.push(
    new ethereum.EventParam(
      "privareData",
      ethereum.Value.fromFixedBytes(privareData)
    )
  )

  return userDefaultedEvent
}

export function createUserRegistredEvent(
  userAddress: Address,
  privateData: Bytes,
  creditLimit: BigInt
): UserRegistred {
  let userRegistredEvent = changetype<UserRegistred>(newMockEvent())

  userRegistredEvent.parameters = new Array()

  userRegistredEvent.parameters.push(
    new ethereum.EventParam(
      "userAddress",
      ethereum.Value.fromAddress(userAddress)
    )
  )
  userRegistredEvent.parameters.push(
    new ethereum.EventParam(
      "privateData",
      ethereum.Value.fromFixedBytes(privateData)
    )
  )
  userRegistredEvent.parameters.push(
    new ethereum.EventParam(
      "creditLimit",
      ethereum.Value.fromUnsignedBigInt(creditLimit)
    )
  )

  return userRegistredEvent
}
