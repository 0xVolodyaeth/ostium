import { ethers } from "hardhat";
import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";
import { AggregatorV3Mock, Wager, ERC20Mock } from "../typechain-types";


export interface Protocol {
	usdc: ERC20Mock;
	wager: Wager;
	bob: SignerWithAddress;
	alice: SignerWithAddress;
	tom: SignerWithAddress;
	priceFeedMock: AggregatorV3Mock;
}

export async function deploy(): Promise<Protocol> {
	const [bob, alice, tom] = await ethers.getSigners();

	const ERC20Mock = await ethers.getContractFactory("ERC20Mock");
	const usdc = await ERC20Mock.connect(bob).deploy();
	await usdc.waitForDeployment();

	const PriceFeedMock = await ethers.getContractFactory("AggregatorV3Mock");
	const priceFeedMock = await PriceFeedMock.connect(bob).deploy();
	await priceFeedMock.waitForDeployment();

	const Wager = await ethers.getContractFactory("Wager");
	const wager = await Wager.connect(bob).deploy(
		usdc.target,
		priceFeedMock.target,
	);

	await wager.waitForDeployment();

	return { usdc, wager, bob, alice, tom, priceFeedMock };
}