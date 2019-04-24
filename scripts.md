##Join channel
peer channel create -f configtx/channel.tx  -c mychannel -o orderer.dstar.com:7050
peer channel join -b mychannel.block -o orderer.dstar.com:7050
CORE_PEER_ADDRESS=peer1.org1.dstar.com:7051
peer channel join -b mychannel.block -o orderer.dstar.com:7050
CORE_PEER_ADDRESS=peer0.org1.dstar.com:7051
peer channel update -f configtx/Org1MSPanchors.tx -o orderer.dstar.com:7050 -c mychannel

##Install chaincode
peer chaincode install -o orderer.dstar.com:7050 -n gocc -l golang -v 1.0 -p github.com/ogpl_test/go/
CORE_PEER_ADDRESS=peer1.org1.dstar.com:7051
peer chaincode install -o orderer.dstar.com:7050 -n gocc -l golang -v 1.0 -p github.com/ogpl_test/go/
peer chaincode instantiate -o orderer.dstar.com:7050 -C mychannel -n gocc -l golang -v 1.0 -c '{"Args":[]}'

peer chaincode invoke -o orderer.dstar.com:7050 -C mychannel -n gocc -c '{"Args":["initLedger"]}'
peer chaincode invoke -o orderer.dstar.com:7050 -C mychannel -n gocc -c '{"Args":["queryAllDatasets"]}'

