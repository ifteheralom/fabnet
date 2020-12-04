/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	_ "github.com/hyperledger/fabric-contract-api-go/contractapi/utils"
	_ "github.com/hyperledger/fabric-protos-go/peer"
	//_"google.golang.org/genproto/googleapis/privacy/dlp/v2"
	"strconv"

	//"golang.org/x/text/message"
	_ "math/rand"
)

// SmartContract provides functions for managing a car
var i=0
var j= 0
var m = 0
type SmartContract struct {
	contractapi.Contract
}

// Car describes basic details of what makes up a car
type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
}
type MetaDataStore struct {
	Doctype string //`json:"docType"`
	User    string //`json:"user"`
	Metadata string //`json:"metaData"`
	//Key string
}
type List struct{
	Tal string
}
type TalList struct {
	Doctype string //`json:"docType"`
	EntityId    string //`json:"user"`
	TList []List//`json:"metaData"`
	Key string
}
type CodeStore struct {
	Doctype string //`json:"docType"`
	ForWhichSP    string ///`json:"forWhichSp"`
	WhichIDP  string //`json:"whichIdp"`
	Code string//`json:"code"`
	Key string
}
type NewCodeStore struct {
	Doctype string //`json:"docType"`
	ForWhichSP    string ///`json:"forWhichSp"`
	WhichIDP  string //`json:"whichIdp"`
	SPCode string//`json:"code"`
	IDPCode string
	SPCheck string//`json:"code"`
	IDPCheck string
	Key string
}


// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string //`json:"Key"`
	Record *Car
}
type QueryResultMetaData struct {
	Key    string //`json:"Key"`
	Record *MetaDataStore
}
type QueryResultCode struct {
	Key string //`json:"Key"`
	Record *CodeStore
}
type QueryResultTalList struct {
	Key string //`json:"Key"`
	Record *TalList
}
type QueryResultNewCode struct {
	Key string //`json:"Key"`
	Record *NewCodeStore
}

// InitLedger adds a base set of cars to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	metaDatas := [] MetaDataStore{
		{ Doctype: "MetaData Store", User: "www.idp1.org", Metadata: "entityid: \"https://mail.service.com/service/extension/samlreceiver \",\n  contacts: [],\n  \"metadata-set\": \"saml20-sp-remote\",\n  AssertionConsumerService: [\n    {\n      Binding: \"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\",\n      Location: \"https://mail.service.com/service/extension/samlreceiver\",\n      index: 0,\n    },\n  ],\n  SingleLogoutService: [],\n  \"validate.authnrequest\": false,\n  \"NameIDFormat\": \"urn:oasis:names:tc:"},
	}
	for i, metaData := range metaDatas {
		metaDataBytes, _ := json.Marshal(metaData)
		err := ctx.GetStub().PutState("MetaData"+strconv.Itoa(i), metaDataBytes)

		if err != nil {
			return fmt.Errorf("failed to put to world state. %s", err.Error())
		}

	}

	talList := TalList{
		Doctype: "TAL List",
		EntityId :  "www.idp.sust.com",
		TList: [] List{
			{Tal : "http://sp1.sust.com/simplesaml/module.php/saml/sp/metadata.php/default-sp" },
			{Tal : "http://sp2.sust.com/simplesaml/module.php/saml/sp/metadata.php/default-sp"},
			{Tal : "http://code.sust.com/simplesaml/module.php/saml/sp/metadata.php/default-sp"},
			{Tal : "http://18.191.122.156:3000/mailmetadata" },
		},
		Key: "0001",
	}
	talListBytes, _ := json.Marshal(talList)
	 ctx.GetStub().PutState(talList.Key,talListBytes)

	talListSp1 := TalList{
		Doctype: "TAL List",
		EntityId :  "www.sp1.sust.com",
		TList: [] List{
			{Tal :  "http://idp.sust.com/simplesaml/saml2/idp/metadata.php" },
		},
		Key: "0002",
	}
	talListBytes1, _ := json.Marshal(talListSp1)
	ctx.GetStub().PutState(talListSp1.Key,talListBytes1)

	talListSp2 := TalList{
		Doctype: "TAL List",
		EntityId :  "www.sp2.sust.com",
		TList: [] List{
			{Tal :  "http://idp.sust.com/simplesaml/saml2/idp/metadata.php"},
		},
		Key: "0003",
	}
	talListBytes2, _ := json.Marshal(talListSp2)
	ctx.GetStub().PutState(talListSp2.Key,talListBytes2)
	talListBytes2,  _ = json.Marshal(talListSp2)
	ctx.GetStub().PutState(talListSp2.Key,talListBytes2)

	return nil
}

