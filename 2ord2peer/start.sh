#!/bin/bash
#
export MSYS_NO_PATHCONV=1
#
#docker stop $(docker ps -a -q)
#docker rm $(docker ps -a -q)
#
docker-compose -f docker-compose.yaml up -d
#
export FABRIC_START_TIMEOUT=10
echo "*** wait for Hyperledger Fabric to start"
echo ${FABRIC_START_TIMEOUT}
sleep ${FABRIC_START_TIMEOUT}
#
echo "*** Create the channel"
#
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel create -o orderer0.example.com:7050 -c mychannel -f /etc/hyperledger/configtx/mychannel.tx
#
echo "*** Join peer0.org1.example.com to the channel"
#
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel fetch 0 genesis.block --channelID mychannel --orderer orderer0.example.com:7050
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel join -b genesis.block
#
echo "*** Join peer1.org1.example.com to the channel"
#
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer1.org1.example.com peer channel fetch 0 genesis.block --channelID mychannel --orderer orderer0.example.com:7050
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer1.org1.example.com peer channel join -b genesis.block
#