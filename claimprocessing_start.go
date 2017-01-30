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
    "time"
    "strconv"
	"encoding/json"
	"strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// ClaimProcessing implementation is a sample PoC code to demonstrate the BlockChain Capabilities
type ClaimProcessing struct {
}

var claim_IndexStr = "claimid_index_key"	
var claimant_IndexStr = "claimantid_index_key"	

// The below structure would help us to track the number of claims by a claimant
type ClaimantIndex struct{
    ClaimantId string
    ClaimIds[] string
}
var arrayOfClaimantIndex[] ClaimantIndex



//Defining the required Structure for the PoC
type ClaimantDetailsType struct{
	ClaimantId string 					
	ClaimantName string 
   
}

type ActorType struct{
	ActorEmpId string 				
	ActorName string 
    ActorRole string 
	ActionDescription string 
}

type ClaimStateType{
    ClaimStatus string 
    ClaimStatusChanged string
}

type Claim struct{
    ClaimId string
    ClaimDate string
    ClaimDescription string 
    ClaimantDetails ClaimantDetailsType
    ClaimAmount string
    ClaimState ClaimStateType
    ActorDetails ActorType
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
func (t *ClaimProcessing) init(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var err error

	return nil, nil
}

// ============================================================================================================================
// Run - Our entry point
// ============================================================================================================================
func (t *ClaimProcessing) invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("Entry to Invoke " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.init(stub, args)
	} 
	fmt.Println("run did not find func: " + function)						//error

	return nil, errors.New("Received unknown function invocation")
}

// ============================================================================================================================
// Query - Our entry point for Queries
// ============================================================================================================================
func (t *ClaimProcessing) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "xxxxxxx" {													//read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query")
}