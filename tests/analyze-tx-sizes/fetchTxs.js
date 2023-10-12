const { Tendermint34Client } = require("@cosmjs/tendermint-rpc");
const { decodeTxRaw } = require("@cosmjs/proto-signing");
const fs = require("fs");
const path = require("path");

const TM_RPC = process.env.TM_RPC || "https://cosmos-rpc.rockrpc.net";
const START_HEIGHT = Number(process.env.START_HEIGHT) || 16674319;

let totalTxSize = 0;
let numTxs = 0;
let largestTxSize = 0;
let largestMsgSize = 0;
const msgSizeMap = {};
const largestTxMap = {};
const txDataTable = {};
const largestMsgs = {};
const largestTxs = [];

const updateAnalysisVariables = (txData, height) => {
  totalTxSize += txData.txSize;
  numTxs++;
  largestTxSize = Math.max(largestTxSize, txData.txSize);

  const composedMsgTypes = txData.msgTypesSizes.map(m => m.type).join(', ');

  largestTxs.push({ size: txData.txSize, height, composedMsgTypes });
  largestTxs.sort((a, b) => b.size - a.size);
  if (largestTxs.length > 50) largestTxs.pop();

  for (const { type, size } of txData.msgTypesSizes) {
    msgSizeMap[type] = (msgSizeMap[type] || 0) + size;
    largestMsgSize = Math.max(largestMsgSize, size);
    largestTxMap[type] = Math.max(largestTxMap[type] || 0, txData.txSize);

    largestMsgs[type] = largestMsgs[type] || [];
    largestMsgs[type].push({ size, height });
    largestMsgs[type].sort((a, b) => b.size - a.size);
    if (largestMsgs[type].length > 10) largestMsgs[type].pop();
  }
};

const saveAndExit = () => {
  const avgTxSize = totalTxSize / numTxs;
  fs.writeFileSync("txDataTable.json", JSON.stringify(txDataTable, null, 2));
  console.log("Transaction data table saved to txDataTable.json");
  fs.writeFileSync("txDataAnalysis.json", JSON.stringify({
    avgTxSize,
    largestTxSize,
    largestMsgSize,
    msgSizeMap,
    largestTxMap,
    largestMsgs,
    largestTxs
  }, null, 2));
  console.log("Transaction data analysis saved to txDataAnalysis.json");
  process.exit(0);
};

process.on("SIGINT", saveAndExit);

const fetchAndDecodeBlockTxs = async (client, height) => {
  console.log(`Fetching transactions for block ${height}...`);
  const block = await client.block(height);
  const txs = block.block.txs || [];

  const txDataArray = txs.map((tx) => {
    const decodedTx = decodeTxRaw(tx);
    const txSize = tx.length;
    const msgTypesSizes = decodedTx.body.messages.map((msg) => {
      const size = msg.value.length;
      const type = msg.typeUrl;
      return { type, size };
    });

    const txData = {
      height,
      txSize,
      msgTypesSizes
    };
    updateAnalysisVariables(txData, height);
    return txData;
  });

  return txDataArray;
};

const analyzeOnly = () => {
  if (!fs.existsSync("txDataTable.json")) {
    console.error("File not found: txDataTable.json");
    process.exit(1);
  }

  const loadedTxDataTable = JSON.parse(fs.readFileSync("txDataTable.json"));
  Object.assign(txDataTable, loadedTxDataTable);

  for (const height in loadedTxDataTable) {
    for (const txData of loadedTxDataTable[height]) {
      updateAnalysisVariables(txData, height);
    }
  }

  saveAndExit();
};

const main = async () => {
  if (process.argv.includes('--analyze-only')) {
    analyzeOnly();
    return;
  }

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
