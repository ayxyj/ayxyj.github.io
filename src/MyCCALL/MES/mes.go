package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//合约信息
type SmartContract struct {
	contractapi.Contract
}
//初始化： A:0 , B:10 , C:10
//1A = 2B + 3C
//节点信息
type Node struct {
	NodeID			string				`json:"nodeID"`
	//所需资源 例如：A初始：2B  3C
	ResNeed			Resource			`json:"resNeed"`
	//已有资源 例如：A初始：0A 0B 0C
	ResOwn			Resource			`json:"resOwn"`
}
//节点资源信息
type Resource struct {
	ResType			[]string				`json:"resType"`
	ResNum			[]int					`json:"resNum"`
}

// GetAllNode返回在世界状态下找到的所有资产
func (s *SmartContract) GetAllNode(ctx contractapi.TransactionContextInterface) ([]*Node ,error)  {
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

//TransferAsset在世界状态中使用给定id更新资产的所有者字段。
func (s *SmartContract) TransferNodeResource(ctx contractapi.TransactionContextInterface , ownNodeID string ,  transferNodeID string , resType string , num  int )  error {
	//'{"Args":["TransferNodeResource","NodeB","NodeA",2]}'
	nodeOwn, err := s.ReadNode(ctx, ownNodeID)
	transferNode, err := s.ReadNode(ctx, transferNodeID)
	if err != nil {
		return err
	}
	//获取拥有资源节点的信息并更新
	//获取转移资源节点的信息并更新
	//根据资源类型所在的索引位置去更新数量的切片resNum[index] -= num  ; 判断数量是否足够
	resTypeArr := nodeOwn.ResOwn.ResType
	var indexOwn int 				//索引位置
	var flagOwn  bool = false  	    //是否存在
	//判断是否存在和索引位置
	for indexOwn=0 ; indexOwn <= len(resTypeArr) ; indexOwn++ {
		if resTypeArr[indexOwn] == resType{
			flagOwn = true
			break
		}
	}
	//判断拥有者资源类型索引位置，没有则添加进去
	var indexTransfer int 				//索引位置
	var flagTransfer  bool = false  	//是否存在
	resTypeArrTransfer := transferNode.ResOwn.ResType
	for indexTransfer=0 ; indexTransfer <= len(resTypeArrTransfer) ; indexTransfer++ {
		if resTypeArrTransfer[indexTransfer] == resType{
			flagTransfer = true
			break
		}
	}
	//修改对应节点的资源数量
	if flagOwn  {
		//资源拥有者数量减少
		resNumArr := nodeOwn.ResOwn.ResNum
		resNumArr[indexOwn] -= num
		//资源获取者数量增加
		if flagTransfer{
			//资源存在，直接修改数量
			resNumArr := transferNode.ResOwn.ResNum
			resNumArr[indexTransfer] += num
		}else{
			//资源不存在，在对应切片中的索引位置添加资源和数量
			resTypeArr := transferNode.ResOwn.ResType
			resTypeArr = append( resTypeArr, resType )//末尾添加
			resNumArr := transferNode.ResOwn.ResNum
			resNumArr  = append( resNumArr , num )//末尾添加
		}
	}
	//资源信息修改完毕，需要更新账本中的信息
	//更新两个节点信息即可
	nodeOwnByBytes, err := json.Marshal(nodeOwn)
	transferNodeByBytes, err1 := json.Marshal(transferNode)
	if err != nil ||  err1 != nil{
		return fmt.Errorf("json marshal faied !")
	}
	err = ctx.GetStub().PutState(ownNodeID, nodeOwnByBytes)
	err = ctx.GetStub().PutState(transferNodeID, transferNodeByBytes)
	return err
}
//DeleteNode从世界状态中删除给定的节点。
func ( s *SmartContract ) DeleteNode( ctx contractapi.TransactionContextInterface , nodeID string )  error {
	//根据节点ID删除节点信息
	exists, err := s.NodeExists(ctx, nodeID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the node %s does not exist", nodeID)
	}

	return ctx.GetStub().DelState( nodeID )
}

//UpdateNode使用提供的参数更新世界状态中的现有节点的资源信息。
func ( s *SmartContract ) UpdateNode( ctx contractapi.TransactionContextInterface , nodeID string , resNeedType  []string , resNeedNum []int , resOwnType []string , resOwnNum []int  ) error {
	//'{"Args":["DeleteNode","NodeA", ["A" , "B" , "C"],[0,0,0],["A" , "B" , "C"],[0,0,0] ]}'
	//更新某个节点的资源需求数量和资源拥有数量

	//判断节点信息是否存在账本
	exists, err := s.NodeExists( ctx , nodeID )
	if err != nil {
		return err
	}
	if !exists  {
		return fmt.Errorf("the node %s does not extis!",nodeID)
	}
	// 用新节点覆盖原有节点
	node := Node{
		NodeID: nodeID,
		ResNeed: Resource{ResType: resNeedType,ResNum: resNeedNum},
		ResOwn:  Resource{ResType: resOwnType,ResNum: resOwnNum},
	}
	nodeByBytes, err := json.Marshal(node)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState( nodeID  , nodeByBytes )
}

//ReadNode返回以给定nodeID存储在世界状态中的信息。
func ( s *SmartContract ) ReadNode( ctx contractapi.TransactionContextInterface , nodeID string ) ( *Node , error )  {
	//根据节点ID获取节点
	nodeByBytes, err := ctx.GetStub().GetState(nodeID)
	if err != nil {
		return nil , err
	}
	if nodeByBytes == nil{
		return  nil , fmt.Errorf("the Node %s does not exits !" , nodeID)
	}
	//将json字节转换为node
	node := new(Node)
	err = json.Unmarshal(nodeByBytes, node)
	if err != nil {
		return nil , err
	}
	//返回节点信息
	return node , nil
}

//创建新的节点信息
func ( s *SmartContract ) CreateNode( ctx contractapi.TransactionContextInterface , nodeID string , resNeedType  []string , resNeedNum []int , resOwnType []string , resOwnNum []int ) error {
	//'{"Args":["CreateNode","NodeD",["A" , "B" , "C"],[1,0,2],["A" , "B" , "C"],[0,0,0]]}'
	//创建节点NodeD ，资源需要 A:1  C:2   , 初始拥有0
	exists, err := s.NodeExists( ctx, nodeID )
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf( "Node already exists ： %s" , nodeID )
	}
	//新建节点信息
	node := Node{
		NodeID: nodeID,
		ResNeed: Resource{ResType: resNeedType,ResNum: resNeedNum},
		ResOwn:  Resource{ResType: resOwnType,ResNum: resOwnNum},
	}
	//格式化节点信息为json
	nodeJsonByBytes, err := json.Marshal( node )
	if err != nil {
		return err
	}
	//写入账本
	return  ctx.GetStub().PutState( nodeID , nodeJsonByBytes )
}

