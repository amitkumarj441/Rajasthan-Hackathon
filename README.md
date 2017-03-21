# Rajasthan-Hackathon

Idea Description : 
Developing {Manhood/Humanity}Coins point system where users can honour each other with ManhoodCoins. Later ManhoodCoins can also be collected by wearable devices, or smart meters, rewarding people saving resources (electricity, water, gas) or living a healthy lifestyle reported by their wearable device. To avoid abuse of the system these points are audited by a tweeting auditor (using twitter). The ManhoodCoins points can be used by companies and government bodies to reward people doing good to their communities, health and to the environment.

# Application Background

## Technical Details

Types of thanks:
  - small = 1 {Manhood/Humanity}Coins
  - medium= 5 {Manhood/Humanity)Coins
  - large = 10{Manhood/Humanity}Coins
  
Attributes of a user:
  1. userID   (unique string, will be used as key)
  2. balance  (int, computed points from the type of thank)
  3. thanklist(string slice (array), array of the thanks received by the user)
  
Attributes of a thank:
  1. Thanker  (the name of the person giving the thank)
  2. ThankType(type of the thank small, medium, large)
  3. message  (a small message explaining the thank, can be empty)
  
## Setting up the development environment 

### Overview

In order to run this project on your system, the hyperledger fabric needs to be setup : [hyperledger/fabric](https://github.com/hyperledger/fabric)

### Prerequisites

  - [Git Client](https://git-scm.com/downloads)
  - [Go](https://golang.org/) - 1.7 or later
  - [Docker](https://www.docker.com/products/overview) - 1.12 or later
  - [Pip](https://pip.pypa.io/en/stable/installing/)
  

### `pip`, `behave` and `docker-compose`
     pip install --upgrade pip
     pip install behave nose docker-compose
     pip install -I flask==0.10.1 python-dateutil==2.2 pytz==2014.3 pyyaml==3.10 couchdb==1.0 flask-cors==2.0.1 requests==2.4.3 pyOpenSSL==16.2.0 pysha3==1.0b1 grpcio==1.0.4
     
     #PIP packages required for some behave tests
     pip install urllib3 ndg-httpsclient pyasn1 ecdsa python-slugify grpcio-tools jinja2 b3j0f.aop six
     
### Steps

### Set your GOPATH
Make sure you have properly setup your Host's [GOPATH environment variable](https://github.com/golang/go/wiki/GOPATH). This allows for both building within the Host and the VM.

### Cloning the Fabric project
Since the Fabric project is a Go project, you'll need to clone the Fabric repo to your $GOPATH/src directory. If your $GOPATH has multiple path components, then you will want to use the first one. There's a little bit of setup needed:
   - `cd $GOPATH/src`
   - `mkdir -p github.com/hyperledger`
   - `cd github.com/hyperledger`

  ### Setting your $GOPATH
  Define these environment variables in your ~/.bashrc
   -  `export GOROOT=/usr/local/go`
   -  `export PATH=${GOROOT}/bin:${PATH}`
   -  `export GOPATH=${HOME}/other_src/gopath`  # typical value change at will
   -  `export PATH=${GOPATH}/bin:${PATH}`
   
Recall that we are using `Gerrit` for source control, which has its own internal git repositories. Hence, we will need to clone from [Gerrit](https://github.com/hyperledger/fabric/blob/master/docs/source/Gerrit/gerrit.md#Working-with-a-local-clone-of-the-repository). For brevity, 
### `the command is as follows:`
      git clone ssh://LFID@gerrit.hyperledger.org:29418/fabric && scp -p -P 29418 LFID@gerrit.hyperledger.org:hooks/commit-msg fabric/.git/hooks/

#### Note:
Of course, you would want to replace LFID with your own [Linux Foundation ID](https://github.com/hyperledger/fabric/blob/master/docs/source/Gerrit/lf-account.md) .

### Building the Fabric
Once you have all the dependencies installed, and have cloned the repository, you can proceed to build and test the fabric.

### Application Setup

The codebase should be placed under /root directory. Once you set up the system, you can start the test script to verify the settings. `/tests/humanity_test.sh`

You should get the following output: TESTS PASSED: 4/4

Now you can start up your network of peers with [Manhood/Humanity}Coins chaincode to serve the backend of the application. `/scripts/4vpeers.sh`

The backend of this application is running GoLang code on the 4 peer blockchain network on the mainframe, similar to IBM's High Security Bussiness Network. The chaincode itself will create users at the init method call. The user names are stored in a separate part of the database too, in order to get a random name for simulating the selection for a reward.

## Use Cases 
  - The front end application (a mobile app) can connect to this network via Rest API calls.
  - Coffe bar rewards every 5th person coming in with a free drink with HumanityCoins greater than 1000 this week.
  - To honour people, researcher, companies and government bodies for doing good to their communities.
