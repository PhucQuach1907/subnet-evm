const hre = require("hardhat");
const { ethers } = require("ethers");

async function main() {
  const recipient = "0x39cB15C562a0490d517a0cF56d5d82b329aDe26B";
  const amount = 1000;

  const contractAddress = process.env.ATI_ERC20_TOKEN_ADDRESS;
  if (!contractAddress) {
    console.error("Please set ATI_ERC20_TOKEN_ADDRESS in your .env file");
    process.exit(1);
  }

  const token = await hre.ethers.getContractAt("ERC20Token", contractAddress);

  const [deployer] = await hre.ethers.getSigners();

  console.log("Minting tokens...");
  console.log("Contract address:", contractAddress);
  console.log("Recipient address:", recipient);
  console.log("Amount to mint:", amount);
  console.log("Using account:", deployer.address);

  const amountWithDecimals = ethers.parseUnits(amount.toString(), 18);
  const tx = await token.mint(recipient, amountWithDecimals);
  await tx.wait();

  console.log("Minting successful!");
  console.log("Transaction hash:", tx.hash);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
