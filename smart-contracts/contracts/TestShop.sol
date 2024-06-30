// contracts/GameItem.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ERC721URIStorage, ERC721} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract TestShop is ERC721URIStorage {
    uint256 private _nextTokenId;

    constructor() ERC721("HadShaver", "HS") {}

    function buy(address buyer) public payable returns (uint256) {
        require(msg.value == 0.001 ether, "In sufficient amount to buy an item!");
        uint256 tokenId = _nextTokenId++;
        _mint(buyer, tokenId);
        _setTokenURI(tokenId, "https://head.shaver/");
        return tokenId;
    }
}
