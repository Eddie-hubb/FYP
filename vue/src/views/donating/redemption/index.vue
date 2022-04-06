<template>
  <div class="container">
    <user-info-bar />
    <div v-if="!loading" class="cards-container">
      <el-card v-for="(val,index) in transactionInfoList" :key="index" class="card-item">
        <div class="item">
          <span>{{ val.createTime }}</span>
        </div>
        <div class="item">
          <span>Amount Of Fee: {{ val.amountOfFee }}</span>
        </div>
        <div class="item">
          <span>Receiver: {{ val.receiver }}</span>
        </div>
        <div class="item">
          <span>Sender: {{ val.sender }}</span>
        </div>
      </el-card>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryRedemptionFeeTransactionList } from '@/api/transaction'
import UserInfoBar from '@/components/UserInfoBar'

export default {
  name: 'RedemptionFee',
  components: {
    UserInfoBar
  },
  data() {
    return {
      loading: true,
      donatingList: []
    }
  },
  computed: {
    ...mapGetters([
      'accountId',
      'userName',
      'balance'
    ])
  },
  created() {
    queryRedemptionFeeTransactionList({ accountID: this.accountId }).then(response => {
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
  .tag {
    float: left;
  }

  .item {
    font-size: 14px;
    margin-bottom: 18px;
    color: #999;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
  }
  .clearfix:after {
    clear: both
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
