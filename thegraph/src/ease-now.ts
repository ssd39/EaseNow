import {
  AmountBorrowed as AmountBorrowedEvent,
  AmountRepaid as AmountRepaidEvent,
  Approval as ApprovalEvent,
  LenderDeposit as LenderDepositEvent,
  LenderRewarded as LenderRewardedEvent,
  LenderWithdraw as LenderWithdrawEvent,
  OwnershipTransferred as OwnershipTransferredEvent,
  Transfer as TransferEvent,
  UserDefaulted as UserDefaultedEvent,
  UserRegistred as UserRegistredEvent
} from "../generated/EaseNow/EaseNow"
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
} from "../generated/schema"

export function handleAmountBorrowed(event: AmountBorrowedEvent): void {
  let entity = new AmountBorrowed(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.userAddress = event.params.userAddress
  entity.amount = event.params.amount
  entity.remainingLimit = event.params.remainingLimit
  entity.merchent = event.params.merchent
  entity.isContract = event.params.isContract

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleAmountRepaid(event: AmountRepaidEvent): void {
  let entity = new AmountRepaid(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.userAddress = event.params.userAddress
  entity.amount = event.params.amount
  entity.remaingDebt = event.params.remaingDebt

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleApproval(event: ApprovalEvent): void {
  let entity = new Approval(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.owner = event.params.owner
  entity.spender = event.params.spender
  entity.value = event.params.value

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleLenderDeposit(event: LenderDepositEvent): void {
  let entity = new LenderDeposit(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.lender = event.params.lender
  entity.amount = event.params.amount
  entity.locktime = event.params.locktime

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleLenderRewarded(event: LenderRewardedEvent): void {
  let entity = new LenderRewarded(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.lender = event.params.lender
  entity.rewardAmount = event.params.rewardAmount

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleLenderWithdraw(event: LenderWithdrawEvent): void {
  let entity = new LenderWithdraw(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.lender = event.params.lender
  entity.amount = event.params.amount

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleOwnershipTransferred(
  event: OwnershipTransferredEvent
): void {
  let entity = new OwnershipTransferred(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.previousOwner = event.params.previousOwner
  entity.newOwner = event.params.newOwner

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleTransfer(event: TransferEvent): void {
  let entity = new Transfer(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.from = event.params.from
  entity.to = event.params.to
  entity.value = event.params.value

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleUserDefaulted(event: UserDefaultedEvent): void {
  let entity = new UserDefaulted(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.userAddress = event.params.userAddress
  entity.defaultAmount = event.params.defaultAmount
  entity.privareData = event.params.privareData

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}

export function handleUserRegistred(event: UserRegistredEvent): void {
  let entity = new UserRegistred(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.userAddress = event.params.userAddress
  entity.privateData = event.params.privateData
  entity.creditLimit = event.params.creditLimit

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}
