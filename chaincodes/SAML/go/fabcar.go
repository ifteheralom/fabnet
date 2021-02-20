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
 
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 sc "github.com/hyperledger/fabric/protos/peer"
 )
 
 // Define the Smart Contract structure
 type SmartContract struct {
 }
 
 var j = 0
 var m = 0
 var i = 0
 
 // Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
 type Car struct {
	 Make   string `json:"make"`
	 Model  string `json:"model"`
	 Colour string `json:"colour"`
	 Owner  string `json:"owner"`
 }
 type metaDataStore struct {
	 Doctype  string //`json:"docType"`
	 User     string //`json:"user"`
	 Metadata string //`json:"metaData"`
	 Key      string
 }
 type list struct {
	 Tal string
 }
 type talList struct {
	 Doctype  string //`json:"docType"`
	 EntityID string //`json:"user"`
	 TList    []list //`json:"metaData"`
	 Key      string
 }
 type codeStore struct {
	 Doctype    string //`json:"docType"`
	 ForWhichSP string ///`json:"forWhichSp"`
	 WhichIDP   string //`json:"whichIdp"`
	 Code       string //`json:"code"`
	 Key        string
 }
 type newCodeStore struct {
	 Doctype    string //`json:"docType"`
	 ForWhichSP string ///`json:"forWhichSp"`
	 WhichIDP   string //`json:"whichIdp"`
	 SPCode     string //`json:"code"`
	 IDPCode    string
	 SPCheck    string //`json:"code"`
	 IDPCheck   string
	 Key        string
 }
 type queryResultMetaData struct {
	 Key    string //`json:"Key"`
	 Record *metaDataStore
 }
 type queryResultCode struct {
	 Key    string //`json:"Key"`
	 Record *codeStore
 }
 type queryResultTalList struct {
	 Key    string //`json:"Key"`
	 Record *talList
 }
 type queryResultNewCode struct {
	 Key    string //`json:"Key"`
	 Record *newCodeStore
 }
 
 /*
  * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
  * Best practice is to have any Ledger initialization in separate function -- see initLedger()
  */
 func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	 return shim.Success(nil)
 }
 
 /*
  * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
  * The calling application program has also specified the particular smart contract function to be called, with arguments
  */
 func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 // Retrieve the requested Smart Contract function and arguments
	 function, args := APIstub.GetFunctionAndParameters()
	 // Route to the appropriate handler function to interact with the ledger appropriately
	 if function == "queryCar" {
		 return s.queryCar(APIstub, args)
	 } else if function == "initLedger" {
		 return s.initLedger(APIstub)
	 } else if function == "createCar" {
		 return s.createCar(APIstub, args)
	 } else if function == "queryAllCars" {
		 return s.queryAllCars(APIstub)
	 } else if function == "changeCarOwner" {
		 return s.changeCarOwner(APIstub, args)
	 } else if function == "storeMetaData" {
		 return s.storeMetaData(APIstub, args)
	 } else if function == "storeTalList" {
		 return s.storeTalList(APIstub, args)
	 } else if function == "talListDelete" {
		 return s.talListDelete(APIstub, args)
	 } else if function == "returnTalList" {
		 return s.returnTalList(APIstub, args)
	 } else if function == "storeCode" {
		 return s.storeCode(APIstub, args)
	 } else if function == "approval" {
		 return s.approval(APIstub, args)
	 } else if function == "metaDataFetch" {
		 return s.metaDataFetch(APIstub, args)
	 } else if function == "removeApproval" {
		 return s.removeApproval(APIstub, args)
	 } else if function == "codeInvoke" {
		 return s.codeInvoke(APIstub, args)
	 } else if function == "codeCheck" {
		 return s.codeCheck(APIstub, args)
	 } else if function == "codeFetch" {
		 return s.codeFetch(APIstub, args)
	 }
 
	 return shim.Error("Invalid Smart Contract function name.")
 }
 
 func (s *SmartContract) queryCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1")
	 }
 
	 carAsBytes, _ := APIstub.GetState(args[0])
	 return shim.Success(carAsBytes)
 }
 
 func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	 metaDatas := metaDataStore{
		 Doctype:  "MetaData Store",
		 User:     "www.idp.org",
		 Metadata: "entityid: \"https://mail.service.com/service/extension/samlreceiver \",\n  contacts: [],\n  \"metadata-set\": \"saml20-sp-remote\",\n  AssertionConsumerService: [\n    {\n      Binding: \"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\",\n      Location: \"https://mail.service.com/service/extension/samlreceiver\",\n      index: 0,\n    },\n  ],\n  SingleLogoutService: [],\n  \"validate.authnrequest\": false,\n  \"NameIDFormat\": \"urn:oasis:names:tc:",
		 Key:      "000",
	 }
 
	 // for i, metaData := range metaDatas {
	 // 	metaDataBytes, _ := json.Marshal(metaData)
	 // 	 APIstub.PutState("MetaData"+strconv.Itoa(i), metaDataBytes)
	 // }
	 metaDataBytes, _ := json.Marshal(metaDatas)
	 APIstub.PutState(metaDatas.Key, metaDataBytes)
 
	 talListIdp := talList{
		 Doctype:  "TAL List",
		 EntityID: "www.idp.sust.com",
		 TList: []list{
			 {Tal: "http://sp1.sust.com/simplesaml/module.php/saml/sp/metadata.php/default-sp"},
			 {Tal: "http://sp2.sust.com/simplesaml/module.php/saml/sp/metadata.php/default-sp"},
			 {Tal: "http://code.sust.com/simplesaml/module.php/saml/sp/metadata.php/default-sp"},
			 {Tal: "http://3.17.180.157:3000/mailmetadata"},
		 },
		 Key: "0001",
	 }
	 talListBytes, _ := json.Marshal(talListIdp)
	 APIstub.PutState(talListIdp.Key, talListBytes)
 
	 talListSp1 := talList{
		 Doctype:  "TAL List",
		 EntityID: "www.sp1.sust.com",
		 TList: []list{
			 {Tal: "http://idp.sust.com/simplesaml/saml2/idp/metadata.php"},
		 },
		 Key: "0002",
	 }
	 talListBytes1, _ := json.Marshal(talListSp1)
	 APIstub.PutState(talListSp1.Key, talListBytes1)
 
	 talListSp2 := talList{
		 Doctype:  "TAL List",
		 EntityID: "www.sp2.sust.com",
		 TList: []list{
			 {Tal: "http://idp.sust.com/simplesaml/saml2/idp/metadata.php"},
		 },
		 Key: "0003",
	 }
	 talListBytes2, _ := json.Marshal(talListSp2)
	 APIstub.PutState(talListSp2.Key, talListBytes2)
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) createCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 5 {
		 return shim.Error("Incorrect number of arguments. Expecting 5")
	 }
 
	 var car = Car{Make: args[1], Model: args[2], Colour: args[3], Owner: args[4]}
 
	 carAsBytes, _ := json.Marshal(car)
	 APIstub.PutState(args[0], carAsBytes)
	 /////create car
	 return shim.Success(nil)
 }
 
 /// store metadata is working absulately fine
 func (s *SmartContract) storeMetaData(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }
	 user := args[0]
	 metaData := args[1]
	 //var b string
	 //result := userFetch(APIstub, args)
	 //fmt.Println(result)
	 j++
	 b := string(j)
	 queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"MetaData Store\",\"User\": \"%s\"}}", user)
	 queryResults, _ := APIstub.GetQueryResult(queryString)
	 var codeData metaDataStore
	 for queryResults.HasNext() {
		 queryResultsData, _ := queryResults.Next()
		 _ = json.Unmarshal(queryResultsData.Value, &codeData)
	 }
	 if codeData.User != user {
		 metadataStore := metaDataStore{
			 Doctype:  "MetaData Store",
			 User:     user,
			 Metadata: metaData,
			 Key:      b,
		 }
		 // metadataStore := metaDataStore{
		 // 	Doctype:  "MetaData Store",
		 // 	User:     user,
		 // 	Metadata: "ACHE already",
		 // 	Key:      b,
		 // }
		 metaDataBytes, _ := json.Marshal(metadataStore)
		 APIstub.PutState(metadataStore.Key, metaDataBytes)
		 //return shim.Success(nil)
	 } else {
		 metadataStore := metaDataStore{
			 Doctype:  "MetaData Store",
			 User:     user,
			 Metadata: "ACHE already",
			 Key:      b,
		 }
		 // metadataStore := metaDataStore{
		 // 	Doctype:  "MetaData Store",
		 // 	User:     user,
		 // 	Metadata: metaData,
		 // 	Key:      b,
		 // }
		 metaDataBytes, _ := json.Marshal(metadataStore)
		 APIstub.PutState(codeData.Key, metaDataBytes)
		 //return shim.Success(nil)
	 }
	 return shim.Success(nil)
 
 }
 
 //talList function works perfectly
 func (s *SmartContract) storeTalList(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }
	 m++
	 key := string(m)
	 entityID := args[0]
	 tal := args[1]
	 queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"TAL List\",\"EntityID\": \"%s\"}}", entityID)
	 queryResults, _ := APIstub.GetQueryResult(queryString)
	 defer queryResults.Close()
	 var codeData talList
	 for queryResults.HasNext() {
		 queryResultsData, _ := queryResults.Next()
		 _ = json.Unmarshal(queryResultsData.Value, &codeData)
	 }
	 //data := MetaDataStore{}
	 // var results []QueryResultTalList
	 tallist := talList{
		 Doctype:  "TAL List",
		 EntityID: entityID,
		 TList: []list{
			 {Tal: tal},
		 },
		 Key: key,
	 }
	 talListBytes, _ := json.Marshal(tallist)
	 if codeData.EntityID == args[0] {
		 l := list{}
		 l.Tal = tal
		 codeData.TList = append(codeData.TList, l)
		 codeDataBytes, _ := json.Marshal(codeData)
		 APIstub.PutState(codeData.Key, codeDataBytes)
	 } else {
		 APIstub.PutState(tallist.Key, talListBytes)
	 }
	 return shim.Success(nil)
 
 }
 
 ///////   storecode run smothly
 
 func (s *SmartContract) storeCode(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 7 {
		 return shim.Error("Incorrect number of arguments. Expecting 7 ")
	 }
	 forWhichSp := args[0]
	 whichIdp := args[1]
	 spCode := args[2]
	 idpCode := args[3]
	 spCheck := args[4]
	 idpCheck := args[5]
	 author := args[6]
	 i++
	 key := string(i)
	 queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"Code Store\",\"ForWhichSP\": \"%s\", \"WhichIDP\": \"%s\"}}", forWhichSp, whichIdp)
	 resultsIterator, _ := APIstub.GetQueryResult(queryString)
	 defer resultsIterator.Close()
	 var codeData newCodeStore
	 for resultsIterator.HasNext() {
		 queryResponse, _ := resultsIterator.Next()
		 _ = json.Unmarshal(queryResponse.Value, codeData)
	 }
	 codeStore := newCodeStore{
		 Doctype:    "Code Store",
		 ForWhichSP: forWhichSp,
		 WhichIDP:   whichIdp,
		 SPCode:     spCode,
		 IDPCode:    idpCode,
		 SPCheck:    spCheck,
		 IDPCheck:   idpCheck,
		 Key:        key,
	 }
 
	 if codeData.ForWhichSP == forWhichSp && codeData.WhichIDP == whichIdp {
		 if author == "sp" {
			 codeData.SPCode = spCode
		 } else if author == "idp" {
			 codeData.IDPCode = idpCode
		 }
		 codeDataBytes, _ := json.Marshal(codeData)
		 APIstub.PutState(codeData.Key, codeDataBytes)
	 } else {
		 codeDataBytes, _ := json.Marshal(codeStore)
		 APIstub.PutState(codeStore.Key, codeDataBytes)
	 }
	 return shim.Success(nil)
 
 }
 
 ///tallIstDelete works perfectly
 func (s *SmartContract) talListDelete(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 1")
	 }
	 entityID := args[0]
	 tal := args[1]
 
	 queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"TAL List\",\"EntityID\": \"%s\"}}", entityID)
	 queryResults, _ := APIstub.GetQueryResult(queryString)
	 var codeData talList
	 for queryResults.HasNext() {
		 queryResultsData, _ := queryResults.Next()
		 _ = json.Unmarshal(queryResultsData.Value, &codeData)
	 }
	 var pos int
	 for _, service := range codeData.TList {
		 if service.Tal == tal {
			 codeData.TList = append(codeData.TList[:pos], codeData.TList[pos+1:]...)
			 if pos > 0 {
				 pos = pos - 1
			 }
			 continue
		 }
		 pos++
	 }
	 codeDataBytes, _ := json.Marshal(codeData)
	 APIstub.PutState(codeData.Key, codeDataBytes)
	 return shim.Success(nil)
 }
 
 //approval function works perfectly
 func (s *SmartContract) approval(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1 ")
	 }
	 author := args[0]
	 queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"Code Store\"}}")
	 resultsIterator, _ := APIstub.GetQueryResult(queryString)
	 defer resultsIterator.Close()
	 codeData := new(newCodeStore)
	 var results []queryResultNewCode
	 //var results []byte
	 //	var queryResults []byte
	 for resultsIterator.HasNext() {
		 queryResponse, _ := resultsIterator.Next()
		 _ = json.Unmarshal(queryResponse.Value, codeData)
		 if codeData.ForWhichSP == author {
			 queryResult := queryResultNewCode{Key: queryResponse.Key, Record: codeData}
			 //queryResult := QueryResultNewCode{Key: codeData.Key, Record: codeData}
			 results = append(results, queryResult)
		 } else if codeData.WhichIDP == author {
			 queryResult := queryResultNewCode{Key: queryResponse.Key, Record: codeData}
			 //queryResult := QueryResultNewCode{Key: codeData.Key, Record: codeData}
			 results = append(results, queryResult)
			 //	queryResult := queryResultNewCode{Key: queryResponse.Key, Record: codeData}
			 //queryResult := QueryResultNewCode{Key: codeData.Key, Record: codeData}
			 //	results = codeData
			 //results = codeData.WhichIDP
		 }
	 }
	 approveResult := results
	 approveResultBytes := new(bytes.Buffer)
	 json.NewEncoder(approveResultBytes).Encode(approveResult)
 
	 return shim.Success(approveResultBytes.Bytes())
 }
 
 ////remove approval works perfectly
 func (s *SmartContract) removeApproval(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 1 ")
	 }
	 forWhichSp := args[0]
	 whichIdp := args[1]
	 queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"Code Store\",\"ForWhichSP\": \"%s\", \"WhichIDP\": \"%s\"}}", forWhichSp, whichIdp)
	 resultsIterator, _ := APIstub.GetQueryResult(queryString)
	 defer resultsIterator.Close()
	 var codeData newCodeStore
	 for resultsIterator.HasNext() {
		 queryResponse, _ := resultsIterator.Next()
		 _ = json.Unmarshal(queryResponse.Value, &codeData)
	 }
	 //var result string
	 if codeData.ForWhichSP == forWhichSp && codeData.WhichIDP == whichIdp {
 
		 APIstub.DelState(codeData.Key)
 
	 }
	 return shim.Success(nil)
 }
 
 //code invoke works perfectly
 func (s *SmartContract) codeInvoke(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 4 {
		 return shim.Error("Incorrect number of arguments. Expecting 4 ")
	 }
	 forWhichSp := args[0]
	 whichIdp := args[1]
	 author := args[2]
	 code := args[3]
 
	 queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"Code Store\",\"ForWhichSP\": \"%s\", \"WhichIDP\": \"%s\"}}", forWhichSp, whichIdp)
	 resultsIterator, _ := APIstub.GetQueryResult(queryString)
	 defer resultsIterator.Close()
	 var codeData newCodeStore
	 for resultsIterator.HasNext() {
		 queryResponse, _ := resultsIterator.Next()
		 _ = json.Unmarshal(queryResponse.Value, &codeData)
	 }
	 //	var result = codeData.Key
	 if codeData.ForWhichSP == forWhichSp && codeData.WhichIDP == whichIdp {
		 if codeData.ForWhichSP == forWhichSp && codeData.WhichIDP == whichIdp {
			 if author == "sp" && codeData.SPCode == code {
				 codeData.SPCheck = "success"
				 codeDataBytes, _ := json.Marshal(codeData)
				 APIstub.PutState(codeData.Key, codeDataBytes)
				 //	result = "sp-success"
				 //result = codeData.SPCheck
			 } else if author == "idp" && codeData.IDPCode == code {
				 codeData.IDPCheck = "success"
				 codeDataBytes, _ := json.Marshal(codeData)
				 APIstub.PutState(codeData.Key, codeDataBytes)
				 //	result = "idp-success"
 
			 } else {
				 //result = "code-failed"
			 }
		 } else {
			 //result = "code-failed"
		 }
	 }
	 //	return  shim.Success([]byte(result))
	 return shim.Success(nil)
 }
 
 //code check works perfectly
 func (s *SmartContract) codeCheck(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 3 {
		 return shim.Error("Incorrect number of arguments. Expecting 4 ")
	 }
	 forWhichSp := args[0]
	 whichIdp := args[1]
	 author := args[2]
	 //code := args[3]
 
	 queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"Code Store\",\"ForWhichSP\": \"%s\", \"WhichIDP\": \"%s\"}}", forWhichSp, whichIdp)
	 resultsIterator, _ := APIstub.GetQueryResult(queryString)
	 defer resultsIterator.Close()
	 var codeData newCodeStore
	 for resultsIterator.HasNext() {
		 queryResponse, _ := resultsIterator.Next()
		 _ = json.Unmarshal(queryResponse.Value, &codeData)
	 }
	 result := "hi"
	 if codeData.ForWhichSP == forWhichSp && codeData.WhichIDP == whichIdp {
 
		 if codeData.ForWhichSP == forWhichSp && codeData.WhichIDP == whichIdp {
			 if author == "sp" && codeData.SPCheck == "success" {
				 result = "sp-success"
				 //result = codeData.SPCheck
			 } else if author == "idp" && codeData.IDPCheck == "success" {
				 result = "idp-success"
 
			 } else {
				 result = "code-failed"
			 }
		 } else {
			 result = "code-failed"
		 }
	 }
	 return shim.Success([]byte(result))
	 //return shim.Success(nil)
 }
 
 func (s *SmartContract) codeFetch(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 4 {
		 return shim.Error("Incorrect number of arguments. Expecting 4 ")
	 }
	 forWhichSp := args[0]
	 whichIdp := args[1]
	 author := args[2]
	 code := args[3]
 
	 queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"Code Store\",\"ForWhichSP\": \"%s\", \"WhichIDP\": \"%s\"}}", forWhichSp, whichIdp)
	 resultsIterator, _ := APIstub.GetQueryResult(queryString)
	 defer resultsIterator.Close()
	 var codeData newCodeStore
	 for resultsIterator.HasNext() {
		 queryResponse, _ := resultsIterator.Next()
		 _ = json.Unmarshal(queryResponse.Value, &codeData)
	 }
	 var result = codeData.Key
	 if codeData.ForWhichSP == forWhichSp && codeData.WhichIDP == whichIdp {
		 if codeData.ForWhichSP == forWhichSp && codeData.WhichIDP == whichIdp {
			 if author == "sp" && codeData.SPCode == code {
				 codeData.SPCheck = "success"
				 codeDataBytes, _ := json.Marshal(codeData)
				 APIstub.PutState(codeData.Key, codeDataBytes)
				 result = "sp-success"
				 //result = codeData.SPCheck
			 } else if author == "idp" && codeData.IDPCode == code {
				 codeData.IDPCheck = "success"
				 codeDataBytes, _ := json.Marshal(codeData)
				 APIstub.PutState(codeData.Key, codeDataBytes)
				 result = "idp-success"
 
			 } else {
				 result = "code-failed"
			 }
		 } else {
			 result = "code-failed"
		 }
	 }
	 return shim.Success([]byte(result))
	 //return shim.Success(nil)
 }
 
 ///perfectly return metadata
 func (s *SmartContract) metaDataFetch(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	 user := args[0]
	 queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"MetaData Store\",\"User\": \"%s\"}}", user)
 
	 ///	queryResults, _ := getQueryResultForQueryString(APIstub, queryString)
	 queryResults, _ := APIstub.GetQueryResult(queryString)
	 var codeData metaDataStore
	 for queryResults.HasNext() {
		 queryResultsData, _ := queryResults.Next()
		 _ = json.Unmarshal(queryResultsData.Value, &codeData)
	 }
	 arr := []byte(codeData.Metadata)
	 return shim.Success(arr)
 }
 
 //perfectly return tallits
 func (s *SmartContract) returnTalList(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	 entityID := args[0]
	 //tal := args[1]
	 queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"TAL List\",\"EntityID\": \"%s\"}}", entityID)
	 queryResults, _ := APIstub.GetQueryResult(queryString)
	 defer queryResults.Close()
	 var codeData talList
	 for queryResults.HasNext() {
		 queryResultsData, _ := queryResults.Next()
		 _ = json.Unmarshal(queryResultsData.Value, &codeData)
	 }
	 tLists := codeData.TList
	 tListsBytes := new(bytes.Buffer)
	 json.NewEncoder(tListsBytes).Encode(tLists)
 
	 return shim.Success(tListsBytes.Bytes())
 }
 
 func (s *SmartContract) talList(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 1")
	 }
	 entityID := args[0]
	 tal := args[1]
 
	 queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"TAL List\",\"EntityID\": \"%s\"}}", entityID)
	 queryResults, _ := APIstub.GetQueryResult(queryString)
	 var codeData talList
	 for queryResults.HasNext() {
		 queryResultsData, _ := queryResults.Next()
		 _ = json.Unmarshal(queryResultsData.Value, &codeData)
	 }
	 var pos int
	 for _, service := range codeData.TList {
		 if service.Tal == tal {
			 codeData.TList = append(codeData.TList[:pos], codeData.TList[pos+1:]...)
			 if pos > 0 {
				 pos = pos - 1
			 }
			 continue
		 }
		 pos++
	 }
	 codeDataBytes, _ := json.Marshal(codeData)
	 APIstub.PutState(codeData.Key, codeDataBytes)
	 return shim.Success(nil)
 }
 
 func userFetch(APIstub shim.ChaincodeStubInterface, args []string) string {
	 user := args[0]
	 queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"MetaData Store\",\"User\": \"%s\"}}", user)
	 queryResults, _ := APIstub.GetQueryResult(queryString)
	 var codeData metaDataStore
	 for queryResults.HasNext() {
		 queryResultsData, _ := queryResults.Next()
		 _ = json.Unmarshal(queryResultsData.Value, &codeData)
	 }
 
	 return codeData.User
 }
 
 func (s *SmartContract) entityFetch(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	 entityID := args[0]
	 //tal := args[1]
	 queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"TAL List\",\"EntityID\": \"%s\"}}", entityID)
	 queryResults, _ := APIstub.GetQueryResult(queryString)
	 defer queryResults.Close()
	 var codeData talList
	 for queryResults.HasNext() {
		 queryResultsData, _ := queryResults.Next()
		 _ = json.Unmarshal(queryResultsData.Value, &codeData)
	 }
	 return shim.Success([]byte(codeData.EntityID))
 }
 
 func getJSONQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
	 start := "{\"values\": "
	 end := "}"
	 data, err := getQueryResultForQueryString(stub, queryString)
	 if err != nil {
		 return nil, err
	 }
	 return []byte(start + string(data) + end), nil
 }
 
 func (s *SmartContract) fetch(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	 //user := strings.ToLower(args[0])
	 user := args[0]
	 queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"MetaData Store\",\"User\": \"%s\"}}", user)
	 queryResults, _ := getQueryResultForQueryString(APIstub, queryString)
	 // if  err != nil {
	 // 	return shim.Error(err.error())
	 // }
 
	 return shim.Success(queryResults)
 }
 func (s *SmartContract) queryAllCars(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 startKey := "CAR0"
	 endKey := "CAR999"
 
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
 
	 fmt.Printf("- queryAllCars:\n%s\n", buffer.String())
 
	 return shim.Success(buffer.Bytes())
 }
 
 func (s *SmartContract) changeCarOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }
 
	 carAsBytes, _ := APIstub.GetState(args[0])
	 car := Car{}
 
	 json.Unmarshal(carAsBytes, &car)
	 car.Owner = args[1]
 
	 carAsBytes, _ = json.Marshal(car)
	 APIstub.PutState(args[0], carAsBytes)
 
	 return shim.Success(nil)
 }
 
 ///get query result by query string
 
 func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
 
	 fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)
 
	 resultsIterator, err := stub.GetQueryResult(queryString)
	 if err != nil {
		 return nil, err
	 }
	 defer resultsIterator.Close()
 
	 // buffer is a JSON array containing QueryRecords
	 var buffer bytes.Buffer
	 buffer.WriteString("[")
 
	 bArrayMemberAlreadyWritten := false
	 for resultsIterator.HasNext() {
		 queryResponse, err := resultsIterator.Next()
		 if err != nil {
			 return nil, err
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
 
	 fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())
 
	 return buffer.Bytes(), nil
 }
 
 // The main function is only relevant in unit test mode. Only included here for completeness.
 func main() {
 
	 // Create a new Smart Contract
	 err := shim.Start(new(SmartContract))
	 if err != nil {
		 fmt.Printf("Error creating new Smart Contract: %s", err)
	 }
 }