package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	bc "github.com/togettoyou/blockchain-real-estate/application/blockchain"
	"github.com/togettoyou/blockchain-real-estate/application/pkg/app"
	"net/http"
)

// type AccountIdBody struct {
// 	AccountId string `json:"accountId"`
// }

// type AccountRequestBody struct {
// 	Account []AccountIdBody `json:"account"`
// }



// TODO 建议分开写接口
// query by portfolioId | query by accountId
// bug: 当参数不符合时不返回任何数据
type PortfolioInfoRequestBody struct {
	AccountId string `json:"accountId"`
	PortfolioId string `json:"portfolioId"`
}

// type TransactionIdBody struct {
// }

type TransactionRequestBody struct {
	AccountId string `json:"accountId"`
	TransactionId string `json:"transactionId"`
}

// type MoneyTransactionIdBody struct {
// 	MoneyTransactionId string `json:"transactionId"`
// }

type MoneyTransactionRequestBody struct {
	AccountId string `json:"accountId"`
	MoneyTransactionId string `json:"transactionId"`
}

// type CommodityTransactionIdBody struct {
// 	CommodityTransactionId string `json:"commodityTransactionId"`
// }

type CommodityTransactionRequestBody struct {
	AccountId string `json:"accountId"`
	CommodityTransactionId string `json:"commodityTransactionId"`

}

// type ServiceChargeTransactionIdBody struct {
// 	ServiceChargeTransactionId string `json:"serviceChargeTransactionId"`
// }

type ServiceChargeTransactionRequestBody struct {
	AccountId string `json:"accountId"`
	ServiceChargeTransactionId string `json:"serviceChargeTransactionId"`

}

// type RedemptionFeeTransactionIdBody struct {
// 	RedemptionFeeTransactionId string `json:"redemptionFeeTransactionId"`
// }

type RedemptionFeeTransactionRequestBody struct {
	AccountId string `json:"accountId"`
	RedemptionFeeTransactionId string `json:"redemptionFeeTransactionId"`
}

type PortfolioRequestBody struct {
	AccountID 		string	`json:"accountID"`
	GoldShare		string	`json:"goldShare"`
	SilverShare		string	`json:"silverShare"`
	PlatinumShare	string	`json:"platinumShare"`
}

type TransactionInfoRequestBody struct {
	PortfolioID 	string	`json:"portfolioID"`//need to update
	BuyerID			string 	`json:"buyerID"`
	CommodityType   string   `json:"commodityType"`
	PurchaseShare   string	`json:"purchaseShare"`
	SellShare		string	`json:"sellShare"`	
}

type TransactionStateRequestBody struct {
	AccountId string `json:"accountId"`
	TransactionID	string	`json:"transactionID"`
	NewState	string `json:"newState"`	
}

type UpdatedNetWorthRequestBody struct {
	GoldNetWorth	string `json:"goldNetWorth"`
	SilverNetWorth	string `json:"silverNetWorth"`
	PlatinumNetWorth	string `json:"platinumNetWorth"`
}

type SuggestedPortfolioRequestBody struct {
	GoldPercentage	string `json:"goldPercentage"`
	SilverPercentage	string `json:"silverPercentage"`
	PlatinumPercentage	string `json:"platinumPercentage"`
	Time	string `json:"time"`
}

type AdjustedPortfolioRequestBody struct {
	PreviousPortfolioID 	string	`json:"previousPortfolioID"`//need to update
	AccountID 		string	`json:"accountID"`
	GoldShare		string	`json:"goldShare"`
	SilverShare		string	`json:"silverShare"`
	PlatinumShare	string	`json:"platinumShare"`
}

// type AccountIdBody struct {
// 	AccountId string `json:"accountId"`
// }

// type AccountRequestBody struct {
// 	Account []AccountIdBody `json:"account"`
// }



