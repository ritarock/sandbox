<template>
  <div id="body">
    <div v-on:click="viewlist">{{message}}</div>
  </ul>
</div>
</template>

<script>
export default {
  name: 'BodyContent',
  data: function(){
    return{
      message: 'link'
    }
  },
  methods: {
    viewlist: function(){
      this.message = this.message.split('').reverse().join('')
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#links {
  color: black;
}
</style>