//create MetaData
func (s *SmartContract) StoreMetaData(ctx contractapi.TransactionContextInterface, user string, metaData string) error {
	//var b string
	result := s.UserFetch(ctx,user)
	j++
	b := string(j)
	if  result != user{
		metaDataStore := MetaDataStore {
			Doctype: "MetaData Store",
			User  : user,
			Metadata : metaData,
			//	Key : key,
		}
		metaDataBytes, _ := json.Marshal(metaDataStore)
		return ctx.GetStub().PutState(b, metaDataBytes)

	} else {
		return nil
	}
}
//which IDP store which SP code

func (s *SmartContract) Approval(ctx contractapi.TransactionContextInterface, author string) []QueryResultNewCode{
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"Code Store\"}}")
	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
	defer resultsIterator.Close()
	codeData := new(NewCodeStore)
	//codeData2 := new(NewCodeStore)
	var results  []QueryResultNewCode
	//results := " "
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		_ = json.Unmarshal(queryResponse.Value, codeData)
		if codeData.ForWhichSP == author  {
			queryResult := QueryResultNewCode{Key: queryResponse.Key, Record: codeData}
			//queryResult := QueryResultNewCode{Key: codeData.Key, Record: codeData}
			results = append(results, queryResult)
			//results = codeData.ForWhichSP
		} else	if codeData.WhichIDP == author {
			queryResult := QueryResultNewCode{Key: queryResponse.Key, Record: codeData}
				//queryResult := QueryResultNewCode{Key: codeData.Key, Record: codeData}
				results = append(results, queryResult)
			//results = codeData.WhichIDP
		}
		//queryResult := QueryResultNewCode{Key: queryResponse.Key, Record: codeData}
		//results = append(results, queryResult)
	}
	//for i2 := range results {
	//	if results[i2].Record == codeData2{
	//		if codeData2.ForWhichSP == author  {
	//			//queryResult := QueryResultNewCode{Key: queryResponse.Key, Record: codeData}
	//			queryResult := QueryResultNewCode{Key: codeData2.Key, Record: codeData2}
	//			results2 = append(results2, queryResult)
	//			//results = codeData.ForWhichSP
	//		} else	if codeData2.WhichIDP == author {
	//			//queryResult := QueryResultNewCode{Key: queryResponse.Key, Record: codeData}
	//			queryResult := QueryResultNewCode{Key: codeData2.Key, Record: codeData2}
	//			results2 = append(results2, queryResult)
	//			//results = codeData.WhichIDP
	//		}
	//	}
	//}

	return results
}
func (s *SmartContract) RemoveApproval(ctx contractapi.TransactionContextInterface, forWhichSp string, whichIdp string) {
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"Code Store\",\"ForWhichSP\": \"%s\", \"WhichIDP\": \"%s\"}}", forWhichSp, whichIdp)
	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
	defer resultsIterator.Close()
	codeData := new(NewCodeStore)
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		_ = json.Unmarshal(queryResponse.Value, codeData)
	}
	//var result string
	if codeData.ForWhichSP == forWhichSp && codeData.WhichIDP == whichIdp {
		//codeData.Doctype = " "
		//codeData.ForWhichSP = " "
		//codeData.WhichIDP =  " "
		//codeData.SPCheck = " "
		//codeData.IDPCheck = " "
		//codeData.SPCode = " "
		//codeData.IDPCode = " "
		codeDataBytes, _ := json.Marshal(codeData)
		ctx.GetStub().PutState(codeData.Key, codeDataBytes)
		ctx.GetStub().DelState(codeData.Key)
	//	result = codeData.Doctype
	}
	//return result
}
func (s *SmartContract) NewCode(ctx contractapi.TransactionContextInterface, forWhichSp string, whichIdp string, author string, code string) string {
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"Code Store\",\"ForWhichSP\": \"%s\", \"WhichIDP\": \"%s\"}}", forWhichSp, whichIdp)
	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
	defer resultsIterator.Close()
	codeData := new(NewCodeStore)
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		_ = json.Unmarshal(queryResponse.Value, codeData)
	}
	var result = codeData.Key
	if codeData.ForWhichSP == forWhichSp && codeData.WhichIDP == whichIdp {
		if codeData.ForWhichSP == forWhichSp && codeData.WhichIDP == whichIdp {
			if author == "sp" && codeData.SPCode == code {
				codeData.SPCheck = "success"
				codeDataBytes, _ := json.Marshal(codeData)
				ctx.GetStub().PutState(codeData.Key, codeDataBytes)
				result = "sp-success"
				//result = codeData.SPCheck
			} else if author == "idp" && codeData.IDPCode == code {
				codeData.IDPCheck = "success"
				codeDataBytes, _ := json.Marshal(codeData)
				ctx.GetStub().PutState(codeData.Key, codeDataBytes)
				result = "idp-success"

			} else {
				result = "code-failed"
			}
		} else {
			result = "code-failed"
		}
	}
	return  result
}
func (s *SmartContract) NewCodeFetch(ctx contractapi.TransactionContextInterface, forWhichSp string, whichIdp string, author string, code string) string{
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"Code Store\",\"ForWhichSP\": \"%s\", \"WhichIDP\": \"%s\"}}", forWhichSp, whichIdp)
	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
	defer resultsIterator.Close()
	codeData := new(NewCodeStore)
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		_ = json.Unmarshal(queryResponse.Value, codeData)
	}
	var result string
	if codeData.ForWhichSP == forWhichSp && codeData.WhichIDP == whichIdp {
		if author == "sp" && codeData.SPCode == code {
			codeData.SPCheck = "success"
			codeDataBytes, _ := json.Marshal(codeData)
			ctx.GetStub().PutState(codeData.Key, codeDataBytes)
			result = "sp-success"
		} else if author == "idp" && codeData.IDPCode == code {
			codeData.IDPCheck = "success"
			codeDataBytes, _ := json.Marshal(codeData)
			ctx.GetStub().PutState(codeData.Key, codeDataBytes)
			result = "idp-success"
		}
	} else {
		result = "code-failed"
	}
	return result
}












