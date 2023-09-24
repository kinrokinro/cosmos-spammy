# Spammy

## Tooling and documentation for the reproduction of:

* banana king
* client update spam
* banana client

In an effort to reporoduce a p2p storm on the cosmos hub testnet. 


## Target

We're looking to transmit a "banana client".  This transaction will 


## P2P storms have occured on the following networks:

* Sentinel
* Luna Classic
  * Led to economic harm 
* Stride


## Notes

Generate a random string (hex output)

```bash
openssl rand -hex 10000000
```

Sign a transcation


Broadcast a transaction (oddly cli won't broadcast, we'll need to chain it)
```bash
gaiad tx broadcast bananasigned1.json --home ~/.gaia-rs
```
