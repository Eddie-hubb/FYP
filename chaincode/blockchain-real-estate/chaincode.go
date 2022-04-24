package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	//"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	//"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/routers"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
	"time"
	"encoding/json"
	"errors"
	"strconv"
	"math"
)

type BlockChainRealEstate struct {
}



// Init 链码初始化
func (t *BlockChainRealEstate) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("链码初始化")
	timeLocal, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		return shim.Error(fmt.Sprintf("时区设置失败%s", err))
	}
	time.Local = timeLocal
	//初始化默认数据
	var accountIds = [6]string{
		"5feceb66ffc8",
		"6b86b273ff34",
		"d4735e3a265e",
		"4e07408562be",
		"4b227777d4dd",
		"ef2d127de37b",
	}
	var userNames = [6]string{"Administrator", "Tranding Manager", "User 1", "Gold Whale", "Silver Whale", "Platinum Whale"}
	var balances = [6]float64{0, 5000000, 5000000, 5000000, 5000000, 5000000}
	var goldShares = [6]float64{0, 0, 0, 10000, 0, 0}
	var silverShares = [6]float64{0, 0, 0, 0, 10000, 0}
	var platinumShares = [6]float64{0, 0, 0, 0, 0, 10000}
	var commodityShareList = [6][]CommodityShareInfo{
		{{gold, 0}, {silver, 0}, {platinum, 0}}, 
		{{gold, 0}, {silver, 0}, {platinum, 0}},
		{{gold, 0}, {silver, 0}, {platinum, 0}},
		{{gold, 0}, {silver, 0}, {platinum, 0}},
		{{gold, 0}, {silver, 0}, {platinum, 0}},
		{{gold, 0}, {silver, 0}, {platinum, 0}}}
	var role = [6]string{"admin", "manager", "investor", "whale", "whale", "whale"}

	
	//初始化账号数据
	for i, val := range accountIds {
		account := &Account{
			AccountId: val,
			UserName:  userNames[i],
			Balance:   balances[i],
			GoldShare:	goldShares[i],
			SilverShare:	silverShares[i],
			PlatinumShare:	platinumShares[i],
			CommodityShareList:	commodityShareList[i],
			Role:	role[i],
		}
		// 写入账本
		if err := utils.WriteLedger(account, stub, AccountKey, []string{val}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}


	return shim.Success(nil)
}

// Invoke 实现Invoke接口调用智能合约
func (t *BlockChainRealEstate) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	
	case "createTransactionInfo":
		return createTransactionInfo(stub, args)
	// case "queryTransactionInfo":
	// 	return queryTransactionInfo(stub, args)
	case "createPortfolioInfo":
		return createPortfolioInfo(stub, args)
	case "createSuggestedPortfolioInfo":
		return createSuggestedPortfolioInfo(stub, args)
	case "updateState":
		return updateState(stub, args)
	case "adjustPortfolio":
		return adjustPortfolio(stub, args)
	case "adjustNetWorth":
		return adjustNetWorth(stub, args)
	case "queryAccountList":
		return QueryAccountList(stub, args)
	case "queryTransactionInfoList":
		return queryTransactionInfoList(stub, args)
	case "queryPortfolioList":
		return queryPortfolioList(stub, args)	
	case "queryMoneyTransactionList":
		return queryMoneyTransactionList(stub, args)
	case "queryCommodityTransactionList":
		return queryCommodityTransactionList(stub, args)
	case "queryRedemptionFeeTransactionList":
		return queryRedemptionFeeTransactionList(stub, args)
	case "queryServiceChargeTransactionList":
		return queryServiceChargeTransactionList(stub, args)
	case "queryCommodityTypeList":
		return queryCommodityTypeList(stub, args)
	case "get":
		return get(stub, args)
	case "set":
		return set(stub, args)

	
	
	

	default:
		return shim.Error(fmt.Sprintf("no this function: %s", funcName))
	}
}


