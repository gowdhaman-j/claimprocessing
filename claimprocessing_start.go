/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
limitations under the License.
See the License for the specific language governing permissions and
*/

package main

import (
	"errors"
	"fmt"
    "strconv"
	"encoding/json"
	
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// ClaimProcessing implementation is a sample PoC code to demonstrate the BlockChain Capabilities
type ClaimProcessing struct {
}

var claimantIndexMap_Key = "claimantIndexMap_Key" 

// The below structure would help us to track the number of claims by a claimant
type ClaimantIndex struct{
    ClaimantId string  `json:"claimantid"`
    ClaimIds[] string  `json:"claimids"`
}


//Defining the required Structure for the PoC

type StateHistory struct{
	ClaimStatus string `json:"claimstatus"`
	ClaimStatusChanged string `json:"claimstatuschanged"`
	ActorName string `json:"actorname"`
	ActionDescription string `json:"actiondescription"`
}

type ClaimantDetailsType struct{
	ClaimantId string `json:"claimantid"`	
	ClaimantName string `json:"claimantname"`

}

type ActorType struct{
	ActorEmpId string `json:"actorempid"`				
	ActorName string `json:"actorname"`
    ActorRole string `json:"actorrole"`
	ActionDescription string `json:"actiondescription"`
}

type ClaimStateType struct{
    ClaimStatus string `json:"claimstatus"`
    ClaimStatusChanged string `json:"claimstatuschanged"`
}

type Claim struct{
    ClaimId string `json:"claimid"`
    ClaimDate string `json:"claimdate"`
    ClaimDescription string `json:"claimdescription"`
    ClaimantDetails ClaimantDetailsType `json:"claimantdetails"`
    ClaimedAmount string `json:"claimedamount"`
	ApprovedAmount string `json:"approvedamount"`
    ClaimState ClaimStateType `json:"claimstate"`
    ActorDetails ActorType `json:"actordetails"`
	StatesHistory []StateHistory `json:"stateshistory"`
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(ClaimProcessing))
	if err != nil {   
		fmt.Printf("Error starting ClaimProcessing chaincode: %s", err)

	}
}


// ============================================================================================================================
// Init ClaimProcessing - Initialize the ClaimProcessing
// ============================================================================================================================
func (t *ClaimProcessing) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var Aval int
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	// Initialize the chaincode
	Aval, err = strconv.Atoi(args[0])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}

	// Write the state to the ledger
	err = stub.PutState("abc", []byte(strconv.Itoa(Aval)))				//making a test var "abc", I find it handy to read/write to it right away to test the network
	if err != nil {
		return nil, err
	}

	claimantIndexOfTypeMap := make(map[string]ClaimantIndex)
	jsonAsBytes, _ := json.Marshal(claimantIndexOfTypeMap)
	err = stub.PutState(claimantIndexMap_Key, jsonAsBytes)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// ============================================================================================================================
// Run - Our entry point
// ============================================================================================================================
func (t *ClaimProcessing) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Entry to Invoke " + function)
	res := []byte("") 
	err := errors.New("")
	tempargs:= []string{args[0]}
	
	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "create_Claim" {		
		fmt.Println("******* Identied the function as create_claim *******")							//create a new claim
			// Let us check whether the claim exist before create
			fmt.Println("******* Before Creation, Check for duplication by calling the getClaim method *******")	
			res,err= t.Query(stub ,"getClaim",tempargs)
				if(err!=nil){// Expecting an error saying Claim does not exist
					  fmt.Println("******* Claim DOES NOT Exist... Hence Going for Create Claim ******* " ,err)
					// Let us create the claim
					return t.create_Claim(stub, args)
				}else if(res!=nil){//Claim exist, hence I'm updating
				fmt.Println("******* Claim EXIST ... Hence Going for update Claim *******")	
					return t.update_Claim(stub, args)
				}
	} else if function == "update_Claim" {										// update the claim 
				
		fmt.Println("******* Identied the function as update_claim *******")							//create a new claim
			// Let us check whether the claim exist before create
			fmt.Println("******* Before Update, Check for for existance of the claim *******")	
			res,err= t.Query(stub ,"getClaim",tempargs)
				if(err!=nil){// Expecting an error saying Claim does not exist
					  fmt.Println("******* Claim DOES NOT Exist... Hence Going for Create Claim ******* " ,err)
					// Let us create the claim
					return t.create_Claim(stub, args)
				}else if(res!=nil){//Claim exist, hence I'm updating
				fmt.Println("******* Claim EXIST ... Hence Going for update Claim *******")	
					return t.update_Claim(stub, args)
				}
	}


	fmt.Println("run did not find func: " + function)						//error

	return res, errors.New("Received unknown function invocation")
}

