import request from '@/utils/request'

export function queryPortfolioList(data) {
  return request({
    url: '/queryPortfolioList',
    method: 'post',
    data
  })
}

export function queryCommodityTemplate() {
  return request({
    url: '/queryCommodityTypeList',
    method: 'GET'
  })
}

export function createPortfolioInfo(data) {
  return request({
    url: '/createPortfolioInfo',
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

