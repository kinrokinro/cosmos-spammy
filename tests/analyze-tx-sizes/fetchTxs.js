const { Tendermint34Client } = require("@cosmjs/tendermint-rpc");
const { decodeTxRaw } = require("@cosmjs/proto-signing");
const fs = require("fs");

const TM_RPC = "https://cosmos-rpc.rockrpc.net";
const START_HEIGHT = 16674319;

let totalTxSize = 0;
let numTxs = 0;
let largestTxSize = 0;
let largestMsgSize = 0;
const msgSizeMap = {};
const largestTxMap = {};
const txDataTable = {};

const saveAndExit = () => {
  const avgTxSize = totalTxSize / numTxs;
  fs.writeFileSync("txDataTable.json", JSON.stringify(txDataTable, null, 2));
  fs.writeFileSync("txDataAnalysis.json", JSON.stringify({
    avgTxSize,
    largestTxSize,
    largestMsgSize,
    msgSizeMap,
    largestTxMap,
  }, null, 2));
  console.log("Transaction data saved to txDataTable.json");
  console.log("Transaction data analysis saved to txDataAnalysis.json");
  process.exit(0);
};

process.on("SIGINT", saveAndExit);

const fetchAndDecodeBlockTxs = async (client, height) => {
  console.log(`Fetching transactions for block ${height}...`);
  const block = await client.block(height);
  const txs = block.block.txs || [];

  const txData = txs.map((tx) => {
    const decodedTx = decodeTxRaw(tx);
    const txSize = tx.length;
    const msgTypesSizes = decodedTx.body.messages.map((msg) => {
      const size = msg.value.length;
      const type = msg.typeUrl;
      return { type, size };
    });

    totalTxSize += txSize;
    numTxs++;
    largestTxSize = Math.max(largestTxSize, txSize);

    for (const { type, size } of msgTypesSizes) {
      msgSizeMap[type] = (msgSizeMap[type] || 0) + size;
      largestMsgSize = Math.max(largestMsgSize, size);
      largestTxMap[type] = Math.max(largestTxMap[type] || 0, txSize);
    }

    return {
      height,
      txSize,
      msgTypesSizes,
    };
  });

  return txData;
};

const main = async () => {
  console.log("Connecting to Tendermint RPC...");
  const client = await Tendermint34Client.connect(TM_RPC);

  console.log("Fetching latest block height...");
  const latestBlock = await client.blockchain(0, 0);
  const latestBlockHeight = latestBlock.lastHeight;

  console.log("Starting to fetch and decode transactions...");
  for (let height = START_HEIGHT || 1; height <= latestBlockHeight; height++) {
    const txData = await fetchAndDecodeBlockTxs(client, height);
    txDataTable[height] = txData;
  }

  saveAndExit();
};

main().catch((err) => {
  console.error(`An error occurred: ${err}`);
});
