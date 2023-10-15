#!/bin/bash

# Constants
APPNAME="celestia-appd"
CHAINID="mocha-4"
FEES=60000
UDENOM="utia"
GAS=600000
CHUNK_SIZE=100
CHUNK_COUNT=10
CHANNEL="channel-0" # This should be adjusted as per your channel setup
IBCTIMEOUTS="--packet-timeout-timestamp 0 --packet-timeout-height 0-100000"
MEMO_SIZE=50000 # This size will make the transaction size approximately 10KB

# Addresses
ADDRESS1="celestia1x7jn3tafxdhk844vgle5ga4plyqqxk39z4zsnk"
ADDRESS2="celestia1695pfdl4uxfy2yjr4kkrxvk4s4h964kn5hxn3k"

# Function to send a chunk of IBC transactions from sender to receiver
send_chunk_ibc() {
    local sender=$1
    local receiver=$2
    local chunk_num=$3
    local -n ref_sequence=$4  # Use nameref for updating sequence number

    for j in $(seq 1 $CHUNK_SIZE); do
        # Generate a random string for the --memo field to pad the transaction
        MEMO_DATA=$(openssl rand -hex $((MEMO_SIZE/2)))

        # Send the IBC transaction and capture the output and errors
        OUTPUT=$($APPNAME tx ibc-transfer transfer transfer $CHANNEL $receiver 1$UDENOM \
            --chain-id $CHAINID --yes \
            --fees $FEES$UDENOM --gas $GAS --sequence $ref_sequence \
            --from $sender \
            --memo $MEMO_DATA $IBCTIMEOUTS 2>&1)

        # Check if the output has the error regarding sequence mismatch
        if [[ $OUTPUT =~ "account sequence mismatch, expected ([0-9]+), got ([0-9]+)" ]]; then
            # Extract the expected sequence
            ref_sequence=${BASH_REMATCH[1]}

            # Resend the transaction with the correct sequence number
            $APPNAME tx ibc-transfer transfer transfer $receiver 1$UDENOM \
                --chain-id $CHAINID --yes \
                --fees $FEES$UDENOM --gas $GAS --sequence $ref_sequence \
                --from $sender \
                --memo $MEMO_DATA $IBCTIMEOUTS
        else
            # Log the output (whether it's an error or successful transaction)
            echo "$OUTPUT"
        fi

        # Log the progress
        echo "Chunk $chunk_num: Sent IBC transaction $j/$CHUNK_SIZE from $sender to $receiver with sequence $ref_sequence"

        # Increment the sequence number
        ref_sequence=$(($ref_sequence+1))
    done
}

# Main loop for sending transactions in alternating patterns
while true; do
    # Get the initial sequence numbers for both addresses
    SEQUENCE1=$(curl -s http://127.0.0.1:5003/cosmos/auth/v1beta1/accounts/$ADDRESS1 | jq --raw-output '.account.sequence')
    SEQUENCE2=$(curl -s http://127.0.0.1:5003/cosmos/auth/v1beta1/accounts/$ADDRESS2 | jq --raw-output '.account.sequence')

    # Send chunks from ADDRESS1 to ADDRESS2 and vice versa, alternating every CHUNK_COUNT chunks
    for i in $(seq 1 $CHUNK_COUNT); do
        send_chunk_ibc $ADDRESS1 $ADDRESS2 $i SEQUENCE1
        sleep 10  # Delay for 10 seconds
        send_chunk_ibc $ADDRESS2 $ADDRESS1 $i SEQUENCE2
        sleep 10  # Delay for 10 seconds
    done
done

