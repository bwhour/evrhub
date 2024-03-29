module.exports = async () => {
  /*******************************************
   *** Set up
   ******************************************/
  const Web3 = require("web3");
  const HDWalletProvider = require("@truffle/hdwallet-provider");

  // Contract abstraction
  const truffleContract = require("truffle-contract");
  const contract = truffleContract(
    require("../build/contracts/BridgeBank.json")
  );
  /*******************************************
   *** Constants
   ******************************************/
  // Config values
  let symbol;
  let NETWORK_ROPSTEN;
  let NETWORK_EVRYNET;
  try {
    NETWORK_ROPSTEN =
      process.argv[4] === "--network" && process.argv[5] === "ropsten";
    NETWORK_EVRYNET =
      process.argv[4] === "--network" && process.argv[5] === "evrynet";

    /*******************************************
     *** checkBridgeProphecy transaction parameters
    ******************************************/
    if (NETWORK_ROPSTEN || NETWORK_EVRYNET) {
      symbol = process.argv[6].toString();
    } else {
      symbol = process.argv[4].toString();
    }
  }catch (error) {
    console.log({error})
    return
  }
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

  console.log("Fetching BridgeBank contract...");
  contract.setProvider(web3.currentProvider);
  try {
    /*******************************************
     *** Contract interaction
    ******************************************/
    console.log("Attempting to send createNewBridgeToken() tx with symbol: '" + symbol + "'...");

    // Get the bridge token's address if it were to be created
    const bridgeTokenAddress = await contract.deployed().then(function(instance) {
      return instance.createNewBridgeToken.call(symbol, {
        from: operator,
        value: 0,
        gas: 3000000 // 300,000 Gwei
      });
    });
    console.log(`from ${operator}`)
    console.log('Should deploy to ' + bridgeTokenAddress)

    //  Create the bridge token
    await contract.deployed().then(function(instance) {
      return instance.createNewBridgeToken(symbol, {
        from: operator,
        value: 0,
        gas: 3000000 // 300,000 Gwei
      });
    });

    console.log("")
    // Check bridge token whitelist
    const isOnWhiteList = await contract.deployed().then(function(instance) {
      return instance.bridgeTokenWhitelist(bridgeTokenAddress, {
        from: operator,
        value: 0,
        gas: 3000000 // 300,000 Gwei
      });
    });

    if (isOnWhiteList) {
      console.log(
        'Bridge Token "' + symbol + '" created at address ' + bridgeTokenAddress + ' and added to whitelist'
      );
    } else {
      console.log(
        "Error: Bridge Token creation and whitelisting was not successful"
      );
    }
  } catch (error) {
    console.error({error})
  }

  return;
};
