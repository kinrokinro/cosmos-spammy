# Cure

Assaf Morami of Secret Network has provided me with an example transaction that adjusts the maximum block size by governance.  

I give him huge credit.  Both myself and Jehan Tremback were not aware this was possible. 


```json
{
    "title": "Reduce Maximum Block Size",
    "description": "This Proposal reduces the maximum block size pursuant to: https://github.com/cometbft/cometbft/security/advisories/GHSA-hq58-p9mv-338c",
    "changes": [
        {
            "subspace": "baseapp",
            "key": "BlockParams",
            "value": {
                "max_bytes": "1048576",
                "max_gas": "8000000"
            }
        }
    ],
    "deposit": "100000000uatom"
}
```