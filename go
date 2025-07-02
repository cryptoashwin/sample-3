feat: deploy basic ERC20 token contract for testing
/contracts
  └── TestToken.sol
/scripts
  └── deploy.js
hardhat.config.js
package.json
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract TestToken is ERC20 {
    constructor(uint256 initialSupply) ERC20("Test Token", "TTK") {
        _mint(msg.sender, initialSupply);
    }
}
const hre = require("hardhat");

async function main() {
  const initialSupply = hre.ethers.utils.parseEther("1000000"); // 1 million tokens

  const TestToken = await hre.ethers.getContractFactory("TestToken");
  const token = await TestToken.deploy(initialSupply);

  await token.deployed();
  console.log("TestToken deployed to:", token.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
