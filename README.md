# Nebula AI mainet token Swap scanner
- Monitoring the Swap envent on NBAI Mainet and trigger the swap to BSC Chain
- Monitoring the Swap envent on BSC Chain and trigger the swap to NBAI Mainet
- Swap NBAI ERC20 Token from ETH Chain to Bsc Chain

## Mapping process
The user sends a certain number of nbai tokens to the transfer contract deployed on the nbai blockchain <tr>
Swan_scan scans the event in the background, and then uses the data in the event to call the mapping contract on the bsc blockchain

## Prerequisite
- Golang1.13 (minimum version)
- Mysql5.5

## Getting Started

*  Clone code to $GOPATH/src
```console
https://github.com/nebulaai/swap-scan.git
```

* Create a .env file in the project root directory
and enter the private key of your bsc blockchain wallet
Input value like this:  <br>
privateKeyForBscAdminWallet=<< your wallet private key in bsc blockchain >> <br>
privateKeyForNbaiAdminWallet=<< your wallet private key in nbai blockchain >>  
```console
cd $GOPATH/src/swan_scan
vi  .env
```
* Enter swan_scan/scan/config directory
change config.toml.example to config.toml 
input your database parameters and blockchain node address
```console
cd $GOPATH/src/swan_scan/scan/config
mv config.toml.example config.toml
```

* Enter payment-bridge directory and execute the make command
```console
cd $GOPATH/src/swan_scan/
GO111MODULE=on make
```

## Run the project   
Enter payment-bridge/scan/build/bin directory <br> 
execute the binary file of the payment-bridge project
```console
cd $GOPATH/src/swan_scan/scan/build/bin
chmod +x swan_scan
./payment_bridge
```


## Database table description

|table                     |description       |
|--------------------------|------------------|
|block_scan_record         |record the block number that has been scanned to the blockchain|
|event_nbai                |record eligible data on nbai            |
|event_bsc                 |record eligible data on bsc             |
|swap_coin_transaction     |transaction that excuted mapping contract   |



## Versioning and Releases

Feature branches and master are designated as **unstable** which are internal-only development builds.

Periodically a build will be designated as **stable** and will be assigned a version number by tagging the repository
using Semantic Versioning in the following format: `vMajor.Minor.Patch`.
