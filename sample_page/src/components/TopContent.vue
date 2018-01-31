<template>
  <div id="head">
    <h1>{{ msg }}</h1>
    <hr />
  </div>
</template>

<script>
export default {
  name: 'TopContent',
  data () {
    return {
      msg: 'Welcome to My Page'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1{
  font-weight: normal;
  text-align: center;
}
</style>