// @Summary 获取账户信息
// @Param account body AccountRequestBody true "account"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/queryAccountList [post]
// func QueryAccountList(c *gin.Context) {
// 	appG := app.Gin{C: c}
// 	body := new(AccountRequestBody)
// 	//解析Body参数
// 	if err := c.ShouldBind(body); err != nil {
// 		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
// 		return
// 	}
// 	var bodyBytes [][]byte
// 	for _, val := range body.Account {
// 		bodyBytes = append(bodyBytes, []byte(val.AccountId))
// 	}
// 	//调用智能合约
// 	resp, err := bc.ChannelQuery("queryAccountList", bodyBytes)
// 	if err != nil {
// 		appG.Response(http.StatusInternalServerError, "失败", err.Error())
// 		return
// 	}
// 	// 反序列化json
// 	var data []map[string]interface{}
// 	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
// 		appG.Response(http.StatusInternalServerError, "失败", err.Error())
// 		return
// 	}
// 	appG.Response(http.StatusOK, "成功", data)
// }



func QueryPortfolioList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(PortfolioInfoRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.AccountId != "" {
		bodyBytes = append(bodyBytes, []byte(body.AccountId))
	}
	if body.PortfolioId != "" {
		bodyBytes = append(bodyBytes, []byte(body.PortfolioId))
	}
	
	//调用智能合约
	resp, err := bc.ChannelQuery("queryPortfolioList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}


func QueryTransactionInfoList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(TransactionRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.AccountId != "" {
		bodyBytes = append(bodyBytes, []byte(body.AccountId))
	}	
	if body.TransactionId != "" {
		bodyBytes = append(bodyBytes, []byte(body.TransactionId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryTransactionInfoList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}


func QueryMoneyTransactionList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(MoneyTransactionRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.AccountId != "" {
		bodyBytes = append(bodyBytes, []byte(body.AccountId))
	}	
	if body.MoneyTransactionId != "" {
		bodyBytes = append(bodyBytes, []byte(body.MoneyTransactionId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryMoneyTransactionList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryCommodityTransactionList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(CommodityTransactionRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.AccountId != "" {
		bodyBytes = append(bodyBytes, []byte(body.AccountId))
	}	
	if body.CommodityTransactionId != "" {
		bodyBytes = append(bodyBytes, []byte(body.CommodityTransactionId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryCommodityTransactionList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryServiceChargeTransactionList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(ServiceChargeTransactionRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.AccountId != "" {
		bodyBytes = append(bodyBytes, []byte(body.AccountId))
	}	
	if body.ServiceChargeTransactionId != "" {
		bodyBytes = append(bodyBytes, []byte(body.ServiceChargeTransactionId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryServiceChargeTransactionList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryRedemptionFeeTransactionList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RedemptionFeeTransactionRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.AccountId != "" {
		bodyBytes = append(bodyBytes, []byte(body.AccountId))
	}	
	if body.RedemptionFeeTransactionId != "" {
		bodyBytes = append(bodyBytes, []byte(body.RedemptionFeeTransactionId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryRedemptionFeeTransactionList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func CreatePortfolioInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(PortfolioRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败1", fmt.Sprintf("参数出错123%s", err.Error()))
		return
	}
	// if body.ObjectOfSale == "" || body.Seller == "" {
	// 	appG.Response(http.StatusBadRequest, "失败", "ObjectOfSale销售对象和Seller发起销售人不能为空")
	// 	return
	// }
	// if body.GoldShare <= 0 || body.SilverShare <= 0 || body.PlatinumShare <= 0 {
	// 	appG.Response(http.StatusBadRequest, "fail", "illegal input")
	// 	return
	// }
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.AccountID))
	bodyBytes = append(bodyBytes, []byte(body.GoldShare))
	bodyBytes = append(bodyBytes, []byte(body.SilverShare))
	bodyBytes = append(bodyBytes, []byte(body.PlatinumShare))
	//调用智能合约
	resp, err := bc.ChannelExecute("createPortfolioInfo", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败2", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败3", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func CreateTransactionInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(TransactionInfoRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	// if body.ObjectOfSale == "" || body.Seller == "" {
	// 	appG.Response(http.StatusBadRequest, "失败", "ObjectOfSale销售对象和Seller发起销售人不能为空")
	// 	return
	// }
	// if body.PurchaseShare <= 0 || body.SellShare <= 0  {
	// 	appG.Response(http.StatusBadRequest, "fail", "illegal input")
	// 	return
	// }
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.PortfolioID))
	bodyBytes = append(bodyBytes, []byte(body.BuyerID))
	bodyBytes = append(bodyBytes, []byte(body.CommodityType))
	bodyBytes = append(bodyBytes, []byte(body.PurchaseShare))
	bodyBytes = append(bodyBytes, []byte(body.SellShare))

	//调用智能合约
	resp, err := bc.ChannelExecute("createTransactionInfo", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func CreateSuggestedPortfolioInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(SuggestedPortfolioRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	// if body.ObjectOfSale == "" || body.Seller == "" {
	// 	appG.Response(http.StatusBadRequest, "失败", "ObjectOfSale销售对象和Seller发起销售人不能为空")
	// 	return
	// }
	// if body.GoldPercentage <= 0 || body.SilverPercentage <= 0  || body.PlatinumPercentage <= 0 {
	// 	appG.Response(http.StatusBadRequest, "fail", "illegal input")
	// 	return
	// }
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.GoldPercentage))
	bodyBytes = append(bodyBytes, []byte(body.SilverPercentage))
	bodyBytes = append(bodyBytes, []byte(body.PlatinumPercentage))
	bodyBytes = append(bodyBytes, []byte(body.Time))

	//调用智能合约
	resp, err := bc.ChannelExecute("createSuggestedPortfolioInfo", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func UpdateState(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(TransactionStateRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	// if body.ObjectOfSale == "" || body.Seller == "" {
	// 	appG.Response(http.StatusBadRequest, "失败", "ObjectOfSale销售对象和Seller发起销售人不能为空")
	// 	return
	// }
	// if body.PurchaseShare <= 0 || body.SellShare <= 0  {
	// 	appG.Response(http.StatusBadRequest, "fail", "illegal input")
	// 	return
	// }
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.AccountId))
	bodyBytes = append(bodyBytes, []byte(body.TransactionID))
	bodyBytes = append(bodyBytes, []byte(body.NewState))

	//调用智能合约
	resp, err := bc.ChannelExecute("updateState", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func AdjustPortfolio(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(AdjustedPortfolioRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	// if body.ObjectOfSale == "" || body.Seller == "" {
	// 	appG.Response(http.StatusBadRequest, "失败", "ObjectOfSale销售对象和Seller发起销售人不能为空")
	// 	return
	// }
	// if body.PurchaseShare <= 0 || body.SellShare <= 0  {
	// 	appG.Response(http.StatusBadRequest, "fail", "illegal input")
	// 	return
	// }
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.PreviousPortfolioID))
	bodyBytes = append(bodyBytes, []byte(body.AccountID))
	bodyBytes = append(bodyBytes, []byte(body.GoldShare))
	bodyBytes = append(bodyBytes, []byte(body.SilverShare))
	bodyBytes = append(bodyBytes, []byte(body.PlatinumShare))


	//调用智能合约
	resp, err := bc.ChannelExecute("adjustPortfolio", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func AdjustNetWorth(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(UpdatedNetWorthRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	// if body.ObjectOfSale == "" || body.Seller == "" {
	// 	appG.Response(http.StatusBadRequest, "失败", "ObjectOfSale销售对象和Seller发起销售人不能为空")
	// 	return
	// }
	// if body.PurchaseShare <= 0 || body.SellShare <= 0  {
	// 	appG.Response(http.StatusBadRequest, "fail", "illegal input")
	// 	return
	// }
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.GoldNetWorth))
	bodyBytes = append(bodyBytes, []byte(body.SilverNetWorth))
	bodyBytes = append(bodyBytes, []byte(body.PlatinumNetWorth))

	//调用智能合约
	resp, err := bc.ChannelExecute("adjustNetWorth", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryCommodityTypeList(c *gin.Context) {
	appG := app.Gin{C: c}
	var bodyBytes [][]byte
	resp, err := bc.ChannelExecute("queryCommodityTypeList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败1", err.Error())
		return
	}
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败2", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功3", data)
}