func (t *ClaimProcessing) create_Claim(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	inputValidationResp := []byte("")

	objClaimAsByes := []byte("")
	
	//   0		1		        2 	            3		     4		       5			6				7			      8				   9		10	      11			12                  13 (Status History) 14           15             16
	//ClaimId, ClaimDate, ClaimDescription, ClaimantId, ClaimantName, ClaimedAmount, ApprovedAmount ,ClaimStatus, ClaimStatusChanged, ActorEmpId, ActorName, ActorRole, ActionDescription, ClaimStatus,  ClaimStatusChanged, Actorname, ActionDescription

	fmt.Println("- start create_claim")

	//Validate the Input
	inputValidationResp,err= t.validteInputs(stub, args) 
	if(err!=nil || inputValidationResp==nil){
		return nil, err
	}

    claimId := args[0] 
	claimantId := args[3]

	// build the ClaimObj
	objClaimAsByes,err= t.buildClaimObj(stub, args) 
	if(err!=nil || objClaimAsByes==nil){
			return nil, errors.New("Error While building the ClaimObject")
	}
	
	//store claim with id as key
	err = stub.PutState(claimId, objClaimAsByes)									
	if (err != nil) {
		//Error while creating the claim
		fmt.Println("Error while creating the claim-->" ,err)
		return nil, err
	}
	
	//Get the claimantIndexMap_Key
	claimantAsBytes, err := stub.GetState(claimantIndexMap_Key)
	if err != nil {
		return nil, errors.New("Failed to get claimant index map")
	}
	claimantIndexOfTypeMap := make(map[string]ClaimantIndex)
	json.Unmarshal(claimantAsBytes, &claimantIndexOfTypeMap)
	
	// Get a value for a key with `name[key]`.
	v1,boolvar := claimantIndexOfTypeMap[claimantId]	
	if(boolvar != true){
		//Not able to find claimantId, hence create a new key-value pair
		// Build the ClaimantIndex structure
		
		arrayofClaimIds := []string{claimId}
		objClaimantIndex := ClaimantIndex{}
		objClaimantIndex.ClaimantId=claimantId
		objClaimantIndex.ClaimIds=arrayofClaimIds
		
		claimantIndexOfTypeMap[claimantId] = objClaimantIndex
	}else{
		//able to find the key with claimantId
		arrayofClaimIds := v1.ClaimIds
		arrayofClaimIds = append(arrayofClaimIds, claimId)
		
		
		objClaimantIndex := ClaimantIndex{}
		objClaimantIndex.ClaimantId=claimantId
		objClaimantIndex.ClaimIds=arrayofClaimIds
		
		delete(claimantIndexOfTypeMap, claimantId)
		
		claimantIndexOfTypeMap[claimantId] = objClaimantIndex
		
	}
	
	jsonAsBytes, _ := json.Marshal(claimantIndexOfTypeMap)
	err = stub.PutState(claimantIndexMap_Key, jsonAsBytes)	

	fmt.Println("- end create_claim")
		
	
	return nil, nil
}

func (t *ClaimProcessing) update_Claim(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	objClaimAsByes := []byte("")
	inputValidationResp := []byte("")
	

	//   0		1		        2 	            3		     4		       5			6				7			      8				   9		10	      11			12                  13 (Status History) 14           15             16
	//ClaimId, ClaimDate, ClaimDescription, ClaimantId, ClaimantName, ClaimedAmount, ApprovedAmount ,ClaimStatus, ClaimStatusChanged, ActorEmpId, ActorName, ActorRole, ActionDescription, ClaimStatus,  ClaimStatusChanged, Actorname, ActionDescription

	fmt.Println("- start update_Claim")

	//Validate the Input
	inputValidationResp,err= t.validteInputs(stub, args) 
	if(err!=nil || inputValidationResp==nil){
		return nil, err
	}

    claimId := args[0] 
	
	// build the ClaimObj
	objClaimAsByes,err= t.buildClaimObj(stub, args) 
	if(err!=nil || objClaimAsByes==nil){
			return nil, errors.New("Error While building the ClaimObject")
	}

	//store claim with id as key
	err = stub.PutState(claimId, objClaimAsByes)									
	if (err != nil) {
		//Error while creating the claim
		fmt.Println("Error while updating the claim-->" ,err)
		return nil, err
	}
	
	fmt.Println("- end update_claimStatus")
	return nil, nil
}

// ============================================================================================================================
// Query - Our entry point for Queries
// ============================================================================================================================
func (t *ClaimProcessing) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Entering Query method " + function)
	res := []byte("")
	err := errors.New("")
	// Handle different functions
	if function == "getClaim" {													
		res,err= t.getClaim(stub, args)
		if(res!=nil){
			//Claim Exist
			return res,nil
		}else if(err!=nil){
			return nil,err
		}
	}else if function == "getClaimByClaimant" {													
		res,err= t.getClaimByClaimant(stub, args)
		if(res!=nil){
			//Claim does not Exist
			return res,nil
		}else if(err!=nil){
			return nil,err
		}
	}
	fmt.Println("query did not find func: " + function)						

	return nil, errors.New("Received unknown function query")
}  

func (t *ClaimProcessing) getClaim(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {  
	var claimId, jsonResp string
	var err error
 fmt.Println("Entering GetClaim method")	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting Claimid in the argument to query")
	}

	claimId = args[0]
	valAsbytes, err := stub.GetState(claimId)									//get the var from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + claimId + "\"}"
		return nil, errors.New(jsonResp)
	}
	fmt.Println("Exitng GetClaim method")	
	return valAsbytes, nil
	

}

