#!/bin/bash

NODE_URL="http://127.0.0.1:26657"
BATCH_SIZE=100
ACCOUNT="72620"
CHANNEL="channel-2"
KEY_NAME="test"
UDENOM="utia"
APPNAME="celestia-appd"
CHAINID="mocha-4"
IBCMEMO=10000  # 50kb in bytes
IBCTIMEOUTS="--packet-timeout-timestamp 0 --packet-timeout-height 0-100000"
FEES="55000"
GAS="550000"
ADDRESS="celestia164cq648zdsehddhg6q8l5djva02fvu38ncxzk0"

# Calculate bytes for RECEIVEADDR
RECIEVEADDR=10000

# Get the current block number
current_block() {
  curl -s $NODE_URL/block | jq -r .result.block.header.height
}

# Get the size of the mempool
mempool_size() {
  curl -s $NODE_URL/unconfirmed_txs?limit=1 | jq -r .result.n_txs
}

# Get the size of the latest block
block_size() {
  curl -s $NODE_URL/block?height=$1 | jq -r .result.block.data.txs | jq length
}

start_block=$(current_block)
echo "Script starting at block height: $start_block"

# Query sequence number initially
SEQUENCE=$(curl http://127.0.0.1:1317/cosmos/auth/v1beta1/accounts/$ADDRESS | jq --raw-output ' .account.sequence ')



# Main loop
while true; do
  last_block=$(current_block)
  last_block_size=$(block_size $last_block)
  current_mempool_size=$(mempool_size)

  echo "Last block height: $last_block, size: $last_block_size transactions"
  echo "Current mempool size: $current_mempool_size transactions"

  # Send transactions in a batch
  for ((i=0; i<$BATCH_SIZE; i++)); do
    # Transaction body generation
    $APPNAME tx ibc-transfer transfer transfer $CHANNEL $ADDRESS 1$UDENOM --account-number $ACCOUNT --keyring-backend test --memo $(openssl rand -hex $IBCMEMO) --chain-id $CHAINID --yes $IBCTIMEOUTS --generate-only --fees $FEES$UDENOM --gas $GAS --from $ADDRESS &> bareibctx.json

    openssl rand -hex $RECIEVEADDR > tmp.txt
    jq --rawfile random_str tmp.txt '.body.messages[0].receiver = $random_str' bareibctx.json > autobanana.json
    rm tmp.txt

    # Sign the transaction
    $APPNAME tx sign autobanana.json --account-number $ACCOUNT --from $ADDRESS --yes --sequence $SEQUENCE --chain-id $CHAINID --keyring-backend test --offline &> ban.json

    # Broadcast the signed transaction
    $APPNAME tx broadcast ban.json --node $NODE_URL --chain-id $CHAINID &> broadcast.log
    cat broadcast.log
    echo "transaction broadcasted"

# Check broadcast.log for "code: 20" to indicate mempool fullness
if cat broadcast.log | grep -q "code: 20"; then
    echo -e "\e[31mMEMPOOL FULL!!!!!!!!!\e[0m"  # Print in red
    sleep 60  # Sleep for 60 seconds
    continue  # Skip the rest of the loop iteration
fi
    

    # If there's an account sequence mismatch, parse the expected value and use it
if cat broadcast.log | grep -q "account sequence mismatch"; then
  SEQUENCE=$(cat broadcast.log | grep -oP 'expected \K\d+')
  echo "we had an account sequence mismatch, adjusting to $SEQUENCE"
else
  ((SEQUENCE++))
fi

  done

  # Check for a new block before looping again
  while [[ $(current_block) -le $last_block ]]; do
    continue
  done
done

