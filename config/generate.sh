#!/usr/bin/env bash

# Prior to generation, need to run:
export PATH=$GOPATH/src/github.com/hyperledger/fabric/build/bin:${PWD}/../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=${PWD}
CHANNEL_NAME=mychannel

# remove previous crypto material and config transactions
rm -fr config/*
rm -fr crypto-config/*

# The below assumes you have the relevant code available to generate the cryto-material
~/fabric-samples/bin/cryptogen generate --config=./crypto-config.yaml
~/fabric-samples/bin/configtxgen -profile OrgsOrdererGenesis -outputBlock genesis.block
~/fabric-samples/bin/configtxgen -profile OrgsChannel -outputCreateChannelTx mychannel.tx -channelID $CHANNEL_NAME
~/fabric-samples/bin/configtxgen -profile OrgsChannel -outputAnchorPeersUpdate Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP

# Rename the key files we use to be key.pem instead of a uuid
for KEY in $(find crypto-config -type f -name "*_sk"); do
    KEY_DIR=$(dirname ${KEY})
    mv ${KEY} ${KEY_DIR}/key.pem
done
