<template>
  <div id="body">
    <div v-on:click="clickLink">{{link}}</div>
    <ul v-if="listLink">
      <li><a href="https://qiita.com/ritarock">qiita</a></li>
      <li><a href="https://github.com/ritarock">github</a></li>
      <li><a href="http://ritarock.hatenablog.com/">hatenablog</a></li>
    </ul>
    <hr />

    <div v-on:click="clickWho">{{who}}</div>
    <ul v-if="linkWho" id="self">
      <li>NAME: ritarock</li>
      <li>MAIL: wkryz1204@gmail.com</li>
    </ul>
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
      listLink: false,
      linkWho: false
    }
  },
  methods: {
    clickLink: function(){
      if (this.listLink) {
        this.link = 'LINKS ▶︎'
        this.listLink = false
      } else {
        this.link = 'LINKS ▼'
        this.listLink = true
      }
    },
    clickWho: function(){
      if (this.linkWho) {
        this.who = 'WHO? ▶︎'
        this. linkWho = false
      } else {
        this.who = 'WHO? ▼'
        this.linkWho = true
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#links {
  color: black;
}
#self {
  list-style:none;
}
</style>
