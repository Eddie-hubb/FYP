<template>
  <div class="container">
    <user-info-bar />
    <div class="create-form">
      <el-form ref="portfolioForm" :model="portfolioForm" label-width="100px" label-position="right">
        <el-form-item label="gold(Troy Ounce)" prop="gold">
          <el-input-number v-model="portfolioForm.gold" :min="0" :step="50" />
        </el-form-item>
        <el-form-item label="silver(Troy Ounce)" prop="silver">
          <el-input-number v-model="portfolioForm.silver" :min="0" :step="50" />
        </el-form-item>
        <el-form-item label="platinum(Troy Ounce)" prop="platinum">
          <el-input-number v-model="portfolioForm.platinum" :min="0" :step="50" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('portfolioForm')">Create</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { createPortfolioInfo } from '@/api/portfolio'
import UserInfoBar from '@/components/UserInfoBar'

export default {
  name: 'BuySelling',
  components: {
    UserInfoBar
  },
  data() {
    return {
      loading: false,
      portfolioForm: {
        gold: 0,
        silver: 0,
        platinum: 0
      }
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
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (this.portfolioForm.gold !== 0 ||
        this.portfolioForm.silver !== 0 ||
        this.portfolioForm.platinum !== 0) {
          this.$confirm('Adjuct portfolio now?', 'Tip', {
            confirmButtonText: 'confirm',
            cancelButtonText: 'cancel',
            type: 'success'
          }).then(() => {
            this.loading = true
            createPortfolioInfo({
              accountID: this.accountId,
              goldShare: this.portfolioForm.gold.toString(),
              silverShare: this.portfolioForm.silver.toString(),
              platinumShare: this.portfolioForm.platinum.toString()
            }).then(response => {
              this.loading = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: 'Success!'
                })
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

  .buy-card {
    width: 280px;
    height: 430px;
    margin: 18px;
  }

  .create-form {
    width: 400px;
  }
</style>
