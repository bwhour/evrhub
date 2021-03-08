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
  const tokenContract = truffleContract(
      require("../build/contracts/BridgeToken.json")
  );

  const NULL_ADDRESS = "0x0000000000000000000000000000000000000000";

  /*******************************************
   *** Constants
   ******************************************/
      // Config values
  const NETWORK_ROPSTEN =
      process.argv[4] === "--network" && process.argv[5] === "ropsten";
  const NETWORK_DEVELOP =
      process.argv[4] === "--network" && process.argv[5] === "develop";
  const NUM_ARGS = process.argv.length - 4;

  /*******************************************
   *** Command line argument error checking
   ***
   *** truffle exec lacks support for dynamic command line arguments:
   *** https://github.com/trufflesuite/truffle/issues/889#issuecomment-522581580
   ******************************************/
  if (NETWORK_ROPSTEN || NETWORK_DEVELOP) {
    if (NUM_ARGS !== 6) {
      return console.error(
          "Error: invalid number of parameters, please try again."
      );
    }
  } else {
    if (NUM_ARGS !== 4) {
      return console.error(
          "Error: must specify recipient address, token address, and amount."
      );
    }
  }

  /*******************************************
   *** Lock transaction parameters
   ******************************************/
  let ethSender = "";
  let evrRecipient = "";
  let coinDenom = "";
  let amount = "";

  if (NETWORK_ROPSTEN || NETWORK_DEVELOP) {
    ethSender = process.argv[6];
    evrRecipient = process.argv[7];
    coinDenom = process.argv[8];
    amount = process.argv[9];
  } else {
    ethSender = process.argv[4];
    evrRecipient = process.argv[5];
    coinDenom = process.argv[6];
    amount = process.argv[7];
  }

  // Convert default 'eth' coin denom into null address
  if (coinDenom === "eth") {
    coinDenom = NULL_ADDRESS;
  }

  /*******************************************
   *** Web3 provider
   *** Set contract provider based on --network flag
   ******************************************/
  let provider;
  if (NETWORK_ROPSTEN) {
    provider = new HDWalletProvider(
        process.env.MNEMONIC,
        "https://ropsten.infura.io/v3/".concat(process.env.INFURA_PROJECT_ID)
    );
  } else {
    provider = new Web3.providers.HttpProvider(process.env.LOCAL_PROVIDER);
  }

  const web3 = new Web3(provider);
  contract.setProvider(web3.currentProvider);
  tokenContract.setProvider(web3.currentProvider);
  try {
    /*******************************************
     *** Contract interaction
     ******************************************/
    let _amount = web3.utils.toWei(amount)

    // Send approve transaction
    if(coinDenom !== NULL_ADDRESS) {
      const bridgeContractAddress = await contract
          .deployed()
          .then(function(instance) {
            return instance.address;
          });

      let instance = await tokenContract.at(coinDenom)
      const { logs } = await instance.approve(bridgeContractAddress, _amount, {
        from: ethSender,
        value: 0,
        gas: 300000 // 300,000 Gwei
      });

      console.log("Sent approval...");

      // Get event logs
      const eventA = logs.find(e => e.event === "Approval");

      // Parse event fields
      const approvalEvent = {
        owner: eventA.args.owner,
        spender: eventA.args.spender,
        value: Number(eventA.args.value)
      };

      console.log(approvalEvent);
    }

    // Send lock transaction
    console.log("Connecting to contract....");
    const { logs } = await contract.deployed().then(function (instance) {
      console.log("Connected to contract, sending lock...");
      return instance.lock(evrRecipient, coinDenom, _amount, {
        from: ethSender,
        value: coinDenom === NULL_ADDRESS ? _amount : 0,
        gas: 300000 // 300,000 Gwei
      });
    });

    console.log("Sent lock...");

    // Get event logs
    const event = logs.find(e => e.event === "LogLock");

    // Parse event fields
    const lockEvent = {
      to: event.args._to,
      from: event.args._from,
      symbol: event.args._symbol,
      token: event.args._token,
      value: Number(event.args._value),
      nonce: Number(event.args._nonce)
    };

    console.log(lockEvent);
  } catch (error) {
    console.error({ error });
  }
  return;
};
