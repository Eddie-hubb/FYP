<template>
  <div class="container">
    <user-info-bar />
    <div class="card-container">
      <el-card v-for="(val,index) in portfolioList" :key="index" class="all-card">
        <div slot="header" class="clearfix">
          <span style="font-size: 12px">{{ val.portfolioID }}</span>
        </div>
        <div class="item">
          <span>{{ val.createTime }}</span>
        </div>
        <div class="item">
          <span>state: </span>
          <el-tag v-if="val.portfolioState.PortfolioStateTypeId == 2" type="error">{{ val.portfolioState.PortfolioStateTypeName }}</el-tag>
          <el-tag v-if="val.portfolioState.PortfolioStateTypeId == 1" type="success">{{ val.portfolioState.PortfolioStateTypeName }}</el-tag>
        </div>
        <div class="item">
          <span>Gold: {{ val.goldShare }} </span>
        </div>
        <div class="item">
          <span>Sliver: {{ val.silverShare }} </span>
        </div>
        <div class="item">
          <span>Platinum: {{ val.platinumShare }} </span>
        </div>

        <el-button v-if="val.portfolioState.PortfolioStateTypeId == 1" type="primary" round @click="adjustPortfolio(val)">Adjust Portfolio</el-button>
      </el-card>
    </div>

    <!-- dialog -->
    <el-dialog v-loading="adjustLoading" title="Adjust PortfolioList" :visible.sync="isAdjust">
      <el-form :model="adjustForm" label-width="100px" :label-position="right">
        <el-form-item label="gold">
          <el-input-number v-model="adjustForm.goldShare" :min="0" :step="50" />
        </el-form-item>
        <el-form-item label="silver">
          <el-input-number v-model="adjustForm.silverShare" :min="0" :step="50" />
        </el-form-item>
        <el-form-item label="platinum">
          <el-input-number v-model="adjustForm.platinumShare" :min="0" :step="50" />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="submitForm('adjustForm')">
            submit
          </el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>

</template>

<script>
import { mapGetters } from 'vuex'
import { queryPortfolioList, adjustPortfolio } from '@/api/portfolio'
import UserInfoBar from '@/components/UserInfoBar'

// import { mockPortfolioList } from '@/utils/mock'
export default {
  name: 'AllSelling',
  components: {
    UserInfoBar
  },
  data() {
    return {
      loading: true,
      adjustLoading: false,
      portfolioList: [],
      adjustForm: {
        accountID: '',
        createTime: '',
        goldShare: 0,
        platinumShare: 0,
        portfolioID: '',
        silverShare: 0
      },
      isAdjust: false
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
    this.getList()
  },
  methods: {
    getList() {
      this.loading = true
      queryPortfolioList({
        accountID: this.accountId
      }).then(res => {
        console.log(res)
        if (res !== null) {
          this.portfolioList = res
        }
        this.loading = false
      })
    },
    adjustPortfolio(val) {
      this.adjustForm = val
      this.isAdjust = true
    },
    submitForm(a) {
      if (this.adjustForm.goldShare !== 0 ||
        this.adjustForm.silverShare !== 0 ||
        this.adjustForm.platinumShare !== 0) {
        this.$confirm('Adjuct portfolio now?', 'Tip', {
          confirmButtonText: 'confirm',
          cancelButtonText: 'cancel',
          type: 'success'
        }).then(() => {
          this.adjustLoading = true
          adjustPortfolio({
            accountID: this.accountId,
            previousPortfolioID: this.adjustForm.portfolioID,
            goldShare: this.adjustForm.goldShare.toString(),
            silverShare: this.adjustForm.silverShare.toString(),
            platinumShare: this.adjustForm.platinumShare.toString()
          }).then(response => {
            this.loading = false
            if (response !== null) {
              this.$message({
                type: 'success',
                message: 'Success!'
              })
              this.isAdjust = false
              this.getList()
            } else {
              this.$message({
                type: 'error',
                message: 'Fail!'
              })
            }
          }).catch(_ => {
            this.loading = false
          })
        }).catch(() => {
          this.loading = false
          this.$message({
            type: 'info',
            message: 'Canceled'
          })
        })
      } else {
        this.$message({
          type: 'error',
          message: 'Empty Portfolio!'
        })
      }
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

  .all-card {
    width: 450px;
    padding: 10px auto;
    margin: 18px;
  }

  .card-container {
    width: 100%;
    display: flex;
    flex-wrap: wrap;
    gap: 20px;
  }
</style>
