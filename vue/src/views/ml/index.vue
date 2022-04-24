<template>
  <div class="container">
    <user-info-bar />
    <div class="form-container">
      <el-form ref="mlForm" :model="mlForm" :rules="rules" label-position="right" label-width="250px">
        <el-form-item label="Start and End date" prop="start_end_date">
          <el-date-picker
            v-model="mlForm.start_end_date"
            type="daterange"
            range-separator="to"
            start-placeholder="Start date"
            end-placeholder="End date"
            value-format="yyyy-MM-dd"
          />
        </el-form-item>
        <el-form-item label="Validation Start and End date" prop="validation_start_end_date">
          <el-date-picker
            v-model="mlForm.validation_start_end_date"
            type="daterange"
            range-separator="to"
            start-placeholder="Start date"
            end-placeholder="End date"
            value-format="yyyy-MM-dd"
          />
        </el-form-item>
        <el-form-item label="Asset" prop="asset">
          <el-checkbox-group v-model="mlForm.asset_list" :min="2">
            <el-checkbox label="Gold" />
            <el-checkbox label="Silver" />
            <el-checkbox label="Platinum" />
          </el-checkbox-group>
        </el-form-item>
        <el-form-item label="Target">
          <el-select v-model="mlForm.target">
            <el-option
              v-for="item in targetOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="Algorithm">
          <el-select v-model="mlForm.algorithm">
            <el-option
              v-for="item in algorithmOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-button v-loading="isLoading" type="primary" @click="submitForm('mlForm')">
            submit
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-dialog v-loading="isLoading" title="Suggested Weight" :visible.sync="showResult">
      <div class="item">
        <span>start date: {{ mlResult.start_date }}</span>
      </div>
      <div class="item">
        <span>end date: {{ mlResult.end_date }}</span>
      </div>
      <div class="item">
        <span>validation start date: {{ mlResult.validation_start_date }}</span>
      </div>
      <div class="item">
        <span>validation end date: {{ mlResult.validation_end_date }}</span>
      </div>
      <div v-for="(val, index) in mlResult.asset_list" :key="index" class="item">
        <el-tag>{{ val }}</el-tag>
        <span> {{ mlResult.suggested_weight[index+1] }}</span>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

import UserInfoBar from '@/components/UserInfoBar'
import axios from 'axios'

export default {
  name: 'MachineLearning',
  components: {
    UserInfoBar
  },
  data() {
    return {
      isFinish: false,
      isLoading: false,
      showResult: false,
      mlResult: {
        start_date: '',
        end_date: '',
        validation_start_date: ' ',
        validation_end_date: '',
        asset_list: []
      },
      pollingTimer: null,
      targetOptions: [
        {
          value: 'Minimal_Volatility',
          label: 'Minimal Volatility'
        }
      ],
      algorithmOptions: [
        {
          value: 'EA',
          label: 'EA'
        }
      ],
      mlForm: {
        start_end_date: [],
        end_date: '',
        validation_start_end_date: [],
        validation_end_date: '',
        asset_list: ['Gold', 'Silver'],
        constraint: {
          minimal_asset_num: 2,
          minimal_asset_weight: 500,
          maximal_asset_weight: 7000,
          account_id: this.accountId
        },
        target: 'Minimal_Volatility',
        algorithm: 'EA'
      },
      rules: {
        start_end_date: [
          { type: 'array', required: true, message: 'Please select date', trigger: 'change' }
        ],
        validation_start_end_date: [
          { type: 'array', required: true, message: 'Please select date', trigger: 'change' }
        ]
      }
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
  beforeDestroy() {
    this.pollingTimer && clearInterval(this.pollingTimer)
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          console.log('h', this.mlForm, this.mlForm.start_end_date)
          this.getRequestId({
            start_date: this.mlForm.start_end_date[0],
            end_date: this.mlForm.start_end_date[1],
            validation_start_date: this.mlForm.validation_start_end_date[0],
            validation_end_date: this.mlForm.validation_start_end_date[1],
            asset_list: this.mlForm.asset_list,
            constraint: {
              minimal_asset_num: 2,
              minimal_asset_weight: 500,
              maximal_asset_weight: 7000,
              account_id: this.accountId
            },
            target: this.mlForm.target,
            algorithm: this.mlForm.algorithm
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    getResult(request_id) {
      axios({
        method: 'GET',
        url: '/result',
        baseURL: 'http://localhost:8888',
        params: {
          request_id: request_id
        }
      }).then((res) => {
        const result = res.data
        console.log('res', result)
        if (result.status === 'finished') {
          this.isLoading = false
          this.isFinish = true
          this.showResult = true
          this.mlResult = result
          console.log('finished')
          clearInterval(this.pollingTimer)
        }
      }).catch(() => {
        this.isLoading = false
      })
    },
    getRequestId(data) {
      axios({
        method: 'post',
        url: '/machine-learning',
        baseURL: 'http://localhost:8888',
        data: data
      }).then((res) => {
        const result = res.data
        console.log('res', result)
        // this.request_id = result.request_id
        if (result.request_id != null) {
          this.pollingRequest(result.request_id)
          this.isLoading = true
        }
      })
    },
    pollingRequest(request_id) {
      const t = setInterval(() => {
        this.getResult(request_id)
      }, 2000)
      this.pollingTimer = t
    }
  }
}
</script>

<style>
.form-container{
    padding: 20px;
}
</style>
