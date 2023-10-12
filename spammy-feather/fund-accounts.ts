import { MnemonicKey, LCDClient, MsgSend, Coins } from '@terra-money/feather.js';

const init = async () => {
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

    // Initialize MnemonicKey and Wallet
    const mk = new MnemonicKey({ mnemonic: "..." });
    const wallet = lcd.wallet(mk);

    for (let newWalletIndex = 104; newWalletIndex < 200; newWalletIndex++) {
        const newmkindex = new MnemonicKey({
            mnemonic: "...",
            index: newWalletIndex
        });
        const newWallet = lcd.wallet(newmkindex);
        const msgSend = new MsgSend(
            wallet.key.accAddress("terra"),
            newWallet.key.accAddress("terra"),
            Coins.fromString("10000000000uluna")
        );

        // Increment account sequence
        // Create and sign transaction
        const tx = await wallet.createAndSignTx({
            msgs: [msgSend],
            chainID: "pisco-1",
        });

        // Broadcast transaction
        const result = await lcd.tx.broadcast(tx, "pisco-1");
        console.log(`Submitted on chain:
         - To Wallet: ${newWallet.key.accAddress("terra")} 
         - With Index: ${newWalletIndex}
         - Tx Hash: ${result.txhash}`
        );
    }
};

try {
    init();
} catch (e) {
    console.log("Initialization failed due to error: ", e);
}