const (
	layout = "2006-01-02"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

type SuggestedPortfolio struct {
	SuggestedPortfolioID 	string	`json:"portfolioID"`
	GoldShare		float64	`json:"goldShare"`
	SilverShare		float64	`json:"silverShare"`
	PlatinumShare	float64	`json:"platinumShare"`
	// CommodityShareList []CommodityShareInfo `json:"commodityShareList"`
	CreateTime    	string  `json:"createTime"`   
}

type Account struct {
	AccountId string  `json:"accountId"` 
	UserName  string  `json:"userName"`  
	Balance   float64 `json:"balance"`   
	GoldShare float64 `json:"goldShare"`
	SilverShare float64 `json:"silverShare"`
	PlatinumShare float64 `json:"platinumShare"`
	CommodityShareList []CommodityShareInfo `json:"commodityShareList"`
	Role 	string	`json:"role"` 
}

type Portfolio struct {
	PortfolioID 	string	`json:"portfolioID"`
	PortfolioState 	PortfolioStateType 	`json:"portfolioState"`
	AccountID 		string	`json:"accountID"`
	GoldShare		float64	`json:"goldShare"`
	SilverShare		float64	`json:"silverShare"`
	PlatinumShare	float64	`json:"platinumShare"`
	// CommodityShareList []CommodityShareInfo `json:"commodityShareList"`
	CreateTime    	string  `json:"createTime"`    
}

type PortfolioStateType struct {
	PortfolioStateTypeName	string	`json:"PortfolioStateTypeName"`
	PortfolioStateTypeId	int	`json:"PortfolioStateTypeId"`
}

type TransactionInfo struct {
	TransactionID   string	`json:"transactionID"`
	PortfolioID 	string	`json:"portfolioID"`//need to update
	BuyerID			string 	`json:"buyerID"`
	CommodityType     Commodity   `json:"commodityType"`
	TransactionStateType TransactionState	`json:"transactionStateType"`
	PurchaseAmount	float64	`json:"purchaseAmount"`
	NetWorth		float64	`json:"netWorth"`
	ServiceCharge 	float64	`json:"serviceCharge"`
	PurchaseShare   float64	`json:"purchaseShare"`
	RedemptionFee 	float64	`json:"redemptionFee"`
	SellAmount		float64	`json:"sellAmount"`
	SellShare		float64	`json:"sellShare"`	
	CreateTime    string  `json:"createTime"`    //创建时间
}

type MoneyTransaction struct{
	MoneyTransactionID string `json:"moneyTransactionID"`
	AmountOfMoney 	float64	`json:"amountOfMoney"`
	Sender			string	`json:"sender"`
	Receiver		string 	`json:"receiver"`
	CreateTime    string  `json:"createTime"`    //创建时间
}

type ServiceChargeTransaction struct {
	ServiceChargeTransactionID	string	`json:"serviceChargeTransactionID"`
	AmountOfCharge	float64	`json:"amountOfCharge"`
	Sender			string	`json:"sender"`
	Receiver		string	`json:"receiver"`
	CreateTime    string  `json:"createTime"`    //创建时间
}

type RedemptionFeeTransaction struct {
	RedemptionFeeTransactionID	string	`json:"eedemptionFeeTransactionID"`
	AmountOfFee		float64	`json:"amountOfFee"`
	Sender			string	`json:"sender"`
	Receiver		string	`json:"receiver"`
	CreateTime    string  `json:"createTime"`    //创建时间
}

type CommodityTransaction struct {
	CommodityTransactionID	string	`json:"commodityTransactionID"`
	CommodityType	Commodity	`json:"commodityType"`
	CommodityShare	float64	`json:"commodityShare"`
	Sender			string	`json:"sender"`
	Receiver 		string	`json:"receiver"`
	CreateTime    	string  `json:"createTime"`    //创建时间
}

type Commodity struct {
	CommodityName string	`json:"commodityName"`
	CommodityID int	`json:"commodityID"`
	CommodityNetWorth float64	`json:"commodityNetWorth"`
}


type CommodityShareInfo struct {
	CommodityType Commodity `json:"commodityType"`
	CommodityShare	float64 `json:"commodityShare"`
}

type TransactionState struct {
	TransactionStateName string	`json:"transactionStateName"`
	TransactionStateID int	`json:"transactionStateID"`
}


const (
	AccountKey         = "account-key"
	TransactionKey	   = "transaction-key"
	PortfolioKey	   = "portfolio-key"
	MoneyTransactionKey = "money-transaction-key"
	CommodityTransactionKey = "commodity-transaction-key"
	ServiceChargeTransactionKey = "service-charge-transaction-key"
	RedemptionFeeTransactionKey = "redemption-fee-transaction-key"
	SuggestedPortfoliokey = "suggested-portfolio-key"
)


var gold = Commodity{"gold", 1, 300.0}
var silver = Commodity{"silver", 2, 200.0}
var platinum = Commodity{"platinum", 3, 100.0}
var commodityList = [3]Commodity{gold, silver, platinum}

var stateInProcess = TransactionState{"inProcess", 1}
var stateSuccess = TransactionState{"success", 2}
var stateFail = TransactionState{"fail", 3}

var portfolioStateWorking = PortfolioStateType{"working", 1}
var portfolioStateExpired = PortfolioStateType{"expired", 2}

var tradingManagerID = "6b86b273ff34"
var goldWhaleID = "4e07408562be"
var silverWhaleID = "4b227777d4dd"
var platinumWhaleID = "ef2d127de37b"

//var currentNetWorth float64 = 100.0

var serviceChargeRate float64 = 0.005
var redemptionFeeRate float64 = 0.005




// Init of the chaincode
// This function is called only one when the chaincode is instantiated.
// So the goal is to prepare the ledger to handle future requests.

func Decimal(value float64) float64 {
    return math.Trunc(value*1e2+0.5) * 1e-2
}

// func CleanString(str string) string {
// 	trimStr := strings.Trim(string, "")
// 	return trimStr
// }

func adjustNetWorth(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	fmt.Println("running the function adjustNetWorth()")
	// var tempCommodityList []Commodity
	var err error
	gold.CommodityNetWorth,  err  = strconv.ParseFloat(args[0],64)
	silver.CommodityNetWorth,  err  = strconv.ParseFloat(args[1],64)
	platinum.CommodityNetWorth,  err  = strconv.ParseFloat(args[2],64)
	// tempCommodityList = append(tempCommodityList, gold)
	// tempCommodityList = append(tempCommodityList, silver)
	// tempCommodityList = append(tempCommodityList, platinum)
	commodityList = [3]Commodity{gold, silver, platinum}
	commodityListByte, err := json.Marshal(commodityList)
	if err != nil {
		return shim.Error(fmt.Sprintf("commodityList-序列化出错: %s", err))
	}
	return shim.Success(commodityListByte)


}



func QueryAccountList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var accountList []Account
	results, err := utils.GetStateByPartialCompositeKeys2(stub, AccountKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var account Account
			err := json.Unmarshal(v, &account)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
			}
			accountList = append(accountList, account)
		}
	}
	accountListByte, err := json.Marshal(accountList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAccountList-序列化出错: %s", err))
	}
	return shim.Success(accountListByte)
}

