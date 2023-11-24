# Mintstation <img width="34" alt="logo" src="https://github.com/cosmostation/mintstation/assets/91711862/571eee8b-1d10-4893-8af8-5c1d647f06ec">

**Mintstation** is a blockchain built using Cosmos SDK and Tendermint.

Mintstation is a network designed for testing the Cosmos SDK. It provides access to the latest features of the Cosmos ecosystem, allowing users to leverage CosmWasm smart contracts.

By providing an environment that maintains the latest features of the Cosmos SDK and supports Cosmwasm smart contracts, Mintstation facilitates a dynamic and insightful experience for individuals and teams working within the Cosmos ecosystem.



## Get started

Build Requirements

```
Go 1.20.+
```

Clone source from repository:

```bash
git clone https://github.com/cosmostation/mintstation
cd mintstation
git checkout {CURRENT_VERSION}
```

Once you're on the correct tag, you can build:

```bash
# from mintstation dir
make install
```

To confirm that the installation has succeeded, you can run:

```bash
mintstationd version

#{CURRENT_VERSION}
```

To configure `chain-id` & `node`)

```bash
#Update config.toml
mintstationd config chain-id mintstation-1
mintstationd config node {NODE_RPC_ENDPOINT:PORT}
```

## Playground
Take a look at the code in Mintstation Playground to experience the easiest way to understand the Cosmos SDK and develop a DApp. This is the simplest way to grasp the Cosmos ecosystem.
![playground](https://github.com/cosmostation/mintstation/assets/91711862/c2534b26-acaf-49db-838d-ec2b1b8cf541)
- Github : https://github.com/comostation/mintstation-playground
- Demo : https://play-mintstation.web.app/

## Learn more

- [Mintstation docs](https://docs.cosmostation.io)
- [Mintstation playground](https://play-mintstation.web.app)
