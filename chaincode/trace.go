package main

import (
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type TraceChaincode struct {
	
}

func (t *TraceChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)
}

func (t *TraceChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	
	fun, args := stub.GetFunctionAndParameters()

	if fun == "addTrace" {
		return t.addTrace(stub, args)
	} else if fun == "queryTraceByCode" {
		return t.queryTraceByCode(stub, args)
	}

	return shim.Error("Parameter error")
}

func (t *TraceChaincode) addTrace(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2{
		return shim.Error("Parameter number is not 2")
	}

	var traceInfo Trace
	err := json.Unmarshal([]byte(args[0]), &traceInfo)
	if err != nil {
		return shim.Error("Error in deserialization")
	}

	// 查重: 身份证号码必须唯一
	_, exist := getTrace(stub, traceInfo.UniqueCode)
	if exist {
		return shim.Error("Product already exist")
	}

	_, bl := putTrace(stub, traceInfo)
	if !bl {
		return shim.Error("Save info error")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("Trace info added successfully"))
}

func getTrace(stub shim.ChaincodeStubInterface, uniqueCode string) (Trace, bool)  {
	var traceInfo Trace

	b, err := stub.GetState(uniqueCode)
	if err != nil {
		return traceInfo, false
	}

	if b == nil {
		return traceInfo, false
	}

	err = json.Unmarshal(b, &traceInfo)
	if err != nil {
		return traceInfo, false
	}

	return traceInfo, true
}

func putTrace(stub shim.ChaincodeStubInterface, traceInfo Trace) ([]byte, bool) {

	b, err := json.Marshal(traceInfo)
	if err != nil {
		return nil, false
	}

	err = stub.PutState(traceInfo.UniqueCode, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

func (t *TraceChaincode) queryTraceByCode(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Parameter number is not 1")
	}

	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Query failure")
	}

	if b == nil {
		return shim.Error("No record found")
	}

	return shim.Success(b)
}

func main(){
	err := shim.Start(new(TraceChaincode))
	if err != nil{
		fmt.Printf("Start TraceChaincode error: %s", err)
	}
}