//func (s *SmartContract) NewCodeFetch(ctx contractapi.TransactionContextInterface, forWhichSp string, whichIdp string, author string, code string) string {
//	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"Code Store\",\"ForWhichSP\": \"%s\", \"WhichIDP\": \"%s\"}}", forWhichSp, whichIdp)
//	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
//	defer resultsIterator.Close()
//	codeData := new(NewCodeStore)
//	for resultsIterator.HasNext() {
//		queryResponse, _ := resultsIterator.Next()
//		_ = json.Unmarshal(queryResponse.Value, codeData)
//	}
//	var result string
//	if codeData.ForWhichSP == forWhichSp && codeData.WhichIDP == whichIdp {
//		//codeData.Doctype = " "
//		//codeData.ForWhichSP = " "
//		//codeData.WhichIDP =  " "
//		//codeData.SPCheck = " "
//		//codeData.IDPCheck = " "
//		//codeData.SPCode = " "
//		//codeData.IDPCode = " "
//		//codeDataBytes, _ := json.Marshal(codeData)
//		//ctx.GetStub().PutState(codeData.Key, codeDataBytes)
//			result = codeData.Doctype
//	}
//	//return result
//
//	if codeData.ForWhichSP == forWhichSp && codeData.WhichIDP == whichIdp {
//		if author == "sp" && codeData.SPCode == code {
//			codeData.SPCheck = "success"
//			codeDataBytes, _ := json.Marshal(codeData)
//			ctx.GetStub().PutState(codeData.Key, codeDataBytes)
//			 result = codeData.SPCode
//		} else if author == "idp" && codeData.IDPCode == code {
//			codeData.IDPCheck = "success"
//			codeDataBytes, _ := json.Marshal(codeData)
//			ctx.GetStub().PutState(codeData.Key, codeDataBytes)
//			result = codeData.IDPCode
//		}
//	} else {
//		result = "1234"
//	}
//	return result
//}
//func (s *SmartContract) StoreCode(ctx contractapi.TransactionContextInterface, forWhichSp string, whichIdp string, code string) error {
//	//b := make([]byte, 4)
//	//rand.Read(b)
//	//key := fmt.Sprintf("%x", b)
//	i++
//
//	key := string(i)
//	codeStore := CodeStore{
//		Doctype: "Code Store",
//		ForWhichSP :  forWhichSp,
//		WhichIDP : whichIdp,
//		Code : code,
//		Key : key,
//	}
//	codeDataBytes, _ := json.Marshal(codeStore)
//	return ctx.GetStub().PutState(codeStore.Key,codeDataBytes)
//}

