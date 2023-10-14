#!/bin/bash


# Constants
APPNAME="celestia-appd"
CHANNEL="channel-0"
ADDRESS="celestia1x7jn3tafxdhk844vgle5ga4plyqqxk39z4zsnk"
IBCMEMO=50000
RECIEVEADDR=100000
GAS=3114054
CHAINID="mocha-4"
IBCTIMEOUTS="--packet-timeout-timestamp 0 --packet-timeout-height 0-100000"
FEES=391405
UDENOM="utia"
# Retrieve account details (Placeholder: Replace with actual curl command if needed)
ACCOUNT=72559


    # Retrieve sequence
    SEQUENCE=$(curl http://127.0.0.1:5003/cosmos/auth/v1beta1/accounts/$ADDRESS | jq --raw-output ' .account.sequence ')
    
    # Inner loop to handle payload delivery
    while true; do

        echo "Sequence number is $SEQUENCE"

        # Generate new transaction body with a random string
        $APPNAME tx ibc-transfer transfer transfer $CHANNEL $ADDRESS 1$UDENOM  \
            --keyring-backend test --memo $(cat /dev/urandom | tr -dc '[:alpha:]' | fold -w ${1:-$RECIEVEADDR} | head -n 1) --chain-id $CHAINID --yes $IBCTIMEOUTS \
            --generate-only --fees $FEES$UDENOM --gas $GAS --from $ADDRESS &> bareibctx.json
        echo "Transaction body generated with $((IBCMEMO*2)) byte IBC memo field"

        # Generate random hex string and set it to the receiver field
        cat /dev/urandom | tr -dc '[:alpha:]' | fold -w ${1:-$RECIEVEADDR} | head -n 1 > tmp.txt
        echo "Random string generated"
        jq --rawfile random_str tmp.txt '.body.messages[0].receiver = $random_str' bareibctx.json > autobanana.json
        echo "$RECIEVEADDR byte random string inserted"
        
        # Remove temporary file
        rm tmp.txt
        echo "Temporary file removed"

        # Sign the transaction
        $APPNAME tx sign autobanana.json --from $ADDRESS --yes --sequence $SEQUENCE \
            --chain-id $CHAINID --keyring-backend test --offline --account-number $ACCOUNT &> ban.json
        echo "Transaction signed"

        # Broadcast the transaction
	
        if $APPNAME tx broadcast ban.json > banana.log
	then
        cat banana.log
        echo "Transaction broadcasted"
	SEQUENCE=$(($SEQUENCE+1))

	else

echo "some kinda bug, sleeping for a minute"
sleep 60
	fi






    done


echo "If you're running it right, the script just restarted"
