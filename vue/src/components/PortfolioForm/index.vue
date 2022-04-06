<template>
  <div>
    <el-form ref="portfolioForm" :model="portfolioForm" label-width="100px" label-position="right">
      <el-form-item label="gold" prop="gold">
        <el-input-number v-model="portfolioForm.gold" :min="0" :step="50" />
      </el-form-item>
      <el-form-item label="silver" prop="silver">
        <el-input-number v-model="portfolioForm.silver" :min="0" :step="50" />
      </el-form-item>
      <el-form-item label="platinum" prop="platinum">
        <el-input-number v-model="portfolioForm.platinum" :min="0" :step="50" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('portfolioForm')">Create</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {

}
</script>

<style>

</style>
