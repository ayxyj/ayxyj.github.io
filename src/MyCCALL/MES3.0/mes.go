package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
)

//合约信息
type SmartContract struct {
	contractapi.Contract
}

//设置节点NodeA ,NodeB ,NodeC
//资源 A,B,C
//初始化： A:0 , B:10 , C:10
////1A = 2B + 3C 在TransferNodeResource中去判断，数量达到时生产A
////节点信息
type Node struct {
	NodeID   string `json:"nodeID"`
	ResTypeA string `json:"resTypeA"`
	ResTypeB string `json:"resTypeB"`
	ResTypeC string `json:"resTypeC"`
	ResNumA  int    `json:"resNumA"`
	ResNumB  int    `json:"resNumB"`
	ResNumC  int    `json:"resNumC"`
}

//需要传入资源参数的函数
//生产C
func (s *SmartContract) ProductResC(ctx contractapi.TransactionContextInterface, ownNodeID string, num int) error {
	//'{"Args":["ProductResC","NodeC","5"]}'
	node, err := s.ReadNode(ctx, ownNodeID)
	if err != nil {
		return err
	}
	if num >= 0 {
		node.ResNumC += num
	} else {
		return fmt.Errorf(" num is error :%d", num)
	}
	//C生产成功，写入账本
	nodeJsonByBytes, err := json.Marshal(node)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(ownNodeID, nodeJsonByBytes)
}

//生产B
func (s *SmartContract) ProductResB(ctx contractapi.TransactionContextInterface, ownNodeID string, num int) error {
	//'{"Args":["ProductResB","NodeB","5"]}'
	node, err := s.ReadNode(ctx, ownNodeID)
	if err != nil {
		return err
	}
	if num >= 0 {
		node.ResNumB += num
	} else {
		return fmt.Errorf(" num is error :%d", num)
	}
	//B生产成功，写入账本
	nodeJsonByBytes, err := json.Marshal(node)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(ownNodeID, nodeJsonByBytes)
}

//生产A
func (s *SmartContract) ProductResA(ctx contractapi.TransactionContextInterface, ownNodeID string) error {
	//'{"Args":["ProductResA","NodeA"]}'
	node, err := s.ReadNode(ctx, ownNodeID)
	if err != nil {
		return err
	}
	var flag bool = false
	//判断B和C的数量是否足够，不够找NodeB和NodeC请求
	if node.ResNumB >= 2 && node.ResNumC >= 3 {
		node.ResNumA++
		node.ResNumB -= 2
		node.ResNumC -= 3
		flag = true
	}else{
		//把请求B和C的任务交给上层网络，在边缘节点进行调用TransferNodeResource
		return fmt.Errorf("The production of A failed because the quantity of B %d and C %d was insufficient",node.ResNumB,node.ResNumC)
	}
	/**
	if !flag && node.ResNumB < 2 {
		//B的数量不够
		transferNum := 2 - node.ResNumB
		err := s.TransferNodeResource(ctx, "NodeB", ownNodeID, "B", transferNum)
		fmt.Println("transfer B!")
		if err != nil {
			return err
		}
	}
	if !flag && node.ResNumC < 3 {
		//C的数量不够
		transferNum := 3 - node.ResNumC
		err := s.TransferNodeResource(ctx, "NodeC", ownNodeID, "C", transferNum)
		fmt.Println("transfer C!")
		if err != nil {
			return err
		}
	}*/
	if flag {
		//A生产成功，写入账本
		nodeJsonByBytes, err := json.Marshal(node)
		if err != nil {
			return err
		}
		return ctx.GetStub().PutState(ownNodeID, nodeJsonByBytes)
	}
	return fmt.Errorf("")
}

//TransferAsset在世界状态中使用给定id更新资产的所有者字段。
func (s *SmartContract) TransferNodeResource(ctx contractapi.TransactionContextInterface, ownNodeID string, transferNodeID string, resType string, num int) error {
	//'{"Args":["TransferNodeResource","NodeB","NodeA","B",2]}'
	nodeOwn, err := s.ReadNode(ctx, ownNodeID)
	transferNode, err := s.ReadNode(ctx, transferNodeID)
	if err != nil {
		return err
	}
	//获取拥有资源节点的信息并更新
	//获取转移资源节点的信息并更新
	//根据资源类型所在的索引位置去更新数量的切片resNum[index] -= num  ; 判断数量是否足够
	//资源数量不够，请求生产对应资源
	var flag bool = false
	if resType == "A" && nodeOwn.ResNumA >= num {
		nodeOwn.ResNumA -= num
		transferNode.ResNumA += num
		flag = true
	}
	if resType == "B" && nodeOwn.ResNumB >= num {
		nodeOwn.ResNumB -= num
		transferNode.ResNumB += num
		flag = true
	}
	if resType == "C" && nodeOwn.ResNumC >= num {
		nodeOwn.ResNumC -= num
		transferNode.ResNumC += num
		flag = true
	}
	//资源信息修改完毕，需要更新账本中的信息
	//更新两个节点信息即可
	if flag {
		nodeOwnByBytes, err := json.Marshal(nodeOwn)
		transferNodeByBytes, err1 := json.Marshal(transferNode)
		if err != nil || err1 != nil {
			return fmt.Errorf("json marshal faied !")
		}
		err = ctx.GetStub().PutState(ownNodeID, nodeOwnByBytes)
		err = ctx.GetStub().PutState(transferNodeID, transferNodeByBytes)
		return err
	} else {
		return fmt.Errorf("resType %s num does not enought!", resType)
	}
	return nil
}