func (s *SmartContract) NewStoreCode(ctx contractapi.TransactionContextInterface, forWhichSp string, whichIdp string, spCode string, idpCode string, spCheck string, idpCheck string, author string) error {
	//b := make([]byte, 4)
	//rand.Read(b)
	//key := fmt.Sprintf("%x", b)
	i++
	key := string(i)
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"Code Store\",\"ForWhichSP\": \"%s\", \"WhichIDP\": \"%s\"}}",forWhichSp,whichIdp)
	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
	defer resultsIterator.Close()
	codeData := new(NewCodeStore)
	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		_ = json.Unmarshal(queryResponse.Value, codeData)
	}
	codeStore := NewCodeStore{
		Doctype: "Code Store",
		ForWhichSP :  forWhichSp,
		WhichIDP : whichIdp,
		SPCode : spCode,
		IDPCode: idpCode,
		SPCheck: spCheck,
		IDPCheck: idpCheck,
		Key : key,
	}

	if codeData.ForWhichSP == forWhichSp && codeData.WhichIDP == whichIdp {
		if author == "sp" {
			codeData.SPCode = spCode
		} else if author == "idp"{
			codeData.IDPCode = idpCode
		}
		codeDataBytes, _ := json.Marshal(codeData)
		return ctx.GetStub().PutState(codeData.Key,codeDataBytes)
	} else{
		codeDataBytes, _ := json.Marshal(codeStore)
		return ctx.GetStub().PutState(codeStore.Key,codeDataBytes)
	}
}

