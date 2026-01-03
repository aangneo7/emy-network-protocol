// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title EmyToken
 * @dev Core utility token for the Emy Network ecosystem.
 */
 
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract EmyToken is ERC20, Ownable {
    constructor() ERC20("Emy Network", "EMY") Ownable(msg.sender) {
        // Minting 1 Billion EMY tokens to the deployer
        _mint(msg.sender, 1000000000 * 10 ** decimals());
    }

    function burn(uint256 amount) public {
        _burn(msg.sender, amount);
    }
}

feat: implement basic ERC20 logic for EMY token