// accountID, GoldShare, SilverShare, PlatinumShare
func createSuggestedPortfolioInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	fmt.Println("running the function createSuggestedPortfolioInfo()")

	if len(args) != 4 {
		return shim.Error("Wrong input")
	}



	var portfolioInfo SuggestedPortfolio
	portfolioInfoID := stub.GetTxID()
	portfolioInfo.SuggestedPortfolioID = portfolioInfoID
	portfolioGoldShare,  err := strconv.ParseFloat(args[0],64)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("arg1")
	}
	portfolioInfo.GoldShare = portfolioGoldShare
	portfolioSilverShare,  err := strconv.ParseFloat(args[1],64)
	portfolioInfo.SilverShare = portfolioSilverShare
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("arg2")
	}
	portfolioPlatinumShare,  err := strconv.ParseFloat(args[2],64)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("arg3")
	}
	portfolioInfo.PlatinumShare = portfolioPlatinumShare
	portfolioInfo.CreateTime = args[3]


	jsonPortfolio, err := json.Marshal(portfolioInfo)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("Error marshalling to JSON")
	}

	// err = stub.PutState(portfolioInfoID, jsonPortfolio)
	// if err != nil {
	// 	return shim.Error("createPortfolio() : Error writing to state")
	// }

	if err := utils.WriteLedger(portfolioInfo, stub, SuggestedPortfoliokey, []string{portfolioInfo.SuggestedPortfolioID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	// Notify listeners that an event "eventInvoke" has been executed
	err = stub.SetEvent("eventInvoke", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	

	return shim.Success(jsonPortfolio)
	
}


// accountID, GoldShare, SilverShare, PlatinumShare
func createPortfolioInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	fmt.Println("running the function createPortfolioInfo()111" + time.Now().Local().Format("2006-01-02 15:04:05"))

	if len(args) != 4 {
		return shim.Error("Wrong input")
	}



	var portfolioInfo Portfolio
	portfolioInfoID := stub.GetTxID()
	portfolioInfo.PortfolioID = portfolioInfoID
	portfolioInfo.PortfolioState = portfolioStateWorking
	portfolioInfo.AccountID = args[0]
	portfolioGoldShare,  err := strconv.ParseFloat(args[1],64)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error(err.Error() + " arg0:"   + args[0])
	}
	portfolioInfo.GoldShare = portfolioGoldShare
	portfolioSilverShare,  err := strconv.ParseFloat(args[2],64)
	portfolioInfo.SilverShare = portfolioSilverShare
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("arg2")
	}
	portfolioPlatinumShare,  err := strconv.ParseFloat(args[3],64)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("arg3")
	}
	portfolioInfo.PlatinumShare = portfolioPlatinumShare
	portfolioInfo.CreateTime = time.Now().Local().Format("2006-01-02 15:04:05")


	if(portfolioInfo.GoldShare > 0){
		var goldString []string
		goldString = append(goldString, portfolioInfoID)//portfolio ID
		goldString = append(goldString, portfolioInfoID + "GoldTransaction")// transaction ID
		goldString = append(goldString, portfolioInfo.AccountID) // BuyerID
		goldString = append(goldString, "1")//type
		goldString = append(goldString, fmt.Sprintf("%f", portfolioInfo.GoldShare))//purchase share
		goldString = append(goldString, "0")//selling share
		createTransactionInfoforPortfolio(stub, goldString)

	}

	if(portfolioInfo.SilverShare > 0){
		var silverString []string
		silverString = append(silverString, portfolioInfoID)//portfolio ID
		silverString = append(silverString, portfolioInfoID + "SilverTransaction")// transaction ID
		silverString = append(silverString, portfolioInfo.AccountID) // BuyerID
		silverString = append(silverString, "2")//type
		silverString = append(silverString, fmt.Sprintf("%f", portfolioInfo.SilverShare))//purchase share
		silverString = append(silverString, "0")//selling share
		createTransactionInfoforPortfolio(stub, silverString)

	}

	if(portfolioInfo.PlatinumShare > 0){
		var platinumString []string
		platinumString = append(platinumString, portfolioInfoID)//portfolio ID
		platinumString = append(platinumString, portfolioInfoID + "PlatinumTransaction")// transaction ID
		platinumString = append(platinumString, portfolioInfo.AccountID) // BuyerID
		platinumString = append(platinumString, "3")//type
		platinumString = append(platinumString, fmt.Sprintf("%f", portfolioInfo.PlatinumShare))//purchase share
		platinumString = append(platinumString, "0")//selling share
		createTransactionInfoforPortfolio(stub, platinumString)
	}
	


	jsonPortfolio, err := json.Marshal(portfolioInfo)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("Error marshalling to JSON")
	}


	if err := utils.WriteLedger(portfolioInfo, stub, PortfolioKey, []string{portfolioInfo.AccountID, portfolioInfo.PortfolioID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	// err = stub.PutState(portfolioInfoID, jsonPortfolio)
	// if err != nil {
	// 	return shim.Error("createPortfolio() : Error writing to state")
	// }

	// Notify listeners that an event "eventInvoke" has been executed
	err = stub.SetEvent("eventInvoke", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	// valAsbytes, err := stub.GetState(portfolioInfoID)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return shim.Error(err.Error())
	// } else if valAsbytes == nil {
	// 	return shim.Error("valAsBytes nil")
	// }

	// var portfolioOutput Portfolio
	// err = json.Unmarshal(valAsbytes, &portfolioOutput)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return nil, errors.New("Error unmarshalling JSON")
	// }
	return shim.Success(jsonPortfolio)
	
}

func createTransactionInfoforPortfolio(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	fmt.Println("running the function createTransactionInfo()")

	if len(args) != 6 {
		return shim.Error("Wrong input")
	}



	var transactionInfo TransactionInfo
	transactionInfo.PortfolioID = args[0]
	transactionInfo.TransactionID = args[1]
	transactionInfo.BuyerID = args[2]
	if(args[3] == "1"){
		transactionInfo.CommodityType = gold
	}else if(args[3] == "2"){
		transactionInfo.CommodityType = silver
	}else{
		transactionInfo.CommodityType = platinum
	}
	// transactionInfo.TransactionStateType = stateInProcess
	//transactionInfo.NetWorth = currentNetWorth
	transactionInfo.TransactionStateType = stateInProcess
	transactionInfo.NetWorth = Decimal(transactionInfo.CommodityType.CommodityNetWorth)
	newPurchaseShare, err := strconv.ParseFloat(args[4],64)//need exception
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("arg1")
	}
	transactionInfo.PurchaseShare = newPurchaseShare
	transactionInfo.PurchaseAmount =  Decimal(transactionInfo.NetWorth * transactionInfo.PurchaseShare)
	transactionInfo.ServiceCharge = serviceChargeRate * transactionInfo.NetWorth * transactionInfo.PurchaseShare
	
	newSellShare, err := strconv.ParseFloat(args[5],64)//need exception
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("arg1")
	} 
	transactionInfo.SellShare = newSellShare	

	transactionInfo.SellAmount =  Decimal(transactionInfo.NetWorth * transactionInfo.SellShare)
	transactionInfo.RedemptionFee = Decimal(redemptionFeeRate * transactionInfo.NetWorth * transactionInfo.SellShare)
	transactionInfo.CreateTime = time.Now().Local().Format("2006-01-02 15:04:05")




	// jsonTransaction, err := json.Marshal(transactionInfo)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return shim.Error("Error marshalling to JSON")
	// }

	// err = stub.PutState(transactionInfo.TransactionID, jsonTransaction)
	// if err != nil {
	// 	return shim.Error("createTransactionInfo() : Error writing to state")
	// }

	if err := utils.WriteLedger(transactionInfo, stub, TransactionKey, []string{transactionInfo.BuyerID, transactionInfo.TransactionID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	var buyerAccount Account
	buyerAccount.getAccount(stub, transactionInfo.BuyerID)

	var tradingManagerAccount Account
	tradingManagerAccount.getAccount(stub, tradingManagerID)

	

	if(transactionInfo.PurchaseShare > 0){		
		costingMoney := transactionInfo.PurchaseAmount + transactionInfo.ServiceCharge

		if buyerAccount.Balance < costingMoney {
			return shim.Error(fmt.Sprintf("Balance is not enought"))
		}
		var moneyTransaction = MoneyTransaction{transactionInfo.TransactionID + "DepositReceivedMoney", transactionInfo.PurchaseAmount, transactionInfo.BuyerID,  tradingManagerID, time.Now().Local().Format("2006-01-02 15:04:05")}
		// jsonTransaction, err := json.Marshal(moneyTransaction)
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	return shim.Error("Error marshalling to JSON")
		// }

		// err = stub.PutState(moneyTransaction.MoneyTransactionID, jsonTransaction)
		// if err != nil {
		// 	return shim.Error("createMoneyTransaction() : Error writing to state")
		// }

		if err := utils.WriteLedger(moneyTransaction, stub, MoneyTransactionKey, []string{moneyTransaction.Sender, moneyTransaction.MoneyTransactionID, moneyTransaction.Receiver}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	

		buyerAccount.Balance -= costingMoney
		if err := utils.WriteLedger(buyerAccount, stub, AccountKey, []string{buyerAccount.AccountId}); err != nil {
			return shim.Error(fmt.Sprintf("fail to reduce user balance %s", err))
		}

		tradingManagerAccount.Balance += costingMoney
		if err := utils.WriteLedger(tradingManagerAccount, stub, AccountKey, []string{tradingManagerAccount.AccountId}); err != nil {
			return shim.Error(fmt.Sprintf("fail to reduce user balance %s", err))
		}



	}else {
		var commodityTransaction = CommodityTransaction{transactionInfo.TransactionID + "DepositReceivedCommodity", transactionInfo.CommodityType, transactionInfo.SellShare,  transactionInfo.BuyerID,  tradingManagerID, time.Now().Local().Format("2006-01-02 15:04:05")}
		// jsonTransaction, err := json.Marshal(commodityTransaction)
		// if err != nil {
		// 	fmt.Println(err.Error()) 
		// if err != nil {
		// 	return shim.Error("createcommodityTransaction() : Error writing to state")
		// }

		// err = stub.PutState(commodityTransaction.CommodityTransactionID, jsonTransaction)

		if err := utils.WriteLedger(commodityTransaction, stub, CommodityTransactionKey, []string{commodityTransaction.Sender, commodityTransaction.CommodityTransactionID, commodityTransaction.Receiver}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}

		if(args[3] == "1"){
			buyerAccount.GoldShare -= transactionInfo.SellShare
			tradingManagerAccount.GoldShare += transactionInfo.SellShare
		}else if(args[3] == "2"){
			buyerAccount.SilverShare -= transactionInfo.SellShare
			tradingManagerAccount.SilverShare += transactionInfo.SellShare
		}else{
			buyerAccount.PlatinumShare -= transactionInfo.SellShare
			tradingManagerAccount.PlatinumShare += transactionInfo.SellShare
		}
		
		if err := utils.WriteLedger(buyerAccount, stub, AccountKey, []string{buyerAccount.AccountId}); err != nil {
			return shim.Error(fmt.Sprintf("fail to reduce user balance %s", err))
		}

		
		if err := utils.WriteLedger(tradingManagerAccount, stub, AccountKey, []string{tradingManagerAccount.AccountId}); err != nil {
			return shim.Error(fmt.Sprintf("fail to reduce user balance %s", err))
		}

	}
	
	// Notify listeners that an event "eventInvoke" has been executed
	err = stub.SetEvent("eventInvoke", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(transactionInfo.TransactionID))
	
}

func createTransactionInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	fmt.Println("running the function createTransactionInfo()")

	if len(args) != 5 {
		return shim.Error("Wrong input")
	}



	var transactionInfo TransactionInfo
	transactionID := stub.GetTxID()
	transactionInfo.TransactionID = transactionID
	transactionInfo.PortfolioID = args[0]
	transactionInfo.BuyerID = args[1]
	if(args[2] == "1"){
		transactionInfo.CommodityType = gold
	}else if(args[2] == "2"){
		transactionInfo.CommodityType = silver
	}else{
		transactionInfo.CommodityType = platinum
	}
	// transactionInfo.TransactionStateType = stateInProcess
	//transactionInfo.NetWorth = currentNetWorth
	transactionInfo.TransactionStateType = stateInProcess
	transactionInfo.NetWorth = Decimal(transactionInfo.CommodityType.CommodityNetWorth)
	newPurchaseShare, err := strconv.ParseFloat(args[3],64)//need exception
	transactionInfo.PurchaseShare = newPurchaseShare
	transactionInfo.PurchaseAmount =  Decimal(transactionInfo.NetWorth * transactionInfo.PurchaseShare)
	transactionInfo.ServiceCharge = serviceChargeRate * transactionInfo.NetWorth * transactionInfo.PurchaseShare
	
	newSellShare, err := strconv.ParseFloat(args[4],64)//need exception
	transactionInfo.SellShare = newSellShare	

	transactionInfo.SellAmount =  Decimal(transactionInfo.NetWorth * transactionInfo.SellShare)
	transactionInfo.RedemptionFee = Decimal(redemptionFeeRate * transactionInfo.NetWorth * transactionInfo.SellShare)
	transactionInfo.CreateTime = time.Now().Local().Format("2006-01-02 15:04:05")




	jsonTransaction, err := json.Marshal(transactionInfo)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("Error marshalling to JSON")
	}

	// err = stub.PutState(transactionID, jsonTransaction)
	// if err != nil {
	// 	return shim.Error("createTransactionInfo() : Error writing to state111")
	// }

	if err := utils.WriteLedger(transactionInfo, stub, TransactionKey, []string{transactionInfo.BuyerID, transactionInfo.TransactionID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}


	valAsbytes, err := utils.GetStateByPartialCompositeKeys(stub, AccountKey, []string{transactionInfo.BuyerID})

	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("createTransactionInfo() : Error writing to state222")
	}else if valAsbytes == nil {
		return shim.Error("createTransactionInfo() : Error writing to state333")
	}
	

	var buyerAccount Account
	

	if err = json.Unmarshal(valAsbytes[0], &buyerAccount); err != nil {
		return shim.Error(fmt.Sprintf("Unmarshal fail: %s", err))
	}
	if buyerAccount.UserName == "Administrator" {
		return shim.Error(fmt.Sprintf("Admin cannot purchase%s", err))
	}

	valAsbytes, err = utils.GetStateByPartialCompositeKeys(stub, AccountKey, []string{tradingManagerID})

	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("createTransactionInfo() : Error writing to state444")
	}else if valAsbytes == nil {
		return shim.Error("createTransactionInfo() : Error writing to state555")
	}
	

	var tradingManagerAccount Account

	if err = json.Unmarshal(valAsbytes[0], &tradingManagerAccount); err != nil {
		return shim.Error(fmt.Sprintf("Unmarshal fail: %s", err))
	}
	if buyerAccount.UserName == "Administrator" {
		return shim.Error(fmt.Sprintf("Admin cannot purchase%s", err))
	}
	

	if(transactionInfo.PurchaseShare > 0){		
		costingMoney := transactionInfo.PurchaseAmount + transactionInfo.ServiceCharge

		if buyerAccount.Balance < costingMoney {
			return shim.Error(fmt.Sprintf("Balance is not enought"))
		}
		var moneyTransaction = MoneyTransaction{transactionInfo.TransactionID + "DepositReceivedMoney", transactionInfo.PurchaseAmount, transactionInfo.BuyerID,  tradingManagerID, time.Now().Local().Format("2006-01-02 15:04:05")}
		// jsonTransaction, err := json.Marshal(moneyTransaction)
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	return shim.Error("Error marshalling to JSON")
		// }

		// err = stub.PutState(moneyTransaction.MoneyTransactionID, jsonTransaction)
		// if err != nil {
		// 	return shim.Error("createMoneyTransaction() : Error writing to state666")
		// }

		if err := utils.WriteLedger(moneyTransaction, stub, MoneyTransactionKey, []string{moneyTransaction.Sender, moneyTransaction.MoneyTransactionID, moneyTransaction.Receiver}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}

		buyerAccount.Balance -= costingMoney
		if err := utils.WriteLedger(buyerAccount, stub, AccountKey, []string{buyerAccount.AccountId}); err != nil {
			return shim.Error(fmt.Sprintf("fail to reduce user balance %s", err))
		}

		tradingManagerAccount.Balance += costingMoney
		if err := utils.WriteLedger(tradingManagerAccount, stub, AccountKey, []string{tradingManagerAccount.AccountId}); err != nil {
			return shim.Error(fmt.Sprintf("fail to reduce user balance %s", err))
		}



	}else {
		var commodityTransaction = CommodityTransaction{transactionInfo.TransactionID + "DepositReceivedCommodity", transactionInfo.CommodityType, transactionInfo.SellShare,  transactionInfo.BuyerID,  tradingManagerID, time.Now().Local().Format("2006-01-02 15:04:05")}
		// jsonTransaction, err := json.Marshal(commodityTransaction)
		// if err != nil {
		// 	fmt.Println(err.Error()) 
		// if err != nil {
		// 	return shim.Error("createcommodityTransaction() : Error writing to state777")
		// }
		// err = stub.PutState(commodityTransaction.CommodityTransactionID, jsonTransaction)
		
		if err := utils.WriteLedger(commodityTransaction, stub, CommodityTransactionKey, []string{commodityTransaction.Sender, commodityTransaction.CommodityTransactionID, commodityTransaction.Receiver}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}

		if(args[3] == "1"){
			buyerAccount.GoldShare -= transactionInfo.SellShare
			tradingManagerAccount.GoldShare += transactionInfo.SellShare
		}else if(args[3] == "2"){
			buyerAccount.SilverShare -= transactionInfo.SellShare
			tradingManagerAccount.SilverShare += transactionInfo.SellShare
		}else{
			buyerAccount.PlatinumShare -= transactionInfo.SellShare
			tradingManagerAccount.PlatinumShare += transactionInfo.SellShare
		}
		
		if err := utils.WriteLedger(buyerAccount, stub, AccountKey, []string{buyerAccount.AccountId}); err != nil {
			return shim.Error(fmt.Sprintf("fail to reduce user balance %s", err))
		}

		
		if err := utils.WriteLedger(tradingManagerAccount, stub, AccountKey, []string{tradingManagerAccount.AccountId}); err != nil {
			return shim.Error(fmt.Sprintf("fail to reduce user balance %s", err))
		}

	}
	
	
	// Notify listeners that an event "eventInvoke" has been executed
	err = stub.SetEvent("eventInvoke", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(jsonTransaction)
	
}

func (portfolio *Portfolio) getPortfolio(stub shim.ChaincodeStubInterface, accountId string, ID string) error{
	valAsbytes, err := utils.GetStateByPartialCompositeKeys2(stub, PortfolioKey, []string{accountId, ID})

	if err != nil {
		fmt.Println(err.Error())
		return err
	} else if valAsbytes == nil {
		return err
	}


	if err = json.Unmarshal(valAsbytes[0], &portfolio); err != nil {
		return err
	}
	
	return nil
}

func adjustPortfolio(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	fmt.Println("running the function createPortfolioInfo()")

	if len(args) != 5{
		return shim.Error("Wrong input")
	}
	
	previousPortfolioID := args[0]
	var previousPortfolio Portfolio
	err = previousPortfolio.getPortfolio(stub, args[1], previousPortfolioID)
	previousPortfolio.PortfolioState = portfolioStateExpired

	var portfolioInfo Portfolio
	portfolioInfoID := stub.GetTxID()
	portfolioInfo.PortfolioID = portfolioInfoID
	portfolioInfo.PortfolioState = portfolioStateWorking
	portfolioInfo.AccountID = args[1]
	portfolioGoldShare,  err := strconv.ParseFloat(args[2],64)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("arg1")
	}
	portfolioInfo.GoldShare = portfolioGoldShare
	portfolioSilverShare,  err := strconv.ParseFloat(args[3],64)
	portfolioInfo.SilverShare = portfolioSilverShare
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("arg2")
	}
	portfolioPlatinumShare,  err := strconv.ParseFloat(args[4],64)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("arg3")
	}
	portfolioInfo.PlatinumShare = portfolioPlatinumShare
	portfolioInfo.CreateTime = time.Now().Local().Format("2006-01-02 15:04:05")


	if(portfolioInfo.GoldShare - previousPortfolio.GoldShare> 0){
		var goldString []string
		goldString = append(goldString, portfolioInfoID)//portfolio ID
		goldString = append(goldString, portfolioInfoID + "GoldTransaction")// transaction ID
		goldString = append(goldString, portfolioInfo.AccountID) // BuyerID
		goldString = append(goldString, "1")//type
		goldString = append(goldString, fmt.Sprintf("%f", portfolioInfo.GoldShare - previousPortfolio.GoldShare))//purchase share
		goldString = append(goldString, "0")//selling share
		createTransactionInfoforPortfolio(stub, goldString)

	}else if(portfolioInfo.GoldShare - previousPortfolio.GoldShare< 0){
		var goldString []string
		goldString = append(goldString, portfolioInfoID)//portfolio ID
		goldString = append(goldString, portfolioInfoID + "GoldTransaction")// transaction ID
		goldString = append(goldString, portfolioInfo.AccountID) // BuyerID
		goldString = append(goldString, "1")//type
		goldString = append(goldString, "0")//purchase share
		goldString = append(goldString, fmt.Sprintf("%f", previousPortfolio.GoldShare - portfolioInfo.GoldShare))//selling share
		createTransactionInfoforPortfolio(stub, goldString)
	}



	if(portfolioInfo.SilverShare - previousPortfolio.SilverShare> 0){
		var silverString []string
		silverString = append(silverString, portfolioInfoID)//portfolio ID
		silverString = append(silverString, portfolioInfoID + "SilverTransaction")// transaction ID
		silverString = append(silverString, portfolioInfo.AccountID) // BuyerID
		silverString = append(silverString, "2")//type
		silverString = append(silverString, fmt.Sprintf("%f", portfolioInfo.SilverShare - previousPortfolio.SilverShare))//purchase share
		silverString = append(silverString, "0")//selling share
		createTransactionInfoforPortfolio(stub, silverString)

	}else if(portfolioInfo.SilverShare - previousPortfolio.SilverShare< 0){
		var silverString []string
		silverString = append(silverString, portfolioInfoID)//portfolio ID
		silverString = append(silverString, portfolioInfoID + "SilverTransaction")// transaction ID
		silverString = append(silverString, portfolioInfo.AccountID) // BuyerID
		silverString = append(silverString, "2")//type
		silverString = append(silverString, "0")//purchase share
		silverString = append(silverString, fmt.Sprintf("%f", previousPortfolio.SilverShare - portfolioInfo.SilverShare))//selling share
		createTransactionInfoforPortfolio(stub, silverString)

	}

	if(portfolioInfo.PlatinumShare - previousPortfolio.PlatinumShare > 0){
		var platinumString []string
		platinumString = append(platinumString, portfolioInfoID)//portfolio ID
		platinumString = append(platinumString, portfolioInfoID + "PlatinumTransaction")// transaction ID
		platinumString = append(platinumString, portfolioInfo.AccountID) // BuyerID
		platinumString = append(platinumString, "3")//type
		platinumString = append(platinumString, fmt.Sprintf("%f", portfolioInfo.PlatinumShare - previousPortfolio.PlatinumShare))//purchase share
		platinumString = append(platinumString, "0")//selling share
		createTransactionInfoforPortfolio(stub, platinumString)
	}else if(portfolioInfo.PlatinumShare - previousPortfolio.PlatinumShare < 0){
		var platinumString []string
		platinumString = append(platinumString, portfolioInfoID)//portfolio ID
		platinumString = append(platinumString, portfolioInfoID + "PlatinumTransaction")// transaction ID
		platinumString = append(platinumString, portfolioInfo.AccountID) // BuyerID
		platinumString = append(platinumString, "3")//type
		platinumString = append(platinumString, "0")//purchase share
		platinumString = append(platinumString, fmt.Sprintf("%f", previousPortfolio.PlatinumShare - portfolioInfo.PlatinumShare ))//selling share
		createTransactionInfoforPortfolio(stub, platinumString)
	}


	jsonPortfolio, err := json.Marshal(portfolioInfo)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("Error marshalling to JSON")
	}

	// err = stub.PutState(portfolioInfoID, jsonPortfolio)
	// if err != nil {
	// 	return shim.Error("createPortfolio() : Error writing to state")
	// }

	if err := utils.WriteLedger(portfolioInfo, stub, PortfolioKey, []string{portfolioInfo.AccountID, portfolioInfo.PortfolioID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	if err := utils.WriteLedger(previousPortfolio, stub, PortfolioKey, []string{previousPortfolio.AccountID, previousPortfolio.PortfolioID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	// Notify listeners that an event "eventInvoke" has been executed
	err = stub.SetEvent("eventInvoke", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	

	return shim.Success(jsonPortfolio)

}



// func queryTransactionInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	var transactionID string
// 	var err error

// 	if len(args) != 1 {
// 		return shim.Error("Wrong input")
// 	}
// 	transactionID = args[0]
// 	valAsbytes, err := stub.GetState(transactionID)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return shim.Error(err.Error())
// 	} else if valAsbytes == nil {
// 		fmt.Println("Transaction does not exist")
// 		return shim.Error("Transaction does not exist")
// 	}

// 	// accountListByte, err := json.Marshal(accountList)
// 	// if err != nil {
// 	// 	return shim.Error(fmt.Sprintf("QueryAccountList-序列化出错: %s", err))
// 	// }

// 	return shim.Success(valAsbytes)
// }



func (transactionInfo *TransactionInfo) updateTransactionState(newState string) error {
	if(newState == "1"){
		transactionInfo.TransactionStateType = stateInProcess
	}else if (newState == "2"){
		transactionInfo.TransactionStateType = stateSuccess
	}else{
		transactionInfo.TransactionStateType = stateFail
	}
	return nil
}

func (account *Account) getAccount(stub shim.ChaincodeStubInterface, ID string) error {
	valAsbytes, err := utils.GetStateByPartialCompositeKeys(stub, AccountKey, []string{ID})

	if err != nil {
		fmt.Println(err.Error())
		return err
	} else if valAsbytes == nil {
		return err
	}


	if err = json.Unmarshal(valAsbytes[0], &account); err != nil {
		return err
	}
	if account.UserName == "Administrator" {
		return err
	}
	return nil
}



func updateState(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var transactionID string
	var err error
	var transactionInfo TransactionInfo

	if len(args) != 3 { 
		return shim.Error("Wrong input")
	}
	accountID := args[0]
	transactionID = args[1]
	err = transactionInfo.queryTransactionInfobyID(stub, accountID, transactionID)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error(err.Error())
	}
	// if transactionInfo == nil {
	// 	fmt.Println("Error reading state : transactionInfo is nil")
	// 	return shim.Error("nil transaction")
	// }
	newTransactionState := args[2]
	err = transactionInfo.updateTransactionState(newTransactionState)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error(err.Error())
	}

	jsonTransaction, err := json.Marshal(transactionInfo)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("error marshalling json" + err.Error())
	}

	// err = stub.PutState(transactionID, jsonTransaction)
	// if err != nil {
	// 	return shim.Error("updateTransactionInfo() : Error put state")
	// }

	if err := utils.WriteLedger(transactionInfo, stub, TransactionKey, []string{transactionInfo.BuyerID, transactionInfo.TransactionID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	if (newTransactionState == "2"){
		var whaleID string
		if(transactionInfo.CommodityType.CommodityID == 1){
			whaleID = goldWhaleID
		}else if(transactionInfo.CommodityType.CommodityID == 2){
			whaleID = silverWhaleID
		}else{
			whaleID = platinumWhaleID
		}

		var whaleAccount Account
		err = whaleAccount.getAccount(stub, whaleID)


		
		var buyerAccount Account
		err = buyerAccount.getAccount(stub, transactionInfo.BuyerID)

		var tradingManagerAccount Account
		err = tradingManagerAccount.getAccount(stub,tradingManagerID)

		// var goldWhale Account
		// err = goldWhale.getAccount(stub, whaleID)
		
		// var silverWhale Account
		// err = silverWhale.getAccount(stub, whaleID)

		// var platinumWhale Account
		// err = platinumWhale.getAccount(stub, whaleID)



		if(transactionInfo.PurchaseShare == 0){		

			var commodityTransactionTrade2Whale = CommodityTransaction{transactionInfo.TransactionID + "PurchaseCommodityTrade2Whale", transactionInfo.CommodityType, transactionInfo.SellShare,  tradingManagerID, whaleID, time.Now().Local().Format("2006-01-02 15:04:05")}


			if err := utils.WriteLedger(commodityTransactionTrade2Whale, stub, CommodityTransactionKey, []string{commodityTransactionTrade2Whale.Receiver, commodityTransactionTrade2Whale.CommodityTransactionID, commodityTransactionTrade2Whale.Sender}); err != nil {
				return shim.Error(fmt.Sprintf("%s", err))
			}

			

			var moneyTransactionWhale2Trade = MoneyTransaction{transactionInfo.TransactionID + "RedemptionMoneyWhale2Trade", transactionInfo.SellAmount, whaleID,  tradingManagerID, time.Now().Local().Format("2006-01-02 15:04:05")}

			
			if err := utils.WriteLedger(moneyTransactionWhale2Trade, stub, MoneyTransactionKey, []string{moneyTransactionWhale2Trade.Sender, moneyTransactionWhale2Trade.MoneyTransactionID, moneyTransactionWhale2Trade.Receiver}); err != nil {
				return shim.Error(fmt.Sprintf("%s", err))
			}

			if err != nil {
				return shim.Error("createMoneyTransaction() : Error writing to state")
			}

			var moneyTransactionTrade2User = MoneyTransaction{transactionInfo.TransactionID + "RedemptionMoneyTrade2User", transactionInfo.SellAmount,  tradingManagerID, transactionInfo.BuyerID, time.Now().Local().Format("2006-01-02 15:04:05")}


			if err := utils.WriteLedger(moneyTransactionTrade2User, stub, MoneyTransactionKey, []string{moneyTransactionTrade2User.MoneyTransactionID, moneyTransactionTrade2User.Sender}); err != nil {
				return shim.Error(fmt.Sprintf("%s", err))
			}

			var redemptionFeeTransaction = RedemptionFeeTransaction{transactionInfo.TransactionID + "RedemptionFeeFromUser", transactionInfo.RedemptionFee,  transactionInfo.BuyerID,  tradingManagerID, time.Now().Local().Format("2006-01-02 15:04:05")}


			if err := utils.WriteLedger(redemptionFeeTransaction, stub, RedemptionFeeTransactionKey, []string{redemptionFeeTransaction.Sender, redemptionFeeTransaction.RedemptionFeeTransactionID, redemptionFeeTransaction.Receiver}); err != nil {
				return shim.Error(fmt.Sprintf("%s", err))
			}

			var serviceChargeTransaction = ServiceChargeTransaction{transactionInfo.TransactionID + "ServiceChargeFromWhale", transactionInfo.SellAmount * serviceChargeRate, whaleAccount.AccountId,  tradingManagerID, time.Now().Local().Format("2006-01-02 15:04:05")}

			if err := utils.WriteLedger(serviceChargeTransaction, stub, ServiceChargeTransactionKey, []string{serviceChargeTransaction.Sender, serviceChargeTransaction.ServiceChargeTransactionID, serviceChargeTransaction.Receiver}); err != nil {
				return shim.Error(fmt.Sprintf("%s", err))
			}

			if(transactionInfo.CommodityType.CommodityID == 1){
				tradingManagerAccount.GoldShare -= transactionInfo.SellShare
				whaleAccount.GoldShare += transactionInfo.SellShare

			}else if(transactionInfo.CommodityType.CommodityID == 2){
				tradingManagerAccount.SilverShare -= transactionInfo.SellShare
				whaleAccount.SilverShare += transactionInfo.SellShare

			}else{
				tradingManagerAccount.PlatinumShare -= transactionInfo.SellShare
				whaleAccount.PlatinumShare += transactionInfo.SellShare
			}
			whaleAccount.Balance -= (transactionInfo.SellAmount + transactionInfo.ServiceCharge)
			buyerAccount.Balance += (transactionInfo.SellAmount - transactionInfo.RedemptionFee)
			tradingManagerAccount.Balance += (transactionInfo.RedemptionFee + transactionInfo.ServiceCharge)


		}else {
			var moneyTransactionTrade2Whale = MoneyTransaction{transactionInfo.TransactionID + "PurchaseMoneyTrade2Whale", transactionInfo.PurchaseAmount,  tradingManagerID, whaleID, time.Now().Local().Format("2006-01-02 15:04:05")}


			if err := utils.WriteLedger(moneyTransactionTrade2Whale, stub, MoneyTransactionKey, []string{moneyTransactionTrade2Whale.Receiver, moneyTransactionTrade2Whale.MoneyTransactionID, moneyTransactionTrade2Whale.Sender}); err != nil {
				return shim.Error(fmt.Sprintf("%s", err))
			}


			var commodityTransactionWhale2Trade = CommodityTransaction{transactionInfo.TransactionID + "CommodityWhale2Trade", transactionInfo.CommodityType, transactionInfo.PurchaseShare, whaleID,  tradingManagerID, time.Now().Local().Format("2006-01-02 15:04:05")}
		

			if err := utils.WriteLedger(commodityTransactionWhale2Trade, stub, CommodityTransactionKey, []string{commodityTransactionWhale2Trade.Sender, commodityTransactionWhale2Trade.CommodityTransactionID, commodityTransactionWhale2Trade.Receiver}); err != nil {
				return shim.Error(fmt.Sprintf("%s", err))
			}

			var commodityTransactionTrade2User = CommodityTransaction{transactionInfo.TransactionID + "CommodityTrade2User", transactionInfo.CommodityType, transactionInfo.PurchaseShare,  tradingManagerID, transactionInfo.BuyerID, time.Now().Local().Format("2006-01-02 15:04:05")}

			if err := utils.WriteLedger(commodityTransactionTrade2User, stub, CommodityTransactionKey, []string{commodityTransactionTrade2User.Receiver, commodityTransactionTrade2User.CommodityTransactionID, commodityTransactionTrade2User.Sender}); err != nil {
				return shim.Error(fmt.Sprintf("%s", err))
			}

			var serviceChargeTransaction = ServiceChargeTransaction{transactionInfo.TransactionID + "ServiceChargeFromUser", transactionInfo.ServiceCharge, transactionInfo.BuyerID,  tradingManagerID, time.Now().Local().Format("2006-01-02 15:04:05")}

			if err := utils.WriteLedger(serviceChargeTransaction, stub, ServiceChargeTransactionKey, []string{serviceChargeTransaction.Sender, serviceChargeTransaction.ServiceChargeTransactionID, serviceChargeTransaction.Receiver}); err != nil {
				return shim.Error(fmt.Sprintf("%s", err))
			}

			var redemptionFeeTransaction = RedemptionFeeTransaction{transactionInfo.TransactionID + "RedemptionFeeFromWhale", transactionInfo.PurchaseAmount * redemptionFeeRate,  whaleAccount.AccountId,  tradingManagerID, time.Now().Local().Format("2006-01-02 15:04:05")}


			if err := utils.WriteLedger(redemptionFeeTransaction, stub, RedemptionFeeTransactionKey, []string{redemptionFeeTransaction.Sender, redemptionFeeTransaction.RedemptionFeeTransactionID, redemptionFeeTransaction.Receiver}); err != nil {
				return shim.Error(fmt.Sprintf("%s", err))
			}
			

			if(transactionInfo.CommodityType.CommodityID == 1){
				tradingManagerAccount.GoldShare -= transactionInfo.PurchaseShare
				whaleAccount.GoldShare -= transactionInfo.PurchaseShare
				buyerAccount.GoldShare += transactionInfo.PurchaseShare
			}else if(transactionInfo.CommodityType.CommodityID == 2){
				tradingManagerAccount.SilverShare -= transactionInfo.PurchaseShare
				whaleAccount.SilverShare -= transactionInfo.PurchaseShare
				buyerAccount.SilverShare += transactionInfo.PurchaseShare

			}else{
				tradingManagerAccount.PlatinumShare -= transactionInfo.PurchaseShare
				whaleAccount.PlatinumShare -= transactionInfo.PurchaseShare
				buyerAccount.PlatinumShare += transactionInfo.PurchaseShare

			}
			whaleAccount.Balance += (transactionInfo.PurchaseAmount - transactionInfo.RedemptionFee)
			tradingManagerAccount.Balance -= (transactionInfo.PurchaseAmount - transactionInfo.RedemptionFee)

		}

		if err := utils.WriteLedger(buyerAccount, stub, AccountKey, []string{buyerAccount.AccountId}); err != nil {
			return shim.Error(fmt.Sprintf("fail to reduce user balance %s", err))
		}
		
		if err := utils.WriteLedger(tradingManagerAccount, stub, AccountKey, []string{tradingManagerAccount.AccountId}); err != nil {
			return shim.Error(fmt.Sprintf("fail to reduce user balance %s", err))
		}

		if err := utils.WriteLedger(whaleAccount, stub, AccountKey, []string{whaleID}); err != nil {
			return shim.Error(fmt.Sprintf("fail to reduce user balance %s", err))
		}

	}

	// Notify listeners that an event "eventInvoke" has been executed
	err = stub.SetEvent("eventInvoke", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}



	return shim.Success(jsonTransaction)

}


func (transactionInfo *TransactionInfo) queryTransactionInfobyID(stub shim.ChaincodeStubInterface, accountID string, ID string) error {
	valAsbytes, err := utils.GetStateByPartialCompositeKeys2(stub, TransactionKey, []string{accountID, ID})

	if err != nil {
		fmt.Println(err.Error())
		return err
	} else if valAsbytes == nil {
		return errors.New("Transaction does not exist, queryTransactionInfobyID")
	}


	if err = json.Unmarshal(valAsbytes[0], &transactionInfo); err != nil {
		return err
	}
	
	return nil
	
}

// QuerySellingList 查询销售(可查询所有，也可根据发起销售人查询)(发起的)(供卖家查询)
func queryTransactionInfoList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var transactionList []TransactionInfo
	results, err := utils.GetStateByPartialCompositeKeys2(stub, TransactionKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var transaction TransactionInfo
			err := json.Unmarshal(v, &transaction)
			if err != nil {
				return shim.Error(fmt.Sprintf("queryTransactionInfoList-反序列化出错: %s", err))
			}
			transactionList = append(transactionList, transaction)
		}
	}
	transactionListByte, err := json.Marshal(transactionList)
	if err != nil {
		return shim.Error(fmt.Sprintf("queryTransactionInfoList-序列化出错: %s", err))
	}
	return shim.Success(transactionListByte)
}

// QuerySellingList 查询销售(可查询所有，也可根据发起销售人查询)(发起的)(供卖家查询)
func queryPortfolioList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var portfolioList []Portfolio
	results, err := utils.GetStateByPartialCompositeKeys2(stub, PortfolioKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var transaction Portfolio
			err := json.Unmarshal(v, &transaction)
			if err != nil {
				return shim.Error(fmt.Sprintf("queryTransactionInfoList-反序列化出错: %s", err))
			}
			portfolioList = append(portfolioList, transaction)
		}
	}
	transactionListByte, err := json.Marshal(portfolioList)
	if err != nil {
		return shim.Error(fmt.Sprintf("queryTransactionInfoList-序列化出错: %s", err))
	}
	return shim.Success(transactionListByte)
}



