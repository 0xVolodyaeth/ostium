// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {AggregatorV3Interface} from "../interfaces/AggregatorV3Interface.sol";

contract AggregatorV3Mock {
    int internal answer;

    constructor() {}

    function latestRoundData()
        external
        view
        returns (uint80, int256, uint256, uint256, uint80)
    {
        return (0, answer, 0, 0, 0);
    }

    function prepareMock(int256 _price) external {
        answer = _price;
    }
}
