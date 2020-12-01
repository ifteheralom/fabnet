#!/bin/bash

set -e

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1
starttime=$(date +%s)
LANGUAGE=${1:-"golang"}
CC_SRC_PATH=github.com/fabcar/go
if [ "$LANGUAGE" = "node" -o "$LANGUAGE" = "NODE" ]; then
	CC_SRC_PATH=/opt/gopath/src/github.com/fabcar/node
fi

# clean the keystore
rm -rf ./hfc-key-store

echo "*** launch network; create channel and join peers to channel"
cd ../1ord1peer
./start.sh
#
echo "*** Install Chaincode in peer0.org1.example.com"
#
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.org1.example.com:7051" cli peer chaincode install -n fabcar -v 1.0 -p "$CC_SRC_PATH" -l "$LANGUAGE"
sleep 10
#
echo "*** Instantiate Chaincode in peer0.org1.example.com or in ANY PEER ONLY ONCE"
#
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.org1.example.com:7051" cli peer chaincode instantiate -o orderer0.example.com:7050 -C mychannel -n fabcar -l "$LANGUAGE" -v 1.0 -c '{"Args":[""]}' -P "OR ('Org1MSP.member','Org2MSP.member')"
sleep 10
#
echo "*** Invoking chaincodes"
#
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp"  -e "CORE_PEER_ADDRESS=peer0.org1.example.com:7051" cli peer chaincode invoke -o orderer0.example.com:7050 -C mychannel -n fabcar -c '{"function":"initLedger","Args":[""]}'
#
export FABRIC_START_TIMEOUT=10
echo "*** wait for Hyperledger Fabric to finish"
echo ${FABRIC_START_TIMEOUT}
sleep ${FABRIC_START_TIMEOUT}

# clean the keystore
rm -rf ./hfc-key-store
node ../enrollAdmin.js
node ../registerUser.js
node ../dapp.js