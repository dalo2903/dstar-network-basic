# D-Star Hyperledger Fabric Network Prototype
1.  Installation
2. Project structure
3. Architecture
4. Running guide


## Prerequisite
1. Docker
	`sudo apt install docker.io`
2. Docker Compose
`sudo apt install docker-compose`
3. Hyperledger Fabric Binaries and Docker Images
`curl -sSL http://bit.ly/2ysbOFE | bash -s 1.4.0`
4. Put Hyperledger Fabric binaries into  {project root folder}/bin
## Project Structure
### chaincode
### config
### crypto-config

## Running guide
1. Generate certificates
run the command: `./generate.sh` to generate artifacts and cryptographic components for the project. Generated files are inside **config** and **crypto-config** folders.
>NOTE: Those are generated only **once** and are used across different PCs.
2. Init Docker Swarm Mode
`docker swarm init`
3. Join other PCs to the swarm as Manager
Join tokens can be get by typing the following command on the init PC:
`docker swarm join-token manager`
4. Create docker containers from `docker-compose.yml` file.
*PC1:
`docker-compose -f docker-compose.yml up -d ca.dstar.com orderer.dstar.com peer0.org1.dstar.com couchdb0`
*PC2:
`docker-compose -f docker-compose.yml up -d peer1.org1.dstar.com couchdb1`
*PC3: 
`docker-compose -f docker-compose.yml up -d peer2.org1.dstar.com couchdb2`