// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {AggregatorV3Interface} from "./interfaces/AggregatorV3Interface.sol";

contract Wager {
    struct Bet {
        // packing into two slots
        address long;
        address short;
        uint96 amount;
        uint96 expiration;
        // packing into one slot
        uint120 createdAt;
        uint128 openingPrice;
        bool isActive;
    }

    error TransferFailed();
    error SameAddressForLongAndShort();
    error BetDoesntExist();
    error BetAlreadyTaken();
    error BetHasNotYetExpired();
    error BetIsExpiredAlready();
    error BetIsActive();
    error NotYourBet();
    error AddressDidntWin();
    error AddressIsNotBetCreator();

    event BetMade(
        address initiator,
        bool long,
        uint256 indexed betId,
        uint96 amount,
        uint96 expiration,
        uint128 openingPrice
    );
    event JoinBet(address indexed joiner, uint256 indexed betId);
    event Withdrawn(address winner, uint256 indexed betId);
    event BetCanceled(address creator, uint256 indexed betId);

    // USDC returns true on transfer so we can omit SafeERC20 library
    uint256 public betId;
    IERC20 internal usdc;
    AggregatorV3Interface internal priceFeed;

    mapping(uint256 => Bet) public bets;

    constructor(address _usdc, address _priceFeed) {
        usdc = IERC20(_usdc);
        priceFeed = AggregatorV3Interface(_priceFeed);
    }

    /// @notice cancel a bet if it has not been taken by the other side
    /// @param _betId id of the bet to cancel
    function cancelBet(uint256 _betId) external {
        Bet storage bet = bets[_betId];

        if (bet.isActive) revert BetIsActive();

        address receiver = bet.long != address(0) ? bet.long : bet.short;
        if (receiver != msg.sender) revert AddressIsNotBetCreator();

        bool sent = usdc.transfer(msg.sender, bet.amount);
        if (!sent) revert TransferFailed();

        emit BetCanceled(msg.sender, _betId);

        delete bets[_betId];
    }

    /// @notice used to open a bet
    /// @param _amount amount of USDC which will be taken from openers address
    /// @param _expiration amount in seconds after which a bet will be closed
    /// @param _long if long is true, then msg.sender takes long, if false, then short
    function openBet(uint96 _amount, uint96 _expiration, bool _long) external {
        bool sent = usdc.transferFrom(msg.sender, address(this), _amount);
        if (!sent) revert TransferFailed();

        uint256 assetPrice = getLatestPrice();

        if (_long) {
            bets[betId] = Bet(
                msg.sender,
                address(0),
                _amount,
                _expiration,
                uint112(block.timestamp),
                uint128(assetPrice),
                false
            );
        } else {
            bets[betId] = Bet(
                address(0),
                msg.sender,
                _amount,
                _expiration,
                uint112(block.timestamp),
                uint128(assetPrice),
                false
            );
        }

        emit BetMade(
            msg.sender,
            _long,
            betId,
            _amount,
            _expiration,
            uint128(assetPrice)
        );

        unchecked {
            ++betId;
        }
    }

    /// @notice joins a bet which has not yet been joined
    /// @param _betId id of the bet to join
    function joinBet(uint256 _betId) external {
        Bet storage bet = bets[_betId];

        if (bet.long == msg.sender && bet.short == msg.sender)
            revert SameAddressForLongAndShort();

        if (bet.long == address(0) && bet.short == address(0))
            revert BetDoesntExist();

        if (bet.createdAt + bet.expiration <= block.timestamp)
            revert BetIsExpiredAlready();

        if (bet.isActive) revert BetAlreadyTaken();

        bool sent = usdc.transferFrom(msg.sender, address(this), bet.amount);
        if (!sent) revert TransferFailed();

        if (bet.short == address(0)) {
            bet.short = msg.sender;
        } else {
            bet.long = msg.sender;
        }

        bet.isActive = true;
        emit JoinBet(msg.sender, _betId);
    }

    /// @notice resolves a bet and withdraws reward from a bet which has been resolved
    /// @param _betId id of the bet to resolve
    function resolveAndWithdraw(uint256 _betId) external {
        Bet storage bet = bets[_betId];

        if (block.timestamp <= bet.expiration + bet.createdAt)
            revert BetHasNotYetExpired();

        if (!(bet.long == msg.sender || bet.short == msg.sender))
            revert NotYourBet();

        uint256 closingPrice = getLatestPrice();

        bool won = (bet.long == msg.sender &&
            closingPrice > bet.openingPrice) ||
            (bet.short == msg.sender && closingPrice < bet.openingPrice);

        if (!won) revert AddressDidntWin();

        bool sent = usdc.transfer(msg.sender, bet.amount * 2);
        if (!sent) revert TransferFailed();

        delete bets[_betId];
        emit Withdrawn(msg.sender, _betId);
    }

    function getLatestPrice() internal view returns (uint256) {
        (, int price, , , ) = priceFeed.latestRoundData();
        return uint256(price);
    }
}
