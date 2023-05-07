#!/bin/sh

set -o errexit -o nounset

CHAINID=$1
GENACCT=$2

if [ -z "$1" ]; then
  echo "Need to input chain id..."
  exit 1
fi

if [ -z "$2" ]; then
  echo "Need to input genesis account address..."
  exit 1
fi

# Build genesis file incl account for passed address
coins="100000000000000000stake"
gstaked init --chain-id $CHAINID $CHAINID
gstaked keys add validator --keyring-backend="test"
gstaked add-genesis-account $(gstaked keys show validator -a --keyring-backend="test") $coins
gstaked add-genesis-account $GENACCT $coins
gstaked gentx validator 5000000000stake --keyring-backend="test" --chain-id $CHAINID
gstaked collect-gentxs

# Set proper defaults and change ports
sed -i 's#"tcp://127.0.0.1:26657"#"tcp://0.0.0.0:26657"#g' ~/.gstaked/config/config.toml
sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' ~/.gaia/config/config.toml
sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' ~/.gaia/config/config.toml
sed -i 's/index_all_keys = false/index_all_keys = true/g' ~/.gstaked/config/config.toml

# Start the gaia
gstaked start --pruning=nothing
