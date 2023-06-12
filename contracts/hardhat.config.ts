import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@nomicfoundation/hardhat-verify";
import dotenv from "dotenv";

const cfg = dotenv.config().parsed!;
console.log(cfg.POLYGONSCAN_KEY)

const config: HardhatUserConfig = {
	solidity: "0.8.18",
	defaultNetwork: "mumbai",
	etherscan: {
		apiKey: cfg.POLYGONSCAN_KEY!
	},
	networks: {
		mumbai: {
			chainId: 80001,
			accounts: {
				mnemonic: cfg.MNEMONIC!,
				path: "m/44'/60'/0'/0",
				initialIndex: 0,
				count: 20
			},
			url: cfg.RPC_URL!,
		}
	}
};

export default config;
