// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import ".deps/npm/@openzeppelin/contracts/token/ERC721/ERC721.sol";
import ".deps/npm/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import ".deps/npm/@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";

contract Diploma is ERC721, ERC721URIStorage, ERC721Burnable {
    uint256 public tokenCounter = 0;
    address public admin;

    constructor() ERC721("Diploma", "DPL") {
        admin = msg.sender;
    }

    modifier onlyAdmin() {
        require(msg.sender == admin, "Only contract admin can mint diploma");
        _;
    }

    function safeMint(address to, string memory uri) public onlyAdmin {
        tokenCounter++;
        uint256 tokenId = tokenCounter;

        _safeMint(to, tokenId);
        _setTokenURI(tokenId, uri);

        // allow the contract admin to transfer all minted tokens
        if (super.isApprovedForAll(to, admin) == false) {
            super._setApprovalForAll(to, admin, true);
        }
    }

    // Override this function so that only contract owner (admin) can mint, transfer or burn diplomas
    function _update(address to, uint256 tokenId, address auth) internal override onlyAdmin returns (address) {
        return super._update(to, tokenId, auth);
    }

    // The following functions are overrides required by Solidity.

    function tokenURI(uint256 tokenId) public view override(ERC721, ERC721URIStorage) returns (string memory) {
        return super.tokenURI(tokenId);
    }

    function supportsInterface(bytes4 interfaceId) public view override(ERC721, ERC721URIStorage) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}