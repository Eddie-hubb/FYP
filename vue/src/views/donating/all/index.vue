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
        <el-button v-if="isManager&&val.transactionStateType.transactionStateID !== 2" type="primary" @click="updateStateToSuccess(val.transactionID, val.buyerID)">
          update state
        </el-button>
      </el-card>
    </div>
    <div v-else>
      <i class="el-icon-loading" />
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryTransactionInfoList, updateState } from '@/api/transaction'
import UserInfoBar from '@/components/UserInfoBar'

export default {
  name: 'Information',
  components: {
    UserInfoBar
  },
  data() {
    return {
      loading: true,
      isManager: false,
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
    if (this.roles[0] === 'manager') {
      this.isManager = true
    }
    this.loadInfoList()
  },
  methods: {
    updateStateToSuccess(id, buyerID) {
      this.$confirm('Confirm to update state to success?', 'Tip', {
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel',
        type: 'info'
      }).then(() => {
        this.loading = true
        updateState({
          accountId: buyerID,
          transactionID: id,
          newState: '2'
        }).then((res) => {
          this.loadInfoList()
        }).catch(() => {
          this.loading = false
        })
      })
    },
    loadInfoList() {
      queryTransactionInfoList({
        accountID: this.roles[0] === 'manager' ? '' : this.accountId
      }).then(response => {
        if (response !== null) {
          this.transactionInfoList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    }
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
