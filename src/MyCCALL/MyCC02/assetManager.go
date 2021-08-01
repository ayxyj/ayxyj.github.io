package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)
type SimpleChainCode struct {

}
type marble struct {
	ObjectType  string		`json:"docType"`
	Name 		string		`json:"name"`
	Color		string 		`json:"color"`
	Size 		string 		`json:"size"`
	Owner		string 		`json:"owner"`
}
func (t *SimpleChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	//不做具体处理
	return shim.Success(nil)
}
func (t *SimpleChainCode)Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	fmt.Println("====================function Name ：====================",funcName)
	if funcName == "initMarble"{
		//创建一个大理石信息并写入账本
		//{"initMarble","marble1","blue","35","tom"}
		return t.initMarble(stub , args)
	}else if funcName == "transferMarble"{
		//更改一个大理石拥有者
		//{"tranferMarble","marble2","jack"}
		return t.transferMarble(stub , args)
	}else if funcName == "transferMarbleBaseOnColor"{
		//返回所有名称在制定字典序范围内的大理石信息
		//{"transferMarbleBaseOnColor","marble1","marble3"}
		return t.transferMarbleBaseOnColor(stub , args)
	}else if funcName == "delete"{
		//删除一个大理石信息
		//{"delete","marble1"}
		return t.delete(stub , args)
	}else if funcName == "readMarble"{
		//从账本中读取一个大理石信息
		//{"readMarble","marble1"}
		return t.readMarble(stub , args)
	}else if funcName == "queryMarblesByOwner"{
		//返回指定拥有者拥有的所有的大理石信息
		//{"queryMarblesByOwner","tom"}
		return t.queryMarblesByOwner(stub , args)
	}else if funcName == "queryMarbles"{
		//富查询(rich query)大理石信息
		//{"queryMarbles","{\"selector\":{\"wner\":\"tom\"}}"}
		return t.queryMarbles(stub , args)
	}else if funcName == "getHistoryForMarble"{
		//返回一个大理石信息
		//{"getHistoryForMarble","marble1"}
		return t.getHistoryForMarble(stub , args)
	}else if funcName == "getMarblesByRangeWithPagination"{
		//分页返回富查询（rich query）大理石信息
		//{"getMarblesByRangeWithPagination","marble1","marble3","10","1"}
		return t.getMarblesByRangeWithPagination(stub , args)
	}else if funcName == "getMarblesByRange"{
		//分页返回所有名称在制定字典序范围内的大理石信息
		//{"getMarblesByRange","maeble1","marble3"}
		return t.getMarblesByRange(stub , args)
	}else if funcName == "queryMarblesWithPagination"{
		//分页返回大理石信息（rich query）
		//{"queryMarblesWithPagination","{\"selector\":{\"wner\":\"tom\"}}","10","1"}
		return t.queryMarblesWithPagination(stub , args)
	}
	return shim.Success(nil)
}

func (t *SimpleChainCode) initMarble(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	marbleName := args[0]  //大理石名称
	marbleAsBytes, err := stub.GetState(marbleName)
	//判断大理石是否存在
	if err != nil{
		return shim.Error("init faile to get marble:"+err.Error())
	}else if  marbleAsBytes != nil {
		fmt.Println("this marble already exists :" + marbleName)
		return shim.Error("this marble already exists：" + marbleName)
	}
	//创建marble,并进行序列化为JSON对象
	objectType := "marble"
	color := args[1]
	size  := args[2]
	owner := args[3]
	marble := &marble{ObjectType: objectType,Color: color,Size: size,Owner: owner}
	marshal, err := json.Marshal(marble)
	if err != nil{
		return shim.Error(err.Error())
	}
	//写入账本
	err = stub.PutState(marbleName, marshal)
	if err !=nil{
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *SimpleChainCode) transferMarble(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChainCode) transferMarbleBaseOnColor(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChainCode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChainCode) readMarble(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChainCode) queryMarblesByOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChainCode) queryMarbles(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChainCode) getHistoryForMarble(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChainCode) getMarblesByRangeWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChainCode) getMarblesByRange(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChainCode) queryMarblesWithPagination(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}

func main()  {
	//主函数需要调用start方法
	err := shim.Start(new(SimpleChainCode))
	if err != nil  {
		fmt.Printf("Err Starting Simple chaincode! %s",err)
	}
}