// QuerySellingList 查询销售(可查询所有，也可根据发起销售人查询)(发起的)(供卖家查询)
func queryMoneyTransactionList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var transactionList []MoneyTransaction
	results, err := utils.GetStateByPartialCompositeKeys2(stub, MoneyTransactionKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var transaction MoneyTransaction
			err := json.Unmarshal(v, &transaction)
			if err != nil {
				return shim.Error(fmt.Sprintf("queryTransactionInfoList-反序列化出错: %s", err))
			}
			transactionList = append(transactionList, transaction)
		}
	}
	transactionListByte, err := json.Marshal(transactionList)
	if err != nil {
		return shim.Error(fmt.Sprintf("queryTransactionInfoList-序列化出错: %s", err))
	}
	return shim.Success(transactionListByte)
}



// QuerySellingList 查询销售(可查询所有，也可根据发起销售人查询)(发起的)(供卖家查询)
func queryCommodityTransactionList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var transactionList []CommodityTransaction
	results, err := utils.GetStateByPartialCompositeKeys2(stub, CommodityTransactionKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var transaction CommodityTransaction
			err := json.Unmarshal(v, &transaction)
			if err != nil {
				return shim.Error(fmt.Sprintf("queryTransactionInfoList-反序列化出错: %s", err))
			}
			transactionList = append(transactionList, transaction)
		}
	}
	transactionListByte, err := json.Marshal(transactionList)
	if err != nil {
		return shim.Error(fmt.Sprintf("queryTransactionInfoList-序列化出错: %s", err))
	}
	return shim.Success(transactionListByte)
}



