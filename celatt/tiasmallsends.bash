#!/bin/bash
set -ue

# Constants
APPNAME="celestia-appd"
CHAINID="mocha-4"
FEES=10000
UDENOM="utia"
GAS=100000
CHUNK_SIZE=100
CHUNK_COUNT=10

# Function to get the initial sequence number for an address
get_initial_sequence() {
    local address=$1
    echo $(curl -s http://127.0.0.1:5003/cosmos/auth/v1beta1/accounts/$address | jq --raw-output '.account.sequence')
}

# Send 1 utia from sender to receiver, ensuring order within the chunk
send_chunk() {
    local sender=$1
    local receiver=$2
    local chunk_num=$3
    local -n ref_sequence=$4  # Use nameref for updating sequence number

    for j in $(seq 1 $CHUNK_SIZE); do
        # Send the transaction and print its output
        $APPNAME tx bank send $sender $receiver 1$UDENOM \
            --chain-id $CHAINID --yes \
            --fees $FEES$UDENOM --gas $GAS --sequence $ref_sequence

        # Log the progress
        echo "Chunk $chunk_num: Sent transaction $j/$CHUNK_SIZE from $sender to $receiver with sequence $ref_sequence"

        # Increment the sequence number
        ref_sequence=$(($ref_sequence+1))
    done
}

while true 
do


# Main script execution
# Step 1: Get the initial sequence numbers
sequence_celestia1=$(get_initial_sequence "celestia1x7jn3tafxdhk844vgle5ga4plyqqxk39z4zsnk")
sequence_celestia1695=$(get_initial_sequence "celestia1695pfdl4uxfy2yjr4kkrxvk4s4h964kn5hxn3k")

# Step 2: Send from celestia1 to celestia1695
echo "Sending funds from celestia1x7jn3tafxdhk844vgle5ga4plyqqxk39z4zsnk to celestia1695pfdl4uxfy2yjr4kkrxvk4s4h964kn5hxn3k"
for i in $(seq 1 $CHUNK_COUNT); do
    send_chunk "celestia1x7jn3tafxdhk844vgle5ga4plyqqxk39z4zsnk" "celestia1695pfdl4uxfy2yjr4kkrxvk4s4h964kn5hxn3k" $i sequence_celestia1 &
done
wait

# Step 3: Wait for 10 seconds
echo "Waiting for 10 seconds..."
sleep 10

# Step 4: Send back from celestia1695 to celestia1
echo "Sending funds back from celestia1695pfdl4uxfy2yjr4kkrxvk4s4h964kn5hxn3k to celestia1x7jn3tafxdhk844vgle5ga4plyqqxk39z4zsnk"
for i in $(seq 1 $CHUNK_COUNT); do
    send_chunk "celestia1695pfdl4uxfy2yjr4kkrxvk4s4h964kn5hxn3k" "celestia1x7jn3tafxdhk844vgle5ga4plyqqxk39z4zsnk" $i sequence_celestia1695 &
done
wait

echo "Script execution completed successfully!"

done

