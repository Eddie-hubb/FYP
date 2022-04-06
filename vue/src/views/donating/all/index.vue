<template>
  <div class="container">
    <user-info-bar />
    <div v-if="!loading" class="cards-container">
      <el-card v-for="(val,index) in transactionInfoList" :key="index" class="card-item">
        <div class="item">
          <span>{{ val.createTime }}</span>
        </div>
        <div class="item">
          <span>Commodity: {{ val.commodityType.commodityName }}</span>
        </div>
        <div class="item">
          <span>Commodity Net Worth: {{ val.commodityType.commodityNetWorth }}</span>
        </div>
        <div v-if="val.purchaseShare !== 0">
          <div class="item">
            <span>Purchase Share: {{ val.purchaseShare }}</span>
          </div>
          <div class="item">
            <span>Purchase Amount: {{ val.purchaseAmount }}</span>
          </div>
        </div>
        <div v-if="val.sellShare !== 0">
          <div class="item">
            <span>Sell Share: {{ val.sellShare }}</span>
          </div>
          <div class="item">
            <span>Sell Amount: {{ val.sellAmount }}</span>
          </div>
        </div>
        <div class="item">
          <span>Redemption Fee: {{ val.redemptionFee }}</span>
        </div>
        <div class="item">
          <span>Service Charge: {{ val.serviceCharge }}</span>
        </div>
        <div class="item">
          <span>State: </span>
          <el-tag v-if="val.transactionStateType.transactionStateID === 1">{{ val.transactionStateType.transactionStateName }}</el-tag>
          <el-tag v-if="val.transactionStateType.transactionStateID === 2" type="success">{{ val.transactionStateType.transactionStateName }}</el-tag>
        </div>
      </el-card>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryTransactionInfoList } from '@/api/transaction'
import UserInfoBar from '@/components/UserInfoBar'

export default {
  name: 'Information',
  components: {
    UserInfoBar
  },
  data() {
    return {
      loading: true,
      transactionInfoList: []
    }
  },
  computed: {
    ...mapGetters([
      'accountId',
      'roles',
      'userName',
      'balance'
    ])
  },
  created() {
    queryTransactionInfoList({
      accountID: this.accountId
    }).then(response => {
      if (response !== null) {
        this.transactionInfoList = response
      }
      this.loading = false
    }).catch(_ => {
      this.loading = false
    })
  },
  methods: {

  }
}

</script>

<style>

  .container{
    width: 100%;
    text-align: center;
    min-height: 100%;
    overflow: hidden;
  }

  .item {
    font-size: 14px;
    margin-bottom: 18px;
    color: #999;
  }

  .cards-container {
    display: flex;
    gap: 20px;
    flex-wrap: wrap;
  }

  .card-item {
    width: 350px;
    padding: 10px auto;
    margin: 18px;
  }
</style>
