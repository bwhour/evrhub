module.exports = async () => {
  /*******************************************
   *** Set up
   ******************************************/
  const Web3 = require("web3");
  const HDWalletProvider = require("@truffle/hdwallet-provider");

  // Contract abstraction
  const truffleContract = require("truffle-contract");
  const evrnetBridgeContract = truffleContract(
    require("../build/contracts/EvrnetBridge.json")
  );
  const oracleContract = truffleContract(
    require("../build/contracts/Oracle.json")
  );
  const bridgeBankContract = truffleContract(
    require("../build/contracts/BridgeBank.json")
  );

  /*******************************************
   *** Constants
   ******************************************/
  // Config values
  const NETWORK_ROPSTEN =
    process.argv[4] === "--network" && process.argv[5] === "ropsten";
  const NETWORK_EVRYNET =
      process.argv[4] === "--network" && process.argv[5] === "evrynet";

  /*******************************************
   *** Web3 provider
   *** Set contract provider based on --network flag
   ******************************************/
  let provider;
  let operator;
  if (NETWORK_ROPSTEN || NETWORK_EVRYNET) {
    provider = new HDWalletProvider(
      process.env.MNEMONIC,
      process.env.HDWALLET_PROVIDER
    );
    operator = process.env.OPERATOR;
  } else {
    provider = new Web3.providers.HttpProvider(process.env.LOCAL_PROVIDER);
    operator = process.env.LOCAL_OPERATOR;
  }

  const web3 = new Web3(provider);

  evrnetBridgeContract.setProvider(web3.currentProvider);
  oracleContract.setProvider(web3.currentProvider);
  bridgeBankContract.setProvider(web3.currentProvider);
  try {

  /*******************************************
   *** Contract interaction
   ******************************************/
  // Get deployed Oracle's address
  const oracleContractAddress = await oracleContract
    .deployed()
    .then(function(instance) {
      return instance.address;
    });

  // Set Oracle
  const { logs: setOracleLogs } = await evrnetBridgeContract
    .deployed()
    .then(function(instance) {
      return instance.setOracle(oracleContractAddress, {
        from: operator,
        value: 0,
        gas: 300000 // 300,000 Gwei
      });
    });
  // Get event logs
  const setOracleEvent = setOracleLogs.find(e => e.event === "LogOracleSet");
  console.log("EvrnetBridge's Oracle set:", setOracleEvent.args._oracle);

  // Get deployed BridgeBank's address
  const bridgeBankContractAddress = await bridgeBankContract
    .deployed()
    .then(function(instance) {
      return instance.address;
    });

  // Set BridgeBank
  const {
    logs: setBridgeBankLogs
  } = await evrnetBridgeContract.deployed().then(function(instance) {
    return instance.setBridgeBank(bridgeBankContractAddress, {
      from: operator,
      value: 0,
      gas: 300000 // 300,000 Gwei
    });
  });

  // Get event logs
  const setBridgeBankEvent = setBridgeBankLogs.find(
    e => e.event === "LogBridgeBankSet"
  );
  console.log(
    "EvrnetBridge's BridgeBank set:",
    setBridgeBankEvent.args._bridgeBank
  );

  return;
} catch (error) {
  console.error({error})
}
};
