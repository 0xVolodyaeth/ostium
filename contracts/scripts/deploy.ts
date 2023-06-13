import { ethers } from "hardhat";

async function main() {
	const priceFeedAddress = "0x007A22900a3B98143368Bd5906f8E17e9867581b";

	const [deployer] = await ethers.getSigners();

	const ERC20Mock = await ethers.getContractFactory("ERC20Mock");
	const erc20mock = await ERC20Mock.deploy();
	await erc20mock.waitForDeployment();

	let tx = await erc20mock.mint(deployer.address, BigInt(1000) * BigInt(10) ** BigInt(18));
	await tx.wait();

	const Wager = await ethers.getContractFactory("Wager");
	const wager = await Wager.deploy(erc20mock.target, priceFeedAddress);
	await wager.waitForDeployment();

	tx = await erc20mock.approve(wager.target, BigInt(1000) * BigInt(10) ** BigInt(18));
	await tx.wait();

	console.log("ERC20Mock address:", erc20mock.target);
	console.log("Wager address:", wager.target);
}

main().catch((error) => {
	console.error(error);
	process.exitCode = 1;
});