//UpdateNode使用提供的参数更新世界状态中的现有节点的资源信息。
func (s *SmartContract) UpdateNode(ctx contractapi.TransactionContextInterface, nodeID string, typeA string, typeB string, typeC string, numA int, numB int, numC int) error {
	//'{"function":"CreateNode","Args":["NodeA","A","B","C","0","0","0"]}'
	//更新某个节点的资源需求数量和资源拥有数量

	//判断节点信息是否存在账本
	exists, err := s.NodeExists(ctx, nodeID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the node %s does not extis!", nodeID)
	}
	// 用新节点覆盖原有节点
	node := Node{
		NodeID:   nodeID,
		ResTypeA: typeA,
		ResTypeB: typeB,
		ResTypeC: typeC,
		ResNumA:  numA,
		ResNumB:  numB,
		ResNumC:  numC,
	}
	nodeByBytes, err := json.Marshal(node)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(nodeID, nodeByBytes)
}

//创建新的节点信息
func (s *SmartContract) CreateNode(ctx contractapi.TransactionContextInterface, nodeID string, typeA string, typeB string, typeC string, numA int, numB int, numC int) error {
	//'{"function":"CreateNode","Args":["NodeA","A","B","C",0,0,0]}'
	//创建节点NodeD ，资源需要 A:1  C:2   , 初始拥有0
	exists, err := s.NodeExists(ctx, nodeID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("Node already exists ： %s", nodeID)
	}
	//新建节点信息
	node := Node{
		NodeID:   nodeID,
		ResTypeA: typeA,
		ResTypeB: typeB,
		ResTypeC: typeC,
		ResNumA:  numA,
		ResNumB:  numB,
		ResNumC:  numC,
	}
	//格式化节点信息为json
	nodeJsonByBytes, err := json.Marshal(node)
	if err != nil {
		return err
	}
	//写入账本
	return ctx.GetStub().PutState(nodeID, nodeJsonByBytes)
}

//不需要传入其他信息的函数

//ReadNode返回以给定nodeID存储在世界状态中的信息。
func (s *SmartContract) ReadNode(ctx contractapi.TransactionContextInterface, nodeID string) (*Node, error) {
	//根据节点ID获取节点
	nodeByBytes, err := ctx.GetStub().GetState(nodeID)
	if err != nil {
		return nil, err
	}
	if nodeByBytes == nil {
		return nil, fmt.Errorf("the Node %s does not exits !", nodeID)
	}
	//将json字节转换为node
	node := new(Node)
	err = json.Unmarshal(nodeByBytes, node)
	if err != nil {
		return nil, err
	}
	//返回节点信息
	return node, nil
}

//判断该节点信息是否存在账本当中
func (s *SmartContract) NodeExists(ctx contractapi.TransactionContextInterface, nodeID string) (bool, error) {
	nodeJsonBytes, err := ctx.GetStub().GetState(nodeID)
	if err != nil {
		return false, err
	}
	return nodeJsonBytes != nil, nil
}

//DeleteNode从世界状态中删除给定的节点。
func (s *SmartContract) DeleteNode(ctx contractapi.TransactionContextInterface, nodeID string) error {
	//根据节点ID删除节点信息
	exists, err := s.NodeExists(ctx, nodeID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the node %s does not exist", nodeID)
	}

	return ctx.GetStub().DelState(nodeID)
}

// GetAllNode返回在世界状态下找到的所有资产
func (s *SmartContract) GetAllNode(ctx contractapi.TransactionContextInterface) ([]*Node, error) {
	// startKey和endKey为空字符串的范围查询执行
	// 对chaincode命名空间中所有资产的开放式查询。
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var nodes []*Node
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var node Node
		err = json.Unmarshal(queryResponse.Value, &node)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, &node)
	}

	return nodes, nil
}

//初始化信息
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	//节点初始化
	//NodeA资源信息、需求信息、其他节点信息,初始化有各有10个资源；需求信息根据生产动态调整即可
	nodes := []Node{
		{NodeID: "NodeA", ResTypeA: "A", ResTypeB: "B", ResTypeC: "C", ResNumA: 0, ResNumB: 0, ResNumC: 0},
		{NodeID: "NodeB", ResTypeA: "A", ResTypeB: "B", ResTypeC: "C", ResNumA: 0, ResNumB: 10, ResNumC: 0},
		{NodeID: "NodeC", ResTypeA: "A", ResTypeB: "B", ResTypeC: "C", ResNumA: 0, ResNumB: 0, ResNumC: 10},
	}
	for _, node := range nodes {
		nodeJsonBytes, err := json.Marshal(node)
		if err != nil {
			return err
		}
		//将节点信息写入账本
		err = ctx.GetStub().PutState(node.NodeID, nodeJsonBytes)
		if err != nil {
			return fmt.Errorf("failed put to world state : %v", err)
		}
	}
	return nil
}

func main() {
	fmt.Println("--------------SmartContract--------------")
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panic("Error creating asset-transfer-basic chaincode :%v")
	}
	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
	}
}