func (s *SmartContract) StoreTalList(ctx contractapi.TransactionContextInterface, entityId string, tal string) error {
	//b := make([]byte, 4)
	//rand.Read(b)
	//key := fmt.Sprintf("%x", b)
	m++
	key := string(m)
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"TAL List\",\"EntityId\": \"%s\"}}", entityId)
	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
	//if err != nil {
	//	return nil, nil
	//}
	defer resultsIterator.Close()
	codeData := new(TalList)
	//data := MetaDataStore{}
	var results []QueryResultTalList

	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		_ = json.Unmarshal(queryResponse.Value, codeData)
		queryResult := QueryResultTalList{Key: queryResponse.Key, Record: codeData}
		results = append(results, queryResult)
	}
	talList := TalList{
		Doctype: "TAL List",
		EntityId :  entityId,
		TList: [] List{
			{Tal : tal },
		},
		Key: key,
	}
	talListBytes, _ := json.Marshal(talList)

	if codeData.EntityId == entityId{
		l := List{}
		l.Tal =tal
		codeData.TList= append(codeData.TList, l)
		//customer.Contact = append(customer.Contact, *c)
		//customer.Contact = append(customer.Contact, *b)
		codeDataBytes, _ := json.Marshal(codeData)
    return ctx.GetStub().PutState(codeData.Key,codeDataBytes)
	} else{
		return ctx.GetStub().PutState(talList.Key,talListBytes)
	}

}

func (s *SmartContract) TalListDelete(ctx contractapi.TransactionContextInterface, entityId string, tal string) {
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"TAL List\",\"EntityId\": \"%s\"}}", entityId)
	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
	//if err != nil {
	//	return nil, nil
	//}
	defer resultsIterator.Close()
	codeData := new(TalList)
	//data := MetaDataStore{}
	var results []QueryResultTalList

	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		_ = json.Unmarshal(queryResponse.Value, codeData)
		queryResult := QueryResultTalList{Key: queryResponse.Key, Record: codeData}
		results = append(results, queryResult)
	}
	var pos int
	for _,service := range codeData.TList {
		if service.Tal == tal {
			codeData.TList= append(codeData.TList[:pos],codeData.TList[pos+1:]...)
			if pos >0{
				pos = pos -1
			}
			continue
		}
		pos ++
	}
	codeDataBytes, _ := json.Marshal(codeData)
	ctx.GetStub().PutState(codeData.Key,codeDataBytes)

	//return results
}

func (s *SmartContract) TalListFetch(ctx contractapi.TransactionContextInterface, entityId string) []QueryResultTalList{
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"TAL List\",\"EntityId\": \"%s\"}}", entityId)
	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
	//if err != nil {
	//	return nil, nil
	//}
	defer resultsIterator.Close()
	codeData := new(TalList)
	//data := MetaDataStore{}
	var results []QueryResultTalList

	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		_ = json.Unmarshal(queryResponse.Value, codeData)
		queryResult := QueryResultTalList{Key: queryResponse.Key, Record: codeData}
		results = append(results, queryResult)
	}

	return results
}
func (s *SmartContract) TalListCheck(ctx contractapi.TransactionContextInterface, entityId string) string{
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"TAL List\",\"EntityId\": \"%s\"}}", entityId)
	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
	//if err != nil {
	//	return nil, nil
	//}
	defer resultsIterator.Close()
	codeData := new(TalList)
	//data := MetaDataStore{}
	var results []QueryResultTalList

	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		_ = json.Unmarshal(queryResponse.Value, codeData)
		queryResult := QueryResultTalList{Key: queryResponse.Key, Record: codeData}
		results = append(results, queryResult)
	}

	return codeData.EntityId

}
func (s *SmartContract) ShowTalList(ctx contractapi.TransactionContextInterface) []QueryResultTalList{
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"TAL List\"}}")
	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
	//if err != nil {
	//	return nil, nil
	//}
	defer resultsIterator.Close()
	codeData := new(TalList)
	//data := MetaDataStore{}
	var results []QueryResultTalList

	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		_ = json.Unmarshal(queryResponse.Value, codeData)
		queryResult := QueryResultTalList{Key: queryResponse.Key, Record: codeData}
		results = append(results, queryResult)
	}

	return results

}