//判断该节点信息是否存在账本当中
func ( s *SmartContract ) NodeExists( ctx contractapi.TransactionContextInterface , nodeID string ) ( bool , error )  {
	nodeJsonBytes, err := ctx.GetStub().GetState(nodeID)
	if err != nil {
		return false , err
	}
	return nodeJsonBytes != nil , nil
}

//初始化信息
func ( s * SmartContract ) InitLedger( ctx contractapi.TransactionContextInterface ) error  {
	//节点初始化
	//NodeA资源信息、需求信息、其他节点信息,初始化有各有10个资源；需求信息根据生产动态调整即可
	nodes := []Node{
		{
			NodeID: "NodeA",
			ResNeed: Resource{ResType: []string{"A","B","C"},ResNum: []int{0,2,3}},
			ResOwn:  Resource{ResType: []string{"A","B","C"},ResNum: []int{0,0,0}},
		},
		{
			NodeID: "NodeB",
			ResNeed: Resource{ResType: []string{"A","B","C"},ResNum: []int{0,0,0}},
			ResOwn:  Resource{ResType: []string{"A","B","C"},ResNum: []int{0,10,0}},
		},
		{
			NodeID: "NodeC",
			ResNeed: Resource{ResType: []string{"A","B","C"},ResNum: []int{0,0,0}},
			ResOwn:  Resource{ResType: []string{"A","B","C"},ResNum: []int{0,0,10}},
		},
	}
	//"{"nodeID":"NodeA","resNeed":{"resType":["A","B","C"],"resNum":[0,2,3]},"resOwn":{"resType":["A","B","C"],"resNum":[0,0,0]}}"
	for _, node := range nodes {
		nodeJsonBytes, err := json.Marshal(node)
		if err !=nil{
			return err
		}
		//将节点信息写入账本
		err = ctx.GetStub().PutState(node.NodeID, nodeJsonBytes)
		if err !=nil {
			return fmt.Errorf("failed put to world state : %v",err)
		}
	}
	return nil
}
func main()  {
	fmt.Println("--------------SmartContract--------------")
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panic("Error creating asset-transfer-basic chaincode :%v")
	}
	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
	}
}