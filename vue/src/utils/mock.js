const accountId = '12346'

export const mockPortfolioList = [
  {
    accountID: accountId,
    portfolioID: '11111',
    createTime: '2022-04-02 10:00:00',
    shareList: [
      {
        name: 'gold',
        share: 100
      },
      {
        name: 'silver',
        share: 100
      },
      {
        name: 'platinum',
        share: 100
      }
    ]
  },
  {
    accountID: accountId,
    portfolioID: '11112',
    createTime: '2022-04-02 11:00:00',
    shareList: [
      {
        name: 'gold',
        share: 900
      },
      {
        name: 'silver',
        share: 0
      },
      {
        name: 'platinum',
        share: 101
      }
    ]
  }

]
