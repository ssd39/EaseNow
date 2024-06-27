import {
  assert,
  describe,
  test,
  clearStore,
  beforeAll,
  afterAll
} from "matchstick-as/assembly/index"
import { Address, BigInt, Bytes } from "@graphprotocol/graph-ts"
import { AmountRepaid } from "../generated/schema"
import { AmountRepaid as AmountRepaidEvent } from "../generated/EaseNow/EaseNow"
import { handleAmountRepaid } from "../src/ease-now"
import { createAmountRepaidEvent } from "./ease-now-utils"

// Tests structure (matchstick-as >=0.5.0)
// https://thegraph.com/docs/en/developer/matchstick/#tests-structure-0-5-0

describe("Describe entity assertions", () => {
  beforeAll(() => {
    let userAddress = Address.fromString(
      "0x0000000000000000000000000000000000000001"
    )
    let amount = BigInt.fromI32(234)
    let remaingDebt = BigInt.fromI32(234)
    let newAmountRepaidEvent = createAmountRepaidEvent(
      userAddress,
      amount,
      remaingDebt
    )
    handleAmountRepaid(newAmountRepaidEvent)
  })

  afterAll(() => {
    clearStore()
  })

  // For more test scenarios, see:
  // https://thegraph.com/docs/en/developer/matchstick/#write-a-unit-test

  test("AmountRepaid created and stored", () => {
    assert.entityCount("AmountRepaid", 1)

    // 0xa16081f360e3847006db660bae1c6d1b2e17ec2a is the default address used in newMockEvent() function
    assert.fieldEquals(
      "AmountRepaid",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "userAddress",
      "0x0000000000000000000000000000000000000001"
    )
    assert.fieldEquals(
      "AmountRepaid",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "amount",
      "234"
    )
    assert.fieldEquals(
      "AmountRepaid",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "remaingDebt",
      "234"
    )

    // More assert options:
    // https://thegraph.com/docs/en/developer/matchstick/#asserts
  })
})
