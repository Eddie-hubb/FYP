<template>
  <el-alert
    type="success"
  >
    <div class="info-bar-container">
      <div class="left-box">
        <p>Account ID: {{ accountId }}</p>
        <p>Username: {{ userName }}</p>
        <p>Balance: {{ balance }}</p>
      </div>
      <div class="right-box">
        <p>Gold Share: {{ goldShare }}</p>
        <p>Sliver Share: {{ silverShare }}</p>
        <p>Platinum Share: {{ platinumShare }}</p>
      </div>
    </div>

  </el-alert>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  computed: {
    ...mapGetters([
      'accountId',
      'userName',
      'balance',
      'goldShare',
      'silverShare',
      'platinumShare'
    ])
  }
}
</script>

<style>
.info-bar-container {
    display: flex;
}

.left-box {
    width: 200px;
}

</style>
