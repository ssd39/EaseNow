const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

module.exports = buildModule("EaseNowModule", (m) => {
  const easeNow = m.contract("EaseNow");
  return { easeNow };
});
