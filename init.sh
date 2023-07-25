#!/bin/sh

#reset
rm -rf $HOME/.mintstationd

mintstationd config chain-id localnet-1
mintstationd config keyring-backend test
mintstationd config output json

#add keys
yes | mintstationd keys add admin
yes | mintstationd keys add validator
yes | mintstationd keys add cosmostation
ADMIN=$(mintstationd keys show admin -a)
COSMOSTATION=$(mintstationd keys show cosmostation -a)
VALIDATOR=$(mintstationd keys show validator -a)

#init
mintstationd init mintstation --chain-id localnet-1

#config cors
config="$HOME/.mintstationd/config/config.toml"
if [ "$(uname)" = "Linux" ]; then
  sed -i "s/cors_allowed_origins = \[\]/cors_allowed_origins = [\"*\"]/g" $config
else
  sed -i '' "s/cors_allowed_origins = \[\]/cors_allowed_origins = [\"*\"]/g" $config
fi

#generate genesis
mintstationd add-genesis-account $ADMIN 10000000000000000stake
mintstationd add-genesis-account $COSMOSTATION 10000000000000000stake
mintstationd add-genesis-account $VALIDATOR 10000000000000000stake
mintstationd gentx validator 10000000000stake --chain-id localnet-1 --keyring-backend test
mintstationd collect-gentxs
mintstationd validate-genesis

#start
mintstationd start
