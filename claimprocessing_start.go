/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
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
	fmt.Printf("Entry to Invoke " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "create_claim" {									//create a new claim
		return t.create_claim(stub, args)
	} else if function == "update_claimStatus" {										// update the claim status
		return t.update_claimStatus(stub, args)
	}


	fmt.Printf("run did not find func: " + function)						//error

	return nil, errors.New("Received unknown function invocation")
}

func (t *ClaimProcessing) create_claim(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//var err error
	//args
	//0			1			2 					3			4			5					6				7				8			9			10			11			12
	//ClaimId, ClaimDate, ClaimDescription, ClaimantId, ClaimantName, ClaimedAmount, ApprovedAmount ClaimStatus, ClaimStatusChanged, ActorEmpId, ActorName, ActorRole, ActionDescription

	fmt.Printf("- start crete_claim")

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
	//Input sanitation ends

	//check if the claim id already exists..
	claimId := args[0]
	var err error
	/*
	claimAsBytes, err := stub.GetState(claimId)
	if err != nil {
		//return nil, errors.New("Failed to get claim id")
	}else{
		objClaim := Claim{}
		json.Unmarshal(claimAsBytes, &objClaim)
		//if objClaim.ClaimId == claimId{
		//	fmt.Printf("This claim arleady exists: " + claimId)
		//	fmt.Printf(objClaim);
			return nil, errors.New("This claim arleady exists")				//all stop a claim by this id exists
		//}
	}
	
	*/
	
	//Now we have to build the Claim structure
	//before that lets build other sub-structures required for Claim struct
	
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
	strClaimantDetailsType := `{"claimantid": "}` + claimantId + `", "claimantname": ",` + claimantName + `"}`
	
	//objClaimantDetailsType := ClaimantDetailsType{}
	//objClaimantDetailsType.ClaimantId = claimantId
	//objClaimantDetailsType.ClaimantName = claimantname
	
	strClaimStateType := `{"claimantstatus": "}` + claimStatus + `", "claimstatuschanged": "` + claimStatusChanged + `"}`
	
	strActorType :=  `{"actorempid": "}` + actorEmpId + `", "actorname": "` + actorName + `", "actorrole": "` + actorRole + `", "actiondescription": "` + actionDesc + `"}`

	//Build Claim structure
	strClaim := `{"claimid": "}` + claimId + `", "claimdate": "` + claimDate + `", "claimdescription": "` + claimDesc + `", "claimantdetails": "` + strClaimantDetailsType + `", "claimedamount": "` + claimedAmount + `", "approvedamount": "` + approvedAmount + `", "claimstate": "` + strClaimStateType + `", "actordetails": "` + strActorType + `"}`
	
	

	err = stub.PutState(claimId, []byte(strClaim))									//store claim with id as key
	if err != nil {
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

	fmt.Printf("- end create_claim")
	return nil, nil
}

func (t *ClaimProcessing) update_claimStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//var err error
	fmt.Printf("- start update_claimStatus")




	fmt.Printf("- end update_claimStatus")
	return nil, nil
}

// ============================================================================================================================
// Query - Our entry point for Queries
// ============================================================================================================================
func (t *ClaimProcessing) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("query is running " + function)

	// Handle different functions
	if function == "getClaim" {													//read a variable
		return t.getClaim(stub, args)
	}else if function == "getClaimByClaimant" {													//read a variable
		return t.getClaimByClaimant(stub, args)
	}
	fmt.Printf("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query")
}  

func (t *ClaimProcessing) getClaim(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var claimId, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
	}

	claimId = args[0]
	valAsbytes, err := stub.GetState(claimId)									//get the var from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + claimId + "\"}"
		return nil, errors.New(jsonResp)
	}
	objClaim := Claim{}
	json.Unmarshal(valAsbytes, &objClaim)
	
	strClaimantDetailsType := `{"claimantid": "}` + objClaim.ClaimantDetails.ClaimantId + `", "claimantname": ",` +  objClaim.ClaimantDetails.ClaimantName + `"}`
	
	
	strClaimStateType := `{"claimantstatus": "}` + objClaim.ClaimState.ClaimStatus + `", "claimstatuschanged": "` + objClaim.ClaimState.ClaimStatusChanged + `"}`
	
	strActorType :=  `{"actorempid": "}` + objClaim.ActorDetails.ActorEmpId + `", "actorname": "` + objClaim.ActorDetails.ActorName + `", "actorrole": "` + objClaim.ActorDetails.ActorRole + `", "actiondescription": "` + objClaim.ActorDetails.ActionDescription + `"}`
	
	jsonResp = `{"claimid": "}` + objClaim.ClaimId + `", "claimdate": "` + objClaim.ClaimDate + `", "claimdescription": "` + objClaim.ClaimDescription + `", "claimantdetails": "` + strClaimantDetailsType + `", "claimedamount": "` + objClaim.ClaimedAmount + `", "approvedamount": "` + objClaim.ApprovedAmount + `", "claimstate": "` + strClaimStateType + `", "actordetails": "` + strActorType + `"}`

	return []byte(jsonResp), nil													//send it onward
}

func (t *ClaimProcessing) getClaimByClaimant(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
	}
	
	valAsbytes, err := stub.GetState(claimantIndexMap_Key)									//get the var from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for claimantIndexMap_Key\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil													//send it onward
}