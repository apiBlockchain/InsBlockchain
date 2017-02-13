/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main
import (
	"errors"
	"fmt"
	"time"
	"encoding/json"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)
// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Address Record
type Location   struct {
	Street   	string   `json:"Street"`
	Unit   		string   `json:"Unit"`
	City 		string   `json:"City"`
	State 		string   `json:"State"`
	Zip 		string   `json:"Zip"`
}

// Check Record
type DataCheck   struct {
	Result   	string   `json:"Status"`
	Message   	string   `json:"Message"`
}


// Employee Record
type Employee   struct {
	Name   		string   `json:"Name"`
	Address 	Location  `json:"Address"`
	Email 		string   `json:"Email"`
	Phone 		string   `json:"Phone"`
	DOB  		 string   `json:"DOB"`
	Gender    	string   `json:"Gender"`
	EmployerID  string   `json:"EmployerID"`
	EmployeeID	string   `json:"EmployeeID"`
	Type		string   `json:"Type"`
	Status		string   `json:"Status"`
	StartDate	string   `json:"StartDate"`
	EndDate		string   `json:"EndDate"`
}


type Member struct {
	MemberName string `json:"MemberName"`
	MemberID  string   `json:"MemberID"`
	MemberDOB  string   `json:"MemberDOB"`
	SubscriberID string `json:"SubscriberID"`
}

//Coverage Record
type Coverage struct {
		CoverageName string  `json:"CoverageName"`
		CoverageType string `json:"CoverageType"`
		CarrierID string `json:"CarrierID"`
		GroupNum string `json:"GroupNum"`
		PlanCode string `json:"PlanCode"`
		SubscriberID string `json:"SubscriberID"`
	  SubscriberName string `json:"subsciberName"`
	  SubscriberDOB string `json:"subscriberDOB"`
		IsPrimary string `json:"isPrimary"`
	  StartDate string   `json:"startDate"`
	  EndDate string   `json:"EndDate"`
		AnnualDeductible int `json:"AnnualDeductible"`
		AnnualBenefitMaximum int `json:"AnnualBenefitMaximum"`
		LifetimeBenefitMaximum string `json:"LifetimeBenefitMaximum"`
		PreventiveCare  string `json:"PreventiveCare "`
		MinorRestorativeCare string `json:"MinorRestorativeCare"`
		MajorRestorativeCare string `json:"MajorRestorativeCare"`
		OrthodonticTreatment string `json:"OrthodonticTreatment"`
		OrthodonticLifetimeBenefitMaximum string `json:"OrthodonticLifetimeBenefitMaximum"`
		AnnualDeductibleBal int `json:"AnnualDeductibleBal"`
		AnnualBenefitMaximumBal int `json:"AnnualBenefitMaximumBal"`
		EmployeeID string `json:"EmployeeID"`
		MemberID string  `json:"MemberID"`
		EmployerID string   `json:"EmployerID"`
		Dependents []Member `json:"Dependents"`
		Premium string `json:"Premium"`
		}
//Array for storing all coverages
type AllCoverages struct{
	Coverages []Coverage `json:"Coverages"`
}

type ALLMembers struct{
	Members []Member `json:"Members"`
}

	var dental Coverage
// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	// Create an example employee
	// Employee Record
	var aliceSheen Employee;


	aliceSheen.Name   			=	"Alice Sheen";
	aliceSheen.Address.Street 	=   "451 Indian Rocks Rd S";
	aliceSheen.Address.City 	=   "Largo";
	aliceSheen.Address.State 	=   "FL";
	aliceSheen.Address.Zip 		=   "33770";

	aliceSheen.Email 			=	"alicesheen@gmail.com";
	aliceSheen.Phone 			=	"727-223-5432";
	aliceSheen.DOB   			=	"08/16/1970";
	aliceSheen.Gender    		=	"Female";
	aliceSheen.EmployerID  		=	"Global Industries";
	aliceSheen.EmployeeID		=	"294048";
	aliceSheen.Type				= 	"Full Time";
	aliceSheen.Status 			= 	"Active";
	aliceSheen.StartDate		= 	"10/14/2008";
	aliceSheen.EndDate			= 	"NA";

	jsonAsBytes, _ := json.Marshal(aliceSheen)
	err = stub.PutState(aliceSheen.EmployeeID, jsonAsBytes)
	if err != nil {
		fmt.Println("Error Creating Bank user account")
		return nil, err
	}