func (s *SmartContract) TalListReturn(ctx contractapi.TransactionContextInterface, entityId string) []List{
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"TAL List\",\"EntityId\": \"%s\"}}", entityId)
	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
	//if err != nil {
	//	return nil, nil
	//}
	defer resultsIterator.Close()
	codeData := new(TalList)
	//data := MetaDataStore{}
	var results []QueryResultTalList

	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		_ = json.Unmarshal(queryResponse.Value, codeData)
		queryResult := QueryResultTalList{Key: queryResponse.Key, Record: codeData}
		results = append(results, queryResult)
	}

	return codeData.TList

	///return results
}
//deleteData
//func (s *SmartContract) DeleteCode(ctx contractapi.TransactionContextInterface,forwhichSp string, whichIdp string) error {
//
//	return ctx.GetStub().DelState(id)
//}
//Query Code
//func (s *SmartContract) QueryCode(ctx contractapi.TransactionContextInterface, idp string) (*CodeStore, error) {
//	codeDataBytes, err := ctx.GetStub().GetState(idp)
//
//	if err != nil {
//		return nil, fmt.Errorf("failed to read from world state. %s", err.Error())
//	}
//
//	if codeDataBytes == nil {
//		return nil, fmt.Errorf("%s does not exist", idp)
//	}
//
//	codeData := new(CodeStore)
//	_ = json.Unmarshal(codeDataBytes, codeData)
//
//	return codeData, nil
//}
func (s *SmartContract) QueryForSpecificUser(ctx contractapi.TransactionContextInterface, user string) []QueryResultMetaData{
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"MetaData Store\",\"User\": \"%s\"}}",user)
	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
	//if err != nil {
	//	return nil, nil
	//}
	defer resultsIterator.Close()
	codeData := new(MetaDataStore)
	//data := MetaDataStore{}
	var results []QueryResultMetaData

	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		_ = json.Unmarshal(queryResponse.Value, codeData)
		queryResult := QueryResultMetaData{Key: queryResponse.Key, Record: codeData}
		results = append(results, queryResult)
	}
	return results
}
//func (s *SmartContract) QueryAllCars(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
//	startKey := ""
//	endKey := ""
//
//	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
//
//	if err != nil {
//		return nil, err
//	}
//	defer resultsIterator.Close()
//
//	var results []QueryResult
//
//	for resultsIterator.HasNext() {
//		queryResponse, err := resultsIterator.Next()
//
//		if err != nil {
//			return nil, err
//		}
//
//		car := new(Car)
//		_ = json.Unmarshal(queryResponse.Value, car)
//
//		queryResult := QueryResult{Key: queryResponse.Key, Record: car}
//		results = append(results, queryResult)
//	}
//
//	return results, nil
//}
//query all metadata
// Query metadata with a specific usernam
func (s *SmartContract) UserFetch(ctx contractapi.TransactionContextInterface, user string) string {
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"MetaData Store\",\"User\": \"%s\"}}",user)
	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
	//if err != nil {
	//	return nil, nil
	//}
	defer resultsIterator.Close()
	codeData := new(MetaDataStore)
	//data := MetaDataStore{}
	var results []QueryResultMetaData

	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()

		//if err != nil {
		//	return nil, nil
		//}
		//codeData := new(MetaDataStore)
		_ = json.Unmarshal(queryResponse.Value, codeData)
		queryResult := QueryResultMetaData{Key: queryResponse.Key, Record: codeData}
		results = append(results, queryResult)
	}
	//data.Doctype=codeData.Doctype
	//data.User = codeData.User
	//data.Metadata =codeData.Metadata
	if codeData.User ==user{
		return codeData.User
	} else {
		return "Not found"
	}
}

func (s *SmartContract) MetaDataFetch(ctx contractapi.TransactionContextInterface, user string) string{
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"MetaData Store\",\"User\": \"%s\"}}",user)
	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
	//if err != nil {
	//	return nil, nil
	//}
	defer resultsIterator.Close()
	codeData := new(MetaDataStore)
	//data := MetaDataStore{}
	var results []QueryResultMetaData

	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()

		//if err != nil {
		//	return nil, nil
		//}
		//codeData := new(MetaDataStore)
		_ = json.Unmarshal(queryResponse.Value, codeData)
		queryResult := QueryResultMetaData{Key: queryResponse.Key, Record: codeData}
		results = append(results, queryResult)
	}
	return codeData.Metadata
}


