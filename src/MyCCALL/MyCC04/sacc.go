

package main

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)
/**
	简单资产管理代码 sacc
	实现资产存取
 */

// SimpleAsset 实现了一个简单的链代码来管理资产
type SimpleAsset struct {
}

// Init在chaincode实例化期间被调用来初始化任何数据。注意，chaincode upgrade也调用这个函数来重置或迁移数据。
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	//从事务提议中获取参数
	//funcName, args := stub.GetFunctionAndParameters()
	args := stub.GetStringArgs()
	if len(args) != 2 {
		//参数个数不对
		return shim.Error("Incorrect arguments. Expecting a key and a value")
	}

	//通过调用stub.PutState()来设置任何变量或资产

	// 我们将键和值存储在总账上
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		//记账失败 ，创建失败
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
	}
	return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. Each transaction is either a 'get' or a 'set' on the asset created by Init function. The Set method may create a new asset by specifying a new key-value pair.
//调用是对链代码上的每个事务调用的。每个事务都是Init函数创建的资产的“获取”或“设置”。Set方法可以通过指定一个新的键值对来创建一个新的资产。
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// 从事务提议中提取函数和参数
	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error
	//我们将使set与get这两个函数名正式生效，并调用这些chaincode应用函数，经由shim.Success或shim.Error函数返回一个合理的响应。这两个shim成员函数可以将响应序列化为gRPC protobuf消息
	if fn == "set" {
		//操作类型 set
		result, err = set(stub, args)
	} else { // assume 'get' even if fn is nil
		//操作类型 get
		result, err = get(stub, args)
	}
	if err != nil {
		//操作失败
		return shim.Error(err.Error())
	}

	// Return the result as success payload
	//返回结果作为成功有效负载
	return shim.Success([]byte(result))
}

// Set stores the asset (both key and value) on the ledger. If the key exists, it will override the value with the new one
//Set将资产(键和值)存储在分类账上。如果该键存在，它将用新的值覆盖该值
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return args[1], nil
}

// Get returns the value of the specified asset key
// Get返回指定资源键的值
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	return string(value), nil
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