//create an array for storing all coverages , and store the array on the blockchain
var coverages AllCoverages
jsonAsBytes, _ = json.Marshal(coverages)
err = stub.PutState("allCvgs", jsonAsBytes)
if err != nil {
	return nil, err
}

//create an array for storing all members and store array on the blockchain

var members ALLMembers
jsonAsBytes, _ = json.Marshal(members)
err = stub.PutState("allMembrs", jsonAsBytes)
if err != nil {
	return nil, err
}
dental.EmployeeID="294048"
dental.MemberID="M-01"
var dep1 Member;
var dep2 Member;
dep1.MemberName ="Megan Sheen";
dep1.MemberID="M-03";
dep1.MemberDOB="08/20/1990";
dep1.SubscriberID="ba2345";
jsonAsBytes, _ = json.Marshal(dep1);
err= stub.PutState(dep1.MemberID, jsonAsBytes);
if err != nil {
	fmt.Println("Error Creating dependents")
	return nil, err
}

dep2.MemberName ="Wade Sheen";
dep2.MemberID="M-02";
dep2.MemberDOB="08/20/1961";
dep2.SubscriberID="ba2345";
jsonAsBytes, _ = json.Marshal(dep2);
err= stub.PutState(dep2.MemberID, jsonAsBytes);
if err != nil {
	fmt.Println("Error Creating dependents")
	return nil, err
}
		dental.Dependents=append(dental.Dependents,dep1);
		dental.Dependents=append(dental.Dependents,dep2);


return nil,nil
}
// ============================================================================================================================
// Run - Our entry point for Invocations - [LEGACY] obc-peer 4/25/2016
// ============================================================================================================================
func (t *SimpleChaincode) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("run is running " + function)
	return t.Invoke(stub, function, args)
}

// ============================================================================================================================
// Invoke - Our entry point for Invocations
// ============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)
	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
		} else if function == "addCoverage" {											//create a transaction
		return t.addCoverage(stub, args)
	}	else if function == "updateCoverage" {											//create a transaction
	return t.updateCoverage(stub, args)
}
	fmt.Println("invoke did not find func: " + function)					//error
	return nil, errors.New("Received unknown function invocation")
}

// ============================================================================================================================
// Query - Our entry point for Queries
// ============================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)
	if function == "getCoverages" { return t.getCoverages(stub, args[1]) }
	if function == "getBlockchainRecord" { return t.getBlockchainRecord(stub, args[1]) }
	if function == "getUserAccount" { return t.getUserAccount(stub, args[1]) }
	if function == "getEmployeeRecord" { return t.getEmployeeRecord(stub, args[1]) }
	if function == "verifyEmployment" { return t.verifyEmployment(stub, args[1]) }
	if function == "verifyCoverage" { return t.verifyCoverage(stub, args[1], args[2]) }

	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query")
}

// ============================================================================================================================
// Get Open Points member account from the blockchain
// ============================================================================================================================
func (t *SimpleChaincode) getUserAccount(stub shim.ChaincodeStubInterface, userId string)([]byte, error){

	fmt.Println("Start getUserAccount")
	fmt.Println("Looking for user with ID " + userId);

	//get the User index
	fdAsBytes, err := stub.GetState(userId)
	if err != nil {
		return nil, errors.New("Failed to get user account from blockchain")
	}

	return fdAsBytes, nil

}

