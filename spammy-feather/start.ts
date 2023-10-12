import { MnemonicKey, LCDClient, MsgTransfer, Coin } from '@terra-money/feather.js';
import * as fs from 'fs';
import { exit } from 'process';

const init = async () => {
    // Read file once before the loop
    const fileContents = fs.readFileSync("./spam-data.txt").toString();

    // Configuration for the LCD client
    const lcdConfig = {
        'pisco-1': {
            lcd: 'http://88.99.146.56:1317',
            chainID: 'pisco-1',
            gasAdjustment: 1.75,
            gasPrices: { uluna: 0.015 },
            prefix: 'terra'
        }
    };

    // Initialize LCD Client
    const lcd = new LCDClient(lcdConfig);

    for (let walletIndex = 100; walletIndex < 200; walletIndex++) {
        newFunction(walletIndex);
    }

    async function newFunction(walletIndex: number) {
        // Initialize MnemonicKey and Wallet
        const mk = new MnemonicKey({
            mnemonic: "...",
            index: walletIndex
        });
        const wallet = lcd.wallet(mk);

        let accSeq = 0;
        let accNumber = 0;
        
        // Get current account sequence
        try {

            let authRes = (await lcd.auth.accountInfo(wallet.key.accAddress("terra")));

            accSeq = authRes.getSequenceNumber();
            accNumber = authRes.getAccountNumber();

            console.log('accSec: ' + accSeq);
            console.log('accNumber: ' + accNumber);
        }
        catch (e) {
            console.error(e);
        }

        while (true) {
            try {
                const msgExecute = new MsgTransfer(
                    "transfer",
                    "channel-351",
                    Coin.fromString("1uluna"),
                    wallet.key.accAddress("terra"),
                    "juno16p6lrlxf7f03c0ka8cv4sznr29rym27uhxw73t",
                    undefined,
                    (new Date().getTime() * 1000000 + 600000000).toString(),
                    fileContents
                );

                // Increment account sequence
                // Create and sign transaction
                const tx = await wallet.createAndSignTx({
                    msgs: [msgExecute],
                    chainID: "pisco-1",
                    sequence: accSeq,
                    accountNumber: accNumber,
                });
                accSeq = accSeq + 1;

                // Broadcast transaction
                const result = await lcd.tx.broadcast(tx, "pisco-1");
                console.log(`Transaction with accSeq ${accSeq} submitted on chain - Tx Hash: ${result.txhash}`);
                // console.log('broadcasted tx with sequence ' + accSeq);
            } catch (e) {
                console.log("An error occurred: ", e);
            }
        }
    }
};

try {
    init();
} catch (e) {
    console.log("Initialization failed due to error: ", e);
}