// QuerySellingList 查询销售(可查询所有，也可根据发起销售人查询)(发起的)(供卖家查询)
func queryRedemptionFeeTransactionList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var transactionList []RedemptionFeeTransaction
	results, err := utils.GetStateByPartialCompositeKeys2(stub, RedemptionFeeTransactionKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var transaction RedemptionFeeTransaction
			err := json.Unmarshal(v, &transaction)
			if err != nil {
				return shim.Error(fmt.Sprintf("queryTransactionInfoList-反序列化出错: %s", err))
			}
			transactionList = append(transactionList, transaction)
		}
	}
	transactionListByte, err := json.Marshal(transactionList)
	if err != nil {
		return shim.Error(fmt.Sprintf("queryTransactionInfoList-序列化出错: %s", err))
	}
	return shim.Success(transactionListByte)
}



// QuerySellingList 查询销售(可查询所有，也可根据发起销售人查询)(发起的)(供卖家查询)
func queryServiceChargeTransactionList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var transactionList []ServiceChargeTransaction
	results, err := utils.GetStateByPartialCompositeKeys2(stub, ServiceChargeTransactionKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var transaction ServiceChargeTransaction
			err := json.Unmarshal(v, &transaction)
			if err != nil {
				return shim.Error(fmt.Sprintf("queryTransactionInfoList-反序列化出错: %s", err))
			}
			transactionList = append(transactionList, transaction)
		}
	}
	transactionListByte, err := json.Marshal(transactionList)
	if err != nil {
		return shim.Error(fmt.Sprintf("queryTransactionInfoList-序列化出错: %s", err))
	}
	return shim.Success(transactionListByte)
}

