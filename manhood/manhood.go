package main

import (
	      "math/rand"
	      "time"
	      "errors"
	      "fmt"
	      "encoding/json"
	      "strconv"
	      "github.com/hyperledger/fabric/core/chaincode/shim"
)

//predefined points for each level of thank:
const (
	     small int = 1 << iota	//1 points
	     medium = 5 * iota	//5 points
	     large = 5 * iota	//10 points
)

//Each user has an ID a balance and an array of thanks
//users can gain points by receiving thanks valued at different levels (small, medium, large)
type entity struct {
	   UserID    string `json:"name"`
	   Balance   int `json:"balance"`		//points received by thanks
	   ThankList []thank `json:"thank"`	//list of thanks received
}

//keylist to contain all keys for the users, so we could select one randomly
type keyList struct {
	   Keys []string `json:"keys"`
}

//Each thank contains:
//- the name of the "giver"
//- one of the three types of thanks: ta, thanks, bigthanks 
//- and a message stating their good deed
type thank struct{
	   Thanker	    string `json:"name"`      //person who gives the "thank"
     ThankType    string `json:"type"`      //number of points given 
  `  Message      string `json:"message"`   //the reason for giving thanks
}

//AddThank method adds a thank to the slice of thanks inside entity struct.
func (e *entity) AddThank(t thank) []thank {
        e.ThankList = append(e.ThankList, t)
        return e.ThankList
}

//AddKey method
func(kl *keyList) AddKey(k string) []string {
         kl.Keys = append(kl.Keys, k)
         return kl.Keys
}

//ManhoodChaincode is the receiver of chaincode functions
type ManhoodChaincode struct{
}

//Init function to initialize chaincode add entities and points to start with to the ledger.
func  (t *ManhoodChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
var userID string         //name of the user to be registered on the chain
var pointsToAdd int       //points to start with
var err error
keyListObj := keyList{}

//get attributes from args
if len(args) %2 != 0 {
return nil, errors.New("Incorrect number of args, Needs to be even: (ID, points)")
}

//fill the db from args
for index := 0; index < len(args); index += 2{
userID = args[index]
pointsToAdd, err = strconv.Atoi(args[index + 1])
if err != nil {
return nil, errors.New("Expecting integer value for intial points")
}
entityObj := entity{}

//add the usernames to the list of users
keyListObj.Keys = append(keyListObj.Keys, userID)
//keyListObj.AddKey(userID)

//fill entity struct
entityObj.UserID = userID
entityObj.ThankList = []thank{}
entityObj.Balance = pointsToAdd

//covert entity struct to entityJSON
entityJson, err := json.Marshal(entityObj)
if err != nil || entityJson == nil {
return nil, errors.New("Converting entity struct to entityJSON failed")
}

//write entity attributes into ledger
err = stub.PutState(userID, entityJson)
if err != nil {
fmt.Printf("Error: could not update ledger")
return nil, err
}

}
//convert keylist struct to keyListJSON
keyListJson, err := json.Marshal(keyListObj)
if err != nil || keyListJson == nil {
return nil, errors.New("Converting entity struct to keyListJSON failed")
}

//write keylist into ledger 
err = stub.PutState("keys", keyListJson)
if err != nil {
fmt.Printf("Error: could not update ledger")

  
  
