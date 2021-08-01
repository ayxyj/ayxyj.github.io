package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"strconv"
)

type SimpleChainCode struct {

}

func (t *SimpleChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response  {
	//在改方法中实现链码初始化或者升级时候的处理逻辑
	//编写时候可以灵活使用stub中的api
	//eturn pb.Response{}
	funcName, args := stub.GetFunctionAndParameters()
	//{"Args":["init","a","100","b","200"]}
	//parameters = init
	//strings[]  = a 100 b 200
	if funcName!="init"{
		return shim.Error("Error argument name:init")
	}
	if len(args) != 4{
		return shim.Error("Incorrent number of arguments ,Exception 4")
	}

	//读取参数
	A := args[0]
	Aval ,err := strconv.Atoi(args[1])
	if err != nil{
		return shim.Error("Exception integer value for asset holding!")
	}
	B := args[2]
	Bval , err := strconv.Atoi(args[3])
	if err != nil{
		return shim.Error("Exception integer value for asset holding!")
	}

	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err!= nil{
		return shim.Error(err.Error())
	}
	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err!= nil{
		return shim.Error(err.Error())
	}
	fmt.Println("=========== success init ===============")
	return shim.Success(nil)
}

func (t *SimpleChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response  {
	//在该方法中实现链码运行中被调用或者查询时候的处理逻辑
	//编写时候可以灵活使用stub中的api
	//return pb.Response{}
	//获取所有参数
	fmt.Println("=========== in invoke ==============")
	funcName, args := stub.GetFunctionAndParameters()
	if funcName =="invoke"{
		return t.Input(stub,args)
	}else if funcName=="delete"{
		return t.delete(stub,args)
	}else if funcName=="query"{
		return t.query(stub,args)
	}
	return shim.Success([]byte("success invoke"))
}
func (t *SimpleChainCode) Input(stub shim.ChaincodeStubInterface, strings []string) pb.Response {
	return shim.Success(nil)
}
func (t *SimpleChainCode) delete(stub shim.ChaincodeStubInterface, strings []string) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChainCode) query(stub shim.ChaincodeStubInterface, strings []string) pb.Response {
	return shim.Success(nil)
}
func main()  {
	fmt.Println("hello world !")
	err := shim.Start(new(SimpleChainCode))
	if err != nil{
		fmt.Printf("error starting simple chaincode! %s",err)
	}
}