func (t *ClaimProcessing) validteInputs(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//Do Input Sanitation
	if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return nil, errors.New("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return nil, errors.New("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return nil, errors.New("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return nil, errors.New("6th argument must be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return nil, errors.New("7th argument must be a non-empty string")
	}
	if len(args[7]) <= 0 {
		return nil, errors.New("8th argument must be a non-empty string")
	}
	if len(args[8]) <= 0 {
		return nil, errors.New("9th argument must be a non-empty string")
	}
	if len(args[9]) <= 0 {
		return nil, errors.New("10th argument must be a non-empty string")
	}
	if len(args[10]) <= 0 {
		return nil, errors.New("11th argument must be a non-empty string")
	}
	if len(args[11]) <= 0 {
		return nil, errors.New("12th argument must be a non-empty string")
	}
	if len(args[12]) <= 0 {
		return nil, errors.New("13th argument must be a non-empty string")
	}
	return []byte("Valid Input"),nil
}

func (t *ClaimProcessing) buildClaimObj(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	var err error
	var objClaim Claim
	var objStateHistory StateHistory
	var arrayOfStatesHistory []StateHistory
	


	claimId := args[0]

	fmt.Println("2222222222222222-->")
	
	claimDate := args[1] 
	claimDesc := args[2]
	claimantId := args[3]
	claimantName := args[4]
	claimedAmount := args[5]
	approvedAmount := args [6]
	claimStatus := args[7]
	claimStatusChanged := args[8]
	actorEmpId := args[9]
	actorName := args[10]
	actorRole := args[11]
	actionDesc := args[12]
	
	
	// Build 3 sub-structures inside Claim structure
	strClaimantDetailsType := `{"claimantid": "` + claimantId + `", "claimantname": "` + claimantName + `"}`
	

	
	strClaimStateType := `{"claimstatus": "` + claimStatus + `", "claimstatuschanged": "` + claimStatusChanged + `"}`
	
	strActorType :=  `{"actorempid": "` + actorEmpId + `", "actorname": "` + actorName + `", "actorrole": "` + actorRole + `", "actiondescription": "` + actionDesc + `"}`
	
	strStateHistoryType := `{"claimstatus": "` + claimStatus + `", "claimstatuschanged": "` + claimStatusChanged + `", "actorname": "` + actorName + `", "actiondescription": "` + actionDesc + `"}`
	err = json.Unmarshal([]byte(strStateHistoryType), &objStateHistory) // Convert the value to be appeneded to JSON
	//Build Claim structure
	fmt.Println("Current varStateHistory.ClaimStatus----->" + objStateHistory.ClaimStatus)
	
	
	valAsbytes, err := stub.GetState(claimId)		
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + claimId + "\"}"
		return nil, errors.New(jsonResp)
	}
	
	
	json.Unmarshal(valAsbytes, &objClaim)
	arrayOfStatesHistory = objClaim.StatesHistory	//get the current value of StatesHistory array and then append
	arrayOfStatesHistory = append(arrayOfStatesHistory, objStateHistory) // append the current state to array
	
	
	strStatesHistory, _ := json.Marshal(arrayOfStatesHistory);
	
	
	strClaim := `{"claimid": "` + claimId + `", "claimdate": "` + claimDate + `", "claimdescription": "` + claimDesc + `", "claimantdetails": ` + strClaimantDetailsType + `, "claimedamount": "` + claimedAmount + `", "approvedamount": "` + approvedAmount + `", "claimstate": ` + strClaimStateType + `, "actordetails": ` + strActorType + `, "stateshistory": `+ string(strStatesHistory) +`}`

	

	return []byte(strClaim),nil
}

func (t *ClaimProcessing) getClaimByClaimant(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonResp string
	var err error
	var jsonAsBytes []byte
	var objClaim Claim
	var arrayOfClaims []Claim
	var arrayofClaimIds []string
	
	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
	}
	
	claimantId := args[0]
	
	claimantAsBytes, err := stub.GetState(claimantIndexMap_Key)									//get the var from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for claimantIndexMap_Key\"}"
		return nil, errors.New(jsonResp)
	}
	
	
	claimantIndexOfTypeMap := make(map[string]ClaimantIndex)
	json.Unmarshal(claimantAsBytes, &claimantIndexOfTypeMap)
	
	// Get a value for a key with `name[key]`.
	objClaimantIndex,boolvar := claimantIndexOfTypeMap[claimantId]
	if(boolvar != true){
		return nil, errors.New("This claimantId does not exist")
	}else{
		arrayofClaimIds = objClaimantIndex.ClaimIds
		for i := range arrayofClaimIds{
			strClaimId := arrayofClaimIds[i]
			claimAsBytes, err := stub.GetState(strClaimId)
			if(err==nil){
				json.Unmarshal(claimAsBytes, &objClaim)
				arrayOfClaims = append(arrayOfClaims, objClaim)
			}
			
		}	
		jsonAsBytes, _ = json.Marshal(arrayOfClaims)
		
	}	
	return jsonAsBytes, nil 
	
}