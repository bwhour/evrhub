pragma solidity ^0.5.0;


contract BridgeRegistry {
    address public evrnetBridge;
    address public bridgeBank;
    address public oracle;
    address public valset;

    event LogContractsRegistered(
        address _evrnetBridge,
        address _bridgeBank,
        address _oracle,
        address _valset
    );

    constructor(
        address _evrnetBridge,
        address _bridgeBank,
        address _oracle,
        address _valset
    ) public {
        evrnetBridge = _evrnetBridge;
        bridgeBank = _bridgeBank;
        oracle = _oracle;
        valset = _valset;

        emit LogContractsRegistered(evrnetBridge, bridgeBank, oracle, valset);
    }
}
