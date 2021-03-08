pragma solidity ^0.5.0;

import "../../../node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./BridgeToken.sol";


/**
 * @title EvrnetBank
 * @dev Manages the deployment and minting of ERC20 compatible BridgeTokens
 *      which represent assets based on the Evrnet blockchain.
 **/

contract EvrnetBank {
    using SafeMath for uint256;

    uint256 public bridgeTokenCount;
    mapping(string => address) controlledBridgeTokens;
    mapping(bytes32 => EvrnetDeposit) evrnetDeposits;

    struct EvrnetDeposit {
        address evrnetSender;
        address payable ethereumRecipient;
        address bridgeTokenAddress;
        uint256 amount;
        bool locked;
    }

    /*
     * @dev: Event declarations
     */
    event LogNewBridgeToken(address _token, string _symbol);

    event LogBridgeTokenMint(
        address _token,
        string _symbol,
        uint256 _amount,
        address _beneficiary
    );

    /*
     * @dev: Constructor, sets bridgeTokenCount
     */
    constructor() public {
        bridgeTokenCount = 0;
    }

    /*
     * @dev: Get a token symbol's corresponding bridge token address.
     *
     * @param _symbol: The token's symbol/denom without 'PEGGY' prefix.
     * @return: Address associated with the given symbol. Returns address(0) if none is found.
     */
    function getBridgeToken(string memory _symbol)
        public
        view
        returns (address)
    {
        return (controlledBridgeTokens[_symbol]);
    }

    /*
     * @dev: Creates a new EvrnetDeposit with a unique ID
     *
     * @param _evrnetSender: The sender's Evrnet address in bytes.
     * @param _ethereumRecipient: The intended recipient's Ethereum address.
     * @param _token: The currency type
     * @param _amount: The amount in the deposit.
     * @return: The newly created EvrnetDeposit's unique id.
     */
    function newEvrnetDeposit(
        address _evrnetSender,
        address payable _ethereumRecipient,
        address _token,
        uint256 _amount
    ) internal returns (bytes32) {
        bytes32 depositID = keccak256(
            abi.encodePacked(_evrnetSender, _ethereumRecipient, _token, _amount)
        );

        evrnetDeposits[depositID] = EvrnetDeposit(
            _evrnetSender,
            _ethereumRecipient,
            _token,
            _amount,
            true
        );

        return depositID;
    }

    /*
     * @dev: Deploys a new BridgeToken contract
     *
     * @param _symbol: The BridgeToken's symbol
     */
    function deployNewBridgeToken(string memory _symbol)
        internal
        returns (address)
    {
        bridgeTokenCount = bridgeTokenCount.add(1);

        // Deploy new bridge token contract
        BridgeToken newBridgeToken = (new BridgeToken)(_symbol);

        // Set address in tokens mapping
        address newBridgeTokenAddress = address(newBridgeToken);
        controlledBridgeTokens[_symbol] = newBridgeTokenAddress;

        emit LogNewBridgeToken(newBridgeTokenAddress, _symbol);
        return newBridgeTokenAddress;
    }

    /*
     * @dev: Mints new evrnet tokens
     *
     * @param _evrnetSender: The sender's Evrnet address in bytes.
     * @param _ethereumRecipient: The intended recipient's Ethereum address.
     * @param _evrnetTokenAddress: The currency type
     * @param _symbol: comsos token symbol
     * @param _amount: number of comsos tokens to be minted
     */
    function mintNewBridgeTokens(
        address _evrnetSender,
        address payable _intendedRecipient,
        address _bridgeTokenAddress,
        string memory _symbol,
        uint256 _amount
    ) internal {
        require(
            controlledBridgeTokens[_symbol] == _bridgeTokenAddress,
            "Token must be a controlled bridge token"
        );

        // Mint bridge tokens
        require(
            BridgeToken(_bridgeTokenAddress).mint(_intendedRecipient, _amount),
            "Attempted mint of bridge tokens failed"
        );

        newEvrnetDeposit(
            _evrnetSender,
            _intendedRecipient,
            _bridgeTokenAddress,
            _amount
        );

        emit LogBridgeTokenMint(
            _bridgeTokenAddress,
            _symbol,
            _amount,
            _intendedRecipient
        );
    }

    /*
     * @dev: Checks if an individual EvrnetDeposit exists.
     *
     * @param _id: The unique EvrnetDeposit's id.
     * @return: Boolean indicating if the EvrnetDeposit exists in memory.
     */
    function isLockedEvrnetDeposit(bytes32 _id) internal view returns (bool) {
        return (evrnetDeposits[_id].locked);
    }

    /*
     * @dev: Gets an item's information
     *
     * @param _Id: The item containing the desired information.
     * @return: Sender's address.
     * @return: Recipient's address in bytes.
     * @return: Token address.
     * @return: Amount of ethereum/erc20 in the item.
     * @return: Unique nonce of the item.
     */
    function getEvrnetDeposit(bytes32 _id)
        internal
        view
        returns (address, address payable, address, uint256)
    {
        EvrnetDeposit memory deposit = evrnetDeposits[_id];

        return (
            deposit.evrnetSender,
            deposit.ethereumRecipient,
            deposit.bridgeTokenAddress,
            deposit.amount
        );
    }
}