// QuerySellingList 查询销售(可查询所有，也可根据发起销售人查询)(发起的)(供卖家查询)
func querySuggestedPortfolioList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var transactionList []SuggestedPortfolio
	results, err := utils.GetStateByPartialCompositeKeys2(stub, SuggestedPortfoliokey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var transaction SuggestedPortfolio
			err := json.Unmarshal(v, &transaction)
			if err != nil {
				return shim.Error(fmt.Sprintf("querySuggestedPortfolioList-反序列化出错: %s", err))
			}
			transactionList = append(transactionList, transaction)
		}
	}
	transactionListByte, err := json.Marshal(transactionList)
	if err != nil {
		return shim.Error(fmt.Sprintf("queryTransactionInfoList-序列化出错: %s", err))
	}
	return shim.Success(transactionListByte)
}

func queryCommodityTypeList(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	fmt.Println("running the function queryCommodityTypeList()")
	// var tempCommodityList []Commodity
	var err error


	commodityListByte, err := json.Marshal(commodityList)
	if err != nil {
		return shim.Error(fmt.Sprintf("commodityList-序列化出错: %s", err))
	}
	return shim.Success(commodityListByte)
}




func set(stub shim.ChaincodeStubInterface, args []string)pb.Response {
	stus := []Student{
		{"lisi", 20, 99},
		{"lisi",20,98},
		{"lisi",21,100},
	}
	for i, _ := range(stus) {
		stu := stus[i]

		key, err := stub.CreateCompositeKey(stu.Name, []string{strconv.Itoa(stu.Age), strconv.Itoa(stu.Score)})
		if err != nil {
			fmt.Println(err)
			return shim.Error(err.Error())
		}

		bytes, err := json.Marshal(stu)
		if err != nil {
			fmt.Println(err)
			return shim.Error(err.Error())
		}
		stub.PutState(key, bytes)

	}
	return shim.Success(nil)
}


func  get(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var stuList []Student
	rs, err := stub.GetStateByPartialCompositeKey("lisi", args)
	if err != nil{
		fmt.Println(err)
		return  shim.Error(err.Error())
	}
	defer rs.Close()

	for rs.HasNext(){
		responseRange, err := rs.Next()

		if err != nil{
			fmt.Println(err)
		}
		stu := new(Student)
		err = json.Unmarshal(responseRange.Value, stu)
		if err != nil{
			fmt.Println(err)
		}
		stuList = append(stuList, *stu)
		fmt.Println(responseRange.Key, stu)
	}
	bytes, err := json.Marshal(stuList)

	return shim.Success(bytes)
}


func main() {
	// Start the chaincode and make it ready for futures requests
	err := shim.Start(new(BlockChainRealEstate))
	if err != nil {
		fmt.Printf("Error starting Heroes Service chaincode: %s", err)
	}
}

