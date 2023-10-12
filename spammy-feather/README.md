# R&D Spammy Feather

Resource to spam CosmosSDK with transactions. This solution will iterate a wallet based on your mnemonic 100 indexes, if you need to fund all these accounts just send funds to the index 0 and run `ts-node fund-accounts.ts` if you want to run the spammy feather run `ts-node start.ts`, remember you need to have funds on all the accounts from the mnemonic starting from index 100 to 199.

This code was completed by @clemensgg and @emidev98 during Cosmoverse 2023 and was used to determine if Luna 2was vulnerable to p2p storms.  Luna 2 is vulnerable to p2p storms despite an 800kb block size.
