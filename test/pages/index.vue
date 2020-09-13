<template>
  <div>
    <h1>Message: {{ message }}</h1>
    <span>
      <nuxt-link to="/table">Table Page</nuxt-link>
    </span>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  data() {
    return {
      message: 'Hello'
    }
  }
})
</script>