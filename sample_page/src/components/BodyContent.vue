<template>
  <div id="body">
    <div v-on:click="viewlist">{{message}}</div>
    <ul v-if="linklist">
      <li>qiita</li>
      <li>github</li>
      <li>hatena</li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'BodyContent',
  data: function(){
    return{
      message: 'link',
      linklist: false
    }
  },
  methods: {
    viewlist: function(){
      // this.message = this.message.split('').reverse().join('')
      console.log(this.linklist)
      this.linklist = true
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
