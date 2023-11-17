APP_HOME="$HOME/.mintstation"
RPC="http://localhost:26657"
CHAIN_ID="localnet-1"
MONIKER="moniker"

rm -rf $APP_HOME

mintstationd init ${MONIKER} --chain-id ${CHAIN_ID} --home ${APP_HOME}

mintstationd config keyring-backend test

sed -i -r 's/minimum-gas-prices = "0stake"/minimum-gas-prices = "0.0001umint"/' ${APP_HOME}/config/app.toml
 
sed -i -e 's/\"stake\"/\"umint\"/g' ${APP_HOME}/config/genesis.json
 
mintstationd keys add validator
 
MY_VALIDATOR_ADDRESS=$(mintstationd keys show validator -a)
 
mintstationd add-genesis-account $MY_VALIDATOR_ADDRESS 10000000000umint
 
mintstationd gentx validator 10000000umint --chain-id localnet-1
 
mintstationd collect-gentxs
 
mintstationd start --home ${APP_HOME}
