# Cure

Assaf Morami of Secret Network has provided me with an example transaction that adjusts the maximum block size by governance.  

I give him huge credit.  Both myself and Jehan Tremback were not aware this was possible. 


```bash
appd tx gov submit-proposal param-change ./path/to/cure.json --from key -y -b block
```



```json
{
    "title": "Reduce Maximum Block Size",
    "description": "This Proposal reduces the maximum block size pursuant to: https://github.com/cometbft/cometbft/security/advisories/GHSA-hq58-p9mv-338c",
    "changes": [
        {
            "subspace": "baseapp",
            "key": "BlockParams",
            "value": {
                "max_bytes": "1048576"
            }
        }
    ],
    "deposit": "100000000uatom"
}
```
