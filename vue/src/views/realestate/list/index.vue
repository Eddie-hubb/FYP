<template>
  <div class="container">
    <user-info-bar />

  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAccountList } from '@/api/account'
import { createSelling } from '@/api/selling'
import { createDonating } from '@/api/donating'
import UserInfoBar from '@/components/UserInfoBar'

export default {
  name: 'Information',
  components: {
    UserInfoBar
  },
  data() {
    var checkArea = (rule, value, callback) => {
      if (value <= 0) {
        callback(new Error('必须大于0'))
      } else {
        callback()
      }
    }
    return {
      loading: true,
      loadingDialog: false,
      realEstateList: [],
      dialogCreateSelling: false,
      dialogCreateDonating: false,
      realForm: {
        price: 0,
        salePeriod: 0
      },
      rules: {
        price: [
          { validator: checkArea, trigger: 'blur' }
        ],
        salePeriod: [
          { validator: checkArea, trigger: 'blur' }
        ]
      },
      DonatingForm: {
        proprietor: ''
      },
      rulesDonating: {
        proprietor: [
          { required: true, message: '请选择业主', trigger: 'change' }
        ]
      },
      accountList: [],
      valItem: {}
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
  },
  methods: {
    openDialog(item) {
      this.dialogCreateSelling = true
      this.valItem = item
    },
    openDonatingDialog(item) {
      this.dialogCreateDonating = true
      this.valItem = item
      queryAccountList().then(response => {
        if (response !== null) {
          // 过滤掉管理员和当前用户
          this.accountList = response.filter(item =>
            item.userName !== '管理员' && item.accountId !== this.accountId
          )
        }
      })
    },
    createSelling(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('是否立即出售?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'success'
          }).then(() => {
            this.loadingDialog = true
            createSelling({
              objectOfSale: this.valItem.realEstateId,
              seller: this.valItem.proprietor,
              price: this.realForm.price,
              salePeriod: this.realForm.salePeriod
            }).then(response => {
              this.loadingDialog = false
              this.dialogCreateSelling = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: '出售成功!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: '出售失败!'
                })
              }
              setTimeout(() => {
                window.location.reload()
              }, 1000)
            }).catch(_ => {
              this.loadingDialog = false
              this.dialogCreateSelling = false
            })
          }).catch(() => {
            this.loadingDialog = false
            this.dialogCreateSelling = false
            this.$message({
              type: 'info',
              message: '已取消出售'
            })
          })
        } else {
          return false
        }
      })
    },
    createDonating(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('是否立即捐赠?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'success'
          }).then(() => {
            this.loadingDialog = true
            createDonating({
              objectOfDonating: this.valItem.realEstateId,
              donor: this.valItem.proprietor,
              grantee: this.DonatingForm.proprietor
            }).then(response => {
              this.loadingDialog = false
              this.dialogCreateDonating = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: '捐赠成功!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: '捐赠失败!'
                })
              }
              setTimeout(() => {
                window.location.reload()
              }, 1000)
            }).catch(_ => {
              this.loadingDialog = false
              this.dialogCreateDonating = false
            })
          }).catch(() => {
            this.loadingDialog = false
            this.dialogCreateDonating = false
            this.$message({
              type: 'info',
              message: '已取消捐赠'
            })
          })
        } else {
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    selectGet(accountId) {
      this.DonatingForm.proprietor = accountId
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

  .realEstate-card {
    width: 280px;
    height: 340px;
    margin: 18px;
  }
</style>
