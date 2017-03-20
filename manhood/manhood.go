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
     	   ThankType        string `json:"type"`      //number of points given 
     	   Message          string `json:"message"`   //the reason for giving thanks
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
return nil, err
}

fmt.Printf("Manhood points chaincode initialization ready.\n")
return nil, nil

//getRandomUser returns a random user from the ledger
func (t *ManhoodChaincode) getRandomUser(stub *shim.ChaincodeStub, function string,args []string) ([]byte, error) {
if len(args) != 0 {
	return nil, errors.New("Invalid number of arguments, expected 0")
}

//set the source for generating a random number
source := rand.NewSource(time.Now().UnixNano())
random := rand.New(source)

//get list of keys from ledger
keysJSON, err := stub.GetState("keys")
if keysJSON == nil || err != nil {
return nil, errors.New("Cannot get user list data from chain.")
}
	
//convert JSON to struct
keyListObj := keyList{}
err = json.Unmarshal(keysJSON, &keyListObj)
if err != nil
return nil, errors.New("Invalid user list data pulled from ledger")
}
	
//print and return an element in json form, from the slice containing a random name
randomUserObj := keyList{}
randomUserObj.Keys = append(randomUserObj.Keys, keyListObj.Keys[random.Intn(len(keyListObj.Keys))])
randomUserJson, err := json.Marshal(randomUserObj)
if err != nil || randomUserJson == nil {
return nil, errors.New("Converting struct to JSON failed")

fmt.Printf("Query Response:%s\n", randomUserJson)
return randomUserJson, nil
}

//addThanks function enables user to receive a "thank", adds points according to the thank level, and adds the "thank"
//to the person's thank list(name of "thanker", type and message).
func (t *ManhoodChaincode) addThanks(stub *shim.ChaincodeStub,function string, args []string) ([]byte, error) {
var userID string
var pointsToAdd int

//check arguments number and type
if len(args) != 2 {
return nil, errors.New("Incorrect number of arguments. Expecting 2")
}
userID = args[0]
//convert JSON to struct
thankJson := []byte(args[1])

var thankObj thank

err := json.Unmarshal(thankJson, &thankObj)
if err != nil {
return nil, errors.New("Invalid thank JSON")
}

//simple sanity check (message part can be ""):
if thankObj.ThankType != "ta" && thankObj.ThankType != "thankyou" && thankObj.ThankType != "bigthanks"{
return nil, errors.New("Invalid thank type! Valids are: ta(1), thankyou(5), bigthanks(10)")
}

if thankObj.Thanker =="" {
return nil, errors.New("No thanker name!")
}

//calculate how many points to add according to "thank level":
switch thankObj.ThankType {
	case "ta" : pointsToAdd = small
	case "thankyou" : pointsToAdd = medium
	case "bigthanks" : pointsToAdd = large
}

//get entity data from ledger:
entityJSON, err := stub.GetState(userID)
if entityJSON == nil {
return nil, errors.New("Error: No account exists for user.")
}

//convert JSON to struct
entityObj := entity{}
err = json.Unmarshal(entityJSON, &entityObj)
if err != nil {
return nil, errors.New("Invalid entity data pulled from ledger.")
}

//add points:
entityObj.Balance = entityObj.Balance + pointsToAdd


//add the thankObject to the thank array of the entityObject:
entityObj.AddThank(thankObj)
entityJSON = nil
entityJSON, err = json.Marshal(entityObj)

if err != nil || entityJSON == nil {
return nil, errors.New("Converting entity struct to JSON failed")
}

//write entity back to ledger
err = stub.PutState(userID, entityJSON)
if err != nil {
return nil, errors.New("Writing updated entity to ledger failed")
}
jsonResp := "{\"msg\": \"Thank added\"}"
fmt.Printf("Invoke Response:%s\n", jsonResp)
return []byte(jsonResp), nil

}

//Invoke function to invoke addThanks, and getRandomUser functions
func (t *ManhoodChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
if function == "addThanks" {
//Add points to a member
return t.addThanks(stub,function, args)
}
return nil, errors.New("Received unknown function invocation")
}

//getUser queries the ledger for a given ID and returns the whole JSON for the userID
func (t *ManhoodChaincode) getUser(stub *shim.ChaincodeStub,function string,args []string) ([]byte, error) {
if len(args) != 1 {
return nil, errors.New("Invalid number of arguments, expected 1")
}
userID := args[0]
//get user data from ledger
dataJson, err := stub.GetState(userID)
if dataJson == nil || err != nil {
return nil, errors.New("Cannot get user data from chain.")
}

fmt.Printf("Query Response: %s\n", dataJson)
return dataJson, nil
}

//getKeys queries the ledger for all user keys and returns it as a JSON
func (t *ManhoodChaincode) getKeys(stub *shim.ChaincodeStub,function string,args []string) ([]byte, error) {
if len(args) != 0 {
return nil, errors.New("Invalid number of arguments, expected 0")
}
//get list of keys from ledger
keysJSON, err := stub.GetState("keys")
if keysJSON == nil || err != nil {
return nil, errors.New("Cannot user list data from chain.")
}

fmt.Printf("Query Response: %s\n", keysJSON)
return keysJSON, nil
}

//query function to return a user, or a list of all user's keys
func (t *ManhoodChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
if function == "getUser" {
//Add points to a member
return t.getUser(stub, function, args)
} else if function == "getKeys" {
return t.getKeys(stub, function, args)
} else if function == "getRandomUser" {
return t.getRandomUser(stub,function, args)
}
return nil, errors.New("Received unknown function invocation")
}

//main function to start chaincode
func main() {
err := shim.Start(new(ManhoodChaincode))
if err != nil {
fmt.Printf("Error starting Manhood chaincode: %s", err)
}
}