func (t *SimpleChaincode) verifyEmployment(stub shim.ChaincodeStubInterface, subscriberId string)([]byte, error){

	// Get the insurance coverage record
	coverageAsBytes, err := stub.GetState(subscriberId)
	if err != nil {
		return nil, errors.New("Failed to get coverage from blockchain")
	}
	var insRecord Coverage
	json.Unmarshal(coverageAsBytes, &insRecord)


	// Get the employee record
	employeeAsBytes, err := stub.GetState(insRecord.EmployeeID)
	if err != nil {
		return nil, errors.New("Failed to get user account from blockchain")
	}

	// Unmarshall the employee record from the blockchain
	var empRecord Employee
	json.Unmarshal(employeeAsBytes, &empRecord)


	// Check that the employment is valid

	var checkResults DataCheck
	checkResults.Result 	= "Passed";
	checkResults.Message 	= "Primary policy holder is a valid employee."

	if empRecord.Type != "Full Time" {
		checkResults.Result 	= "Failed";
		checkResults.Message 	= "Primary policy holder is not a full-time employee."
	}

	if empRecord.Status != "Active" {
		checkResults.Result 	= "Failed";
		checkResults.Message 	= "Primary policy holder is not an active employee."
	}


	resAsBytes, _ := json.Marshal(checkResults)
	return resAsBytes, nil

}


func (t *SimpleChaincode) verifyCoverage(stub shim.ChaincodeStubInterface, subscriberId string, memberId string)([]byte, error){

	dateForDemo, _ := time.Parse("01/02/2006",  "01/15/2017")

	// Get the insurance coverage record
	coverageAsBytes, err := stub.GetState(subscriberId)
	if err != nil {
		return nil, errors.New("Failed to get coverage from blockchain")
	}
	var insRecord Coverage
	json.Unmarshal(coverageAsBytes, &insRecord)

	/// Add some magic to get the right dependent here


	var checkResults DataCheck
	checkResults.Result 	= "Passed";
	checkResults.Message 	= "Coverage is valid."


	// Check if age is > 26 ONLY if the person is NOT the primary member
	if	insRecord.MemberID != memberId {

		numDependents := len(insRecord.Dependents)

		for index := numDependents - 1; index >= 0; index-- {

			if insRecord.Dependents[index].MemberID == memberId{

				memberDobDate, _ := time.Parse("01/02/2006", insRecord.Dependents[index].MemberDOB)
				ageInYears, _, _, _, _, _ 	  := diff(memberDobDate, dateForDemo)

				if ageInYears >= 26 {
					checkResults.Result 	= "Failed";
					checkResults.Message 	= "Dependent is above legal minimum coverage age."
				}

			}

		}
	}

	startDate,_ 	:= time.Parse("2006-01-02",  insRecord.StartDate)
	endDate,_	:= time.Parse("2006-01-02",  insRecord.EndDate)

	// Check if plan is active
	if (!(dateForDemo.After(startDate) && dateForDemo.Before(endDate))) {
		checkResults.Result 	= "Failed";
		checkResults.Message 	= "Policy is not active."
	}

	// Check if annual benefit max has not yet been reached
	if (insRecord.AnnualBenefitMaximumBal  < 1 ) {
		checkResults.Result 	= "Failed";
		checkResults.Message 	= "Annual maximum has already been met."
	}

	resAsBytes, _ := json.Marshal(checkResults)
	return resAsBytes, nil

}


func (t *SimpleChaincode) getBlockchainRecord(stub shim.ChaincodeStubInterface, recordKey string)([]byte, error){

	fmt.Println("Start getBlockchainRecord")
	fmt.Println("Looking for user with ID " + recordKey);

	//get the User index
	fdAsBytes, err := stub.GetState(recordKey)
	if err != nil {
		return nil, errors.New("Failed to get user account from blockchain")
	}

	return fdAsBytes, nil

}


func (t *SimpleChaincode) getEmployeeRecord(stub shim.ChaincodeStubInterface, employeeId string)([]byte, error){

	fmt.Println("Start getEmployeeRecord")
	fmt.Println("Looking for Employee with ID " + employeeId);

	//get the User index
	fdAsBytes, err := stub.GetState(employeeId)
	if err != nil {
		return nil, errors.New("Failed to get Employee account from blockchain")
	}

	return fdAsBytes, nil

}
func (t *SimpleChaincode) getCoverages(stub shim.ChaincodeStubInterface, subscriberID string)([]byte, error){

	fmt.Println("Start Get Coverage")
	fmt.Println("Looking for Coverage for SubscriberID" + subscriberID);

	coverageAsBytes, err := stub.GetState(subscriberID)
	if err != nil {
		return nil, errors.New("Failed to get coverage from blockchain")
	}
	return coverageAsBytes, nil
}
//Update Coverage

