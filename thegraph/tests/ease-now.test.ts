import {
  assert,
  describe,
  test,
  clearStore,
  beforeAll,
  afterAll
} from "matchstick-as/assembly/index"
import { Address, BigInt, Bytes } from "@graphprotocol/graph-ts"
import { AmountBorrowed } from "../generated/schema"
import { AmountBorrowed as AmountBorrowedEvent } from "../generated/EaseNow/EaseNow"
import { handleAmountBorrowed } from "../src/ease-now"
import { createAmountBorrowedEvent } from "./ease-now-utils"

// Tests structure (matchstick-as >=0.5.0)
// https://thegraph.com/docs/en/developer/matchstick/#tests-structure-0-5-0

describe("Describe entity assertions", () => {
  beforeAll(() => {
    let userAddress = Address.fromString(
      "0x0000000000000000000000000000000000000001"
    )
    let amount = BigInt.fromI32(234)
    let remainingLimit = BigInt.fromI32(234)
    let merchent = Address.fromString(
      "0x0000000000000000000000000000000000000001"
    )
    let isContract = "boolean Not implemented"
    let newAmountBorrowedEvent = createAmountBorrowedEvent(
      userAddress,
      amount,
      remainingLimit,
      merchent,
      isContract
    )
    handleAmountBorrowed(newAmountBorrowedEvent)
  })

  afterAll(() => {
    clearStore()
  })

  // For more test scenarios, see:
  // https://thegraph.com/docs/en/developer/matchstick/#write-a-unit-test

  test("AmountBorrowed created and stored", () => {
    assert.entityCount("AmountBorrowed", 1)

    // 0xa16081f360e3847006db660bae1c6d1b2e17ec2a is the default address used in newMockEvent() function
    assert.fieldEquals(
      "AmountBorrowed",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "userAddress",
      "0x0000000000000000000000000000000000000001"
    )
    assert.fieldEquals(
      "AmountBorrowed",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "amount",
      "234"
    )
    assert.fieldEquals(
      "AmountBorrowed",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "remainingLimit",
      "234"
    )
    assert.fieldEquals(
      "AmountBorrowed",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "merchent",
      "0x0000000000000000000000000000000000000001"
    )
    assert.fieldEquals(
      "AmountBorrowed",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "isContract",
      "boolean Not implemented"
    )

    // More assert options:
    // https://thegraph.com/docs/en/developer/matchstick/#asserts
  })
})
