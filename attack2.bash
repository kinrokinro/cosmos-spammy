set -e

SEQUENCE=237014

for i in $( eval echo {"$SEQUENCE"..10000000} )
	 do gaiad tx ibc-transfer transfer transfer channel-58 cosmos140rptve4cr0mxgknzprl86868nfslydfyem3nq 1uatom --from test --keyring-backend test --home ~/.gaia-rs --memo $(openssl rand -hex 50000) --chain-id provider --gas 1109809 --yes --fees 2775uatom --sequence $i

	
done
