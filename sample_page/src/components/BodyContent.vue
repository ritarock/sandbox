<template>
  <div id="body">
    <div v-on:click="clickLink">{{link}}</div>
    <ul v-if="linklist">
      <li><a href="https://qiita.com/ritarock">qiita</a></li>
      <li><a href="https://github.com/ritarock">github</a></li>
      <li><a href="http://ritarock.hatenablog.com/">hatenablog</a></li>
    </ul>
    <hr />
    <div>{{who}}</div>
    <hr />
  </div>
</template>

<script>
export default {
  name: 'BodyContent',
  data: function(){
    return{
      link: 'LINKS ▶︎',
      who: 'WHO? ▶︎',
      linklist: false
    }
  },
  methods: {
    clickLink: function(){
      // this.message = this.message.split('').reverse().join('')
      // console.log(this.linklist)
      this.linklist = true
      this.link = 'LINKS ▼'
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
