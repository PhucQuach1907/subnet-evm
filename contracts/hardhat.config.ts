import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@nomicfoundation/hardhat-ethers";
import "./tasks";
import * as dotenv from "dotenv";
dotenv.config();

const config: HardhatUserConfig = {
  solidity: {
    compilers: [
      {
        version: "0.8.24",
        settings: {
          evmVersion: "shanghai",
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
    ],
  },
  networks: {
    atichain: {
      url: process.env.RPC_URI,
      accounts: [process.env.PRIVATE_KEY],
      chainId: parseInt(process.env.ATI_CHAIN_ID || "1337"),
      gasPrice: "auto",
      gas: "auto",
    },
    cchain: {
      url: process.env.C_CHAIN_RPC_URI,
      accounts: [process.env.PRIVATE_KEY],
      chainId: 43113,
      gasPrice: "auto",
      gas: "auto",
    },
    hardhat: {
      chainId: 1337,
    },
  },
  mocha: {
    timeout: 30000,
  },
};

export default config;
