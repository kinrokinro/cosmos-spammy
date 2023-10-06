set -ue


loop () {

for i in $( eval echo {"$SEQUENCE"..10000} )
do

# Make a new transaction body with a random string
gaiad tx ibc-transfer transfer transfer channel-58 cosmos1fjzgfyt8way9sp7hktnv2jv73j697gvz3fyptm 1uatom  --keyring-backend test --home ~/.gaia-rs --memo $(openssl rand -hex 50000) --chain-id provider --yes   --packet-timeout-timestamp 0 --generate-only --fees 6969uatom --gas 1505075 --from cosmos1fjzgfyt8way9sp7hktnv2jv73j697gvz3fyptm &> bareibctx.json

# Step 1: Generate the random hex string and save it to a temporary file
openssl rand -hex 200000 > tmp.txt

# Step 2: Use jq with --arg to set the receiver field
jq --rawfile random_str tmp.txt '.body.messages[0].receiver = $random_str' bareibctx.json > autobanana.json

# Step 3: remove the temporary file afterwards
rm tmp.txt

# Step 4: Sign the transaction
gaiad tx sign autobanana.json --home /root/.gaia-rs --from cosmos18hmramafeyg3xu3j8m6s4w38sgt93r29v7c8d5 --yes --sequence 16081 --chain-id provider --keyring-backend test &> signedbanana.json

# Step 5: Broadcast the transaction
gaiad tx broadcast signedbanana.json --home ~/.gaia-rs

done

}

# get sequence number 
sequence () {

SEQUENCE=$(curl http://127.0.0.1:1317/cosmos/auth/v1beta1/accounts/cosmos18hmramafeyg3xu3j8m6s4w38sgt93r29v7c8d5 | jq --raw-output ' .account.sequence ')
echo "new sequence number is $SEQUENCE"
loop

}

sequence

