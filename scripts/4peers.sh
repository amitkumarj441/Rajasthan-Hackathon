#!/bin/bash
################################################################################################
#                                                                                              #
# Shell script to start up 4 validating peers using their .env files as environment variables. #
# Deploy humanity chaincode to the ledger, and write the hash received to a file.              #
#																							   #	
# Amit Kumar Jaiswal																			   #
# 20/03/2017                                                                                   #
################################################################################################

export GOROOT=/root/go
export GOPATH=/root/git
export HLDGPATH=$GOPATH/src/github.com/hyperledger/fabric/build/bin
export CCROOT=/root/humanity

RED='\033[0;31m'
CYAN='\033[0;36m'
NC='\033[0m'

#Clear things up:
read -p "${RED}This script will stop all running peers. Are you sure?(y/n)${NC}" -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]
then
	echo Clearing things up...
	echo Kill all previous processes
	killall node > /dev/null 2>&1
	killall peer > /dev/null 2>&1
	killall membersrvc > /dev/null 2>&1

	echo Delete all previous values in Ledger Database... 
	rm -rf /var/hyperledger/production
	
	echo Docker clean up:
	docker stop $(docker ps -a -q)      #stop all containers
	docker rm $(docker ps -a -q)		#remove all containers
	docker rmi $(docker images | grep dev-test) > /dev/null 2>&1    #remove all dev-test images 

    TIMESTAMP=$(date +%d-%m-%Y-%H%M%S)
    cd /logs
    mkdir $TIMESTAMP
    mv *.log $TIMESTAMP > /dev/null 2>&1
    if [ $? -eq 0 ]; then
        echo "${CYAN}Previous log files copied to $TIMESTAMP folder${NC}"
    else
      echo There were no previous log files.
    fi
fi
else 
	exit 1 

#Copy most recent chaincode to Fabric handler
#cd $CCROOT
#\cp -u humanity.go ~/git/src/github.com/hyperledger/fabric/examples/chaincode/go/humanity/humanity.go

cd $HLDGPATH
#Start Membership and Security Services
echo "Starting Membership and Security Server.."
./membersrvc > $CCROOT/logs/MemberSrvc.log 2>&1 &

#Start validating peers, each with its own .env file
echo "Starting validating peers:"
#VP0:
echo "Starting HyperLedger Fabric Validating Peer 1/4"
docker run --rm --env-file $CCROOT/env/vpo.env hyperledger/fabric-peer peer node start --logging-level=debug > $CCROOT/logs/vp0.log 2>$CCROOT/logs/vp0-err.log & 
echo "Waiting for initialization..."
sleep 20

#VP1:
echo "Starting HyperLedger Fabric Validating Peer 2/4"
docker run --rm --env-file $CCROOT/env/vp1.env hyperledger/fabric-peer peer node start --logging-level=debug > $CCROOT/logs/vp1.log 2> $CCROOT/logs/vp1-err.log & 
echo "Waiting for initialization..."
sleep 20

#VP2:
echo "Starting HyperLedger Fabric Validating Peer 3/4"
docker run --rm --env-file $CCROOT/env/vp2.env hyperledger/fabric-peer peer node start --logging-level=debug > $CCROOT/logs/vp2.log 2> $CCROOT/logs/vp2-err.log & 
echo "Waiting for initialization..."
sleep 20

#VP3:
echo "Starting HyperLedger Fabric Validating Peer 4/4"
docker run --rm --env-file $CCROOT/env/vp3.env hyperledger/fabric-peer peer node start --logging-level=debug > $CCROOT/logs/vp3.log 2> $CCROOT/logs/vp3-err.log & 
echo "Waiting for initialization..."
sleep 20

#TODO:change to manhood:
echo "deploying ManhoodCoins chaincode:"
CHAIN_NAME=`CORE_PEER_ADDRESS=0.0.0.0:30303 peer chaincode deploy -p github.com/hyperledger/fabric/examples/chaincode/go/humanity -c '{"Function":"Init", "Args": ["Laszlo","100","Zoe","40","Juci","100"]}'`

#write chaincode hash to file
echo $CHAIN >> $CCROOT/logs/chain_name.log

echo "${CYAN}ManhoodCoins chain:$CHAIN${NC}"

echo "Manhood chaincode ready." 