func (t *SimpleChaincode) updateCoverage(stub shim.ChaincodeStubInterface, args []string)([]byte, error){

			var dentalcoverage Coverage
			var subscriberID string
  		var  benefitMaximumBal int
			var deductibleBal int
			subscriberID=args[0];
  	  deductibleBal, err := strconv.Atoi(args[1]);
			benefitMaximumBal, err =strconv.Atoi(args[2]);
			coverageAsBytes, err := stub.GetState(subscriberID)
			fmt.Println("In Update coverage");
			if err != nil {
				return nil, errors.New("Failed to get Coverage from blopckchain")
			}
		json.Unmarshal(coverageAsBytes,&dentalcoverage)
		dentalcoverage.AnnualDeductibleBal=deductibleBal;
		dentalcoverage.AnnualBenefitMaximumBal=benefitMaximumBal;
		dentalcvgAsBytes, _ := json.Marshal(dentalcoverage)
		err = stub.PutState(dental.SubscriberID,dentalcvgAsBytes)
		if err != nil {
		fmt.Println("Error updating coverage")
		return nil, err
		}
			return nil,nil

		}
func (t *SimpleChaincode) addCoverage(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
fmt.Println("In addcoverage chaincode");
	//var dental Coverage
	var AnnualBenefitMaximum int
	dental.CoverageName=args[0]
	dental.CoverageType=args[1]
 	dental.CarrierID=args[2]
 	dental.GroupNum=args[3]
	dental.PlanCode=args[4]
 	dental.SubscriberID=args[5]
 	dental.SubscriberName=args[6]
	dental.SubscriberDOB=args[7]
	dental.IsPrimary=args[8]
	dental.EndDate=args[9]
	dental.StartDate=args[10]
  AnnualDeductible, err := strconv.Atoi(args[11])
	dental.AnnualDeductible=AnnualDeductible
	AnnualBenefitMaximum, err= strconv.Atoi(args[12])
	dental.AnnualBenefitMaximum=AnnualBenefitMaximum
	dental.LifetimeBenefitMaximum=args[13]
	dental.PreventiveCare =args[14]
	dental.MinorRestorativeCare=args[15]
	dental.MajorRestorativeCare=args[16]
	dental.OrthodonticTreatment=args[17]
	dental.OrthodonticLifetimeBenefitMaximum=args[18]
	dental.AnnualDeductibleBal=850
	dental.AnnualBenefitMaximumBal=5000
	//dental.EmployeeID="294048"
	//dental.MemberID="M-01"
	dental.EmployerID=args[19]

		// // Adding Dependents
		// var dep1 Member;
		// var dep2 Member;
		// dep1.MemberName ="Megan Sheen";
		// dep1.MemberID="M-03";
		// dep1.MemberDOB="08/20/1990";
		// dep1.SubscriberID="ba2345";
		// jsonAsBytes, _ := json.Marshal(dep1);
		// err= stub.PutState(dep1.MemberID, jsonAsBytes);
		// if err != nil {
		// 	fmt.Println("Error Creating dependents")
		// 	return nil, err
		// }
		//
		// dep2.MemberName ="Wade Sheen";
		// dep2.MemberID="M-02";
		// dep2.MemberDOB="08/20/1961";
		// dep2.SubscriberID="ba2345";
		// jsonAsBytes, _ = json.Marshal(dep2);
		// err= stub.PutState(dep2.MemberID, jsonAsBytes);
		// if err != nil {
		// 	fmt.Println("Error Creating dependents")
		// 	return nil, err
		// }
   // 			dental.Dependents=append(dental.Dependents,dep1);
		// 		dental.Dependents=append(dental.Dependents,dep2);
			  dental.Premium=args[20]
//new code for single struct
dentalAsBytes, _ := json.Marshal(dental)
err = stub.PutState(dental.SubscriberID, dentalAsBytes)
if err != nil {
	fmt.Println("Error adding Coverage")
	return nil, err
}

//new code for single struct
	return nil,nil


}
	// end add coverage
func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