func (s *SmartContract) DeleteCodeSp(ctx contractapi.TransactionContextInterface, forWhichSp string, WhichIdp string) {
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"Code Store\",\"ForWhichSP\": \"%s\", \"WhichIDP\": \"%s\"}}",forWhichSp,WhichIdp)
	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
	//if err != nil {
	//	return nil, nil
	//}
	defer resultsIterator.Close()
	codeData := new(CodeStore)
	//data := MetaDataStore{}
	var results []QueryResultCode

	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		_ = json.Unmarshal(queryResponse.Value, codeData)
		queryResult := QueryResultCode{Key: queryResponse.Key, Record: codeData}
		results = append(results, queryResult)
	}
	codeData.Doctype = " "
	codeData.WhichIDP = " "
	codeData.ForWhichSP = " "
	codeData.Code = " "
	codeDataBytes, _ := json.Marshal(codeData)
	ctx.GetStub().PutState(codeData.Key,codeDataBytes)

}

//func (s *SmartContract) CodeFetch(ctx contractapi.TransactionContextInterface, forWhichSp string, WhichIdp string) string {
//	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"Code Store\",\"ForWhichSP\": \"%s\", \"WhichIDP\": \"%s\"}}",forWhichSp,WhichIdp)
//	resultsIterator, _ := ctx.GetStub().GetQueryResult(queryString)
//	defer resultsIterator.Close()
//	codeData := new(CodeStore)
//	for resultsIterator.HasNext() {
//		queryResponse, _ := resultsIterator.Next()
//		_ = json.Unmarshal(queryResponse.Value, codeData)
//	}
//	return codeData.Code
//}

func (s *SmartContract) AllMetaData (ctx contractapi.TransactionContextInterface) ([]QueryResultMetaData, error){
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"MetaData Store\"}}")
	fmt.Println(queryString)
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	var results []QueryResultMetaData

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		metaData := new(MetaDataStore)
		_ = json.Unmarshal(queryResponse.Value, metaData)

		queryResult := QueryResultMetaData{Key: queryResponse.Key, Record: metaData}
		results = append(results, queryResult)
	}

	return results, nil
}
func (s *SmartContract) AllCodeData (ctx contractapi.TransactionContextInterface) ([]QueryResultNewCode, error){
	queryString := fmt.Sprintf("{\"selector\": {\"Doctype\": \"Code Store\"}}")
	fmt.Println(queryString)
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	var results []QueryResultNewCode

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		metaData := new(NewCodeStore)
		_ = json.Unmarshal(queryResponse.Value, metaData)

		queryResult := QueryResultNewCode{Key: queryResponse.Key, Record: metaData}
		results = append(results, queryResult)
	}

	return results, nil
}

// ChangeCarOwner updates the owner field of car with given id in world state
//func (s *SmartContract) ChangeCarOwner(ctx contractapi.TransactionContextInterface, carNumber string, newOwner string) error {
//	car, err := s.QueryCar(ctx, carNumber)
//
//	if err != nil {
//		return err
//	}
//
//	car.Owner = newOwner
//
//	carAsBytes, _ := json.Marshal(car)
//
//	return ctx.GetStub().PutState(carNumber, carAsBytes)
//}
///changing metadata
//func (s *SmartContract) ChangeMetaData(ctx contractapi.TransactionContextInterface, user string, newMetaData string) error {
//	metaData, err := s.QueryMetaData(ctx, user)
//
//	if err != nil {
//		return err
//	}
//
//metaData.Metadata =newMetaData
//
//metaDataBytes, _ := json.Marshal(metaData)
//
//	return ctx.GetStub().PutState(metaData.User, metaDataBytes)
//}
func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
}
