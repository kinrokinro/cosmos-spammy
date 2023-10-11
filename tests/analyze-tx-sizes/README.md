Nodejs script to output transaction and message size data to JSON files.

Install dependencies: `npm install`
Configure RPC and Start Block in the script:
```
const TM_RPC = "<YOUR-TM-RPC>";
const START_HEIGHT = <BLOCK-START-HEIGHT>;
```
Run script: `node fetchTxs.js`
Script runs until last block, or interrupt and save with `Crtl+C`
