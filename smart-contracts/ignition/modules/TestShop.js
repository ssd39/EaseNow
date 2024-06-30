const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

module.exports = buildModule("TestShopModule", (m) => {
  const testShop = m.contract("TestShop");
  return { testShop };
});
