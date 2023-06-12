import {
	time,
	loadFixture,
} from "@nomicfoundation/hardhat-toolbox/network-helpers";
import { expect } from "chai";
import { ethers } from "hardhat";

import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";
import { ZeroAddress } from "ethers";
import { AggregatorV3Mock, Wager, ERC20Mock } from "../typechain-types";
import { deploy } from "./helpers.test";

describe("Wager", function () {
	const usdcBalance = BigInt(1000) * BigInt(10) ** BigInt(18);
	const betAmount = BigInt(500) * BigInt(10) ** BigInt(18);
	const expiration = BigInt(10);
	const wbtcPrice = BigInt(25000) * BigInt(10) ** BigInt(8);

	let usdc: ERC20Mock;
	let wager: Wager;
	let bob: SignerWithAddress;
	let alice: SignerWithAddress;
	let tom: SignerWithAddress;
	let oracle: AggregatorV3Mock;

	beforeEach(async () => {
		const protocol = await loadFixture(deploy);

		usdc = protocol.usdc;
		wager = protocol.wager;
		bob = protocol.bob;
		alice = protocol.alice;
		tom = protocol.tom;
		oracle = protocol.priceFeedMock;

		await usdc.connect(bob).mint(bob.address, usdcBalance);
		await usdc.connect(bob).approve(wager.target, usdcBalance);

		await usdc.connect(alice).mint(alice.address, usdcBalance);
		await usdc.connect(alice).approve(wager.target, usdcBalance);

		await oracle.connect(bob).prepareMock(wbtcPrice);
	});

	describe("Opening", function () {
		it("Should open a long bet", async function () {
			await expect(wager.connect(bob).openBet(betAmount, expiration, true)).to
				.emit(wager, "BetMade")
				.withArgs(bob.address, true, 0, betAmount, expiration, wbtcPrice);

			const latestBlockTimestamp = await time.latest();
			expect(await wager.bets(0)).to.deep.equal([
				bob.address,
				ZeroAddress,
				betAmount,
				expiration,
				latestBlockTimestamp,
				wbtcPrice,
				false,
			]);
		});

		it("Should open a short bet", async function () {
			await expect(wager.connect(bob).openBet(betAmount, expiration, false)).to
				.emit(wager, "BetMade")
				.withArgs(bob.address, false, 0, betAmount, expiration, wbtcPrice);

			const latestBlockTimestamp = await time.latest();
			expect(await wager.bets(0)).to.deep.equal([
				ZeroAddress,
				bob.address,
				betAmount,
				expiration,
				latestBlockTimestamp,
				wbtcPrice,
				false,
			]);
		});
	});

	describe("Canceling", function () {
		it("Should open a long bet and cancel it succesfully	", async function () {
			await wager.connect(bob).openBet(betAmount, expiration, true)

			await expect(wager.connect(bob).cancelBet(0)).to
				.emit(wager, "BetCanceled")
				.withArgs(bob.address, 0);

			expect(await usdc.balanceOf(bob.address)).to.equal(usdcBalance);
		});

		it("Should open a long bet and cancel it and get reverted because it's active", async function () {
			await wager.connect(bob).openBet(betAmount, expiration, true)
			await wager.connect(alice).joinBet(0);

			await expect(wager.connect(bob).cancelBet(0)).to.be
				.revertedWithCustomError(wager, "BetIsActive()");
		});

		it("Should open a long bet and cancel and be reveted because a caller is not a creator", async function () {
			await wager.connect(bob).openBet(betAmount, expiration, true)

			await expect(wager.connect(alice).cancelBet(0)).to.be
				.revertedWithCustomError(wager, "AddressIsNotBetCreator()");
		});
	});

	describe("Joining", function () {
		it("Should join a bet successfully", async function () {
			await wager.connect(bob).openBet(betAmount, expiration, true)
			const latestBlockTimestamp = await time.latest();

			await expect(wager.connect(alice).joinBet(0)).to
				.emit(wager, "JoinBet")
				.withArgs(alice.address, 0);

			expect(await usdc.balanceOf(alice.address)).to.equal(usdcBalance - betAmount);

			expect(await wager.bets(0)).to.deep.equal([
				bob.address,
				alice.address,
				betAmount,
				expiration,
				latestBlockTimestamp,
				wbtcPrice,
				true,
			]);
		});

		it("Should join a bet and be reverted because it is expired ", async function () {
			await wager.connect(bob).openBet(betAmount, expiration, true)
			await time.increase(expiration + BigInt(1));

			await expect(wager.connect(alice).joinBet(0)).to.be
				.revertedWithCustomError(wager, "BetIsExpiredAlready()");
		});

		it("Should join a bet and be reverted because it is taken already ", async function () {
			await wager.connect(bob).openBet(betAmount, expiration, true)
			await wager.connect(alice).joinBet(0)

			await expect(wager.connect(tom).joinBet(0)).to.be
				.revertedWithCustomError(wager, "BetAlreadyTaken()");
		});

		it("Should join a bet and be reverted because the bet doen't exist ", async function () {
			await expect(wager.connect(tom).joinBet(10)).to.be
				.revertedWithCustomError(wager, "BetDoesntExist()");
		});
	});

	describe("Withdrawing", function () {
		it("Should resolve a long bet successfully", async function () {
			await wager.connect(bob).openBet(betAmount, expiration, true);
			await wager.connect(alice).joinBet(0);

			await time.increase(expiration + BigInt(10));
			await oracle.prepareMock(wbtcPrice * BigInt(2));

			await expect(wager.connect(bob).resolveAndWithdraw(0)).to
				.emit(wager, "Withdrawn")
				.withArgs(bob.address, 0);

			expect(await usdc.balanceOf(bob.address)).to.equal(usdcBalance + betAmount);
		});

		it("Should resolve a short bet successfully", async function () {
			await wager.connect(bob).openBet(betAmount, expiration, false);
			await wager.connect(alice).joinBet(0);

			await time.increase(expiration + BigInt(10));
			await oracle.prepareMock(wbtcPrice * BigInt(2));

			await expect(wager.connect(alice).resolveAndWithdraw(0)).to
				.emit(wager, "Withdrawn")
				.withArgs(alice.address, 0);

			expect(await usdc.balanceOf(alice.address)).to.equal(usdcBalance + betAmount);
		});

		it("Should resolve a short bet and be reverted because address didn't win", async function () {
			await wager.connect(bob).openBet(betAmount, expiration, false);
			await wager.connect(alice).joinBet(0);

			await time.increase(expiration + BigInt(10));
			await oracle.prepareMock(wbtcPrice * BigInt(2));

			await expect(wager.connect(bob).resolveAndWithdraw(0)).to.be
				.revertedWithCustomError(wager, "AddressDidntWin()");
		});

		it("Should resolve and be reverted because the bet is not expired", async function () {
			await wager.connect(bob).openBet(betAmount, expiration, true);
			await wager.connect(alice).joinBet(0);

			await expect(wager.connect(bob).resolveAndWithdraw(0)).to.be
				.revertedWithCustomError(wager, "BetHasNotYetExpired()");
		});

		it("Should resolve and be reverted because the is not an addresses's bet", async function () {
			await wager.connect(bob).openBet(betAmount, expiration, true);
			await wager.connect(alice).joinBet(0);

			await time.increase(expiration + BigInt(10));

			await expect(wager.connect(tom).resolveAndWithdraw(0)).to.be
				.revertedWithCustomError(wager, "NotYourBet()");
		});

	});
});