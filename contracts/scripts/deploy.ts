import { ethers } from "hardhat";

async function main() {
	const priceFeedAddress = "0x007A22900a3B98143368Bd5906f8E17e9867581b";

	const ERC20Mock = await ethers.getContractFactory("ERC20Mock");
	const erc20mock = await ERC20Mock.deploy();
	await erc20mock.waitForDeployment();

	const Wager = await ethers.getContractFactory("Wager");
	const wager = await Wager.deploy(priceFeedAddress, erc20mock.target);
	await wager.waitForDeployment();

	console.log("ERC20Mock address:", erc20mock.target);
	console.log("Wager address:", wager.target);
}

main().catch((error) => {
	console.error(error);
	process.exitCode = 1;
});
