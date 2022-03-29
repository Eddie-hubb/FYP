import request from '@/utils/request'

export function queryTransactionInfoList(data) {
  return request({
    url: '/queryTransactionInfoList',
    method: 'post',
    data
  })
}

export function queryMoneyTransactionList(data) {
    return request({
      url: '/queryMoneyTransactionList',
      method: 'post',
      data
    })
  }
export function queryCommodityTransactionList(data) {
  return request({
    url: '/queryCommodityTransactionList',
    method: 'post',
    data
  })
}


export function queryServiceChargeTransactionList(data) {
    return request({
      url: '/queryServiceChargeTransactionList',
      method: 'post',
      data
    })
  }

  
export function queryRedemptionFeeTransactionList(data) {
    return request({
      url: '/queryRedemptionFeeTransactionList',
      method: 'post',
      data
    })
  }

  
export function createPortfolioInfo(data) {
    return request({
      url: '/createPortfolioInfo',
      method: 'post',
      data
    })
  }

  
export function createTransactionInfo(data) {
    return request({
      url: '/createTransactionInfo',
      method: 'post',
      data
    })
  }

  
export function createSuggestedPortfolioInfo(data) {
    return request({
      url: '/createSuggestedPortfolioInfo',
      method: 'post',
      data
    })
  }

  
export function updateState(data) {
    return request({
      url: '/updateState',
      method: 'post',
      data
    })
  }

  
export function adjustPortfolio(data) {
    return request({
      url: '/adjustPortfolio',
      method: 'post',
      data
    })
  }

export function adjustNetWorth(data) {
    return request({
        url: '/adjustNetWorth',
        method: 'post',
        data
    })
}
  
