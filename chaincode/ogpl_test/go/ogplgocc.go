/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

 package main

 /* Imports
  * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
  * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
  */
 import (
	 "bytes"
	 "encoding/json"
	 "fmt"
	 "strconv"
 
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 sc "github.com/hyperledger/fabric/protos/peer"
 )
 
 // Define the Smart Contract structure
 type SmartContract struct {
 }
 
 // Define the dataset structure, with 4 properties.  Structure tags are used by encoding/json library
 type Dataset struct {
	 Name   	string `json:"name"`
	 Size		string `json:"size"`
	 Extension 	string `json:"extension"`
	 Timestamp  string `json:"timestamp"`
	 Owner  	string `json:"owner"`
	 Checksum	string `json:"checksum"`
 }
 
 /*
  * The Init method is called when the Smart Contract "fabdataset" is instantiated by the blockchain network
  * Best practice is to have any Ledger initialization in separate function -- see initLedger()
  */
 func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	 return shim.Success(nil)
 }
 
 /*
  * The Invoke method is called as a result of an application request to run the Smart Contract "fabdataset"
  * The calling application program has also specified the particular smart contract function to be called, with arguments
  */
 func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 // Retrieve the requested Smart Contract function and arguments
	 function, args := APIstub.GetFunctionAndParameters()
	 // Route to the appropriate handler function to interact with the ledger appropriately
	 if function == "queryDataset" {
		 return s.queryDataset(APIstub, args)
	 } else if function == "initLedger" {
		 return s.initLedger(APIstub)
	 } else if function == "createDataset" {
		 return s.createDataset(APIstub, args)
	 } else if function == "queryAllDatasets" {
		 return s.queryAllDatasets(APIstub)
	 } else if function == "changeDatasetOwner" {
		 return s.changeDatasetOwner(APIstub, args)
	 }
 
	 return shim.Error("Invalid Smart Contract function name.")
 }
 
 func (s *SmartContract) queryDataset(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1")
	 }
 
	 datasetAsBytes, _ := APIstub.GetState(args[0])
	 return shim.Success(datasetAsBytes)
 }
 
 func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	 datasets := []Dataset{
		Dataset{Name: "DatasetA", Size: "1400000", Owner: "Kha", Extension: "csv", Checksum:"abc"},
		Dataset{Name: "DatasetB", Size: "4400000", Owner: "Kha", Extension: "csv", Checksum:"abc"},
		Dataset{Name: "DatasetC", Size: "2400000", Owner: "Kha", Extension: "csv", Checksum:"abc"},
	 }
 
	 i := 0
	 for i < len(datasets) {
		 fmt.Println("i is ", i)
		 datasetAsBytes, _ := json.Marshal(datasets[i])
		 APIstub.PutState("DATASET"+strconv.Itoa(i), datasetAsBytes)
		 fmt.Println("Added", datasets[i])
		 i = i + 1
	 }
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) createDataset(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 6 {
		 return shim.Error("Incorrect number of arguments. Expecting 6")
	 }
	 var dataset = Dataset{Name: args[1], Size: args[2], Owner: args[3], Extension: args[4], Checksum: args[5]}
 
	 datasetAsBytes, _ := json.Marshal(dataset)
	 APIstub.PutState(args[0], datasetAsBytes)
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) queryAllDatasets(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 startKey := "DATASET0"
	 endKey := "DATASET999"
 
	 resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	 if err != nil {
		 return shim.Error(err.Error())
	 }
	 defer resultsIterator.Close()
 
	 // buffer is a JSON array containing QueryResults
	 var buffer bytes.Buffer
	 buffer.WriteString("[")
 
	 bArrayMemberAlreadyWritten := false
	 for resultsIterator.HasNext() {
		 queryResponse, err := resultsIterator.Next()
		 if err != nil {
			 return shim.Error(err.Error())
		 }
		 // Add a comma before array members, suppress it for the first array member
		 if bArrayMemberAlreadyWritten == true {
			 buffer.WriteString(",")
		 }
		 buffer.WriteString("{\"Key\":")
		 buffer.WriteString("\"")
		 buffer.WriteString(queryResponse.Key)
		 buffer.WriteString("\"")
 
		 buffer.WriteString(", \"Record\":")
		 // Record is a JSON object, so we write as-is
		 buffer.WriteString(string(queryResponse.Value))
		 buffer.WriteString("}")
		 bArrayMemberAlreadyWritten = true
	 }
	 buffer.WriteString("]")
 
	 fmt.Printf("- queryAllDatasets:\n%s\n", buffer.String())
 
	 return shim.Success(buffer.Bytes())
 }
 
 func (s *SmartContract) changeDatasetOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }
 
	 datasetAsBytes, _ := APIstub.GetState(args[0])
	 dataset := Dataset{}
 
	 json.Unmarshal(datasetAsBytes, &dataset)
	 dataset.Owner = args[1]
 
	 datasetAsBytes, _ = json.Marshal(dataset)
	 APIstub.PutState(args[0], datasetAsBytes)
 
	 return shim.Success(nil)
 }
 
 // The main function is only relevant in unit test mode. Only included here for completeness.
 func main() {
 
	 // Create a new Smart Contract
	 err := shim.Start(new(SmartContract))
	 if err != nil {
		 fmt.Printf("Error creating new Smart Contract: %s", err)
	 }
 }
 